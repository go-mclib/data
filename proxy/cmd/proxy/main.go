package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// CapturedPacket stores packet data for test generation
type CapturedPacket struct {
	Direction string `json:"direction"` // "c2s" or "s2c"
	State     string `json:"state"`
	PacketID  string `json:"packet_id"` // hex format "0x00"
	WireData  string `json:"wire_data"` // hex encoded packet data (excludes length, compression, id)
}

// PacketFilter controls which packets are captured
type PacketFilter struct {
	States    map[string]bool // nil means all states
	PacketIDs map[int]bool    // nil means all packet IDs
}

func (f *PacketFilter) Match(state string, packetID int) bool {
	if f.States != nil && !f.States[state] {
		return false
	}
	if f.PacketIDs != nil && !f.PacketIDs[packetID] {
		return false
	}
	return true
}

// PacketCapture manages captured packets for a session
type PacketCapture struct {
	mu      sync.Mutex
	packets []CapturedPacket
	dir     string
	filter  *PacketFilter
}

func NewPacketCapture(outputDir string, filter *PacketFilter) *PacketCapture {
	return &PacketCapture{
		packets: make([]CapturedPacket, 0),
		dir:     outputDir,
		filter:  filter,
	}
}

func (pc *PacketCapture) Add(pkt CapturedPacket, state string, packetID int, wire *jp.WirePacket) {
	if !pc.filter.Match(state, packetID) {
		return
	}
	pc.mu.Lock()
	defer pc.mu.Unlock()

	// save only packet data (excludes length, compression header, and packet ID)
	pkt.WireData = hex.EncodeToString(wire.Data)
	pc.packets = append(pc.packets, pkt)
}

func (pc *PacketCapture) Save(filename string) error {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	if len(pc.packets) == 0 {
		return nil // don't save empty captures
	}

	if err := os.MkdirAll(pc.dir, 0755); err != nil {
		return err
	}

	path := filepath.Join(pc.dir, filename)
	data, err := json.MarshalIndent(pc.packets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (pc *PacketCapture) Count() int {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	return len(pc.packets)
}

// ProxySession handles a single client<->server connection pair
type ProxySession struct {
	clientConn *jp.Conn
	serverConn *jp.Conn
	capture    *PacketCapture
	verbose    bool
	logger     *log.Logger

	// protocol state tracking per direction (transitions happen independently)
	mu                   sync.RWMutex
	c2sState             jp.State // state for client-to-server packets
	s2cState             jp.State // state for server-to-client packets
	compressionThreshold int

	// compressionReady is closed when compression settings are finalized
	// (after S2C LoginFinished). C2S goroutine waits on this before reading
	// LoginAcknowledged to avoid race condition with compression format.
	compressionReady     chan struct{}
	compressionReadyOnce sync.Once
}

func NewProxySession(clientConn, serverConn net.Conn, capture *PacketCapture, verbose bool) *ProxySession {
	return &ProxySession{
		clientConn:           jp.NewConn(clientConn),
		serverConn:           jp.NewConn(serverConn),
		capture:              capture,
		verbose:              verbose,
		logger:               log.New(os.Stdout, "[proxy] ", log.LstdFlags),
		c2sState:             jp.StateHandshake,
		s2cState:             jp.StateHandshake,
		compressionThreshold: -1,
		compressionReady:     make(chan struct{}),
	}
}

// signalCompressionReady signals that compression settings are finalized
func (s *ProxySession) signalCompressionReady() {
	s.compressionReadyOnce.Do(func() {
		close(s.compressionReady)
	})
}

// waitCompressionReady waits until compression settings are finalized
func (s *ProxySession) waitCompressionReady() {
	<-s.compressionReady
}

func (s *ProxySession) setState(direction string, state jp.State) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if direction == "c2s" {
		s.c2sState = state
	} else {
		s.s2cState = state
	}
}

func (s *ProxySession) setBothStates(state jp.State) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.c2sState = state
	s.s2cState = state
}

func (s *ProxySession) getState(direction string) jp.State {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if direction == "c2s" {
		return s.c2sState
	}
	return s.s2cState
}

func (s *ProxySession) setCompression(threshold int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.compressionThreshold = threshold
}

func (s *ProxySession) getCompression() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.compressionThreshold
}

func stateToString(state jp.State) string {
	switch state {
	case jp.StateHandshake:
		return "handshake"
	case jp.StateStatus:
		return "status"
	case jp.StateLogin:
		return "login"
	case jp.StateConfiguration:
		return "configuration"
	case jp.StatePlay:
		return "play"
	default:
		return "unknown"
	}
}

// Run starts bidirectional packet forwarding
func (s *ProxySession) Run() {
	var wg sync.WaitGroup
	wg.Add(2)

	// client -> server
	go func() {
		defer wg.Done()
		s.forward(s.clientConn, s.serverConn, "c2s")
	}()

	// server -> client
	go func() {
		defer wg.Done()
		s.forward(s.serverConn, s.clientConn, "s2c")
	}()

	wg.Wait()
}

func (s *ProxySession) forward(src, dst *jp.Conn, direction string) {
	defer func() {
		_ = src.Close()
		_ = dst.Close()
	}()

	// Track if we need to wait for compression sync (C2S only)
	waitForCompression := false

	for {
		// C2S: wait for compression settings before reading LoginAcknowledged
		if waitForCompression {
			s.waitCompressionReady()
			waitForCompression = false
		}

		compression := s.getCompression()
		wire, err := jp.ReadWirePacketFrom(src, compression)
		if err != nil {
			if s.verbose {
				s.logger.Printf("%s: read error: %v", direction, err)
			}
			return
		}

		state := s.getState(direction)
		packetID := int(wire.PacketID)
		if s.verbose {
			s.logger.Printf("%s: state=%s id=0x%02X len=%d",
				direction, stateToString(state), packetID, len(wire.Data))
		}

		// capture packet (filter applied inside Add)
		s.capture.Add(CapturedPacket{
			Direction: direction,
			State:     stateToString(state),
			PacketID:  fmt.Sprintf("0x%02X", packetID),
		}, stateToString(state), packetID, wire)

		// handle state transitions based on packet type
		s.handleStateTransition(wire, direction)

		// C2S: after Login Hello, wait for compression before reading next packet
		// This ensures compression is set before reading LoginAcknowledged
		if direction == "c2s" && state == jp.StateLogin && packetID == 0x00 {
			waitForCompression = true
		}

		// S2C: after LoginFinished, signal that compression is ready
		// (compression was set by LoginCompression, or not used)
		if direction == "s2c" && state == jp.StateLogin && packetID == 0x02 {
			s.signalCompressionReady()
		}

		// forward packet
		if err := wire.WriteTo(dst, compression); err != nil {
			if s.verbose {
				s.logger.Printf("%s: write error: %v", direction, err)
			}
			return
		}
	}
}

// handleStateTransition updates protocol state based on terminal packets.
// State is tracked per direction since transitions happen independently.
func (s *ProxySession) handleStateTransition(wire *jp.WirePacket, direction string) {
	state := s.getState(direction)
	packetID := int(wire.PacketID)

	switch state {
	case jp.StateHandshake:
		// C2S Handshake (0x00) contains intent field - both directions transition together
		if direction == "c2s" && packetID == 0x00 {
			buf := ns.NewReader(wire.Data)
			_, _ = buf.ReadVarInt()    // protocol version
			_, _ = buf.ReadString(255) // server address
			_, _ = buf.ReadUint16()    // server port
			intent, err := buf.ReadVarInt()
			if err == nil {
				switch intent {
				case 1:
					s.setBothStates(jp.StateStatus)
					if s.verbose {
						s.logger.Printf("state -> status")
					}
				case 2:
					s.setBothStates(jp.StateLogin)
					if s.verbose {
						s.logger.Printf("state -> login")
					}
				}
			}
		}

	case jp.StateLogin:
		// S2C LoginCompression (0x03): set compression threshold
		if direction == "s2c" && packetID == 0x03 {
			buf := ns.NewReader(wire.Data)
			threshold, err := buf.ReadVarInt()
			if err == nil {
				s.setCompression(int(threshold))
				if s.verbose {
					s.logger.Printf("compression threshold -> %d", threshold)
				}
			}
		}
		// S2C LoginFinished (0x02): server transitions to configuration
		if direction == "s2c" && packetID == 0x02 {
			s.setState("s2c", jp.StateConfiguration)
			if s.verbose {
				s.logger.Printf("s2c state -> configuration")
			}
		}
		// C2S LoginAcknowledged (0x03): client transitions to configuration
		if direction == "c2s" && packetID == 0x03 {
			s.setState("c2s", jp.StateConfiguration)
			if s.verbose {
				s.logger.Printf("c2s state -> configuration")
			}
		}

	case jp.StateConfiguration:
		// S2C FinishConfiguration (0x03): server transitions to play
		if direction == "s2c" && packetID == 0x03 {
			s.setState("s2c", jp.StatePlay)
			if s.verbose {
				s.logger.Printf("s2c state -> play")
			}
		}
		// C2S FinishConfiguration (0x03): client transitions to play
		if direction == "c2s" && packetID == 0x03 {
			s.setState("c2s", jp.StatePlay)
			if s.verbose {
				s.logger.Printf("c2s state -> play")
			}
		}

	case jp.StatePlay:
		// S2C StartConfiguration (0x74): server transitions to configuration
		if direction == "s2c" && packetID == 0x74 {
			s.setState("s2c", jp.StateConfiguration)
			if s.verbose {
				s.logger.Printf("s2c state -> configuration (from play)")
			}
		}
		// C2S ConfigurationAcknowledged (0x0F): client transitions to configuration
		if direction == "c2s" && packetID == 0x0F {
			s.setState("c2s", jp.StateConfiguration)
			if s.verbose {
				s.logger.Printf("c2s state -> configuration (from play)")
			}
		}
	}
}

func parseStates(s string) map[string]bool {
	if s == "" {
		return nil
	}
	states := make(map[string]bool)
	for _, state := range strings.Split(s, ",") {
		state = strings.TrimSpace(strings.ToLower(state))
		if state != "" {
			states[state] = true
		}
	}
	if len(states) == 0 {
		return nil
	}
	return states
}

func parsePacketIDs(s string) map[int]bool {
	if s == "" {
		return nil
	}
	ids := make(map[int]bool)
	for _, idStr := range strings.Split(s, ",") {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		// handle hex format (0x00) or decimal
		var id int64
		var err error
		if strings.HasPrefix(idStr, "0x") || strings.HasPrefix(idStr, "0X") {
			id, err = strconv.ParseInt(idStr[2:], 16, 32)
		} else {
			id, err = strconv.ParseInt(idStr, 10, 32)
		}
		if err == nil {
			ids[int(id)] = true
		}
	}
	if len(ids) == 0 {
		return nil
	}
	return ids
}

func main() {
	var (
		listenPort  = flag.Int("port", 25565, "port to listen on")
		targetAddr  = flag.String("target", "", "target server address (host:port)")
		outputDir   = flag.String("output", "captures", "output directory for captured packets")
		verbose     = flag.Bool("verbose", false, "enable verbose logging")
		stateFilter = flag.String("state", "", "comma-separated states to capture (e.g., login,play)")
		idFilter    = flag.String("packetId", "", "comma-separated packet IDs to capture (e.g., 0x00,0x01)")
	)
	flag.Parse()

	if *targetAddr == "" {
		log.Fatal("target server address is required (-target)")
	}

	filter := &PacketFilter{
		States:    parseStates(*stateFilter),
		PacketIDs: parsePacketIDs(*idFilter),
	}

	// log filter settings
	if filter.States != nil {
		states := make([]string, 0, len(filter.States))
		for s := range filter.States {
			states = append(states, s)
		}
		log.Printf("filtering states: %s", strings.Join(states, ", "))
	}
	if filter.PacketIDs != nil {
		ids := make([]string, 0, len(filter.PacketIDs))
		for id := range filter.PacketIDs {
			ids = append(ids, fmt.Sprintf("0x%02X", id))
		}
		log.Printf("filtering packet IDs: %s", strings.Join(ids, ", "))
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	log.Printf("listening on :%d, proxying to %s", *listenPort, *targetAddr)
	log.Printf("packet captures will be saved to %s/", *outputDir)

	// single capture for entire proxy run
	capture := NewPacketCapture(*outputDir, filter)
	startTime := time.Now()

	// save capture on exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("shutting down...")
		saveCapture(capture, startTime)
		os.Exit(0)
	}()

	connNum := 0
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		connNum++
		log.Printf("connection %d: client connected from %s", connNum, clientConn.RemoteAddr())

		go func(num int, client net.Conn) {
			defer client.Close()

			// connect to target server
			serverConn, err := net.Dial("tcp", *targetAddr)
			if err != nil {
				log.Printf("connection %d: failed to connect to server: %v", num, err)
				return
			}
			defer serverConn.Close()

			log.Printf("connection %d: connected to server %s", num, *targetAddr)

			session := NewProxySession(client, serverConn, capture, *verbose)
			session.Run()

			log.Printf("connection %d: closed (%d total packets captured)", num, capture.Count())
		}(connNum, clientConn)
	}
}

func saveCapture(capture *PacketCapture, startTime time.Time) {
	if capture.Count() > 0 {
		filename := fmt.Sprintf("session_%s.json", startTime.Format("20060102_150405"))
		if err := capture.Save(filename); err != nil {
			log.Printf("failed to save capture: %v", err)
		} else {
			log.Printf("saved %d packets to %s", capture.Count(), filename)
		}
	} else {
		log.Println("no packets matched filter, nothing saved")
	}
}

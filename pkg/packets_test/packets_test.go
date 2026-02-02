package packets_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/go-mclib/data/pkg/packets"
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// PacketFixture represents a captured packet for testing
type PacketFixture struct {
	Direction string `json:"direction"`
	State     string `json:"state"`
	PacketID  string `json:"packet_id"` // hex format "0x00"
	RawData   string `json:"raw_data"`  // hex encoded
}

func parsePacketID(s string) (int, error) {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		id, err := strconv.ParseInt(s[2:], 16, 32)
		return int(id), err
	}
	id, err := strconv.ParseInt(s, 10, 32)
	return int(id), err
}

// LoadFixtures loads packet fixtures from a JSON capture file
func LoadFixtures(t *testing.T, filename string) []PacketFixture {
	t.Helper()
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		t.Fatalf("failed to load fixtures: %v", err)
	}

	var fixtures []PacketFixture
	if err := json.Unmarshal(data, &fixtures); err != nil {
		t.Fatalf("failed to parse fixtures: %v", err)
	}
	return fixtures
}

// testPacketRoundTrip tests that a packet can be read and written identically
func testPacketRoundTrip(t *testing.T, p jp.Packet, rawData []byte) {
	t.Helper()

	buf := ns.NewReader(rawData)
	if err := p.Read(buf); err != nil {
		t.Fatalf("Read failed: %v", err)
	}

	outBuf := ns.NewWriter()
	if err := p.Write(outBuf); err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	written := outBuf.Bytes()
	if !bytes.Equal(rawData, written) {
		t.Errorf("round-trip mismatch:\n  input:  %s\n  output: %s",
			hex.EncodeToString(rawData),
			hex.EncodeToString(written))
	}
}

// TestPacketsFromCapture runs round-trip tests on all captured packets
func TestPacketsFromCapture(t *testing.T) {
	fixtureDir := "fixtures"
	entries, err := os.ReadDir(fixtureDir)
	if err != nil {
		if os.IsNotExist(err) {
			t.Skip("no fixtures directory found - run proxy to capture packets first")
		}
		t.Fatalf("failed to read fixtures dir: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			fixtures := LoadFixtures(t, filepath.Join(fixtureDir, entry.Name()))
			if fixtures == nil {
				t.Skip("no fixtures loaded")
			}

			for i, f := range fixtures {
				packetID, err := parsePacketID(f.PacketID)
				if err != nil {
					t.Errorf("fixture %d: invalid packet ID %q: %v", i, f.PacketID, err)
					continue
				}

				registryKey := f.State + "_" + f.Direction
				registry, ok := packets.PacketRegistries[registryKey]
				if !ok {
					t.Logf("fixture %d: unknown state/direction %s, skipping", i, registryKey)
					continue
				}

				factory, ok := registry[packetID]
				if !ok {
					t.Logf("fixture %d: unknown packet ID %s in %s, skipping", i, f.PacketID, registryKey)
					continue
				}

				rawData, err := hex.DecodeString(f.RawData)
				if err != nil {
					t.Errorf("fixture %d: failed to decode hex: %v", i, err)
					continue
				}

				t.Run(f.PacketID, func(t *testing.T) {
					p := factory()
					testPacketRoundTrip(t, p, rawData)
				})
			}
		})
	}
}

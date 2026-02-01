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

	"github.com/go-mclib/data/packets"
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

// packetFactory creates a new packet instance by state, direction, and ID
type packetFactory func() jp.Packet

// registries map state+direction+id to packet factories
var packetRegistries = map[string]map[int]packetFactory{
	"handshake_c2s": {
		int(packets.C2SIntentionID): func() jp.Packet { return &packets.C2SIntention{} },
	},
	"status_c2s": {
		int(packets.C2SStatusRequestID):     func() jp.Packet { return &packets.C2SStatusRequest{} },
		int(packets.C2SPingRequestStatusID): func() jp.Packet { return &packets.C2SPingRequestStatus{} },
	},
	"status_s2c": {
		int(packets.S2CStatusResponseID):     func() jp.Packet { return &packets.S2CStatusResponse{} },
		int(packets.S2CPongResponseStatusID): func() jp.Packet { return &packets.S2CPongResponseStatus{} },
	},
	"login_c2s": {
		int(packets.C2SHelloID):               func() jp.Packet { return &packets.C2SHello{} },
		int(packets.C2SKeyID):                 func() jp.Packet { return &packets.C2SKey{} },
		int(packets.C2SCustomQueryAnswerID):   func() jp.Packet { return &packets.C2SCustomQueryAnswer{} },
		int(packets.C2SLoginAcknowledgedID):   func() jp.Packet { return &packets.C2SLoginAcknowledged{} },
		int(packets.C2SCookieResponseLoginID): func() jp.Packet { return &packets.C2SCookieResponseLogin{} },
	},
	"login_s2c": {
		int(packets.S2CLoginDisconnectLoginID): func() jp.Packet { return &packets.S2CLoginDisconnectLogin{} },
		int(packets.S2CHelloID):                func() jp.Packet { return &packets.S2CHello{} },
		int(packets.S2CLoginFinishedID):        func() jp.Packet { return &packets.S2CLoginFinished{} },
		int(packets.S2CLoginCompressionID):     func() jp.Packet { return &packets.S2CLoginCompression{} },
		int(packets.S2CCustomQueryID):          func() jp.Packet { return &packets.S2CCustomQuery{} },
		int(packets.S2CCookieRequestLoginID):   func() jp.Packet { return &packets.S2CCookieRequestLogin{} },
	},
	"configuration_c2s": {
		int(packets.C2SClientInformationConfigurationID): func() jp.Packet { return &packets.C2SClientInformationConfiguration{} },
		int(packets.C2SCookieResponseConfigurationID):    func() jp.Packet { return &packets.C2SCookieResponseConfiguration{} },
		int(packets.C2SCustomPayloadConfigurationID):     func() jp.Packet { return &packets.C2SCustomPayloadConfiguration{} },
		int(packets.C2SFinishConfigurationID):            func() jp.Packet { return &packets.C2SFinishConfiguration{} },
		int(packets.C2SKeepAliveConfigurationID):         func() jp.Packet { return &packets.C2SKeepAliveConfiguration{} },
		int(packets.C2SPongConfigurationID):              func() jp.Packet { return &packets.C2SPongConfiguration{} },
		int(packets.C2SResourcePackConfigurationID):      func() jp.Packet { return &packets.C2SResourcePackConfiguration{} },
		int(packets.C2SSelectKnownPacksID):               func() jp.Packet { return &packets.C2SSelectKnownPacks{} },
	},
	"configuration_s2c": {
		int(packets.S2CCookieRequestConfigurationID):       func() jp.Packet { return &packets.S2CCookieRequestConfiguration{} },
		int(packets.S2CCustomPayloadConfigurationID):       func() jp.Packet { return &packets.S2CCustomPayloadConfiguration{} },
		int(packets.S2CDisconnectConfigurationID):          func() jp.Packet { return &packets.S2CDisconnectConfiguration{} },
		int(packets.S2CFinishConfigurationID):              func() jp.Packet { return &packets.S2CFinishConfiguration{} },
		int(packets.S2CKeepAliveConfigurationID):           func() jp.Packet { return &packets.S2CKeepAliveConfiguration{} },
		int(packets.S2CPingConfigurationID):                func() jp.Packet { return &packets.S2CPingConfiguration{} },
		int(packets.S2CResetChatID):                        func() jp.Packet { return &packets.S2CResetChat{} },
		int(packets.S2CResourcePackPopConfigurationID):     func() jp.Packet { return &packets.S2CResourcePackPopConfiguration{} },
		int(packets.S2CResourcePackPushConfigurationID):    func() jp.Packet { return &packets.S2CResourcePackPushConfiguration{} },
		int(packets.S2CStoreCookieConfigurationID):         func() jp.Packet { return &packets.S2CStoreCookieConfiguration{} },
		int(packets.S2CTransferConfigurationID):            func() jp.Packet { return &packets.S2CTransferConfiguration{} },
		int(packets.S2CUpdateEnabledFeaturesID):            func() jp.Packet { return &packets.S2CUpdateEnabledFeatures{} },
		int(packets.S2CUpdateTagsConfigurationID):          func() jp.Packet { return &packets.S2CUpdateTagsConfiguration{} },
		int(packets.S2CSelectKnownPacksID):                 func() jp.Packet { return &packets.S2CSelectKnownPacks{} },
		int(packets.S2CCustomReportDetailsConfigurationID): func() jp.Packet { return &packets.S2CCustomReportDetailsConfiguration{} },
		int(packets.S2CServerLinksConfigurationID):         func() jp.Packet { return &packets.S2CServerLinksConfiguration{} },
		int(packets.S2CClearDialogConfigurationID):         func() jp.Packet { return &packets.S2CClearDialogConfiguration{} },
		int(packets.S2CCodeOfConductID):                    func() jp.Packet { return &packets.S2CCodeOfConduct{} },
	},
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
				registry, ok := packetRegistries[registryKey]
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

package packets_test

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

type packetCodec interface {
	Read(*ns.PacketBuffer) error
	Write(*ns.PacketBuffer) error
}

// These payloads were emitted by the corresponding vanilla 26.3 Java
// STREAM_CODEC, independently of the Go encoder.
func TestVanilla263PacketFixtures(t *testing.T) {
	effects := ns.PrefixedArray[ns.Identifier]{"minecraft:creeper", "minecraft:spider"}
	cases := []struct {
		name   string
		wire   string
		packet packetCodec
		want   packetCodec
	}{
		{"respawn_absent", "00136d696e6563726166743a6f766572776f726c64000000000000002a01000001000c3f03", &packets.S2CRespawn{}, &packets.S2CRespawn{DimensionName: "minecraft:overworld", HashedSeed: 42, GameMode: 1, PreviousGameMode: -1, IsFlat: true, PortalCooldown: 12, SeaLevel: 63, DataKept: 3}},
		{"respawn_spectator", "00136d696e6563726166743a6f766572776f726c64000000000000002a00040001000c3f01", &packets.S2CRespawn{}, &packets.S2CRespawn{DimensionName: "minecraft:overworld", HashedSeed: 42, PreviousGameMode: 3, IsFlat: true, PortalCooldown: 12, SeaLevel: 63, DataKept: 1}},
		{"advancement_hidden", "0101106d696e6563726166743a68696464656e0000010109637269746572696f6e013fc00000c0200000000001", &packets.S2CUpdateAdvancements{}, &packets.S2CUpdateAdvancements{Reset: true, ShowAdvancements: true, Added: []packets.PositionedAdvancement{{Id: "minecraft:hidden", Requirements: [][]ns.String{{"criterion"}}, SendsTelemetryEvent: true, X: 1.5, Y: -2.5}}}},
		{"advancement_display", "0001116d696e6563726166743a76697369626c65010e6d696e6563726166743a726f6f74010800055469746c6508000b4465736372697074696f6e0102000001000000072c6d696e6563726166743a6775692f616476616e63656d656e74732f6261636b67726f756e64732f73746f6e650102056669727374067365636f6e640040500000c090000001116d696e6563726166743a72656d6f7665640000", &packets.S2CUpdateAdvancements{}, &packets.S2CUpdateAdvancements{Added: []packets.PositionedAdvancement{{Id: "minecraft:visible", Parent: ns.Some(ns.Identifier("minecraft:root")), Display: &packets.AdvancementDisplay{Title: ns.TextComponent{Text: "Title"}, Description: ns.TextComponent{Text: "Description"}, Icon: ns.NewSlot(1, 2), Frame: packets.AdvancementFrameChallenge, Background: ns.Some(ns.Identifier("minecraft:gui/advancements/backgrounds/stone")), ShowToast: true, Hidden: true}, Requirements: [][]ns.String{{"first", "second"}}, X: 3.25, Y: -4.5}}, Removed: []ns.Identifier{"minecraft:removed"}}},
		{"teleport", "ac023ff80000000000004050500000000000c00600000000000042b40000c1f00000", &packets.C2SAcceptTeleportation{}, &packets.C2SAcceptTeleportation{TeleportId: 300, X: 1.5, Y: 65.25, Z: -2.75, Yaw: 90, Pitch: -30}},
		{"punch", "", &packets.C2SPunch{}, &packets.C2SPunch{}},
		{"sign", "fffffbc000020041056669727374067365636f6e640006666f7572746801", &packets.C2SSignUpdate{}, &packets.C2SSignUpdate{Location: ns.Position{X: -17, Y: 65, Z: 32}, Line1: "first", Line2: "second", Line3: "", Line4: "fourth", Slot: packets.SignTextFront}},
		{"move_linear", "ac02010002fffd0004", &packets.S2CMoveEntityPos{}, &packets.S2CMoveEntityPos{EntityId: 300, Delta: packets.VecDelta{X: 2, Y: -3, Z: 4}, OnGround: true}},
		{"move_stepped", "ac0204010002fffd000403fffb0006fff940e0", &packets.S2CMoveEntityPosRot{}, &packets.S2CMoveEntityPosRot{EntityId: 300, Delta: packets.VecDelta{Steps: []packets.VecDeltaStep{{Ticks: 1, X: 2, Y: -3, Z: 4}, {Ticks: 3, X: -5, Y: 6, Z: -7}}}, Yaw: 64, Pitch: 224}},
		{"move_rot", "ac020140e0", &packets.S2CMoveEntityRot{}, &packets.S2CMoveEntityRot{EntityId: 300, OnGround: true, Yaw: 64, Pitch: 224}},
		{"sync_linear", "ac02003ff80000000000004050500000000000c00600000000000042b40000c1f0000001", &packets.S2CEntityPositionSync{}, &packets.S2CEntityPositionSync{EntityId: 300, Path: packets.PositionPath{X: 1.5, Y: 65.25, Z: -2.75}, Yaw: 90, Pitch: -30, OnGround: true}},
		{"sync_stepped", "ac0201023ff80000000000004050500000000000c0060000000000000140040000000000004050900000000000c00e0000000000000342b40000c1f0000000", &packets.S2CEntityPositionSync{}, &packets.S2CEntityPositionSync{EntityId: 300, Path: packets.PositionPath{Steps: []packets.PositionStep{{X: 1.5, Y: 65.25, Z: -2.75, TickOffset: 1}, {X: 2.5, Y: 66.25, Z: -3.75, TickOffset: 3}}}, Yaw: 90, Pitch: -30}},
		{"swing", "ac02010209", &packets.S2CSwingAnimation{}, &packets.S2CSwingAnimation{EntityId: 300, Hand: packets.SwingOffHand, Type: packets.SwingAnimationStab, Duration: 9}},
		{"post_effects_play", "02116d696e6563726166743a63726565706572106d696e6563726166743a737069646572", &packets.S2CPostEffectsPlay{}, &packets.S2CPostEffectsPlay{Effects: effects}},
		{"post_effects_configuration", "02116d696e6563726166743a63726565706572106d696e6563726166743a737069646572", &packets.S2CPostEffectsConfiguration{}, &packets.S2CPostEffectsConfiguration{Effects: effects}},
		{"particles", "15001234563fc0000001003ff80000000000004050500000000000c0060000000000003dcccccd3e4ccccd3e99999a3ecccccd3f0000003f19999aac0202", &packets.S2CLevelParticles{}, &packets.S2CLevelParticles{ParticleId: 21, Data: []byte{0, 0x12, 0x34, 0x56, 0x3f, 0xc0, 0, 0}, LongDistance: true, X: 1.5, Y: 65.25, Z: -2.75, OffsetX: 0.1, OffsetY: 0.2, OffsetZ: 0.3, MaxSpeedX: 0.4, MaxSpeedY: 0.5, MaxSpeedZ: 0.6, ParticleCount: 300, RandomizationType: packets.ParticleRandomizationAlternativeWithSpeed}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			wire, err := hex.DecodeString(tc.wire)
			if err != nil {
				t.Fatal(err)
			}
			reader := bytes.NewReader(append(append([]byte(nil), wire...), 0x5a))
			if err := tc.packet.Read(ns.NewReaderFrom(reader)); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tc.packet, tc.want) {
				t.Fatalf("decoded = %#v; want %#v", tc.packet, tc.want)
			}
			if next, err := reader.ReadByte(); err != nil || next != 0x5a || reader.Len() != 0 {
				t.Fatalf("packet consumed wrong number of bytes: sentinel=%x, remaining=%d, err=%v", next, reader.Len(), err)
			}
			writer := ns.NewWriter()
			if err := tc.want.Write(writer); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(writer.Bytes(), wire) {
				t.Fatalf("encoded = %x; vanilla = %x", writer.Bytes(), wire)
			}
		})
	}
}

func TestMovementRejectsMalformedPaths(t *testing.T) {
	for _, tc := range []struct {
		name   string
		wire   string
		packet packetCodec
	}{
		{"negative_properties", "01ffffffff0f", &packets.S2CMoveEntityPos{}},
		{"huge_relative_count_truncated", "01ffffffff07", &packets.S2CMoveEntityPos{}},
		{"missing_relative_coordinate", "01020100010002", &packets.S2CMoveEntityPos{}},
		{"unknown_absolute_kind", "0102", &packets.S2CEntityPositionSync{}},
		{"empty_absolute_steps", "010100", &packets.S2CEntityPositionSync{}},
		{"negative_absolute_count", "0101ffffffff0f", &packets.S2CEntityPositionSync{}},
		{"huge_absolute_count_truncated", "0101ffffffff07", &packets.S2CEntityPositionSync{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wire, err := hex.DecodeString(tc.wire)
			if err != nil {
				t.Fatal(err)
			}
			if err := tc.packet.Read(ns.NewReader(wire)); err == nil {
				t.Fatal("accepted malformed movement path")
			}
		})
	}
	p := packets.S2CEntityPositionSync{Path: packets.PositionPath{Steps: []packets.PositionStep{}}}
	if err := p.Write(ns.NewWriter()); err == nil {
		t.Fatal("accepted an empty stepped position path")
	}
}

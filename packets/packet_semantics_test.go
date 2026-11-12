package packets_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestLoginOptionalPreviousGameMode(t *testing.T) {
	// Hand-written 26.3 login payload. Previous survival is encoded as 1, not
	// as the old signed byte 0; the booleans and spawn data must stay aligned.
	wire, err := hex.DecodeString("000000070001136d696e6563726166743a6f766572776f726c64140a0800010000136d696e6563726166743a6f766572776f726c64000000000000002a0301010000053f0100")
	if err != nil {
		t.Fatal(err)
	}
	var packet packets.S2CLogin
	reader := bytes.NewReader(append(append([]byte(nil), wire...), 0x5a))
	if err := packet.Read(ns.NewReaderFrom(reader)); err != nil {
		t.Fatal(err)
	}
	if packet.EntityID != 7 || packet.GameMode != 3 || packet.PreviousGameMode != 0 || !bool(packet.IsDebug) || bool(packet.IsFlat) || packet.DeathLocation.Present || packet.PortalCooldown != 5 || packet.SeaLevel != 63 || !bool(packet.OnlineMode) || bool(packet.EnforcesSecureChat) {
		t.Fatalf("login spawn info = %+v", packet)
	}
	if sentinel, err := reader.ReadByte(); err != nil || sentinel != 0x5a || reader.Len() != 0 {
		t.Fatalf("login payload boundary: sentinel=%x, err=%v", sentinel, err)
	}
	out := ns.NewWriter()
	if err := packet.Write(out); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(out.Bytes(), wire) {
		t.Fatalf("login wire = %x; want %x", out.Bytes(), wire)
	}
}

func TestPlayerActionAndAnimateOrdinals(t *testing.T) {
	for _, tc := range []struct {
		action ns.VarInt
		wire   byte
	}{
		{packets.PlayerActionStartDestroyBlock, 0},
		{packets.PlayerActionChangeDestroyDirection, 1},
		{packets.PlayerActionAbortDestroyBlock, 2},
		{packets.PlayerActionStopDestroyBlock, 3},
		{packets.PlayerActionDropAllItems, 4},
		{packets.PlayerActionDropItem, 5},
		{packets.PlayerActionReleaseUseItem, 6},
		{packets.PlayerActionSwapItemWithOffhand, 7},
		{packets.PlayerActionStab, 8},
	} {
		p := packets.C2SPlayerAction{Status: tc.action, Face: 2, Sequence: 1}
		out := ns.NewWriter()
		if err := p.Write(out); err != nil {
			t.Fatal(err)
		}
		want := []byte{tc.wire, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1}
		if !bytes.Equal(out.Bytes(), want) {
			t.Fatalf("action %d wire = %x; want %x", tc.action, out.Bytes(), want)
		}
	}
	for _, tc := range []struct {
		animation ns.Uint8
		wire      byte
	}{
		{packets.AnimateWakeUp, 0},
		{packets.AnimateCriticalHit, 1},
		{packets.AnimateMagicCriticalHit, 2},
	} {
		p := packets.S2CAnimate{EntityID: 1, Animation: tc.animation}
		out := ns.NewWriter()
		if err := p.Write(out); err != nil {
			t.Fatal(err)
		}
		if want := []byte{1, tc.wire}; !bytes.Equal(out.Bytes(), want) {
			t.Fatalf("animation wire = %x; want %x", out.Bytes(), want)
		}
	}
}

func TestTransientBlockAndSignSlotFixtures(t *testing.T) {
	for _, tc := range []struct {
		packet packetCodec
		wire   []byte
	}{
		{&packets.S2CAddTransientBlock{BlockState: 300}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0xac, 2}},
		{&packets.S2COpenSignEditor{Slot: packets.SignTextBack}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{&packets.S2COpenSignEditor{Slot: packets.SignTextFront}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 1}},
	} {
		out := ns.NewWriter()
		if err := tc.packet.Write(out); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(out.Bytes(), tc.wire) {
			t.Fatalf("%T wire = %x; want %x", tc.packet, out.Bytes(), tc.wire)
		}
	}
}

func TestExplosionPlaySoundIsTerminal(t *testing.T) {
	// Radius=4, block count=0, no knockback, explosion particle id=29,
	// sound holder=700, one weighted particle (dust options, scale/speed,
	// weight=2). The final sound flag is outside this unframed tail.
	tail, err := hex.DecodeString("4080000000000000001dbc050115001234563fc000003f8000003f00000002")
	if err != nil {
		t.Fatal(err)
	}
	for _, sound := range []bool{false, true} {
		want := append(make([]byte, 24), tail...)
		flag := byte(0)
		if sound {
			flag = 1
		}
		want = append(want, flag)
		packet := packets.S2CExplode{Data: tail, PlaySound: ns.Boolean(sound)}
		out := ns.NewWriter()
		if err := packet.Write(out); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(out.Bytes(), want) {
			t.Fatalf("explosion wire = %x; want %x", out.Bytes(), want)
		}
		var decoded packets.S2CExplode
		if err := decoded.Read(ns.NewReader(want)); err != nil {
			t.Fatal(err)
		}
		if bool(decoded.PlaySound) != sound || !bytes.Equal(decoded.Data, tail) {
			t.Fatalf("decoded explosion = %+v", decoded)
		}
	}
	var truncated packets.S2CExplode
	if err := truncated.Read(ns.NewReader(make([]byte, 24))); err == nil {
		t.Fatal("accepted explosion without tail or sound flag")
	}
}

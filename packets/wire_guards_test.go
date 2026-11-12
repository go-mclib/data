package packets_test

import (
	"math"
	"testing"

	"github.com/go-mclib/data/packetids"
	"github.com/go-mclib/data/packets"
	jp "github.com/go-mclib/protocol/javaproto"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// countPayload builds a body of prefix fields followed by a bare VarInt count.
func countPayload(t *testing.T, count ns.VarInt, prefix func(*ns.PacketBuffer) error) []byte {
	t.Helper()
	w := ns.NewWriter()
	if prefix != nil {
		if err := prefix(w); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.WriteVarInt(count); err != nil {
		t.Fatal(err)
	}
	return w.Bytes()
}

// Crafted counts must fail cleanly instead of panicking in make or requesting
// multi-GiB allocations.
func TestWireCountGuards(t *testing.T) {
	containerClickPrefix := func(w *ns.PacketBuffer) error {
		for _, step := range []func() error{
			func() error { return w.WriteVarInt(0) },
			func() error { return w.WriteVarInt(0) },
			func() error { return w.WriteInt16(0) },
			func() error { return w.WriteInt8(0) },
			func() error { return w.WriteVarInt(0) },
		} {
			if err := step(); err != nil {
				return err
			}
		}
		return nil
	}
	loginPrefix := func(w *ns.PacketBuffer) error {
		if err := w.WriteUUID(ns.UUID{}); err != nil {
			return err
		}
		return w.WriteString("Steve")
	}
	cases := []struct {
		name   string
		packet packetCodec
		limit  ns.VarInt // vanilla cap; limit+1 must be rejected
		prefix func(*ns.PacketBuffer) error
	}{
		{"C2SSelectKnownPacks", &packets.C2SSelectKnownPacks{}, 64, nil},
		{"C2SContainerClick", &packets.C2SContainerClick{}, 128, containerClickPrefix},
		{"C2SDebugSubscriptionRequest", &packets.C2SDebugSubscriptionRequest{}, 32, nil},
		{"S2CCustomReportDetailsConfiguration", &packets.S2CCustomReportDetailsConfiguration{}, 32, nil},
		{"S2CLoginFinished", &packets.S2CLoginFinished{}, 16, loginPrefix},
		{"C2SSetGameRule", &packets.C2SSetGameRule{}, 65536, nil},
		{"S2CUpdateEnabledFeatures", &packets.S2CUpdateEnabledFeatures{}, 65536, nil},
		{"S2CPlayerInfoRemove", &packets.S2CPlayerInfoRemove{}, 65536, nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, count := range []ns.VarInt{-1, tc.limit + 1, math.MaxInt32} {
				if err := tc.packet.Read(ns.NewReader(countPayload(t, count, tc.prefix))); err == nil {
					t.Errorf("count %d accepted", count)
				}
			}
		})
	}
}

func TestDamageEventSourcePositionIsThreeDoubles(t *testing.T) {
	in := &packets.S2CDamageEvent{EntityID: 7, SourcePosition: ns.PrefixedOptional[[3]ns.Float64]{Present: true, Value: [3]ns.Float64{1.5, -2, 3}}}
	w := ns.NewWriter()
	if err := in.Write(w); err != nil {
		t.Fatal(err)
	}
	// 4 single-byte VarInts + presence byte + 3 raw doubles, no length prefix
	if got, want := len(w.Bytes()), 4+1+24; got != want {
		t.Fatalf("encoded length %d, want %d", got, want)
	}
	var out packets.S2CDamageEvent
	if err := out.Read(ns.NewReader(w.Bytes())); err != nil {
		t.Fatal(err)
	}
	if out.SourcePosition != in.SourcePosition {
		t.Fatalf("source position %+v, want %+v", out.SourcePosition, in.SourcePosition)
	}
}

func TestDeleteChatFullSignatureRoundTrip(t *testing.T) {
	sig := make(ns.ByteArray, 256)
	sig[0] = 0xFF // would be misread as a length prefix
	sig[255] = 1
	w := ns.NewWriter()
	if err := (&packets.S2CDeleteChat{Signature: sig}).Write(w); err != nil {
		t.Fatal(err)
	}
	if got := len(w.Bytes()); got != 1+256 {
		t.Fatalf("encoded length %d, want 257", got)
	}
	var out packets.S2CDeleteChat
	if err := out.Read(ns.NewReader(w.Bytes())); err != nil {
		t.Fatal(err)
	}
	if string(out.Signature) != string(sig) {
		t.Fatal("signature mismatch")
	}
	if err := (&packets.S2CDeleteChat{Signature: sig[:10]}).Write(ns.NewWriter()); err == nil {
		t.Fatal("short full signature accepted")
	}
}

func TestDebugSampleRoundTrip(t *testing.T) {
	in := &packets.S2CDebugSample{Sample: []ns.Int64{1, -2, 3}, SampleType: 1}
	w := ns.NewWriter()
	if err := in.Write(w); err != nil {
		t.Fatal(err)
	}
	if got, want := len(w.Bytes()), 1+3*8+1; got != want {
		t.Fatalf("encoded length %d, want %d", got, want)
	}
	var out packets.S2CDebugSample
	if err := out.Read(ns.NewReader(w.Bytes())); err != nil {
		t.Fatal(err)
	}
	if len(out.Sample) != 3 || out.Sample[1] != -2 || out.SampleType != 1 {
		t.Fatalf("got %+v", out)
	}
}

func TestUpdateTagsPlaySharesConfigurationCodec(t *testing.T) {
	tags := []packets.TagRegistry{{Registry: "minecraft:block", Tags: []packets.Tag{{TagName: "minecraft:logs", Entries: []ns.VarInt{1, 2}}}}}
	cfg := ns.NewWriter()
	if err := (&packets.S2CUpdateTagsConfiguration{ArrayOfTags: tags}).Write(cfg); err != nil {
		t.Fatal(err)
	}
	var out packets.S2CUpdateTagsPlay
	if err := out.Read(ns.NewReader(cfg.Bytes())); err != nil {
		t.Fatal(err)
	}
	if len(out.ArrayOfTags) != 1 || len(out.ArrayOfTags[0].Tags[0].Entries) != 2 {
		t.Fatalf("got %+v", out.ArrayOfTags)
	}
}

func TestWriteRejectsMissingRequiredData(t *testing.T) {
	if err := (&packets.C2SChat{Message: "hi"}).Write(ns.NewWriter()); err == nil {
		t.Error("C2SChat with nil Acknowledged accepted")
	}
	if err := (&packets.C2SChatCommandSigned{Command: "me"}).Write(ns.NewWriter()); err == nil {
		t.Error("C2SChatCommandSigned with nil Acknowledged accepted")
	}
	chat := &packets.S2CPlayerChat{FilterMask: packets.FilterMask{Type: packets.FilterMaskPartiallyFiltered}}
	if err := chat.Write(ns.NewWriter()); err == nil {
		t.Error("partially filtered S2CPlayerChat without mask accepted")
	}
}

func TestNewPacket(t *testing.T) {
	pk, ok := packets.NewPacket(jp.StatePlay, jp.C2S, packetids.C2SChatID)
	if _, isChat := pk.(*packets.C2SChat); !ok || !isChat {
		t.Fatalf("NewPacket(play, c2s, chat) = %T, %v", pk, ok)
	}
	if _, ok := packets.NewPacket(jp.StatePlay, jp.S2C, -1); ok {
		t.Fatal("unknown packet ID accepted")
	}
}

package packets

import (
	"bytes"
	"reflect"
	"testing"

	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestPlayerInfoUpdateAllActionsRoundTrip(t *testing.T) {
	steve, alex := ns.UUID{1, 2, 3}, ns.UUID{4, 5, 6}
	session := ns.UUID{9}
	display := ns.TextComponent{Text: "Knight ", Extra: []ns.TextComponent{{Text: "Steve", Color: "gold"}}}

	b := ns.NewWriter()
	b.WriteByte(0xff)
	b.WriteVarInt(2)
	for i, id := range []ns.UUID{steve, alex} {
		b.WriteUUID(id)
		// add_player: name + properties
		b.WriteString([]ns.String{"Steve", "Alex"}[i])
		b.WriteVarInt(2)
		b.WriteString("textures")
		b.WriteString("skin")
		b.WriteBool(true)
		b.WriteString("sig")
		b.WriteString("other")
		b.WriteString("v")
		b.WriteBool(false)
		// initialize_chat
		b.WriteBool(i == 0)
		if i == 0 {
			b.WriteUUID(session)
			b.WriteInt64(1234)
			b.WriteByteArray([]byte{1, 2})
			b.WriteByteArray([]byte{3})
		}
		b.WriteVarInt(1)    // game mode
		b.WriteBool(true)   // listed
		b.WriteVarInt(42)   // latency
		b.WriteBool(i == 0) // display name present
		if i == 0 {
			b.WriteTextComponent(display)
		}
		b.WriteVarInt(-3) // list order
		b.WriteBool(true) // hat
	}
	raw := b.Bytes()

	var got S2CPlayerInfoUpdate
	if err := got.Read(ns.NewReader(raw)); err != nil {
		t.Fatal(err)
	}
	props := []GameProfileProperty{{Name: "textures", Value: "skin", Signature: ns.Some[ns.String]("sig")}, {Name: "other", Value: "v"}}
	want := S2CPlayerInfoUpdate{Actions: 0xff, Entries: []PlayerInfoEntry{
		{UUID: steve, Name: "Steve", Properties: props, ChatSession: ns.Some(RemoteChatSession{SessionID: session, ExpiresAt: 1234, PublicKey: []byte{1, 2}, KeySignature: []byte{3}}),
			GameMode: 1, Listed: true, Latency: 42, DisplayName: ns.Some(display), ListOrder: -3, ShowHat: true},
		{UUID: alex, Name: "Alex", Properties: props, GameMode: 1, Listed: true, Latency: 42, ListOrder: -3, ShowHat: true},
	}}
	if got.Entries[0].DisplayName.Value.String() != "Knight Steve" {
		t.Fatalf("display name = %q", got.Entries[0].DisplayName.Value.String())
	}
	// TextComponent decoding fills defaults; compare it by re-encoding below.
	got.Entries[0].DisplayName.Value, want.Entries[0].DisplayName.Value = ns.TextComponent{}, ns.TextComponent{}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("decoded:\n got %+v\nwant %+v", got, want)
	}
	want.Entries[0].DisplayName.Value = display
	w := ns.NewWriter()
	if err := want.Write(w); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(w.Bytes(), raw) {
		t.Fatalf("re-encoded differs:\n got %x\nwant %x", w.Bytes(), raw)
	}
}

// Fields of unset actions are neither read nor written.
func TestPlayerInfoUpdatePartialActions(t *testing.T) {
	id := ns.UUID{7}
	raw := []byte{byte(PlayerInfoUpdateLatency | PlayerInfoUpdateHat), 1}
	raw = append(raw, id[:]...)
	raw = append(raw, 0x2a, 1) // latency 42, hat true
	var got S2CPlayerInfoUpdate
	if err := got.Read(ns.NewReader(raw)); err != nil {
		t.Fatal(err)
	}
	if !got.Actions.Has(PlayerInfoUpdateLatency) || got.Actions.Has(PlayerInfoAddPlayer) || len(got.Entries) != 1 ||
		got.Entries[0].UUID != id || got.Entries[0].Latency != 42 || !bool(got.Entries[0].ShowHat) {
		t.Fatalf("decoded %+v", got)
	}
	got.Entries[0].Name, got.Entries[0].GameMode = "ignored", 3
	w := ns.NewWriter()
	if err := got.Write(w); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(w.Bytes(), raw) {
		t.Fatalf("re-encoded %x, want %x", w.Bytes(), raw)
	}
}

func TestPlayerInfoUpdateRejectsMalformed(t *testing.T) {
	addPlayer := func(props ns.VarInt) []byte {
		b := ns.NewWriter()
		b.WriteByte(byte(PlayerInfoAddPlayer))
		b.WriteVarInt(1)
		b.WriteUUID(ns.UUID{})
		b.WriteString("Steve")
		b.WriteVarInt(props)
		return b.Bytes()
	}
	for name, raw := range map[string][]byte{
		"negative entries":   {0x10, 0x7f, 0xff, 0xff, 0xff, 0x0f},
		"huge entries":       {0x10, 0x80, 0x80, 0x80, 0x80, 0x04},
		"too many props":     addPlayer(17),
		"negative props":     addPlayer(-1),
		"truncated property": addPlayer(1),
		"long name":          append([]byte{byte(PlayerInfoAddPlayer), 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 17}, bytes.Repeat([]byte{'a'}, 17)...),
	} {
		var p S2CPlayerInfoUpdate
		if err := p.Read(ns.NewReader(raw)); err == nil {
			t.Errorf("%s: accepted %x", name, raw)
		}
	}
	p := S2CPlayerInfoUpdate{Actions: PlayerInfoAddPlayer, Entries: []PlayerInfoEntry{{Properties: make([]GameProfileProperty, 17)}}}
	if err := p.Write(ns.NewWriter()); err == nil {
		t.Fatal("serialized 17 profile properties")
	}
}

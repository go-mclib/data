package packets

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func parserID(t *testing.T, name string) ns.VarInt {
	t.Helper()
	id := registries.CommandArgumentType.Get(name)
	if id < 0 {
		t.Fatalf("parser %s not registered", name)
	}
	return ns.VarInt(id)
}

// commandGraphBytes crafts a graph exercising every node type, redirects,
// suggestions, the restricted flag and every parser property layout.
func commandGraphBytes(t *testing.T) ([]byte, S2CCommands) {
	t.Helper()
	b := ns.NewWriter()
	args := []struct {
		name  string
		write func()
		props CommandParserProperties
	}{
		{"brigadier:integer", func() { b.WriteByte(3); b.WriteInt32(-5); b.WriteInt32(7) }, CommandParserProperties{Flags: 3, IntMin: -5, IntMax: 7}},
		{"brigadier:long", func() { b.WriteByte(1); b.WriteInt64(-1 << 40) }, CommandParserProperties{Flags: 1, IntMin: -1 << 40}},
		{"brigadier:float", func() { b.WriteByte(1); b.WriteFloat32(0.5) }, CommandParserProperties{Flags: 1, FloatMin: 0.5}},
		{"brigadier:double", func() { b.WriteByte(2); b.WriteFloat64(2.25) }, CommandParserProperties{Flags: 2, FloatMax: 2.25}},
		{"brigadier:string", func() { b.WriteVarInt(2) }, CommandParserProperties{StringMode: 2}},
		{"minecraft:entity", func() { b.WriteByte(3) }, CommandParserProperties{Flags: 3}},
		{"minecraft:score_holder", func() { b.WriteByte(1) }, CommandParserProperties{Flags: 1}},
		{"minecraft:time", func() { b.WriteInt32(20) }, CommandParserProperties{TimeMin: 20}},
		{"minecraft:resource", func() { b.WriteIdentifier("minecraft:mob_effect") }, CommandParserProperties{Registry: "minecraft:mob_effect"}},
		{"minecraft:resource_selector", func() { b.WriteIdentifier("minecraft:recipe") }, CommandParserProperties{Registry: "minecraft:recipe"}},
		{"minecraft:message", func() {}, CommandParserProperties{}},
	}
	want := S2CCommands{RootIndex: 0}
	nodes := 2 + len(args)
	b.WriteVarInt(ns.VarInt(nodes))
	// root -> literal "cmd" (restricted, executable) -> every argument
	b.WriteByte(0)
	b.WriteVarInt(1)
	b.WriteVarInt(1)
	want.Nodes = append(want.Nodes, CommandNode{Children: []ns.VarInt{1}})
	b.WriteByte(CommandNodeLiteral | CommandNodeExecutable | CommandNodeRestricted)
	b.WriteVarInt(ns.VarInt(len(args)))
	literal := CommandNode{Flags: CommandNodeLiteral | CommandNodeExecutable | CommandNodeRestricted, Name: "cmd"}
	for i := range args {
		b.WriteVarInt(ns.VarInt(2 + i))
		literal.Children = append(literal.Children, ns.VarInt(2+i))
	}
	b.WriteString("cmd")
	want.Nodes = append(want.Nodes, literal)
	for i, arg := range args {
		node := CommandNode{Flags: CommandNodeArgument, Children: []ns.VarInt{}, Name: ns.String("a" + string(rune('a'+i))), ParserID: parserID(t, arg.name), Properties: arg.props}
		if i == 0 {
			// redirect back to the literal, with a custom suggestion provider
			node.Flags |= CommandNodeRedirect | CommandNodeCustomSuggestions
			node.Redirect = 1
			node.SuggestionsType = "minecraft:ask_server"
		}
		b.WriteByte(node.Flags)
		b.WriteVarInt(0)
		if i == 0 {
			b.WriteVarInt(1)
		}
		b.WriteString(node.Name)
		b.WriteVarInt(node.ParserID)
		arg.write()
		if i == 0 {
			b.WriteIdentifier("minecraft:ask_server")
		}
		want.Nodes = append(want.Nodes, node)
	}
	b.WriteVarInt(0)
	want.Nodes[0].Children = []ns.VarInt{1}
	return b.Bytes(), want
}

func TestCommandsRoundTrip(t *testing.T) {
	raw, want := commandGraphBytes(t)
	var got S2CCommands
	if err := got.Read(ns.NewReader(raw)); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("decoded graph:\n got %+v\nwant %+v", got, want)
	}
	w := ns.NewWriter()
	if err := got.Write(w); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(w.Bytes(), raw) {
		t.Fatalf("re-encoded graph differs:\n got %x\nwant %x", w.Bytes(), raw)
	}
}

func TestCommandsRejectsMalformedGraphs(t *testing.T) {
	message := parserID(t, "minecraft:message")
	str := parserID(t, "brigadier:string")
	graph := func(build func(b *ns.PacketBuffer)) []byte {
		b := ns.NewWriter()
		build(b)
		return b.Bytes()
	}
	root := func(b *ns.PacketBuffer, children ...ns.VarInt) {
		b.WriteByte(0)
		b.WriteVarInt(ns.VarInt(len(children)))
		for _, c := range children {
			b.WriteVarInt(c)
		}
	}
	valid := graph(func(b *ns.PacketBuffer) { b.WriteVarInt(1); root(b); b.WriteVarInt(0) })
	for name, raw := range map[string][]byte{
		"negative node count": graph(func(b *ns.PacketBuffer) { b.WriteVarInt(-1) }),
		"huge node count":     graph(func(b *ns.PacketBuffer) { b.WriteVarInt(1 << 30) }),
		"empty graph":         graph(func(b *ns.PacketBuffer) { b.WriteVarInt(0); b.WriteVarInt(0) }),
		"child count":         graph(func(b *ns.PacketBuffer) { b.WriteVarInt(1); b.WriteByte(0); b.WriteVarInt(2) }),
		"child bounds":        graph(func(b *ns.PacketBuffer) { b.WriteVarInt(1); root(b, 1); b.WriteVarInt(0) }),
		"redirect bounds": graph(func(b *ns.PacketBuffer) {
			b.WriteVarInt(1)
			b.WriteByte(8)
			b.WriteVarInt(0)
			b.WriteVarInt(5)
			b.WriteVarInt(0)
		}),
		"root bounds": graph(func(b *ns.PacketBuffer) { b.WriteVarInt(1); root(b); b.WriteVarInt(1) }),
		"root type": graph(func(b *ns.PacketBuffer) {
			b.WriteVarInt(1)
			b.WriteByte(1)
			b.WriteVarInt(0)
			b.WriteString("x")
			b.WriteVarInt(0)
		}),
		"node type 3": graph(func(b *ns.PacketBuffer) {
			b.WriteVarInt(2)
			root(b)
			b.WriteByte(3)
			b.WriteVarInt(0)
			b.WriteVarInt(0)
		}),
		"literal suggestions": graph(func(b *ns.PacketBuffer) {
			b.WriteVarInt(2)
			root(b, 1)
			b.WriteByte(1 | 0x10)
			b.WriteVarInt(0)
			b.WriteString("x")
			b.WriteVarInt(0)
		}),
		"unknown parser": graph(func(b *ns.PacketBuffer) {
			b.WriteVarInt(2)
			root(b, 1)
			b.WriteByte(2)
			b.WriteVarInt(0)
			b.WriteString("x")
			b.WriteVarInt(1 << 20)
			b.WriteVarInt(0)
		}),
		"string mode": graph(func(b *ns.PacketBuffer) {
			b.WriteVarInt(2)
			root(b, 1)
			b.WriteByte(2)
			b.WriteVarInt(0)
			b.WriteString("x")
			b.WriteVarInt(str)
			b.WriteVarInt(3)
			b.WriteVarInt(0)
		}),
		"missing suggestion": graph(func(b *ns.PacketBuffer) {
			b.WriteVarInt(2)
			root(b, 1)
			b.WriteByte(2 | 0x10)
			b.WriteVarInt(0)
			b.WriteString("x")
			b.WriteVarInt(message)
		}),
		"trailing byte": append(append([]byte{}, valid...), 1),
		"truncated":     valid[:len(valid)-1],
	} {
		var p S2CCommands
		if err := p.Read(ns.NewReader(raw)); err == nil {
			t.Errorf("%s: accepted %x", name, raw)
		}
	}
	var p S2CCommands
	if err := p.Read(ns.NewReader(valid)); err != nil {
		t.Fatal(err)
	}
	for name, bad := range map[string]S2CCommands{
		"root":           {Nodes: []CommandNode{{Flags: CommandNodeLiteral}}},
		"child":          {Nodes: []CommandNode{{Children: []ns.VarInt{3}}}},
		"unknown parser": {Nodes: []CommandNode{{Children: []ns.VarInt{1}}, {Flags: CommandNodeArgument, ParserID: 1 << 20}}},
	} {
		if err := bad.Write(ns.NewWriter()); err == nil {
			t.Errorf("%s: serialized invalid graph", name)
		}
	}
}

func TestSignedCommandRejectsOversizedSignatureArray(t *testing.T) {
	for _, count := range []ns.VarInt{-1, 9} {
		b := ns.NewWriter()
		b.WriteString("msg alice hi")
		b.WriteInt64(0)
		b.WriteInt64(0)
		b.WriteVarInt(count)
		var p C2SChatCommandSigned
		if err := p.Read(ns.NewReader(b.Bytes())); err == nil {
			t.Fatalf("accepted signature count %d", count)
		}
	}
	p := C2SChatCommandSigned{Command: "msg alice hi", ArgumentSignatures: make([]CommandArgumentSignature, 9), Acknowledged: ns.NewFixedBitSet(20)}
	if err := p.Write(ns.NewWriter()); err == nil {
		t.Fatal("serialized oversized signature array")
	}
}

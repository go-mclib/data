package packets

import (
	"errors"
	"fmt"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// Command node flag bits (ClientboundCommandsPacket).
const (
	CommandNodeTypeMask          byte = 0x03
	CommandNodeRoot              byte = 0
	CommandNodeLiteral           byte = 1
	CommandNodeArgument          byte = 2
	CommandNodeExecutable        byte = 0x04
	CommandNodeRedirect          byte = 0x08
	CommandNodeCustomSuggestions byte = 0x10
	CommandNodeRestricted        byte = 0x20
)

// CommandNode is one Brigadier node of the commands graph. Children and
// Redirect index S2CCommands.Nodes.
type CommandNode struct {
	Flags    byte
	Children []ns.VarInt
	// Redirect is meaningful when Flags has CommandNodeRedirect.
	Redirect ns.VarInt
	// Name is the literal or argument name (absent for the root).
	Name ns.String
	// ParserID (a command_argument_type registry id) and Properties describe
	// argument nodes.
	ParserID   ns.VarInt
	Properties CommandParserProperties
	// SuggestionsType is an argument node's suggestion provider, present when
	// Flags has CommandNodeCustomSuggestions.
	SuggestionsType ns.Identifier
}

// Type returns the node type: CommandNodeRoot, CommandNodeLiteral or
// CommandNodeArgument.
func (n *CommandNode) Type() byte { return n.Flags & CommandNodeTypeMask }

// CommandParserProperties holds the parser-specific properties an argument
// node carries (ArgumentTypeInfo.serializeToNetwork). Parsers registered as
// singletons have none; which fields apply depends on the parser:
type CommandParserProperties struct {
	// Flags: brigadier numeric parsers' min (0x01) / max (0x02) presence,
	// minecraft:entity single (0x01) / players only (0x02), and
	// minecraft:score_holder multiple (0x01).
	Flags byte
	// IntMin, IntMax: brigadier:integer and brigadier:long bounds.
	IntMin, IntMax int64
	// FloatMin, FloatMax: brigadier:float and brigadier:double bounds.
	FloatMin, FloatMax float64
	// StringMode: brigadier:string StringType (0 single word, 1 quotable
	// phrase, 2 greedy phrase).
	StringMode ns.VarInt
	// TimeMin: minecraft:time minimum ticks.
	TimeMin ns.Int32
	// Registry: the registry key of minecraft:resource, resource_key,
	// resource_or_tag, resource_or_tag_key and resource_selector.
	Registry ns.Identifier
}

// Numeric parser and entity parser property flags.
const (
	CommandNumberHasMin      byte = 0x01
	CommandNumberHasMax      byte = 0x02
	CommandEntitySingle      byte = 0x01
	CommandEntityPlayersOnly byte = 0x02
)

func (p *S2CCommands) read(buf *ns.PacketBuffer) error {
	count, err := readCount(buf, "command node", maxListCount)
	if err != nil {
		return err
	}
	p.Nodes = make([]CommandNode, count)
	for i := range p.Nodes {
		if err := p.Nodes[i].read(buf, count); err != nil {
			return err
		}
	}
	if p.RootIndex, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if err := p.validate(); err != nil {
		return err
	}
	// Vanilla's PacketDecoder rejects a frame the packet did not consume.
	if _, err := buf.ReadByte(); err == nil {
		return errors.New("trailing command graph bytes")
	}
	return nil
}

func (p *S2CCommands) write(buf *ns.PacketBuffer) error {
	if len(p.Nodes) > maxListCount {
		return fmt.Errorf("command node count %d exceeds %d", len(p.Nodes), maxListCount)
	}
	if err := p.validate(); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(len(p.Nodes))); err != nil {
		return err
	}
	for i := range p.Nodes {
		if err := p.Nodes[i].write(buf); err != nil {
			return err
		}
	}
	return buf.WriteVarInt(p.RootIndex)
}

// validate checks graph structure, so consumers may index Nodes by Children,
// Redirect and RootIndex without bounds checks.
func (p *S2CCommands) validate() error {
	count := ns.VarInt(len(p.Nodes))
	for i := range p.Nodes {
		n := &p.Nodes[i]
		if n.Type() == CommandNodeTypeMask {
			return fmt.Errorf("command node %d: invalid type", i)
		}
		if n.Flags&CommandNodeCustomSuggestions != 0 && n.Type() != CommandNodeArgument {
			return fmt.Errorf("command node %d: suggestions on non-argument node", i)
		}
		for _, child := range n.Children {
			if child < 0 || child >= count {
				return fmt.Errorf("command node %d: child %d out of bounds", i, child)
			}
		}
		if n.Flags&CommandNodeRedirect != 0 && (n.Redirect < 0 || n.Redirect >= count) {
			return fmt.Errorf("command node %d: redirect %d out of bounds", i, n.Redirect)
		}
	}
	if p.RootIndex < 0 || p.RootIndex >= count || p.Nodes[p.RootIndex].Type() != CommandNodeRoot {
		return fmt.Errorf("invalid command root %d", p.RootIndex)
	}
	return nil
}

func (n *CommandNode) read(buf *ns.PacketBuffer, nodes int) error {
	var err error
	if n.Flags, err = buf.ReadByte(); err != nil {
		return err
	}
	children, err := readCount(buf, "command child", nodes)
	if err != nil {
		return err
	}
	n.Children = make([]ns.VarInt, children)
	for i := range n.Children {
		if n.Children[i], err = buf.ReadVarInt(); err != nil {
			return err
		}
	}
	if n.Flags&CommandNodeRedirect != 0 {
		if n.Redirect, err = buf.ReadVarInt(); err != nil {
			return err
		}
	}
	switch n.Type() {
	case CommandNodeLiteral:
		n.Name, err = buf.ReadString(32767)
		return err
	case CommandNodeArgument:
		if n.Name, err = buf.ReadString(32767); err != nil {
			return err
		}
		if n.ParserID, err = buf.ReadVarInt(); err != nil {
			return err
		}
		if err := n.Properties.read(buf, n.ParserID); err != nil {
			return err
		}
		if n.Flags&CommandNodeCustomSuggestions != 0 {
			n.SuggestionsType, err = buf.ReadIdentifier()
		}
		return err
	case CommandNodeRoot:
		return nil
	default:
		return errors.New("invalid command node type")
	}
}

func (n *CommandNode) write(buf *ns.PacketBuffer) error {
	if err := buf.WriteByte(n.Flags); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(len(n.Children))); err != nil {
		return err
	}
	for _, child := range n.Children {
		if err := buf.WriteVarInt(child); err != nil {
			return err
		}
	}
	if n.Flags&CommandNodeRedirect != 0 {
		if err := buf.WriteVarInt(n.Redirect); err != nil {
			return err
		}
	}
	if n.Type() == CommandNodeRoot {
		return nil
	}
	if err := buf.WriteString(n.Name); err != nil {
		return err
	}
	if n.Type() != CommandNodeArgument {
		return nil
	}
	if err := buf.WriteVarInt(n.ParserID); err != nil {
		return err
	}
	if err := n.Properties.write(buf, n.ParserID); err != nil {
		return err
	}
	if n.Flags&CommandNodeCustomSuggestions != 0 {
		return buf.WriteIdentifier(n.SuggestionsType)
	}
	return nil
}

// commandParser resolves a parser id. Unknown parsers cannot be skipped:
// property lengths are not on the wire.
func commandParser(id ns.VarInt) (string, error) {
	name := registries.CommandArgumentType.ByID(int32(id))
	if name == "" {
		return "", fmt.Errorf("unsupported command parser ID %d", id)
	}
	return name, nil
}

func (p *CommandParserProperties) read(buf *ns.PacketBuffer, id ns.VarInt) error {
	parser, err := commandParser(id)
	if err != nil {
		return err
	}
	switch parser {
	case "brigadier:float", "brigadier:double", "brigadier:integer", "brigadier:long":
		if p.Flags, err = buf.ReadByte(); err != nil {
			return err
		}
		if p.Flags&^(CommandNumberHasMin|CommandNumberHasMax) != 0 {
			return errors.New("invalid numeric parser flags")
		}
		if p.Flags&CommandNumberHasMin != 0 {
			if err := readCommandBound(buf, parser, &p.IntMin, &p.FloatMin); err != nil {
				return err
			}
		}
		if p.Flags&CommandNumberHasMax != 0 {
			if err := readCommandBound(buf, parser, &p.IntMax, &p.FloatMax); err != nil {
				return err
			}
		}
	case "brigadier:string":
		if p.StringMode, err = buf.ReadVarInt(); err != nil {
			return err
		}
		if p.StringMode < 0 || p.StringMode > 2 {
			return errors.New("invalid string parser mode")
		}
	case "minecraft:entity":
		if p.Flags, err = buf.ReadByte(); err != nil {
			return err
		}
		if p.Flags&^(CommandEntitySingle|CommandEntityPlayersOnly) != 0 {
			return errors.New("invalid entity parser flags")
		}
	case "minecraft:score_holder":
		p.Flags, err = buf.ReadByte()
	case "minecraft:time":
		p.TimeMin, err = buf.ReadInt32()
	case "minecraft:resource_or_tag", "minecraft:resource_or_tag_key", "minecraft:resource", "minecraft:resource_key", "minecraft:resource_selector":
		p.Registry, err = buf.ReadIdentifier()
	}
	return err
}

func readCommandBound(buf *ns.PacketBuffer, parser string, i *int64, f *float64) error {
	switch parser {
	case "brigadier:float":
		v, err := buf.ReadFloat32()
		*f = float64(v)
		return err
	case "brigadier:double":
		v, err := buf.ReadFloat64()
		*f = float64(v)
		return err
	case "brigadier:integer":
		v, err := buf.ReadInt32()
		*i = int64(v)
		return err
	default: // brigadier:long
		v, err := buf.ReadInt64()
		*i = int64(v)
		return err
	}
}

func (p *CommandParserProperties) write(buf *ns.PacketBuffer, id ns.VarInt) error {
	parser, err := commandParser(id)
	if err != nil {
		return err
	}
	switch parser {
	case "brigadier:float", "brigadier:double", "brigadier:integer", "brigadier:long":
		if err := buf.WriteByte(p.Flags); err != nil {
			return err
		}
		if p.Flags&CommandNumberHasMin != 0 {
			if err := writeCommandBound(buf, parser, p.IntMin, p.FloatMin); err != nil {
				return err
			}
		}
		if p.Flags&CommandNumberHasMax != 0 {
			return writeCommandBound(buf, parser, p.IntMax, p.FloatMax)
		}
		return nil
	case "brigadier:string":
		return buf.WriteVarInt(p.StringMode)
	case "minecraft:entity", "minecraft:score_holder":
		return buf.WriteByte(p.Flags)
	case "minecraft:time":
		return buf.WriteInt32(p.TimeMin)
	case "minecraft:resource_or_tag", "minecraft:resource_or_tag_key", "minecraft:resource", "minecraft:resource_key", "minecraft:resource_selector":
		return buf.WriteIdentifier(p.Registry)
	}
	return nil
}

func writeCommandBound(buf *ns.PacketBuffer, parser string, i int64, f float64) error {
	switch parser {
	case "brigadier:float":
		return buf.WriteFloat32(ns.Float32(f))
	case "brigadier:double":
		return buf.WriteFloat64(ns.Float64(f))
	case "brigadier:integer":
		return buf.WriteInt32(ns.Int32(i))
	default: // brigadier:long
		return buf.WriteInt64(ns.Int64(i))
	}
}

package packets

import (
	"fmt"

	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// PlayerInfoActions is the EnumSet<ClientboundPlayerInfoUpdatePacket.Action>
// of a player_info_update: a FixedBitSet(8), one bit per action ordinal. Each
// entry carries the fields of every set action, in ordinal order.
type PlayerInfoActions byte

// Player info update actions, in vanilla ordinal (and wire) order.
const (
	PlayerInfoAddPlayer PlayerInfoActions = 1 << iota
	PlayerInfoInitializeChat
	PlayerInfoUpdateGameMode
	PlayerInfoUpdateListed
	PlayerInfoUpdateLatency
	PlayerInfoUpdateDisplayName
	PlayerInfoUpdateListOrder
	PlayerInfoUpdateHat
)

// Has reports whether every action in a is set.
func (s PlayerInfoActions) Has(a PlayerInfoActions) bool { return s&a == a }

// maxProfileProperties is ByteBufCodecs.GAME_PROFILE_PROPERTIES' list limit.
const maxProfileProperties = 16

// RemoteChatSession is RemoteChatSession.Data: the session id and the
// player's profile public key (ProfilePublicKey.Data).
type RemoteChatSession struct {
	SessionID    ns.UUID
	ExpiresAt    ns.Int64
	PublicKey    ns.ByteArray
	KeySignature ns.ByteArray
}

// PlayerInfoEntry is one player_info_update entry. Only the fields of the
// packet's actions are on the wire; the rest are ignored when writing and
// left zero when reading.
type PlayerInfoEntry struct {
	UUID ns.UUID
	// PlayerInfoAddPlayer
	Name       ns.String
	Properties []GameProfileProperty
	// PlayerInfoInitializeChat
	ChatSession ns.PrefixedOptional[RemoteChatSession]
	// PlayerInfoUpdateGameMode
	GameMode ns.VarInt
	// PlayerInfoUpdateListed
	Listed ns.Boolean
	// PlayerInfoUpdateLatency
	Latency ns.VarInt
	// PlayerInfoUpdateDisplayName
	DisplayName ns.PrefixedOptional[ns.TextComponent]
	// PlayerInfoUpdateListOrder
	ListOrder ns.VarInt
	// PlayerInfoUpdateHat
	ShowHat ns.Boolean
}

func (p *S2CPlayerInfoUpdate) read(buf *ns.PacketBuffer) error {
	actions, err := buf.ReadByte()
	if err != nil {
		return err
	}
	p.Actions = PlayerInfoActions(actions)
	count, err := readCount(buf, "player info entry", maxListCount)
	if err != nil {
		return err
	}
	p.Entries = make([]PlayerInfoEntry, count)
	for i := range p.Entries {
		if err := p.Entries[i].read(buf, p.Actions); err != nil {
			return err
		}
	}
	return nil
}

func (p *S2CPlayerInfoUpdate) write(buf *ns.PacketBuffer) error {
	if len(p.Entries) > maxListCount {
		return fmt.Errorf("player info entry count %d exceeds %d", len(p.Entries), maxListCount)
	}
	if err := buf.WriteByte(byte(p.Actions)); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(len(p.Entries))); err != nil {
		return err
	}
	for i := range p.Entries {
		if err := p.Entries[i].write(buf, p.Actions); err != nil {
			return err
		}
	}
	return nil
}

func (e *PlayerInfoEntry) read(buf *ns.PacketBuffer, actions PlayerInfoActions) error {
	var err error
	if e.UUID, err = buf.ReadUUID(); err != nil {
		return err
	}
	if actions.Has(PlayerInfoAddPlayer) {
		if e.Name, err = buf.ReadString(16); err != nil {
			return err
		}
		if e.Properties, err = readProfileProperties(buf); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoInitializeChat) {
		if err := e.ChatSession.DecodeWith(buf, readRemoteChatSession); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateGameMode) {
		if e.GameMode, err = buf.ReadVarInt(); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateListed) {
		if e.Listed, err = buf.ReadBool(); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateLatency) {
		if e.Latency, err = buf.ReadVarInt(); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateDisplayName) {
		if err := e.DisplayName.DecodeWith(buf, (*ns.PacketBuffer).ReadTextComponent); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateListOrder) {
		if e.ListOrder, err = buf.ReadVarInt(); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateHat) {
		if e.ShowHat, err = buf.ReadBool(); err != nil {
			return err
		}
	}
	return nil
}

func (e *PlayerInfoEntry) write(buf *ns.PacketBuffer, actions PlayerInfoActions) error {
	if err := buf.WriteUUID(e.UUID); err != nil {
		return err
	}
	if actions.Has(PlayerInfoAddPlayer) {
		if err := buf.WriteString(e.Name); err != nil {
			return err
		}
		if err := writeProfileProperties(buf, e.Properties); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoInitializeChat) {
		if err := e.ChatSession.EncodeWith(buf, writeRemoteChatSession); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateGameMode) {
		if err := buf.WriteVarInt(e.GameMode); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateListed) {
		if err := buf.WriteBool(e.Listed); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateLatency) {
		if err := buf.WriteVarInt(e.Latency); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateDisplayName) {
		if err := e.DisplayName.EncodeWith(buf, (*ns.PacketBuffer).WriteTextComponent); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateListOrder) {
		if err := buf.WriteVarInt(e.ListOrder); err != nil {
			return err
		}
	}
	if actions.Has(PlayerInfoUpdateHat) {
		if err := buf.WriteBool(e.ShowHat); err != nil {
			return err
		}
	}
	return nil
}

func readRemoteChatSession(buf *ns.PacketBuffer) (RemoteChatSession, error) {
	var s RemoteChatSession
	var err error
	if s.SessionID, err = buf.ReadUUID(); err != nil {
		return s, err
	}
	if s.ExpiresAt, err = buf.ReadInt64(); err != nil {
		return s, err
	}
	if s.PublicKey, err = buf.ReadByteArray(512); err != nil {
		return s, err
	}
	s.KeySignature, err = buf.ReadByteArray(4096)
	return s, err
}

func writeRemoteChatSession(buf *ns.PacketBuffer, s RemoteChatSession) error {
	if err := buf.WriteUUID(s.SessionID); err != nil {
		return err
	}
	if err := buf.WriteInt64(s.ExpiresAt); err != nil {
		return err
	}
	if err := buf.WriteByteArray(s.PublicKey); err != nil {
		return err
	}
	return buf.WriteByteArray(s.KeySignature)
}

// readProfileProperties reads ByteBufCodecs.GAME_PROFILE_PROPERTIES.
func readProfileProperties(buf *ns.PacketBuffer) ([]GameProfileProperty, error) {
	count, err := readCount(buf, "profile property", maxProfileProperties)
	if err != nil {
		return nil, err
	}
	props := make([]GameProfileProperty, count)
	for i := range props {
		if props[i].Name, err = buf.ReadString(64); err != nil {
			return nil, err
		}
		if props[i].Value, err = buf.ReadString(32767); err != nil {
			return nil, err
		}
		if err = props[i].Signature.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.String, error) {
			return b.ReadString(1024)
		}); err != nil {
			return nil, err
		}
	}
	return props, nil
}

// writeProfileProperties writes ByteBufCodecs.GAME_PROFILE_PROPERTIES.
func writeProfileProperties(buf *ns.PacketBuffer, props []GameProfileProperty) error {
	if len(props) > maxProfileProperties {
		return fmt.Errorf("profile property count %d exceeds %d", len(props), maxProfileProperties)
	}
	if err := buf.WriteVarInt(ns.VarInt(len(props))); err != nil {
		return err
	}
	for _, prop := range props {
		if err := buf.WriteString(prop.Name); err != nil {
			return err
		}
		if err := buf.WriteString(prop.Value); err != nil {
			return err
		}
		if err := prop.Signature.EncodeWith(buf, (*ns.PacketBuffer).WriteString); err != nil {
			return err
		}
	}
	return nil
}

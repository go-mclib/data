package packets

import (
	"github.com/go-mclib/data/pkg/data/items"
	packets_data "github.com/go-mclib/data/pkg/data/packets"
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
	"github.com/go-mclib/protocol/nbt"
)

// C2SAcceptTeleportation represents "Confirm Teleportation".
//
// Sent by client as confirmation of Synchronize Player Position.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Confirm_Teleportation
type C2SAcceptTeleportation struct {
	TeleportId ns.VarInt
}

func (p *C2SAcceptTeleportation) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SAcceptTeleportationID)
}
func (p *C2SAcceptTeleportation) State() jp.State { return jp.StatePlay }
func (p *C2SAcceptTeleportation) Bound() jp.Bound { return jp.C2S }

func (p *C2SAcceptTeleportation) Read(buf *ns.PacketBuffer) error {
	var err error
	p.TeleportId, err = buf.ReadVarInt()
	return err
}

func (p *C2SAcceptTeleportation) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.TeleportId)
}

// C2SBlockEntityTagQuery represents "Query Block Entity Tag".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Query_Block_Entity_Tag
type C2SBlockEntityTagQuery struct {
	TransactionId ns.VarInt
	Location      ns.Position
}

func (p *C2SBlockEntityTagQuery) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SBlockEntityTagQueryID)
}
func (p *C2SBlockEntityTagQuery) State() jp.State { return jp.StatePlay }
func (p *C2SBlockEntityTagQuery) Bound() jp.Bound { return jp.C2S }

func (p *C2SBlockEntityTagQuery) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.TransactionId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Location, err = buf.ReadPosition()
	return err
}

func (p *C2SBlockEntityTagQuery) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.TransactionId); err != nil {
		return err
	}
	return buf.WritePosition(p.Location)
}

// C2SBundleItemSelected represents "Bundle Item Selected".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Bundle_Item_Selected
type C2SBundleItemSelected struct {
	SlotOfBundle ns.VarInt
	SlotInBundle ns.VarInt
}

func (p *C2SBundleItemSelected) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SBundleItemSelectedID)
}
func (p *C2SBundleItemSelected) State() jp.State { return jp.StatePlay }
func (p *C2SBundleItemSelected) Bound() jp.Bound { return jp.C2S }

func (p *C2SBundleItemSelected) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.SlotOfBundle, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.SlotInBundle, err = buf.ReadVarInt()
	return err
}

func (p *C2SBundleItemSelected) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.SlotOfBundle); err != nil {
		return err
	}
	return buf.WriteVarInt(p.SlotInBundle)
}

// C2SChangeDifficulty represents "Change Difficulty".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Change_Difficulty
type C2SChangeDifficulty struct {
	NewDifficulty ns.Uint8
}

func (p *C2SChangeDifficulty) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SChangeDifficultyID) }
func (p *C2SChangeDifficulty) State() jp.State { return jp.StatePlay }
func (p *C2SChangeDifficulty) Bound() jp.Bound { return jp.C2S }

func (p *C2SChangeDifficulty) Read(buf *ns.PacketBuffer) error {
	var err error
	p.NewDifficulty, err = buf.ReadUint8()
	return err
}

func (p *C2SChangeDifficulty) Write(buf *ns.PacketBuffer) error {
	return buf.WriteUint8(p.NewDifficulty)
}

// C2SChangeGameMode represents "Change Game Mode".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Change_Game_Mode
type C2SChangeGameMode struct {
	GameMode ns.VarInt
}

func (p *C2SChangeGameMode) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SChangeGameModeID) }
func (p *C2SChangeGameMode) State() jp.State { return jp.StatePlay }
func (p *C2SChangeGameMode) Bound() jp.Bound { return jp.C2S }

func (p *C2SChangeGameMode) Read(buf *ns.PacketBuffer) error {
	var err error
	p.GameMode, err = buf.ReadVarInt()
	return err
}

func (p *C2SChangeGameMode) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.GameMode)
}

// C2SChatAck represents "Acknowledge Message".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Acknowledge_Message
type C2SChatAck struct {
	MessageCount ns.VarInt
}

func (p *C2SChatAck) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SChatAckID) }
func (p *C2SChatAck) State() jp.State { return jp.StatePlay }
func (p *C2SChatAck) Bound() jp.Bound { return jp.C2S }

func (p *C2SChatAck) Read(buf *ns.PacketBuffer) error {
	var err error
	p.MessageCount, err = buf.ReadVarInt()
	return err
}

func (p *C2SChatAck) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.MessageCount)
}

// C2SChatCommand represents "Chat Command".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chat_Command
type C2SChatCommand struct {
	Command ns.String
}

func (p *C2SChatCommand) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SChatCommandID) }
func (p *C2SChatCommand) State() jp.State { return jp.StatePlay }
func (p *C2SChatCommand) Bound() jp.Bound { return jp.C2S }

func (p *C2SChatCommand) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Command, err = buf.ReadString(32767)
	return err
}

func (p *C2SChatCommand) Write(buf *ns.PacketBuffer) error {
	return buf.WriteString(p.Command)
}

// C2SChatCommandSigned represents "Signed Chat Command".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Signed_Chat_Command
type C2SChatCommandSigned struct {
	Command      ns.String
	Timestamp    ns.Int64
	Salt         ns.Int64
	Signature    ns.ByteArray
	MessageCount ns.VarInt
	Acknowledged *ns.FixedBitSet
	Checksum     ns.Int8
}

func (p *C2SChatCommandSigned) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SChatCommandSignedID) }
func (p *C2SChatCommandSigned) State() jp.State { return jp.StatePlay }
func (p *C2SChatCommandSigned) Bound() jp.Bound { return jp.C2S }

func (p *C2SChatCommandSigned) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Command, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.Timestamp, err = buf.ReadInt64(); err != nil {
		return err
	}
	if p.Salt, err = buf.ReadInt64(); err != nil {
		return err
	}
	if p.Signature, err = buf.ReadByteArray(256); err != nil {
		return err
	}
	if p.MessageCount, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Acknowledged = ns.NewFixedBitSet(20)
	ackBytes, err := buf.ReadFixedByteArray(3)
	if err != nil {
		return err
	}
	p.Acknowledged = ns.FixedBitSetFromBytes(ackBytes, 20)
	p.Checksum, err = buf.ReadInt8()
	return err
}

func (p *C2SChatCommandSigned) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteString(p.Command); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.Timestamp); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.Salt); err != nil {
		return err
	}
	if err := buf.WriteByteArray(p.Signature); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.MessageCount); err != nil {
		return err
	}
	if err := buf.WriteFixedByteArray(p.Acknowledged.Bytes()); err != nil {
		return err
	}
	return buf.WriteInt8(p.Checksum)
}

// C2SChat represents "Chat Message".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chat_Message
type C2SChat struct {
	Message      ns.String
	Timestamp    ns.Int64
	Salt         ns.Int64
	Signature    ns.PrefixedOptional[ns.ByteArray]
	MessageCount ns.VarInt
	Acknowledged *ns.FixedBitSet
	Checksum     ns.Int8
}

func (p *C2SChat) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SChatID) }
func (p *C2SChat) State() jp.State { return jp.StatePlay }
func (p *C2SChat) Bound() jp.Bound { return jp.C2S }

func (p *C2SChat) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Message, err = buf.ReadString(256); err != nil {
		return err
	}
	if p.Timestamp, err = buf.ReadInt64(); err != nil {
		return err
	}
	if p.Salt, err = buf.ReadInt64(); err != nil {
		return err
	}
	if err = p.Signature.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.ByteArray, error) {
		return b.ReadFixedByteArray(256)
	}); err != nil {
		return err
	}
	if p.MessageCount, err = buf.ReadVarInt(); err != nil {
		return err
	}
	ackBytes, err := buf.ReadFixedByteArray(3)
	if err != nil {
		return err
	}
	p.Acknowledged = ns.FixedBitSetFromBytes(ackBytes, 20)
	p.Checksum, err = buf.ReadInt8()
	return err
}

func (p *C2SChat) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteString(p.Message); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.Timestamp); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.Salt); err != nil {
		return err
	}
	if err := p.Signature.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.ByteArray) error {
		return b.WriteFixedByteArray(v)
	}); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.MessageCount); err != nil {
		return err
	}
	if err := buf.WriteFixedByteArray(p.Acknowledged.Bytes()); err != nil {
		return err
	}
	return buf.WriteInt8(p.Checksum)
}

// C2SChatSessionUpdate represents "Player Session".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Session
type C2SChatSessionUpdate struct {
	SessionId    ns.UUID
	ExpiresAt    ns.Int64
	PublicKey    ns.ByteArray
	KeySignature ns.ByteArray
}

func (p *C2SChatSessionUpdate) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SChatSessionUpdateID) }
func (p *C2SChatSessionUpdate) State() jp.State { return jp.StatePlay }
func (p *C2SChatSessionUpdate) Bound() jp.Bound { return jp.C2S }

func (p *C2SChatSessionUpdate) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.SessionId, err = buf.ReadUUID(); err != nil {
		return err
	}
	if p.ExpiresAt, err = buf.ReadInt64(); err != nil {
		return err
	}
	if p.PublicKey, err = buf.ReadByteArray(512); err != nil {
		return err
	}
	p.KeySignature, err = buf.ReadByteArray(4096)
	return err
}

func (p *C2SChatSessionUpdate) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteUUID(p.SessionId); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.ExpiresAt); err != nil {
		return err
	}
	if err := buf.WriteByteArray(p.PublicKey); err != nil {
		return err
	}
	return buf.WriteByteArray(p.KeySignature)
}

// C2SChunkBatchReceived represents "Chunk Batch Received".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chunk_Batch_Received
type C2SChunkBatchReceived struct {
	ChunksPerTick ns.Float32
}

func (p *C2SChunkBatchReceived) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SChunkBatchReceivedID)
}
func (p *C2SChunkBatchReceived) State() jp.State { return jp.StatePlay }
func (p *C2SChunkBatchReceived) Bound() jp.Bound { return jp.C2S }

func (p *C2SChunkBatchReceived) Read(buf *ns.PacketBuffer) error {
	var err error
	p.ChunksPerTick, err = buf.ReadFloat32()
	return err
}

func (p *C2SChunkBatchReceived) Write(buf *ns.PacketBuffer) error {
	return buf.WriteFloat32(p.ChunksPerTick)
}

// C2SClientCommand represents "Client Status".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Client_Status
type C2SClientCommand struct {
	ActionId ns.VarInt
}

func (p *C2SClientCommand) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SClientCommandID) }
func (p *C2SClientCommand) State() jp.State { return jp.StatePlay }
func (p *C2SClientCommand) Bound() jp.Bound { return jp.C2S }

func (p *C2SClientCommand) Read(buf *ns.PacketBuffer) error {
	var err error
	p.ActionId, err = buf.ReadVarInt()
	return err
}

func (p *C2SClientCommand) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.ActionId)
}

// C2SClientTickEnd represents "Client Tick End".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Client_Tick_End
type C2SClientTickEnd struct{}

func (p *C2SClientTickEnd) ID() ns.VarInt                { return ns.VarInt(packets_data.C2SClientTickEndID) }
func (p *C2SClientTickEnd) State() jp.State              { return jp.StatePlay }
func (p *C2SClientTickEnd) Bound() jp.Bound              { return jp.C2S }
func (p *C2SClientTickEnd) Read(*ns.PacketBuffer) error  { return nil }
func (p *C2SClientTickEnd) Write(*ns.PacketBuffer) error { return nil }

// C2SClientInformationPlay represents "Client Information (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Client_Information_(Play)
type C2SClientInformationPlay struct {
	Locale              ns.String
	ViewDistance        ns.Int8
	ChatMode            ns.VarInt
	ChatColors          ns.Boolean
	DisplayedSkinParts  ns.Uint8
	MainHand            ns.VarInt
	EnableTextFiltering ns.Boolean
	AllowServerListings ns.Boolean
	ParticleStatus      ns.VarInt
}

func (p *C2SClientInformationPlay) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SClientInformationPlayID)
}
func (p *C2SClientInformationPlay) State() jp.State { return jp.StatePlay }
func (p *C2SClientInformationPlay) Bound() jp.Bound { return jp.C2S }

func (p *C2SClientInformationPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Locale, err = buf.ReadString(16); err != nil {
		return err
	}
	if p.ViewDistance, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.ChatMode, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.ChatColors, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.DisplayedSkinParts, err = buf.ReadUint8(); err != nil {
		return err
	}
	if p.MainHand, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.EnableTextFiltering, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.AllowServerListings, err = buf.ReadBool(); err != nil {
		return err
	}
	p.ParticleStatus, err = buf.ReadVarInt()
	return err
}

func (p *C2SClientInformationPlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteString(p.Locale); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.ViewDistance); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.ChatMode); err != nil {
		return err
	}
	if err := buf.WriteBool(p.ChatColors); err != nil {
		return err
	}
	if err := buf.WriteUint8(p.DisplayedSkinParts); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.MainHand); err != nil {
		return err
	}
	if err := buf.WriteBool(p.EnableTextFiltering); err != nil {
		return err
	}
	if err := buf.WriteBool(p.AllowServerListings); err != nil {
		return err
	}
	return buf.WriteVarInt(p.ParticleStatus)
}

// C2SCommandSuggestion represents "Command Suggestions Request".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Command_Suggestions_Request
type C2SCommandSuggestion struct {
	TransactionId ns.VarInt
	Text          ns.String
}

func (p *C2SCommandSuggestion) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SCommandSuggestionID) }
func (p *C2SCommandSuggestion) State() jp.State { return jp.StatePlay }
func (p *C2SCommandSuggestion) Bound() jp.Bound { return jp.C2S }

func (p *C2SCommandSuggestion) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.TransactionId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Text, err = buf.ReadString(32500)
	return err
}

func (p *C2SCommandSuggestion) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.TransactionId); err != nil {
		return err
	}
	return buf.WriteString(p.Text)
}

// C2SConfigurationAcknowledged represents "Acknowledge Configuration".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Acknowledge_Configuration
type C2SConfigurationAcknowledged struct{}

func (p *C2SConfigurationAcknowledged) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SConfigurationAcknowledgedID)
}
func (p *C2SConfigurationAcknowledged) State() jp.State              { return jp.StatePlay }
func (p *C2SConfigurationAcknowledged) Bound() jp.Bound              { return jp.C2S }
func (p *C2SConfigurationAcknowledged) Read(*ns.PacketBuffer) error  { return nil }
func (p *C2SConfigurationAcknowledged) Write(*ns.PacketBuffer) error { return nil }

// C2SContainerButtonClick represents "Click Container Button".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Click_Container_Button
type C2SContainerButtonClick struct {
	WindowId ns.VarInt
	ButtonId ns.VarInt
}

func (p *C2SContainerButtonClick) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SContainerButtonClickID)
}
func (p *C2SContainerButtonClick) State() jp.State { return jp.StatePlay }
func (p *C2SContainerButtonClick) Bound() jp.Bound { return jp.C2S }

func (p *C2SContainerButtonClick) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.ButtonId, err = buf.ReadVarInt()
	return err
}

func (p *C2SContainerButtonClick) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	return buf.WriteVarInt(p.ButtonId)
}

// C2SContainerClick represents "Click Container".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Click_Container
type C2SContainerClick struct {
	WindowId    ns.VarInt
	StateId     ns.VarInt
	Slot        ns.Int16
	Button      ns.Int8
	Mode        ns.VarInt
	SlotData    ns.Slot
	CarriedItem ns.Slot
}

func (p *C2SContainerClick) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SContainerClickID) }
func (p *C2SContainerClick) State() jp.State { return jp.StatePlay }
func (p *C2SContainerClick) Bound() jp.Bound { return jp.C2S }

func (p *C2SContainerClick) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.StateId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Slot, err = buf.ReadInt16(); err != nil {
		return err
	}
	if p.Button, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.Mode, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.SlotData, err = buf.ReadSlot(items.Decoder()); err != nil {
		return err
	}
	p.CarriedItem, err = buf.ReadSlot(items.Decoder())
	return err
}

func (p *C2SContainerClick) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.StateId); err != nil {
		return err
	}
	if err := buf.WriteInt16(p.Slot); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.Button); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Mode); err != nil {
		return err
	}
	if err := buf.WriteSlot(p.SlotData); err != nil {
		return err
	}
	return buf.WriteSlot(p.CarriedItem)
}

// C2SContainerClose represents "Close Container".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Close_Container
type C2SContainerClose struct {
	WindowId ns.VarInt
}

func (p *C2SContainerClose) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SContainerCloseID) }
func (p *C2SContainerClose) State() jp.State { return jp.StatePlay }
func (p *C2SContainerClose) Bound() jp.Bound { return jp.C2S }

func (p *C2SContainerClose) Read(buf *ns.PacketBuffer) error {
	var err error
	p.WindowId, err = buf.ReadVarInt()
	return err
}

func (p *C2SContainerClose) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.WindowId)
}

// C2SContainerSlotStateChanged represents "Change Container Slot State".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Change_Container_Slot_State
type C2SContainerSlotStateChanged struct {
	SlotId      ns.VarInt
	WindowId    ns.VarInt
	SlotEnabled ns.Boolean
}

func (p *C2SContainerSlotStateChanged) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SContainerSlotStateChangedID)
}
func (p *C2SContainerSlotStateChanged) State() jp.State { return jp.StatePlay }
func (p *C2SContainerSlotStateChanged) Bound() jp.Bound { return jp.C2S }

func (p *C2SContainerSlotStateChanged) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.SlotId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.SlotEnabled, err = buf.ReadBool()
	return err
}

func (p *C2SContainerSlotStateChanged) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.SlotId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	return buf.WriteBool(p.SlotEnabled)
}

// C2SCookieResponsePlay represents "Cookie Response (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Cookie_Response_(Play)
type C2SCookieResponsePlay struct {
	Key     ns.Identifier
	Payload ns.PrefixedOptional[ns.ByteArray]
}

func (p *C2SCookieResponsePlay) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SCookieResponsePlayID)
}
func (p *C2SCookieResponsePlay) State() jp.State { return jp.StatePlay }
func (p *C2SCookieResponsePlay) Bound() jp.Bound { return jp.C2S }

func (p *C2SCookieResponsePlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Key, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	return p.Payload.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.ByteArray, error) {
		return b.ReadByteArray(5120)
	})
}

func (p *C2SCookieResponsePlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteIdentifier(p.Key); err != nil {
		return err
	}
	return p.Payload.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.ByteArray) error {
		return b.WriteByteArray(v)
	})
}

// C2SCustomPayloadPlay represents "Serverbound Plugin Message (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Serverbound_Plugin_Message_(Play)
type C2SCustomPayloadPlay struct {
	Channel ns.Identifier
	Data    ns.ByteArray
}

func (p *C2SCustomPayloadPlay) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SCustomPayloadPlayID) }
func (p *C2SCustomPayloadPlay) State() jp.State { return jp.StatePlay }
func (p *C2SCustomPayloadPlay) Bound() jp.Bound { return jp.C2S }

func (p *C2SCustomPayloadPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Channel, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(32767)
	return err
}

func (p *C2SCustomPayloadPlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteIdentifier(p.Channel); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// C2SDebugSubscriptionRequest represents "Debug Subscription Request".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Debug_Subscription_Request
type C2SDebugSubscriptionRequest struct {
	Subscriptions []ns.VarInt
}

func (p *C2SDebugSubscriptionRequest) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SDebugSubscriptionRequestID)
}
func (p *C2SDebugSubscriptionRequest) State() jp.State { return jp.StatePlay }
func (p *C2SDebugSubscriptionRequest) Bound() jp.Bound { return jp.C2S }

func (p *C2SDebugSubscriptionRequest) Read(buf *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	p.Subscriptions = make([]ns.VarInt, count)
	for i := range p.Subscriptions {
		if p.Subscriptions[i], err = buf.ReadVarInt(); err != nil {
			return err
		}
	}
	return nil
}

func (p *C2SDebugSubscriptionRequest) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(ns.VarInt(len(p.Subscriptions))); err != nil {
		return err
	}
	for _, sub := range p.Subscriptions {
		if err := buf.WriteVarInt(sub); err != nil {
			return err
		}
	}
	return nil
}

// C2SEditBook represents "Edit Book".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Edit_Book
type C2SEditBook struct {
	Slot    ns.VarInt
	Entries []ns.String
	Title   ns.PrefixedOptional[ns.String]
}

func (p *C2SEditBook) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SEditBookID) }
func (p *C2SEditBook) State() jp.State { return jp.StatePlay }
func (p *C2SEditBook) Bound() jp.Bound { return jp.C2S }

func (p *C2SEditBook) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Slot, err = buf.ReadVarInt(); err != nil {
		return err
	}
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	p.Entries = make([]ns.String, count)
	for i := range p.Entries {
		if p.Entries[i], err = buf.ReadString(8192); err != nil {
			return err
		}
	}
	return p.Title.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.String, error) {
		return b.ReadString(128)
	})
}

func (p *C2SEditBook) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.Slot); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(len(p.Entries))); err != nil {
		return err
	}
	for _, entry := range p.Entries {
		if err := buf.WriteString(entry); err != nil {
			return err
		}
	}
	return p.Title.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.String) error {
		return b.WriteString(v)
	})
}

// C2SEntityTagQuery represents "Query Entity Tag".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Query_Entity_Tag
type C2SEntityTagQuery struct {
	TransactionId ns.VarInt
	EntityId      ns.VarInt
}

func (p *C2SEntityTagQuery) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SEntityTagQueryID) }
func (p *C2SEntityTagQuery) State() jp.State { return jp.StatePlay }
func (p *C2SEntityTagQuery) Bound() jp.Bound { return jp.C2S }

func (p *C2SEntityTagQuery) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.TransactionId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.EntityId, err = buf.ReadVarInt()
	return err
}

func (p *C2SEntityTagQuery) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.TransactionId); err != nil {
		return err
	}
	return buf.WriteVarInt(p.EntityId)
}

// C2SInteract represents "Interact".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Interact
type C2SInteract struct {
	EntityId        ns.VarInt
	Type            ns.VarInt
	TargetX         ns.Float32 // only if Type is 2 (interact at)
	TargetY         ns.Float32 // only if Type is 2
	TargetZ         ns.Float32 // only if Type is 2
	Hand            ns.VarInt  // only if Type is 0 or 2
	SneakKeyPressed ns.Boolean
}

func (p *C2SInteract) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SInteractID) }
func (p *C2SInteract) State() jp.State { return jp.StatePlay }
func (p *C2SInteract) Bound() jp.Bound { return jp.C2S }

func (p *C2SInteract) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Type, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Type == 2 {
		if p.TargetX, err = buf.ReadFloat32(); err != nil {
			return err
		}
		if p.TargetY, err = buf.ReadFloat32(); err != nil {
			return err
		}
		if p.TargetZ, err = buf.ReadFloat32(); err != nil {
			return err
		}
	}
	if p.Type == 0 || p.Type == 2 {
		if p.Hand, err = buf.ReadVarInt(); err != nil {
			return err
		}
	}
	p.SneakKeyPressed, err = buf.ReadBool()
	return err
}

func (p *C2SInteract) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Type); err != nil {
		return err
	}
	if p.Type == 2 {
		if err := buf.WriteFloat32(p.TargetX); err != nil {
			return err
		}
		if err := buf.WriteFloat32(p.TargetY); err != nil {
			return err
		}
		if err := buf.WriteFloat32(p.TargetZ); err != nil {
			return err
		}
	}
	if p.Type == 0 || p.Type == 2 {
		if err := buf.WriteVarInt(p.Hand); err != nil {
			return err
		}
	}
	return buf.WriteBool(p.SneakKeyPressed)
}

// C2SJigsawGenerate represents "Jigsaw Generate".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Jigsaw_Generate
type C2SJigsawGenerate struct {
	Location    ns.Position
	Levels      ns.VarInt
	KeepJigsaws ns.Boolean
}

func (p *C2SJigsawGenerate) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SJigsawGenerateID) }
func (p *C2SJigsawGenerate) State() jp.State { return jp.StatePlay }
func (p *C2SJigsawGenerate) Bound() jp.Bound { return jp.C2S }

func (p *C2SJigsawGenerate) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Levels, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.KeepJigsaws, err = buf.ReadBool()
	return err
}

func (p *C2SJigsawGenerate) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Levels); err != nil {
		return err
	}
	return buf.WriteBool(p.KeepJigsaws)
}

// C2SKeepAlivePlay represents "Serverbound Keep Alive (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Serverbound_Keep_Alive_(Play)
type C2SKeepAlivePlay struct {
	KeepAliveId ns.Int64
}

func (p *C2SKeepAlivePlay) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SKeepAlivePlayID) }
func (p *C2SKeepAlivePlay) State() jp.State { return jp.StatePlay }
func (p *C2SKeepAlivePlay) Bound() jp.Bound { return jp.C2S }

func (p *C2SKeepAlivePlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.KeepAliveId, err = buf.ReadInt64()
	return err
}

func (p *C2SKeepAlivePlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt64(p.KeepAliveId)
}

// C2SLockDifficulty represents "Lock Difficulty".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Lock_Difficulty
type C2SLockDifficulty struct {
	Locked ns.Boolean
}

func (p *C2SLockDifficulty) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SLockDifficultyID) }
func (p *C2SLockDifficulty) State() jp.State { return jp.StatePlay }
func (p *C2SLockDifficulty) Bound() jp.Bound { return jp.C2S }

func (p *C2SLockDifficulty) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Locked, err = buf.ReadBool()
	return err
}

func (p *C2SLockDifficulty) Write(buf *ns.PacketBuffer) error {
	return buf.WriteBool(p.Locked)
}

// C2SMovePlayerPos represents "Set Player Position".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Player_Position
type C2SMovePlayerPos struct {
	X     ns.Float64
	FeetY ns.Float64
	Z     ns.Float64
	Flags ns.Int8
}

func (p *C2SMovePlayerPos) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SMovePlayerPosID) }
func (p *C2SMovePlayerPos) State() jp.State { return jp.StatePlay }
func (p *C2SMovePlayerPos) Bound() jp.Bound { return jp.C2S }

func (p *C2SMovePlayerPos) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.FeetY, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Z, err = buf.ReadFloat64(); err != nil {
		return err
	}
	p.Flags, err = buf.ReadInt8()
	return err
}

func (p *C2SMovePlayerPos) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.FeetY); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	return buf.WriteInt8(p.Flags)
}

// C2SMovePlayerPosRot represents "Set Player Position and Rotation".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Player_Position_And_Rotation
type C2SMovePlayerPosRot struct {
	X     ns.Float64
	FeetY ns.Float64
	Z     ns.Float64
	Yaw   ns.Float32
	Pitch ns.Float32
	Flags ns.Int8
}

func (p *C2SMovePlayerPosRot) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SMovePlayerPosRotID) }
func (p *C2SMovePlayerPosRot) State() jp.State { return jp.StatePlay }
func (p *C2SMovePlayerPosRot) Bound() jp.Bound { return jp.C2S }

func (p *C2SMovePlayerPosRot) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.FeetY, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Z, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Yaw, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.Flags, err = buf.ReadInt8()
	return err
}

func (p *C2SMovePlayerPosRot) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.FeetY); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Yaw); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Pitch); err != nil {
		return err
	}
	return buf.WriteInt8(p.Flags)
}

// C2SMovePlayerRot represents "Set Player Rotation".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Player_Rotation
type C2SMovePlayerRot struct {
	Yaw   ns.Float32
	Pitch ns.Float32
	Flags ns.Int8
}

func (p *C2SMovePlayerRot) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SMovePlayerRotID) }
func (p *C2SMovePlayerRot) State() jp.State { return jp.StatePlay }
func (p *C2SMovePlayerRot) Bound() jp.Bound { return jp.C2S }

func (p *C2SMovePlayerRot) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Yaw, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.Flags, err = buf.ReadInt8()
	return err
}

func (p *C2SMovePlayerRot) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat32(p.Yaw); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Pitch); err != nil {
		return err
	}
	return buf.WriteInt8(p.Flags)
}

// C2SMovePlayerStatusOnly represents "Set Player Movement Flags".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Player_Movement_Flags
type C2SMovePlayerStatusOnly struct {
	Flags ns.Int8
}

func (p *C2SMovePlayerStatusOnly) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SMovePlayerStatusOnlyID)
}
func (p *C2SMovePlayerStatusOnly) State() jp.State { return jp.StatePlay }
func (p *C2SMovePlayerStatusOnly) Bound() jp.Bound { return jp.C2S }

func (p *C2SMovePlayerStatusOnly) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Flags, err = buf.ReadInt8()
	return err
}

func (p *C2SMovePlayerStatusOnly) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt8(p.Flags)
}

// C2SMoveVehicle represents "Move Vehicle (serverbound)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Move_Vehicle_(Serverbound)
type C2SMoveVehicle struct {
	X        ns.Float64
	Y        ns.Float64
	Z        ns.Float64
	Yaw      ns.Float32
	Pitch    ns.Float32
	OnGround ns.Boolean
}

func (p *C2SMoveVehicle) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SMoveVehicleID) }
func (p *C2SMoveVehicle) State() jp.State { return jp.StatePlay }
func (p *C2SMoveVehicle) Bound() jp.Bound { return jp.C2S }

func (p *C2SMoveVehicle) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Y, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Z, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Yaw, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.OnGround, err = buf.ReadBool()
	return err
}

func (p *C2SMoveVehicle) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Y); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Yaw); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Pitch); err != nil {
		return err
	}
	return buf.WriteBool(p.OnGround)
}

// C2SPaddleBoat represents "Paddle Boat".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Paddle_Boat
type C2SPaddleBoat struct {
	LeftPaddleTurning  ns.Boolean
	RightPaddleTurning ns.Boolean
}

func (p *C2SPaddleBoat) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SPaddleBoatID) }
func (p *C2SPaddleBoat) State() jp.State { return jp.StatePlay }
func (p *C2SPaddleBoat) Bound() jp.Bound { return jp.C2S }

func (p *C2SPaddleBoat) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.LeftPaddleTurning, err = buf.ReadBool(); err != nil {
		return err
	}
	p.RightPaddleTurning, err = buf.ReadBool()
	return err
}

func (p *C2SPaddleBoat) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteBool(p.LeftPaddleTurning); err != nil {
		return err
	}
	return buf.WriteBool(p.RightPaddleTurning)
}

// C2SPickItemFromBlock represents "Pick Item From Block".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Pick_Item_From_Block
type C2SPickItemFromBlock struct {
	Location    ns.Position
	IncludeData ns.Boolean
}

func (p *C2SPickItemFromBlock) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SPickItemFromBlockID) }
func (p *C2SPickItemFromBlock) State() jp.State { return jp.StatePlay }
func (p *C2SPickItemFromBlock) Bound() jp.Bound { return jp.C2S }

func (p *C2SPickItemFromBlock) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	p.IncludeData, err = buf.ReadBool()
	return err
}

func (p *C2SPickItemFromBlock) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	return buf.WriteBool(p.IncludeData)
}

// C2SPickItemFromEntity represents "Pick Item From Entity".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Pick_Item_From_Entity
type C2SPickItemFromEntity struct {
	EntityId    ns.VarInt
	IncludeData ns.Boolean
}

func (p *C2SPickItemFromEntity) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SPickItemFromEntityID)
}
func (p *C2SPickItemFromEntity) State() jp.State { return jp.StatePlay }
func (p *C2SPickItemFromEntity) Bound() jp.Bound { return jp.C2S }

func (p *C2SPickItemFromEntity) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.IncludeData, err = buf.ReadBool()
	return err
}

func (p *C2SPickItemFromEntity) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteBool(p.IncludeData)
}

// C2SPingRequestPlay represents "Ping Request (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Ping_Request_(Play)
type C2SPingRequestPlay struct {
	Payload ns.Int64
}

func (p *C2SPingRequestPlay) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SPingRequestPlayID) }
func (p *C2SPingRequestPlay) State() jp.State { return jp.StatePlay }
func (p *C2SPingRequestPlay) Bound() jp.Bound { return jp.C2S }

func (p *C2SPingRequestPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Payload, err = buf.ReadInt64()
	return err
}

func (p *C2SPingRequestPlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt64(p.Payload)
}

// C2SPlaceRecipe represents "Place Recipe".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Place_Recipe
type C2SPlaceRecipe struct {
	WindowId ns.VarInt
	RecipeId ns.VarInt
	MakeAll  ns.Boolean
}

func (p *C2SPlaceRecipe) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SPlaceRecipeID) }
func (p *C2SPlaceRecipe) State() jp.State { return jp.StatePlay }
func (p *C2SPlaceRecipe) Bound() jp.Bound { return jp.C2S }

func (p *C2SPlaceRecipe) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.RecipeId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.MakeAll, err = buf.ReadBool()
	return err
}

func (p *C2SPlaceRecipe) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.RecipeId); err != nil {
		return err
	}
	return buf.WriteBool(p.MakeAll)
}

// C2SPlayerAbilities represents "Player Abilities (serverbound)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Abilities_(Serverbound)
type C2SPlayerAbilities struct {
	Flags ns.Int8
}

func (p *C2SPlayerAbilities) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SPlayerAbilitiesID) }
func (p *C2SPlayerAbilities) State() jp.State { return jp.StatePlay }
func (p *C2SPlayerAbilities) Bound() jp.Bound { return jp.C2S }

func (p *C2SPlayerAbilities) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Flags, err = buf.ReadInt8()
	return err
}

func (p *C2SPlayerAbilities) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt8(p.Flags)
}

// C2SPlayerAction represents "Player Action".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Action
type C2SPlayerAction struct {
	Status   ns.VarInt
	Location ns.Position
	Face     ns.Int8
	Sequence ns.VarInt
}

func (p *C2SPlayerAction) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SPlayerActionID) }
func (p *C2SPlayerAction) State() jp.State { return jp.StatePlay }
func (p *C2SPlayerAction) Bound() jp.Bound { return jp.C2S }

func (p *C2SPlayerAction) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Status, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Face, err = buf.ReadInt8(); err != nil {
		return err
	}
	p.Sequence, err = buf.ReadVarInt()
	return err
}

func (p *C2SPlayerAction) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.Status); err != nil {
		return err
	}
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.Face); err != nil {
		return err
	}
	return buf.WriteVarInt(p.Sequence)
}

// C2SPlayerCommand represents "Player Command".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Command
type C2SPlayerCommand struct {
	EntityId  ns.VarInt
	ActionId  ns.VarInt
	JumpBoost ns.VarInt
}

func (p *C2SPlayerCommand) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SPlayerCommandID) }
func (p *C2SPlayerCommand) State() jp.State { return jp.StatePlay }
func (p *C2SPlayerCommand) Bound() jp.Bound { return jp.C2S }

func (p *C2SPlayerCommand) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.ActionId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.JumpBoost, err = buf.ReadVarInt()
	return err
}

func (p *C2SPlayerCommand) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.ActionId); err != nil {
		return err
	}
	return buf.WriteVarInt(p.JumpBoost)
}

// C2SPlayerInput represents "Player Input".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Input
type C2SPlayerInput struct {
	Flags ns.Uint8
}

func (p *C2SPlayerInput) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SPlayerInputID) }
func (p *C2SPlayerInput) State() jp.State { return jp.StatePlay }
func (p *C2SPlayerInput) Bound() jp.Bound { return jp.C2S }

func (p *C2SPlayerInput) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Flags, err = buf.ReadUint8()
	return err
}

func (p *C2SPlayerInput) Write(buf *ns.PacketBuffer) error {
	return buf.WriteUint8(p.Flags)
}

// C2SPlayerLoaded represents "Player Loaded".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Loaded
type C2SPlayerLoaded struct{}

func (p *C2SPlayerLoaded) ID() ns.VarInt                { return ns.VarInt(packets_data.C2SPlayerLoadedID) }
func (p *C2SPlayerLoaded) State() jp.State              { return jp.StatePlay }
func (p *C2SPlayerLoaded) Bound() jp.Bound              { return jp.C2S }
func (p *C2SPlayerLoaded) Read(*ns.PacketBuffer) error  { return nil }
func (p *C2SPlayerLoaded) Write(*ns.PacketBuffer) error { return nil }

// C2SPongPlay represents "Pong (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Pong_(Play)
type C2SPongPlay struct {
	Id ns.Int32
}

func (p *C2SPongPlay) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SPongPlayID) }
func (p *C2SPongPlay) State() jp.State { return jp.StatePlay }
func (p *C2SPongPlay) Bound() jp.Bound { return jp.C2S }

func (p *C2SPongPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Id, err = buf.ReadInt32()
	return err
}

func (p *C2SPongPlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt32(p.Id)
}

// C2SRecipeBookChangeSettings represents "Change Recipe Book Settings".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Change_Recipe_Book_Settings
type C2SRecipeBookChangeSettings struct {
	BookId       ns.VarInt
	BookOpen     ns.Boolean
	FilterActive ns.Boolean
}

func (p *C2SRecipeBookChangeSettings) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SRecipeBookChangeSettingsID)
}
func (p *C2SRecipeBookChangeSettings) State() jp.State { return jp.StatePlay }
func (p *C2SRecipeBookChangeSettings) Bound() jp.Bound { return jp.C2S }

func (p *C2SRecipeBookChangeSettings) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.BookId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.BookOpen, err = buf.ReadBool(); err != nil {
		return err
	}
	p.FilterActive, err = buf.ReadBool()
	return err
}

func (p *C2SRecipeBookChangeSettings) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.BookId); err != nil {
		return err
	}
	if err := buf.WriteBool(p.BookOpen); err != nil {
		return err
	}
	return buf.WriteBool(p.FilterActive)
}

// C2SRecipeBookSeenRecipe represents "Set Seen Recipe".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Seen_Recipe
type C2SRecipeBookSeenRecipe struct {
	RecipeId ns.VarInt
}

func (p *C2SRecipeBookSeenRecipe) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SRecipeBookSeenRecipeID)
}
func (p *C2SRecipeBookSeenRecipe) State() jp.State { return jp.StatePlay }
func (p *C2SRecipeBookSeenRecipe) Bound() jp.Bound { return jp.C2S }

func (p *C2SRecipeBookSeenRecipe) Read(buf *ns.PacketBuffer) error {
	var err error
	p.RecipeId, err = buf.ReadVarInt()
	return err
}

func (p *C2SRecipeBookSeenRecipe) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.RecipeId)
}

// C2SRenameItem represents "Rename Item".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Rename_Item
type C2SRenameItem struct {
	ItemName ns.String
}

func (p *C2SRenameItem) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SRenameItemID) }
func (p *C2SRenameItem) State() jp.State { return jp.StatePlay }
func (p *C2SRenameItem) Bound() jp.Bound { return jp.C2S }

func (p *C2SRenameItem) Read(buf *ns.PacketBuffer) error {
	var err error
	p.ItemName, err = buf.ReadString(50)
	return err
}

func (p *C2SRenameItem) Write(buf *ns.PacketBuffer) error {
	return buf.WriteString(p.ItemName)
}

// C2SResourcePackPlay represents "Resource Pack Response (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Resource_Pack_Response_(Play)
type C2SResourcePackPlay struct {
	Uuid   ns.UUID
	Result ns.VarInt
}

func (p *C2SResourcePackPlay) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SResourcePackPlayID) }
func (p *C2SResourcePackPlay) State() jp.State { return jp.StatePlay }
func (p *C2SResourcePackPlay) Bound() jp.Bound { return jp.C2S }

func (p *C2SResourcePackPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Uuid, err = buf.ReadUUID(); err != nil {
		return err
	}
	p.Result, err = buf.ReadVarInt()
	return err
}

func (p *C2SResourcePackPlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteUUID(p.Uuid); err != nil {
		return err
	}
	return buf.WriteVarInt(p.Result)
}

// C2SSeenAdvancements represents "Seen Advancements".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Seen_Advancements
type C2SSeenAdvancements struct {
	Action ns.VarInt
	TabId  ns.Identifier // only present if Action is 0
}

func (p *C2SSeenAdvancements) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSeenAdvancementsID) }
func (p *C2SSeenAdvancements) State() jp.State { return jp.StatePlay }
func (p *C2SSeenAdvancements) Bound() jp.Bound { return jp.C2S }

func (p *C2SSeenAdvancements) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Action, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Action == 0 {
		p.TabId, err = buf.ReadIdentifier()
	}
	return err
}

func (p *C2SSeenAdvancements) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.Action); err != nil {
		return err
	}
	if p.Action == 0 {
		return buf.WriteIdentifier(p.TabId)
	}
	return nil
}

// C2SSelectTrade represents "Select Trade".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Select_Trade
type C2SSelectTrade struct {
	SelectedSlot ns.VarInt
}

func (p *C2SSelectTrade) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSelectTradeID) }
func (p *C2SSelectTrade) State() jp.State { return jp.StatePlay }
func (p *C2SSelectTrade) Bound() jp.Bound { return jp.C2S }

func (p *C2SSelectTrade) Read(buf *ns.PacketBuffer) error {
	var err error
	p.SelectedSlot, err = buf.ReadVarInt()
	return err
}

func (p *C2SSelectTrade) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.SelectedSlot)
}

// C2SSetBeacon represents "Set Beacon Effect".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Beacon_Effect
type C2SSetBeacon struct {
	PrimaryEffect   ns.PrefixedOptional[ns.VarInt]
	SecondaryEffect ns.PrefixedOptional[ns.VarInt]
}

func (p *C2SSetBeacon) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSetBeaconID) }
func (p *C2SSetBeacon) State() jp.State { return jp.StatePlay }
func (p *C2SSetBeacon) Bound() jp.Bound { return jp.C2S }

func (p *C2SSetBeacon) Read(buf *ns.PacketBuffer) error {
	if err := p.PrimaryEffect.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.VarInt, error) {
		return b.ReadVarInt()
	}); err != nil {
		return err
	}
	return p.SecondaryEffect.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.VarInt, error) {
		return b.ReadVarInt()
	})
}

func (p *C2SSetBeacon) Write(buf *ns.PacketBuffer) error {
	if err := p.PrimaryEffect.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.VarInt) error {
		return b.WriteVarInt(v)
	}); err != nil {
		return err
	}
	return p.SecondaryEffect.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.VarInt) error {
		return b.WriteVarInt(v)
	})
}

// C2SSetCarriedItem represents "Set Held Item (serverbound)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Held_Item_(Serverbound)
type C2SSetCarriedItem struct {
	Slot ns.Int16
}

func (p *C2SSetCarriedItem) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSetCarriedItemID) }
func (p *C2SSetCarriedItem) State() jp.State { return jp.StatePlay }
func (p *C2SSetCarriedItem) Bound() jp.Bound { return jp.C2S }

func (p *C2SSetCarriedItem) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Slot, err = buf.ReadInt16()
	return err
}

func (p *C2SSetCarriedItem) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt16(p.Slot)
}

// C2SSetCommandBlock represents "Program Command Block".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Program_Command_Block
type C2SSetCommandBlock struct {
	Location ns.Position
	Command  ns.String
	Mode     ns.VarInt
	Flags    ns.Int8
}

func (p *C2SSetCommandBlock) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSetCommandBlockID) }
func (p *C2SSetCommandBlock) State() jp.State { return jp.StatePlay }
func (p *C2SSetCommandBlock) Bound() jp.Bound { return jp.C2S }

func (p *C2SSetCommandBlock) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Command, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.Mode, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Flags, err = buf.ReadInt8()
	return err
}

func (p *C2SSetCommandBlock) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteString(p.Command); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Mode); err != nil {
		return err
	}
	return buf.WriteInt8(p.Flags)
}

// C2SSetCommandMinecart represents "Program Command Block Minecart".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Program_Command_Block_Minecart
type C2SSetCommandMinecart struct {
	EntityId    ns.VarInt
	Command     ns.String
	TrackOutput ns.Boolean
}

func (p *C2SSetCommandMinecart) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SSetCommandMinecartID)
}
func (p *C2SSetCommandMinecart) State() jp.State { return jp.StatePlay }
func (p *C2SSetCommandMinecart) Bound() jp.Bound { return jp.C2S }

func (p *C2SSetCommandMinecart) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Command, err = buf.ReadString(32767); err != nil {
		return err
	}
	p.TrackOutput, err = buf.ReadBool()
	return err
}

func (p *C2SSetCommandMinecart) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteString(p.Command); err != nil {
		return err
	}
	return buf.WriteBool(p.TrackOutput)
}

// C2SSetCreativeModeSlot represents "Set Creative Mode Slot".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Creative_Mode_Slot
type C2SSetCreativeModeSlot struct {
	Slot        ns.Int16
	ClickedItem ns.Slot
}

func (p *C2SSetCreativeModeSlot) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SSetCreativeModeSlotID)
}
func (p *C2SSetCreativeModeSlot) State() jp.State { return jp.StatePlay }
func (p *C2SSetCreativeModeSlot) Bound() jp.Bound { return jp.C2S }

func (p *C2SSetCreativeModeSlot) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Slot, err = buf.ReadInt16(); err != nil {
		return err
	}
	// uses length-prefixed format (OPTIONAL_UNTRUSTED_STREAM_CODEC in JE source code)
	p.ClickedItem, err = buf.ReadSlot(items.DecoderDelimited())
	return err
}

func (p *C2SSetCreativeModeSlot) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt16(p.Slot); err != nil {
		return err
	}
	// uses length-prefixed format (OPTIONAL_UNTRUSTED_STREAM_CODEC in JE source code)
	return items.WriteRawSlotDelimited(buf, p.ClickedItem)
}

// C2SSetJigsawBlock represents "Program Jigsaw Block".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Program_Jigsaw_Block
type C2SSetJigsawBlock struct {
	Location          ns.Position
	Name              ns.Identifier
	Target            ns.Identifier
	Pool              ns.Identifier
	FinalState        ns.String
	JointType         ns.String
	SelectionPriority ns.VarInt
	PlacementPriority ns.VarInt
}

func (p *C2SSetJigsawBlock) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSetJigsawBlockID) }
func (p *C2SSetJigsawBlock) State() jp.State { return jp.StatePlay }
func (p *C2SSetJigsawBlock) Bound() jp.Bound { return jp.C2S }

func (p *C2SSetJigsawBlock) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Name, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	if p.Target, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	if p.Pool, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	if p.FinalState, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.JointType, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.SelectionPriority, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.PlacementPriority, err = buf.ReadVarInt()
	return err
}

func (p *C2SSetJigsawBlock) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteIdentifier(p.Name); err != nil {
		return err
	}
	if err := buf.WriteIdentifier(p.Target); err != nil {
		return err
	}
	if err := buf.WriteIdentifier(p.Pool); err != nil {
		return err
	}
	if err := buf.WriteString(p.FinalState); err != nil {
		return err
	}
	if err := buf.WriteString(p.JointType); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SelectionPriority); err != nil {
		return err
	}
	return buf.WriteVarInt(p.PlacementPriority)
}

// C2SSetStructureBlock represents "Program Structure Block".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Program_Structure_Block
type C2SSetStructureBlock struct {
	Location  ns.Position
	Action    ns.VarInt
	Mode      ns.VarInt
	Name      ns.String
	OffsetX   ns.Int8
	OffsetY   ns.Int8
	OffsetZ   ns.Int8
	SizeX     ns.Int8
	SizeY     ns.Int8
	SizeZ     ns.Int8
	Mirror    ns.VarInt
	Rotation  ns.VarInt
	Metadata  ns.String
	Integrity ns.Float32
	Seed      ns.VarLong
	Flags     ns.Int8
}

func (p *C2SSetStructureBlock) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSetStructureBlockID) }
func (p *C2SSetStructureBlock) State() jp.State { return jp.StatePlay }
func (p *C2SSetStructureBlock) Bound() jp.Bound { return jp.C2S }

func (p *C2SSetStructureBlock) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Action, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Mode, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Name, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.OffsetX, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.OffsetY, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.OffsetZ, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.SizeX, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.SizeY, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.SizeZ, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.Mirror, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Rotation, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Metadata, err = buf.ReadString(128); err != nil {
		return err
	}
	if p.Integrity, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Seed, err = buf.ReadVarLong(); err != nil {
		return err
	}
	p.Flags, err = buf.ReadInt8()
	return err
}

func (p *C2SSetStructureBlock) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Action); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Mode); err != nil {
		return err
	}
	if err := buf.WriteString(p.Name); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.OffsetX); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.OffsetY); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.OffsetZ); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.SizeX); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.SizeY); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.SizeZ); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Mirror); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Rotation); err != nil {
		return err
	}
	if err := buf.WriteString(p.Metadata); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Integrity); err != nil {
		return err
	}
	if err := buf.WriteVarLong(p.Seed); err != nil {
		return err
	}
	return buf.WriteInt8(p.Flags)
}

// C2SSetTestBlock represents "Set Test Block".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Test_Block
type C2SSetTestBlock struct {
	Position ns.Position
	Mode     ns.VarInt
	Message  ns.String
}

func (p *C2SSetTestBlock) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSetTestBlockID) }
func (p *C2SSetTestBlock) State() jp.State { return jp.StatePlay }
func (p *C2SSetTestBlock) Bound() jp.Bound { return jp.C2S }

func (p *C2SSetTestBlock) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Position, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Mode, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Message, err = buf.ReadString(32767)
	return err
}

func (p *C2SSetTestBlock) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Position); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Mode); err != nil {
		return err
	}
	return buf.WriteString(p.Message)
}

// C2SSignUpdate represents "Update Sign".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Sign
type C2SSignUpdate struct {
	Location    ns.Position
	IsFrontText ns.Boolean
	Line1       ns.String
	Line2       ns.String
	Line3       ns.String
	Line4       ns.String
}

func (p *C2SSignUpdate) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSignUpdateID) }
func (p *C2SSignUpdate) State() jp.State { return jp.StatePlay }
func (p *C2SSignUpdate) Bound() jp.Bound { return jp.C2S }

func (p *C2SSignUpdate) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.IsFrontText, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.Line1, err = buf.ReadString(384); err != nil {
		return err
	}
	if p.Line2, err = buf.ReadString(384); err != nil {
		return err
	}
	if p.Line3, err = buf.ReadString(384); err != nil {
		return err
	}
	p.Line4, err = buf.ReadString(384)
	return err
}

func (p *C2SSignUpdate) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteBool(p.IsFrontText); err != nil {
		return err
	}
	if err := buf.WriteString(p.Line1); err != nil {
		return err
	}
	if err := buf.WriteString(p.Line2); err != nil {
		return err
	}
	if err := buf.WriteString(p.Line3); err != nil {
		return err
	}
	return buf.WriteString(p.Line4)
}

// C2SSwing represents "Swing Arm".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Swing_Arm
type C2SSwing struct {
	Hand ns.VarInt
}

func (p *C2SSwing) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SSwingID) }
func (p *C2SSwing) State() jp.State { return jp.StatePlay }
func (p *C2SSwing) Bound() jp.Bound { return jp.C2S }

func (p *C2SSwing) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Hand, err = buf.ReadVarInt()
	return err
}

func (p *C2SSwing) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.Hand)
}

// C2STeleportToEntity represents "Teleport To Entity".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Teleport_To_Entity
type C2STeleportToEntity struct {
	TargetPlayer ns.UUID
}

func (p *C2STeleportToEntity) ID() ns.VarInt   { return ns.VarInt(packets_data.C2STeleportToEntityID) }
func (p *C2STeleportToEntity) State() jp.State { return jp.StatePlay }
func (p *C2STeleportToEntity) Bound() jp.Bound { return jp.C2S }

func (p *C2STeleportToEntity) Read(buf *ns.PacketBuffer) error {
	var err error
	p.TargetPlayer, err = buf.ReadUUID()
	return err
}

func (p *C2STeleportToEntity) Write(buf *ns.PacketBuffer) error {
	return buf.WriteUUID(p.TargetPlayer)
}

// C2STestInstanceBlockAction represents "Test Instance Block Action".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Test_Instance_Block_Action
type C2STestInstanceBlockAction struct {
	Position       ns.Position
	Action         ns.VarInt
	Test           ns.PrefixedOptional[ns.Identifier]
	SizeX          ns.VarInt
	SizeY          ns.VarInt
	SizeZ          ns.VarInt
	Rotation       ns.VarInt
	IgnoreEntities ns.Boolean
	Status         ns.VarInt
	ErrorMessage   ns.PrefixedOptional[ns.TextComponent]
}

func (p *C2STestInstanceBlockAction) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2STestInstanceBlockActionID)
}
func (p *C2STestInstanceBlockAction) State() jp.State { return jp.StatePlay }
func (p *C2STestInstanceBlockAction) Bound() jp.Bound { return jp.C2S }

func (p *C2STestInstanceBlockAction) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Position, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Action, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if err = p.Test.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.Identifier, error) {
		return b.ReadIdentifier()
	}); err != nil {
		return err
	}
	if p.SizeX, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.SizeY, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.SizeZ, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Rotation, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.IgnoreEntities, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.Status, err = buf.ReadVarInt(); err != nil {
		return err
	}
	return p.ErrorMessage.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.TextComponent, error) {
		return b.ReadTextComponent()
	})
}

func (p *C2STestInstanceBlockAction) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Position); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Action); err != nil {
		return err
	}
	if err := p.Test.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.Identifier) error {
		return b.WriteIdentifier(v)
	}); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SizeX); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SizeY); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SizeZ); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Rotation); err != nil {
		return err
	}
	if err := buf.WriteBool(p.IgnoreEntities); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Status); err != nil {
		return err
	}
	return p.ErrorMessage.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.TextComponent) error {
		return b.WriteTextComponent(v)
	})
}

// C2SUseItemOn represents "Use Item On".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Use_Item_On
type C2SUseItemOn struct {
	Hand            ns.VarInt
	Location        ns.Position
	Face            ns.VarInt
	CursorPositionX ns.Float32
	CursorPositionY ns.Float32
	CursorPositionZ ns.Float32
	InsideBlock     ns.Boolean
	WorldBorderHit  ns.Boolean
	Sequence        ns.VarInt
}

func (p *C2SUseItemOn) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SUseItemOnID) }
func (p *C2SUseItemOn) State() jp.State { return jp.StatePlay }
func (p *C2SUseItemOn) Bound() jp.Bound { return jp.C2S }

func (p *C2SUseItemOn) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Hand, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Face, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.CursorPositionX, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.CursorPositionY, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.CursorPositionZ, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.InsideBlock, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.WorldBorderHit, err = buf.ReadBool(); err != nil {
		return err
	}
	p.Sequence, err = buf.ReadVarInt()
	return err
}

func (p *C2SUseItemOn) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.Hand); err != nil {
		return err
	}
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Face); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.CursorPositionX); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.CursorPositionY); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.CursorPositionZ); err != nil {
		return err
	}
	if err := buf.WriteBool(p.InsideBlock); err != nil {
		return err
	}
	if err := buf.WriteBool(p.WorldBorderHit); err != nil {
		return err
	}
	return buf.WriteVarInt(p.Sequence)
}

// C2SUseItem represents "Use Item".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Use_Item
type C2SUseItem struct {
	Hand     ns.VarInt
	Sequence ns.VarInt
	Yaw      ns.Float32
	Pitch    ns.Float32
}

func (p *C2SUseItem) ID() ns.VarInt   { return ns.VarInt(packets_data.C2SUseItemID) }
func (p *C2SUseItem) State() jp.State { return jp.StatePlay }
func (p *C2SUseItem) Bound() jp.Bound { return jp.C2S }

func (p *C2SUseItem) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Hand, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Sequence, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Yaw, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.Pitch, err = buf.ReadFloat32()
	return err
}

func (p *C2SUseItem) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.Hand); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Sequence); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Yaw); err != nil {
		return err
	}
	return buf.WriteFloat32(p.Pitch)
}

// C2SCustomClickActionPlay represents "Custom Click Action (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Custom_Click_Action_(Play)
type C2SCustomClickActionPlay struct {
	Id      ns.Identifier
	Payload nbt.Tag
}

func (p *C2SCustomClickActionPlay) ID() ns.VarInt {
	return ns.VarInt(packets_data.C2SCustomClickActionPlayID)
}
func (p *C2SCustomClickActionPlay) State() jp.State { return jp.StatePlay }
func (p *C2SCustomClickActionPlay) Bound() jp.Bound { return jp.C2S }

func (p *C2SCustomClickActionPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Id, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	remaining, err := buf.ReadByteArray(1048576)
	if err != nil {
		return err
	}
	p.Payload, err = nbt.DecodeNetwork(remaining)
	return err
}

func (p *C2SCustomClickActionPlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteIdentifier(p.Id); err != nil {
		return err
	}
	data, err := nbt.EncodeNetwork(p.Payload)
	if err != nil {
		return err
	}
	return buf.WriteFixedByteArray(data)
}

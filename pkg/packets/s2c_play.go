package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
	"github.com/go-mclib/protocol/nbt"
)

const (
	S2CBundleDelimiterID ns.VarInt = iota
	S2CAddEntityID
	S2CAnimateID
	S2CAwardStatsID
	S2CBlockChangedAckID
	S2CBlockDestructionID
	S2CBlockEntityDataID
	S2CBlockEventID
	S2CBlockUpdateID
	S2CBossEventID
	S2CChangeDifficultyID
	S2CChunkBatchFinishedID
	S2CChunkBatchStartID
	S2CChunksBiomesID
	S2CClearTitlesID
	S2CCommandSuggestionsID
	S2CCommandsID
	S2CContainerCloseID
	S2CContainerSetContentID
	S2CContainerSetDataID
	S2CContainerSetSlotID
	S2CCookieRequestPlayID
	S2CCooldownID
	S2CCustomChatCompletionsID
	S2CCustomPayloadPlayID
	S2CDamageEventID
	S2CDebugBlockValueID
	S2CDebugChunkValueID
	S2CDebugEntityValueID
	S2CDebugEventID
	S2CDebugSampleID
	S2CDeleteChatID
	S2CDisconnectPlayID
	S2CDisguisedChatID
	S2CEntityEventID
	S2CEntityPositionSyncID
	S2CExplodeID
	S2CForgetLevelChunkID
	S2CGameEventID
	S2CGameTestHighlightPosID
	S2CHorseScreenOpenID
	S2CHurtAnimationID
	S2CInitializeBorderID
	S2CKeepAlivePlayID
	S2CLevelChunkWithLightID
	S2CLevelEventID
	S2CLevelParticlesID
	S2CLightUpdateID
	S2CLoginPlayID
	S2CMapItemDataID
	S2CMerchantOffersID
	S2CMoveEntityPosID
	S2CMoveEntityPosRotID
	S2CMoveMinecartAlongTrackID
	S2CMoveEntityRotID
	S2CMoveVehicleID
	S2COpenBookID
	S2COpenScreenID
	S2COpenSignEditorID
	S2CPingPlayID
	S2CPongResponsePlayID
	S2CPlaceGhostRecipeID
	S2CPlayerAbilitiesID
	S2CPlayerChatID
	S2CPlayerCombatEndID
	S2CPlayerCombatEnterID
	S2CPlayerCombatKillID
	S2CPlayerInfoRemoveID
	S2CPlayerInfoUpdateID
	S2CPlayerLookAtID
	S2CPlayerPositionID
	S2CPlayerRotationID
	S2CRecipeBookAddID
	S2CRecipeBookRemoveID
	S2CRecipeBookSettingsID
	S2CRemoveEntitiesID
	S2CRemoveMobEffectID
	S2CResetScoreID
	S2CResourcePackPopPlayID
	S2CResourcePackPushPlayID
	S2CRespawnID
	S2CRotateHeadID
	S2CSectionBlocksUpdateID
	S2CSelectAdvancementsTabID
	S2CServerDataID
	S2CSetActionBarTextID
	S2CSetBorderCenterID
	S2CSetBorderLerpSizeID
	S2CSetBorderSizeID
	S2CSetBorderWarningDelayID
	S2CSetBorderWarningDistanceID
	S2CSetCameraID
	S2CSetChunkCacheCenterID
	S2CSetChunkCacheRadiusID
	S2CSetCursorItemID
	S2CSetDefaultSpawnPositionID
	S2CSetDisplayObjectiveID
	S2CSetEntityDataID
	S2CSetEntityLinkID
	S2CSetEntityMotionID
	S2CSetEquipmentID
	S2CSetExperienceID
	S2CSetHealthID
	S2CSetHeldSlotID
	S2CSetObjectiveID
	S2CSetPassengersID
	S2CSetPlayerInventoryID
	S2CSetPlayerTeamID
	S2CSetScoreID
	S2CSetSimulationDistanceID
	S2CSetSubtitleTextID
	S2CSetTimeID
	S2CSetTitleTextID
	S2CSetTitlesAnimationID
	S2CSoundEntityID
	S2CSoundID
	S2CStartConfigurationID
	S2CStopSoundID
	S2CStoreCookiePlayID
	S2CSystemChatID
	S2CTabListID
	S2CTagQueryID
	S2CTakeItemEntityID
	S2CTeleportEntityID
	S2CTestInstanceBlockStatusID
	S2CTickingStateID
	S2CTickingStepID
	S2CTransferPlayID
	S2CUpdateAdvancementsID
	S2CUpdateAttributesID
	S2CUpdateMobEffectID
	S2CUpdateRecipesID
	S2CUpdateTagsPlayID
	S2CProjectilePowerID
	S2CCustomReportDetailsID
	S2CServerLinksID
	S2CWaypointID
	S2CClearDialogPlayID
	S2CShowDialogPlayID
)

// S2CBundleDelimiter represents "Bundle Delimiter".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Bundle_Delimiter
type S2CBundleDelimiter struct{}

func (p *S2CBundleDelimiter) ID() ns.VarInt                { return S2CBundleDelimiterID }
func (p *S2CBundleDelimiter) State() jp.State              { return jp.StatePlay }
func (p *S2CBundleDelimiter) Bound() jp.Bound              { return jp.S2C }
func (p *S2CBundleDelimiter) Read(*ns.PacketBuffer) error  { return nil }
func (p *S2CBundleDelimiter) Write(*ns.PacketBuffer) error { return nil }

// S2CAddEntity represents "Spawn Entity".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Spawn_Entity
type S2CAddEntity struct {
	EntityId   ns.VarInt
	EntityUuid ns.UUID
	Type       ns.VarInt
	X          ns.Float64
	Y          ns.Float64
	Z          ns.Float64
	Velocity   ns.LpVec3
	Pitch      ns.Angle
	Yaw        ns.Angle
	HeadYaw    ns.Angle
	Data       ns.VarInt
}

func (p *S2CAddEntity) ID() ns.VarInt   { return S2CAddEntityID }
func (p *S2CAddEntity) State() jp.State { return jp.StatePlay }
func (p *S2CAddEntity) Bound() jp.Bound { return jp.S2C }

func (p *S2CAddEntity) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.EntityUuid, err = buf.ReadUUID(); err != nil {
		return err
	}
	if p.Type, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Y, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Z, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Velocity, err = buf.ReadLpVec3(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadAngle(); err != nil {
		return err
	}
	if p.Yaw, err = buf.ReadAngle(); err != nil {
		return err
	}
	if p.HeadYaw, err = buf.ReadAngle(); err != nil {
		return err
	}
	p.Data, err = buf.ReadVarInt()
	return err
}

func (p *S2CAddEntity) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteUUID(p.EntityUuid); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Type); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Y); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	if err := buf.WriteLpVec3(p.Velocity); err != nil {
		return err
	}
	if err := buf.WriteAngle(p.Pitch); err != nil {
		return err
	}
	if err := buf.WriteAngle(p.Yaw); err != nil {
		return err
	}
	if err := buf.WriteAngle(p.HeadYaw); err != nil {
		return err
	}
	return buf.WriteVarInt(p.Data)
}

// S2CAnimate represents "Entity Animation".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Entity_Animation
type S2CAnimate struct {
	EntityId  ns.VarInt
	Animation ns.Uint8
}

func (p *S2CAnimate) ID() ns.VarInt   { return S2CAnimateID }
func (p *S2CAnimate) State() jp.State { return jp.StatePlay }
func (p *S2CAnimate) Bound() jp.Bound { return jp.S2C }

func (p *S2CAnimate) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Animation, err = buf.ReadUint8()
	return err
}

func (p *S2CAnimate) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteUint8(p.Animation)
}

// S2CAwardStats represents "Award Statistics".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Award_Statistics
type S2CAwardStats struct {
	Statistics ns.ByteArray
}

func (p *S2CAwardStats) ID() ns.VarInt   { return S2CAwardStatsID }
func (p *S2CAwardStats) State() jp.State { return jp.StatePlay }
func (p *S2CAwardStats) Bound() jp.Bound { return jp.S2C }

func (p *S2CAwardStats) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Statistics, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CAwardStats) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Statistics)
}

// S2CBlockChangedAck represents "Acknowledge Block Change".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Acknowledge_Block_Change
type S2CBlockChangedAck struct {
	SequenceId ns.VarInt
}

func (p *S2CBlockChangedAck) ID() ns.VarInt   { return S2CBlockChangedAckID }
func (p *S2CBlockChangedAck) State() jp.State { return jp.StatePlay }
func (p *S2CBlockChangedAck) Bound() jp.Bound { return jp.S2C }

func (p *S2CBlockChangedAck) Read(buf *ns.PacketBuffer) error {
	var err error
	p.SequenceId, err = buf.ReadVarInt()
	return err
}

func (p *S2CBlockChangedAck) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.SequenceId)
}

// S2CBlockDestruction represents "Set Block Destroy Stage".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Block_Destroy_Stage
type S2CBlockDestruction struct {
	EntityId     ns.VarInt
	Location     ns.Position
	DestroyStage ns.Uint8
}

func (p *S2CBlockDestruction) ID() ns.VarInt   { return S2CBlockDestructionID }
func (p *S2CBlockDestruction) State() jp.State { return jp.StatePlay }
func (p *S2CBlockDestruction) Bound() jp.Bound { return jp.S2C }

func (p *S2CBlockDestruction) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	p.DestroyStage, err = buf.ReadUint8()
	return err
}

func (p *S2CBlockDestruction) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	return buf.WriteUint8(p.DestroyStage)
}

// S2CBlockEntityData represents "Block Entity Data".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Block_Entity_Data
type S2CBlockEntityData struct {
	Location ns.Position
	Type     ns.VarInt
	NbtData  nbt.Tag
}

func (p *S2CBlockEntityData) ID() ns.VarInt   { return S2CBlockEntityDataID }
func (p *S2CBlockEntityData) State() jp.State { return jp.StatePlay }
func (p *S2CBlockEntityData) Bound() jp.Bound { return jp.S2C }

func (p *S2CBlockEntityData) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Type, err = buf.ReadVarInt(); err != nil {
		return err
	}
	remaining, err := buf.ReadByteArray(1048576)
	if err != nil {
		return err
	}
	p.NbtData, err = nbt.DecodeNetwork(remaining)
	return err
}

func (p *S2CBlockEntityData) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Type); err != nil {
		return err
	}
	data, err := nbt.EncodeNetwork(p.NbtData)
	if err != nil {
		return err
	}
	return buf.WriteFixedByteArray(data)
}

// S2CBlockEvent represents "Block Action".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Block_Action
type S2CBlockEvent struct {
	Location        ns.Position
	ActionId        ns.Uint8
	ActionParameter ns.Uint8
	BlockType       ns.VarInt
}

func (p *S2CBlockEvent) ID() ns.VarInt   { return S2CBlockEventID }
func (p *S2CBlockEvent) State() jp.State { return jp.StatePlay }
func (p *S2CBlockEvent) Bound() jp.Bound { return jp.S2C }

func (p *S2CBlockEvent) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.ActionId, err = buf.ReadUint8(); err != nil {
		return err
	}
	if p.ActionParameter, err = buf.ReadUint8(); err != nil {
		return err
	}
	p.BlockType, err = buf.ReadVarInt()
	return err
}

func (p *S2CBlockEvent) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteUint8(p.ActionId); err != nil {
		return err
	}
	if err := buf.WriteUint8(p.ActionParameter); err != nil {
		return err
	}
	return buf.WriteVarInt(p.BlockType)
}

// S2CBlockUpdate represents "Block Update".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Block_Update
type S2CBlockUpdate struct {
	Location ns.Position
	BlockId  ns.VarInt
}

func (p *S2CBlockUpdate) ID() ns.VarInt   { return S2CBlockUpdateID }
func (p *S2CBlockUpdate) State() jp.State { return jp.StatePlay }
func (p *S2CBlockUpdate) Bound() jp.Bound { return jp.S2C }

func (p *S2CBlockUpdate) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	p.BlockId, err = buf.ReadVarInt()
	return err
}

func (p *S2CBlockUpdate) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	return buf.WriteVarInt(p.BlockId)
}

// S2CBossEvent represents "Boss Bar".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Boss_Bar
type S2CBossEvent struct {
	Uuid   ns.UUID
	Action ns.VarInt
	Data   ns.ByteArray
}

func (p *S2CBossEvent) ID() ns.VarInt   { return S2CBossEventID }
func (p *S2CBossEvent) State() jp.State { return jp.StatePlay }
func (p *S2CBossEvent) Bound() jp.Bound { return jp.S2C }

func (p *S2CBossEvent) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Uuid, err = buf.ReadUUID(); err != nil {
		return err
	}
	if p.Action, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CBossEvent) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteUUID(p.Uuid); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Action); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CChangeDifficulty represents "Change Difficulty".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Change_Difficulty
type S2CChangeDifficulty struct {
	Difficulty       ns.Uint8
	DifficultyLocked ns.Boolean
}

func (p *S2CChangeDifficulty) ID() ns.VarInt   { return S2CChangeDifficultyID }
func (p *S2CChangeDifficulty) State() jp.State { return jp.StatePlay }
func (p *S2CChangeDifficulty) Bound() jp.Bound { return jp.S2C }

func (p *S2CChangeDifficulty) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Difficulty, err = buf.ReadUint8(); err != nil {
		return err
	}
	p.DifficultyLocked, err = buf.ReadBool()
	return err
}

func (p *S2CChangeDifficulty) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteUint8(p.Difficulty); err != nil {
		return err
	}
	return buf.WriteBool(p.DifficultyLocked)
}

// S2CChunkBatchFinished represents "Chunk Batch Finished".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chunk_Batch_Finished
type S2CChunkBatchFinished struct {
	BatchSize ns.VarInt
}

func (p *S2CChunkBatchFinished) ID() ns.VarInt   { return S2CChunkBatchFinishedID }
func (p *S2CChunkBatchFinished) State() jp.State { return jp.StatePlay }
func (p *S2CChunkBatchFinished) Bound() jp.Bound { return jp.S2C }

func (p *S2CChunkBatchFinished) Read(buf *ns.PacketBuffer) error {
	var err error
	p.BatchSize, err = buf.ReadVarInt()
	return err
}

func (p *S2CChunkBatchFinished) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.BatchSize)
}

// S2CChunkBatchStart represents "Chunk Batch Start".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chunk_Batch_Start
type S2CChunkBatchStart struct{}

func (p *S2CChunkBatchStart) ID() ns.VarInt                { return S2CChunkBatchStartID }
func (p *S2CChunkBatchStart) State() jp.State              { return jp.StatePlay }
func (p *S2CChunkBatchStart) Bound() jp.Bound              { return jp.S2C }
func (p *S2CChunkBatchStart) Read(*ns.PacketBuffer) error  { return nil }
func (p *S2CChunkBatchStart) Write(*ns.PacketBuffer) error { return nil }

// S2CChunksBiomes represents "Chunk Biomes".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chunk_Biomes
type S2CChunksBiomes struct {
	ChunkBiomeData ns.ByteArray
}

func (p *S2CChunksBiomes) ID() ns.VarInt   { return S2CChunksBiomesID }
func (p *S2CChunksBiomes) State() jp.State { return jp.StatePlay }
func (p *S2CChunksBiomes) Bound() jp.Bound { return jp.S2C }

func (p *S2CChunksBiomes) Read(buf *ns.PacketBuffer) error {
	var err error
	p.ChunkBiomeData, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CChunksBiomes) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.ChunkBiomeData)
}

// S2CClearTitles represents "Clear Titles".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clear_Titles
type S2CClearTitles struct {
	Reset ns.Boolean
}

func (p *S2CClearTitles) ID() ns.VarInt   { return S2CClearTitlesID }
func (p *S2CClearTitles) State() jp.State { return jp.StatePlay }
func (p *S2CClearTitles) Bound() jp.Bound { return jp.S2C }

func (p *S2CClearTitles) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Reset, err = buf.ReadBool()
	return err
}

func (p *S2CClearTitles) Write(buf *ns.PacketBuffer) error {
	return buf.WriteBool(p.Reset)
}

// S2CCommandSuggestions represents "Command Suggestions Response".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Command_Suggestions_Response
type S2CCommandSuggestions struct {
	Id      ns.VarInt
	Start   ns.VarInt
	Length  ns.VarInt
	Matches ns.ByteArray
}

func (p *S2CCommandSuggestions) ID() ns.VarInt   { return S2CCommandSuggestionsID }
func (p *S2CCommandSuggestions) State() jp.State { return jp.StatePlay }
func (p *S2CCommandSuggestions) Bound() jp.Bound { return jp.S2C }

func (p *S2CCommandSuggestions) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Id, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Start, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Length, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Matches, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CCommandSuggestions) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.Id); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Start); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Length); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Matches)
}

// S2CCommands represents "Commands".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Commands
type S2CCommands struct {
	Data ns.ByteArray
}

func (p *S2CCommands) ID() ns.VarInt   { return S2CCommandsID }
func (p *S2CCommands) State() jp.State { return jp.StatePlay }
func (p *S2CCommands) Bound() jp.Bound { return jp.S2C }

func (p *S2CCommands) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CCommands) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Data)
}

// S2CContainerClose represents "Close Container".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Close_Container
type S2CContainerClose struct {
	WindowId ns.VarInt
}

func (p *S2CContainerClose) ID() ns.VarInt   { return S2CContainerCloseID }
func (p *S2CContainerClose) State() jp.State { return jp.StatePlay }
func (p *S2CContainerClose) Bound() jp.Bound { return jp.S2C }

func (p *S2CContainerClose) Read(buf *ns.PacketBuffer) error {
	var err error
	p.WindowId, err = buf.ReadVarInt()
	return err
}

func (p *S2CContainerClose) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.WindowId)
}

// S2CContainerSetContent represents "Set Container Content".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Container_Content
type S2CContainerSetContent struct {
	WindowId    ns.VarInt
	StateId     ns.VarInt
	Slots       []ns.Slot
	CarriedItem ns.Slot
}

func (p *S2CContainerSetContent) ID() ns.VarInt   { return S2CContainerSetContentID }
func (p *S2CContainerSetContent) State() jp.State { return jp.StatePlay }
func (p *S2CContainerSetContent) Bound() jp.Bound { return jp.S2C }

func (p *S2CContainerSetContent) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.StateId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	slotCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	p.Slots = make([]ns.Slot, slotCount)
	for i := range p.Slots {
		if p.Slots[i], err = buf.ReadSlot(); err != nil {
			return err
		}
	}
	p.CarriedItem, err = buf.ReadSlot()
	return err
}

func (p *S2CContainerSetContent) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.StateId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(len(p.Slots))); err != nil {
		return err
	}
	for _, slot := range p.Slots {
		if err := buf.WriteSlot(slot); err != nil {
			return err
		}
	}
	return buf.WriteSlot(p.CarriedItem)
}

// S2CContainerSetData represents "Set Container Property".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Container_Property
type S2CContainerSetData struct {
	WindowId ns.VarInt
	Property ns.Int16
	Value    ns.Int16
}

func (p *S2CContainerSetData) ID() ns.VarInt   { return S2CContainerSetDataID }
func (p *S2CContainerSetData) State() jp.State { return jp.StatePlay }
func (p *S2CContainerSetData) Bound() jp.Bound { return jp.S2C }

func (p *S2CContainerSetData) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Property, err = buf.ReadInt16(); err != nil {
		return err
	}
	p.Value, err = buf.ReadInt16()
	return err
}

func (p *S2CContainerSetData) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	if err := buf.WriteInt16(p.Property); err != nil {
		return err
	}
	return buf.WriteInt16(p.Value)
}

// S2CContainerSetSlot represents "Set Container Slot".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Container_Slot
type S2CContainerSetSlot struct {
	WindowId ns.VarInt
	StateId  ns.VarInt
	Slot     ns.Int16
	SlotData ns.Slot
}

func (p *S2CContainerSetSlot) ID() ns.VarInt   { return S2CContainerSetSlotID }
func (p *S2CContainerSetSlot) State() jp.State { return jp.StatePlay }
func (p *S2CContainerSetSlot) Bound() jp.Bound { return jp.S2C }

func (p *S2CContainerSetSlot) Read(buf *ns.PacketBuffer) error {
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
	p.SlotData, err = buf.ReadSlot()
	return err
}

func (p *S2CContainerSetSlot) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.StateId); err != nil {
		return err
	}
	if err := buf.WriteInt16(p.Slot); err != nil {
		return err
	}
	return buf.WriteSlot(p.SlotData)
}

// S2CCookieRequestPlay represents "Cookie Request (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Cookie_Request_(Play)
type S2CCookieRequestPlay struct {
	Key ns.Identifier
}

func (p *S2CCookieRequestPlay) ID() ns.VarInt   { return S2CCookieRequestPlayID }
func (p *S2CCookieRequestPlay) State() jp.State { return jp.StatePlay }
func (p *S2CCookieRequestPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CCookieRequestPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Key, err = buf.ReadIdentifier()
	return err
}

func (p *S2CCookieRequestPlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteIdentifier(p.Key)
}

// S2CCooldown represents "Set Cooldown".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Cooldown
type S2CCooldown struct {
	CooldownGroup ns.Identifier
	CooldownTicks ns.VarInt
}

func (p *S2CCooldown) ID() ns.VarInt   { return S2CCooldownID }
func (p *S2CCooldown) State() jp.State { return jp.StatePlay }
func (p *S2CCooldown) Bound() jp.Bound { return jp.S2C }

func (p *S2CCooldown) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.CooldownGroup, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	p.CooldownTicks, err = buf.ReadVarInt()
	return err
}

func (p *S2CCooldown) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteIdentifier(p.CooldownGroup); err != nil {
		return err
	}
	return buf.WriteVarInt(p.CooldownTicks)
}

// S2CCustomChatCompletions represents "Chat Suggestions".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chat_Suggestions
type S2CCustomChatCompletions struct {
	Action  ns.VarInt
	Entries ns.ByteArray
}

func (p *S2CCustomChatCompletions) ID() ns.VarInt   { return S2CCustomChatCompletionsID }
func (p *S2CCustomChatCompletions) State() jp.State { return jp.StatePlay }
func (p *S2CCustomChatCompletions) Bound() jp.Bound { return jp.S2C }

func (p *S2CCustomChatCompletions) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Action, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Entries, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CCustomChatCompletions) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.Action); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Entries)
}

// S2CCustomPayloadPlay represents "Clientbound Plugin Message (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clientbound_Plugin_Message_(Play)
type S2CCustomPayloadPlay struct {
	Channel ns.Identifier
	Data    ns.ByteArray
}

func (p *S2CCustomPayloadPlay) ID() ns.VarInt   { return S2CCustomPayloadPlayID }
func (p *S2CCustomPayloadPlay) State() jp.State { return jp.StatePlay }
func (p *S2CCustomPayloadPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CCustomPayloadPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Channel, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CCustomPayloadPlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteIdentifier(p.Channel); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CDamageEvent represents "Damage Event".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Damage_Event
type S2CDamageEvent struct {
	EntityId       ns.VarInt
	SourceTypeId   ns.VarInt
	SourceCauseId  ns.VarInt
	SourceDirectId ns.VarInt
	SourcePosition ns.PrefixedOptional[ns.ByteArray]
}

func (p *S2CDamageEvent) ID() ns.VarInt   { return S2CDamageEventID }
func (p *S2CDamageEvent) State() jp.State { return jp.StatePlay }
func (p *S2CDamageEvent) Bound() jp.Bound { return jp.S2C }

func (p *S2CDamageEvent) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.SourceTypeId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.SourceCauseId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.SourceDirectId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	return p.SourcePosition.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.ByteArray, error) {
		return b.ReadByteArray(24)
	})
}

func (p *S2CDamageEvent) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SourceTypeId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SourceCauseId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SourceDirectId); err != nil {
		return err
	}
	return p.SourcePosition.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.ByteArray) error {
		return b.WriteByteArray(v)
	})
}

// S2CDebugBlockValue represents "Debug Block Value".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Debug_Block_Value
type S2CDebugBlockValue struct {
	Location ns.Position
	Update   ns.ByteArray
}

func (p *S2CDebugBlockValue) ID() ns.VarInt   { return S2CDebugBlockValueID }
func (p *S2CDebugBlockValue) State() jp.State { return jp.StatePlay }
func (p *S2CDebugBlockValue) Bound() jp.Bound { return jp.S2C }

func (p *S2CDebugBlockValue) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	p.Update, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CDebugBlockValue) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Update)
}

// S2CDebugChunkValue represents "Debug Chunk Value".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Debug_Chunk_Value
type S2CDebugChunkValue struct {
	ChunkZ ns.Int32
	ChunkX ns.Int32
	Update ns.ByteArray
}

func (p *S2CDebugChunkValue) ID() ns.VarInt   { return S2CDebugChunkValueID }
func (p *S2CDebugChunkValue) State() jp.State { return jp.StatePlay }
func (p *S2CDebugChunkValue) Bound() jp.Bound { return jp.S2C }

func (p *S2CDebugChunkValue) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.ChunkZ, err = buf.ReadInt32(); err != nil {
		return err
	}
	if p.ChunkX, err = buf.ReadInt32(); err != nil {
		return err
	}
	p.Update, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CDebugChunkValue) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt32(p.ChunkZ); err != nil {
		return err
	}
	if err := buf.WriteInt32(p.ChunkX); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Update)
}

// S2CDebugEntityValue represents "Debug Entity Value".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Debug_Entity_Value
type S2CDebugEntityValue struct {
	EntityId ns.VarInt
	Update   ns.ByteArray
}

func (p *S2CDebugEntityValue) ID() ns.VarInt   { return S2CDebugEntityValueID }
func (p *S2CDebugEntityValue) State() jp.State { return jp.StatePlay }
func (p *S2CDebugEntityValue) Bound() jp.Bound { return jp.S2C }

func (p *S2CDebugEntityValue) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Update, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CDebugEntityValue) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Update)
}

// S2CDebugEvent represents "Debug Event".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Debug_Event
type S2CDebugEvent struct {
	Event ns.ByteArray
}

func (p *S2CDebugEvent) ID() ns.VarInt   { return S2CDebugEventID }
func (p *S2CDebugEvent) State() jp.State { return jp.StatePlay }
func (p *S2CDebugEvent) Bound() jp.Bound { return jp.S2C }

func (p *S2CDebugEvent) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Event, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CDebugEvent) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Event)
}

// S2CDebugSample represents "Debug Sample".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Debug_Sample
type S2CDebugSample struct {
	Sample     ns.ByteArray
	SampleType ns.VarInt
}

func (p *S2CDebugSample) ID() ns.VarInt   { return S2CDebugSampleID }
func (p *S2CDebugSample) State() jp.State { return jp.StatePlay }
func (p *S2CDebugSample) Bound() jp.Bound { return jp.S2C }

func (p *S2CDebugSample) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Sample, err = buf.ReadByteArray(1048576); err != nil {
		return err
	}
	p.SampleType, err = buf.ReadVarInt()
	return err
}

func (p *S2CDebugSample) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteByteArray(p.Sample); err != nil {
		return err
	}
	return buf.WriteVarInt(p.SampleType)
}

// S2CDeleteChat represents "Delete Message".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Delete_Message
type S2CDeleteChat struct {
	MessageId ns.VarInt
	Signature ns.ByteArray
}

func (p *S2CDeleteChat) ID() ns.VarInt   { return S2CDeleteChatID }
func (p *S2CDeleteChat) State() jp.State { return jp.StatePlay }
func (p *S2CDeleteChat) Bound() jp.Bound { return jp.S2C }

func (p *S2CDeleteChat) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.MessageId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.MessageId == 0 {
		p.Signature, err = buf.ReadByteArray(256)
	}
	return err
}

func (p *S2CDeleteChat) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.MessageId); err != nil {
		return err
	}
	if p.MessageId == 0 {
		return buf.WriteFixedByteArray(p.Signature)
	}
	return nil
}

// S2CDisconnectPlay represents "Disconnect (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Disconnect_(Play)
type S2CDisconnectPlay struct {
	Reason ns.TextComponent
}

func (p *S2CDisconnectPlay) ID() ns.VarInt   { return S2CDisconnectPlayID }
func (p *S2CDisconnectPlay) State() jp.State { return jp.StatePlay }
func (p *S2CDisconnectPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CDisconnectPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Reason, err = buf.ReadTextComponent()
	return err
}

func (p *S2CDisconnectPlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteTextComponent(p.Reason)
}

// S2CDisguisedChat represents "Disguised Chat Message".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Disguised_Chat_Message
type S2CDisguisedChat struct {
	Message    ns.TextComponent
	ChatType   ns.ByteArray
	SenderName ns.TextComponent
	TargetName ns.PrefixedOptional[ns.TextComponent]
}

func (p *S2CDisguisedChat) ID() ns.VarInt   { return S2CDisguisedChatID }
func (p *S2CDisguisedChat) State() jp.State { return jp.StatePlay }
func (p *S2CDisguisedChat) Bound() jp.Bound { return jp.S2C }

func (p *S2CDisguisedChat) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Message, err = buf.ReadTextComponent(); err != nil {
		return err
	}
	if p.ChatType, err = buf.ReadByteArray(1048576); err != nil {
		return err
	}
	if p.SenderName, err = buf.ReadTextComponent(); err != nil {
		return err
	}
	return p.TargetName.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.TextComponent, error) {
		return b.ReadTextComponent()
	})
}

func (p *S2CDisguisedChat) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteTextComponent(p.Message); err != nil {
		return err
	}
	if err := buf.WriteByteArray(p.ChatType); err != nil {
		return err
	}
	if err := buf.WriteTextComponent(p.SenderName); err != nil {
		return err
	}
	return p.TargetName.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.TextComponent) error {
		return b.WriteTextComponent(v)
	})
}

// S2CEntityEvent represents "Entity Event".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Entity_Event
type S2CEntityEvent struct {
	EntityId     ns.Int32
	EntityStatus ns.Int8
}

func (p *S2CEntityEvent) ID() ns.VarInt   { return S2CEntityEventID }
func (p *S2CEntityEvent) State() jp.State { return jp.StatePlay }
func (p *S2CEntityEvent) Bound() jp.Bound { return jp.S2C }

func (p *S2CEntityEvent) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadInt32(); err != nil {
		return err
	}
	p.EntityStatus, err = buf.ReadInt8()
	return err
}

func (p *S2CEntityEvent) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt32(p.EntityId); err != nil {
		return err
	}
	return buf.WriteInt8(p.EntityStatus)
}

// S2CEntityPositionSync represents "Teleport Entity".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Teleport_Entity
type S2CEntityPositionSync struct {
	EntityId  ns.VarInt
	X         ns.Float64
	Y         ns.Float64
	Z         ns.Float64
	VelocityX ns.Float64
	VelocityY ns.Float64
	VelocityZ ns.Float64
	Yaw       ns.Float32
	Pitch     ns.Float32
	OnGround  ns.Boolean
}

func (p *S2CEntityPositionSync) ID() ns.VarInt   { return S2CEntityPositionSyncID }
func (p *S2CEntityPositionSync) State() jp.State { return jp.StatePlay }
func (p *S2CEntityPositionSync) Bound() jp.Bound { return jp.S2C }

func (p *S2CEntityPositionSync) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Y, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Z, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.VelocityX, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.VelocityY, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.VelocityZ, err = buf.ReadFloat64(); err != nil {
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

func (p *S2CEntityPositionSync) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Y); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.VelocityX); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.VelocityY); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.VelocityZ); err != nil {
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

// S2CExplode represents "Explosion".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Explosion
type S2CExplode struct {
	X    ns.Float64
	Y    ns.Float64
	Z    ns.Float64
	Data ns.ByteArray
}

func (p *S2CExplode) ID() ns.VarInt   { return S2CExplodeID }
func (p *S2CExplode) State() jp.State { return jp.StatePlay }
func (p *S2CExplode) Bound() jp.Bound { return jp.S2C }

func (p *S2CExplode) Read(buf *ns.PacketBuffer) error {
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
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CExplode) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Y); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CForgetLevelChunk represents "Unload Chunk".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Unload_Chunk
type S2CForgetLevelChunk struct {
	ChunkZ ns.Int32
	ChunkX ns.Int32
}

func (p *S2CForgetLevelChunk) ID() ns.VarInt   { return S2CForgetLevelChunkID }
func (p *S2CForgetLevelChunk) State() jp.State { return jp.StatePlay }
func (p *S2CForgetLevelChunk) Bound() jp.Bound { return jp.S2C }

func (p *S2CForgetLevelChunk) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.ChunkZ, err = buf.ReadInt32(); err != nil {
		return err
	}
	p.ChunkX, err = buf.ReadInt32()
	return err
}

func (p *S2CForgetLevelChunk) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt32(p.ChunkZ); err != nil {
		return err
	}
	return buf.WriteInt32(p.ChunkX)
}

// S2CGameEvent represents "Game Event".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Game_Event
type S2CGameEvent struct {
	Event ns.Uint8
	Value ns.Float32
}

func (p *S2CGameEvent) ID() ns.VarInt   { return S2CGameEventID }
func (p *S2CGameEvent) State() jp.State { return jp.StatePlay }
func (p *S2CGameEvent) Bound() jp.Bound { return jp.S2C }

func (p *S2CGameEvent) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Event, err = buf.ReadUint8(); err != nil {
		return err
	}
	p.Value, err = buf.ReadFloat32()
	return err
}

func (p *S2CGameEvent) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteUint8(p.Event); err != nil {
		return err
	}
	return buf.WriteFloat32(p.Value)
}

// S2CGameTestHighlightPos represents "Game Test Highlight Position" (debug packet).
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Hitboxes
type S2CGameTestHighlightPos struct {
	Data ns.ByteArray
}

func (p *S2CGameTestHighlightPos) ID() ns.VarInt   { return S2CGameTestHighlightPosID }
func (p *S2CGameTestHighlightPos) State() jp.State { return jp.StatePlay }
func (p *S2CGameTestHighlightPos) Bound() jp.Bound { return jp.S2C }

func (p *S2CGameTestHighlightPos) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CGameTestHighlightPos) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Data)
}

// S2CHorseScreenOpen represents "Open Horse Screen".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Open_Horse_Screen
type S2CHorseScreenOpen struct {
	WindowId              ns.VarInt
	InventoryColumnsCount ns.VarInt
	EntityId              ns.Int32
}

func (p *S2CHorseScreenOpen) ID() ns.VarInt   { return S2CHorseScreenOpenID }
func (p *S2CHorseScreenOpen) State() jp.State { return jp.StatePlay }
func (p *S2CHorseScreenOpen) Bound() jp.Bound { return jp.S2C }

func (p *S2CHorseScreenOpen) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.InventoryColumnsCount, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.EntityId, err = buf.ReadInt32()
	return err
}

func (p *S2CHorseScreenOpen) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.InventoryColumnsCount); err != nil {
		return err
	}
	return buf.WriteInt32(p.EntityId)
}

// S2CHurtAnimation represents "Hurt Animation".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Hurt_Animation
type S2CHurtAnimation struct {
	EntityId ns.VarInt
	Yaw      ns.Float32
}

func (p *S2CHurtAnimation) ID() ns.VarInt   { return S2CHurtAnimationID }
func (p *S2CHurtAnimation) State() jp.State { return jp.StatePlay }
func (p *S2CHurtAnimation) Bound() jp.Bound { return jp.S2C }

func (p *S2CHurtAnimation) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Yaw, err = buf.ReadFloat32()
	return err
}

func (p *S2CHurtAnimation) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteFloat32(p.Yaw)
}

// S2CInitializeBorder represents "Initialize World Border".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Initialize_World_Border
type S2CInitializeBorder struct {
	X                      ns.Float64
	Z                      ns.Float64
	OldDiameter            ns.Float64
	NewDiameter            ns.Float64
	Speed                  ns.VarLong
	PortalTeleportBoundary ns.VarInt
	WarningBlocks          ns.VarInt
	WarningTime            ns.VarInt
}

func (p *S2CInitializeBorder) ID() ns.VarInt   { return S2CInitializeBorderID }
func (p *S2CInitializeBorder) State() jp.State { return jp.StatePlay }
func (p *S2CInitializeBorder) Bound() jp.Bound { return jp.S2C }

func (p *S2CInitializeBorder) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Z, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.OldDiameter, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.NewDiameter, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Speed, err = buf.ReadVarLong(); err != nil {
		return err
	}
	if p.PortalTeleportBoundary, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.WarningBlocks, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.WarningTime, err = buf.ReadVarInt()
	return err
}

func (p *S2CInitializeBorder) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.OldDiameter); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.NewDiameter); err != nil {
		return err
	}
	if err := buf.WriteVarLong(p.Speed); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.PortalTeleportBoundary); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.WarningBlocks); err != nil {
		return err
	}
	return buf.WriteVarInt(p.WarningTime)
}

// S2CKeepAlivePlay represents "Clientbound Keep Alive (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clientbound_Keep_Alive_(Play)
type S2CKeepAlivePlay struct {
	KeepAliveId ns.Int64
}

func (p *S2CKeepAlivePlay) ID() ns.VarInt   { return S2CKeepAlivePlayID }
func (p *S2CKeepAlivePlay) State() jp.State { return jp.StatePlay }
func (p *S2CKeepAlivePlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CKeepAlivePlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.KeepAliveId, err = buf.ReadInt64()
	return err
}

func (p *S2CKeepAlivePlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt64(p.KeepAliveId)
}

// S2CLevelChunkWithLight represents "Chunk Data and Update Light".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chunk_Data_And_Update_Light
type S2CLevelChunkWithLight struct {
	ChunkX ns.Int32
	ChunkZ ns.Int32
	Data   ns.ByteArray
}

func (p *S2CLevelChunkWithLight) ID() ns.VarInt   { return S2CLevelChunkWithLightID }
func (p *S2CLevelChunkWithLight) State() jp.State { return jp.StatePlay }
func (p *S2CLevelChunkWithLight) Bound() jp.Bound { return jp.S2C }

func (p *S2CLevelChunkWithLight) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.ChunkX, err = buf.ReadInt32(); err != nil {
		return err
	}
	if p.ChunkZ, err = buf.ReadInt32(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(2097152)
	return err
}

func (p *S2CLevelChunkWithLight) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt32(p.ChunkX); err != nil {
		return err
	}
	if err := buf.WriteInt32(p.ChunkZ); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CLevelEvent represents "World Event".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#World_Event
type S2CLevelEvent struct {
	Event                 ns.Int32
	Location              ns.Position
	Data                  ns.Int32
	DisableRelativeVolume ns.Boolean
}

func (p *S2CLevelEvent) ID() ns.VarInt   { return S2CLevelEventID }
func (p *S2CLevelEvent) State() jp.State { return jp.StatePlay }
func (p *S2CLevelEvent) Bound() jp.Bound { return jp.S2C }

func (p *S2CLevelEvent) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Event, err = buf.ReadInt32(); err != nil {
		return err
	}
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Data, err = buf.ReadInt32(); err != nil {
		return err
	}
	p.DisableRelativeVolume, err = buf.ReadBool()
	return err
}

func (p *S2CLevelEvent) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt32(p.Event); err != nil {
		return err
	}
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteInt32(p.Data); err != nil {
		return err
	}
	return buf.WriteBool(p.DisableRelativeVolume)
}

// S2CLevelParticles represents "Particle".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Particle
type S2CLevelParticles struct {
	LongDistance  ns.Boolean
	AlwaysVisible ns.Boolean
	X             ns.Float64
	Y             ns.Float64
	Z             ns.Float64
	OffsetX       ns.Float32
	OffsetY       ns.Float32
	OffsetZ       ns.Float32
	MaxSpeed      ns.Float32
	ParticleCount ns.Int32
	ParticleId    ns.VarInt
	Data          ns.ByteArray
}

func (p *S2CLevelParticles) ID() ns.VarInt   { return S2CLevelParticlesID }
func (p *S2CLevelParticles) State() jp.State { return jp.StatePlay }
func (p *S2CLevelParticles) Bound() jp.Bound { return jp.S2C }

func (p *S2CLevelParticles) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.LongDistance, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.AlwaysVisible, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Y, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Z, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.OffsetX, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.OffsetY, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.OffsetZ, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.MaxSpeed, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.ParticleCount, err = buf.ReadInt32(); err != nil {
		return err
	}
	if p.ParticleId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CLevelParticles) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteBool(p.LongDistance); err != nil {
		return err
	}
	if err := buf.WriteBool(p.AlwaysVisible); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Y); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.OffsetX); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.OffsetY); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.OffsetZ); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.MaxSpeed); err != nil {
		return err
	}
	if err := buf.WriteInt32(p.ParticleCount); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.ParticleId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CLightUpdate represents "Update Light".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Light
type S2CLightUpdate struct {
	ChunkX ns.VarInt
	ChunkZ ns.VarInt
	Data   ns.ByteArray
}

func (p *S2CLightUpdate) ID() ns.VarInt   { return S2CLightUpdateID }
func (p *S2CLightUpdate) State() jp.State { return jp.StatePlay }
func (p *S2CLightUpdate) Bound() jp.Bound { return jp.S2C }

func (p *S2CLightUpdate) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.ChunkX, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.ChunkZ, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CLightUpdate) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.ChunkX); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.ChunkZ); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CLoginPlay represents "Login (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Login_(Play)
type S2CLoginPlay struct {
	EntityId            ns.Int32
	IsHardcore          ns.Boolean
	DimensionNames      ns.PrefixedArray[ns.Identifier]
	MaxPlayers          ns.VarInt
	ViewDistance        ns.VarInt
	SimulationDistance  ns.VarInt
	ReducedDebugInfo    ns.Boolean
	EnableRespawnScreen ns.Boolean
	DoLimitedCrafting   ns.Boolean
	DimensionType       ns.VarInt
	DimensionName       ns.Identifier
	HashedSeed          ns.Int64
	GameMode            ns.Uint8
	PreviousGameMode    ns.Int8
	IsDebug             ns.Boolean
	IsFlat              ns.Boolean
	DeathLocation       ns.PrefixedOptional[ns.GlobalPos]
	PortalCooldown      ns.VarInt
	SeaLevel            ns.VarInt
	EnforcesSecureChat  ns.Boolean
}

func (p *S2CLoginPlay) ID() ns.VarInt   { return S2CLoginPlayID }
func (p *S2CLoginPlay) State() jp.State { return jp.StatePlay }
func (p *S2CLoginPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CLoginPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadInt32(); err != nil {
		return err
	}
	if p.IsHardcore, err = buf.ReadBool(); err != nil {
		return err
	}
	if err = p.DimensionNames.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.Identifier, error) {
		return b.ReadIdentifier()
	}); err != nil {
		return err
	}
	if p.MaxPlayers, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.ViewDistance, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.SimulationDistance, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.ReducedDebugInfo, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.EnableRespawnScreen, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.DoLimitedCrafting, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.DimensionType, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.DimensionName, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	if p.HashedSeed, err = buf.ReadInt64(); err != nil {
		return err
	}
	if p.GameMode, err = buf.ReadUint8(); err != nil {
		return err
	}
	if p.PreviousGameMode, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.IsDebug, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.IsFlat, err = buf.ReadBool(); err != nil {
		return err
	}
	if err = p.DeathLocation.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.GlobalPos, error) {
		return b.ReadGlobalPos()
	}); err != nil {
		return err
	}
	if p.PortalCooldown, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.SeaLevel, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.EnforcesSecureChat, err = buf.ReadBool()
	return err
}

func (p *S2CLoginPlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt32(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteBool(p.IsHardcore); err != nil {
		return err
	}
	if err := p.DimensionNames.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.Identifier) error {
		return b.WriteIdentifier(v)
	}); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.MaxPlayers); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.ViewDistance); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SimulationDistance); err != nil {
		return err
	}
	if err := buf.WriteBool(p.ReducedDebugInfo); err != nil {
		return err
	}
	if err := buf.WriteBool(p.EnableRespawnScreen); err != nil {
		return err
	}
	if err := buf.WriteBool(p.DoLimitedCrafting); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.DimensionType); err != nil {
		return err
	}
	if err := buf.WriteIdentifier(p.DimensionName); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.HashedSeed); err != nil {
		return err
	}
	if err := buf.WriteUint8(p.GameMode); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.PreviousGameMode); err != nil {
		return err
	}
	if err := buf.WriteBool(p.IsDebug); err != nil {
		return err
	}
	if err := buf.WriteBool(p.IsFlat); err != nil {
		return err
	}
	if err := p.DeathLocation.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.GlobalPos) error {
		return b.WriteGlobalPos(v)
	}); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.PortalCooldown); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SeaLevel); err != nil {
		return err
	}
	return buf.WriteBool(p.EnforcesSecureChat)
}

// S2CMapItemData represents "Map Data".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Map_Data
type S2CMapItemData struct {
	MapId ns.VarInt
	Data  ns.ByteArray
}

func (p *S2CMapItemData) ID() ns.VarInt   { return S2CMapItemDataID }
func (p *S2CMapItemData) State() jp.State { return jp.StatePlay }
func (p *S2CMapItemData) Bound() jp.Bound { return jp.S2C }

func (p *S2CMapItemData) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.MapId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CMapItemData) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.MapId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CMerchantOffers represents "Merchant Offers".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Merchant_Offers
type S2CMerchantOffers struct {
	WindowId ns.VarInt
	Data     ns.ByteArray
}

func (p *S2CMerchantOffers) ID() ns.VarInt   { return S2CMerchantOffersID }
func (p *S2CMerchantOffers) State() jp.State { return jp.StatePlay }
func (p *S2CMerchantOffers) Bound() jp.Bound { return jp.S2C }

func (p *S2CMerchantOffers) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CMerchantOffers) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CMoveEntityPos represents "Update Entity Position".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Entity_Position
type S2CMoveEntityPos struct {
	EntityId ns.VarInt
	DeltaX   ns.Int16
	DeltaY   ns.Int16
	DeltaZ   ns.Int16
	OnGround ns.Boolean
}

func (p *S2CMoveEntityPos) ID() ns.VarInt   { return S2CMoveEntityPosID }
func (p *S2CMoveEntityPos) State() jp.State { return jp.StatePlay }
func (p *S2CMoveEntityPos) Bound() jp.Bound { return jp.S2C }

func (p *S2CMoveEntityPos) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.DeltaX, err = buf.ReadInt16(); err != nil {
		return err
	}
	if p.DeltaY, err = buf.ReadInt16(); err != nil {
		return err
	}
	if p.DeltaZ, err = buf.ReadInt16(); err != nil {
		return err
	}
	p.OnGround, err = buf.ReadBool()
	return err
}

func (p *S2CMoveEntityPos) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteInt16(p.DeltaX); err != nil {
		return err
	}
	if err := buf.WriteInt16(p.DeltaY); err != nil {
		return err
	}
	if err := buf.WriteInt16(p.DeltaZ); err != nil {
		return err
	}
	return buf.WriteBool(p.OnGround)
}

// S2CMoveEntityPosRot represents "Update Entity Position and Rotation".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Entity_Position_And_Rotation
type S2CMoveEntityPosRot struct {
	EntityId ns.VarInt
	DeltaX   ns.Int16
	DeltaY   ns.Int16
	DeltaZ   ns.Int16
	Yaw      ns.Angle
	Pitch    ns.Angle
	OnGround ns.Boolean
}

func (p *S2CMoveEntityPosRot) ID() ns.VarInt   { return S2CMoveEntityPosRotID }
func (p *S2CMoveEntityPosRot) State() jp.State { return jp.StatePlay }
func (p *S2CMoveEntityPosRot) Bound() jp.Bound { return jp.S2C }

func (p *S2CMoveEntityPosRot) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.DeltaX, err = buf.ReadInt16(); err != nil {
		return err
	}
	if p.DeltaY, err = buf.ReadInt16(); err != nil {
		return err
	}
	if p.DeltaZ, err = buf.ReadInt16(); err != nil {
		return err
	}
	if p.Yaw, err = buf.ReadAngle(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadAngle(); err != nil {
		return err
	}
	p.OnGround, err = buf.ReadBool()
	return err
}

func (p *S2CMoveEntityPosRot) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteInt16(p.DeltaX); err != nil {
		return err
	}
	if err := buf.WriteInt16(p.DeltaY); err != nil {
		return err
	}
	if err := buf.WriteInt16(p.DeltaZ); err != nil {
		return err
	}
	if err := buf.WriteAngle(p.Yaw); err != nil {
		return err
	}
	if err := buf.WriteAngle(p.Pitch); err != nil {
		return err
	}
	return buf.WriteBool(p.OnGround)
}

// S2CMoveMinecartAlongTrack represents "Move Minecart Along Track".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Move_Minecart_Along_Track
type S2CMoveMinecartAlongTrack struct {
	EntityId ns.VarInt
	Data     ns.ByteArray
}

func (p *S2CMoveMinecartAlongTrack) ID() ns.VarInt   { return S2CMoveMinecartAlongTrackID }
func (p *S2CMoveMinecartAlongTrack) State() jp.State { return jp.StatePlay }
func (p *S2CMoveMinecartAlongTrack) Bound() jp.Bound { return jp.S2C }

func (p *S2CMoveMinecartAlongTrack) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CMoveMinecartAlongTrack) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CMoveEntityRot represents "Update Entity Rotation".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Entity_Rotation
type S2CMoveEntityRot struct {
	EntityId ns.VarInt
	Yaw      ns.Angle
	Pitch    ns.Angle
	OnGround ns.Boolean
}

func (p *S2CMoveEntityRot) ID() ns.VarInt   { return S2CMoveEntityRotID }
func (p *S2CMoveEntityRot) State() jp.State { return jp.StatePlay }
func (p *S2CMoveEntityRot) Bound() jp.Bound { return jp.S2C }

func (p *S2CMoveEntityRot) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Yaw, err = buf.ReadAngle(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadAngle(); err != nil {
		return err
	}
	p.OnGround, err = buf.ReadBool()
	return err
}

func (p *S2CMoveEntityRot) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteAngle(p.Yaw); err != nil {
		return err
	}
	if err := buf.WriteAngle(p.Pitch); err != nil {
		return err
	}
	return buf.WriteBool(p.OnGround)
}

// S2CMoveVehicle represents "Move Vehicle (clientbound)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Move_Vehicle_(Clientbound)
type S2CMoveVehicle struct {
	X     ns.Float64
	Y     ns.Float64
	Z     ns.Float64
	Yaw   ns.Float32
	Pitch ns.Float32
}

func (p *S2CMoveVehicle) ID() ns.VarInt   { return S2CMoveVehicleID }
func (p *S2CMoveVehicle) State() jp.State { return jp.StatePlay }
func (p *S2CMoveVehicle) Bound() jp.Bound { return jp.S2C }

func (p *S2CMoveVehicle) Read(buf *ns.PacketBuffer) error {
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
	p.Pitch, err = buf.ReadFloat32()
	return err
}

func (p *S2CMoveVehicle) Write(buf *ns.PacketBuffer) error {
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
	return buf.WriteFloat32(p.Pitch)
}

// S2COpenBook represents "Open Book".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Open_Book
type S2COpenBook struct {
	Hand ns.VarInt
}

func (p *S2COpenBook) ID() ns.VarInt   { return S2COpenBookID }
func (p *S2COpenBook) State() jp.State { return jp.StatePlay }
func (p *S2COpenBook) Bound() jp.Bound { return jp.S2C }

func (p *S2COpenBook) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Hand, err = buf.ReadVarInt()
	return err
}

func (p *S2COpenBook) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.Hand)
}

// S2COpenScreen represents "Open Screen".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Open_Screen
type S2COpenScreen struct {
	WindowId    ns.VarInt
	WindowType  ns.VarInt
	WindowTitle ns.TextComponent
}

func (p *S2COpenScreen) ID() ns.VarInt   { return S2COpenScreenID }
func (p *S2COpenScreen) State() jp.State { return jp.StatePlay }
func (p *S2COpenScreen) Bound() jp.Bound { return jp.S2C }

func (p *S2COpenScreen) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.WindowType, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.WindowTitle, err = buf.ReadTextComponent()
	return err
}

func (p *S2COpenScreen) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.WindowType); err != nil {
		return err
	}
	return buf.WriteTextComponent(p.WindowTitle)
}

// S2COpenSignEditor represents "Open Sign Editor".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Open_Sign_Editor
type S2COpenSignEditor struct {
	Location    ns.Position
	IsFrontText ns.Boolean
}

func (p *S2COpenSignEditor) ID() ns.VarInt   { return S2COpenSignEditorID }
func (p *S2COpenSignEditor) State() jp.State { return jp.StatePlay }
func (p *S2COpenSignEditor) Bound() jp.Bound { return jp.S2C }

func (p *S2COpenSignEditor) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	p.IsFrontText, err = buf.ReadBool()
	return err
}

func (p *S2COpenSignEditor) Write(buf *ns.PacketBuffer) error {
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	return buf.WriteBool(p.IsFrontText)
}

// S2CPingPlay represents "Ping (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Ping_(Play)
type S2CPingPlay struct {
	Id ns.Int32
}

func (p *S2CPingPlay) ID() ns.VarInt   { return S2CPingPlayID }
func (p *S2CPingPlay) State() jp.State { return jp.StatePlay }
func (p *S2CPingPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CPingPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Id, err = buf.ReadInt32()
	return err
}

func (p *S2CPingPlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt32(p.Id)
}

// S2CPongResponsePlay represents "Ping Response (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Ping_Response_(Play)
type S2CPongResponsePlay struct {
	Payload ns.Int64
}

func (p *S2CPongResponsePlay) ID() ns.VarInt   { return S2CPongResponsePlayID }
func (p *S2CPongResponsePlay) State() jp.State { return jp.StatePlay }
func (p *S2CPongResponsePlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CPongResponsePlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Payload, err = buf.ReadInt64()
	return err
}

func (p *S2CPongResponsePlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt64(p.Payload)
}

// S2CPlaceGhostRecipe represents "Place Ghost Recipe".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Place_Ghost_Recipe
type S2CPlaceGhostRecipe struct {
	WindowId      ns.VarInt
	RecipeDisplay ns.ByteArray
}

func (p *S2CPlaceGhostRecipe) ID() ns.VarInt   { return S2CPlaceGhostRecipeID }
func (p *S2CPlaceGhostRecipe) State() jp.State { return jp.StatePlay }
func (p *S2CPlaceGhostRecipe) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlaceGhostRecipe) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WindowId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.RecipeDisplay, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CPlaceGhostRecipe) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.WindowId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.RecipeDisplay)
}

// S2CPlayerAbilities represents "Player Abilities (clientbound)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Abilities_(Clientbound)
type S2CPlayerAbilities struct {
	Flags               ns.Int8
	FlyingSpeed         ns.Float32
	FieldOfViewModifier ns.Float32
}

func (p *S2CPlayerAbilities) ID() ns.VarInt   { return S2CPlayerAbilitiesID }
func (p *S2CPlayerAbilities) State() jp.State { return jp.StatePlay }
func (p *S2CPlayerAbilities) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlayerAbilities) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Flags, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.FlyingSpeed, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.FieldOfViewModifier, err = buf.ReadFloat32()
	return err
}

func (p *S2CPlayerAbilities) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt8(p.Flags); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.FlyingSpeed); err != nil {
		return err
	}
	return buf.WriteFloat32(p.FieldOfViewModifier)
}

// S2CPlayerChat represents "Player Chat Message".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Chat_Message
type S2CPlayerChat struct {
	GlobalIndex     ns.VarInt
	Sender          ns.UUID
	Index           ns.VarInt
	Signature       ns.PrefixedOptional[MessageSignature]
	Body            SignedMessageBody
	UnsignedContent ns.PrefixedOptional[ns.TextComponent]
	FilterMask      FilterMask
	ChatType        ChatTypeBound
}

// MessageSignature is a 256-byte cryptographic signature for chat messages.
type MessageSignature [256]byte

// SignedMessageBody contains the signed content of a chat message.
type SignedMessageBody struct {
	Content   ns.String
	Timestamp ns.Int64 // epoch milliseconds
	Salt      ns.Int64
	LastSeen  LastSeenMessagesPacked
}

// LastSeenMessagesPacked contains references to previously seen messages.
type LastSeenMessagesPacked struct {
	Entries ns.PrefixedArray[MessageSignaturePacked]
}

// MessageSignaturePacked is either a cache index or a full signature.
type MessageSignaturePacked struct {
	ID            ns.VarInt         // id + 1; if 0, FullSignature follows
	FullSignature *MessageSignature // only present if ID == 0
}

// FilterMask indicates how a message should be filtered.
type FilterMask struct {
	Type FilterMaskType
	Mask *ns.BitSet // only present if Type == FilterMaskPartiallyFiltered
}

// FilterMaskType indicates the type of filter mask.
type FilterMaskType ns.VarInt

const (
	FilterMaskPassThrough FilterMaskType = iota
	FilterMaskFullyFiltered
	FilterMaskPartiallyFiltered
)

// ChatTypeBound contains the chat type and sender information.
type ChatTypeBound struct {
	// Registry ID. In vanilla: 1 = normal chat, 3 = whisper, 5 = /say, 2 = /me, ...
	// TODO: accumulate the registry data in client store and reference in higher level funcs
	ChatType   ns.VarInt
	Name       ns.TextComponent                      // sender's display name
	TargetName ns.PrefixedOptional[ns.TextComponent] // target's display name (for whispers)
}

func (p *S2CPlayerChat) ID() ns.VarInt   { return S2CPlayerChatID }
func (p *S2CPlayerChat) State() jp.State { return jp.StatePlay }
func (p *S2CPlayerChat) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlayerChat) Read(buf *ns.PacketBuffer) error {
	var err error

	if p.GlobalIndex, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Sender, err = buf.ReadUUID(); err != nil {
		return err
	}
	if p.Index, err = buf.ReadVarInt(); err != nil {
		return err
	}

	// Signature (optional)
	if err = p.Signature.DecodeWith(buf, func(b *ns.PacketBuffer) (MessageSignature, error) {
		var sig MessageSignature
		data, err := b.ReadFixedByteArray(256)
		if err != nil {
			return sig, err
		}
		copy(sig[:], data)
		return sig, nil
	}); err != nil {
		return err
	}

	// Body
	if p.Body.Content, err = buf.ReadString(256); err != nil {
		return err
	}
	if p.Body.Timestamp, err = buf.ReadInt64(); err != nil {
		return err
	}
	if p.Body.Salt, err = buf.ReadInt64(); err != nil {
		return err
	}

	// LastSeen
	if err = p.Body.LastSeen.Entries.DecodeWith(buf, func(b *ns.PacketBuffer) (MessageSignaturePacked, error) {
		var msp MessageSignaturePacked
		id, err := b.ReadVarInt()
		if err != nil {
			return msp, err
		}
		msp.ID = id
		if id == 0 {
			var sig MessageSignature
			data, err := b.ReadFixedByteArray(256)
			if err != nil {
				return msp, err
			}
			copy(sig[:], data)
			msp.FullSignature = &sig
		}
		return msp, nil
	}); err != nil {
		return err
	}

	// UnsignedContent (optional)
	if err = p.UnsignedContent.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.TextComponent, error) {
		return b.ReadTextComponent()
	}); err != nil {
		return err
	}

	// FilterMask
	filterType, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	p.FilterMask.Type = FilterMaskType(filterType)
	if p.FilterMask.Type == FilterMaskPartiallyFiltered {
		bitset := &ns.BitSet{}
		if err = bitset.Decode(buf); err != nil {
			return err
		}
		p.FilterMask.Mask = bitset
	}

	// ChatType
	if p.ChatType.ChatType, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.ChatType.Name, err = buf.ReadTextComponent(); err != nil {
		return err
	}
	if err = p.ChatType.TargetName.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.TextComponent, error) {
		return b.ReadTextComponent()
	}); err != nil {
		return err
	}

	return nil
}

func (p *S2CPlayerChat) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.GlobalIndex); err != nil {
		return err
	}
	if err := buf.WriteUUID(p.Sender); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Index); err != nil {
		return err
	}

	// Signature
	if err := p.Signature.EncodeWith(buf, func(b *ns.PacketBuffer, sig MessageSignature) error {
		return b.WriteFixedByteArray(sig[:])
	}); err != nil {
		return err
	}

	// Body
	if err := buf.WriteString(p.Body.Content); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.Body.Timestamp); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.Body.Salt); err != nil {
		return err
	}

	// LastSeen
	if err := p.Body.LastSeen.Entries.EncodeWith(buf, func(b *ns.PacketBuffer, msp MessageSignaturePacked) error {
		if err := b.WriteVarInt(msp.ID); err != nil {
			return err
		}
		if msp.ID == 0 && msp.FullSignature != nil {
			return b.WriteFixedByteArray(msp.FullSignature[:])
		}
		return nil
	}); err != nil {
		return err
	}

	// UnsignedContent
	if err := p.UnsignedContent.EncodeWith(buf, func(b *ns.PacketBuffer, tc ns.TextComponent) error {
		return b.WriteTextComponent(tc)
	}); err != nil {
		return err
	}

	// FilterMask
	if err := buf.WriteVarInt(ns.VarInt(p.FilterMask.Type)); err != nil {
		return err
	}
	if p.FilterMask.Type == FilterMaskPartiallyFiltered && p.FilterMask.Mask != nil {
		if err := p.FilterMask.Mask.Encode(buf); err != nil {
			return err
		}
	}

	// ChatType
	if err := buf.WriteVarInt(p.ChatType.ChatType); err != nil {
		return err
	}
	if err := buf.WriteTextComponent(p.ChatType.Name); err != nil {
		return err
	}
	if err := p.ChatType.TargetName.EncodeWith(buf, func(b *ns.PacketBuffer, tc ns.TextComponent) error {
		return b.WriteTextComponent(tc)
	}); err != nil {
		return err
	}

	return nil
}

// GetMessage returns the chat message content.
// If there's unsigned content, it returns that; otherwise returns the signed body content.
func (p *S2CPlayerChat) GetMessage() string {
	if p.UnsignedContent.Present {
		return p.UnsignedContent.Value.PlainText()
	}
	return string(p.Body.Content)
}

// GetSenderName returns the sender's display name as plain text.
func (p *S2CPlayerChat) GetSenderName() string {
	return p.ChatType.Name.PlainText()
}

// S2CPlayerCombatEnd represents "End Combat".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#End_Combat
type S2CPlayerCombatEnd struct {
	Duration ns.VarInt
}

func (p *S2CPlayerCombatEnd) ID() ns.VarInt   { return S2CPlayerCombatEndID }
func (p *S2CPlayerCombatEnd) State() jp.State { return jp.StatePlay }
func (p *S2CPlayerCombatEnd) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlayerCombatEnd) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Duration, err = buf.ReadVarInt()
	return err
}

func (p *S2CPlayerCombatEnd) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.Duration)
}

// S2CPlayerCombatEnter represents "Enter Combat".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Enter_Combat
type S2CPlayerCombatEnter struct{}

func (p *S2CPlayerCombatEnter) ID() ns.VarInt                { return S2CPlayerCombatEnterID }
func (p *S2CPlayerCombatEnter) State() jp.State              { return jp.StatePlay }
func (p *S2CPlayerCombatEnter) Bound() jp.Bound              { return jp.S2C }
func (p *S2CPlayerCombatEnter) Read(*ns.PacketBuffer) error  { return nil }
func (p *S2CPlayerCombatEnter) Write(*ns.PacketBuffer) error { return nil }

// S2CPlayerCombatKill represents "Combat Death".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Combat_Death
type S2CPlayerCombatKill struct {
	PlayerId ns.VarInt
	Message  ns.TextComponent
}

func (p *S2CPlayerCombatKill) ID() ns.VarInt   { return S2CPlayerCombatKillID }
func (p *S2CPlayerCombatKill) State() jp.State { return jp.StatePlay }
func (p *S2CPlayerCombatKill) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlayerCombatKill) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.PlayerId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Message, err = buf.ReadTextComponent()
	return err
}

func (p *S2CPlayerCombatKill) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.PlayerId); err != nil {
		return err
	}
	return buf.WriteTextComponent(p.Message)
}

// S2CPlayerInfoRemove represents "Player Info Remove".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Info_Remove
type S2CPlayerInfoRemove struct {
	Uuids ns.ByteArray
}

func (p *S2CPlayerInfoRemove) ID() ns.VarInt   { return S2CPlayerInfoRemoveID }
func (p *S2CPlayerInfoRemove) State() jp.State { return jp.StatePlay }
func (p *S2CPlayerInfoRemove) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlayerInfoRemove) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Uuids, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CPlayerInfoRemove) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Uuids)
}

// S2CPlayerInfoUpdate represents "Player Info Update".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Info_Update
type S2CPlayerInfoUpdate struct {
	Data ns.ByteArray
}

func (p *S2CPlayerInfoUpdate) ID() ns.VarInt   { return S2CPlayerInfoUpdateID }
func (p *S2CPlayerInfoUpdate) State() jp.State { return jp.StatePlay }
func (p *S2CPlayerInfoUpdate) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlayerInfoUpdate) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CPlayerInfoUpdate) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Data)
}

// S2CPlayerLookAt represents "Look At".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Look_At
type S2CPlayerLookAt struct {
	FeetEyes       ns.VarInt
	TargetX        ns.Float64
	TargetY        ns.Float64
	TargetZ        ns.Float64
	IsEntity       ns.Boolean
	EntityId       ns.VarInt
	EntityFeetEyes ns.VarInt
}

func (p *S2CPlayerLookAt) ID() ns.VarInt   { return S2CPlayerLookAtID }
func (p *S2CPlayerLookAt) State() jp.State { return jp.StatePlay }
func (p *S2CPlayerLookAt) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlayerLookAt) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.FeetEyes, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.TargetX, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.TargetY, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.TargetZ, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.IsEntity, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.IsEntity {
		if p.EntityId, err = buf.ReadVarInt(); err != nil {
			return err
		}
		p.EntityFeetEyes, err = buf.ReadVarInt()
	}
	return err
}

func (p *S2CPlayerLookAt) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.FeetEyes); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.TargetX); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.TargetY); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.TargetZ); err != nil {
		return err
	}
	if err := buf.WriteBool(p.IsEntity); err != nil {
		return err
	}
	if p.IsEntity {
		if err := buf.WriteVarInt(p.EntityId); err != nil {
			return err
		}
		return buf.WriteVarInt(p.EntityFeetEyes)
	}
	return nil
}

// S2CPlayerPosition represents "Synchronize Player Position".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Synchronize_Player_Position
type S2CPlayerPosition struct {
	TeleportId ns.VarInt
	X          ns.Float64
	Y          ns.Float64
	Z          ns.Float64
	VelocityX  ns.Float64
	VelocityY  ns.Float64
	VelocityZ  ns.Float64
	Yaw        ns.Float32
	Pitch      ns.Float32
	Flags      ns.Int32
}

func (p *S2CPlayerPosition) ID() ns.VarInt   { return S2CPlayerPositionID }
func (p *S2CPlayerPosition) State() jp.State { return jp.StatePlay }
func (p *S2CPlayerPosition) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlayerPosition) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.TeleportId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Y, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Z, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.VelocityX, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.VelocityY, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.VelocityZ, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Yaw, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.Flags, err = buf.ReadInt32()
	return err
}

func (p *S2CPlayerPosition) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.TeleportId); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Y); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.VelocityX); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.VelocityY); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.VelocityZ); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Yaw); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Pitch); err != nil {
		return err
	}
	return buf.WriteInt32(p.Flags)
}

// S2CPlayerRotation represents "Player Rotation".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Rotation
type S2CPlayerRotation struct {
	Yaw           ns.Float32
	RelativeYaw   ns.Boolean
	Pitch         ns.Float32
	RelativePitch ns.Boolean
}

func (p *S2CPlayerRotation) ID() ns.VarInt   { return S2CPlayerRotationID }
func (p *S2CPlayerRotation) State() jp.State { return jp.StatePlay }
func (p *S2CPlayerRotation) Bound() jp.Bound { return jp.S2C }

func (p *S2CPlayerRotation) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Yaw, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.RelativeYaw, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.RelativePitch, err = buf.ReadBool()
	return err
}

func (p *S2CPlayerRotation) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat32(p.Yaw); err != nil {
		return err
	}
	if err := buf.WriteBool(p.RelativeYaw); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Pitch); err != nil {
		return err
	}
	return buf.WriteBool(p.RelativePitch)
}

// S2CRecipeBookAdd represents "Recipe Book Add".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Recipe_Book_Add
type S2CRecipeBookAdd struct {
	Data ns.ByteArray
}

func (p *S2CRecipeBookAdd) ID() ns.VarInt   { return S2CRecipeBookAddID }
func (p *S2CRecipeBookAdd) State() jp.State { return jp.StatePlay }
func (p *S2CRecipeBookAdd) Bound() jp.Bound { return jp.S2C }

func (p *S2CRecipeBookAdd) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CRecipeBookAdd) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Data)
}

// S2CRecipeBookRemove represents "Recipe Book Remove".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Recipe_Book_Remove
type S2CRecipeBookRemove struct {
	Recipes ns.ByteArray
}

func (p *S2CRecipeBookRemove) ID() ns.VarInt   { return S2CRecipeBookRemoveID }
func (p *S2CRecipeBookRemove) State() jp.State { return jp.StatePlay }
func (p *S2CRecipeBookRemove) Bound() jp.Bound { return jp.S2C }

func (p *S2CRecipeBookRemove) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Recipes, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CRecipeBookRemove) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Recipes)
}

// S2CRecipeBookSettings represents "Recipe Book Settings".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Recipe_Book_Settings
type S2CRecipeBookSettings struct {
	CraftingRecipeBookOpen             ns.Boolean
	CraftingRecipeBookFilterActive     ns.Boolean
	SmeltingRecipeBookOpen             ns.Boolean
	SmeltingRecipeBookFilterActive     ns.Boolean
	BlastFurnaceRecipeBookOpen         ns.Boolean
	BlastFurnaceRecipeBookFilterActive ns.Boolean
	SmokerRecipeBookOpen               ns.Boolean
	SmokerRecipeBookFilterActive       ns.Boolean
}

func (p *S2CRecipeBookSettings) ID() ns.VarInt   { return S2CRecipeBookSettingsID }
func (p *S2CRecipeBookSettings) State() jp.State { return jp.StatePlay }
func (p *S2CRecipeBookSettings) Bound() jp.Bound { return jp.S2C }

func (p *S2CRecipeBookSettings) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.CraftingRecipeBookOpen, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.CraftingRecipeBookFilterActive, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.SmeltingRecipeBookOpen, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.SmeltingRecipeBookFilterActive, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.BlastFurnaceRecipeBookOpen, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.BlastFurnaceRecipeBookFilterActive, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.SmokerRecipeBookOpen, err = buf.ReadBool(); err != nil {
		return err
	}
	p.SmokerRecipeBookFilterActive, err = buf.ReadBool()
	return err
}

func (p *S2CRecipeBookSettings) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteBool(p.CraftingRecipeBookOpen); err != nil {
		return err
	}
	if err := buf.WriteBool(p.CraftingRecipeBookFilterActive); err != nil {
		return err
	}
	if err := buf.WriteBool(p.SmeltingRecipeBookOpen); err != nil {
		return err
	}
	if err := buf.WriteBool(p.SmeltingRecipeBookFilterActive); err != nil {
		return err
	}
	if err := buf.WriteBool(p.BlastFurnaceRecipeBookOpen); err != nil {
		return err
	}
	if err := buf.WriteBool(p.BlastFurnaceRecipeBookFilterActive); err != nil {
		return err
	}
	if err := buf.WriteBool(p.SmokerRecipeBookOpen); err != nil {
		return err
	}
	return buf.WriteBool(p.SmokerRecipeBookFilterActive)
}

// S2CRemoveEntities represents "Remove Entities".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Remove_Entities
type S2CRemoveEntities struct {
	EntityIds ns.ByteArray
}

func (p *S2CRemoveEntities) ID() ns.VarInt   { return S2CRemoveEntitiesID }
func (p *S2CRemoveEntities) State() jp.State { return jp.StatePlay }
func (p *S2CRemoveEntities) Bound() jp.Bound { return jp.S2C }

func (p *S2CRemoveEntities) Read(buf *ns.PacketBuffer) error {
	var err error
	p.EntityIds, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CRemoveEntities) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.EntityIds)
}

// S2CRemoveMobEffect represents "Remove Entity Effect".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Remove_Entity_Effect
type S2CRemoveMobEffect struct {
	EntityId ns.VarInt
	EffectId ns.VarInt
}

func (p *S2CRemoveMobEffect) ID() ns.VarInt   { return S2CRemoveMobEffectID }
func (p *S2CRemoveMobEffect) State() jp.State { return jp.StatePlay }
func (p *S2CRemoveMobEffect) Bound() jp.Bound { return jp.S2C }

func (p *S2CRemoveMobEffect) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.EffectId, err = buf.ReadVarInt()
	return err
}

func (p *S2CRemoveMobEffect) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteVarInt(p.EffectId)
}

// S2CResetScore represents "Reset Score".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Reset_Score
type S2CResetScore struct {
	EntityName    ns.String
	ObjectiveName ns.PrefixedOptional[ns.String]
}

func (p *S2CResetScore) ID() ns.VarInt   { return S2CResetScoreID }
func (p *S2CResetScore) State() jp.State { return jp.StatePlay }
func (p *S2CResetScore) Bound() jp.Bound { return jp.S2C }

func (p *S2CResetScore) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityName, err = buf.ReadString(32767); err != nil {
		return err
	}
	return p.ObjectiveName.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.String, error) {
		return b.ReadString(32767)
	})
}

func (p *S2CResetScore) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteString(p.EntityName); err != nil {
		return err
	}
	return p.ObjectiveName.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.String) error {
		return b.WriteString(v)
	})
}

// S2CResourcePackPopPlay represents "Remove Resource Pack (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Remove_Resource_Pack_(Play)
type S2CResourcePackPopPlay struct {
	Uuid ns.PrefixedOptional[ns.UUID]
}

func (p *S2CResourcePackPopPlay) ID() ns.VarInt   { return S2CResourcePackPopPlayID }
func (p *S2CResourcePackPopPlay) State() jp.State { return jp.StatePlay }
func (p *S2CResourcePackPopPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CResourcePackPopPlay) Read(buf *ns.PacketBuffer) error {
	return p.Uuid.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.UUID, error) {
		return b.ReadUUID()
	})
}

func (p *S2CResourcePackPopPlay) Write(buf *ns.PacketBuffer) error {
	return p.Uuid.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.UUID) error {
		return b.WriteUUID(v)
	})
}

// S2CResourcePackPushPlay represents "Add Resource Pack (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Add_Resource_Pack_(Play)
type S2CResourcePackPushPlay struct {
	Uuid          ns.UUID
	Url           ns.String
	Hash          ns.String
	Forced        ns.Boolean
	PromptMessage ns.PrefixedOptional[ns.TextComponent]
}

func (p *S2CResourcePackPushPlay) ID() ns.VarInt   { return S2CResourcePackPushPlayID }
func (p *S2CResourcePackPushPlay) State() jp.State { return jp.StatePlay }
func (p *S2CResourcePackPushPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CResourcePackPushPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Uuid, err = buf.ReadUUID(); err != nil {
		return err
	}
	if p.Url, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.Hash, err = buf.ReadString(40); err != nil {
		return err
	}
	if p.Forced, err = buf.ReadBool(); err != nil {
		return err
	}
	return p.PromptMessage.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.TextComponent, error) {
		return b.ReadTextComponent()
	})
}

func (p *S2CResourcePackPushPlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteUUID(p.Uuid); err != nil {
		return err
	}
	if err := buf.WriteString(p.Url); err != nil {
		return err
	}
	if err := buf.WriteString(p.Hash); err != nil {
		return err
	}
	if err := buf.WriteBool(p.Forced); err != nil {
		return err
	}
	return p.PromptMessage.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.TextComponent) error {
		return b.WriteTextComponent(v)
	})
}

// S2CRespawn represents "Respawn".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Respawn
type S2CRespawn struct {
	DimensionType    ns.VarInt
	DimensionName    ns.Identifier
	HashedSeed       ns.Int64
	GameMode         ns.Uint8
	PreviousGameMode ns.Int8
	IsDebug          ns.Boolean
	IsFlat           ns.Boolean
	DeathLocation    ns.PrefixedOptional[ns.ByteArray]
	PortalCooldown   ns.VarInt
	SeaLevel         ns.VarInt
	DataKept         ns.Int8
}

func (p *S2CRespawn) ID() ns.VarInt   { return S2CRespawnID }
func (p *S2CRespawn) State() jp.State { return jp.StatePlay }
func (p *S2CRespawn) Bound() jp.Bound { return jp.S2C }

func (p *S2CRespawn) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.DimensionType, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.DimensionName, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	if p.HashedSeed, err = buf.ReadInt64(); err != nil {
		return err
	}
	if p.GameMode, err = buf.ReadUint8(); err != nil {
		return err
	}
	if p.PreviousGameMode, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.IsDebug, err = buf.ReadBool(); err != nil {
		return err
	}
	if p.IsFlat, err = buf.ReadBool(); err != nil {
		return err
	}
	if err = p.DeathLocation.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.ByteArray, error) {
		return b.ReadByteArray(1048576)
	}); err != nil {
		return err
	}
	if p.PortalCooldown, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.SeaLevel, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.DataKept, err = buf.ReadInt8()
	return err
}

func (p *S2CRespawn) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.DimensionType); err != nil {
		return err
	}
	if err := buf.WriteIdentifier(p.DimensionName); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.HashedSeed); err != nil {
		return err
	}
	if err := buf.WriteUint8(p.GameMode); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.PreviousGameMode); err != nil {
		return err
	}
	if err := buf.WriteBool(p.IsDebug); err != nil {
		return err
	}
	if err := buf.WriteBool(p.IsFlat); err != nil {
		return err
	}
	if err := p.DeathLocation.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.ByteArray) error {
		return b.WriteByteArray(v)
	}); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.PortalCooldown); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SeaLevel); err != nil {
		return err
	}
	return buf.WriteInt8(p.DataKept)
}

// S2CRotateHead represents "Set Head Rotation".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Head_Rotation
type S2CRotateHead struct {
	EntityId ns.VarInt
	HeadYaw  ns.Angle
}

func (p *S2CRotateHead) ID() ns.VarInt   { return S2CRotateHeadID }
func (p *S2CRotateHead) State() jp.State { return jp.StatePlay }
func (p *S2CRotateHead) Bound() jp.Bound { return jp.S2C }

func (p *S2CRotateHead) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.HeadYaw, err = buf.ReadAngle()
	return err
}

func (p *S2CRotateHead) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteAngle(p.HeadYaw)
}

// S2CSectionBlocksUpdate represents "Update Section Blocks".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Section_Blocks
type S2CSectionBlocksUpdate struct {
	ChunkSectionPosition ns.Int64
	Blocks               ns.ByteArray
}

func (p *S2CSectionBlocksUpdate) ID() ns.VarInt   { return S2CSectionBlocksUpdateID }
func (p *S2CSectionBlocksUpdate) State() jp.State { return jp.StatePlay }
func (p *S2CSectionBlocksUpdate) Bound() jp.Bound { return jp.S2C }

func (p *S2CSectionBlocksUpdate) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.ChunkSectionPosition, err = buf.ReadInt64(); err != nil {
		return err
	}
	p.Blocks, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CSectionBlocksUpdate) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt64(p.ChunkSectionPosition); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Blocks)
}

// S2CSelectAdvancementsTab represents "Select Advancements Tab".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Select_Advancements_Tab
type S2CSelectAdvancementsTab struct {
	Identifier ns.PrefixedOptional[ns.Identifier]
}

func (p *S2CSelectAdvancementsTab) ID() ns.VarInt   { return S2CSelectAdvancementsTabID }
func (p *S2CSelectAdvancementsTab) State() jp.State { return jp.StatePlay }
func (p *S2CSelectAdvancementsTab) Bound() jp.Bound { return jp.S2C }

func (p *S2CSelectAdvancementsTab) Read(buf *ns.PacketBuffer) error {
	return p.Identifier.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.Identifier, error) {
		return b.ReadIdentifier()
	})
}

func (p *S2CSelectAdvancementsTab) Write(buf *ns.PacketBuffer) error {
	return p.Identifier.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.Identifier) error {
		return b.WriteIdentifier(v)
	})
}

// S2CServerData represents "Server Data".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Server_Data
type S2CServerData struct {
	Motd ns.TextComponent
	Icon ns.PrefixedOptional[ns.ByteArray]
}

func (p *S2CServerData) ID() ns.VarInt   { return S2CServerDataID }
func (p *S2CServerData) State() jp.State { return jp.StatePlay }
func (p *S2CServerData) Bound() jp.Bound { return jp.S2C }

func (p *S2CServerData) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Motd, err = buf.ReadTextComponent(); err != nil {
		return err
	}
	return p.Icon.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.ByteArray, error) {
		return b.ReadByteArray(1048576)
	})
}

func (p *S2CServerData) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteTextComponent(p.Motd); err != nil {
		return err
	}
	return p.Icon.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.ByteArray) error {
		return b.WriteByteArray(v)
	})
}

// S2CSetActionBarText represents "Set Action Bar Text".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Action_Bar_Text
type S2CSetActionBarText struct {
	Text ns.TextComponent
}

func (p *S2CSetActionBarText) ID() ns.VarInt   { return S2CSetActionBarTextID }
func (p *S2CSetActionBarText) State() jp.State { return jp.StatePlay }
func (p *S2CSetActionBarText) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetActionBarText) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Text, err = buf.ReadTextComponent()
	return err
}

func (p *S2CSetActionBarText) Write(buf *ns.PacketBuffer) error {
	return buf.WriteTextComponent(p.Text)
}

// S2CSetBorderCenter represents "Set Border Center".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Center
type S2CSetBorderCenter struct {
	X ns.Float64
	Z ns.Float64
}

func (p *S2CSetBorderCenter) ID() ns.VarInt   { return S2CSetBorderCenterID }
func (p *S2CSetBorderCenter) State() jp.State { return jp.StatePlay }
func (p *S2CSetBorderCenter) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetBorderCenter) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	p.Z, err = buf.ReadFloat64()
	return err
}

func (p *S2CSetBorderCenter) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	return buf.WriteFloat64(p.Z)
}

// S2CSetBorderLerpSize represents "Set Border Lerp Size".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Lerp_Size
type S2CSetBorderLerpSize struct {
	OldDiameter ns.Float64
	NewDiameter ns.Float64
	Speed       ns.VarLong
}

func (p *S2CSetBorderLerpSize) ID() ns.VarInt   { return S2CSetBorderLerpSizeID }
func (p *S2CSetBorderLerpSize) State() jp.State { return jp.StatePlay }
func (p *S2CSetBorderLerpSize) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetBorderLerpSize) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.OldDiameter, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.NewDiameter, err = buf.ReadFloat64(); err != nil {
		return err
	}
	p.Speed, err = buf.ReadVarLong()
	return err
}

func (p *S2CSetBorderLerpSize) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat64(p.OldDiameter); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.NewDiameter); err != nil {
		return err
	}
	return buf.WriteVarLong(p.Speed)
}

// S2CSetBorderSize represents "Set Border Size".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Size
type S2CSetBorderSize struct {
	Diameter ns.Float64
}

func (p *S2CSetBorderSize) ID() ns.VarInt   { return S2CSetBorderSizeID }
func (p *S2CSetBorderSize) State() jp.State { return jp.StatePlay }
func (p *S2CSetBorderSize) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetBorderSize) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Diameter, err = buf.ReadFloat64()
	return err
}

func (p *S2CSetBorderSize) Write(buf *ns.PacketBuffer) error {
	return buf.WriteFloat64(p.Diameter)
}

// S2CSetBorderWarningDelay represents "Set Border Warning Delay".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Warning_Delay
type S2CSetBorderWarningDelay struct {
	WarningTime ns.VarInt
}

func (p *S2CSetBorderWarningDelay) ID() ns.VarInt   { return S2CSetBorderWarningDelayID }
func (p *S2CSetBorderWarningDelay) State() jp.State { return jp.StatePlay }
func (p *S2CSetBorderWarningDelay) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetBorderWarningDelay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.WarningTime, err = buf.ReadVarInt()
	return err
}

func (p *S2CSetBorderWarningDelay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.WarningTime)
}

// S2CSetBorderWarningDistance represents "Set Border Warning Distance".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Warning_Distance
type S2CSetBorderWarningDistance struct {
	WarningBlocks ns.VarInt
}

func (p *S2CSetBorderWarningDistance) ID() ns.VarInt   { return S2CSetBorderWarningDistanceID }
func (p *S2CSetBorderWarningDistance) State() jp.State { return jp.StatePlay }
func (p *S2CSetBorderWarningDistance) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetBorderWarningDistance) Read(buf *ns.PacketBuffer) error {
	var err error
	p.WarningBlocks, err = buf.ReadVarInt()
	return err
}

func (p *S2CSetBorderWarningDistance) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.WarningBlocks)
}

// S2CSetCamera represents "Set Camera".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Camera
type S2CSetCamera struct {
	CameraId ns.VarInt
}

func (p *S2CSetCamera) ID() ns.VarInt   { return S2CSetCameraID }
func (p *S2CSetCamera) State() jp.State { return jp.StatePlay }
func (p *S2CSetCamera) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetCamera) Read(buf *ns.PacketBuffer) error {
	var err error
	p.CameraId, err = buf.ReadVarInt()
	return err
}

func (p *S2CSetCamera) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.CameraId)
}

// S2CSetChunkCacheCenter represents "Set Center Chunk".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Center_Chunk
type S2CSetChunkCacheCenter struct {
	ChunkX ns.VarInt
	ChunkZ ns.VarInt
}

func (p *S2CSetChunkCacheCenter) ID() ns.VarInt   { return S2CSetChunkCacheCenterID }
func (p *S2CSetChunkCacheCenter) State() jp.State { return jp.StatePlay }
func (p *S2CSetChunkCacheCenter) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetChunkCacheCenter) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.ChunkX, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.ChunkZ, err = buf.ReadVarInt()
	return err
}

func (p *S2CSetChunkCacheCenter) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.ChunkX); err != nil {
		return err
	}
	return buf.WriteVarInt(p.ChunkZ)
}

// S2CSetChunkCacheRadius represents "Set Render Distance".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Render_Distance
type S2CSetChunkCacheRadius struct {
	ViewDistance ns.VarInt
}

func (p *S2CSetChunkCacheRadius) ID() ns.VarInt   { return S2CSetChunkCacheRadiusID }
func (p *S2CSetChunkCacheRadius) State() jp.State { return jp.StatePlay }
func (p *S2CSetChunkCacheRadius) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetChunkCacheRadius) Read(buf *ns.PacketBuffer) error {
	var err error
	p.ViewDistance, err = buf.ReadVarInt()
	return err
}

func (p *S2CSetChunkCacheRadius) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.ViewDistance)
}

// S2CSetCursorItem represents "Set Cursor Item".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Cursor_Item
type S2CSetCursorItem struct {
	CarriedItem ns.Slot
}

func (p *S2CSetCursorItem) ID() ns.VarInt   { return S2CSetCursorItemID }
func (p *S2CSetCursorItem) State() jp.State { return jp.StatePlay }
func (p *S2CSetCursorItem) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetCursorItem) Read(buf *ns.PacketBuffer) error {
	var err error
	p.CarriedItem, err = buf.ReadSlot()
	return err
}

func (p *S2CSetCursorItem) Write(buf *ns.PacketBuffer) error {
	return buf.WriteSlot(p.CarriedItem)
}

// S2CSetDefaultSpawnPosition represents "Set Default Spawn Position".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Default_Spawn_Position
type S2CSetDefaultSpawnPosition struct {
	DimensionName ns.Identifier
	Location      ns.Position
	Yaw           ns.Float32
	Pitch         ns.Float32
}

func (p *S2CSetDefaultSpawnPosition) ID() ns.VarInt   { return S2CSetDefaultSpawnPositionID }
func (p *S2CSetDefaultSpawnPosition) State() jp.State { return jp.StatePlay }
func (p *S2CSetDefaultSpawnPosition) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetDefaultSpawnPosition) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.DimensionName, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	if p.Location, err = buf.ReadPosition(); err != nil {
		return err
	}
	if p.Yaw, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.Pitch, err = buf.ReadFloat32()
	return err
}

func (p *S2CSetDefaultSpawnPosition) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteIdentifier(p.DimensionName); err != nil {
		return err
	}
	if err := buf.WritePosition(p.Location); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Yaw); err != nil {
		return err
	}
	return buf.WriteFloat32(p.Pitch)
}

// S2CSetDisplayObjective represents "Display Objective".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Display_Objective
type S2CSetDisplayObjective struct {
	Position  ns.VarInt
	ScoreName ns.String
}

func (p *S2CSetDisplayObjective) ID() ns.VarInt   { return S2CSetDisplayObjectiveID }
func (p *S2CSetDisplayObjective) State() jp.State { return jp.StatePlay }
func (p *S2CSetDisplayObjective) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetDisplayObjective) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Position, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.ScoreName, err = buf.ReadString(32767)
	return err
}

func (p *S2CSetDisplayObjective) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.Position); err != nil {
		return err
	}
	return buf.WriteString(p.ScoreName)
}

// S2CSetEntityData represents "Set Entity Metadata".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Entity_Metadata
type S2CSetEntityData struct {
	EntityId ns.VarInt
	Metadata ns.ByteArray
}

func (p *S2CSetEntityData) ID() ns.VarInt   { return S2CSetEntityDataID }
func (p *S2CSetEntityData) State() jp.State { return jp.StatePlay }
func (p *S2CSetEntityData) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetEntityData) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Metadata, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CSetEntityData) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Metadata)
}

// S2CSetEntityLink represents "Link Entities".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Link_Entities
type S2CSetEntityLink struct {
	AttachedEntityId ns.Int32
	HoldingEntityId  ns.Int32
}

func (p *S2CSetEntityLink) ID() ns.VarInt   { return S2CSetEntityLinkID }
func (p *S2CSetEntityLink) State() jp.State { return jp.StatePlay }
func (p *S2CSetEntityLink) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetEntityLink) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.AttachedEntityId, err = buf.ReadInt32(); err != nil {
		return err
	}
	p.HoldingEntityId, err = buf.ReadInt32()
	return err
}

func (p *S2CSetEntityLink) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt32(p.AttachedEntityId); err != nil {
		return err
	}
	return buf.WriteInt32(p.HoldingEntityId)
}

// S2CSetEntityMotion represents "Set Entity Velocity".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Entity_Velocity
type S2CSetEntityMotion struct {
	EntityId ns.VarInt
	Velocity ns.LpVec3
}

func (p *S2CSetEntityMotion) ID() ns.VarInt   { return S2CSetEntityMotionID }
func (p *S2CSetEntityMotion) State() jp.State { return jp.StatePlay }
func (p *S2CSetEntityMotion) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetEntityMotion) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Velocity, err = buf.ReadLpVec3()
	return err
}

func (p *S2CSetEntityMotion) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteLpVec3(p.Velocity)
}

// S2CSetEquipment represents "Set Equipment".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Equipment
type S2CSetEquipment struct {
	EntityId ns.VarInt
	Data     ns.ByteArray
}

func (p *S2CSetEquipment) ID() ns.VarInt   { return S2CSetEquipmentID }
func (p *S2CSetEquipment) State() jp.State { return jp.StatePlay }
func (p *S2CSetEquipment) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetEquipment) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CSetEquipment) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CSetExperience represents "Set Experience".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Experience
type S2CSetExperience struct {
	ExperienceBar   ns.Float32
	Level           ns.VarInt
	TotalExperience ns.VarInt
}

func (p *S2CSetExperience) ID() ns.VarInt   { return S2CSetExperienceID }
func (p *S2CSetExperience) State() jp.State { return jp.StatePlay }
func (p *S2CSetExperience) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetExperience) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.ExperienceBar, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Level, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.TotalExperience, err = buf.ReadVarInt()
	return err
}

func (p *S2CSetExperience) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat32(p.ExperienceBar); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Level); err != nil {
		return err
	}
	return buf.WriteVarInt(p.TotalExperience)
}

// S2CSetHealth represents "Set Health".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Health
type S2CSetHealth struct {
	Health         ns.Float32
	Food           ns.VarInt
	FoodSaturation ns.Float32
}

func (p *S2CSetHealth) ID() ns.VarInt   { return S2CSetHealthID }
func (p *S2CSetHealth) State() jp.State { return jp.StatePlay }
func (p *S2CSetHealth) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetHealth) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Health, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Food, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.FoodSaturation, err = buf.ReadFloat32()
	return err
}

func (p *S2CSetHealth) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat32(p.Health); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Food); err != nil {
		return err
	}
	return buf.WriteFloat32(p.FoodSaturation)
}

// S2CSetHeldSlot represents "Set Held Item (clientbound)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Held_Item_(Clientbound)
type S2CSetHeldSlot struct {
	Slot ns.VarInt
}

func (p *S2CSetHeldSlot) ID() ns.VarInt   { return S2CSetHeldSlotID }
func (p *S2CSetHeldSlot) State() jp.State { return jp.StatePlay }
func (p *S2CSetHeldSlot) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetHeldSlot) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Slot, err = buf.ReadVarInt()
	return err
}

func (p *S2CSetHeldSlot) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.Slot)
}

// S2CSetObjective represents "Update Objectives".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Objectives
type S2CSetObjective struct {
	ObjectiveName ns.String
	Mode          ns.Int8
	Data          ns.ByteArray
}

func (p *S2CSetObjective) ID() ns.VarInt   { return S2CSetObjectiveID }
func (p *S2CSetObjective) State() jp.State { return jp.StatePlay }
func (p *S2CSetObjective) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetObjective) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.ObjectiveName, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.Mode, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.Mode == 0 || p.Mode == 2 {
		p.Data, err = buf.ReadByteArray(1048576)
	}
	return err
}

func (p *S2CSetObjective) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteString(p.ObjectiveName); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.Mode); err != nil {
		return err
	}
	if p.Mode == 0 || p.Mode == 2 {
		return buf.WriteByteArray(p.Data)
	}
	return nil
}

// S2CSetPassengers represents "Set Passengers".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Passengers
type S2CSetPassengers struct {
	EntityId   ns.VarInt
	Passengers ns.ByteArray
}

func (p *S2CSetPassengers) ID() ns.VarInt   { return S2CSetPassengersID }
func (p *S2CSetPassengers) State() jp.State { return jp.StatePlay }
func (p *S2CSetPassengers) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetPassengers) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Passengers, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CSetPassengers) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Passengers)
}

// S2CSetPlayerInventory represents "Set Player Inventory Slot".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Player_Inventory_Slot
type S2CSetPlayerInventory struct {
	Slot     ns.VarInt
	SlotData ns.Slot
}

func (p *S2CSetPlayerInventory) ID() ns.VarInt   { return S2CSetPlayerInventoryID }
func (p *S2CSetPlayerInventory) State() jp.State { return jp.StatePlay }
func (p *S2CSetPlayerInventory) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetPlayerInventory) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Slot, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.SlotData, err = buf.ReadSlot()
	return err
}

func (p *S2CSetPlayerInventory) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.Slot); err != nil {
		return err
	}
	return buf.WriteSlot(p.SlotData)
}

// S2CSetPlayerTeam represents "Update Teams".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Teams
type S2CSetPlayerTeam struct {
	TeamName ns.String
	Method   ns.Int8
	Data     ns.ByteArray
}

func (p *S2CSetPlayerTeam) ID() ns.VarInt   { return S2CSetPlayerTeamID }
func (p *S2CSetPlayerTeam) State() jp.State { return jp.StatePlay }
func (p *S2CSetPlayerTeam) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetPlayerTeam) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.TeamName, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.Method, err = buf.ReadInt8(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CSetPlayerTeam) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteString(p.TeamName); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.Method); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CSetScore represents "Update Score".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Score
type S2CSetScore struct {
	EntityName    ns.String
	ObjectiveName ns.String
	Value         ns.VarInt
	Data          ns.ByteArray
}

func (p *S2CSetScore) ID() ns.VarInt   { return S2CSetScoreID }
func (p *S2CSetScore) State() jp.State { return jp.StatePlay }
func (p *S2CSetScore) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetScore) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityName, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.ObjectiveName, err = buf.ReadString(32767); err != nil {
		return err
	}
	if p.Value, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CSetScore) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteString(p.EntityName); err != nil {
		return err
	}
	if err := buf.WriteString(p.ObjectiveName); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Value); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CSetSimulationDistance represents "Set Simulation Distance".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Simulation_Distance
type S2CSetSimulationDistance struct {
	SimulationDistance ns.VarInt
}

func (p *S2CSetSimulationDistance) ID() ns.VarInt   { return S2CSetSimulationDistanceID }
func (p *S2CSetSimulationDistance) State() jp.State { return jp.StatePlay }
func (p *S2CSetSimulationDistance) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetSimulationDistance) Read(buf *ns.PacketBuffer) error {
	var err error
	p.SimulationDistance, err = buf.ReadVarInt()
	return err
}

func (p *S2CSetSimulationDistance) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.SimulationDistance)
}

// S2CSetSubtitleText represents "Set Subtitle Text".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Subtitle_Text
type S2CSetSubtitleText struct {
	SubtitleText ns.TextComponent
}

func (p *S2CSetSubtitleText) ID() ns.VarInt   { return S2CSetSubtitleTextID }
func (p *S2CSetSubtitleText) State() jp.State { return jp.StatePlay }
func (p *S2CSetSubtitleText) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetSubtitleText) Read(buf *ns.PacketBuffer) error {
	var err error
	p.SubtitleText, err = buf.ReadTextComponent()
	return err
}

func (p *S2CSetSubtitleText) Write(buf *ns.PacketBuffer) error {
	return buf.WriteTextComponent(p.SubtitleText)
}

// S2CSetTime represents "Update Time".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Time
type S2CSetTime struct {
	WorldAge            ns.Int64
	TimeOfDay           ns.Int64
	TimeOfDayIncreasing ns.Boolean
}

func (p *S2CSetTime) ID() ns.VarInt   { return S2CSetTimeID }
func (p *S2CSetTime) State() jp.State { return jp.StatePlay }
func (p *S2CSetTime) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetTime) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.WorldAge, err = buf.ReadInt64(); err != nil {
		return err
	}
	if p.TimeOfDay, err = buf.ReadInt64(); err != nil {
		return err
	}
	p.TimeOfDayIncreasing, err = buf.ReadBool()
	return err
}

func (p *S2CSetTime) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt64(p.WorldAge); err != nil {
		return err
	}
	if err := buf.WriteInt64(p.TimeOfDay); err != nil {
		return err
	}
	return buf.WriteBool(p.TimeOfDayIncreasing)
}

// S2CSetTitleText represents "Set Title Text".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Title_Text
type S2CSetTitleText struct {
	TitleText ns.TextComponent
}

func (p *S2CSetTitleText) ID() ns.VarInt   { return S2CSetTitleTextID }
func (p *S2CSetTitleText) State() jp.State { return jp.StatePlay }
func (p *S2CSetTitleText) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetTitleText) Read(buf *ns.PacketBuffer) error {
	var err error
	p.TitleText, err = buf.ReadTextComponent()
	return err
}

func (p *S2CSetTitleText) Write(buf *ns.PacketBuffer) error {
	return buf.WriteTextComponent(p.TitleText)
}

// S2CSetTitlesAnimation represents "Set Title Animation Times".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Title_Animation_Times
type S2CSetTitlesAnimation struct {
	FadeIn  ns.Int32
	Stay    ns.Int32
	FadeOut ns.Int32
}

func (p *S2CSetTitlesAnimation) ID() ns.VarInt   { return S2CSetTitlesAnimationID }
func (p *S2CSetTitlesAnimation) State() jp.State { return jp.StatePlay }
func (p *S2CSetTitlesAnimation) Bound() jp.Bound { return jp.S2C }

func (p *S2CSetTitlesAnimation) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.FadeIn, err = buf.ReadInt32(); err != nil {
		return err
	}
	if p.Stay, err = buf.ReadInt32(); err != nil {
		return err
	}
	p.FadeOut, err = buf.ReadInt32()
	return err
}

func (p *S2CSetTitlesAnimation) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt32(p.FadeIn); err != nil {
		return err
	}
	if err := buf.WriteInt32(p.Stay); err != nil {
		return err
	}
	return buf.WriteInt32(p.FadeOut)
}

// S2CSoundEntity represents "Entity Sound Effect".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Entity_Sound_Effect
type S2CSoundEntity struct {
	SoundEvent    ns.ByteArray
	SoundCategory ns.VarInt
	EntityId      ns.VarInt
	Volume        ns.Float32
	Pitch         ns.Float32
	Seed          ns.Int64
}

func (p *S2CSoundEntity) ID() ns.VarInt   { return S2CSoundEntityID }
func (p *S2CSoundEntity) State() jp.State { return jp.StatePlay }
func (p *S2CSoundEntity) Bound() jp.Bound { return jp.S2C }

func (p *S2CSoundEntity) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.SoundEvent, err = buf.ReadByteArray(1048576); err != nil {
		return err
	}
	if p.SoundCategory, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Volume, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.Seed, err = buf.ReadInt64()
	return err
}

func (p *S2CSoundEntity) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteByteArray(p.SoundEvent); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SoundCategory); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Volume); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Pitch); err != nil {
		return err
	}
	return buf.WriteInt64(p.Seed)
}

// S2CSound represents "Sound Effect".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Sound_Effect
type S2CSound struct {
	SoundEvent      ns.ByteArray
	SoundCategory   ns.VarInt
	EffectPositionX ns.Int32
	EffectPositionY ns.Int32
	EffectPositionZ ns.Int32
	Volume          ns.Float32
	Pitch           ns.Float32
	Seed            ns.Int64
}

func (p *S2CSound) ID() ns.VarInt   { return S2CSoundID }
func (p *S2CSound) State() jp.State { return jp.StatePlay }
func (p *S2CSound) Bound() jp.Bound { return jp.S2C }

func (p *S2CSound) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.SoundEvent, err = buf.ReadByteArray(1048576); err != nil {
		return err
	}
	if p.SoundCategory, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.EffectPositionX, err = buf.ReadInt32(); err != nil {
		return err
	}
	if p.EffectPositionY, err = buf.ReadInt32(); err != nil {
		return err
	}
	if p.EffectPositionZ, err = buf.ReadInt32(); err != nil {
		return err
	}
	if p.Volume, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.Seed, err = buf.ReadInt64()
	return err
}

func (p *S2CSound) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteByteArray(p.SoundEvent); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.SoundCategory); err != nil {
		return err
	}
	if err := buf.WriteInt32(p.EffectPositionX); err != nil {
		return err
	}
	if err := buf.WriteInt32(p.EffectPositionY); err != nil {
		return err
	}
	if err := buf.WriteInt32(p.EffectPositionZ); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Volume); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Pitch); err != nil {
		return err
	}
	return buf.WriteInt64(p.Seed)
}

// S2CStartConfiguration represents "Start Configuration".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Start_Configuration
type S2CStartConfiguration struct{}

func (p *S2CStartConfiguration) ID() ns.VarInt                { return S2CStartConfigurationID }
func (p *S2CStartConfiguration) State() jp.State              { return jp.StatePlay }
func (p *S2CStartConfiguration) Bound() jp.Bound              { return jp.S2C }
func (p *S2CStartConfiguration) Read(*ns.PacketBuffer) error  { return nil }
func (p *S2CStartConfiguration) Write(*ns.PacketBuffer) error { return nil }

// S2CStopSound represents "Stop Sound".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Stop_Sound
type S2CStopSound struct {
	Flags  ns.Int8
	Source ns.VarInt
	Sound  ns.Identifier
}

func (p *S2CStopSound) ID() ns.VarInt   { return S2CStopSoundID }
func (p *S2CStopSound) State() jp.State { return jp.StatePlay }
func (p *S2CStopSound) Bound() jp.Bound { return jp.S2C }

func (p *S2CStopSound) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Flags, err = buf.ReadInt8(); err != nil {
		return err
	}
	if p.Flags&0x01 != 0 {
		if p.Source, err = buf.ReadVarInt(); err != nil {
			return err
		}
	}
	if p.Flags&0x02 != 0 {
		p.Sound, err = buf.ReadIdentifier()
	}
	return err
}

func (p *S2CStopSound) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt8(p.Flags); err != nil {
		return err
	}
	if p.Flags&0x01 != 0 {
		if err := buf.WriteVarInt(p.Source); err != nil {
			return err
		}
	}
	if p.Flags&0x02 != 0 {
		return buf.WriteIdentifier(p.Sound)
	}
	return nil
}

// S2CStoreCookiePlay represents "Store Cookie (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Store_Cookie_(Play)
type S2CStoreCookiePlay struct {
	Key     ns.Identifier
	Payload ns.ByteArray
}

func (p *S2CStoreCookiePlay) ID() ns.VarInt   { return S2CStoreCookiePlayID }
func (p *S2CStoreCookiePlay) State() jp.State { return jp.StatePlay }
func (p *S2CStoreCookiePlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CStoreCookiePlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Key, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	p.Payload, err = buf.ReadByteArray(5120)
	return err
}

func (p *S2CStoreCookiePlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteIdentifier(p.Key); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Payload)
}

// S2CSystemChat represents "System Chat Message".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#System_Chat_Message
type S2CSystemChat struct {
	Content ns.TextComponent
	Overlay ns.Boolean
}

func (p *S2CSystemChat) ID() ns.VarInt   { return S2CSystemChatID }
func (p *S2CSystemChat) State() jp.State { return jp.StatePlay }
func (p *S2CSystemChat) Bound() jp.Bound { return jp.S2C }

func (p *S2CSystemChat) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Content, err = buf.ReadTextComponent(); err != nil {
		return err
	}
	p.Overlay, err = buf.ReadBool()
	return err
}

func (p *S2CSystemChat) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteTextComponent(p.Content); err != nil {
		return err
	}
	return buf.WriteBool(p.Overlay)
}

// S2CTabList represents "Set Tab List Header And Footer".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Tab_List_Header_And_Footer
type S2CTabList struct {
	Header ns.TextComponent
	Footer ns.TextComponent
}

func (p *S2CTabList) ID() ns.VarInt   { return S2CTabListID }
func (p *S2CTabList) State() jp.State { return jp.StatePlay }
func (p *S2CTabList) Bound() jp.Bound { return jp.S2C }

func (p *S2CTabList) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Header, err = buf.ReadTextComponent(); err != nil {
		return err
	}
	p.Footer, err = buf.ReadTextComponent()
	return err
}

func (p *S2CTabList) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteTextComponent(p.Header); err != nil {
		return err
	}
	return buf.WriteTextComponent(p.Footer)
}

// S2CTagQuery represents "Tag Query Response".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Tag_Query_Response
type S2CTagQuery struct {
	TransactionId ns.VarInt
	Nbt           nbt.Tag
}

func (p *S2CTagQuery) ID() ns.VarInt   { return S2CTagQueryID }
func (p *S2CTagQuery) State() jp.State { return jp.StatePlay }
func (p *S2CTagQuery) Bound() jp.Bound { return jp.S2C }

func (p *S2CTagQuery) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.TransactionId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	remaining, err := buf.ReadByteArray(1048576)
	if err != nil {
		return err
	}
	p.Nbt, err = nbt.DecodeNetwork(remaining)
	return err
}

func (p *S2CTagQuery) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.TransactionId); err != nil {
		return err
	}
	data, err := nbt.EncodeNetwork(p.Nbt)
	if err != nil {
		return err
	}
	return buf.WriteFixedByteArray(data)
}

// S2CTakeItemEntity represents "Pickup Item".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Pickup_Item
type S2CTakeItemEntity struct {
	CollectedEntityId ns.VarInt
	CollectorEntityId ns.VarInt
	PickupItemCount   ns.VarInt
}

func (p *S2CTakeItemEntity) ID() ns.VarInt   { return S2CTakeItemEntityID }
func (p *S2CTakeItemEntity) State() jp.State { return jp.StatePlay }
func (p *S2CTakeItemEntity) Bound() jp.Bound { return jp.S2C }

func (p *S2CTakeItemEntity) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.CollectedEntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.CollectorEntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.PickupItemCount, err = buf.ReadVarInt()
	return err
}

func (p *S2CTakeItemEntity) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.CollectedEntityId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.CollectorEntityId); err != nil {
		return err
	}
	return buf.WriteVarInt(p.PickupItemCount)
}

// S2CTeleportEntity represents "Synchronize Vehicle Position".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Synchronize_Vehicle_Position
type S2CTeleportEntity struct {
	EntityId  ns.VarInt
	X         ns.Float64
	Y         ns.Float64
	Z         ns.Float64
	VelocityX ns.Float64
	VelocityY ns.Float64
	VelocityZ ns.Float64
	Yaw       ns.Float32
	Pitch     ns.Float32
	Flags     ns.Int8
	OnGround  ns.Boolean
}

func (p *S2CTeleportEntity) ID() ns.VarInt   { return S2CTeleportEntityID }
func (p *S2CTeleportEntity) State() jp.State { return jp.StatePlay }
func (p *S2CTeleportEntity) Bound() jp.Bound { return jp.S2C }

func (p *S2CTeleportEntity) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.X, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Y, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Z, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.VelocityX, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.VelocityY, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.VelocityZ, err = buf.ReadFloat64(); err != nil {
		return err
	}
	if p.Yaw, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Pitch, err = buf.ReadFloat32(); err != nil {
		return err
	}
	if p.Flags, err = buf.ReadInt8(); err != nil {
		return err
	}
	p.OnGround, err = buf.ReadBool()
	return err
}

func (p *S2CTeleportEntity) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.X); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Y); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.Z); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.VelocityX); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.VelocityY); err != nil {
		return err
	}
	if err := buf.WriteFloat64(p.VelocityZ); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Yaw); err != nil {
		return err
	}
	if err := buf.WriteFloat32(p.Pitch); err != nil {
		return err
	}
	if err := buf.WriteInt8(p.Flags); err != nil {
		return err
	}
	return buf.WriteBool(p.OnGround)
}

// S2CTestInstanceBlockStatus represents "Test Instance Block Status".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Test_Instance_Block_Status
type S2CTestInstanceBlockStatus struct {
	Status ns.TextComponent
	Size   ns.PrefixedOptional[ns.ByteArray]
}

func (p *S2CTestInstanceBlockStatus) ID() ns.VarInt   { return S2CTestInstanceBlockStatusID }
func (p *S2CTestInstanceBlockStatus) State() jp.State { return jp.StatePlay }
func (p *S2CTestInstanceBlockStatus) Bound() jp.Bound { return jp.S2C }

func (p *S2CTestInstanceBlockStatus) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Status, err = buf.ReadTextComponent(); err != nil {
		return err
	}
	return p.Size.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.ByteArray, error) {
		return b.ReadByteArray(24)
	})
}

func (p *S2CTestInstanceBlockStatus) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteTextComponent(p.Status); err != nil {
		return err
	}
	return p.Size.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.ByteArray) error {
		return b.WriteByteArray(v)
	})
}

// S2CTickingState represents "Set Ticking State".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Ticking_State
type S2CTickingState struct {
	TickRate ns.Float32
	IsFrozen ns.Boolean
}

func (p *S2CTickingState) ID() ns.VarInt   { return S2CTickingStateID }
func (p *S2CTickingState) State() jp.State { return jp.StatePlay }
func (p *S2CTickingState) Bound() jp.Bound { return jp.S2C }

func (p *S2CTickingState) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.TickRate, err = buf.ReadFloat32(); err != nil {
		return err
	}
	p.IsFrozen, err = buf.ReadBool()
	return err
}

func (p *S2CTickingState) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteFloat32(p.TickRate); err != nil {
		return err
	}
	return buf.WriteBool(p.IsFrozen)
}

// S2CTickingStep represents "Step Tick".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Step_Tick
type S2CTickingStep struct {
	TickSteps ns.VarInt
}

func (p *S2CTickingStep) ID() ns.VarInt   { return S2CTickingStepID }
func (p *S2CTickingStep) State() jp.State { return jp.StatePlay }
func (p *S2CTickingStep) Bound() jp.Bound { return jp.S2C }

func (p *S2CTickingStep) Read(buf *ns.PacketBuffer) error {
	var err error
	p.TickSteps, err = buf.ReadVarInt()
	return err
}

func (p *S2CTickingStep) Write(buf *ns.PacketBuffer) error {
	return buf.WriteVarInt(p.TickSteps)
}

// S2CTransferPlay represents "Transfer (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Transfer_(Play)
type S2CTransferPlay struct {
	Host ns.String
	Port ns.VarInt
}

func (p *S2CTransferPlay) ID() ns.VarInt   { return S2CTransferPlayID }
func (p *S2CTransferPlay) State() jp.State { return jp.StatePlay }
func (p *S2CTransferPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CTransferPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.Host, err = buf.ReadString(32767); err != nil {
		return err
	}
	p.Port, err = buf.ReadVarInt()
	return err
}

func (p *S2CTransferPlay) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteString(p.Host); err != nil {
		return err
	}
	return buf.WriteVarInt(p.Port)
}

// S2CUpdateAdvancements represents "Update Advancements".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Advancements
type S2CUpdateAdvancements struct {
	Data ns.ByteArray
}

func (p *S2CUpdateAdvancements) ID() ns.VarInt   { return S2CUpdateAdvancementsID }
func (p *S2CUpdateAdvancements) State() jp.State { return jp.StatePlay }
func (p *S2CUpdateAdvancements) Bound() jp.Bound { return jp.S2C }

func (p *S2CUpdateAdvancements) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CUpdateAdvancements) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Data)
}

// S2CUpdateAttributes represents "Update Attributes".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Attributes
type S2CUpdateAttributes struct {
	EntityId ns.VarInt
	Data     ns.ByteArray
}

func (p *S2CUpdateAttributes) ID() ns.VarInt   { return S2CUpdateAttributesID }
func (p *S2CUpdateAttributes) State() jp.State { return jp.StatePlay }
func (p *S2CUpdateAttributes) Bound() jp.Bound { return jp.S2C }

func (p *S2CUpdateAttributes) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CUpdateAttributes) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteByteArray(p.Data)
}

// S2CUpdateMobEffect represents "Entity Effect".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Entity_Effect
type S2CUpdateMobEffect struct {
	EntityId  ns.VarInt
	EffectId  ns.VarInt
	Amplifier ns.VarInt
	Duration  ns.VarInt
	Flags     ns.Int8
}

func (p *S2CUpdateMobEffect) ID() ns.VarInt   { return S2CUpdateMobEffectID }
func (p *S2CUpdateMobEffect) State() jp.State { return jp.StatePlay }
func (p *S2CUpdateMobEffect) Bound() jp.Bound { return jp.S2C }

func (p *S2CUpdateMobEffect) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.EffectId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Amplifier, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if p.Duration, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Flags, err = buf.ReadInt8()
	return err
}

func (p *S2CUpdateMobEffect) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.EffectId); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Amplifier); err != nil {
		return err
	}
	if err := buf.WriteVarInt(p.Duration); err != nil {
		return err
	}
	return buf.WriteInt8(p.Flags)
}

// S2CUpdateRecipes represents "Update Recipes".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Recipes
type S2CUpdateRecipes struct {
	Data ns.ByteArray
}

func (p *S2CUpdateRecipes) ID() ns.VarInt   { return S2CUpdateRecipesID }
func (p *S2CUpdateRecipes) State() jp.State { return jp.StatePlay }
func (p *S2CUpdateRecipes) Bound() jp.Bound { return jp.S2C }

func (p *S2CUpdateRecipes) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CUpdateRecipes) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Data)
}

// S2CUpdateTagsPlay represents "Update Tags (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Tags_(Play)
type S2CUpdateTagsPlay struct {
	Data ns.ByteArray
}

func (p *S2CUpdateTagsPlay) ID() ns.VarInt   { return S2CUpdateTagsPlayID }
func (p *S2CUpdateTagsPlay) State() jp.State { return jp.StatePlay }
func (p *S2CUpdateTagsPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CUpdateTagsPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CUpdateTagsPlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Data)
}

// S2CProjectilePower represents "Projectile Power".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Projectile_Power
type S2CProjectilePower struct {
	EntityId ns.VarInt
	Power    ns.Float64
}

func (p *S2CProjectilePower) ID() ns.VarInt   { return S2CProjectilePowerID }
func (p *S2CProjectilePower) State() jp.State { return jp.StatePlay }
func (p *S2CProjectilePower) Bound() jp.Bound { return jp.S2C }

func (p *S2CProjectilePower) Read(buf *ns.PacketBuffer) error {
	var err error
	if p.EntityId, err = buf.ReadVarInt(); err != nil {
		return err
	}
	p.Power, err = buf.ReadFloat64()
	return err
}

func (p *S2CProjectilePower) Write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(p.EntityId); err != nil {
		return err
	}
	return buf.WriteFloat64(p.Power)
}

// S2CCustomReportDetails represents "Custom Report Details".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Custom_Report_Details
type S2CCustomReportDetails struct {
	Details ns.ByteArray
}

func (p *S2CCustomReportDetails) ID() ns.VarInt   { return S2CCustomReportDetailsID }
func (p *S2CCustomReportDetails) State() jp.State { return jp.StatePlay }
func (p *S2CCustomReportDetails) Bound() jp.Bound { return jp.S2C }

func (p *S2CCustomReportDetails) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Details, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CCustomReportDetails) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Details)
}

// S2CServerLinks represents "Server Links".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Server_Links
type S2CServerLinks struct {
	Links ns.ByteArray
}

func (p *S2CServerLinks) ID() ns.VarInt   { return S2CServerLinksID }
func (p *S2CServerLinks) State() jp.State { return jp.StatePlay }
func (p *S2CServerLinks) Bound() jp.Bound { return jp.S2C }

func (p *S2CServerLinks) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Links, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CServerLinks) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Links)
}

// S2CWaypoint represents "Waypoint".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Waypoint
type S2CWaypoint struct {
	Data ns.ByteArray
}

func (p *S2CWaypoint) ID() ns.VarInt   { return S2CWaypointID }
func (p *S2CWaypoint) State() jp.State { return jp.StatePlay }
func (p *S2CWaypoint) Bound() jp.Bound { return jp.S2C }

func (p *S2CWaypoint) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Data, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CWaypoint) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Data)
}

// S2CClearDialogPlay represents "Clear Dialog (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clear_Dialog_(Play)
type S2CClearDialogPlay struct{}

func (p *S2CClearDialogPlay) ID() ns.VarInt                { return S2CClearDialogPlayID }
func (p *S2CClearDialogPlay) State() jp.State              { return jp.StatePlay }
func (p *S2CClearDialogPlay) Bound() jp.Bound              { return jp.S2C }
func (p *S2CClearDialogPlay) Read(*ns.PacketBuffer) error  { return nil }
func (p *S2CClearDialogPlay) Write(*ns.PacketBuffer) error { return nil }

// S2CShowDialogPlay represents "Show Dialog (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Show_Dialog_(Play)
type S2CShowDialogPlay struct {
	Dialog ns.ByteArray
}

func (p *S2CShowDialogPlay) ID() ns.VarInt   { return S2CShowDialogPlayID }
func (p *S2CShowDialogPlay) State() jp.State { return jp.StatePlay }
func (p *S2CShowDialogPlay) Bound() jp.Bound { return jp.S2C }

func (p *S2CShowDialogPlay) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Dialog, err = buf.ReadByteArray(1048576)
	return err
}

func (p *S2CShowDialogPlay) Write(buf *ns.PacketBuffer) error {
	return buf.WriteByteArray(p.Dialog)
}

package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/net_structures"
)

// S2CBundleDelimiter represents "Bundle Delimiter".
//
// > The delimiter for a bundle of packets. When received, the client should store every subsequent packet it receives, and wait until another delimiter is received. Once that happens, the client is guaranteed to
// > process every packet in the bundle on the same tick, and the client should stop storing packets.
// >
// > As of 1.20.6, the vanilla server only uses this to ensure Spawn Entity and associated packets used to configure the entity happen on the same tick. Each entity gets a separate bundle.
// >
// > The vanilla client doesn't allow more than 4096 packets in the same bundle.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Bundle_Delimiter
var S2CBundleDelimiter = jp.NewPacket(jp.StatePlay, jp.S2C, 0x00)

type S2CBundleDelimiterData struct {
	// No fields
}

// S2CAddEntity represents "Spawn Entity".
//
// > Sent by the server when an entity (aside from Experience Orb ) is created.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Spawn_Entity
var S2CAddEntity = jp.NewPacket(jp.StatePlay, jp.S2C, 0x01)

type S2CAddEntityData struct {
	// A unique integer ID mostly used in the protocol to identify the entity.
	EntityId ns.VarInt
	// A unique identifier that is mostly used in persistence and places where the uniqueness matters more.
	EntityUuid ns.UUID
	// ID in the minecraft:entity_type registry (see "type" field in Entity metadata#Entities ).
	Type ns.VarInt
	//
	X ns.Double
	//
	Y ns.Double
	//
	Z ns.Double
	// To get the real pitch, you must divide this by (256.0F / 360.0F)
	Pitch ns.Angle
	// To get the real yaw, you must divide this by (256.0F / 360.0F)
	Yaw ns.Angle
	// Only used by living entities, where the head of the entity may differ from the general body rotation.
	HeadYaw ns.Angle
	// Meaning dependent on the value of the Type field, see Object Data for details.
	Data ns.VarInt
	// Same units as Set Entity Velocity .
	VelocityX ns.Short
}

// S2CAnimate represents "Entity Animation".
//
// > Sent whenever an entity should change animation.
// >
// > Animation can be one of the following values:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Entity_Animation
var S2CAnimate = jp.NewPacket(jp.StatePlay, jp.S2C, 0x02)

type S2CAnimateData struct {
	// Player ID.
	EntityId ns.VarInt
	// Animation ID (see below).
	Animation ns.UnsignedByte
}

// S2CAwardStats represents "Award Statistics".
//
// > Sent as a response to Client Status (id 1). Will only send the changed values if previously requested.
// >
// > Categories (these are namespaced, but with : replaced with . ):
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Award_Statistics
var S2CAwardStats = jp.NewPacket(jp.StatePlay, jp.S2C, 0x03)

type S2CAwardStatsData struct {
	// Prefixed Array
	Statistics ns.PrefixedArray[struct {
		CategoryId  ns.VarInt
		StatisticId ns.VarInt
		Value       ns.VarInt
	}]
}

// S2CBlockChangedAck represents "Acknowledge Block Change".
//
// > Acknowledges a user-initiated block change. After receiving this packet, the client will display the block state sent by the server instead of the one predicted by the client.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Acknowledge_Block_Change
var S2CBlockChangedAck = jp.NewPacket(jp.StatePlay, jp.S2C, 0x04)

type S2CBlockChangedAckData struct {
	// Represents the sequence to acknowledge, this is used for properly syncing block changes to the client after interactions.
	SequenceId ns.VarInt
}

// S2CBlockDestruction represents "Set Block Destroy Stage".
//
// > 0–9 are the displayable destroy stages and each other number means that there is no animation on this coordinate.
// >
// > Block break animations can still be applied on air; the animation will remain visible although there is no block being broken. However, if this is applied to a transparent block, odd graphical effects may happen,
// > including water losing its transparency. (An effect similar to this can be seen in normal gameplay when breaking ice blocks)
// >
// > If you need to display several break animations at the same time you have to give each of them a unique Entity ID. The entity ID does not need to correspond to an actual entity on the client. It is valid to use a
// > randomly generated number.
// >
// > When removing break animation, you must use the ID of the entity that set it.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Block_Destroy_Stage
var S2CBlockDestruction = jp.NewPacket(jp.StatePlay, jp.S2C, 0x05)

type S2CBlockDestructionData struct {
	// The ID of the entity breaking the block.
	EntityId ns.VarInt
	// Block Position.
	Location ns.Position
	// 0–9 to set it, any other value to remove it.
	DestroyStage ns.UnsignedByte
}

// S2CBlockEntityData represents "Block Entity Data".
//
// > Sets the block entity associated with the block at the given location.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Block_Entity_Data
var S2CBlockEntityData = jp.NewPacket(jp.StatePlay, jp.S2C, 0x06)

type S2CBlockEntityDataData struct {
	//
	Location ns.Position
	// ID in the minecraft:block_entity_type registry
	Type ns.VarInt
	// Data to set.
	NbtData ns.NBT
}

// S2CBlockEvent represents "Block Action".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Block_Action
var S2CBlockEvent = jp.NewPacket(jp.StatePlay, jp.S2C, 0x07)

type S2CBlockEventData struct {
	// Block coordinates.
	Location ns.Position
	// Varies depending on block — see Block Actions .
	ActionId ns.UnsignedByte
	// Varies depending on block — see Block Actions .
	ActionParameter ns.UnsignedByte
	// ID in the minecraft:block registry. This value is unused by the vanilla client, as it will infer the type of block based on the given position.
	BlockType ns.VarInt
}

// S2CBlockUpdate represents "Block Update".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Block_Update
var S2CBlockUpdate = jp.NewPacket(jp.StatePlay, jp.S2C, 0x08)

type S2CBlockUpdateData struct {
	// Block Coordinates.
	Location ns.Position
	// The new block state ID for the block as given in the block state registry .
	BlockId ns.VarInt
}

// S2CBossEvent represents "Boss Bar".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Boss_Bar
var S2CBossEvent = jp.NewPacket(jp.StatePlay, jp.S2C, 0x09)

type S2CBossEventData struct {
	// Unique ID for this bar.
	Uuid ns.UUID
	// Determines the layout of the remaining packet.
	Action ns.VarInt
	// From 0 to 1. Values greater than 1 do not crash a vanilla client, and start rendering part of a second health bar at around 1.5.
	Health ns.Float
	// Color ID (see below).
	Color ns.VarInt
	// Type of division (see below).
	Division ns.VarInt
	// Bit mask. 0x01: should darken sky, 0x02: is dragon bar (used to play end music), 0x04: create fog (previously was also controlled by 0x02).
	Flags ns.UnsignedByte
	// as above
	Dividers ns.VarInt
}

// S2CChangeDifficulty represents "Change Difficulty".
//
// > Changes the difficulty setting in the client's option menu
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Change_Difficulty
var S2CChangeDifficulty = jp.NewPacket(jp.StatePlay, jp.S2C, 0x0A)

type S2CChangeDifficultyData struct {
	// 0: peaceful, 1: easy, 2: normal, 3: hard.
	Difficulty ns.UnsignedByte
	//
	DifficultyLocked ns.Boolean
}

// S2CChunkBatchFinished represents "Chunk Batch Finished".
//
// > Marks the end of a chunk batch. The vanilla client marks the time it receives this packet and calculates the elapsed duration since the beginning of the chunk batch . The server
// > uses this duration and the batch size received in this packet to estimate the number of milliseconds elapsed per chunk received. This value is then used to calculate the desired number of chunks per tick through
// > the formula 25 / millisPerChunk , which is reported to the server through Chunk Batch Received . This likely uses 25 instead of the normal tick duration
// > of 50 so chunk processing will only use half of the client's and network's bandwidth.
// >
// > The vanilla client uses the samples from the latest 15 batches to estimate the milliseconds per chunk number.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chunk_Batch_Finished
var S2CChunkBatchFinished = jp.NewPacket(jp.StatePlay, jp.S2C, 0x0B)

type S2CChunkBatchFinishedData struct {
	// Number of chunks.
	BatchSize ns.VarInt
}

// S2CChunkBatchStart represents "Chunk Batch Start".
//
// > Marks the start of a chunk batch. The vanilla client marks and stores the time it receives this packet.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chunk_Batch_Start
var S2CChunkBatchStart = jp.NewPacket(jp.StatePlay, jp.S2C, 0x0C)

type S2CChunkBatchStartData struct {
	// No fields
}

// S2CChunksBiomes represents "Chunk Biomes".
//
// > Note: The order of X and Z is inverted, because the client reads them as one big-endian Long , with Z being the upper 32 bits.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chunk_Biomes
var S2CChunksBiomes = jp.NewPacket(jp.StatePlay, jp.S2C, 0x0D)

type S2CChunksBiomesData struct {
	// Prefixed Array
	ChunkBiomeData ns.PrefixedArray[struct {
		ChunkZ ns.Int
		ChunkX ns.Int
		Data   ns.PrefixedByteArray
	}]
}

// S2CClearTitles represents "Clear Titles".
//
// > Clear the client's current title information, with the option to also reset it.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clear_Titles
var S2CClearTitles = jp.NewPacket(jp.StatePlay, jp.S2C, 0x0E)

type S2CClearTitlesData struct {
	//
	Reset ns.Boolean
}

// S2CCommandSuggestions represents "Command Suggestions Response".
//
// > The server responds with a list of auto-completions of the last word sent to it. In the case of regular chat, this is a player username. Command names and parameters are also supported. The client sorts these
// > alphabetically before listing them.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Command_Suggestions_Response
var S2CCommandSuggestions = jp.NewPacket(jp.StatePlay, jp.S2C, 0x0F)

type S2CCommandSuggestionsData struct {
	// Transaction ID.
	Id ns.VarInt
	// Start of the text to replace.
	Start ns.VarInt
	// Length of the text to replace.
	Length ns.VarInt
	// Tooltip to display.
	Tooltip ns.PrefixedOptional[ns.TextComponent]
}

// S2CCommands represents "Commands".
//
// > Lists all of the commands on the server, and how they are parsed.
// >
// > This is a directed graph, with one root node. Each redirect or child node must refer only to nodes that have already been declared.
// >
// > For more information on this packet, see the Command Data article.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Commands
var S2CCommands = jp.NewPacket(jp.StatePlay, jp.S2C, 0x10)

type S2CCommandsData struct {
	// An array of nodes.
	Nodes ns.PrefixedArray[ns.ByteArray] // TODO: Node
	// Index of the root node in the previous array.
	RootIndex ns.VarInt
}

// S2CContainerClose represents "Close Container".
//
// > This packet is sent from the server to the client when a window is forcibly closed, such as when a chest is destroyed while it's open. The vanilla client disregards the provided window ID and closes any active
// > window.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Close_Container
var S2CContainerClose = jp.NewPacket(jp.StatePlay, jp.S2C, 0x11)

type S2CContainerCloseData struct {
	// This is the ID of the window that was closed. 0 for inventory.
	WindowId ns.VarInt
}

// S2CContainerSetContent represents "Set Container Content".
//
// > Replaces the contents of a container window. Sent by the server upon initialization of a container window or the player's inventory, and in response to state ID mismatches (see #Click Container ).
// >
// > See inventory windows for further information about how slots
// > are indexed. Use Open Screen to open the container on the client.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Container_Content
var S2CContainerSetContent = jp.NewPacket(jp.StatePlay, jp.S2C, 0x12)

type S2CContainerSetContentData struct {
	// The ID of window which items are being sent for. 0 for player inventory. The client ignores any packets targeting a Window ID other than the current one. However, an exception is made for the player inventory, which may be targeted at any time. (The vanilla server does not appear to utilize this special case.)
	WindowId ns.VarInt
	// A server-managed sequence number used to avoid desynchronization; see #Click Container .
	StateId ns.VarInt
	// Item being dragged with the mouse.
	CarriedItem ns.Slot
}

// S2CContainerSetData represents "Set Container Property".
//
// > This packet is used to inform the client that part of a GUI window should be updated.
// >
// > The meaning of the Property field depends on the type of the window. The following table shows the known combinations of window type and property, and how the value is to be interpreted.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Container_Property
var S2CContainerSetData = jp.NewPacket(jp.StatePlay, jp.S2C, 0x13)

type S2CContainerSetDataData struct {
	//
	WindowId ns.VarInt
	// The property to be updated, see below.
	Property ns.Short
	// The new value for the property, see below.
	Value ns.Short
}

// S2CContainerSetSlot represents "Set Container Slot".
//
// > Sent by the server when an item in a slot (in a window) is added/removed.
// >
// > If Window ID is 0, the hotbar and offhand slots (slots 36 through 45) may be updated even when a different container window is open. (The vanilla server does not appear to utilize this special case.) Updates are
// > also restricted to those slots when the player is looking at a creative inventory tab other than the survival inventory. (The vanilla server does not handle this restriction in any way, leading to MC-242392 .)
// >
// > If Window ID is -1, the item being dragged with the mouse is set. In this case, State ID and Slot are ignored.
// >
// > If Window ID is -2, any slot in the player's inventory can be updated irrespective of the current container window. In this case, State ID is ignored, and the vanilla server uses a bogus value of 0. Used by the
// > vanilla server to implement the #Pick Item functionality.
// >
// > When a container window is open, the server never sends updates targeting Window ID 0—all of the window types include slots for the player inventory. The client must
// > automatically apply changes targeting the inventory portion of a container window to the main inventory; the server does not resend them for ID 0 when the window is closed. However, since the armor and offhand
// > slots are only present on ID 0, updates to those slots occurring while a window is open must be deferred by the server until the window's closure.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Container_Slot
var S2CContainerSetSlot = jp.NewPacket(jp.StatePlay, jp.S2C, 0x14)

type S2CContainerSetSlotData struct {
	// The window which is being updated. 0 for player inventory. The client ignores any packets targeting a Window ID other than the current one; see below for exceptions.
	WindowId ns.VarInt
	// A server-managed sequence number used to avoid desynchronization; see #Click Container .
	StateId ns.VarInt
	// The slot that should be updated.
	Slot ns.Short
	//
	SlotData ns.Slot
}

// S2CCookieRequestPlay represents "Cookie Request (play)".
//
// > Requests a cookie that was previously stored.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Cookie_Request_(Play)
var S2CCookieRequestPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x15)

type S2CCookieRequestPlayData struct {
	// The identifier of the cookie.
	Key ns.Identifier
}

// S2CCooldown represents "Set Cooldown".
//
// > Applies a cooldown period to all items with the given type. Used by the vanilla server with enderpearls. This packet should be sent when the cooldown starts and also when the cooldown ends (to compensate for
// > lag), although the client will end the cooldown automatically. Can be applied to any item, note that interactions still get sent to the server with the item but the client does not play the animation nor attempt
// > to predict results (i.e block placing).
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Cooldown
var S2CCooldown = jp.NewPacket(jp.StatePlay, jp.S2C, 0x16)

type S2CCooldownData struct {
	// Identifier of the item (minecraft:stone) or the cooldown group ("use_cooldown" item component)
	CooldownGroup ns.Identifier
	// Number of ticks to apply a cooldown for, or 0 to clear the cooldown.
	CooldownTicks ns.VarInt
}

// S2CCustomChatCompletions represents "Chat Suggestions".
//
// > Unused by the vanilla server. Likely provided for custom servers to send chat message completions to clients.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chat_Suggestions
var S2CCustomChatCompletions = jp.NewPacket(jp.StatePlay, jp.S2C, 0x17)

type S2CCustomChatCompletionsData struct {
	// 0: Add, 1: Remove, 2: Set
	Action ns.VarInt
	//
	Entries ns.PrefixedArray[ns.String]
}

// S2CCustomPayloadPlay represents "Clientbound Plugin Message (play)".
//
// > Mods and plugins can use this to send their data. Minecraft itself uses several plugin channels . These internal channels are in the minecraft namespace.
// >
// > More information on how it works on Dinnerbone's blog . More documentation about
// > internal and popular registered channels are here .
// >
// > In vanilla clients, the maximum data length is 1048576 bytes.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clientbound_Plugin_Message_(Play)
var S2CCustomPayloadPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x18)

type S2CCustomPayloadPlayData struct {
	// Name of the plugin channel used to send the data.
	Channel ns.Identifier
	// Any data. The length of this array must be inferred from the packet length.
	Data ns.ByteArray
}

// S2CDamageEvent represents "Damage Event".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Damage_Event
var S2CDamageEvent = jp.NewPacket(jp.StatePlay, jp.S2C, 0x19)

type S2CDamageEventData struct {
	// The ID of the entity taking damage
	EntityId ns.VarInt
	// The type of damage in the minecraft:damage_type registry, defined by the Registry Data packet.
	SourceTypeId ns.VarInt
	// The ID + 1 of the entity responsible for the damage, if present. If not present, the value is 0
	SourceCauseId ns.VarInt
	// The ID + 1 of the entity that directly dealt the damage, if present. If not present, the value is 0. If this field is present: and damage was dealt indirectly, such as by the use of a projectile, this field will contain the ID of such projectile; and damage was dealt dirctly, such as by manually attacking, this field will contain the same value as Source Cause ID.
	SourceDirectId ns.VarInt
}

// S2CDebugSample represents "Debug Sample".
//
// > Sample data that is sent periodically after the client has subscribed with Debug Sample Subscription .
// >
// > The vanilla server only sends debug samples to players that are server operators.
// >
// > Types:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Debug_Sample
var S2CDebugSample = jp.NewPacket(jp.StatePlay, jp.S2C, 0x1A)

type S2CDebugSampleData struct {
	// Array of type-dependent samples.
	Sample ns.PrefixedArray[ns.Long]
	// See below.
	SampleType ns.VarInt
}

// S2CDeleteChat represents "Delete Message".
//
// > Removes a message from the client's chat. This only works for messages with signatures, system messages cannot be deleted with this packet.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Delete_Message
var S2CDeleteChat = jp.NewPacket(jp.StatePlay, jp.S2C, 0x1B)

type S2CDeleteChatData struct {
	// The message Id + 1, used for validating message signature. The next field is present only when value of this field is equal to 0.
	MessageId ns.VarInt
	// The previous message's signature. Always 256 bytes and not length-prefixed.
	Signature ns.Optional[ns.ByteArray]
}

// S2CDisconnectPlay represents "Disconnect (play)".
//
// > Sent by the server before it disconnects a client. The client assumes that the server has already closed the connection by the time the packet arrives.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Disconnect_(Play)
var S2CDisconnectPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x1C)

type S2CDisconnectPlayData struct {
	// Displayed to the client when the connection terminates.
	Reason ns.TextComponent
}

// S2CDisguisedChat represents "Disguised Chat Message".
//
// > Sends the client a chat message, but without any message signing information.
// >
// > The vanilla server uses this packet when the console is communicating with players through commands, such as /say , /tell , /me , among others.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Disguised_Chat_Message
var S2CDisguisedChat = jp.NewPacket(jp.StatePlay, jp.S2C, 0x1D)

type S2CDisguisedChatData struct {
	// This is used as the content parameter when formatting the message on the client.
	Message ns.TextComponent
	// Either the type of chat in the minecraft:chat_type registry, defined by the Registry Data packet, or an inline definition.
	ChatType ns.Or[ns.Identifier, ns.ChatType]
	// The name of the one sending the message, usually the sender's display name. This is used as the sender parameter when formatting the message on the client.
	SenderName ns.TextComponent
	// The name of the one receiving the message, usually the receiver's display name. This is used as the target parameter when formatting the message on the client.
	TargetName ns.PrefixedOptional[ns.TextComponent]
}

// S2CEntityEvent represents "Entity Event".
//
// > Entity statuses generally trigger an animation for an entity. The available statuses vary by the entity's type (and are available to subclasses of that type as well).
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Entity_Event
var S2CEntityEvent = jp.NewPacket(jp.StatePlay, jp.S2C, 0x1E)

type S2CEntityEventData struct {
	//
	EntityId ns.Int
	// See Entity statuses for a list of which statuses are valid for each type of entity.
	EntityStatus ns.Byte
}

// S2CEntityPositionSync represents "Teleport Entity".
//
// > This packet is sent by the server when an entity moves more than 8 blocks.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Teleport_Entity
var S2CEntityPositionSync = jp.NewPacket(jp.StatePlay, jp.S2C, 0x1F)

type S2CEntityPositionSyncData struct {
	//
	EntityId ns.VarInt
	//
	X ns.Double
	//
	Y ns.Double
	//
	Z ns.Double
	//
	VelocityX ns.Double
	//
	VelocityY ns.Double
	//
	VelocityZ ns.Double
	// Rotation on the X axis, in degrees.
	Yaw ns.Float
	// Rotation on the Y axis, in degrees.
	Pitch ns.Float
	//
	OnGround ns.Boolean
}

// S2CExplode represents "Explosion".
//
// > Sent when an explosion occurs (creepers, TNT, and ghast fireballs).
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Explosion
var S2CExplode = jp.NewPacket(jp.StatePlay, jp.S2C, 0x20)

type S2CExplodeData struct {
	//
	X ns.Double
	//
	Y ns.Double
	//
	Z ns.Double
	// The particle ID listed in Particles .
	ExplosionParticleId ns.VarInt
	// Particle data as specified in Particles .
	ExplosionParticleData ns.ByteArray // TODO: ParticleData
	// ID in the minecraft:sound_event registry, or an inline definition.
	ExplosionSound ns.Or[ns.Identifier, ns.SoundEvent]
}

// S2CForgetLevelChunk represents "Unload Chunk".
//
// > Tells the client to unload a chunk column.
// >
// > Note: The order is inverted, because the client reads this packet as one big-endian Long , with Z being the upper 32 bits.
// >
// > It is legal to send this packet even if the given chunk is not currently loaded.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Unload_Chunk
var S2CForgetLevelChunk = jp.NewPacket(jp.StatePlay, jp.S2C, 0x21)

type S2CForgetLevelChunkData struct {
	// Block coordinate divided by 16, rounded down.
	ChunkZ ns.Int
	// Block coordinate divided by 16, rounded down.
	ChunkX ns.Int
}

// S2CGameEvent represents "Game Event".
//
// > Used for a wide variety of game events, such as weather, respawn availability (from bed and respawn anchor ), game mode, some game rules, and demo messages.
// >
// > Events :
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Game_Event
var S2CGameEvent = jp.NewPacket(jp.StatePlay, jp.S2C, 0x22)

type S2CGameEventData struct {
	// See below.
	Event ns.UnsignedByte
	// Depends on Event.
	Value ns.Float
}

// S2CHorseScreenOpen represents "Open Horse Screen".
//
// > This packet is used exclusively for opening the horse GUI. Open Screen is used for all other GUIs. The client will not open the inventory if the Entity ID does not point to an
// > horse-like animal.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Open_Horse_Screen
var S2CHorseScreenOpen = jp.NewPacket(jp.StatePlay, jp.S2C, 0x23)

type S2CHorseScreenOpenData struct {
	// Same as the field of Open Screen .
	WindowId ns.VarInt
	// How many columns of horse inventory slots exist in the GUI, 3 slots per column.
	InventoryColumnsCount ns.VarInt
	// The "owner" entity of the GUI. The client should close the GUI if the owner entity dies or is cleared.
	EntityId ns.Int
}

// S2CHurtAnimation represents "Hurt Animation".
//
// > Plays a bobbing animation for the entity receiving damage.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Hurt_Animation
var S2CHurtAnimation = jp.NewPacket(jp.StatePlay, jp.S2C, 0x24)

type S2CHurtAnimationData struct {
	// The ID of the entity taking damage
	EntityId ns.VarInt
	// The direction the damage is coming from in relation to the entity
	Yaw ns.Float
}

// S2CInitializeBorder represents "Initialize World Border".
//
// > The vanilla client determines how solid to display the warning by comparing to whichever is higher, the warning distance or whichever is lower, the distance from the current diameter to the target diameter or the
// > place the border will be after warningTime seconds. In pseudocode:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Initialize_World_Border
var S2CInitializeBorder = jp.NewPacket(jp.StatePlay, jp.S2C, 0x25)

type S2CInitializeBorderData struct {
	//
	X ns.Double
	//
	Z ns.Double
	// Current length of a single side of the world border, in meters.
	OldDiameter ns.Double
	// Target length of a single side of the world border, in meters.
	NewDiameter ns.Double
	// Number of real-time milli seconds until New Diameter is reached. It appears that vanilla server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0.
	Speed ns.VarLong
	// Resulting coordinates from a portal teleport are limited to ±value. Usually 29999984.
	PortalTeleportBoundary ns.VarInt
	// In meters.
	WarningBlocks ns.VarInt
	// In seconds as set by /worldborder warning time .
	WarningTime ns.VarInt
}

// S2CKeepAlivePlay represents "Clientbound Keep Alive (play)".
//
// > The server will frequently send out a keep-alive, each containing a random ID. The client must respond with the same payload (see Serverbound Keep Alive ). If the
// > client does not respond to a Keep Alive packet within 15 seconds after it was sent, the server kicks the client. Vice versa, if the server does not send any keep-alives for 20 seconds, the client will disconnect
// > and yields a "Timed out" exception.
// >
// > The vanilla server uses a system-dependent time in milliseconds to generate the keep alive ID value.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clientbound_Keep_Alive_(Play)
var S2CKeepAlivePlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x26)

type S2CKeepAlivePlayData struct {
	//
	KeepAliveId ns.Long
}

// S2CLevelChunkWithLight represents "Chunk Data and Update Light".
//
// > Sent when a chunk comes into the client's view distance, specifying its terrain, lighting and block entities.
// >
// > The chunk must be within the view area previously specified with Set Center Chunk ; see that packet for details.
// >
// > It is not strictly necessary to send all block entities in this packet; it is still legal to send them with Block Entity Data later.
// >
// > Unlike the Update Light packet which uses the same format, setting the bit corresponding to a section to 0 in both of the block light or sky light masks does not appear to be useful,
// > and the results in testing have been highly inconsistent.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Chunk_Data_And_Update_Light
var S2CLevelChunkWithLight = jp.NewPacket(jp.StatePlay, jp.S2C, 0x27)

type S2CLevelChunkWithLightData struct {
	// Chunk coordinate (block coordinate divided by 16, rounded down)
	ChunkX ns.Int
	// Chunk coordinate (block coordinate divided by 16, rounded down)
	ChunkZ ns.Int
	//
	Data ns.ChunkData
	//
	Light ns.LightData
}

// S2CLevelEvent represents "World Event".
//
// > Sent when a client is to play a sound or particle effect.
// >
// > By default, the Minecraft client adjusts the volume of sound effects based on distance. The final boolean field is used to disable this, and instead the effect is played from 2 blocks away in the correct
// > direction. Currently this is only used for effect 1023 (wither spawn), effect 1028 (enderdragon death), and effect 1038 (end portal opening); it is ignored on other effects.
// >
// > Events:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#World_Event
var S2CLevelEvent = jp.NewPacket(jp.StatePlay, jp.S2C, 0x28)

type S2CLevelEventData struct {
	// The event, see below.
	Event ns.Int
	// The location of the event.
	Location ns.Position
	// Extra data for certain events, see below.
	Data ns.Int
	// See above.
	DisableRelativeVolume ns.Boolean
}

// S2CLevelParticles represents "Particle".
//
// > Displays the named particle
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Particle
var S2CLevelParticles = jp.NewPacket(jp.StatePlay, jp.S2C, 0x29)

type S2CLevelParticlesData struct {
	// If true, particle distance increases from 256 to 65536.
	LongDistance ns.Boolean
	// Whether this particle should always be visible.
	AlwaysVisible ns.Boolean
	// X position of the particle.
	X ns.Double
	// Y position of the particle.
	Y ns.Double
	// Z position of the particle.
	Z ns.Double
	// This is added to the X position after being multiplied by random.nextGaussian() .
	OffsetX ns.Float
	// This is added to the Y position after being multiplied by random.nextGaussian() .
	OffsetY ns.Float
	// This is added to the Z position after being multiplied by random.nextGaussian() .
	OffsetZ ns.Float
	//
	MaxSpeed ns.Float
	// The number of particles to create.
	ParticleCount ns.Int
	// The particle ID listed in Particles .
	ParticleId ns.VarInt
	// Particle data as specified in Particles .
	Data ns.ByteArray // TODO: ParticleData
}

// S2CLightUpdate represents "Update Light".
//
// > Updates light levels for a chunk. See Light for information on how lighting works in Minecraft.
// >
// > A bit will never be set in both the block light mask and the empty block light mask, though it may be present in neither of them (if the block light does not need to be updated for the corresponding chunk
// > section). The same applies to the sky light mask and the empty sky light mask.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Light
var S2CLightUpdate = jp.NewPacket(jp.StatePlay, jp.S2C, 0x2A)

type S2CLightUpdateData struct {
	// Chunk coordinate (block coordinate divided by 16, rounded down)
	ChunkX ns.VarInt
	// Chunk coordinate (block coordinate divided by 16, rounded down)
	ChunkZ ns.VarInt
	//
	Data ns.LightData
}

// S2CLoginPlay represents "Login (play)".
//
// > See protocol encryption for information on logging in.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Login_(Play)
var S2CLoginPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x2B)

type S2CLoginPlayData struct {
	// The player's Entity ID (EID).
	EntityId ns.Int
	//
	IsHardcore ns.Boolean
	// Identifiers for all dimensions on the server.
	DimensionNames ns.PrefixedArray[ns.Identifier]
	// Was once used by the client to draw the player list, but now is ignored.
	MaxPlayers ns.VarInt
	// Render distance (2-32).
	ViewDistance ns.VarInt
	// The distance that the client will process specific things, such as entities.
	SimulationDistance ns.VarInt
	// If true, a vanilla client shows reduced information on the debug screen . For servers in development, this should almost always be false.
	ReducedDebugInfo ns.Boolean
	// Set to false when the doImmediateRespawn gamerule is true.
	EnableRespawnScreen ns.Boolean
	// Whether players can only craft recipes they have already unlocked. Currently unused by the client.
	DoLimitedCrafting ns.Boolean
	// The ID of the type of dimension in the minecraft:dimension_type registry, defined by the Registry Data packet.
	DimensionType ns.VarInt
	// Name of the dimension being spawned into.
	DimensionName ns.Identifier
	// First 8 bytes of the SHA-256 hash of the world's seed. Used client side for biome noise
	HashedSeed ns.Long
	// 0: Survival, 1: Creative, 2: Adventure, 3: Spectator.
	GameMode ns.UnsignedByte
	// -1: Undefined (null), 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. The previous game mode. Vanilla client uses this for the debug (F3 + N & F3 + F4) game mode switch. (More information needed)
	PreviousGameMode ns.Byte
	// True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks.
	IsDebug ns.Boolean
	// True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63.
	IsFlat ns.Boolean
	// If true, then the next two fields are present.
	HasDeathLocation ns.Boolean
	// Name of the dimension the player died in.
	DeathDimensionName ns.Optional[ns.Identifier]
	// The location that the player died at.
	DeathLocation ns.Optional[ns.Position]
	// The number of ticks until the player can use the last used portal again. Looks like it's an attempt to fix MC-180.
	PortalCooldown ns.VarInt
	//
	SeaLevel ns.VarInt
	//
	EnforcesSecureChat ns.Boolean
}

// S2CMapItemData represents "Map Data".
//
// > Updates a rectangular area on a map item.
// >
// > For icons, a direction of 0 is a vertical icon and increments by 22.5° (360/16).
// >
// > Types are based off of rows and columns in map_icons.png :
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Map_Data
var S2CMapItemData = jp.NewPacket(jp.StatePlay, jp.S2C, 0x2C)

type S2CMapItemDataData struct {
	// Map ID of the map being modified
	MapId ns.VarInt
	// From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel)
	Scale ns.Byte
	// True if the map has been locked in a cartography table
	Locked ns.Boolean
	// Map coordinates: -128 for furthest left, +127 for furthest right
	X ns.Byte
	// Map coordinates: -128 for highest, +127 for lowest
	Z ns.Byte
	// 0-15
	Direction   ns.Byte
	DisplayName ns.PrefixedOptional[ns.TextComponent]
	// Only if Columns is more than 0; number of rows updated
	Rows ns.Optional[ns.UnsignedByte]
	// Only if Columns is more than 0; see Map item format
	Data ns.Optional[ns.PrefixedArray[ns.ByteArray]]
}

// S2CMerchantOffers represents "Merchant Offers".
//
// > The list of trades a villager NPC is offering.
// >
// > Trade Item:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Merchant_Offers
var S2CMerchantOffers = jp.NewPacket(jp.StatePlay, jp.S2C, 0x2D)

type S2CMerchantOffersData struct {
	// The ID of the window that is open; this is an int rather than a byte.
	WindowId ns.VarInt
	// The item the player will receive from this villager trade.
	OutputItem ns.Slot
	// The second item the player has to supply for this villager trade.
	InputItem2 ns.PrefixedOptional[ns.ByteArray] // TODO: TradeItem
	// True if the trade is disabled; false if the trade is enabled.
	TradeDisabled ns.Boolean
	// Number of times the trade has been used so far. If equal to the maximum number of trades, the client will display a red X.
	NumberOfTradeUses ns.Int
	// Number of times this trade can be used before it's exhausted.
	MaximumNumberOfTradeUses ns.Int
	// Amount of XP the villager will earn each time the trade is used.
	Xp ns.Int
	// Can be zero or negative. The number is added to the price when an item is discounted due to player reputation or other effects.
	SpecialPrice ns.Int
	// Can be low (0.05) or high (0.2). Determines how much demand, player reputation, and temporary effects will adjust the price.
	PriceMultiplier ns.Float
	// If positive, causes the price to increase. Negative values seem to be treated the same as zero.
	Demand ns.Int
	// Appears on the trade GUI; meaning comes from the translation key merchant.level. + level. 1: Novice, 2: Apprentice, 3: Journeyman, 4: Expert, 5: Master.
	VillagerLevel ns.VarInt
	// Total experience for this villager (always 0 for the wandering trader).
	Experience ns.VarInt
	// True if this is a regular villager; false for the wandering trader. When false, hides the villager level and some other GUI elements.
	IsRegularVillager ns.Boolean
	// True for regular villagers and false for the wandering trader. If true, the "Villagers restock up to two times per day." message is displayed when hovering over disabled trades.
	CanRestock ns.Boolean
}

// S2CMoveEntityPos represents "Update Entity Position".
//
// > This packet is sent by the server when an entity moves a small distance. The change in position is represented as a fixed-point number with 12 fraction bits and 4 integer bits.
// > As such, the maximum movement distance along each axis is 8 blocks in the negative direction, or 7.999755859375 blocks in the positive direction. If the movement exceeds these limits, Teleport Entity should be sent instead.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Entity_Position
var S2CMoveEntityPos = jp.NewPacket(jp.StatePlay, jp.S2C, 0x2E)

type S2CMoveEntityPosData struct {
	//
	EntityId ns.VarInt
	// Change in X position as currentX * 4096 - prevX * 4096 .
	DeltaX ns.Short
	// Change in Y position as currentY * 4096 - prevY * 4096 .
	DeltaY ns.Short
	// Change in Z position as currentZ * 4096 - prevZ * 4096 .
	DeltaZ ns.Short
	//
	OnGround ns.Boolean
}

// S2CMoveEntityPosRot represents "Update Entity Position and Rotation".
//
// > This packet is sent by the server when an entity rotates and moves. See #Update Entity Position for how the position is encoded.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Entity_Position_And_Rotation
var S2CMoveEntityPosRot = jp.NewPacket(jp.StatePlay, jp.S2C, 0x2F)

type S2CMoveEntityPosRotData struct {
	//
	EntityId ns.VarInt
	// Change in X position as currentX * 4096 - prevX * 4096 .
	DeltaX ns.Short
	// Change in Y position as currentY * 4096 - prevY * 4096 .
	DeltaY ns.Short
	// Change in Z position as currentZ * 4096 - prevZ * 4096 .
	DeltaZ ns.Short
	// New angle, not a delta.
	Yaw ns.Angle
	// New angle, not a delta.
	Pitch ns.Angle
	//
	OnGround ns.Boolean
}

// S2CMoveMinecartAlongTrack represents "Move Minecart Along Track".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Move_Minecart_Along_Track
var S2CMoveMinecartAlongTrack = jp.NewPacket(jp.StatePlay, jp.S2C, 0x30)

type S2CMoveMinecartAlongTrackData struct {
	//
	EntityId ns.VarInt
	//
	Y ns.Double
	//
	Z ns.Double
	//
	VelocityX ns.Double
	//
	VelocityY ns.Double
	//
	VelocityZ ns.Double
	//
	Yaw ns.Angle
	//
	Pitch ns.Angle
}

// S2CMoveEntityRot represents "Update Entity Rotation".
//
// > This packet is sent by the server when an entity rotates.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Entity_Rotation
var S2CMoveEntityRot = jp.NewPacket(jp.StatePlay, jp.S2C, 0x31)

type S2CMoveEntityRotData struct {
	//
	EntityId ns.VarInt
	// New angle, not a delta.
	Yaw ns.Angle
	// New angle, not a delta.
	Pitch ns.Angle
	//
	OnGround ns.Boolean
}

// S2CMoveVehicle represents "Move Vehicle".
//
// > Note that all fields use absolute positioning and do not allow for relative positioning.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Move_Vehicle
var S2CMoveVehicle = jp.NewPacket(jp.StatePlay, jp.S2C, 0x32)

type S2CMoveVehicleData struct {
	// Absolute position (X coordinate).
	X ns.Double
	// Absolute position (Y coordinate).
	Y ns.Double
	// Absolute position (Z coordinate).
	Z ns.Double
	// Absolute rotation on the vertical axis, in degrees.
	Yaw ns.Float
	// Absolute rotation on the horizontal axis, in degrees.
	Pitch ns.Float
}

// S2COpenBook represents "Open Book".
//
// > Sent when a player right clicks with a signed book. This tells the client to open the book GUI.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Open_Book
var S2COpenBook = jp.NewPacket(jp.StatePlay, jp.S2C, 0x33)

type S2COpenBookData struct {
	// 0: Main hand, 1: Off hand .
	Hand ns.VarInt
}

// S2COpenScreen represents "Open Screen".
//
// > This is sent to the client when it should open an inventory, such as a chest, workbench, furnace, or other container. Resending this packet with already existing window id, will update the window title and window
// > type without closing the window.
// >
// > This message is not sent to clients opening their own inventory, nor do clients inform the server in any way when doing so. From the server's perspective, the inventory is always "open" whenever no other windows
// > are.
// >
// > For horses, use Open Horse Screen .
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Open_Screen
var S2COpenScreen = jp.NewPacket(jp.StatePlay, jp.S2C, 0x34)

type S2COpenScreenData struct {
	// An identifier for the window to be displayed. vanilla server implementation is a counter, starting at 1. There can only be one window at a time; this is only used to ignore outdated packets targeting already-closed windows. Note also that the Window ID field in most other packets is only a single byte, and indeed, the vanilla server wraps around after 100.
	WindowId ns.VarInt
	// The window type to use for display. Contained in the minecraft:menu registry; see Inventory for the different values.
	WindowType ns.VarInt
	// The title of the window.
	WindowTitle ns.TextComponent
}

// S2COpenSignEditor represents "Open Sign Editor".
//
// > Sent when the client has placed a sign and is allowed to send Update Sign . There must already be a sign at the given location (which the client does not do automatically) - send a Block Update first.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Open_Sign_Editor
var S2COpenSignEditor = jp.NewPacket(jp.StatePlay, jp.S2C, 0x35)

type S2COpenSignEditorData struct {
	//
	Location ns.Position
	// Whether the opened editor is for the front or on the back of the sign
	IsFrontText ns.Boolean
}

// S2CPingPlay represents "Ping (play)".
//
// > Packet is not used by the vanilla server. When sent to the client, client responds with a Pong packet with the same id.
// >
// > Unlike Keep Alive this packet is handled synchronously with game logic on the vanilla client, and can thus be used to reliably detect which serverbound packets were
// > sent after the ping and all preceding clientbound packets were received and handled.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Ping_(Play)
var S2CPingPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x36)

type S2CPingPlayData struct {
	//
	Id ns.Int
}

// S2CPongResponsePlay represents "Ping Response (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Ping_Response_(Play)
var S2CPongResponsePlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x37)

type S2CPongResponsePlayData struct {
	// Should be the same as sent by the client.
	Payload ns.Long
}

// S2CPlaceGhostRecipe represents "Place Ghost Recipe".
//
// > Response to the serverbound packet ( Place Recipe ), with the same recipe ID. Appears to be used to notify the UI.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Place_Ghost_Recipe
var S2CPlaceGhostRecipe = jp.NewPacket(jp.StatePlay, jp.S2C, 0x38)

type S2CPlaceGhostRecipeData struct {
	//
	WindowId ns.VarInt
	//
	RecipeDisplay ns.RecipeDisplay
}

// S2CPlayerAbilities represents "Player Abilities (clientbound)".
//
// > The latter 2 floats are used to indicate the flying speed and field of view respectively, while the first byte is used to determine the value of 4 booleans.
// >
// > About the flags:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Abilities_(Clientbound)
var S2CPlayerAbilities = jp.NewPacket(jp.StatePlay, jp.S2C, 0x39)

type S2CPlayerAbilitiesData struct {
	// Bit field, see below.
	Flags ns.Byte
	// 0.05 by default.
	FlyingSpeed ns.Float
	// Modifies the field of view, like a speed potion. A vanilla server will use the same value as the movement speed sent in the Update Attributes packet, which defaults to 0.1 for players.
	FieldOfViewModifier ns.Float
}

// S2CPlayerChat represents "Player Chat Message".
//
// > Sends the client a chat message from a player.
// >
// > Currently a lot is unknown about this packet, blank descriptions are for those that are unknown
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Chat_Message
var S2CPlayerChat = jp.NewPacket(jp.StatePlay, jp.S2C, 0x3A)

type S2CPlayerChatData struct {
	GlobalIndex ns.VarInt
	// Used by the vanilla client for the disableChat launch option. Setting both longs to 0 will always display the message regardless of the setting.
	Sender ns.UUID
	//
	Index ns.VarInt
	// Cryptography, the signature consists of the Sender UUID, Session UUID from the Player Session packet, Index, Salt, Timestamp in epoch seconds, the length of the original chat content, the original content itself, the length of Previous Messages, and all of the Previous message signatures. These values are hashed with SHA-256 and signed using the RSA cryptosystem. Modifying any of these values in the packet will cause this signature to fail. This buffer is always 256 bytes long and it is not length-prefixed.
	MessageSignatureBytes ns.PrefixedOptional[ns.ByteArray]
	// Raw (optionally) signed sent message content. This is used as the content parameter when formatting the message on the client.
	Message ns.String
	// Represents the time the message was signed as milliseconds since the epoch , used to check if the message was received within 2 minutes of it being sent.
	Timestamp ns.Long
	// Cryptography, used for validating the message signature.
	Salt ns.Long
	// The previous message's signature. Contains the same type of data as Message Signature bytes (256 bytes) above. Not length-prefxied.
	SignatureData ns.PrefixedArray[struct {
		// The message Id + 1, used for validating message signature. The next field is present only when value of this field is equal to 0.
		MessageID ns.VarInt
		// The previous message's signature. Contains the same type of data as Message Signature bytes (256 bytes) above. Not length-prefxied.
		Signature ns.Optional[ns.ByteArray]
	}]
	// The original message content, before filtering.
	UnsignedContent ns.PrefixedOptional[ns.TextComponent]
	// If the message has been filtered
	FilterType ns.VarInt
	// Only present if the Filter Type is Partially Filtered. Specifies the indexes at which characters in the original message string should be replaced with the # symbol (i.e. filtered) by the vanilla client
	FilterTypeBits ns.Optional[ns.BitSet]
	// Either the type of chat in the minecraft:chat_type registry, defined by the Registry Data packet, or an inline definition. 
	ChatType ns.Or[ns.Identifier, ns.ChatType]
	// The name of the one sending the message, usually the sender's display name. This is used as the sender parameter when formatting the message on the client.
	SenderName ns.TextComponent
	// The name of the one receiving the message, usually the receiver's display name. This is used as the target parameter when formatting the message on the client.
	TargetName ns.PrefixedOptional[ns.TextComponent]
}

// S2CPlayerCombatEnd represents "End Combat".
//
// > Unused by the vanilla client. This data was once used for twitch.tv metadata circa 1.8.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#End_Combat
var S2CPlayerCombatEnd = jp.NewPacket(jp.StatePlay, jp.S2C, 0x3B)

type S2CPlayerCombatEndData struct {
	// Length of the combat in ticks.
	Duration ns.VarInt
}

// S2CPlayerCombatEnter represents "Enter Combat".
//
// > Unused by the vanilla client. This data was once used for twitch.tv metadata circa 1.8.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Enter_Combat
var S2CPlayerCombatEnter = jp.NewPacket(jp.StatePlay, jp.S2C, 0x3C)

type S2CPlayerCombatEnterData struct {
	// No fields
}

// S2CPlayerCombatKill represents "Combat Death".
//
// > Used to send a respawn screen.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Combat_Death
var S2CPlayerCombatKill = jp.NewPacket(jp.StatePlay, jp.S2C, 0x3D)

type S2CPlayerCombatKillData struct {
	// Entity ID of the player that died (should match the client's entity ID).
	PlayerId ns.VarInt
	// The death message.
	Message ns.TextComponent
}

// S2CPlayerInfoRemove represents "Player Info Remove".
//
// > Used by the server to remove players from the player list.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Info_Remove
var S2CPlayerInfoRemove = jp.NewPacket(jp.StatePlay, jp.S2C, 0x3E)

type S2CPlayerInfoRemoveData struct {
	// UUIDs of players to remove.
	Uuids ns.PrefixedArray[ns.UUID]
}

// S2CPlayerInfoUpdate represents "Player Info Update".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Info_Update
var S2CPlayerInfoUpdate = jp.NewPacket(jp.StatePlay, jp.S2C, 0x3F)

type S2CPlayerInfoUpdateData struct {
	// Determines what actions are present.
	Actions ns.EnumSet
	// The length of this array is determined by the number of Player Actions that give a non-zero value when applying its mask to the actions flag. For example given the decimal number 5, binary 00000101. The masks 0x01 and 0x04 would return a non-zero value, meaning the Player Actions array would include two actions: Add Player and Update Game Mode.
	PlayerActions ns.ByteArray // TODO: https://minecraft.wiki/w/Java_Edition_protocol/Packets#player-info:player-actions
}

// S2CPlayerLookAt represents "Look At".
//
// > Used to rotate the client player to face the given location or entity (for /teleport [<targets>] <x> <y> <z> facing ).
// >
// > If the entity given by entity ID cannot be found, this packet should be treated as if is entity was false.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Look_At
var S2CPlayerLookAt = jp.NewPacket(jp.StatePlay, jp.S2C, 0x40)

type S2CPlayerLookAtData struct {
	// Values are feet=0, eyes=1. If set to eyes, aims using the head position; otherwise aims using the feet position.
	FeetEyes ns.VarInt
	// x coordinate of the point to face towards.
	TargetX ns.Double
	// y coordinate of the point to face towards.
	TargetY ns.Double
	// z coordinate of the point to face towards.
	TargetZ ns.Double
	// If true, additional information about an entity is provided.
	IsEntity ns.Boolean
	// Only if is entity is true — the entity to face towards.
	EntityId ns.Optional[ns.VarInt]
	// Whether to look at the entity's eyes or feet. Same values and meanings as before, just for the entity's head/feet.
	EntityFeetEyes ns.Optional[ns.VarInt]
}

// S2CPlayerPosition represents "Synchronize Player Position".
//
// > Teleports the client, e.g. during login, when using an ender pearl, in response to invalid move packets, etc.
// >
// > Due to latency, the server may receive outdated movement packets sent before the client was aware of the teleport. To account for this, the server ignores all movement packets from the client until a Confirm Teleportation packet with an ID matching the one sent in the teleport packet is received.
// >
// > Yaw is measured in degrees, and does not follow classical trigonometry rules. The unit circle of yaw on the XZ-plane starts at (0, 1) and turns counterclockwise, with 90 at (-1, 0), 180 at (0, -1) and 270 at (1,
// > 0). Additionally, yaw is not clamped to between 0 and 360 degrees; any number is valid, including negative numbers and numbers greater than 360 (see MC-90097 ).
// >
// > Pitch is measured in degrees, where 0 is looking straight ahead, -90 is looking straight up, and 90 is looking straight down.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Synchronize_Player_Position
var S2CPlayerPosition = jp.NewPacket(jp.StatePlay, jp.S2C, 0x41)

type S2CPlayerPositionData struct {
	// Client should confirm this packet with Confirm Teleportation containing the same Teleport ID.
	TeleportId ns.VarInt
	// Absolute or relative position, depending on Flags.
	X ns.Double
	// Absolute or relative position, depending on Flags.
	Y ns.Double
	// Absolute or relative position, depending on Flags.
	Z ns.Double
	//
	VelocityX ns.Double
	//
	VelocityY ns.Double
	//
	VelocityZ ns.Double
	// Absolute or relative rotation on the X axis, in degrees.
	Yaw ns.Float
	// Absolute or relative rotation on the Y axis, in degrees.
	Pitch ns.Float
	//
	Flags ns.TeleportFlags
}

// S2CPlayerRotation represents "Player Rotation".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Player_Rotation
var S2CPlayerRotation = jp.NewPacket(jp.StatePlay, jp.S2C, 0x42)

type S2CPlayerRotationData struct {
	// Rotation on the X axis, in degrees.
	Yaw ns.Float
	// Rotation on the Y axis, in degrees.
	Pitch ns.Float
}

// S2CRecipeBookAdd represents "Recipe Book Add".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Recipe_Book_Add
var S2CRecipeBookAdd = jp.NewPacket(jp.StatePlay, jp.S2C, 0x43)

type S2CRecipeBookAddData struct {
	// Prefixed Array
	Recipes ns.PrefixedArray[struct {
		RecipeId    ns.VarInt
		Display     ns.RecipeDisplay
		GroupId     ns.VarInt
		CategoryId  ns.VarInt
		Ingredients ns.PrefixedOptional[ns.PrefixedArray[ns.ByteArray]] // TODO: ID Set (https://minecraft.wiki/w/Java_Edition_protocol/Packets#ID_Set)
		Flags       ns.Byte
	}]
	// Replace or Add to known recipes
	Replace ns.Boolean
}

// S2CRecipeBookRemove represents "Recipe Book Remove".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Recipe_Book_Remove
var S2CRecipeBookRemove = jp.NewPacket(jp.StatePlay, jp.S2C, 0x44)

type S2CRecipeBookRemoveData struct {
	// IDs of recipes to remove.
	Recipes ns.PrefixedArray[ns.VarInt]
}

// S2CRecipeBookSettings represents "Recipe Book Settings".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Recipe_Book_Settings
var S2CRecipeBookSettings = jp.NewPacket(jp.StatePlay, jp.S2C, 0x45)

type S2CRecipeBookSettingsData struct {
	// If true, then the crafting recipe book will be open when the player opens its inventory.
	CraftingRecipeBookOpen ns.Boolean
	// If true, then the filtering option is active when the players opens its inventory.
	CraftingRecipeBookFilterActive ns.Boolean
	// If true, then the smelting recipe book will be open when the player opens its inventory.
	SmeltingRecipeBookOpen ns.Boolean
	// If true, then the filtering option is active when the players opens its inventory.
	SmeltingRecipeBookFilterActive ns.Boolean
	// If true, then the blast furnace recipe book will be open when the player opens its inventory.
	BlastFurnaceRecipeBookOpen ns.Boolean
	// If true, then the filtering option is active when the players opens its inventory.
	BlastFurnaceRecipeBookFilterActive ns.Boolean
	// If true, then the smoker recipe book will be open when the player opens its inventory.
	SmokerRecipeBookOpen ns.Boolean
	// If true, then the filtering option is active when the players opens its inventory.
	SmokerRecipeBookFilterActive ns.Boolean
}

// S2CRemoveEntities represents "Remove Entities".
//
// > Sent by the server when an entity is to be destroyed on the client.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Remove_Entities
var S2CRemoveEntities = jp.NewPacket(jp.StatePlay, jp.S2C, 0x46)

type S2CRemoveEntitiesData struct {
	// The list of entities to destroy.
	EntityIds ns.PrefixedArray[ns.VarInt]
}

// S2CRemoveMobEffect represents "Remove Entity Effect".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Remove_Entity_Effect
var S2CRemoveMobEffect = jp.NewPacket(jp.StatePlay, jp.S2C, 0x47)

type S2CRemoveMobEffectData struct {
	//
	EntityId ns.VarInt
	// See this table .
	EffectId ns.VarInt
}

// S2CResetScore represents "Reset Score".
//
// > This is sent to the client when it should remove a scoreboard item.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Reset_Score
var S2CResetScore = jp.NewPacket(jp.StatePlay, jp.S2C, 0x48)

type S2CResetScoreData struct {
	// The entity whose score this is. For players, this is their username; for other entities, it is their UUID.
	EntityName ns.String
	// The name of the objective the score belongs to.
	ObjectiveName ns.PrefixedOptional[ns.String]
}

// S2CResourcePackPopPlay represents "Remove Resource Pack (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Remove_Resource_Pack_(Play)
var S2CResourcePackPopPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x49)

type S2CResourcePackPopPlayData struct {
	// The UUID of the resource pack to be removed.
	Uuid ns.Optional[ns.UUID]
}

// S2CResourcePackPushPlay represents "Add Resource Pack (play)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Add_Resource_Pack_(Play)
var S2CResourcePackPushPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x4A)

type S2CResourcePackPushPlayData struct {
	// The unique identifier of the resource pack.
	Uuid ns.UUID
	// The URL to the resource pack.
	Url ns.String
	// A 40 character hexadecimal, case-insensitive SHA-1 hash of the resource pack file. If it's not a 40 character hexadecimal string, the client will not use it for hash verification and likely waste bandwidth.
	Hash ns.String
	// The vanilla client will be forced to use the resource pack from the server. If they decline they will be kicked from the server.
	Forced ns.Boolean
	// This is shown in the prompt making the client accept or decline the resource pack.
	PromptMessage ns.PrefixedOptional[ns.TextComponent]
}

// S2CRespawn represents "Respawn".
//
// > To change the player's dimension (overworld/nether/end), send them a respawn packet with the appropriate dimension, followed by prechunks/chunks for the new dimension, and finally a position and look packet. You
// > do not need to unload chunks, the client will do it automatically.
// >
// > The background of the loading screen is determined based on the Dimension Name specified in this packet, and the one specified in the previous Login or Respawn packet. If either the current or the previous
// > dimension is minecraft:nether , the Nether portal background is used. Otherwise, if the current or the previous dimension is minecraft:the_end , the End portal background is used. If the
// > player is dead (health is 0), the default background is always used.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Respawn
var S2CRespawn = jp.NewPacket(jp.StatePlay, jp.S2C, 0x4B)

type S2CRespawnData struct {
	// The ID of type of dimension in the minecraft:dimension_type registry, defined by the Registry Data packet.
	DimensionType ns.VarInt
	// Name of the dimension being spawned into.
	DimensionName ns.Identifier
	// First 8 bytes of the SHA-256 hash of the world's seed. Used client side for biome noise
	HashedSeed ns.Long
	// 0: Survival, 1: Creative, 2: Adventure, 3: Spectator.
	GameMode ns.UnsignedByte
	// -1: Undefined (null), 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. The previous game mode. Vanilla client uses this for the debug (F3 + N & F3 + F4) game mode switch. (More information needed)
	PreviousGameMode ns.Byte
	// True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks.
	IsDebug ns.Boolean
	// True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63.
	IsFlat ns.Boolean
	// If true, then the next two fields are present.
	HasDeathLocation ns.Boolean
	// Name of the dimension the player died in.
	DeathDimensionName ns.Optional[ns.Identifier]
	// The location that the player died at.
	DeathLocation ns.Optional[ns.Position]
	// The number of ticks until the player can use the portal again.
	PortalCooldown ns.VarInt
	//
	SeaLevel ns.VarInt
	// Bit mask. 0x01: Keep attributes, 0x02: Keep metadata. Tells which data should be kept on the client side once the player has respawned. In the vanilla implementation, this is context dependent: normal respawns (after death) keep no data; exiting the end poem/credits keeps the attributes; other dimension changes (portals or teleports) keep all data.
	DataKept ns.Byte
}

// S2CRotateHead represents "Set Head Rotation".
//
// > Changes the direction an entity's head is facing.
// >
// > While sending the Entity Look packet changes the vertical rotation of the head, sending this packet appears to be necessary to rotate the head horizontally.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Head_Rotation
var S2CRotateHead = jp.NewPacket(jp.StatePlay, jp.S2C, 0x4C)

type S2CRotateHeadData struct {
	//
	EntityId ns.VarInt
	// New angle, not a delta.
	HeadYaw ns.Angle
}

// S2CSectionBlocksUpdate represents "Update Section Blocks".
//
// > Chunk section position is encoded:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Section_Blocks
var S2CSectionBlocksUpdate = jp.NewPacket(jp.StatePlay, jp.S2C, 0x4D)

type S2CSectionBlocksUpdateData struct {
	// Chunk section coordinate (encoded chunk x and z with each 22 bits, and section y with 20 bits, from left to right).
	ChunkSectionPosition ns.Long
	// Each entry is composed of the block state id, shifted left by 12, and the relative block position in the chunk section (4 bits for x, z, and y, from left to right).
	Blocks ns.PrefixedArray[ns.VarLong]
}

// S2CSelectAdvancementsTab represents "Select Advancements Tab".
//
// > Sent by the server to indicate that the client should switch advancement tab. Sent either when the client switches tab in the GUI or when an advancement in another tab is made.
// >
// > The Identifier must be one of the following if no custom data pack is loaded:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Select_Advancements_Tab
var S2CSelectAdvancementsTab = jp.NewPacket(jp.StatePlay, jp.S2C, 0x4E)

type S2CSelectAdvancementsTabData struct {
	// See below.
	Identifier ns.PrefixedOptional[ns.Identifier]
}

// S2CServerData represents "Server Data".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Server_Data
var S2CServerData = jp.NewPacket(jp.StatePlay, jp.S2C, 0x4F)

type S2CServerDataData struct {
	//
	Motd ns.TextComponent
	// Icon bytes in the PNG format.
	Icon ns.PrefixedOptional[ns.PrefixedByteArray]
}

// S2CSetActionBarText represents "Set Action Bar Text".
//
// > Displays a message above the hotbar. Equivalent to System Chat Message with Overlay set to true, except that chat message blocking isn't performed. Used by the
// > vanilla server only to implement the /title command.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Action_Bar_Text
var S2CSetActionBarText = jp.NewPacket(jp.StatePlay, jp.S2C, 0x50)

type S2CSetActionBarTextData struct {
	// No fields
}

// S2CSetBorderCenter represents "Set Border Center".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Center
var S2CSetBorderCenter = jp.NewPacket(jp.StatePlay, jp.S2C, 0x51)

type S2CSetBorderCenterData struct {
	//
	X ns.Double
	//
	Z ns.Double
}

// S2CSetBorderLerpSize represents "Set Border Lerp Size".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Lerp_Size
var S2CSetBorderLerpSize = jp.NewPacket(jp.StatePlay, jp.S2C, 0x52)

type S2CSetBorderLerpSizeData struct {
	// Current length of a single side of the world border, in meters.
	OldDiameter ns.Double
	// Target length of a single side of the world border, in meters.
	NewDiameter ns.Double
	// Number of real-time milli seconds until New Diameter is reached. It appears that vanilla server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0.
	Speed ns.VarLong
}

// S2CSetBorderSize represents "Set Border Size".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Size
var S2CSetBorderSize = jp.NewPacket(jp.StatePlay, jp.S2C, 0x53)

type S2CSetBorderSizeData struct {
	// Length of a single side of the world border, in meters.
	Diameter ns.Double
}

// S2CSetBorderWarningDelay represents "Set Border Warning Delay".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Warning_Delay
var S2CSetBorderWarningDelay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x54)

type S2CSetBorderWarningDelayData struct {
	// In seconds as set by /worldborder warning time .
	WarningTime ns.VarInt
}

// S2CSetBorderWarningDistance represents "Set Border Warning Distance".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Border_Warning_Distance
var S2CSetBorderWarningDistance = jp.NewPacket(jp.StatePlay, jp.S2C, 0x55)

type S2CSetBorderWarningDistanceData struct {
	// In meters.
	WarningBlocks ns.VarInt
}

// S2CSetCamera represents "Set Camera".
//
// > Sets the entity that the player renders from. This is normally used when the player left-clicks an entity while in spectator mode.
// >
// > The player's camera will move with the entity and look where it is looking. The entity is often another player, but can be any type of entity. The player is unable to move this entity (move packets will act as if
// > they are coming from the other entity).
// >
// > If the given entity is not loaded by the player, this packet is ignored. To return control to the player, send this packet with their entity ID.
// >
// > The vanilla server resets this (sends it back to the default entity) whenever the spectated entity is killed or the player sneaks, but only if they were spectating an entity. It also sends this packet whenever
// > the player switches out of spectator mode (even if they weren't spectating an entity).
// >
// > The vanilla client also loads certain shaders for given entities:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Camera
var S2CSetCamera = jp.NewPacket(jp.StatePlay, jp.S2C, 0x56)

type S2CSetCameraData struct {
	// ID of the entity to set the client's camera to.
	CameraId ns.VarInt
}

// S2CSetChunkCacheCenter represents "Set Center Chunk".
//
// > Sets the center position of the client's chunk loading area. The area is square-shaped, spanning 2 × server view distance + 7 chunks on both axes (width, not radius!). Since the area's width is always an odd
// > number, there is no ambiguity as to which chunk is the center.
// >
// > The vanilla client never renders or simulates chunks located outside the loading area, but keeps them in memory (unless explicitly unloaded by the server while still in range), and only automatically unloads a
// > chunk when another chunk is loaded at coordinates congruent to the old chunk's coordinates modulo (2 × server view distance + 7). This means that a chunk may reappear after leaving and later re-entering the
// > loading area through successive uses of this packet, unless it is replaced in the meantime by a different chunk in the same "slot".
// >
// > The vanilla client ignores attempts to load or unload chunks located outside the loading area. This applies even to unloads targeting chunks that are still loaded, but currently located outside the loading area
// > (per the previous paragraph).
// >
// > The vanilla server does not rely on any specific behavior for chunks leaving the loading area, and custom clients need not replicate the above exactly. A client may instead choose to immediately unload any chunks
// > outside the loading area, to use a different modulus, or to ignore the loading area completely and keep chunks loaded regardless of their location until the server requests to unload them. Servers aiming for
// > maximal interoperability should always explicitly unload any loaded chunks before they go outside the loading area.
// >
// > The center chunk is normally the chunk the player is in, but apart from the implications on chunk loading, the (vanilla) client takes no issue with this not being the case. Indeed, as long as chunks are sent only
// > within the default loading area centered on the world origin, it is not necessary to send this packet at all. This may be useful for servers with small bounded worlds, such as minigames, since it ensures chunks
// > never need to be resent after the client has joined, saving on bandwidth.
// >
// > The vanilla server sends this packet whenever the player moves across a chunk border horizontally, and also (according to testing) for any integer change in the vertical axis, even if it doesn't go across a chunk
// > section border.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Center_Chunk
var S2CSetChunkCacheCenter = jp.NewPacket(jp.StatePlay, jp.S2C, 0x57)

type S2CSetChunkCacheCenterData struct {
	// Chunk X coordinate of the loading area center.
	ChunkX ns.VarInt
	// Chunk Z coordinate of the loading area center.
	ChunkZ ns.VarInt
}

// S2CSetChunkCacheRadius represents "Set Render Distance".
//
// > Sent by the integrated singleplayer server when changing render distance. This packet is sent by the server when the client reappears in the overworld after leaving the end.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Render_Distance
var S2CSetChunkCacheRadius = jp.NewPacket(jp.StatePlay, jp.S2C, 0x58)

type S2CSetChunkCacheRadiusData struct {
	// Render distance (2-32).
	ViewDistance ns.VarInt
}

// S2CSetCursorItem represents "Set Cursor Item".
//
// > Replaces or sets the inventory item that's being dragged with the mouse.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Cursor_Item
var S2CSetCursorItem = jp.NewPacket(jp.StatePlay, jp.S2C, 0x59)

type S2CSetCursorItemData struct {
	//
	CarriedItem ns.Slot
}

// S2CSetDefaultSpawnPosition represents "Set Default Spawn Position".
//
// > Sent by the server after login to specify the coordinates of the spawn point (the point at which players spawn at, and which the compass points to). It can be sent at any time to update the point compasses point
// > at.
// >
// > The client uses this as the default position of the player upon spawning, though it's a good idea to always override this default by sending Synchronize Player Position .
// > When converting the position to floating point, 0.5 is added to the x and z coordinates and 1.0 to the y coordinate, so as to place the player centered on top of the specified block position.
// >
// > Before receiving this packet, the client uses the default position 8, 64, 8, and angle 0.0 (resulting in a default player spawn position of 8.5, 65.0, 8.5).
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Default_Spawn_Position
var S2CSetDefaultSpawnPosition = jp.NewPacket(jp.StatePlay, jp.S2C, 0x5A)

type S2CSetDefaultSpawnPositionData struct {
	// Spawn location.
	Location ns.Position
	// The angle at which to respawn at.
	Angle ns.Float
}

// S2CSetDisplayObjective represents "Display Objective".
//
// > This is sent to the client when it should display a scoreboard.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Display_Objective
var S2CSetDisplayObjective = jp.NewPacket(jp.StatePlay, jp.S2C, 0x5B)

type S2CSetDisplayObjectiveData struct {
	// The position of the scoreboard. 0: list, 1: sidebar, 2: below name, 3 - 18: team specific sidebar, indexed as 3 + team color.
	Position ns.VarInt
	// The unique name for the scoreboard to be displayed.
	ScoreName ns.String
}

// S2CSetEntityData represents "Set Entity Metadata".
//
// > Updates one or more metadata properties for an existing entity. Any properties not
// > included in the Metadata field are left unchanged.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Entity_Metadata
var S2CSetEntityData = jp.NewPacket(jp.StatePlay, jp.S2C, 0x5C)

type S2CSetEntityDataData struct {
	//
	EntityId ns.VarInt
	//
	Metadata ns.EntityMetadata
}

// S2CSetEntityLink represents "Link Entities".
//
// > This packet is sent when an entity has been leashed to another entity.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Link_Entities
var S2CSetEntityLink = jp.NewPacket(jp.StatePlay, jp.S2C, 0x5D)

type S2CSetEntityLinkData struct {
	// Attached entity's EID.
	AttachedEntityId ns.Int
	// ID of the entity holding the lead. Set to -1 to detach.
	HoldingEntityId ns.Int
}

// S2CSetEntityMotion represents "Set Entity Velocity".
//
// > Velocity is in units of 1/8000 of a block per server tick (50ms); for example, -1343 would move (-1343 / 8000) = −0.167875 blocks per tick (or −3.3575 blocks per second).
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Entity_Velocity
var S2CSetEntityMotion = jp.NewPacket(jp.StatePlay, jp.S2C, 0x5E)

type S2CSetEntityMotionData struct {
	//
	EntityId ns.VarInt
	// Velocity on the X axis.
	VelocityX ns.Short
	// Velocity on the Y axis.
	VelocityY ns.Short
	// Velocity on the Z axis.
	VelocityZ ns.Short
}

// S2CSetEquipment represents "Set Equipment".
//
// > Equipment slot can be one of the following:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Equipment
var S2CSetEquipment = jp.NewPacket(jp.StatePlay, jp.S2C, 0x5F)

type S2CSetEquipmentData struct {
	// Entity's ID.
	EntityId ns.VarInt
	//
	Item ns.Slot
}

// S2CSetExperience represents "Set Experience".
//
// > Sent by the server when the client should change experience levels.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Experience
var S2CSetExperience = jp.NewPacket(jp.StatePlay, jp.S2C, 0x60)

type S2CSetExperienceData struct {
	// Between 0 and 1.
	ExperienceBar ns.Float
	//
	Level ns.VarInt
	// See Experience#Leveling up on the Minecraft Wiki for Total Experience to Level conversion.
	TotalExperience ns.VarInt
}

// S2CSetHealth represents "Set Health".
//
// > Sent by the server to set the health of the player it is sent to.
// >
// > Food saturation acts as a food “overcharge”. Food values will not decrease while the saturation is over zero. New players logging in or respawning
// > automatically get a saturation of 5.0. Eating food increases the saturation as well as the food bar.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Health
var S2CSetHealth = jp.NewPacket(jp.StatePlay, jp.S2C, 0x61)

type S2CSetHealthData struct {
	// 0 or less = dead, 20 = full HP.
	Health ns.Float
	// 0–20.
	Food ns.VarInt
	// Seems to vary from 0.0 to 5.0 in integer increments.
	FoodSaturation ns.Float
}

// S2CSetHeldSlot represents "Set Held Item (clientbound)".
//
// > Sent to change the player's slot selection.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Held_Item_(Clientbound)
var S2CSetHeldSlot = jp.NewPacket(jp.StatePlay, jp.S2C, 0x62)

type S2CSetHeldSlotData struct {
	// The slot which the player has selected (0–8).
	Slot ns.VarInt
}

// S2CSetObjective represents "Update Objectives".
//
// > This is sent to the client when it should create a new scoreboard objective or remove one.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Objectives
var S2CSetObjective = jp.NewPacket(jp.StatePlay, jp.S2C, 0x63)

type S2CSetObjectiveData struct {
	// A unique name for the objective.
	ObjectiveName ns.String
	// 0 to create the scoreboard. 1 to remove the scoreboard. 2 to update the display text.
	Mode ns.Byte
	// Only if mode is 0 or 2.The text to be displayed for the score.
	ObjectiveValue ns.Optional[ns.TextComponent]
	// Only if mode is 0 or 2. 0 = "integer", 1 = "hearts".
	Type ns.Optional[ns.VarInt]
	// Only if mode is 0 or 2. Whether this objective has a set number format for the scores.
	HasNumberFormat ns.Optional[ns.Boolean]
	// Only if mode is 0 or 2 and the previous boolean is true. Determines how the score number should be formatted.
	NumberFormat ns.Optional[ns.VarInt]
	// Show nothing.
	StylingContent ns.ByteArray // TODO
}

// S2CSetPassengers represents "Set Passengers".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Passengers
var S2CSetPassengers = jp.NewPacket(jp.StatePlay, jp.S2C, 0x64)

type S2CSetPassengersData struct {
	// Vehicle's EID.
	EntityId ns.VarInt
	// EIDs of entity's passengers.
	Passengers ns.PrefixedArray[ns.VarInt]
}

// S2CSetPlayerInventory represents "Set Player Inventory Slot".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Player_Inventory_Slot
var S2CSetPlayerInventory = jp.NewPacket(jp.StatePlay, jp.S2C, 0x65)

type S2CSetPlayerInventoryData struct {
	//
	Slot ns.VarInt
	//
	SlotData ns.Slot
}

// S2CSetPlayerTeam represents "Update Teams".
//
// > Creates and updates teams.
// >
// > Team Color: The color of a team defines how the names of the team members are visualized; any formatting code can be used. The following table lists all the possible values.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Teams
var S2CSetPlayerTeam = jp.NewPacket(jp.StatePlay, jp.S2C, 0x66)

type S2CSetPlayerTeamData struct {
	// A unique name for the team. (Shared with scoreboard).
	TeamName ns.String
	// Determines the layout of the remaining packet.
	Method ns.Byte
	// Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team.
	FriendlyFlags ns.Byte
	// 0 = ALWAYS, 1 = NEVER, 2 = HIDE_FOR_OTHER_TEAMS, 3 = HIDE_FOR_OWN_TEAMS
	NameTagVisibility ns.VarInt
	// 0 = ALWAYS, 1 = NEVER, 2 = PUSH_OTHER_TEAMS, 3 = PUSH_OWN_TEAM
	CollisionRule ns.VarInt
	// Used to color the name of players on the team; see below.
	TeamColor ns.VarInt
	// Displayed before the names of players that are part of this team.
	TeamPrefix ns.TextComponent
	// Displayed after the names of players that are part of this team.
	TeamSuffix ns.TextComponent
	// Identifiers for the entities in this team. For players, this is their username; for other entities, it is their UUID.
	Entities ns.PrefixedArray[ns.String]
}

// S2CSetScore represents "Update Score".
//
// > This is sent to the client when it should update a scoreboard item.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Score
var S2CSetScore = jp.NewPacket(jp.StatePlay, jp.S2C, 0x67)

type S2CSetScoreData struct {
	// The entity whose score this is. For players, this is their username; for other entities, it is their UUID.
	EntityName ns.String
	// The name of the objective the score belongs to.
	ObjectiveName ns.String
	// The score to be displayed next to the entry.
	Value ns.VarInt
	// The custom display name.
	DisplayName ns.PrefixedOptional[ns.TextComponent]
	// Determines how the score number should be formatted.
	NumberFormat ns.PrefixedOptional[ns.VarInt]
	// Show nothing.
	StylingContent ns.ByteArray // TODO
}

// S2CSetSimulationDistance represents "Set Simulation Distance".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Simulation_Distance
var S2CSetSimulationDistance = jp.NewPacket(jp.StatePlay, jp.S2C, 0x68)

type S2CSetSimulationDistanceData struct {
	// The distance that the client will process specific things, such as entities.
	SimulationDistance ns.VarInt
}

// S2CSetSubtitleText represents "Set Subtitle Text".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Subtitle_Text
var S2CSetSubtitleText = jp.NewPacket(jp.StatePlay, jp.S2C, 0x69)

type S2CSetSubtitleTextData struct {
	//
	SubtitleText ns.TextComponent
}

// S2CSetTime represents "Update Time".
//
// > Time is based on ticks, where 20 ticks happen every second. There are 24000 ticks in a day, making Minecraft days exactly 20 minutes long.
// >
// > The time of day is based on the timestamp modulo 24000. 0 is sunrise, 6000 is noon, 12000 is sunset, and 18000 is midnight.
// >
// > The default SMP server increments the time by 20 every second.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Time
var S2CSetTime = jp.NewPacket(jp.StatePlay, jp.S2C, 0x6A)

type S2CSetTimeData struct {
	// In ticks; not changed by server commands.
	WorldAge ns.Long
	// The world (or region) time, in ticks.
	TimeOfDay ns.Long
	// If true, the client should automatically advance the time of day according to its ticking rate.
	TimeOfDayIncreasing ns.Boolean
}

// S2CSetTitleText represents "Set Title Text".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Title_Text
var S2CSetTitleText = jp.NewPacket(jp.StatePlay, jp.S2C, 0x6B)

type S2CSetTitleTextData struct {
	//
	TitleText ns.TextComponent
}

// S2CSetTitlesAnimation represents "Set Title Animation Times".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Title_Animation_Times
var S2CSetTitlesAnimation = jp.NewPacket(jp.StatePlay, jp.S2C, 0x6C)

type S2CSetTitlesAnimationData struct {
	// Ticks to spend fading in.
	FadeIn ns.Int
	// Ticks to keep the title displayed.
	Stay ns.Int
	// Ticks to spend fading out, not when to start fading out.
	FadeOut ns.Int
}

// S2CSoundEntity represents "Entity Sound Effect".
//
// > Plays a sound effect from an entity, either by hardcoded ID or Identifier. Sound IDs and names can be found here .
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Entity_Sound_Effect
var S2CSoundEntity = jp.NewPacket(jp.StatePlay, jp.S2C, 0x6D)

type S2CSoundEntityData struct {
	// ID in the minecraft:sound_event registry, or an inline definition.
	SoundEvent ns.Or[ns.Identifier, ns.SoundEvent]
	// The category that this sound will be played from ( current categories ).
	SoundCategory ns.VarInt
	//
	EntityId ns.VarInt
	// 1.0 is 100%, capped between 0.0 and 1.0 by vanilla clients.
	Volume ns.Float
	// Float between 0.5 and 2.0 by vanilla clients.
	Pitch ns.Float
	// Seed used to pick sound variant.
	Seed ns.Long
}

// S2CSound represents "Sound Effect".
//
// > Plays a sound effect at the given location, either by hardcoded ID or Identifier. Sound IDs and names can be found here .
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Sound_Effect
var S2CSound = jp.NewPacket(jp.StatePlay, jp.S2C, 0x6E)

type S2CSoundData struct {
	// ID in the minecraft:sound_event registry, or an inline definition.
	SoundEvent ns.Or[ns.Identifier, ns.SoundEvent]
	// The category that this sound will be played from ( current categories ).
	SoundCategory ns.VarInt
	// Effect X multiplied by 8 ( fixed-point number with only 3 bits dedicated to the fractional part).
	EffectPositionX ns.Int
	// Effect Y multiplied by 8 ( fixed-point number with only 3 bits dedicated to the fractional part).
	EffectPositionY ns.Int
	// Effect Z multiplied by 8 ( fixed-point number with only 3 bits dedicated to the fractional part).
	EffectPositionZ ns.Int
	// 1.0 is 100%, capped between 0.0 and 1.0 by vanilla clients.
	Volume ns.Float
	// Float between 0.5 and 2.0 by vanilla clients.
	Pitch ns.Float
	// Seed used to pick sound variant.
	Seed ns.Long
}

// S2CStartConfiguration represents "Start Configuration".
//
// > Sent during gameplay in order to redo the configuration process. The client must respond with Acknowledge Configuration for the process to start.
// >
// > This packet switches the connection state to configuration .
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Start_Configuration
var S2CStartConfiguration = jp.NewPacket(jp.StatePlay, jp.S2C, 0x6F)

type S2CStartConfigurationData struct {
	// No fields
}

// S2CStopSound represents "Stop Sound".
//
// > Categories:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Stop_Sound
var S2CStopSound = jp.NewPacket(jp.StatePlay, jp.S2C, 0x70)

type S2CStopSoundData struct {
	// Controls which fields are present.
	Flags ns.Byte
	// Only if flags is 3 or 1 (bit mask 0x1). See below. If not present, then sounds from all sources are cleared.
	Source ns.Optional[ns.VarInt]
	// Only if flags is 2 or 3 (bit mask 0x2). A sound effect name, see Custom Sound Effect . If not present, then all sounds are cleared.
	Sound ns.Optional[ns.Identifier]
}

// S2CStoreCookiePlay represents "Store Cookie (play)".
//
// > Stores some arbitrary data on the client, which persists between server transfers. The vanilla client only accepts cookies of up to 5 kiB in size.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Store_Cookie_(Play)
var S2CStoreCookiePlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x71)

type S2CStoreCookiePlayData struct {
	// The identifier of the cookie.
	Key ns.Identifier
	// The data of the cookie.
	Payload ns.PrefixedByteArray
}

// S2CSystemChat represents "System Chat Message".
//
// > Sends the client a raw system message.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#System_Chat_Message
var S2CSystemChat = jp.NewPacket(jp.StatePlay, jp.S2C, 0x72)

type S2CSystemChatData struct {
	// Limited to 262144 bytes.
	Content ns.TextComponent
	// Whether the message is an actionbar or chat message. See also #Set Action Bar Text .
	Overlay ns.Boolean
}

// S2CTabList represents "Set Tab List Header And Footer".
//
// > This packet may be used by custom servers to display additional information above/below the player list. It is never sent by the vanilla server.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Tab_List_Header_And_Footer
var S2CTabList = jp.NewPacket(jp.StatePlay, jp.S2C, 0x73)

type S2CTabListData struct {
	// To remove the header, send a empty text component: {"text":""} .
	Header ns.TextComponent
	// To remove the footer, send a empty text component: {"text":""} .
	Footer ns.TextComponent
}

// S2CTagQuery represents "Tag Query Response".
//
// > Sent in response to Query Block Entity Tag or Query Entity Tag .
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Tag_Query_Response
var S2CTagQuery = jp.NewPacket(jp.StatePlay, jp.S2C, 0x74)

type S2CTagQueryData struct {
	// Can be compared to the one sent in the original query packet.
	TransactionId ns.VarInt
	// The NBT of the block or entity. May be a TAG_END (0) in which case no NBT is present.
	Nbt ns.NBT
}

// S2CTakeItemEntity represents "Pickup Item".
//
// > Sent by the server when someone picks up an item lying on the ground — its sole purpose appears to be the animation of the item flying towards you. It doesn't destroy the entity in the client memory, and it
// > doesn't add it to your inventory. The server only checks for items to be picked up after each Set Player Position (and Set Player Position And Rotation ) packet sent by the client. The collector entity can be any entity; it does not have to be a player. The collected entity also can
// > be any entity, but the vanilla server only uses this for items, experience orbs, and the different varieties of arrows.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Pickup_Item
var S2CTakeItemEntity = jp.NewPacket(jp.StatePlay, jp.S2C, 0x75)

type S2CTakeItemEntityData struct {
	//
	CollectedEntityId ns.VarInt
	//
	CollectorEntityId ns.VarInt
	// Seems to be 1 for XP orbs, otherwise the number of items in the stack.
	PickupItemCount ns.VarInt
}

// S2CTeleportEntity represents "Synchronize Vehicle Position".
//
// > Teleports the entity on the client without changing the reference point of movement deltas in future Update Entity Position packets. Seems to be used to make relative
// > adjustments to vehicle positions; more information needed.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Synchronize_Vehicle_Position
var S2CTeleportEntity = jp.NewPacket(jp.StatePlay, jp.S2C, 0x76)

type S2CTeleportEntityData struct {
	//
	EntityId ns.VarInt
	//
	X ns.Double
	//
	Y ns.Double
	//
	Z ns.Double
	//
	VelocityX ns.Double
	//
	VelocityY ns.Double
	//
	VelocityZ ns.Double
	// Rotation on the Y axis, in degrees.
	Yaw ns.Float
	// Rotation on the Y axis, in degrees.
	Pitch ns.Float
	//
	Flags ns.TeleportFlags
	//
	OnGround ns.Boolean
}

// S2CTestInstanceBlockStatus represents "Test Instance Block Status".
//
// > Updates the status of the currently open Test Instance Block screen, if any.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Test_Instance_Block_Status
var S2CTestInstanceBlockStatus = jp.NewPacket(jp.StatePlay, jp.S2C, 0x77)

type S2CTestInstanceBlockStatusData struct {
	//
	Status ns.TextComponent
	//
	HasSize ns.Boolean
	// Only present if Has Size is true.
	SizeX ns.Optional[ns.Double]
	// Only present if Has Size is true.
	SizeY ns.Optional[ns.Double]
	// Only present if Has Size is true.
	SizeZ ns.Optional[ns.Double]
}

// S2CTickingState represents "Set Ticking State".
//
// > Used to adjust the ticking rate of the client, and whether it's frozen.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Set_Ticking_State
var S2CTickingState = jp.NewPacket(jp.StatePlay, jp.S2C, 0x78)

type S2CTickingStateData struct {
	//
	TickRate ns.Float
	//
	IsFrozen ns.Boolean
}

// S2CTickingStep represents "Step Tick".
//
// > Advances the client processing by the specified number of ticks. Has no effect unless client ticking is frozen.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Step_Tick
var S2CTickingStep = jp.NewPacket(jp.StatePlay, jp.S2C, 0x79)

type S2CTickingStepData struct {
	//
	TickSteps ns.VarInt
}

// S2CTransferPlay represents "Transfer (play)".
//
// > Notifies the client that it should transfer to the given server. Cookies previously stored are preserved between server transfers.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Transfer_(Play)
var S2CTransferPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x7A)

type S2CTransferPlayData struct {
	// The hostname or IP of the server.
	Host ns.String
	// The port of the server.
	Port ns.VarInt
}

// S2CUpdateAdvancements represents "Update Advancements".
//
// > Advancement structure:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Advancements
var S2CUpdateAdvancements = jp.NewPacket(jp.StatePlay, jp.S2C, 0x7B)

type S2CUpdateAdvancementsData struct {
	// Whether to reset/clear the current advancements.
	ResetClear ns.Boolean
	// See below
	Value ns.ByteArray // TODO: Advancement
	// The identifiers of the advancements that should be removed.
	Identifiers ns.PrefixedArray[ns.Identifier]
	// Value ns.ByteArray // TODO: AdvancementProgress
	ShowAdvancements ns.Boolean
}

// S2CUpdateAttributes represents "Update Attributes".
//
// > Sets attributes on the given entity.
// >
// > Modifier Data structure:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Attributes
var S2CUpdateAttributes = jp.NewPacket(jp.StatePlay, jp.S2C, 0x7C)

type S2CUpdateAttributesData struct {
	//
	EntityId ns.VarInt
	// See below.
	Value ns.Double
	// See Attribute#Modifiers . Modifier Data defined below.
	Modifiers ns.PrefixedArray[ns.ByteArray] // TODO: ModifierData
}

// S2CUpdateMobEffect represents "Entity Effect".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Entity_Effect
var S2CUpdateMobEffect = jp.NewPacket(jp.StatePlay, jp.S2C, 0x7D)

type S2CUpdateMobEffectData struct {
	//
	EntityId ns.VarInt
	// See this table .
	EffectId ns.VarInt
	// Vanilla client displays effect level as Amplifier + 1.
	Amplifier ns.VarInt
	// Duration in ticks. (-1 for infinite)
	Duration ns.VarInt
	// Bit field, see below.
	Flags ns.Byte
}

// S2CUpdateRecipes represents "Update Recipes".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Recipes
var S2CUpdateRecipes = jp.NewPacket(jp.StatePlay, jp.S2C, 0x7E)

type S2CUpdateRecipesData struct {
	// Prefixed Array
	PropertySets ns.PrefixedArray[struct {
		PropertySetId ns.Identifier
		Items         ns.PrefixedArray[ns.VarInt]
	}]
	//
	SlotDisplay ns.SlotDisplay
}

// S2CUpdateTagsPlay represents "Update Tags (play)".
//
// > A tag looks like this:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Tags_(Play)
var S2CUpdateTagsPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x7F)

type S2CUpdateTagsPlayData struct {
	// Prefixed Array
	RegistryToTagsMap ns.PrefixedArray[struct {
		Registry ns.Identifier
		Tags     ns.Array[Tag]
	}]
}

type Tag struct {
	Identifier ns.Identifier
	Entries    ns.PrefixedArray[ns.VarInt]
}

// S2CProjectilePower represents "Projectile Power".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Projectile_Power
var S2CProjectilePower = jp.NewPacket(jp.StatePlay, jp.S2C, 0x80)

type S2CProjectilePowerData struct {
	//
	EntityId ns.VarInt
	//
	Power ns.Double
}

// S2CCustomReportDetails represents "Custom Report Details".
//
// > Contains a list of key-value text entries that are included in any crash or disconnection report generated during connection to the server.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Custom_Report_Details
var S2CCustomReportDetails = jp.NewPacket(jp.StatePlay, jp.S2C, 0x81)

type S2CCustomReportDetailsData struct {
	// Prefixed Array (32)
	Details ns.PrefixedArray[struct {
		Title       ns.String
		Description ns.String
	}]
}

// S2CServerLinks represents "Server Links".
//
// > This packet contains a list of links that the vanilla client will display in the menu available from the pause menu. Link labels can be built-in or custom (i.e., any text).
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Server_Links
var S2CServerLinks = jp.NewPacket(jp.StatePlay, jp.S2C, 0x82)

type S2CServerLinksData struct {
	// Prefixed Array
	Links ns.PrefixedArray[struct {
		// Determines if the following label is built-in (from enum) or custom (text component).
		IsBuiltin ns.Boolean
		// See `ServerLink*` enums.
		Label ns.Or[ns.VarInt, ns.TextComponent]
		// Valid URL.
		Url ns.String
	}]
}

const (
	// Displayed on connection error screen; included as a comment in the disconnection report.
	ServerLinkBugReport ns.VarInt = iota
	ServerLinkCommunityGuidelines
	ServerLinkSupport
	ServerLinkStatus
	ServerLinkFeedback
	ServerLinkCommunity
	ServerLinkWebsite
	ServerLinkForums
	ServerLinkNews
	ServerLinkAnnouncements
)

// S2CWaypoint represents "Waypoint".
//
// > Adds, removes, or updates an entry that will be tracked on the player locator bar.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Waypoint
var S2CWaypoint = jp.NewPacket(jp.StatePlay, jp.S2C, 0x83)

type S2CWaypointData struct {
	// 0: track, 1: untrack, 2: update.
	Operation ns.VarInt
	// Something that uniquely identifies this specific waypoint.
	Identifier ns.Or[ns.UUID, ns.String]
	// Path to the waypoint style JSON: assets/<namespace>/waypoint_style/<value>.json.
	IconStyle ns.Identifier
	// Defines how the following field is read.
	WaypointType ns.VarInt
}

// S2CClearDialogPlay represents "Clear Dialog (play)".
//
// > If we're currently in a dialog screen, then this removes the current screen and switches back to the previous one.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clear_Dialog_(Play)
var S2CClearDialogPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x84)

type S2CClearDialogPlayData struct {
	// No fields
}

// S2CShowDialogPlay represents "Show Dialog (play)".
//
// > Show a custom dialog screen to the client.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Show_Dialog_(Play)
var S2CShowDialogPlay = jp.NewPacket(jp.StatePlay, jp.S2C, 0x85)

type S2CShowDialogPlayData struct {
	// ID in the minecraft:dialog registry, or an inline definition as described at Registry_data#Dialog .
	Dialog ns.Or[ns.Identifier, ns.NBT]
}

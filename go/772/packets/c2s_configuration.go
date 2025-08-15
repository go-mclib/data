package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/net_structures"
)

// C2SClientInformationConfiguration represents "Client Information (configuration)".
//
// > Sent when the player connects, or when settings are changed.
// >
// > Displayed Skin Parts flags:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Client_Information_(Configuration)
var C2SClientInformationConfiguration = jp.NewPacket(jp.StateConfiguration, jp.C2S, 0x00)

type C2SClientInformationConfigurationData struct {
	// e.g. en_GB .
	Locale ns.String
	// Client-side render distance, in chunks.
	ViewDistance ns.Byte
	// 0: enabled, 1: commands only, 2: hidden. See Chat#Client chat mode for more information.
	ChatMode ns.VarInt
	// “Colors” multiplayer setting. The vanilla server stores this value but does nothing with it (see MC-64867 ). Third-party servers such as Hypixel disable all coloring in chat and system messages when it is false.
	ChatColors ns.Boolean
	// Bit mask, see below.
	DisplayedSkinParts ns.UnsignedByte
	// 0: Left, 1: Right.
	MainHand ns.VarInt
	// Enables filtering of text on signs and written book titles. The vanilla client sets this according to the profanityFilterPreferences.profanityFilterOn account attribute indicated by the /player/attributes Mojang API endpoint . In offline mode it is always false.
	EnableTextFiltering ns.Boolean
	// Servers usually list online players, this option should let you not show up in that list.
	AllowServerListings ns.Boolean
	// 0: all, 1: decreased, 2: minimal
	ParticleStatus ns.VarInt
}

// C2SCookieResponseConfiguration represents "Cookie Response (configuration)".
//
// > Response to a Cookie Request (configuration) from the server. The vanilla server only accepts responses of up to 5 kiB in size.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Cookie_Response_(Configuration)
var C2SCookieResponseConfiguration = jp.NewPacket(jp.StateConfiguration, jp.C2S, 0x01)

type C2SCookieResponseConfigurationData struct {
	// The identifier of the cookie.
	Key ns.Identifier
	// The data of the cookie.
	Payload ns.PrefixedOptional[ns.PrefixedArray[ns.Byte]]
}

// C2SCustomPayloadConfiguration represents "Serverbound Plugin Message (configuration)".
//
// > Mods and plugins can use this to send their data. Minecraft itself uses some plugin channels . These internal channels are in the minecraft namespace.
// >
// > More documentation on this: https://dinnerbone.com/blog/2012/01/13/minecraft-plugin-channels-messaging/
// >
// > Note that the length of Data is known only from the packet length, since the packet has no length field of any kind.
// >
// > In vanilla server, the maximum data length is 32767 bytes.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Serverbound_Plugin_Message_(Configuration)
var C2SCustomPayloadConfiguration = jp.NewPacket(jp.StateConfiguration, jp.C2S, 0x02)

type C2SCustomPayloadConfigurationData struct {
	// Name of the plugin channel used to send the data.
	Channel ns.Identifier
	// Any data, depending on the channel. minecraft: channels are documented here . The length of this array must be inferred from the packet length.
	Data ns.Array[ns.Byte]
}

// C2SFinishConfiguration represents "Acknowledge Finish Configuration".
//
// > Sent by the client to notify the server that the configuration process has finished. It is sent in response to the server's Finish Configuration .
// >
// > This packet switches the connection state to play .
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Acknowledge_Finish_Configuration
var C2SFinishConfiguration = jp.NewPacket(jp.StateConfiguration, jp.C2S, 0x03)

type C2SFinishConfigurationData struct {
	// No fields
}

// C2SKeepAliveConfiguration represents "Serverbound Keep Alive (configuration)".
//
// > The server will frequently send out a keep-alive (see Clientbound Keep Alive ), each containing a random ID. The client must respond with the same packet.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Serverbound_Keep_Alive_(Configuration)
var C2SKeepAliveConfiguration = jp.NewPacket(jp.StateConfiguration, jp.C2S, 0x04)

type C2SKeepAliveConfigurationData struct {
	//
	KeepAliveId ns.Long
}

// C2SPongConfiguration represents "Pong (configuration)".
//
// > Response to the clientbound packet ( Ping ) with the same id.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Pong_(Configuration)
var C2SPongConfiguration = jp.NewPacket(jp.StateConfiguration, jp.C2S, 0x05)

type C2SPongConfigurationData struct {
	//
	Id ns.Int
}

// C2SResourcePackConfiguration represents "Resource Pack Response (configuration)".
//
// > Result can be one of the following values:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Resource_Pack_Response_(Configuration)
var C2SResourcePackConfiguration = jp.NewPacket(jp.StateConfiguration, jp.C2S, 0x06)

type C2SResourcePackConfigurationData struct {
	// The unique identifier of the resource pack received in the Add Resource Pack (configuration) request.
	Uuid ns.UUID
	// Result ID (see below).
	Result ns.VarInt
}

// C2SSelectKnownPacks represents "Serverbound Known Packs".
//
// > Informs the server of which data packs are present on the client. The client sends this in response to Clientbound Known Packs .
// >
// > If the client specifies a pack in this packet, the server should omit its contained data from the Registry Data packet.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Serverbound_Known_Packs
var C2SSelectKnownPacks = jp.NewPacket(jp.StateConfiguration, jp.C2S, 0x07)

type C2SSelectKnownPacksData struct {
	// Prefixed Array
	KnownPacks ns.PrefixedArray[ns.String]
}

// C2SCustomClickActionConfiguration represents "Custom Click Action (configuration)".
//
// > Sent when the client clicks a Text Component with the minecraft:custom click action. This is meant as an alternative to running a command, but will not have any
// > effect on vanilla servers.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Custom_Click_Action_(Configuration)
var C2SCustomClickActionConfiguration = jp.NewPacket(jp.StateConfiguration, jp.C2S, 0x08)

type C2SCustomClickActionConfigurationData struct {
	// The identifier for the click action.
	Id ns.Identifier
	// The data to send with the click action. May be a TAG_END (0).
	Payload ns.NBT
}

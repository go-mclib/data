package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/net_structures"
)

// S2CCookieRequestConfiguration represents "Cookie Request (configuration)".
//
// > Requests a cookie that was previously stored.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Cookie_Request_(Configuration)
var S2CCookieRequestConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x00)

type S2CCookieRequestConfigurationData struct {
	// The identifier of the cookie.
	Key ns.Identifier
}

// S2CCustomPayloadConfiguration represents "Clientbound Plugin Message (configuration)".
//
// > Mods and plugins can use this to send their data. Minecraft itself uses several plugin channels . These internal channels are in the minecraft namespace.
// >
// > More information on how it works on Dinnerbone's blog . More documentation about internal and popular registered channels are here .
// >
// > In vanilla clients, the maximum data length is 1048576 bytes.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clientbound_Plugin_Message_(Configuration)
var S2CCustomPayloadConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x01)

type S2CCustomPayloadConfigurationData struct {
	// Name of the plugin channel used to send the data.
	Channel ns.Identifier
	// Any data. The length of this array must be inferred from the packet length.
	Data ns.ByteArray
}

// S2CDisconnectConfiguration represents "Disconnect (configuration)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Disconnect_(Configuration)
var S2CDisconnectConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x02)

type S2CDisconnectConfigurationData struct {
	// The reason why the player was disconnected.
	Reason ns.TextComponent
}

// S2CFinishConfiguration represents "Finish Configuration".
//
// > Sent by the server to notify the client that the configuration process has finished. The client answers with Acknowledge Finish Configuration whenever it is ready
// > to continue.
// >
// > This packet switches the connection state to play .
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Finish_Configuration
var S2CFinishConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x03)

type S2CFinishConfigurationData struct {
	// No fields
}

// S2CKeepAliveConfiguration represents "Clientbound Keep Alive (configuration)".
//
// > The server will frequently send out a keep-alive, each containing a random ID. The client must respond with the same payload (see Serverbound Keep Alive ). If
// > the client does not respond to a Keep Alive packet within 15 seconds after it was sent, the server kicks the client. Vice versa, if the server does not send any keep-alives for 20 seconds, the client will
// > disconnect and yields a "Timed out" exception.
// >
// > The vanilla server uses a system-dependent time in milliseconds to generate the keep alive ID value.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clientbound_Keep_Alive_(Configuration)
var S2CKeepAliveConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x04)

type S2CKeepAliveConfigurationData struct {
	//
	KeepAliveId ns.Long
}

// S2CPingConfiguration represents "Ping (configuration)".
//
// > Packet is not used by the vanilla server. When sent to the client, client responds with a Pong packet with the same id.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Ping_(Configuration)
var S2CPingConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x05)

type S2CPingConfigurationData struct {
	//
	Id ns.Int
}

// S2CResetChat represents "Reset Chat".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Reset_Chat
var S2CResetChat = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x06)

type S2CResetChatData struct {
	// No fields
}

// S2CRegistryData represents "Registry Data".
//
// > Represents certain registries that are sent from the server and are applied on the client.
// >
// > See Registry Data for details.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Registry_Data
var S2CRegistryData = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x07)

type S2CRegistryDataData struct {
	//
	RegistryId ns.Identifier
	// Entry data.
	Data ns.PrefixedOptional[ns.NBT]
}

// S2CResourcePackPopConfiguration represents "Remove Resource Pack (configuration)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Remove_Resource_Pack_(Configuration)
var S2CResourcePackPopConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x08)

type S2CResourcePackPopConfigurationData struct {
	// The UUID of the resource pack to be removed. If not present every resource pack will be removed.
	Uuid ns.PrefixedOptional[ns.UUID]
}

// S2CResourcePackPushConfiguration represents "Add Resource Pack (configuration)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Add_Resource_Pack_(Configuration)
var S2CResourcePackPushConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x09)

type S2CResourcePackPushConfigurationData struct {
	// The unique identifier of the resource pack.
	Uuid ns.UUID
	// The URL to the resource pack.
	Url ns.String
	// A 40 character hexadecimal, case-insensitive SHA-1 hash of the resource pack file. If it's not a 40 character hexadecimal string, the client will not use it for hash verification and likely waste bandwidth.
	Hash ns.String
	// The vanilla client will be forced to use the resource pack from the server. If they decline they will be kicked from the server.
	Forced ns.Boolean
	// This is shown in the prompt making the client accept or decline the resource pack (only if present).
	PromptMessage ns.PrefixedOptional[ns.TextComponent]
}

// S2CStoreCookieConfiguration represents "Store Cookie (configuration)".
//
// > Stores some arbitrary data on the client, which persists between server transfers. The vanilla client only accepts cookies of up to 5 kiB in size.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Store_Cookie_(Configuration)
var S2CStoreCookieConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x0A)

type S2CStoreCookieConfigurationData struct {
	// The identifier of the cookie.
	Key ns.Identifier
	// The data of the cookie.
	Payload ns.PrefixedByteArray
}

// S2CTransferConfiguration represents "Transfer (configuration)".
//
// > Notifies the client that it should transfer to the given server. Cookies previously stored are preserved between server transfers.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Transfer_(Configuration)
var S2CTransferConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x0B)

type S2CTransferConfigurationData struct {
	// The hostname or IP of the server.
	Host ns.String
	// The port of the server.
	Port ns.VarInt
}

// S2CUpdateEnabledFeatures represents "Feature Flags".
//
// > Used to enable and disable features, generally experimental ones, on the client.
// >
// > There is one special feature flag, which is in most versions:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Feature_Flags
var S2CUpdateEnabledFeatures = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x0C)

type S2CUpdateEnabledFeaturesData struct {
	//
	FeatureFlags ns.PrefixedArray[ns.Identifier]
}

// S2CUpdateTagsConfiguration represents "Update Tags (configuration)".
//
// > Tag arrays look like:
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Update_Tags_(Configuration)
var S2CUpdateTagsConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x0D)

type S2CUpdateTagsConfigurationData struct {
	// Prefixed Array
	ArrayOfTags ns.PrefixedArray[struct {
		Registry   ns.Identifier
		ArrayOfTag ns.Array[Tag]
	}]
}

// S2CSelectKnownPacks represents "Clientbound Known Packs".
//
// > Informs the client of which data packs are present on the server. The client is expected to respond with its own Serverbound Known Packs packet. The vanilla server does not
// > continue with Configuration until it receives a response.
// >
// > The vanilla client requires the minecraft:core pack with version 1.21.8 for a normal login sequence. This packet must be sent before the Registry Data packets.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clientbound_Known_Packs
var S2CSelectKnownPacks = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x0E)

type S2CSelectKnownPacksData struct {
	// Prefixed Array
	KnownPacks ns.PrefixedArray[struct {
		Namespace ns.String
		Id        ns.String
		Version   ns.String
	}]
}

// S2CCustomReportDetailsConfiguration represents "Custom Report Details (configuration)".
//
// > Contains a list of key-value text entries that are included in any crash or disconnection report generated during connection to the server.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Custom_Report_Details_(Configuration)
var S2CCustomReportDetailsConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x0F)

type S2CCustomReportDetailsConfigurationData struct {
	// Prefixed Array (32)
	Details ns.PrefixedArray[struct {
		Title       ns.String
		Description ns.String
	}]
}

// S2CServerLinksConfiguration represents "Server Links (configuration)".
//
// > This packet contains a list of links that the vanilla client will display in the menu available from the pause menu. Link labels can be built-in or custom (i.e., any text).
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Server_Links_(Configuration)
var S2CServerLinksConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x10)

type S2CServerLinksConfigurationData struct {
	// Prefixed Array
	Links ns.PrefixedArray[struct {
		Label ns.Or[ns.VarInt, ns.TextComponent]
		Url   ns.String
	}]
}

// S2CClearDialogConfiguration represents "Clear Dialog (configuration)".
//
// > If we're currently in a dialog screen, then this removes the current screen and switches back to the previous one.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Clear_Dialog_(Configuration)
var S2CClearDialogConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x11)

type S2CClearDialogConfigurationData struct {
	// No fields
}

// S2CShowDialogConfiguration represents "Show Dialog (configuration)".
//
// > Show a custom dialog screen to the client.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Show_Dialog_(Configuration)
var S2CShowDialogConfiguration = jp.NewPacket(jp.StateConfiguration, jp.S2C, 0x12)

type S2CShowDialogConfigurationData struct {
	// Inline definition as described at Registry_data#Dialog .
	Dialog ns.NBT
}

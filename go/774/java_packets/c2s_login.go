package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/net_structures"
)

// C2SHello represents "Login Start".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Login_Start
var C2SHello = jp.NewPacket(jp.StateLogin, jp.C2S, 0x00)

type C2SHelloData struct {
	// Player's Username.
	Name ns.String
	// The UUID of the player logging in. Unused by the vanilla server.
	PlayerUuid ns.UUID
}

// C2SKey represents "Encryption Response".
//
// > See protocol encryption for details.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Encryption_Response
var C2SKey = jp.NewPacket(jp.StateLogin, jp.C2S, 0x01)

type C2SKeyData struct {
	// Shared Secret value, encrypted with the server's public key.
	SharedSecret ns.PrefixedArray[ns.Byte]
	// Verify Token value, encrypted with the same public key as the shared secret.
	VerifyToken ns.PrefixedArray[ns.Byte]
}

// C2SCustomQueryAnswer represents "Login Plugin Response".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Login_Plugin_Response
var C2SCustomQueryAnswer = jp.NewPacket(jp.StateLogin, jp.C2S, 0x02)

type C2SCustomQueryAnswerData struct {
	// Should match ID from server.
	MessageId ns.VarInt
	// Any data, depending on the channel. Only present if the client understood the request. Typically this would be a sequence of fields using standard data types, but some unofficial channels have unusual formats. There is no length prefix that applies to all channel types, but the format specific to the channel may or may not include one or more length prefixes (e.g. for strings).
	Data ns.PrefixedOptional[ns.ByteArray]
}

// C2SLoginAcknowledged represents "Login Acknowledged".
//
// > Acknowledgement to the Login Success packet sent by the server.
// >
// > This packet switches the connection state to configuration .
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Login_Acknowledged
var C2SLoginAcknowledged = jp.NewPacket(jp.StateLogin, jp.C2S, 0x03)

type C2SLoginAcknowledgedData struct {
	// No fields
}

// C2SCookieResponseLogin represents "Cookie Response (login)".
//
// > Response to a Cookie Request (login) from the server. The vanilla server only accepts responses of up to 5 kiB in size.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Cookie_Response_(Login)
var C2SCookieResponseLogin = jp.NewPacket(jp.StateLogin, jp.C2S, 0x04)

type C2SCookieResponseLoginData struct {
	// The identifier of the cookie.
	Key ns.Identifier
	// The data of the cookie.
	Payload ns.PrefixedOptional[ns.PrefixedArray[ns.Byte]]
}

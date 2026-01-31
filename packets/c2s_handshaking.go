package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// C2SIntention represents "Handshake".
//
// > This packet causes the server to switch into the target state. It should be sent right after opening the TCP connection to prevent the server from disconnecting.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Handshake
var C2SIntention = jp.NewPacket(jp.StateHandshake, jp.C2S, 0x00)

type C2SIntentionData struct {
	// See protocol version numbers (currently 773 in Minecraft 1.21.10).
	ProtocolVersion ns.VarInt
	// Hostname or IP, e.g. localhost or 127.0.0.1, that was used to connect. The vanilla server does not use this information. This is the name obtained after SRV record resolution, except in 1.17 (and no older or newer version) and during server list ping ( MC-278651 ), where it is the host portion of the address specified by the user directly. In 1.17.1 and later if a literal IP address is specified by the user, reverse DNS lookup is attempted, and the result is used as the value of this field if successful.
	ServerAddress ns.String
	// Default is 25565. The vanilla server does not use this information.
	ServerPort ns.UnsignedShort
	// 1 for Status , 2 for Login , 3 for Transfer .
	Intent ns.VarInt
}

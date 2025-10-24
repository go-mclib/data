package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/net_structures"
)

// C2SIntention represents "Handshake".
//
// > This packet causes the server to switch into the target state. It should be sent right after opening the TCP connection to prevent the server from disconnecting.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Handshake
var C2SIntention = jp.NewPacket(jp.StateHandshake, jp.C2S, 0x00)

type C2SIntentionData struct {
	// See protocol version numbers (currently 772 in Minecraft 1.21.8).
	ProtocolVersion ns.VarInt
	// Hostname or IP, e.g. localhost or 127.0.0.1, that was used to connect. The vanilla server does not use this information. Note that SRV records are a simple redirect, e.g. if _minecraft._tcp.example.com points to mc.example.org, users connecting to example.com will provide example.org as the server address in addition to connecting to it.
	ServerAddress ns.String
	// Default is 25565. The vanilla server does not use this information.
	ServerPort ns.UnsignedShort
	// 1 for Status , 2 for Login , 3 for Transfer .
	Intent ns.VarInt
}

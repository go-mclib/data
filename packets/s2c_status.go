package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// S2CStatusResponse represents "Status Response".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Status_Response
var S2CStatusResponse = jp.NewPacket(jp.StateStatus, jp.S2C, 0x00)

type S2CStatusResponseData struct {
	// See Java Edition protocol/Server List Ping#Status Response ; as with all strings, this is prefixed by its length as a VarInt .
	JsonResponse ns.String
}

// S2CPongResponseStatus represents "Pong Response (status)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Pong_Response_(Status)
var S2CPongResponseStatus = jp.NewPacket(jp.StateStatus, jp.S2C, 0x01)

type S2CPongResponseStatusData struct {
	// Should match the one sent by the client.
	Timestamp ns.Long
}

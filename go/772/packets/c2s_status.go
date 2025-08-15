package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/net_structures"
)

// C2SStatusRequest represents "Status Request".
//
// > The status can only be requested once immediately after the handshake, before any ping. The server won't respond otherwise.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Status_Request
var C2SStatusRequest = jp.NewPacket(jp.StateStatus, jp.C2S, 0x00)

type C2SStatusRequestData struct {
	// No fields
}

// C2SPingRequestStatus represents "Ping Request (status)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Ping_Request_(Status)
var C2SPingRequestStatus = jp.NewPacket(jp.StateStatus, jp.C2S, 0x01)

type C2SPingRequestStatusData struct {
	// May be any number, but vanilla clients will always use the timestamp in milliseconds.
	Timestamp ns.Long
}

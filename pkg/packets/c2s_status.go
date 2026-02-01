package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

const (
	C2SStatusRequestID ns.VarInt = iota
	C2SPingRequestStatusID
)

// C2SStatusRequest represents "Status Request".
//
// The status can only be requested once, immediately after the handshake,
// before any ping. The server won't respond otherwise.
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Status_Request
type C2SStatusRequest struct{}

func (p *C2SStatusRequest) ID() ns.VarInt                { return C2SStatusRequestID }
func (p *C2SStatusRequest) State() jp.State              { return jp.StateStatus }
func (p *C2SStatusRequest) Bound() jp.Bound              { return jp.C2S }
func (p *C2SStatusRequest) Read(*ns.PacketBuffer) error  { return nil }
func (p *C2SStatusRequest) Write(*ns.PacketBuffer) error { return nil }

// C2SPingRequestStatus represents "Ping Request (status)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Ping_Request_(Status)
type C2SPingRequestStatus struct {
	// May be any number, but vanilla clients will always use the timestamp in milliseconds.
	Timestamp ns.Int64
}

func (p *C2SPingRequestStatus) ID() ns.VarInt   { return C2SPingRequestStatusID }
func (p *C2SPingRequestStatus) State() jp.State { return jp.StateStatus }
func (p *C2SPingRequestStatus) Bound() jp.Bound { return jp.C2S }

func (p *C2SPingRequestStatus) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Timestamp, err = buf.ReadInt64()
	return err
}

func (p *C2SPingRequestStatus) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt64(p.Timestamp)
}

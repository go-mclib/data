package packets

import (
	"github.com/go-mclib/data/pkg/data/packet_ids"
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// S2CStatusResponse represents "Status Response".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Status_Response
type S2CStatusResponse struct {
	// See Server List Ping; as with all strings, this is prefixed by its length as a VarInt.
	JsonResponse ns.String
}

func (p *S2CStatusResponse) ID() ns.VarInt   { return ns.VarInt(packet_ids.S2CStatusResponseID) }
func (p *S2CStatusResponse) State() jp.State { return jp.StateStatus }
func (p *S2CStatusResponse) Bound() jp.Bound { return jp.S2C }

func (p *S2CStatusResponse) Read(buf *ns.PacketBuffer) error {
	var err error
	p.JsonResponse, err = buf.ReadString(32767)
	return err
}

func (p *S2CStatusResponse) Write(buf *ns.PacketBuffer) error {
	return buf.WriteString(p.JsonResponse)
}

// S2CPongResponseStatus represents "Pong Response (status)".
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#Pong_Response_(Status)
type S2CPongResponseStatus struct {
	// Should match the one sent by the client.
	Timestamp ns.Int64
}

func (p *S2CPongResponseStatus) ID() ns.VarInt {
	return ns.VarInt(packet_ids.S2CPongResponseStatusID)
}
func (p *S2CPongResponseStatus) State() jp.State { return jp.StateStatus }
func (p *S2CPongResponseStatus) Bound() jp.Bound { return jp.S2C }

func (p *S2CPongResponseStatus) Read(buf *ns.PacketBuffer) error {
	var err error
	p.Timestamp, err = buf.ReadInt64()
	return err
}

func (p *S2CPongResponseStatus) Write(buf *ns.PacketBuffer) error {
	return buf.WriteInt64(p.Timestamp)
}

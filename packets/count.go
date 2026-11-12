package packets

import (
	"fmt"

	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// maxListCount bounds wire lists that vanilla decodes without a size limit
// (`ByteBufCodecs.list()`), so a hostile peer cannot request a huge allocation.
const maxListCount = 65536

// readCount reads a VarInt element count and rejects values outside 0..max.
func readCount(buf *ns.PacketBuffer, label string, max int) (int, error) {
	count, err := buf.ReadVarInt()
	if err != nil {
		return 0, err
	}
	if count < 0 || int(count) > max {
		return 0, fmt.Errorf("%s count %d outside 0..%d", label, count, max)
	}
	return int(count), nil
}

package items

import (
	"fmt"
	"strings"

	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// readHolderSet reads a vanilla HolderSet (an ns.IDSet) and resolves its IDs
// to registry names. A tag is returned as a single "#name" entry.
func readHolderSet(buf *ns.PacketBuffer, kind string, nameOf func(int32) string) ([]string, error) {
	var set ns.IDSet
	if err := set.Decode(buf); err != nil {
		return nil, fmt.Errorf("%s holder set: %w", kind, err)
	}
	if set.IsTag {
		return []string{"#" + string(set.TagName)}, nil
	}
	result := make([]string, 0, len(set.IDs))
	for _, id := range set.IDs {
		name := nameOf(int32(id))
		if name == "" {
			return nil, fmt.Errorf("unknown %s %d", kind, id)
		}
		result = append(result, name)
	}
	return result, nil
}

// writeHolderSet writes entries in the format read by readHolderSet.
func writeHolderSet(w *ns.PacketBuffer, kind string, entries []string, idOf func(string) int32) error {
	var set ns.IDSet
	if tag, ok := singleTagRef(entries); ok {
		set = ns.IDSet{IsTag: true, TagName: ns.Identifier(tag)}
	} else {
		set.IDs = make([]ns.VarInt, 0, len(entries))
		for _, name := range entries {
			id := idOf(name)
			if id < 0 {
				return fmt.Errorf("unknown %s %q", kind, name)
			}
			set.IDs = append(set.IDs, ns.VarInt(id))
		}
	}
	return set.Encode(w)
}

func readOptionalHolderSet(buf *ns.PacketBuffer, kind string, nameOf func(int32) string) ([]string, error) {
	present, err := buf.ReadBool()
	if err != nil || !present {
		return nil, err
	}
	return readHolderSet(buf, kind, nameOf)
}

// writeOptionalHolderSet writes a nil entries slice as absent.
func writeOptionalHolderSet(w *ns.PacketBuffer, kind string, entries []string, idOf func(string) int32) error {
	w.WriteBool(ns.Boolean(entries != nil))
	if entries == nil {
		return nil
	}
	return writeHolderSet(w, kind, entries, idOf)
}

// singleTagRef returns the tag name when entries is a single "#tag" reference.
func singleTagRef(entries []string) (string, bool) {
	if len(entries) != 1 {
		return "", false
	}
	return strings.CutPrefix(entries[0], "#")
}

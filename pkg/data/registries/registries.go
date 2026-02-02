package registries

// Registry represents a Minecraft registry.
type Registry struct {
	ProtocolID int32
	entries    map[string]int32
	byID       map[int32]string
}

// Get returns the protocol ID for an entry, or -1 if not found.
func (r *Registry) Get(id string) int32 {
	if v, ok := r.entries[id]; ok {
		return v
	}
	return -1
}

// ByID returns the entry name for a protocol ID, or empty string if not found.
func (r *Registry) ByID(protocolID int32) string {
	return r.byID[protocolID]
}

// Size returns the number of entries in the registry.
func (r *Registry) Size() int {
	return len(r.entries)
}

func newRegistry(protocolID int32, entries map[string]int32) *Registry {
	byID := make(map[int32]string, len(entries))
	for k, v := range entries {
		byID[v] = k
	}
	return &Registry{ProtocolID: protocolID, entries: entries, byID: byID}
}

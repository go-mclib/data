package registries

import (
	"iter"
	"maps"
)

// Registry represents a Minecraft registry.
type Registry struct {
	Identifier string // e.g., "minecraft:block"
	ProtocolID int32  // ID in minecraft:root, or -1 for data-driven (synchronized) registries
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

// HolderID returns the network holder ID of an entry reference: its protocol
// ID + 1, since holder ID 0 marks an inline (direct) value. It returns -1 if
// the entry is not found.
func (r *Registry) HolderID(id string) int32 {
	if v, ok := r.entries[id]; ok {
		return v + 1
	}
	return -1
}

// HolderName returns the entry name referenced by a network holder ID (see
// HolderID), or empty string for 0 (an inline value) or an unknown ID.
func (r *Registry) HolderName(holderID int32) string {
	if holderID <= 0 {
		return ""
	}
	return r.byID[holderID-1]
}

// Size returns the number of entries in the registry.
func (r *Registry) Size() int {
	return len(r.entries)
}

// Entries iterates over all entries in the registry.
func (r *Registry) Entries() iter.Seq2[string, int32] {
	return maps.All(r.entries)
}

func newRegistry(identifier string, protocolID int32, entries map[string]int32) *Registry {
	byID := make(map[int32]string, len(entries))
	for k, v := range entries {
		byID[v] = k
	}
	return &Registry{Identifier: identifier, ProtocolID: protocolID, entries: entries, byID: byID}
}

// newSynchronizedRegistry builds a data-driven registry whose protocol IDs are
// the indices of names.
func newSynchronizedRegistry(identifier string, names []string) *Registry {
	entries := make(map[string]int32, len(names))
	for id, name := range names {
		entries[name] = int32(id)
	}
	return newRegistry(identifier, -1, entries)
}

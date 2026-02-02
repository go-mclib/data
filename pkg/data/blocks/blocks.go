package blocks

// BlockID returns the protocol ID for a block string identifier, or -1 if not found.
func BlockID(id string) int32 {
	if v, ok := blockByName[id]; ok {
		return v
	}
	return -1
}

// BlockName returns the string identifier for a block protocol ID, or empty string if not found.
func BlockName(protocolID int32) string {
	return blockByID[protocolID]
}

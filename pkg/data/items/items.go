package items

// ItemID returns the protocol ID for an item string identifier, or -1 if not found.
func ItemID(id string) int32 {
	if v, ok := itemByName[id]; ok {
		return v
	}
	return -1
}

// ItemName returns the string identifier for an item protocol ID, or empty string if not found.
func ItemName(protocolID int32) string {
	return itemByID[protocolID]
}

// DefaultComponents returns the default components for an item, or nil if not found.
func DefaultComponents(itemID int32) *Components {
	return defaultComponents[itemID]
}

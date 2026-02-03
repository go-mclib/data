package entities

// EntityTypeID returns the protocol ID for an entity type string identifier, or -1 if not found.
func EntityTypeID(name string) int32 {
	if v, ok := entityByName[name]; ok {
		return v
	}
	return -1
}

// EntityTypeName returns the string identifier for an entity type protocol ID, or empty string if not found.
func EntityTypeName(protocolID int32) string {
	return entityByID[protocolID]
}

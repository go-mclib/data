package main

// JSON structures for server reports

type RegistryJSON struct {
	ProtocolID int32                        `json:"protocol_id"`
	Default    string                       `json:"default,omitempty"`
	Entries    map[string]RegistryEntryJSON `json:"entries"`
}

type RegistryEntryJSON struct {
	ProtocolID int32 `json:"protocol_id"`
}

type BlockJSON struct {
	Definition BlockDefinitionJSON `json:"definition"`
	Properties map[string][]string `json:"properties"`
	States     []BlockStateJSON    `json:"states"`
}

type BlockDefinitionJSON struct {
	Type string `json:"type"`
}

type BlockStateJSON struct {
	ID         int32             `json:"id"`
	Default    bool              `json:"default,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

type ItemJSON struct {
	Components map[string]any `json:"components"`
}

// PacketsJSON structure: phase -> bound -> packet_name -> {protocol_id}
type PacketsJSON map[string]map[string]map[string]PacketEntryJSON

type PacketEntryJSON struct {
	ProtocolID int32 `json:"protocol_id"`
}

// ComponentMetadata defines wire format info for a component.
type ComponentMetadata struct {
	WireType    string `json:"wireType"`
	Passthrough bool   `json:"passthrough,omitempty"`
}

// ComponentMetadataFile is the structure of component_metadata.include.json.
type ComponentMetadataFile struct {
	Components     map[string]ComponentMetadata `json:"components"`
	EntityVariants []string                     `json:"entityVariants"`
}

package packets

import (
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// StatusResponse is the JSON document carried by [S2CStatusResponse]
// (vanilla ServerStatus). Description is a text component: it decodes from
// both the plain-string and the component-object form.
type StatusResponse struct {
	Description        ns.TextComponent `json:"description"`
	Players            StatusPlayers    `json:"players"`
	Version            StatusVersion    `json:"version"`
	Favicon            string           `json:"favicon,omitempty"`
	EnforcesSecureChat bool             `json:"enforcesSecureChat,omitempty"`
}

// StatusPlayers is the player count and sample of a [StatusResponse].
type StatusPlayers struct {
	Max    int            `json:"max"`
	Online int            `json:"online"`
	Sample []StatusPlayer `json:"sample,omitempty"`
}

// StatusPlayer is one sampled player of a [StatusResponse].
type StatusPlayer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// StatusVersion is the game version and protocol of a [StatusResponse].
type StatusVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

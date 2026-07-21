package misc

type ServerStatusResponse struct {
	Description        string              `json:"description"`
	Players            ServerStatusPlayers `json:"players"`
	Version            ServerStatusVersion `json:"version"`
	Favicon            string              `json:"favicon,omitempty"`
	EnforcesSecureChat bool                `json:"enforcesSecureChat,omitempty"`
}

type ServerStatusPlayers struct {
	Max    int                  `json:"max"`
	Online int                  `json:"online"`
	Sample []ServerStatusPlayer `json:"sample,omitempty"`
}

type ServerStatusPlayer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ServerStatusVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

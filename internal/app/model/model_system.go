package model

// System ...
type System struct {
	State              string `json:"state"`
	PatchLevel         string `json:"patch_level"`
	ThreatLevel        int    `json:"threat_level"`
	NetworkState       string `json:"network_state"`
	NetworkThreatLevel int    `json:"network_threat_level"`
}

// Validate ...
func (s *System) Validate() error {
	return nil
}

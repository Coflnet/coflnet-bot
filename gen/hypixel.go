package gen

// HypixelPlayerResponse / HypixelPlayerResponse this code is generated with json 2 go
type HypixelPlayerResponse struct {
	Player struct {
		ID          string `json:"_id,omitempty"`
		SocialMedia struct {
			Links struct {
				Discord string `json:"DISCORD,omitempty"`
			} `json:"links,omitempty"`
		} `json:"socialMedia,omitempty"`
	} `json:"player,omitempty"`
}

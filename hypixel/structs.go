package hypixel

type HypixelApiResponse struct {
	Success bool `json:"success"`
	Player  struct {
		ID          string `json:"_id"`
		SocialMedia struct {
			Links struct {
				Discord string `json:"DISCORD"`
			} `json:"links"`
		} `json:"socialMedia"`
	} `json:"player"`
}

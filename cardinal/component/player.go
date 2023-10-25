package component

type PlayerComponent struct {
	Nickname string `json:"nickname"`
}

func (PlayerComponent) Name() string {
	return "Player"
}

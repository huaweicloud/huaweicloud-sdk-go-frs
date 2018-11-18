package common

type VideoResult struct {
	Alive   bool     `json:"alive"`
	Actions []Action `json:"actions"`
	Picture string   `json:"picture"`
}

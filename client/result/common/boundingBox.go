package common

type BoundingBox struct {
	Width    int `json:"width"`
	Height   int `json:"height"`
	TopLeftX int `json:"top_left_x"`
	TopLeftY int `json:"top_left_y"`
}

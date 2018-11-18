package common

type DetectFace struct {
	SimpleFace
	Attributes *Attributes `json:"attributes,omitempty"`
	Landmark   *Landmark   `json:"landmark,omitempty"`
}

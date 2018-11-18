package common

type Attributes struct {
	HeadPose *HeadPose `json:"headpose,omitempty"`
	Gender   *string   `json:"gender,omitempty"`
	Age      *int      `json:"age,omitempty"`
	Dress    *Dress    `json:"dress,omitempty"`
	Smile    *string   `json:"smile,omitempty"`
}

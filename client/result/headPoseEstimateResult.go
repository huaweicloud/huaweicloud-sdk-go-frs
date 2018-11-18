package result

type HeadPoseEstimateResult struct {
	Yaw   float64 `json:"yaw"`
	Roll  float64 `json:"roll"`
	Pitch float64 `json:"pitch"`
}

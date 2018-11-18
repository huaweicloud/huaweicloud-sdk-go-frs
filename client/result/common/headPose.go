package common

type HeadPose struct {
	YawAngle   float64 `json:"yaw_angle"`
	RollAngle  float64 `json:"roll_angle"`
	PitchAngle float64 `json:"pitch_angle"`
}

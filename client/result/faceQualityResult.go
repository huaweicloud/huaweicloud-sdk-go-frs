package result

type FaceQualityResult struct {
	Blur BlurClassifyResult     `json:"blur"`
	Pose HeadPoseEstimateResult `json:"pose"`
}

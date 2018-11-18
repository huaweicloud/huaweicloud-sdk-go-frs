package common

type Landmark struct {
	EyesContour    PointList `json:"eyes_contour"`
	MouthContour   PointList `json:"mouth_contour"`
	FaceContour    PointList `json:"face_contour"`
	EyeBrowContour PointList `json:"eyebrow_contour"`
	NoseContour    PointList `json:"nose_contour"`
}

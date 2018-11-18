package result

type DeleteFaceResult struct {
	FaceNumber  int    `json:"face_number"`
	FaceSetId   string `json:"face_set_id"`
	FaceSetName string `json:"face_set_name"`
}

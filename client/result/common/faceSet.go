package common

type FaceSet struct {
	FaceNumber      *int                         `json:"face_number,omitempty"`
	FaceSetId       *string                      `json:"face_set_id,omitempty"`
	FaceSetName     *string                      `json:"face_set_name,omitempty"`
	CreateDate      *string                      `json:"create_date,omitempty"`
	FaceSetCapacity *int                         `json:"face_set_capacity,omitempty"`
	ExternalFields  map[string]map[string]string `json:"external_fields,omitempty"`
}

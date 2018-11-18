package common

type Face struct {
	SimpleFace
	ExternalImageId string                 `json:"external_image_id"`
	FaceId          string                 `json:"face_id"`
	ExternalFields  map[string]interface{} `json:"external_fields,omitempty"`
}

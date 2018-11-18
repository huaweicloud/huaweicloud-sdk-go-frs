package result

import "github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result/common"

type AddFaceResult struct {
	FaceSetId   string        `json:"face_set_id"`
	FaceSetName string        `json:"face_set_name"`
	Faces       []common.Face `json:"faces"`
}

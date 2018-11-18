package result

import "github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result/common"

type GetAllFaceSetsResult struct {
	FaceSetsInfo []common.FaceSet `json:"face_sets_info"`
}

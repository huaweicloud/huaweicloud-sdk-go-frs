package result

import "github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result/common"

type SearchFaceResult struct {
	Faces []common.ComplexFace `json:"faces"`
}

package result

import "github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result/common"

type DetectFaceResult struct {
	Faces []common.DetectFace `json:"faces"`
}

package result

import "github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result/common"

type CompareFaceResult struct {
	Similarity float64           `json:"similarity"`
	Image1Face common.SimpleFace `json:"image1_face"`
	Image2Face common.SimpleFace `json:"image2_face"`
}

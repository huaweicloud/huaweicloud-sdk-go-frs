package result

import "github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result/common"

type LiveDetectResult struct {
	VideoResult common.VideoResult `json:"video-result"`
	WarningList []common.Warning   `json:"warning-list"`
}

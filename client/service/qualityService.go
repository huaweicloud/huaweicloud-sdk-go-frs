package service

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-frs/access"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/common"
)

type QualityService struct {
	accessService *access.AccessService
	projectId     string
}

func newQualityService(accessService *access.AccessService, projectId string) *QualityService {
	return &QualityService{accessService, projectId}
}

func (qualityService *QualityService) faceQuality(image string, imageType int) (*result.FaceQualityResult, error) {
	uri := fmt.Sprintf(_FACE_QUALITY_URI, qualityService.projectId)
	jsonObj := common.NewJsonObj()
	switch imageType {
	case BASE64:
		jsonObj.Put("image_base64", image)
		break
	case OBSURL:
		jsonObj.Put("image_url", image)
		break
	}
	response, err := qualityService.accessService.Post(uri, map[string]string{}, jsonObj.GetString())
	if nil != err {
		return nil, err
	}
	faceQualityResult := &result.FaceQualityResult{}
	err = common.ResponseToObj(response, faceQualityResult)
	return faceQualityResult, err
}

func (qualityService *QualityService) FaceQualityByBase64(imageBase64 string) (*result.FaceQualityResult, error) {
	return qualityService.faceQuality(imageBase64, BASE64)
}

func (qualityService *QualityService) FaceQualityByFile(imagePath string) (*result.FaceQualityResult, error) {
	imageBase64, err := common.LoadImageToBase64(imagePath)
	if nil != err {
		return nil, err
	}
	return qualityService.FaceQualityByBase64(imageBase64)
}

func (qualityService *QualityService) FaceQualityByObsUrl(imageUrl string) (*result.FaceQualityResult, error) {
	return qualityService.faceQuality(imageUrl, OBSURL)
}

func (qualityService *QualityService) blurClassify(image string, imageType int) (*result.BlurClassifyResult, error) {
	uri := fmt.Sprintf(_BLUR_CLASSIFY_URI, qualityService.projectId)
	jsonObj := common.NewJsonObj()
	switch imageType {
	case BASE64:
		jsonObj.Put("image_base64", image)
		break
	case OBSURL:
		jsonObj.Put("image_url", image)
		break
	}
	response, err := qualityService.accessService.Post(uri, map[string]string{}, jsonObj.GetString())
	if nil != err {
		return nil, err
	}
	blurClassifyResult := &result.BlurClassifyResult{}
	err = common.ResponseToObj(response, blurClassifyResult)
	return blurClassifyResult, err
}

func (qualityService *QualityService) BlurClassifyByBase64(imageBase64 string) (*result.BlurClassifyResult, error) {
	return qualityService.blurClassify(imageBase64, BASE64)
}

func (qualityService *QualityService) BlurClassifyByFile(imagePath string) (*result.BlurClassifyResult, error) {
	imageBase64, err := common.LoadImageToBase64(imagePath)
	if nil != err {
		return nil, err
	}
	return qualityService.BlurClassifyByBase64(imageBase64)
}

func (qualityService *QualityService) BlurClassifyByObsUrl(imageUrl string) (*result.BlurClassifyResult, error) {
	return qualityService.blurClassify(imageUrl, OBSURL)
}

func (qualityService *QualityService) headPoseEstimate(image string, imageType int) (*result.HeadPoseEstimateResult, error) {
	uri := fmt.Sprintf(_HEAD_POSE_ESTIMATE, qualityService.projectId)
	jsonObj := common.NewJsonObj()
	switch imageType {
	case BASE64:
		jsonObj.Put("image_base64", image)
		break
	case OBSURL:
		jsonObj.Put("image_url", image)
		break
	}
	response, err := qualityService.accessService.Post(uri, map[string]string{}, jsonObj.GetString())
	if nil != err {
		return nil, err
	}
	headPoseEstimateResult := &result.HeadPoseEstimateResult{}
	err = common.ResponseToObj(response, headPoseEstimateResult)
	return headPoseEstimateResult, err
}

func (qualityService *QualityService) HeadPoseEstimateByBase64(imageBase64 string) (*result.HeadPoseEstimateResult, error) {
	return qualityService.headPoseEstimate(imageBase64, BASE64)
}

func (qualityService *QualityService) HeadPoseEstimateByFile(imagePath string) (*result.HeadPoseEstimateResult, error) {
	imageBase64, err := common.LoadImageToBase64(imagePath)
	if nil != err {
		return nil, err
	}
	return qualityService.HeadPoseEstimateByBase64(imageBase64)
}

func (qualityService *QualityService) HeadPoseEstimateByObsUrl(imageUrl string) (*result.HeadPoseEstimateResult, error) {
	return qualityService.headPoseEstimate(imageUrl, OBSURL)
}

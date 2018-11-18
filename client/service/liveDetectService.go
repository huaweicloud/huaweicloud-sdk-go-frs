package service

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/huaweicloud/huaweicloud-sdk-go-frs/access"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/common"
)

type LiveDetectService struct {
	accessService *access.AccessService
	projectId     string
}

func newLiveDetectService(accessService *access.AccessService, projectId string) *LiveDetectService {
	return &LiveDetectService{accessService, projectId}
}

func (liveDetectService *LiveDetectService) liveDetect(video string, videoType int, actions string, actionTime string) (*result.LiveDetectResult, error) {
	uri := fmt.Sprintf(_LIVE_DETECT_URI, liveDetectService.projectId)
	jsonObj := common.NewJsonObj()
	switch videoType {
	case BASE64:
		jsonObj.Put("video_base64", video)
		break
	case OBSURL:
		jsonObj.Put("video_url", video)
		break
	}
	jsonObj.Put("actions", actions)
	if "" != actionTime {
		jsonObj.Put("action_time", actionTime)
	}
	response, err := liveDetectService.accessService.Post(uri, map[string]string{}, jsonObj.GetString())
	if nil != err {
		return nil, err
	}
	liveDetectResult := &result.LiveDetectResult{}
	err = common.ResponseToObj(response, liveDetectResult)
	return liveDetectResult, err
}

func (liveDetectService *LiveDetectService) LiveDetectByBase64WithActTime(videoBase64 string, actions string, actionTime string) (*result.LiveDetectResult, error) {
	return liveDetectService.liveDetect(videoBase64, BASE64, actions, actionTime)
}

func (liveDetectService *LiveDetectService) LiveDetectByBase64(videoBase64 string, actions string) (*result.LiveDetectResult, error) {
	return liveDetectService.LiveDetectByBase64WithActTime(videoBase64, actions, "")
}

func (liveDetectService *LiveDetectService) LiveDetectByFileWithActTime(videoPath string, actions string, actionTime string) (*result.LiveDetectResult, error) {
	uri := fmt.Sprintf(_LIVE_DETECT_URI, liveDetectService.projectId)
	buffer := &bytes.Buffer{}
	mpWriter := multipart.NewWriter(buffer)
	//Add file
	formFile, formFileErr := mpWriter.CreateFormFile("video_file", "video_file")
	if nil != formFileErr {
		return nil, formFileErr
	}
	file, fileErr := os.Open(videoPath)
	if nil != fileErr {
		return nil, fileErr
	}
	_, fileCopyErr := io.Copy(formFile, file)
	if nil != fileCopyErr {
		return nil, fileCopyErr
	}
	file.Close()
	//Add actions
	if "" != actions {
		mpWriter.WriteField("actions", actions)
	}
	//Add external field
	if "" != actionTime {
		mpWriter.WriteField("action_time", actionTime)
	}
	//Close writer
	mpWriter.Close()

	header := map[string]string{}
	header["content-type"] = mpWriter.FormDataContentType()
	response, err := liveDetectService.accessService.Post(uri, header, buffer)
	if nil != err {
		return nil, err
	}
	liveDetectResult := &result.LiveDetectResult{}
	err = common.ResponseToObj(response, liveDetectResult)
	return liveDetectResult, err
}

func (liveDetectService *LiveDetectService) LiveDetectByFile(videoPath string, actions string) (*result.LiveDetectResult, error) {
	return liveDetectService.LiveDetectByFileWithActTime(videoPath, actions, "")
}

func (liveDetectService *LiveDetectService) LiveDetectByObsUrlWithActTime(videoUrl string, actions string, actionTime string) (*result.LiveDetectResult, error) {
	return liveDetectService.liveDetect(videoUrl, OBSURL, actions, actionTime)
}

func (liveDetectService *LiveDetectService) LiveDetectByObsUrl(videoUrl string, actions string) (*result.LiveDetectResult, error) {
	return liveDetectService.LiveDetectByObsUrlWithActTime(videoUrl, actions, "")
}

package v2

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

type DetectServiceV2 struct {
	accessService *access.AccessService
	projectId     string
}

func newDetectServiceV2(accessService *access.AccessService, projectId string) *DetectServiceV2 {
	return &DetectServiceV2{accessService, projectId}
}

func (detectService *DetectServiceV2) detectFace(image string, imageType int, attributes string) (*result.DetectFaceResult, error) {
	uri := fmt.Sprintf(_FACE_DETECT_URI, detectService.projectId)
	jsonObj := common.NewJsonObj()
	switch imageType {
	case BASE64:
		jsonObj.Put("image_base64", image)
		break
	case OBSURL:
		jsonObj.Put("image_url", image)
		break
	}
	if "" != attributes {
		jsonObj.Put("attributes", attributes)
	}
	response, err := detectService.accessService.Post(uri, map[string]string{}, jsonObj.GetString())
	if nil != err {
		return nil, err
	}
	detectFaceResult := &result.DetectFaceResult{}
	err = common.ResponseToObj(response, detectFaceResult)
	return detectFaceResult, err
}

func (detectService *DetectServiceV2) detectFaceByBase64WithAttr(imageBase64 string, attributes string) (*result.DetectFaceResult, error) {
	return detectService.detectFace(imageBase64, BASE64, attributes)
}

func (detectService *DetectServiceV2) DetectFaceByBase64(imageBase64 string) (*result.DetectFaceResult, error) {
	return detectService.detectFaceByBase64WithAttr(imageBase64, "")
}

func (detectService *DetectServiceV2) detectFaceByFileWithAttr(imagePath string, attributes string) (*result.DetectFaceResult, error) {
	uri := fmt.Sprintf(_FACE_DETECT_URI, detectService.projectId)
	buffer := &bytes.Buffer{}
	mpWriter := multipart.NewWriter(buffer)
	//Add file
	formFile, formFileErr := mpWriter.CreateFormFile("image_file", imagePath)
	if nil != formFileErr {
		return nil, formFileErr
	}
	file, fileErr := os.Open(imagePath)
	if nil != fileErr {
		return nil, fileErr
	}
	_, fileCopyErr := io.Copy(formFile, file)
	if nil != fileCopyErr {
		return nil, fileCopyErr
	}
	file.Close()
	//Add attribute
	if "" != attributes {
		mpWriter.WriteField("attributes", attributes)
	}
	//Close writer
	mpWriter.Close()

	header := map[string]string{}
	header["content-type"] = mpWriter.FormDataContentType()
	response, err := detectService.accessService.Post(uri, header, buffer)
	if nil != err {
		return nil, err
	}
	detectFaceResult := &result.DetectFaceResult{}
	err = common.ResponseToObj(response, detectFaceResult)
	return detectFaceResult, err
}

func (detectService *DetectServiceV2) DetectFaceByFile(imagePath string) (*result.DetectFaceResult, error) {
	return detectService.detectFaceByFileWithAttr(imagePath, "")
}

func (detectService *DetectServiceV2) detectFaceByObsUrlWithAttr(obsUrl string, attributes string) (*result.DetectFaceResult, error) {
	return detectService.detectFace(obsUrl, OBSURL, attributes)
}
func (detectService *DetectServiceV2) DetectFaceByObsUrl(obsUrl string) (*result.DetectFaceResult, error) {
	return detectService.detectFaceByObsUrlWithAttr(obsUrl, "")
}

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

type CompareService struct {
	accessService *access.AccessService
	projectId     string
}

func newCompareService(accessService *access.AccessService, projectId string) *CompareService {
	return &CompareService{accessService, projectId}
}

func (compareService *CompareService) compareFace(image1 string, image2 string, imageType int) (*result.CompareFaceResult, error) {
	uri := fmt.Sprintf(_FACE_COMPARE_URI, compareService.projectId)
	jsonObj := common.NewJsonObj()
	switch imageType {
	case BASE64:
		jsonObj.Put("image1_base64", image1)
		jsonObj.Put("image2_base64", image2)
		break
	case OBSURL:
		jsonObj.Put("image1_url", image1)
		jsonObj.Put("image2_url", image2)
		break
	}
	response, err := compareService.accessService.Post(uri, map[string]string{}, jsonObj.GetString())
	if nil != err {
		return nil, err
	}
	compareFaceResult := &result.CompareFaceResult{}
	err = common.ResponseToObj(response, compareFaceResult)
	return compareFaceResult, err
}

func (compareService *CompareService) CompareFaceByFile(image1Path string, image2Path string) (*result.CompareFaceResult, error) {
	uri := fmt.Sprintf(_FACE_COMPARE_URI, compareService.projectId)
	buffer := &bytes.Buffer{}
	mpWriter := multipart.NewWriter(buffer)
	//Add file1
	formFile1, formFile1Err := mpWriter.CreateFormFile("image1_file", image1Path)
	if nil != formFile1Err {
		return nil, formFile1Err
	}
	file1, file1Err := os.Open(image1Path)
	if nil != file1Err {
		return nil, file1Err
	}
	_, file1CopyErr := io.Copy(formFile1, file1)
	if nil != file1CopyErr {
		return nil, file1CopyErr
	}
	file1.Close()
	//Add file2
	formFile2, formFile2Err := mpWriter.CreateFormFile("image2_file", image2Path)
	if nil != formFile2Err {
		return nil, formFile2Err
	}
	file2, file2Err := os.Open(image2Path)
	if nil != file2Err {
		return nil, file2Err
	}
	_, file2CopyErr := io.Copy(formFile2, file2)
	if nil != file2CopyErr {
		return nil, file2CopyErr
	}
	file2.Close()
	//Close writer
	mpWriter.Close()

	header := map[string]string{}
	header["content-type"] = mpWriter.FormDataContentType()
	response, err := compareService.accessService.Post(uri, header, buffer)
	if nil != err {
		return nil, err
	}
	compareFaceResult := &result.CompareFaceResult{}
	err = common.ResponseToObj(response, compareFaceResult)
	return compareFaceResult, err
}

func (compareService *CompareService) CompareFaceByBase64(image1Base64 string, image2Base64 string) (*result.CompareFaceResult, error) {
	return compareService.compareFace(image1Base64, image2Base64, BASE64)
}

func (compareService *CompareService) CompareFaceByObsUrl(image1Url string, image2Url string) (*result.CompareFaceResult, error) {
	return compareService.compareFace(image1Url, image2Url, OBSURL)
}

package service

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/huaweicloud/huaweicloud-sdk-go-frs/access"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/param"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/common"
)

type FaceService struct {
	accessService *access.AccessService
	projectId     string
}

func newFaceService(accessService *access.AccessService, projectId string) *FaceService {
	return &FaceService{accessService, projectId}
}

func (faceService *FaceService) addFace(faceSetName string, externalImageId string, image string, imageType int, addExternalFields *param.AddExternalFields) (*result.AddFaceResult, error) {
	uri := fmt.Sprintf(_FACE_ADD_URI, faceService.projectId, faceSetName)
	jsonObj := common.NewJsonObj()
	switch imageType {
	case BASE64:
		jsonObj.Put("image_base64", image)
		break
	case OBSURL:
		jsonObj.Put("image_url", image)
		break
	}
	if "" != externalImageId {
		jsonObj.Put("external_image_id", externalImageId)
	}
	if nil != addExternalFields {
		jsonObj.Put("external_fields", addExternalFields.GetString())
	}
	response, err := faceService.accessService.Post(uri, map[string]string{}, jsonObj.GetString())
	if nil != err {
		return nil, err
	}
	addFaceResult := &result.AddFaceResult{}
	err = common.ResponseToObj(response, addFaceResult)
	return addFaceResult, err
}

func (faceService *FaceService) AddFaceByBae64Full(faceSetName string, externalImageId string, imageBase64 string, addExternalFields *param.AddExternalFields) (*result.AddFaceResult, error) {
	return faceService.addFace(faceSetName, externalImageId, imageBase64, BASE64, addExternalFields)
}

func (faceService *FaceService) AddFaceByBae64WithExFields(faceSetName string, imageBase64 string, addExternalFields *param.AddExternalFields) (*result.AddFaceResult, error) {
	return faceService.AddFaceByBae64Full(faceSetName, "", imageBase64, addExternalFields)
}

func (faceService *FaceService) AddFaceByBae64WithExImgId(faceSetName string, externalImageId string, imageBase64 string) (*result.AddFaceResult, error) {
	return faceService.AddFaceByBae64Full(faceSetName, externalImageId, imageBase64, nil)
}

func (faceService *FaceService) AddFaceByBae64(faceSetName string, imageBase64 string) (*result.AddFaceResult, error) {
	return faceService.AddFaceByBae64Full(faceSetName, "", imageBase64, nil)
}

//
func (faceService *FaceService) AddFaceByFileFull(faceSetName string, externalImageId string, imagePath string, addExternalFields *param.AddExternalFields) (*result.AddFaceResult, error) {
	uri := fmt.Sprintf(_FACE_ADD_URI, faceService.projectId, faceSetName)
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
	//Add external image id
	if "" != externalImageId {
		mpWriter.WriteField("external_image_id", externalImageId)
	}
	//Add external field
	if nil != addExternalFields {
		mpWriter.WriteField("external_fields", addExternalFields.GetString())
	}
	//Close writer
	mpWriter.Close()

	header := map[string]string{}
	header["content-type"] = mpWriter.FormDataContentType()
	response, err := faceService.accessService.Post(uri, header, buffer)
	if nil != err {
		return nil, err
	}
	addFaceResult := &result.AddFaceResult{}
	err = common.ResponseToObj(response, addFaceResult)
	return addFaceResult, err
}

func (faceService *FaceService) AddFaceByFileWithExFields(faceSetName string, imagePath string, addExternalFields *param.AddExternalFields) (*result.AddFaceResult, error) {
	return faceService.AddFaceByFileFull(faceSetName, "", imagePath, addExternalFields)
}

func (faceService *FaceService) AddFaceByFileWithExImgId(faceSetName string, externalImageId string, imagePath string) (*result.AddFaceResult, error) {
	return faceService.AddFaceByFileFull(faceSetName, externalImageId, imagePath, nil)
}

func (faceService *FaceService) AddFaceByFile(faceSetName string, imagePath string) (*result.AddFaceResult, error) {
	return faceService.AddFaceByFileFull(faceSetName, "", imagePath, nil)
}

//
func (faceService *FaceService) AddFaceByObsUrlFull(faceSetName string, externalImageId string, imageUrl string, addExternalFields *param.AddExternalFields) (*result.AddFaceResult, error) {
	return faceService.addFace(faceSetName, externalImageId, imageUrl, OBSURL, addExternalFields)
}

func (faceService *FaceService) AddFaceByObsUrlWithExFields(faceSetName string, imageUrl string, addExternalFields *param.AddExternalFields) (*result.AddFaceResult, error) {
	return faceService.AddFaceByObsUrlFull(faceSetName, "", imageUrl, addExternalFields)
}

func (faceService *FaceService) AddFaceByObsUrlWithExImgId(faceSetName string, externalImageId string, imageUrl string) (*result.AddFaceResult, error) {
	return faceService.AddFaceByObsUrlFull(faceSetName, externalImageId, imageUrl, nil)
}

func (faceService *FaceService) AddFaceByObsUrl(faceSetName string, imageUrl string) (*result.AddFaceResult, error) {
	return faceService.AddFaceByObsUrlFull(faceSetName, "", imageUrl, nil)
}

//
func (faceService *FaceService) GetFaces(faceSetName string, offset int, limit int) (*result.GetFaceResult, error) {
	uri := fmt.Sprintf(_FACE_GET_RANGE_URI, faceService.projectId, faceSetName, offset, limit)
	response, err := faceService.accessService.Get(uri)
	if nil != err {
		return nil, err
	}
	getFaceResult := &result.GetFaceResult{}
	err = common.ResponseToObj(response, getFaceResult)
	return getFaceResult, err
}

func (faceService *FaceService) GetFace(faceSetName string, faceId string) (*result.GetFaceResult, error) {
	uri := fmt.Sprintf(_FACE_GET_ONE_URI, faceService.projectId, faceSetName, faceId)
	response, err := faceService.accessService.Get(uri)
	if nil != err {
		return nil, err
	}
	getFaceResult := &result.GetFaceResult{}
	err = common.ResponseToObj(response, getFaceResult)
	return getFaceResult, err
}

func (faceService *FaceService) DeleteFaceByFaceId(faceSetName string, faceId string) (*result.DeleteFaceResult, error) {
	uri := fmt.Sprintf(_FACE_DELETE_BY_FACE_ID_URI, faceService.projectId, faceSetName, faceId)
	response, err := faceService.accessService.Delete(uri)
	if nil != err {
		return nil, err
	}
	deleteFaceResult := &result.DeleteFaceResult{}
	err = common.ResponseToObj(response, deleteFaceResult)
	return deleteFaceResult, err
}

func (faceService *FaceService) DeleteFaceByExternalImageId(faceSetName string, externalImageId string) (*result.DeleteFaceResult, error) {
	uri := fmt.Sprintf(_FACE_DELETE_BY_EXT_ID_URI, faceService.projectId, faceSetName, externalImageId)
	response, err := faceService.accessService.Delete(uri)
	if nil != err {
		return nil, err
	}
	deleteFaceResult := &result.DeleteFaceResult{}
	err = common.ResponseToObj(response, deleteFaceResult)
	return deleteFaceResult, err
}

func (faceService *FaceService) DeleteFaceByFieldId(faceSetName string, fieldId string, fieldValue string) (*result.DeleteFaceResult, error) {
	uri := fmt.Sprintf(_FACE_DELETE_BY_FIELD_ID_URI, faceService.projectId, faceSetName, fieldId, fieldValue)
	response, err := faceService.accessService.Delete(uri)
	if nil != err {
		return nil, err
	}
	deleteFaceResult := &result.DeleteFaceResult{}
	err = common.ResponseToObj(response, deleteFaceResult)
	return deleteFaceResult, err
}

package v2

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-frs/access"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/param"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/common"
)

type FaceSetServiceV2 struct {
	accessService *access.AccessService
	projectId     string
}

func newFaceSetServiceV2(accessService *access.AccessService, projectId string) *FaceSetServiceV2 {
	return &FaceSetServiceV2{accessService, projectId}
}

func (faceSetService *FaceSetServiceV2) CreateFaceSetWithExtFields(faceSetName string, faceSetCapacity int, createExtrnalFields *param.CreateExternalFields) (*result.CreateFaceSetResult, error) {
	uri := fmt.Sprintf(_FACE_SET_CREATE_URI, faceSetService.projectId)
	jsonObj := common.NewJsonObj()
	jsonObj.Put("face_set_name", faceSetName)
	if 0 != faceSetCapacity {
		jsonObj.Put("face_set_capacity", faceSetCapacity)
	}
	if nil != createExtrnalFields {
		jsonObj.Put("external_fields", createExtrnalFields.GetValue())
	}
	response, err := faceSetService.accessService.Post(uri, map[string]string{}, jsonObj.GetString())
	if nil != err {
		return nil, err
	}
	createFaceSetResult := &result.CreateFaceSetResult{}
	err = common.ResponseToObj(response, createFaceSetResult)
	return createFaceSetResult, err
}

func (faceSetService *FaceSetServiceV2) CreateFaceSetWithCapacity(faceSetName string, faceSetCapacity int) (*result.CreateFaceSetResult, error) {
	return faceSetService.CreateFaceSetWithExtFields(faceSetName, faceSetCapacity, nil)
}

func (faceSetService *FaceSetServiceV2) CreateFaceSet(faceSetName string) (*result.CreateFaceSetResult, error) {
	return faceSetService.CreateFaceSetWithExtFields(faceSetName, 0, nil)
}

func (faceSetService *FaceSetServiceV2) GetAllFaceSets() (*result.GetAllFaceSetsResult, error) {
	uri := fmt.Sprintf(_FACE_SET_GET_ALL_URI, faceSetService.projectId)
	response, err := faceSetService.accessService.Get(uri)
	if nil != err {
		return nil, err
	}
	getAllFaceSetsResult := &result.GetAllFaceSetsResult{}
	err = common.ResponseToObj(response, getAllFaceSetsResult)
	return getAllFaceSetsResult, err
}

func (faceSetService *FaceSetServiceV2) GetFaceSet(faceSetName string) (*result.GetFaceSetResult, error) {
	uri := fmt.Sprintf(_FACE_SET_GET_ONE_URI, faceSetService.projectId, faceSetName)
	response, err := faceSetService.accessService.Get(uri)
	if nil != err {
		return nil, err
	}
	getFaceSetResult := &result.GetFaceSetResult{}
	err = common.ResponseToObj(response, getFaceSetResult)
	return getFaceSetResult, err
}

func (faceSetService *FaceSetServiceV2) DeleteFaceSet(faceSetName string) (*result.DeleteFaceSetResult, error) {
	uri := fmt.Sprintf(_FACE_SET_DELETE_URI, faceSetService.projectId, faceSetName)
	response, err := faceSetService.accessService.Delete(uri, "")
	if nil != err {
		return nil, err
	}
	deleteFaceSetResult := &result.DeleteFaceSetResult{}
	err = common.ResponseToObj(response, deleteFaceSetResult)
	return deleteFaceSetResult, err
}

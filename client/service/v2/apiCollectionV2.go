package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/access"
)

type ApiCollectionV2 struct {
	compareService    *CompareServiceV2
	detectService     *DetectServiceV2
	faceService       *FaceServiceV2
	faceSetService    *FaceSetServiceV2
	liveDetectService *LiveDetectServiceV2
	searchService     *SearchServiceV2
}

func NewApiCollectionV2(accessService *access.AccessService, projectId string) *ApiCollectionV2 {
	api := &ApiCollectionV2{
		compareService:    newCompareServiceV2(accessService, projectId),
		detectService:     newDetectServiceV2(accessService, projectId),
		faceService:       newFaceServiceV2(accessService, projectId),
		faceSetService:    newFaceSetServiceV2(accessService, projectId),
		liveDetectService: newLiveDetectServiceV2(accessService, projectId),
		searchService:     newSearchServiceV2(accessService, projectId)}
	return api
}

func (api *ApiCollectionV2) GetCompareService() *CompareServiceV2 {
	return api.compareService
}

func (api *ApiCollectionV2) GetDetectService() *DetectServiceV2 {
	return api.detectService
}

func (api *ApiCollectionV2) GetFaceService() *FaceServiceV2 {
	return api.faceService
}

func (api *ApiCollectionV2) GetFaceSetService() *FaceSetServiceV2 {
	return api.faceSetService
}

func (api *ApiCollectionV2) GetLiveDetectService() *LiveDetectServiceV2 {
	return api.liveDetectService
}

func (api *ApiCollectionV2) GetSearchService() *SearchServiceV2 {
	return api.searchService
}

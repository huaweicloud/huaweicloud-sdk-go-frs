package service

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/access"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/param"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/service/v2"
)

type FrsClient struct {
	compareService    *CompareService
	detectService     *DetectService
	faceService       *FaceService
	faceSetService    *FaceSetService
	liveDetectService *LiveDetectService
	qualityService    *QualityService
	searchService     *SearchService

	apiCollectionV2 *v2.ApiCollectionV2
}

func NewFrsClient(authInfo *param.AuthInfo, projectId string) *FrsClient {
	frsClient := &FrsClient{}
	accessService := access.NewAccessService(authInfo)
	frsClient.initService(accessService, projectId)
	return frsClient
}

func NewFrsClientWithProxy(authInfo *param.AuthInfo, projectId string, proxyHostInfo *param.ProxyHostInfo) *FrsClient {
	frsClient := &FrsClient{}
	accessService := access.NewAccessServiceWithProxy(authInfo, proxyHostInfo)
	frsClient.initService(accessService, projectId)
	return frsClient
}

func (frsClient *FrsClient) initService(accessService *access.AccessService, projectId string) {
	frsClient.compareService = newCompareService(accessService, projectId)
	frsClient.detectService = newDetectService(accessService, projectId)
	frsClient.faceService = newFaceService(accessService, projectId)
	frsClient.faceSetService = newFaceSetService(accessService, projectId)
	frsClient.liveDetectService = newLiveDetectService(accessService, projectId)
	frsClient.qualityService = newQualityService(accessService, projectId)
	frsClient.searchService = newSearchService(accessService, projectId)

	frsClient.apiCollectionV2 = v2.NewApiCollectionV2(accessService, projectId)
}

func (frsClient *FrsClient) GetCompareService() *CompareService {
	return frsClient.compareService
}

func (frsClient *FrsClient) GetDetectService() *DetectService {
	return frsClient.detectService
}

func (frsClient *FrsClient) GetFaceService() *FaceService {
	return frsClient.faceService
}

func (frsClient *FrsClient) GetFaceSetService() *FaceSetService {
	return frsClient.faceSetService
}

func (frsClient *FrsClient) GetLiveDetectService() *LiveDetectService {
	return frsClient.liveDetectService
}

func (frsClient *FrsClient) GetQualityService() *QualityService {
	return frsClient.qualityService
}

func (frsClient *FrsClient) GetSearchService() *SearchService {
	return frsClient.searchService
}

func (frsClient *FrsClient) GetV2() *v2.ApiCollectionV2 {
	return frsClient.apiCollectionV2
}

package service

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/huaweicloud/huaweicloud-sdk-go-frs/access"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/param"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/result"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/common"
)

type SearchService struct {
	accessService *access.AccessService
	projectId     string
}

func newSearchService(accessService *access.AccessService, projectId string) *SearchService {
	return &SearchService{accessService, projectId}
}

func (searchService *SearchService) searchFace(faceSetName string, image string, imageType int, topN int, threshold float64,
	searchSort *param.SearchSort, searchReturnFields *param.SearchReturnFields, filter string) (*result.SearchFaceResult, error) {
	uri := fmt.Sprintf(_FACE_SEARCH_URI, searchService.projectId, faceSetName)
	jsonObj := common.NewJsonObj()
	switch imageType {
	case BASE64:
		jsonObj.Put("image_base64", image)
		break
	case OBSURL:
		jsonObj.Put("image_url", image)
		break
	case FACEID:
		jsonObj.Put("face_id", image)
		break
	}
	//topn
	if -1 != topN {
		jsonObj.Put("top_n", topN)
	}
	//threshold
	if 0 != threshold {
		jsonObj.Put("threshold", threshold)
	}
	//search sort
	if nil != searchSort {
		jsonObj.Put("sort", searchSort.GetValue())
	}
	//search return fields
	if nil != searchReturnFields {
		jsonObj.Put("return_fields", searchReturnFields.GetValue())
	}
	//filter
	if "" != filter {
		jsonObj.Put("filter", filter)
	}
	response, err := searchService.accessService.Post(uri, map[string]string{}, jsonObj.GetString())
	if nil != err {
		return nil, err
	}
	searchFaceResult := &result.SearchFaceResult{}
	err = common.ResponseToObj(response, searchFaceResult)
	return searchFaceResult, err
}

//base64
func (searchService *SearchService) SearchFaceByBase64Full(faceSetName string, imageBase64 string, topN int, threshold float64,
	searchSort *param.SearchSort, searchReturnFields *param.SearchReturnFields, filter string) (*result.SearchFaceResult, error) {
	return searchService.searchFace(faceSetName, imageBase64, BASE64, topN, threshold, searchSort, searchReturnFields, filter)
}

func (searchService *SearchService) SearchFaceByBase64Ext(faceSetName string, imageBase64 string, topN int, threshold float64) (*result.SearchFaceResult, error) {
	return searchService.SearchFaceByBase64Full(faceSetName, imageBase64, topN, threshold, nil, nil, "")
}

func (searchService *SearchService) SearchFaceByBase64(faceSetName string, imageBase64 string) (*result.SearchFaceResult, error) {
	return searchService.SearchFaceByBase64Full(faceSetName, imageBase64, -1, 0, nil, nil, "")
}

//file
func (searchService *SearchService) SearchFaceByFileFull(faceSetName string, imagePath string, topN int, threshold float64,
	searchSort *param.SearchSort, searchReturnFields *param.SearchReturnFields, filter string) (*result.SearchFaceResult, error) {
	uri := fmt.Sprintf(_FACE_SEARCH_URI, searchService.projectId, faceSetName)
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
	//topn
	if -1 != topN {
		mpWriter.WriteField("top_n", strconv.Itoa(topN))
	}
	//threshold
	if 0 != threshold {
		mpWriter.WriteField("threshold", strconv.FormatFloat(threshold, 'f', -1, 64))
	}
	//search sort
	if nil != searchSort {
		mpWriter.WriteField("sort", searchSort.GetString())
	}
	//search return fields
	if nil != searchReturnFields {
		mpWriter.WriteField("return_fields", searchReturnFields.GetString())
	}
	//filter
	if "" != filter {
		mpWriter.WriteField("filter", filter)
	}
	//Close writer
	mpWriter.Close()

	header := map[string]string{}
	header["content-type"] = mpWriter.FormDataContentType()
	response, err := searchService.accessService.Post(uri, header, buffer)
	if nil != err {
		return nil, err
	}
	searchFaceResult := &result.SearchFaceResult{}
	err = common.ResponseToObj(response, searchFaceResult)
	return searchFaceResult, err
}

func (searchService *SearchService) SearchFaceByFileExt(faceSetName string, imagePath string, topN int, threshold float64) (*result.SearchFaceResult, error) {
	return searchService.SearchFaceByFileFull(faceSetName, imagePath, topN, threshold, nil, nil, "")
}

func (searchService *SearchService) SearchFaceByFile(faceSetName string, imagePath string) (*result.SearchFaceResult, error) {
	return searchService.SearchFaceByFileFull(faceSetName, imagePath, -1, 0, nil, nil, "")
}

//obs url
func (searchService *SearchService) SearchFaceByObsUrlFull(faceSetName string, imageUrl string, topN int, threshold float64,
	searchSort *param.SearchSort, searchReturnFields *param.SearchReturnFields, filter string) (*result.SearchFaceResult, error) {
	return searchService.searchFace(faceSetName, imageUrl, OBSURL, topN, threshold, searchSort, searchReturnFields, filter)
}

func (searchService *SearchService) SearchFaceByObsUrlExt(faceSetName string, imageUrl string, topN int, threshold float64) (*result.SearchFaceResult, error) {
	return searchService.SearchFaceByObsUrlFull(faceSetName, imageUrl, topN, threshold, nil, nil, "")
}

func (searchService *SearchService) SearchFaceByObsUrl(faceSetName string, imageUrl string) (*result.SearchFaceResult, error) {
	return searchService.SearchFaceByObsUrlFull(faceSetName, imageUrl, -1, 0, nil, nil, "")
}

//face id
func (searchService *SearchService) SearchFaceByFaceIdFull(faceSetName string, faceId string, topN int, threshold float64,
	searchSort *param.SearchSort, searchReturnFields *param.SearchReturnFields, filter string) (*result.SearchFaceResult, error) {
	return searchService.searchFace(faceSetName, faceId, FACEID, topN, threshold, searchSort, searchReturnFields, filter)
}

func (searchService *SearchService) SearchFaceByFaceIdExt(faceSetName string, faceId string, topN int, threshold float64) (*result.SearchFaceResult, error) {
	return searchService.SearchFaceByFaceIdFull(faceSetName, faceId, topN, threshold, nil, nil, "")
}

func (searchService *SearchService) SearchFaceByFaceId(faceSetName string, faceId string) (*result.SearchFaceResult, error) {
	return searchService.SearchFaceByFaceIdFull(faceSetName, faceId, -1, 0, nil, nil, "")
}

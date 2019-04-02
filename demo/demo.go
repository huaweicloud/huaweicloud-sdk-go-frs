package main

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/param"
	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/service"
)

func demoV1() {
	//Init frs client
	endpoint := "https://face.cn-north-1.myhuaweicloud.com"
	ak := "ak"
	sk := "sk"
	projectId := "project id"
	authInfo := &param.AuthInfo{EndPoint: endpoint, Ak: ak, Sk: sk}
	frs := service.NewFrsClient(authInfo, projectId)

	//Get all service
	frs.GetCompareService()
	frs.GetDetectService()
	frs.GetFaceService()
	frs.GetFaceSetService()
	frs.GetLiveDetectService()
	frs.GetQualityService()
	frs.GetSearchService()

	//Use frs compare service
	res, err := frs.GetCompareService().CompareFaceByFile("image1", "image2")
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(res.Similarity)
		fmt.Println(res.Image1Face.BoundingBox)
		fmt.Println(res.Image2Face.BoundingBox)
	}

	//Use frs detect service
	res2, err := frs.GetDetectService().DetectFaceByBase64("base64")
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(res2)
	}
}

func demoV2() {
	//Init frs client
	endpoint := "https://face.cn-north-1.myhuaweicloud.com"
	ak := "ak"
	sk := "sk"
	projectId := "project id"
	authInfo := &param.AuthInfo{EndPoint: endpoint, Ak: ak, Sk: sk}
	frs := service.NewFrsClient(authInfo, projectId)

	//Get all service
	frs.GetV2().GetCompareService()
	frs.GetV2().GetDetectService()
	frs.GetV2().GetFaceService()
	frs.GetV2().GetFaceSetService()
	frs.GetV2().GetLiveDetectService()
	frs.GetV2().GetSearchService()

	//Use frs compare service
	res, err := frs.GetV2().GetCompareService().CompareFaceByFile("image1", "image2")
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(res.Similarity)
		fmt.Println(res.Image1Face.BoundingBox)
		fmt.Println(res.Image2Face.BoundingBox)
	}

	//Use frs detect service
	res2, err := frs.GetV2().GetDetectService().DetectFaceByBase64("base64")
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(res2)
	}
}

func main() {
	demoV1()
	demoV2()
}

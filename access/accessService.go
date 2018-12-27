package access

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/huaweicloud/huaweicloud-sdk-go-frs/client/param"
)

type AccessService struct {
	authInfo   *param.AuthInfo
	httpClient *http.Client
}

func NewAccessServiceWithProxy(authInfo *param.AuthInfo, proxyHostInfo *param.ProxyHostInfo) *AccessService {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyHostInfo.Proxy)
	}
	transport := &http.Transport{Proxy: proxy}
	httpClient := &http.Client{Transport: transport}
	accessService := &AccessService{authInfo: authInfo, httpClient: httpClient}
	return accessService
}

func NewAccessService(authInfo *param.AuthInfo) *AccessService {
	httpClient := &http.Client{}
	accessService := &AccessService{authInfo: authInfo, httpClient: httpClient}
	return accessService
}

func (accessService *AccessService) Get(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", accessService.authInfo.EndPoint+url, nil)
	if nil != err {
		return nil, err
	}
	response, err := accessService.access(request)
	return response, err
}

func (accessService *AccessService) Post(url string, header map[string]string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest("POST", accessService.authInfo.EndPoint+url, body)
	if nil != err {
		return nil, err
	}
	for key, value := range header {
		request.Header.Add(key, value)
	}
	response, err := accessService.access(request)
	return response, err
}

func (accessService *AccessService) Put(url string, body string) (*http.Response, error) {
	request, err := http.NewRequest("PUT", accessService.authInfo.EndPoint+url, strings.NewReader(body))
	if nil != err {
		return nil, err
	}
	response, err := accessService.access(request)
	return response, err
}

func (accessService *AccessService) Delete(url string, body string) (*http.Response, error) {
	request, err := http.NewRequest("DELETE", accessService.authInfo.EndPoint+url, strings.NewReader(body))
	if nil != err {
		return nil, err
	}
	response, err := accessService.access(request)
	return response, err
}

func (accessService *AccessService) access(request *http.Request) (*http.Response, error) {
	s := Signer{
		AppKey:    accessService.authInfo.Ak,
		AppSecret: accessService.authInfo.Sk,
	}
	if request.Header.Get("content-type") == "" {
		request.Header.Add("content-type", "application/json; charset=utf-8")
	}
	request.Header.Add("x-stage", "RELEASE")
	request.Header.Add("Host", accessService.authInfo.EndPoint)
	err := s.Sign(request)
	if nil != err {
		return nil, err
	}
	return accessService.httpClient.Do(request)
}

package common

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

///////////////Tools
func ResponseToObj(response *http.Response, v interface{}) error {
	s, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err
	}
	if response.StatusCode/200 != 1 {
		return errors.New(fmt.Sprintf("Http error, status code: %d, detail: %s", response.StatusCode, s))
	} else {
		json.Unmarshal(s, v)
		return nil
	}
}

func LoadImageToBase64(imagePath string) (string, error) {
	data, err := ioutil.ReadFile(imagePath)
	if nil != err {
		return "", errors.New("Read file " + imagePath + " failed, detail:" + err.Error())
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

////////////////Json obj
type jsonObj struct {
	prop map[string]interface{}
}

func NewJsonObj() *jsonObj {
	return &jsonObj{prop: map[string]interface{}{}}
}

func (j *jsonObj) Put(key string, value interface{}) {
	j.prop[key] = value
}

func (j *jsonObj) GetString() io.Reader {
	data, err := json.Marshal(j.prop)
	if nil != err {
		return nil
	}
	return strings.NewReader(string(data))
}

/////////////////

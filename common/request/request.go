package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/open-kingfisher/king-utils/common"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

func Get(signing, url string) (interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set(common.HeaderSigning, signing)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	status := response.StatusCode
	type ResponseData struct {
		Code int        `json:"code"`
		Data v1.Service `json:"data"`
		Msg  string     `json:"msg"`
	}
	responseData := ResponseData{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(responseData.Msg)
	}
	return responseData.Data, nil
}

func ListResource(signing, address, path, service, cluster, namespace string) (interface{}, error) {
	uri := fmt.Sprintf("http://%s%s%s?cluster=%s&namespace=%s", address, path, service, cluster, namespace)
	return Get(signing, uri)
}

func GetResource(signing, address, path, service, name, cluster, namespace string) (interface{}, error) {
	uri := fmt.Sprintf("http://%s%s%s/%s?cluster=%s&namespace=%s", address, path, service, name, cluster, namespace)
	return Get(signing, uri)
}

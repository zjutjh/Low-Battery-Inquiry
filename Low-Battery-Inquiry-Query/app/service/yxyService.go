package service

import (
	"Low-Battery-Inquiry-Query/app/fetch"
	"Low-Battery-Inquiry-Query/config/api/yxyApi"
	"encoding/json"
	"errors"
)

type YxyResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func FetchHandleOfGet(url yxyApi.YxyApi) (*YxyResponse, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.Get(yxyApi.YxyHost + string(url))
	if err != nil {
		return nil, errors.New("web error")
	}
	rc := YxyResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, errors.New("request error")
	}
	return &rc, nil
}

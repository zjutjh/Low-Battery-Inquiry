package service

import (
	"Low-Battery-Inquiry-Query/config/api/yxyApi"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/url"
)

type AuthResp struct {
	Token          string `json:"token"`
	ID             string `json:"id"`
	MobilePhone    string `json:"mobile_phone"`
	Sex            int    `json:"sex"`
	Platform       string `json:"platform"`
	ThirdOpenid    string `json:"third_openid"`
	SchoolCode     string `json:"school_code"`
	SchoolName     string `json:"school_name"`
	UserName       string `json:"user_name"`
	UserType       string `json:"user_type"`
	JobNo          string `json:"job_no"`
	UserIDCard     string `json:"user_id_card"`
	UserClass      string `json:"user_class"`
	BindCardStatus int    `json:"bind_card_status"`
}

type RoomInfo struct {
	ID           string `json:"id"`
	SchoolCode   string `json:"school_code"`
	SchoolName   string `json:"school_name"`
	AreaID       string `json:"area_id"`
	AreaName     string `json:"area_name"`
	BuildingCode string `json:"building_code"`
	BuildingName string `json:"building_name"`
	FloorCode    string `json:"floor_code"`
	FloorName    string `json:"floor_name"`
	RoomCode     string `json:"room_code"`
	RoomName     string `json:"room_name"`
	BindType     string `json:"bind_type"`
	CreateTime   string `json:"create_time"`
}

type EleBalance struct {
	SchoolCode      string  `json:"school_code"`
	AreaID          string  `json:"area_id"`
	BuildingCode    string  `json:"building_code"`
	FloorCode       string  `json:"floor_code"`
	RoomCode        string  `json:"room_code"`
	DisplayRoomName string  `json:"display_room_name"`
	Soc             float64 `json:"soc"`
	SocAmount       float64 `json:"soc_amount"`
	Surplus         float64 `json:"surplus"`
	SurplusAmount   float64 `json:"surplus_amount"`
	Subsidy         float64 `json:"subsidy"`
	SubsidyAmount   float64 `json:"subsidy_amount"`
	MdType          string  `json:"md_type"`
	MdName          string  `json:"md_name"`
	RoomStatus      string  `json:"room_status"`
}

func Auth(uid string) (*string, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.Auth))
	if err != nil {
		return nil, err
	}
	params.Set("uid", uid)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	var data AuthResp
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data.Token, nil
}

func ElectricityBalance(token string) (*EleBalance, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.ElectricityBalance))
	if err != nil {
		return nil, err
	}
	params.Set("token", token)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		fmt.Println(resp)
		return nil, errors.New("balance data fetch error")
	}
	var balance EleBalance
	data, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(data, &balance)
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

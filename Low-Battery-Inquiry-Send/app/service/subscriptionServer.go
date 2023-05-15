package service

import (
	"Low-Battery-Inquiry-Send/app/fetch"
	"Low-Battery-Inquiry-Send/app/models"
	"Low-Battery-Inquiry-Send/config/api/wechatApi"
	"Low-Battery-Inquiry-Send/config/config"
	"Low-Battery-Inquiry-Send/config/database"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type SubmitResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type SubSubData struct {
	Value string `json:"value"`
}

type SubData struct {
	CharacterString1 SubSubData `json:"character_string1"`
	Thing2           SubSubData `json:"thing2"`
	Thing3           SubSubData `json:"thing3"`
}

type Data struct {
	Touser     string  `json:"touser"`
	TemplateId string  `json:"template_id"`
	Data       SubData `json:"data"`
}

func QueryRecords() ([]models.User, error) {
	var records []models.User
	result := database.DB.Model(models.User{}).
		Where("id IN (?)", database.DB.Model(models.LowBatteryQueryRecord{}).Distinct("user_id").Select("user_id")).
		Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

func QueryRecordByUserId(userId int) (models.LowBatteryQueryRecord, error) {
	var record models.LowBatteryQueryRecord
	result := database.DB.Model(models.LowBatteryQueryRecord{
		UserID: userId,
	}).First(&record)
	if result.Error != nil {
		return models.LowBatteryQueryRecord{}, result.Error
	}
	return record, nil
}

func DeleteRecord(id int) error {
	result := database.DB.Delete(models.LowBatteryQueryRecord{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAccessToken() (*AccessTokenResponse, error) {
	params := url.Values{}
	Url, err := url.Parse(string(wechatApi.AccessToken))
	if err != nil {
		return nil, err
	}
	params.Set("grant_type", "client_credential")
	params.Set("appid", config.Config.GetString("wechat.appid"))
	params.Set("secret", config.Config.GetString("wechat.appsecret"))
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	f := fetch.Fetch{}
	f.Init()
	res, err := f.Get(urlPath)
	if err != nil {
		return nil, errors.New("web error")
	}
	resp := AccessTokenResponse{}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return nil, errors.New("request error")
	}
	return &resp, nil
}

func SendInquiry(payload models.LowBatteryRemindPayload, accessToken string) error {
	params := url.Values{}
	Url, err := url.Parse(string(wechatApi.SentSubscribe))
	if err != nil {
		return err
	}
	templateId := config.Config.GetString("wechat.templateid")
	if err != nil {
		return err
	}
	params.Set("access_token", accessToken)
	fmt.Println(payload.Address)
	s := strings.Split(payload.Address, "层")[1]
	if strings.HasPrefix(payload.Address, "屏峰校区") && s[0] == 'x' {
		s = "西" + s[1:]
	} else if strings.HasPrefix(payload.Address, "屏峰校区") {
		s = "东" + s[1:]
	}
	if strings.HasPrefix(payload.Address, "朝晖校区") && s[0] == 's' {
		s = "尚" + s[1:]
	} else if strings.HasPrefix(payload.Address, "朝晖校区") && s[0] == 'z' {
		s = "综" + s[1:]
	} else if strings.HasPrefix(payload.Address, "朝晖校区") {
		s = "梦" + s[1:]
	}
	data := Data{TemplateId: templateId, Touser: payload.UserOpenID, Data: SubData{
		CharacterString1: SubSubData{Value: payload.Value},
		Thing2:           SubSubData{Value: s},
		Thing3:           SubSubData{Value: payload.Remark},
	}}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	f := fetch.Fetch{}
	f.Init()
	res, err := f.PostJsonForm(urlPath, data)
	if err != nil {
		return err
	}
	resp := SubmitResponse{}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return err
	}
	fmt.Println(string(rune(payload.UserId)) + resp.Errmsg)
	return nil
}

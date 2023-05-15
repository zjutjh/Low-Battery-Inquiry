package yxyApi

import "Low-Battery-Inquiry-Query/config/config"

var YxyHost = config.Config.GetString("yxy.host")

type YxyApi string

const (
	Auth               YxyApi = "v1/app/auth"
	ElectricityBalance YxyApi = "v1/app/electricity/subsidy/by_user"
)

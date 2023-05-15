package wechatApi

type WechatApi string

const (
	AccessToken   WechatApi = "https://api.weixin.qq.com/cgi-bin/token"
	SentSubscribe WechatApi = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
)

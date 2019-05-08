package e

import "errors"

var (
	ErrorNilParam   = errors.New("check param")
	ErrorNoOpenId   = errors.New("check param: openid")
	ErrorNoToken    = errors.New("check param: access_token")
	ErrorInitClient = errors.New("call NewClientWithParam first")
)

const (
	UnifiedOrderApiUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	QueryOrderApiUrl   = "https://api.mch.weixin.qq.com/pay/orderquery"
)

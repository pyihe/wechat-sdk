package wechat_sdk

import "errors"

var (
	ErrParams    = errors.New("param is empty")
	ErrTradeType = errors.New("need trade_type")
	ErrOpenId    = errors.New("JSAPI need openid")
	ErrCheckSign = errors.New("check sign fail")
	ErrAppId     = errors.New("different appid")
)

const (
	postContentType = "application/xml;charset=utf-8"
	signTypeMD5     = "MD5"
	signType256     = "HMAC-SHA256"
)

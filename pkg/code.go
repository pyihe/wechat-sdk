package pkg

import "errors"

var (
	ErrParams    = errors.New("param is empty")
	ErrTradeType = errors.New("need trade_type")
	ErrOpenId    = errors.New("JSAPI need openid")
	ErrCheckSign = errors.New("check sign fail")
	ErrAppId     = errors.New("different appid")
)

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
	publicKey       = "public.pem" //企业转账到银行卡的AES公钥保存文件名
	signTypeMD5     = "MD5"
	signType256     = "HMAC-SHA256"

	h5     = "MWEB"   //H5支付
	app    = "APP"    //APP支付
	jsApi  = "JSAPI"  //JSAPI支付
	native = "NATIVE" //Native(扫码)支付
)

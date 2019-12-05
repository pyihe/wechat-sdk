package dev

import (
	"errors"
	"io"
)

var (
	defaultPayer *myPayer
)

type ResultParam interface {
	Param(key string) (interface{}, error) //根据key获取微信返回的对应的value
	ListParam() Params                     //将微信返回的参数转换为PayParams
	//checkWxSign(signType string) (bool, error) //校验微信返回的签名，签名方式默认MD5
}

type WePayer interface {
	//支付相关
	//统一下单
	UnifiedOrder(param Params) (ResultParam, error) //统一下单
	//查询订单
	QueryOrder(param Params) (ResultParam, error) //查询订单
	//关闭订单
	CloseOrder(param Params) (ResultParam, error)
	//退款
	RefundOrder(param Params, certPath string) (ResultParam, error)
	//退款查询
	RefundQuery(param Params) (ResultParam, error)
	//解析退款通知
	RefundNotify(body io.ReadCloser) (ResultParam, error)

	//小程序相关
	//获取授权access_token
	GetAccessTokenForMini() (ResultParam, error) //获取小程序接口凭证，使用者自己保存token，过期重新获取
	//获取微信信息
	GetUserInfoForMini(code, dataStr, ivStr string) (ResultParam, error)
	//获取微信手机号码
	GetUserPhoneForMini(code string, dataStr string, ivStr string) (ResultParam, error)
	//获取session_key
	GetSessionKeyAndOpenId(code string) (ResultParam, error)
}

type option func(*myPayer)

type myPayer struct {
	appId  string //appid
	mchId  string //mchid
	secret string //secret用于获取token
	apiKey string //用于支付
}

func NewPayer(options ...option) WePayer {
	defaultPayer = &myPayer{}
	for _, option := range options {
		option(defaultPayer)
	}
	return defaultPayer
}

func WithAppId(appId string) option {
	return func(payer *myPayer) {
		payer.appId = appId
	}
}

func WithMchId(mchId string) option {
	return func(payer *myPayer) {
		payer.mchId = mchId
	}
}

func WithSecret(secret string) option {
	return func(payer *myPayer) {
		payer.secret = secret
	}
}

func WithApiKey(key string) option {
	return func(payer *myPayer) {
		payer.apiKey = key
	}
}

func (m *myPayer) checkForPay() error {
	if m.appId == "" {
		return errors.New("need appid")
	}
	if m.mchId == "" {
		return errors.New("need mch_id")
	}
	if m.apiKey == "" {
		return errors.New("need api key")
	}
	return nil
}

func (m *myPayer) checkForAccess() error {
	if m.appId == "" {
		return errors.New("need appid")
	}
	if m.secret == "" {
		return errors.New("need secret")
	}
	return nil
}

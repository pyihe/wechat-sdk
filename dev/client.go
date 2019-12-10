package dev

import (
	"errors"
	"io"
)

var (
	defaultPayer *myPayer
)

type ResultParam interface {
	//获取整型数据, 如支付/退款金额
	GetInt64(key string, base int) (value int64, err error)
	//获取字符串数据, 如订单号
	GetString(key string) (value string, err error)
	//
	Data() map[string]string
}

type WePayer interface {
	//支付相关
	//统一下单
	UnifiedOrder(param Param) (ResultParam, error)
	//扫码下单
	UnifiedMicro(param Param) (ResultParam, error)
	//撤销订单
	ReverseOrder(param Param, p12CertPath string) (ResultParam, error)
	//查询订单
	UnifiedQuery(param Param) (ResultParam, error)
	//关闭订单
	CloseOrder(param Param) (ResultParam, error)
	//退款
	RefundOrder(param Param, p12CertPath string) (ResultParam, error)
	//退款查询
	RefundQuery(param Param) (ResultParam, error)
	//解析退款通知, 结果将不会返回req_info
	RefundNotify(body io.Reader) (ResultParam, error)
	//下单对账单
	DownloadBill(param Param, fileSavePath string) error
	//下载资金账单
	DownloadFundFlow(param Param, p12CertPath string, fileSavePath string) error

	//小程序相关
	//获取授权access_token
	GetAccessTokenForMini() (Param, error) //获取小程序接口凭证，使用者自己保存token，过期重新获取
	//获取微信信息
	GetUserInfoForMini(code, dataStr, ivStr string) (Param, error)
	//获取微信手机号码
	GetUserPhoneForMini(code string, dataStr string, ivStr string) (Param, error)
	//获取session_key
	GetSessionKeyAndOpenId(code string) (Param, error)
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

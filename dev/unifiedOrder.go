package dev

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	统一下单
*/

var (
	unifiedMustParam     = []string{"appid", "mch_id", "nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "notify_url", "trade_type"}
	unifiedOptionalParam = []string{"device_info", "sign_type", "detail", "attach", "fee_type", "time_start", "time_expire", "goods_tag", "limit_pay", "receipt", "openid"}
)

const unifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"

type unifiedResult struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	TradeType  string `xml:"trade_type"`
	PrepayId   string `xml:"prepay_id"`
	CodeUrl    string `xml:"code_url"`
}

func (u *unifiedResult) Param(key string) (interface{}, error) {
	var err error
	switch key {
	case "return_code":
		return u.ReturnCode, err
	case "return_msg":
		return u.ReturnMsg, err
	case "appid":
		return u.Appid, err
	case "mch_id":
		return u.MchId, err
	case "device_info":
		return u.DeviceInfo, err
	case "nonce_str":
		return u.NonceStr, err
	case "sign":
		return u.Sign, err
	case "result_code":
		return u.ResultCode, err
	case "err_code":
		return u.ErrCode, err
	case "err_code_des":
		return u.ErrCodeDes, err
	case "trade_type":
		return u.TradeType, err
	case "prepay_id":
		return u.PrepayId, err
	case "code_url":
		return u.CodeUrl, err
	default:
		err = errors.New(fmt.Sprintf("invalid key: %s", key))
		return nil, err
	}
}

func (u unifiedResult) ListParam() Params {
	p := make(Params)

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).IsZero() {
			tagName := strings.Split(string(t.Field(i).Tag), "\"")[1]
			p[tagName] = v.Field(i).Interface()
		}
	}
	return p
}

//统一下单
func (m *myPayer) UnifiedOrder(param Params) (ResultParam, error) {
	if param == nil {
		return nil, e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)
	//获取交易类型和签名类型
	var tradeType string
	var signType = e.SignTypeMD5 //默认MD5签名方式
	if t, ok := param["trade_type"]; ok {
		tradeType = t.(string)
	} else {
		return nil, e.ErrTradeType
	}
	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	//校验参数是否传对了
	if tradeType == "JSAPI" {
		if _, ok := param["openid"]; !ok {
			return nil, e.ErrOpenId
		}
	}
	//这里校验是否包含必传的参数
	for _, v := range unifiedMustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New(fmt.Sprintf("need %s", v))
		}
	}
	//这里校验是否包含不必要的参数
	for key := range param {
		if !util.HaveInArray(unifiedMustParam, key) && !util.HaveInArray(unifiedOptionalParam, key) {
			return nil, errors.New(fmt.Sprintf("no need %s param", key))
		}
	}

	sign, err := param.Sign(signType)
	if err != nil {
		return nil, err
	}
	//将签名添加到需要发送的参数里
	param.Add("sign", sign)
	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var result *unifiedResult
	var request = &util.PostRequest{
		Body:        reader,
		Url:         unifiedOrderUrl,
		ContentType: "application/xml;charset=utf-8",
	}
	err = util.PostToWx(request, &result)
	if err != nil {
		return nil, err
	}

	if result.ReturnCode != "SUCCESS" {
		return nil, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != "SUCCESS" {
		return nil, errors.New(result.ErrCodeDes)
	}

	sign, err = result.ListParam().Sign(signType)
	if err != nil || sign != result.Sign {
		return nil, e.ErrCheckSign
	}
	return result, err
}

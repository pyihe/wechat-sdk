package dev

import (
	"errors"
	"fmt"

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

//统一下单
func (m *myPayer) UnifiedOrder(param Param) (ResultParam, error) {
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

	sign := param.Sign(signType)
	//将签名添加到需要发送的参数里
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         unifiedOrderUrl,
		ContentType: "application/xml;charset=utf-8",
	}
	result, err := postToWx(request)
	if err != nil {
		return nil, err
	}

	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return nil, errors.New(returnMsg)
	}

	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return nil, errors.New(errDes)
	}
	sign = result.Sign(signType)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, e.ErrCheckSign
	}
	return result, err
}

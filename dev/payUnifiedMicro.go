package dev

import (
	"errors"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

func (m *myPayer) UnifiedMicro(param Param) (ResultParam, error) {
	if param == nil {
		return nil, e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//获取交易类型和签名类型
	var (
		microMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "body", "out_trade_no", "total_fee", "spbill_create_ip", "auth_code"}
		microOptionalParam = []string{"device_info", "sign_type", "detail", "attach", "fee_type", "goods_tag", "limit_pay", "time_start", "time_expire", "receipt", "scene_info"}
		signType           = e.SignTypeMD5 //默认MD5签名方式
	)

	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	for _, v := range microMustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}

	for key := range param {
		if !util.HaveInArray(microMustParam, key) && !util.HaveInArray(microOptionalParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/micropay",
		ContentType: e.PostContentType,
	}
	response, err := postToWx(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
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
	return result, nil
}

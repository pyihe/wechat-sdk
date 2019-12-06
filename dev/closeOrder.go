package dev

import (
	"errors"
	"fmt"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	关闭订单
*/

var (
	closeMustParam     = []string{"appid", "mch_id", "out_trade_no", "nonce_str", "sign"}
	closeOptionalParam = []string{"sign_type"}
)

const closeOrderUrl = "https://api.mch.weixin.qq.com/pay/closeorder"

//关闭订单
func (m *myPayer) CloseOrder(param Param) (ResultParam, error) {
	if param == nil {
		return nil, e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var signType = e.SignTypeMD5
	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	for _, v := range closeMustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New(fmt.Sprintf("need %s", v))
		}
	}
	for key := range param {
		if !util.HaveInArray(closeMustParam, key) && !util.HaveInArray(closeOptionalParam, key) {
			return nil, errors.New(fmt.Sprintf("no need %s", key))
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
		Url:         closeOrderUrl,
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
	return result, nil
}

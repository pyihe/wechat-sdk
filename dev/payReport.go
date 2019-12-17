package dev

import (
	"errors"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	交易保障
*/

func (m *myPayer) Report(param Param) error {
	if param == nil {
		return e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var signType = e.SignTypeMD5
	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	var (
		reportMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "interface_url", "execute_time", "return_code", "return_msg", "result_code", "user_ip"}
		reportOptionalParam = []string{"device_info", "sign_type", "err_code", "err_code_des", "out_trade_no", "time"}
		reportMicroParam    = []string{"appid", "mch_id", "nonce_str", "sign", "interface_url", "trades", "user_ip"}
		reportMicroOptional = []string{"device_info"}
	)

	if v := param.Get("trade"); v != nil {
		for _, k := range reportMicroParam {
			if k == "sign" {
				continue
			}
			if _, ok := param[k]; !ok {
				return errors.New("need param: " + k)
			}
		}
		for key := range param {
			if !util.HaveInArray(reportMicroParam, key) && !util.HaveInArray(reportMicroOptional, key) {
				return errors.New("no need param: " + key)
			}
		}
	} else {
		for _, k := range reportMustParam {
			if k == "sign" {
				continue
			}
			if _, ok := param[k]; !ok {
				return errors.New("need param: " + k)
			}
		}
		for key := range param {
			if !util.HaveInArray(reportMustParam, key) && !util.HaveInArray(reportOptionalParam, key) {
				return errors.New("no need param: " + key)
			}
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/payitil/report",
		ContentType: e.PostContentType,
	}
	response, err := postToWx(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return errors.New(returnMsg)
	}

	if resultCode, err := result.GetString("result_code"); err == nil && resultCode != "SUCCESS" {
		return errors.New("report fail")
	}
	return nil
}

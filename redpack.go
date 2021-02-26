package wechat_sdk

import (
	"errors"
	"fmt"

	"github.com/pyihe/util/certs"
)

func (m *myPayer) SendRedPack(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	//读取证书
	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("wxappid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		mustParam = map[string]struct{}{
			"nonce_str":    {},
			"sign":         {},
			"mch_billno":   {},
			"mch_id":       {},
			"wxappid":      {},
			"send_name":    {},
			"re_openid":    {},
			"total_amount": {},
			"total_num":    {},
			"wishing":      {},
			"client_ip":    {},
			"act_name":     {},
			"remark":       {},
		}
		optionalParam = map[string]struct{}{
			"scene_id":  {},
			"risk_info": {},
		}
		signType = signTypeMD5
	)

	//判断是否缺少必须的参数
	for k := range mustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, fmt.Errorf("need param: %s", k)
		}
	}

	//判断是否包含了非法的参数
	if err = param.RangeIn(func(key string) bool {
		_, ok := mustParam[key]
		if !ok {
			_, ok = optionalParam[key]
		}
		return ok
	}); err != nil {
		return nil, err
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack",
		ContentType: postContentType,
	}
	response, err := postToWxWithCert(request, cert)
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
	if resultSign, _ := result.GetString("sign"); resultSign != sign {
		return nil, ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) SendGroupRedPack(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	//读取证书
	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("wxappid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		mustParam = map[string]struct{}{
			"nonce_str":    {},
			"sign":         {},
			"mch_billno":   {},
			"mch_id":       {},
			"wxappid":      {},
			"send_name":    {},
			"re_openid":    {},
			"total_amount": {},
			"total_num":    {},
			"amt_type":     {},
			"wishing":      {},
			"act_name":     {},
			"remark":       {},
		}
		optionalParam = map[string]struct{}{
			"scene_id":  {},
			"risk_info": {},
		}
		signType = signTypeMD5
	)

	//判断是否缺少必须的参数
	for k := range mustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, fmt.Errorf("need param: %s", k)
		}
	}

	//判断是否包含了非法的参数
	if err = param.RangeIn(func(key string) bool {
		_, ok := mustParam[key]
		if !ok {
			_, ok = optionalParam[key]
		}
		return ok
	}); err != nil {
		return nil, err
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack",
		ContentType: postContentType,
	}
	response, err := postToWxWithCert(request, cert)
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
	if resultSign, _ := result.GetString("sign"); resultSign != sign {
		return nil, ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) GetRedPackRecords(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	//读取证书
	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		mustParam = map[string]struct{}{
			"nonce_str":  {},
			"sign":       {},
			"mch_billno": {},
			"mch_id":     {},
			"appid":      {},
			"bill_type":  {},
		}
		signType = signTypeMD5
	)

	//判断是否缺少必须的参数
	for k := range mustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, fmt.Errorf("need param: %s", k)
		}
	}

	//判断是否包含非法参数
	if err = param.RangeIn(func(key string) bool {
		_, ok := mustParam[key]
		return ok
	}); err != nil {
		return nil, err
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo",
		ContentType: postContentType,
	}
	response, err := postToWxWithCert(request, cert)
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
	if resultSign, _ := result.GetString("sign"); resultSign != sign {
		return nil, ErrCheckSign
	}
	return result, nil
}

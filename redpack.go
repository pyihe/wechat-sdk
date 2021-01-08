package wechat_sdk

import (
	"errors"

	"github.com/pyihe/util/certs"
	"github.com/pyihe/util/utils"
	"github.com/pyihe/wechat-sdk/pkg"
)

func (m *myPayer) SendRedPack(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
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
		mustParam     = []string{"nonce_str", "sign", "mch_billno", "mch_id", "wxappid", "send_name", "re_openid", "total_amount", "total_num", "wishing", "client_ip", "act_name", "remark"}
		optionalParam = []string{"scene_id", "risk_info"}
		signType      = pkg.SignTypeMD5
	)

	for _, v := range mustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}

	for key := range param {
		if !utils.Contain(mustParam, key) && !utils.Contain(optionalParam, key) {
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
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack",
		ContentType: pkg.PostContentType,
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
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) SendGroupRedPack(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
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
		mustParam     = []string{"nonce_str", "sign", "mch_billno", "mch_id", "wxappid", "send_name", "re_openid", "total_amount", "total_num", "amt_type", "wishing", "act_name", "remark"}
		optionalParam = []string{"scene_id", "risk_info"}
		signType      = pkg.SignTypeMD5
	)

	for _, v := range mustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}

	for key := range param {
		if !utils.Contain(mustParam, key) && !utils.Contain(optionalParam, key) {
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
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack",
		ContentType: pkg.PostContentType,
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
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) GetRedPackRecords(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
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
		mustParam = []string{"nonce_str", "sign", "mch_billno", "mch_id", "appid", "bill_type"}
		signType  = pkg.SignTypeMD5
	)
	for _, v := range mustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}
	for key := range param {
		if !utils.Contain(mustParam, key) {
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
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo",
		ContentType: pkg.PostContentType,
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
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

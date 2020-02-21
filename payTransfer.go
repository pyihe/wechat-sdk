package wechat_sdk

import (
	"errors"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	企业支付,包括企业付款到零钱、查询企业付款到零钱、企业付款到银行卡
*/

func (m *myPayer) Transfers(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	cert, err := util.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("mch_appid", m.appId)
	param.Add("mchid", m.mchId)

	var transMustParam = []string{"mch_appid", "mchid", "nonce_str", "sign", "partner_trade_no", "openid", "check_name", "amount", "desc", "spbill_create_ip"}
	for _, k := range transMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, errors.New("need param: " + k)
		}
	}

	var transOptionalParam = []string{"device_info", "re_user_name"}
	for k := range param {
		if !util.HaveInArray(transMustParam, k) && !util.HaveInArray(transOptionalParam, k) {
			return nil, errors.New("no need param: " + k)
		}
	}

	sign := param.Sign(e.SignTypeMD5)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers",
		ContentType: e.PostContentType,
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
	return result, nil
}

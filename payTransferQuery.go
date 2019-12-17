package wechat_sdk

import (
	"errors"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	查询转账到零钱
*/

func (m *myPayer) TransfersQuery(param Param, p12CertPath string) (ResultParam, error) {
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

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var queryTransferMustParam = []string{"nonce_str", "sign", "partner_trade_no", "mch_id", "appid"}
	for _, k := range queryTransferMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, errors.New("need param: " + k)
		}
	}

	for k := range param {
		if !util.HaveInArray(queryTransferMustParam, k) {
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
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo",
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

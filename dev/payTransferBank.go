package dev

import (
	"encoding/base64"
	"errors"
	"io/ioutil"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	企业付款到银行卡
*/

func (m *myPayer) TransferBank(param Param, p12CertPath string, publicKeyPath string) (ResultParam, error) {
	if param == nil {
		return nil, e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	cert, err := util.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("mch_id", m.mchId)

	var bankMustParam = []string{"mch_id", "partner_trade_no", "nonce_str", "sign", "enc_bank_no", "enc_true_name", "bank_code", "amount"}
	for _, k := range bankMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, errors.New("need param: " + k)
		}
	}

	var bankOptionalParam = []string{"desc"}
	for k := range param {
		if !util.HaveInArray(bankMustParam, k) && !util.HaveInArray(bankOptionalParam, k) {
			return nil, errors.New("no need param: " + k)
		}
	}

	bankCard := param.Get("enc_bank_no").(string)
	bankName := param.Get("enc_true_name").(string)
	encryptBankCard, err := util.RsaEncrypt([]byte(bankCard), publicKey)
	if err != nil {
		return nil, err
	}
	encryptBankName, err := util.RsaEncrypt([]byte(bankName), publicKey)
	if err != nil {
		return nil, err
	}
	param.Add("enc_bank_no", base64.StdEncoding.EncodeToString(encryptBankCard))
	param.Add("enc_true_name", base64.StdEncoding.EncodeToString(encryptBankName))

	sign := param.Sign(e.SignTypeMD5)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank",
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
	sign = result.Sign(e.SignTypeMD5)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, e.ErrCheckSign
	}
	return result, nil
}

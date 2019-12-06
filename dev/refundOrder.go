package dev

import (
	"errors"
	"fmt"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	订单退款申请
*/

var (
	refundMustParams     = []string{"appid", "mch_id", "nonce_str", "sign", "out_refund_no", "total_fee", "refund_fee"}
	refundOneParams      = []string{"transaction_id", "out_trade_no"}
	refundOptionalParams = []string{"sign_type", "refund_fee_type", "refund_desc", "refund_account", "notify_url"}
)

const refundApiUrl = "https://api.mch.weixin.qq.com/secapi/pay/refund"

func (m *myPayer) RefundOrder(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, errors.New("param is empty")
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	//读取证书
	cert, err := util.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var signType = e.SignTypeMD5

	//校验订单号
	var count = 0
	for _, k := range queryOneParam {
		if v := param.Get(k); v != nil {
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("need order number: transaction_id or out_trade_no")
	} else if count > 1 {
		return nil, errors.New("just one order number: transaction_id or out_trade_no")
	}
	for _, k := range refundMustParams {
		if k == "sign" {
			continue
		}
		if param.Get(k) == nil {
			return nil, errors.New(fmt.Sprintf("need %s", k))
		}
	}

	for k := range param {
		if k == "sign_type" {
			signType = param[k].(string)
		}
		if !util.HaveInArray(refundMustParams, k) && !util.HaveInArray(refundOptionalParams, k) && !util.HaveInArray(refundOneParams, k) {
			return nil, errors.New(fmt.Sprintf("no need %s", k))
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
		Url:         refundApiUrl,
		ContentType: "application/xml;charset=utf-8",
	}

	result, err := postToWxWithCert(request, cert)
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
	if resultSign, _ := result.GetString("sign"); resultSign != sign {
		return nil, e.ErrCheckSign
	}
	return result, nil
}

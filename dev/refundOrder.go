package dev

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

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

type refundResult struct {
	ReturnCode          string `xml:"return_code"`
	ReturnMsg           string `xml:"return_msg"`
	ResultCode          string `xml:"result_code"`
	ErrCode             string `xml:"err_code"`
	ErrCodeDes          string `xml:"err_code_des"`
	Appid               string `xml:"appid"`
	MchId               string `xml:"mch_id"`
	NonceStr            string `xml:"nonce_str"`
	Sign                string `xml:"sign"`
	TransactionId       string `xml:"transaction_id"`
	OutTradeNo          string `xml:"out_trade_no"`
	OutRefundNo         string `xml:"out_refund_no"`
	RefundId            string `xml:"refund_id"`
	RefundFee           int64  `xml:"refund_fee"`
	SettlementRefundFee int64  `xml:"settlement_refund_fee"`
	TotalFee            int64  `xml:"total_fee"`
	SettlementTotalFee  int64  `xml:"settlement_total_fee"`
	FeeType             string `xml:"fee_type"`
	CashFee             int64  `xml:"cash_fee"`
	CashFeeType         string `xml:"cash_fee_type"`
	CashRefundFee       int64  `xml:"cash_refund_fee"`
	CouponTypeN         string `xml:"coupon_type_$n"`
	CouponRefundFee     int64  `xml:"coupon_refund_fee"`
	CouponRefundFeeN    int64  `xml:"coupon_refund_fee_$n"`
	CouponRefundCount   int64  `xml:"coupon_refund_count"`
	CouponRefundIdN     string `xml:"coupon_refund_id_$n"`
}

func (r *refundResult) Param(key string) (interface{}, error) {
	var err error
	switch key {
	case "return_code":
		return r.ReturnCode, err
	case "return_msg":
		return r.ReturnMsg, err
	case "result_code":
		return r.ResultCode, err
	case "err_code":
		return r.ErrCode, err
	case "err_code_des":
		return r.ErrCodeDes, err
	case "appid":
		return r.Appid, err
	case "mch_id":
		return r.MchId, err
	case "nonce_str":
		return r.NonceStr, err
	case "sign":
		return r.Sign, err
	case "transaction_id":
		return r.TransactionId, err
	case "out_trade_no":
		return r.OutTradeNo, err
	case "out_refund_no":
		return r.OutRefundNo, err
	case "refund_id":
		return r.RefundId, err
	case "refund_fee":
		return r.RefundFee, err
	case "settlement_refund_fee":
		return r.SettlementRefundFee, err
	case "total_fee":
		return r.TotalFee, err
	case "settlement_total_fee":
		return r.SettlementTotalFee, err
	case "fee_type":
		return r.FeeType, err
	case "cash_fee":
		return r.CashFee, err
	case "cash_fee_type":
		return r.CashFeeType, err
	case "cash_refund_fee":
		return r.CashRefundFee, err
	case "coupon_type_$n":
		return r.CouponTypeN, err
	case "coupon_refund_fee":
		return r.CouponRefundFee, err
	case "coupon_refund_fee_":
		return r.CouponRefundFeeN, err
	case "coupon_refund_count":
		return r.CouponRefundCount, err
	case "coupon_refund_id_$n":
		return r.CouponRefundIdN, err
	default:
		err = errors.New(fmt.Sprintf("invalid key: %s", key))
		return nil, err
	}
}

func (r refundResult) ListParam() Params {
	p := make(Params)

	t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)

	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).IsZero() {
			tagName := strings.Split(string(t.Field(i).Tag), "\"")[1]
			p[tagName] = v.Field(i).Interface()
		}
	}
	return p
}

func (m *myPayer) RefundOrder(param Params, p12CertPath string) (ResultParam, error) {
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

	sign, err := param.Sign(signType)
	if err != nil {
		return nil, err
	}
	param.Add("sign", sign)
	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var result *refundResult
	var request = &util.PostRequest{
		Body:        reader,
		Url:         refundApiUrl,
		ContentType: "application/xml;charset=utf-8",
	}

	err = util.PostToWxWithCert(request, cert, &result)
	if err != nil {
		return nil, err
	}
	if result.ReturnCode != "SUCCESS" {
		return nil, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != "SUCCESS" {
		return nil, errors.New(result.ErrCodeDes)
	}
	if result.Appid != m.appId {
		return nil, errors.New(fmt.Sprintf("you got appid:%s from WeiXin", result.Appid))
	}
	if result.MchId != m.mchId {
		return nil, errors.New(fmt.Sprintf("you got mch_id:%s from WeiXin", result.MchId))
	}
	return result, nil
}

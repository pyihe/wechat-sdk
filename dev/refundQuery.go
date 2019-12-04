package dev

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

/*
	查询退款
*/

var (
	refundQueryMustParam     = []string{"appid", "mch_id", "nonce_str", "sign"}
	refundQueryOneParam      = []string{"transaction_id", "out_trade_no", "out_refund_no", "refund_id"}
	refundQueryOptionalParam = []string{"sign_type", "offset"}
)

const queryRefundUrl = "https://api.mch.weixin.qq.com/pay/refundquery"

type queryRefundResult struct {
	ReturnCode           string `xml:"return_code"`
	ReturnMsg            string `xml:"return_msg"`
	ResultCode           string `xml:"result_code"`
	ErrCode              string `xml:"err_code"`
	ErrCodeDes           string `xml:"err_code_des"`
	Appid                string `xml:"appid"`
	MchId                string `xml:"mch_id"`
	NonceStr             string `xml:"nonce_str"`
	Sign                 string `xml:"sign"`
	TotalRefundCount     int64  `xml:"total_refund_count"`
	TransactionId        string `xml:"transaction_id"`
	OutTradeNo           string `xml:"out_trade_no"`
	TotalFee             int64  `xml:"total_fee"`
	SettlementTotalFee   int64  `xml:"settlement_total_fee"`
	FeeType              string `xml:"fee_type"`
	CashFee              int64  `xml:"cash_fee"`
	RefundCount          int64  `xml:"refund_count"`
	OutRefundNoN         string `xml:"out_refund_no_$n"`
	RefundIdN            string `xml:"refund_id_$n"`
	RefundChannelN       string `xml:"refund_channel_$n"`
	RefundFeeN           int64  `xml:"refund_fee_$n"`
	SettlementRefundFeeN int64  `xml:"settlement_refund_fee_$n"`
	CouponTypeNM         string `xml:"coupon_type_$n_$m"`
	CouponRefundFeeN     int64  `xml:"coupon_refund_fee_$n"`
	CouponRefundCountN   int64  `xml:"coupon_refund_count_$n"`
	CouponRefundIdNM     string `xml:"coupon_refund_id_$n_$m"`
	CouponRefundFeeNM    int64  `xml:"coupon_refund_fee_$n_$m"`
	RefundStatusN        string `xml:"refund_status_$n"`
	RefundAccountN       string `xml:"refund_account_$n"`
	RefundRecvAccountN   string `xml:"refund_recv_account_$n"`
	RefundSuccessTimeN   string `xml:"refund_success_time_$n"`
}

func (q *queryRefundResult) Param(key string) (interface{}, error) {
	var err error
	switch key {
	case "return_code":
		return q.ReturnCode, err
	case "return_msg":
		return q.ReturnMsg, err
	case "result_code":
		return q.ResultCode, err
	case "err_code":
		return q.ErrCode, err
	case "err_code_des":
		return q.ErrCodeDes, err
	case "appid":
		return q.Appid, err
	case "mch_id":
		return q.MchId, err
	case "nonce_str":
		return q.NonceStr, err
	case "sign":
		return q.Sign, err
	case "total_refund_count":
		return q.TotalRefundCount, err
	case "transaction_id":
		return q.TransactionId, err
	case "out_trade_no":
		return q.OutTradeNo, err
	case "total_fee":
		return q.TotalFee, err
	case "settlement_total_fee":
		return q.SettlementTotalFee, err
	case "fee_type":
		return q.FeeType, err
	case "cash_fee":
		return q.CashFee, err
	case "refund_count":
		return q.RefundCount, err
	case "out_refund_no_$n":
		return q.OutRefundNoN, err
	case "refund_id_$n":
		return q.RefundIdN, err
	case "refund_channel_$n":
		return q.RefundChannelN, err
	case "refund_fee_$n":
		return q.RefundFeeN, err
	case "settlement_refund_fee_$n":
		return q.SettlementRefundFeeN, err
	case "coupon_type_$n_$m":
		return q.CouponTypeNM, err
	case "coupon_refund_fee_$n":
		return q.CouponRefundFeeN, err
	case "coupon_refund_count_$n":
		return q.CouponRefundCountN, err
	case "coupon_refund_id_$n_$m":
		return q.CouponRefundIdNM, err
	case "coupon_refund_fee_$n_$m":
		return q.CouponRefundFeeNM, err
	case "refund_status_$n":
		return q.RefundStatusN, err
	case "refund_account_$n":
		return q.RefundAccountN, err
	case "refund_recv_accout_$n":
		return q.RefundRecvAccountN, err
	case "refund_success_time_$n":
		return q.RefundSuccessTimeN, err
	default:
		err = errors.New(fmt.Sprintf("invalid key: %s", key))
		return nil, err
	}
}

func (q queryRefundResult) ListParam() Params {
	p := make(Params)

	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).IsZero() {
			tagName := strings.Split(string(t.Field(i).Tag), "\"")[1]
			p[tagName] = v.Field(i).Interface()
		}
	}
	return p
}

func (m *myPayer) QueryRefund(param Params) (ResultParam, error) {
	if param == nil {
		return nil, errors.New("param is empty")
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)
	//TODO
	return nil, nil
}

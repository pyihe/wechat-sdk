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
	订单查询
*/

var (
	queryMustParam     = []string{"appid", "mch_id", "nonce_str"}
	queryOptionalParam = []string{"sign_type", "transaction_id", "out_trade_no"}
)

const queryApiUrl = "https://api.mch.weixin.qq.com/pay/orderquery"

type queryResult struct {
	ReturnCode         string `xml:"return_code"`
	ReturnMsg          string `xml:"return_msg"`
	Appid              string `xml:"appid"`
	MchId              string `xml:"mch_id"`
	NonceStr           string `xml:"nonce_str"`
	Sign               string `xml:"sign"`
	ResultCode         string `xml:"result_code"`
	ErrCode            string `xml:"err_code"`
	ErrCodeDes         string `xml:"err_code_des"`
	DeviceInfo         string `xml:"device_info"`
	Openid             string `xml:"openid"`
	IsSubscribe        string `xml:"is_subscribe"`
	TradeType          string `xml:"trade_type"`
	TradeState         string `xml:"trade_state"`
	BankType           string `xml:"bank_type"`
	TotalFee           int64  `xml:"total_fee"`
	FeeType            string `xml:"fee_type"`
	CashFee            int64  `xml:"cash_fee"`
	CashFeeType        string `xml:"cash_fee_type"`
	SettlementTotalFee int64  `xml:"settlement_total_fee"`
	CouponFee          int64  `xml:"coupon_fee"`
	CouponCount        int64  `xml:"coupon_count"`
	CouponId           string `xml:"coupon_id_$n"`
	CouponType         string `xml:"coupon_type_$n"`
	CouponFeeN         int64  `xml:"coupon_fee_$n"`
	TransactionId      string `xml:"transaction_id"`
	OutTradeNo         string `xml:"out_trade_no"`
	Attach             string `xml:"attach"`
	TimeEnd            string `xml:"time_end"`
	TradeStateDesc     string `xml:"trade_state_desc"`
}

func (q *queryResult) Param(key string) (interface{}, error) {
	var err error
	switch key {
	case "return_code":
		return q.ReturnCode, err
	case "return_msg":
		return q.ReturnMsg, err
	case "appid":
		return q.Appid, err
	case "mch_id":
		return q.MchId, err
	case "nonce_str":
		return q.NonceStr, err
	case "sign":
		return q.Sign, err
	case "result_code":
		return q.ResultCode, err
	case "err_code":
		return q.ErrCode, err
	case "err_code_des":
		return q.ErrCodeDes, err
	case "device_info":
		return q.DeviceInfo, err
	case "openid":
		return q.Openid, err
	case "is_subscribe":
		return q.IsSubscribe, err
	case "trade_type":
		return q.TradeType, err
	case "trade_state":
		return q.TradeState, err
	case "bank_type":
		return q.BankType, err
	case "total_fee":
		return q.TotalFee, err
	case "fee_type":
		return q.FeeType, err
	case "cash_fee":
		return q.CashFee, err
	case "cash_fee_type":
		return q.CashFeeType, err
	case "settlement_total_fee":
		return q.SettlementTotalFee, err
	case "coupon_fee":
		return q.CouponFee, err
	case "coupon_count":
		return q.CouponCount, err
	case "coupon_id_$n":
		return q.CouponId, err
	case "coupon_type_$n":
		return q.CouponType, err
	case "coupon_fee_$n":
		return q.CouponFeeN, err
	case "transaction_id":
		return q.TransactionId, err
	case "out_trade_no":
		return q.OutTradeNo, err
	case "attach":
		return q.Attach, err
	case "time_end":
		return q.TimeEnd, err
	case "trade_state_desc":
		return q.TradeStateDesc, err
	default:
		err = errors.New(fmt.Sprintf("invalid key: %s", key))
		return nil, err
	}
}

func (q queryResult) ListParam() Params {
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

func (m *myPayer) QueryOrder(param Params) (ResultParam, error) {
	if param == nil {
		return nil, errors.New("param is empty")
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//var paramNames []string
	var signType = e.SignTypeMD5 //此处默认MD5

	//校验订单号
	if _, ok := param["transaction_id"]; !ok {
		if _, ok = param["out_trade_no"]; !ok {
			return nil, errors.New("lack of order number")
		}
	}

	//校验其他参数
	for _, k := range queryMustParam {
		if param.Get(k) == nil {
			return nil, errors.New(fmt.Sprintf("need %s", k))
		}
	}

	for k := range param {
		if k == "sign" {
			continue
		}
		if k == "sign_type" {
			signType = param[k].(string)
		}
		if !util.HaveInArray(queryMustParam, k) && !util.HaveInArray(queryOptionalParam, k) {
			return nil, errors.New(fmt.Sprintf("no need %s param", k))
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

	var result *queryResult
	var request = &util.PostRequest{
		Body:        reader,
		Url:         queryApiUrl,
		ContentType: "application/xml;charset=utf-8",
	}

	err = util.PostToWx(request, result)
	if err != nil {
		return nil, err
	}

	if result.ReturnCode != "SUCCESS" {
		return nil, errors.New(result.ReturnMsg)
	}

	if result.ResultCode != "SUCCESS" {
		return nil, errors.New(result.ErrCodeDes)
	}

	sign, err = result.ListParam().Sign(signType)
	if err != nil || sign != result.Sign {
		return nil, e.ErrCheckSign
	}

	return result, nil
}

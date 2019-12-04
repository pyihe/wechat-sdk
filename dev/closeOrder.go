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
	关闭订单
*/

var (
	closeMustParam     = []string{"appid", "mch_id", "out_trade_no", "nonce_str", "sign"}
	closeOptionalParam = []string{"sign_type"}
)

const closeOrderUrl = "https://api.mch.weixin.qq.com/pay/closeorder"

type closeResult struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ResultMsg  string `xml:"result_msg"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

func (c *closeResult) Param(key string) (interface{}, error) {
	var err error
	switch key {
	case "return_code":
		return c.ReturnCode, err
	case "return_msg":
		return c.ReturnMsg, err
	case "appid":
		return c.Appid, err
	case "mch_id":
		return c.MchId, err
	case "nonce_str":
		return c.NonceStr, err
	case "sign":
		return c.Sign, err
	case "result_code":
		return c.ResultCode, err
	case "result_msg":
		return c.ResultMsg, err
	case "err_code":
		return c.ErrCode, err
	case "err_code_des":
		return c.ErrCodeDes, err
	default:
		err = errors.New(fmt.Sprintf("invalid key: %s", key))
		return nil, err
	}
}

func (c closeResult) ListParam() Params {
	p := make(Params)

	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)

	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).IsZero() {
			tagName := strings.Split(string(t.Field(i).Tag), "\"")[1]
			p[tagName] = v.Field(i).Interface()
		}
	}
	return p
}

//关闭订单
func (m *myPayer) CloseOrder(param Params) (ResultParam, error) {
	if param == nil {
		return nil, e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var signType = e.SignTypeMD5
	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	for _, v := range closeMustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New(fmt.Sprintf("need %s", v))
		}
	}
	for key := range param {
		if !util.HaveInArray(closeMustParam, key) && !util.HaveInArray(closeOptionalParam, key) {
			return nil, errors.New(fmt.Sprintf("no need %s", key))
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

	var result *closeResult
	var request = &util.PostRequest{
		Body:        reader,
		Url:         closeOrderUrl,
		ContentType: "application/xml;charset=utf-8",
	}
	err = util.PostToWx(request, &result)
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
	if sign != result.Sign {
		return nil, e.ErrCheckSign
	}
	return result, nil
}

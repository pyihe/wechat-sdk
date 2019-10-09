package wechat_sdk

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type myPayRe struct {
	Code       string `xml:"return_code"`
	Msg        string `xml:"return_msg"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrMsg     string `xml:"err_code_desc"`
	Payer      string `xml:"openid"`
	TradeType  string `xml:"trade_type"`
	TradeState string `xml:"trade_state"`
	TotalFee   string `xml:"total_fee"`
	EndTime    string `xml:"time_end"`
	TradeDesc  string `xml:"trade_state_desc"`
}

func (this *myPayRe) Param(p string) string {
	if this == nil {
		return ""
	}
	switch p {
	case "return_code":
		return this.Code
	case "return_msg":
		return this.Msg
	case "result_code":
		return this.ResultCode
	case "err_code":
		return this.ErrCode
	case "err_code_desc":
		return this.ErrMsg
	case "openid":
		return this.Payer
	case "trade_type":
		return this.TradeType
	case "trade_state":
		return this.TradeState
	case "total_fee":
		return this.TotalFee
	case "time_end":
		return this.EndTime
	case "trade_state_desc":
		return this.TradeDesc
	default:
		return ""
	}
}

func postQueryOrder(url string, contentType string, body io.Reader) (*myPayRe, error) {
	if body == nil {
		return nil, errors.New("body cannot be nil")
	}

	response, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("http StatusCode: %v", response.StatusCode))
	}

	byteContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result *myPayRe
	err = xml.Unmarshal(byteContent, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

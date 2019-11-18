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
	Code          string `xml:"return_code"`
	Msg           string `xml:"return_msg"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	ErrCode       string `xml:"err_code"`
	ErrMsg        string `xml:"err_code_desc"`
	Payer         string `xml:"openid"`
	TradeType     string `xml:"trade_type"`
	TradeState    string `xml:"trade_state"`
	TotalFee      string `xml:"total_fee"`
	EndTime       string `xml:"time_end"`
	TradeDesc     string `xml:"trade_state_desc"`
	TransactionId string `xml:"transaction_id"`
}

func (m *myPayRe) Param(p string) string {
	if m == nil {
		return ""
	}
	switch p {
	case "return_code":
		return m.Code
	case "return_msg":
		return m.Msg
	case "result_code":
		return m.ResultCode
	case "err_code":
		return m.ErrCode
	case "err_code_desc":
		return m.ErrMsg
	case "openid":
		return m.Payer
	case "trade_type":
		return m.TradeType
	case "trade_state":
		return m.TradeState
	case "total_fee":
		return m.TotalFee
	case "time_end":
		return m.EndTime
	case "trade_state_desc":
		return m.TradeDesc
	case "transaction_id":
		return m.TransactionId
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

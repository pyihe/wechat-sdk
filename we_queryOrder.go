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
	code       string `xml:"return_code"`
	msg        string `xml:"return_msg"`
	sign       string `xml:"sign"`
	resultCode string `xml:"result_code"`
	errCode    string `xml:"err_code"`
	errMsg     string `xml:"err_code_desc"`
	payer      string `xml:"openid"`
	tradeType  string `xml:"trade_type"`
	tradeState string `xml:"trade_state"`
	totalFee   string `xml:"total_fee"`
	endTime    string `xml:"time_end"`
	tradeDesc  string `xml:"trade_state_desc"`
}

func (this *myPayRe) Param(p string) string {
	if this == nil {
		return ""
	}
	switch p {
	case "return_code":
		return this.code
	case "return_msg":
		return this.msg
	case "result_code":
		return this.resultCode
	case "err_code":
		return this.errCode
	case "err_code_desc":
		return this.errMsg
	case "openid":
		return this.payer
	case "trade_type":
		return this.tradeType
	case "trade_state":
		return this.tradeState
	case "total_fee":
		return this.totalFee
	case "time_end":
		return this.endTime
	case "trade_state_desc":
		return this.tradeDesc
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

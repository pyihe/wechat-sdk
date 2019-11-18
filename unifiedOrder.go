package wechat_sdk

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type ResultParam interface {
	Param(p string) string
}

type maps map[string]interface{}

func (m maps) marshalXML(writer io.Writer) error {
	if writer == nil {
		return errors.New("writer cannot be nil")
	}
	var err error
	if _, err = io.WriteString(writer, "<xml>"); err != nil {
		return err
	}

	for k, v := range m {
		if _, err = io.WriteString(writer, "<"+k+">"); err != nil {
			return err
		}
		if err = xml.EscapeText(writer, []byte(fmt.Sprintf("%v", v))); err != nil {
			return err
		}
		if _, err = io.WriteString(writer, "</"+k+">"); err != nil {
			return err
		}
	}

	if _, err = io.WriteString(writer, "</xml>"); err != nil {
		return err
	}
	return err
}

func postUnifiedOrder(url string, contentType string, body io.Reader) (*unifiedorderReply, error) {
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

	var result *unifiedorderReply
	err = xml.Unmarshal(byteContent, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type unifiedorderReply struct {
	Code       string `xml:"return_code"`
	Msg        string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	TradeType  string `xml:"trade_type"`
	PrepayId   string `xml:"prepay_id"`
	MwebUrl    string `xml:"mweb_url"`
}

func (u *unifiedorderReply) Param(p string) string {
	switch p {
	case "return_code":
		return u.Code
	case "return_msg":
		return u.Msg
	case "appid":
		return u.AppId
	case "mch_id":
		return u.MchId
	case "device_info":
		return u.DeviceInfo
	case "nonce_str":
		return u.NonceStr
	case "sign":
		return u.Sign
	case "result_code":
		return u.ResultCode
	case "err_code":
		return u.ErrCode
	case "err_code_des":
		return u.ErrCodeDes
	case "trade_type":
		return u.TradeType
	case "prepay_id":
		return u.PrepayId
	default:
		return ""
	}
}

//HMAC-SHA256签名方式
func signHMACSHA256(s string) string {
	h := hmac.New(sha256.New, []byte(c.appSecret))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//MD5签名方式
func signMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

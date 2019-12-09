package dev

import (
	"encoding/base64"
	"errors"
	"io"
	"strings"

	"github.com/hong008/wechat-sdk/pkg/util"
)

//处理退款通知
//对结果中对req_info执行解密：
func (m *myPayer) RefundNotify(body io.Reader) (ResultParam, error) {
	result := ParseXMLReader(body)
	if len(result) == 0 {
		return nil, errors.New("reader has nothing")
	}

	var reqInfoStr string
	if reqInfoStr, _ = result.GetString("req_info"); reqInfoStr == "" {
		return nil, errors.New("wx response without req_info")
	}

	//1. 对加密串A做base64解码，得到加密串B
	reqInfo, err := base64.StdEncoding.DecodeString(reqInfoStr)
	if err != nil {
		return nil, err
	}

	//2. 对商户key做md5，得到32位小写key*
	md5Key := strings.ToLower(util.SignMd5(m.apiKey))
	//3. 用key*对加密串B做AES-256-ECB解密（PKCS7Padding）
	realData, err := util.AES256ECBDecrypt(reqInfo, []byte(md5Key))
	if err != nil {
		return nil, err
	}
	xmlReader := strings.NewReader(string(realData))
	reqData := ParseXMLReader(xmlReader)
	for k, v := range reqData {
		result[k] = v
	}
	delete(result, "req_info")
	return result, nil
}

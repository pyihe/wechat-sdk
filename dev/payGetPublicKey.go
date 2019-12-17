package dev

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
	uuid "github.com/satori/go.uuid"
)

/*
	获取RSA加密公钥
*/

func (m *myPayer) GetPublicKey(p12CertPath string, targetPath string) error {
	if m.mchId == "" {
		return e.ErrParams
	}

	cert, err := util.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return err
	}

	param := NewParam()
	u4 := uuid.NewV4()
	s := strings.Replace(fmt.Sprintf("%s", u4), "-", "", -1)
	nonceStr := s[:16]

	param.Add("mch_id", m.mchId)
	param.Add("nonce_str", nonceStr)
	param.Add("sign_type", e.SignTypeMD5)
	param.Add("sign", param.Sign(e.SignTypeMD5))

	reader, err := param.MarshalXML()
	if err != nil {
		return err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://fraud.mch.weixin.qq.com/risk/getpublickey",
		ContentType: e.PostContentType,
	}
	response, err := postToWxWithCert(request, cert)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return errors.New(returnMsg)
	}
	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return errors.New(errDes)
	}
	keyValue, err := result.GetString("pub_key")
	if err != nil {
		return err
	}
	//判断是否存在目标目录，如果不存在则创建
	if !strings.HasSuffix(targetPath, "/") {
		if targetPath == "" {
			targetPath = "./"
		} else {
			targetPath += "/"
		}
	}
	if err = util.MakeNewPath(targetPath); err != nil {
		return err
	}

	f, err := os.Create(targetPath + e.PublicKey)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(keyValue))
	if err != nil {
		return err
	}
	return nil
}

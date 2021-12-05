package v3

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pyihe/wechat-sdk/v3/vars"
)

func (we *WeChatClient) DownloadCertificates(savePath string) (vars.Kvs, error) {
	if we.serialNo == "" {
		return nil, vars.ErrNoSerialNo
	}
	if we.apiKey == "" {
		return nil, vars.ErrNoApiV3Key
	}
	var abUrl = "/v3/certificates"
	// 获取签名信息
	signParam, err := we.signSHA256WithRSA("GET", abUrl, nil)
	if err != nil {
		return nil, err
	}
	// 发起http请求
	response, err := we.requestWithSign("GET", we.apiDomain+abUrl, signParam, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	// 反序列化body到指定结构
	var data struct {
		Data []struct {
			EffectiveTime      time.Time `json:"effective_time"` // 证书生效时间
			ExpireTime         time.Time `json:"expire_time"`    // 证书过期时间
			SerialNo           string    `json:"serial_no"`      // 证书序列号
			EncryptCertificate struct {
				Algorithm      string `json:"algorithm"`       // 加密算法
				AssociatedData string `json:"associated_data"` // 附加数据包
				CipherText     string `json:"ciphertext"`      // 密文
				Nonce          string `json:"nonce"`           // 加密随机向量
			} `json:"encrypt_certificate"` // 加密信息
		} `json:"data"` // 返回的数据
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	//
	var result = vars.NewKvs()
	for _, encryptData := range data.Data {
		// 解密
		cipherText := encryptData.EncryptCertificate.CipherText
		associateData := encryptData.EncryptCertificate.AssociatedData
		nonce := encryptData.EncryptCertificate.Nonce
		plainText, err := we.decryptAEADAES256GCM(cipherText, associateData, nonce)
		if err != nil {
			return nil, err
		}

		var serialNo string
		var publicKey *rsa.PublicKey

		// 同步公钥信息到内存
		serialNo, publicKey, err = unmarshalPublicKey(plainText)
		if err != nil {
			return nil, err
		}
		if we.synchronizeTag {
			if we.publicKeys == nil {
				we.publicKeys = make(map[string]*rsa.PublicKey)
			}
			we.publicKeys[serialNo] = publicKey
		}
		result.Add(serialNo, publicKey)

		// 同步到本地
		fileName := fmt.Sprintf("public_key_%s.pem", encryptData.ExpireTime.Format("2006_01_02"))
		if err = writeToFile(savePath, fileName, plainText); err != nil {
			return nil, err
		}
	}
	return result, nil
}

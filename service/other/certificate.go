package other

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pyihe/go-pkg/maps"
	"github.com/pyihe/wechat-sdk/model/other"
	"github.com/pyihe/wechat-sdk/pkg/files"
	"github.com/pyihe/wechat-sdk/pkg/rsas"
	"github.com/pyihe/wechat-sdk/service"
	"github.com/pyihe/wechat-sdk/vars"
)

// DownloadCertificates 下载证书API
// 请求参数说明:
// savePath: 文件存放路径
// 返回参数:
// requestId: API请求唯一ID
// certs: key为证书序列号, value为*rsa.PublicKey
// err: error
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay5_1.shtml
func DownloadCertificates(config *service.Config, savePath string) (certsResponse *other.CertificateResponse, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.ApiKey == "" {
		err = vars.ErrNoApiV3Key
		return
	}
	// 发起带签名的请求
	response, err := service.RequestWithSign(config, "GET", "/v3/certificates", nil)
	if err != nil {
		return
	}

	certsResponse = new(other.CertificateResponse)
	certsResponse.RequestId = response.Header.Get("Request-ID")
	certsResponse.Certificates = maps.NewParam()
	fmt.Printf("resp: %+v\n", *certsResponse)
	content, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	if err = json.Unmarshal(content, &certsResponse); err != nil {
		return
	}
	for _, encryptData := range certsResponse.Data {
		// 解密
		cipherText := encryptData.EncryptCertificate.CipherText
		associateData := encryptData.EncryptCertificate.AssociatedData
		nonce := encryptData.EncryptCertificate.Nonce
		var serialNo string
		var plainText []byte
		var certificate *x509.Certificate

		fmt.Printf("apiKey: %v\n", config.ApiKey)
		fmt.Printf("cipherText: %s\n", cipherText)
		fmt.Printf("assData: %v\n", associateData)
		fmt.Printf("nonce: %v\n", nonce)
		plainText, err = rsas.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
		if err != nil {
			fmt.Printf("xxxx: %v\n", err)
			return
		}

		// 同步公钥信息到内存
		serialNo, certificate, err = files.UnmarshalCertificate(plainText)
		if err != nil {
			return
		}
		certsResponse.Certificates.Add(serialNo, certificate)
		if config.SyncCertificateTag {
			config.Certificates.Add(serialNo, certificate.PublicKey.(*rsa.PublicKey))
		}
		// 同步到本地
		fileName := fmt.Sprintf("public_key_%s.pem", encryptData.ExpireTime.Format("2006_01_02"))
		if err = files.WritToFile(savePath, fileName, plainText); err != nil {
			return
		}
	}
	return
}

package certificate

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pyihe/wechat-sdk/v3/pkg"
	"github.com/pyihe/wechat-sdk/v3/pkg/errors"

	"github.com/pyihe/wechat-sdk/v3/pkg/aess"
	"github.com/pyihe/wechat-sdk/v3/pkg/files"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// DownloadCertificates 下载证书API
// 请求参数说明:
// savePath: 文件存放路径
// 返回参数:
// requestId: API请求唯一ID
// certs: key为证书序列号, value为*rsa.PublicKey
// err: error
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay5_1.shtml
func DownloadCertificates(config *service.Config, savePath string) (certsResponse *CertResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if config.GetApiKey() == "" {
		err = errors.ErrNoApiV3Key
		return
	}
	// 发起带签名的请求
	response, err := config.RequestWithSign(http.MethodGet, "/v3/certificates", nil)
	if err != nil {
		return
	}
	certsResponse = new(CertResponse)
	certsResponse.RequestId = response.Header.Get("Request-ID")
	certsResponse.Certificates = pkg.NewParam()
	content, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	if err = json.Unmarshal(content, &certsResponse); err != nil {
		return
	}

	var cipher = config.GetMerchantCipher()
	var key = config.GetApiKey()
	var syncTag = config.GetSyncCertificateTag()

	for _, encryptData := range certsResponse.Data {
		// 解密
		cipherText := encryptData.EncryptCertificate.CipherText
		associateData := encryptData.EncryptCertificate.AssociatedData
		nonce := encryptData.EncryptCertificate.Nonce
		var serialNo string
		var plainText []byte
		var certificate *x509.Certificate

		plainText, err = aess.DecryptAEADAES256GCM(cipher, key, cipherText, associateData, nonce)
		if err != nil {
			return
		}

		// 同步公钥信息到内存
		serialNo, certificate, err = files.UnmarshalCertificate(plainText)
		if err != nil {
			return
		}
		certsResponse.Certificates.Add(serialNo, certificate)
		if syncTag {
			config.AddCertificate(serialNo, certificate)
		}
		// 同步到本地
		fileName := fmt.Sprintf("public_key_%s.pem", encryptData.ExpireTime.Format("2006_01_02"))
		if err = files.WritToFile(savePath, fileName, plainText); err != nil {
			return
		}
	}
	return
}

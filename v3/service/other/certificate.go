package other

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pyihe/go-pkg/maps"
	"github.com/pyihe/wechat-sdk/v3/pkg/files"
	"github.com/pyihe/wechat-sdk/v3/pkg/rsas"
	"github.com/pyihe/wechat-sdk/v3/service"
	"github.com/pyihe/wechat-sdk/v3/vars"
)

// DownloadCertificates 下载证书API
// 请求参数说明:
// savePath: 文件存放路径
// 返回参数:
// requestId: API请求唯一ID
// certs: key为证书序列号, value为*rsa.PublicKey
// err: error
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay5_1.shtml
func DownloadCertificates(config *service.Config, savePath string) (requestId string, certs maps.Param, err error) {
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

	requestId = response.Header.Get("Request-ID")
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
	content, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	if err = json.Unmarshal(content, &data); err != nil {
		return
	}

	certs = maps.NewParam()
	for _, encryptData := range data.Data {
		// 解密
		cipherText := encryptData.EncryptCertificate.CipherText
		associateData := encryptData.EncryptCertificate.AssociatedData
		nonce := encryptData.EncryptCertificate.Nonce
		var serialNo string
		var plainText []byte
		var publicKey *rsa.PublicKey

		plainText, err = rsas.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
		if err != nil {
			return
		}

		// 同步公钥信息到内存
		serialNo, publicKey, err = files.UnmarshalRSAPublicKey(plainText)
		if err != nil {
			return
		}

		certs.Add(serialNo, publicKey)
		if config.SyncCertificateTag {
			config.Certificates.Add(serialNo, publicKey)
		}
		// 同步到本地
		fileName := fmt.Sprintf("public_key_%s.pem", encryptData.ExpireTime.Format("2006_01_02"))
		if err = files.WritToFile(savePath, fileName, plainText); err != nil {
			return
		}
	}
	return
}

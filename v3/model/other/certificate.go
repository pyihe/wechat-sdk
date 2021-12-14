package other

import (
	"time"

	"github.com/pyihe/go-pkg/maps"
	"github.com/pyihe/wechat-sdk/v3/model"
)

// Resource 微信返回的证书数据，包括加密数据
type Resource struct {
	EffectiveTime      time.Time           `json:"effective_time"`                // 证书生效时间
	ExpireTime         time.Time           `json:"expire_time"`                   // 证书过期时间
	SerialNo           string              `json:"serial_no"`                     // 证书序列号
	EncryptCertificate *EncryptCertificate `json:"encrypt_certificate,omitempty"` // 加密数据
}

// EncryptCertificate 微信返回的加密数据
type EncryptCertificate struct {
	Algorithm      string `json:"algorithm"`       // 加密算法
	AssociatedData string `json:"associated_data"` // 附加数据包
	CipherText     string `json:"ciphertext"`      // 密文
	Nonce          string `json:"nonce"`           // 加密随机向量
}

// CertificateResponse 用于证书下载API返回的数据
type CertificateResponse struct {
	model.WechatError             // 如果请求失败，微信返回的错误信息将存储在这里
	RequestId         string      `json:"request_id,omitempty"`   // 存放请求的唯一ID
	Data              []*Resource `json:"data,omitempty"`         // 微信返回的原始数据
	Certificates      maps.Param  `json:"certificates,omitempty"` // 解密后的证书, key=证书序列号; value=证书*x509.Certificate
}

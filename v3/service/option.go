package service

import (
	"net/http"

	"github.com/pyihe/wechat-sdk/v3/model/pay/combine"

	"github.com/pyihe/wechat-sdk/v3/model/pay/merchant"

	"github.com/pyihe/go-pkg/maps"
	"github.com/pyihe/secret"
	"github.com/pyihe/wechat-sdk/v3/pkg/files"
	"github.com/pyihe/wechat-sdk/v3/vars"
)

type Option func(*Config)

type Config struct {
	// v3版本API域名公共部分
	domain string

	// 标志是否在下载证书的同时，将证书加载同步更新到内存Config中，在证书更替时，如果不同步到内存的话，以后的方法调用需要手动更换证书信息
	SyncCertificateTag bool

	// 服务商应用ID（服务商申请的公众号appid）
	//SpAppId string
	//
	//// 服务商户号
	//SpMchId string
	//
	//// 子商户应用ID
	//SubAppId string
	//
	//// 子商户号
	//SubMchId string

	// 商户在微信平台的唯一ID
	AppId string

	// 商户号, 包括直连商户、服务商或渠道商的商户号mchid，用于生成签名
	MchId string

	// app secret
	Secret string

	// v3 key
	ApiKey string

	// 商户API证书序列号
	SerialNo string

	// http client
	HttpClient *http.Client

	// 用于签名与验证签名
	Cipher secret.Cipher

	// 服务类型, 商户平台或者服务商平台
	Platform vars.Platform

	// 微信平台公钥证书, key为serialNo, value为*rsa.PublicKey
	Certificates maps.Param

	// 处理商户平台支付通知的handler
	PrepayNotifyHandler func(prepayOrder *merchant.PrepayOrder) error

	// 处理商户平台退款通知的handler
	RefundNotifyHandler func(refundOrder *merchant.RefundOrder) error

	// 处理商户平台合单支付通知的handler
	CombinePrepayNotifyHandler func(prepayOrder *combine.PrepayOrder) error
}

func NewConfig(opts ...Option) *Config {
	var c = &Config{
		domain:     "https://api.mch.weixin.qq.com",
		HttpClient: http.DefaultClient,
		Cipher:     secret.NewCipher(),
	}
	for _, op := range opts {
		op(c)
	}

	c.init()

	return c
}

func (c *Config) init() {
	if c.Platform.Valid() == false {
		panic("请选择正确的服务platform!")
	}
}

func WithSyncCertificate() Option {
	return func(config *Config) {
		config.SyncCertificateTag = true
	}
}

func WithPlatform(platform vars.Platform) Option {
	return func(config *Config) {
		config.Platform = platform
	}
}

//func WithSpAppId(spAppId string) Option {
//	return func(config *Config) {
//		config.SpAppId = spAppId
//	}
//}
//
//func WithSpMchId(spMchId string) Option {
//	return func(config *Config) {
//		config.SpMchId = spMchId
//	}
//}
//
//func WithSubAppId(subAppId string) Option {
//	return func(config *Config) {
//		config.SubAppId = subAppId
//	}
//}
//
//func WithSubMchId(subMchId string) Option {
//	return func(config *Config) {
//		config.SubMchId = subMchId
//	}
//}

func WithAppId(appId string) Option {
	return func(config *Config) {
		config.AppId = appId
	}
}

func WithMchId(mchId string) Option {
	return func(config *Config) {
		config.MchId = mchId
	}
}

func WithAppSecret(secret string) Option {
	return func(config *Config) {
		config.Secret = secret
	}
}

func WithApiV3Key(apiKey string) Option {
	return func(config *Config) {
		config.ApiKey = apiKey
	}
}

func WithSerialNo(serialNo string) Option {
	return func(config *Config) {
		config.SerialNo = serialNo
	}
}

func WithPrivateKey(file string) Option {
	return func(config *Config) {
		privateKey, err := files.LoadRSAPrivateKey(file)
		if err != nil {
			panic(err)
		}
		if err = config.Cipher.SetRSAPrivateKey(privateKey, secret.PKCSLevel8); err != nil {
			panic(err)
		}
	}
}

func WithPublicKey(file string) Option {
	return func(config *Config) {
		serialNo, publicKey, err := files.LoadRSAPublicKeyWithSerialNo(file)
		if err != nil {
			panic(err)
		}
		if err = config.Cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8); err != nil {
			panic(err)
		}
		if config.Certificates == nil {
			config.Certificates = maps.NewParam()
		}
		config.Certificates.Add(serialNo, publicKey)
	}
}

func WithHttpClient(client *http.Client) Option {
	return func(config *Config) {
		config.HttpClient = client
	}
}

func WithPrepayNotifyHandler(handler func(response *merchant.PrepayOrder) error) Option {
	return func(config *Config) {
		config.PrepayNotifyHandler = handler
	}
}

func WithRefundNotifyHandler(handler func(refundOrder *merchant.RefundOrder) error) Option {
	return func(config *Config) {
		config.RefundNotifyHandler = handler
	}
}

func WithCombinePrepayNotifyHandler(handler func(order *combine.PrepayOrder) error) Option {
	return func(config *Config) {
		config.CombinePrepayNotifyHandler = handler
	}
}

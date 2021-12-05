package v3

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"

	"github.com/pyihe/wechat-sdk/v3/vars"

	"github.com/pyihe/secret"
)

type keyFiles struct {
	privateFile string // 商户私钥文件路径，用于签名
	publicFile  string // 微信支付平台公钥，用于验证签名、加密信息等
}

// WeChatClient 发起请求
type WeChatClient struct {
	/*是否在下载公钥证书的同时将公钥同步到内存publicKeys中，为true时，调用方每次调用接口时，都会同步更新publicKeys的内容，否则将不会更新。
	添加synchronizeTag的原因是考虑到证书过期，不同步更新至内存的话，外层调用需要重新实例化一个WeChatClient，否则将无法在使用证书相关的接口*/
	synchronizeTag bool

	// v3版本API域名公共部分
	apiDomain string

	// 商户在微信平台的唯一ID
	appId string

	// 商户号
	mchId string

	// app secret
	secret string

	// v3 key
	apiKey string

	// 商户API证书序列号
	serialNo string

	// 密钥文件路径
	keyFiles *keyFiles

	// http client
	httpClient *http.Client

	// 用于签名与验证签名
	cipher secret.Cipher

	// 这里保存微信平台的公钥信息，key为证书序列号serial_no，value为*rsa.PublicKey
	publicKeys map[string]*rsa.PublicKey

	// 预支付API异步通知处理方法
	prepayNotifyFn func(order *vars.OrderResponse) error

	// 退款异步通知处理方法
	refundNotifyFn func(order *vars.OrderResponse) error
}

type Options func(client *WeChatClient)

func NewWechatPayer(appId string, options ...Options) *WeChatClient {
	var client = &WeChatClient{
		apiDomain:  "https://api.mch.weixin.qq.com",
		appId:      appId,
		httpClient: http.DefaultClient,
		cipher:     secret.NewCipher(),
	}
	for _, opt := range options {
		opt(client)
	}

	client.init()

	return client
}

func (we *WeChatClient) init() {
	if we.keyFiles != nil {
		if privateFile := we.keyFiles.privateFile; privateFile != "" {
			//  这里加载商户私钥，如果有的话
			if err := we.cipher.SetRSAPrivateKey(privateFile, secret.PKCSLevel8); err != nil {
				// 抱歉，这里如果出错的话直接panic
				panic(err)
			}
		}
		if publicFile := we.keyFiles.publicFile; publicFile != "" {
			if we.publicKeys == nil {
				we.publicKeys = make(map[string]*rsa.PublicKey)
			}
			data, err := ioutil.ReadFile(publicFile)
			if err != nil {
				panic(err)
			}
			serialNo, publicKey, err := unmarshalPublicKey(data)
			if err != nil {
				// 同样，如果这里加载出错，直接panic
				panic(err)
			}
			we.publicKeys[serialNo] = publicKey
		}
		// 清空证书文件信息
		we.keyFiles = nil
	}
}

func WithMchId(mchId string) Options {
	return func(client *WeChatClient) {
		client.mchId = mchId
	}
}

func WithSecret(secret string) Options {
	return func(client *WeChatClient) {
		client.secret = secret
	}
}

func WithV3Key(key string) Options {
	return func(client *WeChatClient) {
		client.apiKey = key
	}
}

func WithSerialNo(serialNo string) Options {
	return func(client *WeChatClient) {
		client.serialNo = serialNo
	}
}

func WithPrivateKey(file string) Options {
	return func(client *WeChatClient) {
		if client.keyFiles == nil {
			client.keyFiles = &keyFiles{}
		}
		client.keyFiles.privateFile = file
	}
}

func WithPublicKey(file string) Options {
	return func(client *WeChatClient) {
		if client.keyFiles == nil {
			client.keyFiles = &keyFiles{}
		}
		client.keyFiles.publicFile = file
	}
}

func WithSynchronize(b bool) Options {
	return func(client *WeChatClient) {
		client.synchronizeTag = b
	}
}

func WithPrepayNotifyHandler(fn func(response *vars.OrderResponse) error) Options {
	return func(client *WeChatClient) {
		client.prepayNotifyFn = fn
	}
}

func WithRefundNotifyHandler(fn func(response *vars.OrderResponse) error) Options {
	return func(client *WeChatClient) {
		client.refundNotifyFn = fn
	}
}

package v3

import "net/http"

// 支付客户端
type weChatClient struct {
	httpClient *http.Client // http client
	appId      string       // appid
	mchId      string       // 商户号
	secret     string       // app secret
	apiKey     string       // v3 key
	serialNo   string       // 商户API证书序列号
}

type Options func(client *weChatClient)

func NewWechatPayer(appId string, options ...Options) {
	var client = &weChatClient{
		httpClient: http.DefaultClient,
		appId:      appId,
	}
	for _, opt := range options {
		opt(client)
	}
}

func WithHttpClient(hClient *http.Client) Options {
	return func(client *weChatClient) {
		client.httpClient = hClient
	}
}

func WithMchId(mchId string) Options {
	return func(client *weChatClient) {
		client.mchId = mchId
	}
}

func WithSecret(secret string) Options {
	return func(client *weChatClient) {
		client.secret = secret
	}
}

func WithV3Key(key string) Options {
	return func(client *weChatClient) {
		client.apiKey = key
	}
}

func WithSerialNo(serialNo string) Options {
	return func(client *weChatClient) {
		client.serialNo = serialNo
	}
}

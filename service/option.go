package service

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/maps"
	"github.com/pyihe/go-pkg/rands"
	"github.com/pyihe/secret"
	"github.com/pyihe/wechat-sdk/v3/model"
	"github.com/pyihe/wechat-sdk/v3/pkg/aess"
	"github.com/pyihe/wechat-sdk/v3/pkg/files"
	"github.com/pyihe/wechat-sdk/v3/pkg/rsas"
)

const (
	ContentTypeJSON = "application/json"
	PostContentType = "application/xml;charset=utf-8"
)

var (
	ErrNoAppId           = errors.New("请提供appid!")
	ErrNoRequest         = errors.New("请求参数为空!")
	ErrNoSecret          = errors.New("请提供secret!")
	ErrNoSerialNo        = errors.New("请提供商户证书序列号!")
	ErrNoMchId           = errors.New("请提供商户号!")
	ErrNoApiV3Key        = errors.New("请提供商户API密钥!")
	ErrInvalidSessionKey = errors.New("获取session_key失败!")
	ErrRequestAgain      = errors.New("请稍后再次请求!")
	ErrInitConfig        = errors.New("请初始化config!")
	ErrInvalidResource   = errors.New("未获取到通知资源数据!")
	ErrInvalidHashType   = errors.New("暂不支持的哈希类型!")
)

type Option func(*Config)

func WithSyncCertificate() Option {
	return func(config *Config) {
		config.syncCertificateTag = true
	}
}

func WithAppId(appId string) Option {
	return func(config *Config) {
		config.appId = appId
	}
}

func WithMchId(mchId string) Option {
	return func(config *Config) {
		config.mchId = mchId
	}
}

func WithAppSecret(secret string) Option {
	return func(config *Config) {
		config.secret = secret
	}
}

func WithApiV3Key(apiKey string) Option {
	return func(config *Config) {
		config.apiKey = apiKey
	}
}

func WithSerialNo(serialNo string) Option {
	return func(config *Config) {
		config.serialNo = serialNo
	}
}

func WithPrivateKey(file string) Option {
	return func(config *Config) {
		privateKey, err := files.LoadRSAPrivateKey(file)
		if err != nil {
			panic(err)
		}
		if err = config.cipher.SetRSAPrivateKey(privateKey, secret.PKCSLevel8); err != nil {
			panic(err)
		}
	}
}

func WithPublicKey(file string) Option {
	return func(config *Config) {
		cert, err := files.LoadCertificate(file)
		if err != nil {
			panic(err)
		}
		publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			panic("加载证书失败: 请确认证书是否为RSA PublicKey!")
		}
		serialNo := cert.SerialNumber.Text(16)
		if err = config.cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8); err != nil {
			panic(err)
		}
		config.certificates.Add(serialNo, publicKey)
	}
}

func WithHttpClient(client *http.Client) Option {
	return func(config *Config) {
		config.httpClient = client
	}
}

type Config struct {
	// v3版本API域名公共部分
	domain string

	// 标志是否在下载证书的同时，将证书加载同步更新到内存Config中，在证书更替时，如果不同步到内存的话，以后的方法调用需要手动更换证书信息
	syncCertificateTag bool

	// 商户在微信平台的唯一ID
	appId string

	// 商户号, 包括直连商户、服务商或渠道商的商户号mchid，用于生成签名
	mchId string

	// app secret
	secret string

	// v3 key
	apiKey string

	// 商户API证书序列号
	serialNo string

	// http client
	httpClient *http.Client

	// 用于签名与验证签名
	cipher secret.Cipher

	// 用于验证hash值
	hasher secret.Hasher

	// 微信平台公钥证书, key为serialNo, value为*rsa.PublicKey
	certificates maps.Param
}

func NewConfig(opts ...Option) *Config {
	var c = &Config{
		domain:       "https://api.mch.weixin.qq.com",
		httpClient:   http.DefaultClient,
		cipher:       secret.NewCipher(),
		hasher:       secret.NewHasher(),
		certificates: maps.NewParam(),
	}
	for _, op := range opts {
		op(c)
	}
	return c
}

func (c *Config) GetDomain() string {
	return c.domain
}

func (c *Config) GetSyncCertificateTag() bool {
	return c.syncCertificateTag
}

func (c *Config) GetAppId() string {
	return c.appId
}

func (c *Config) GetMchId() string {
	return c.mchId
}

func (c *Config) GetSecret() string {
	return c.secret
}

func (c *Config) GetApiKey() string {
	return c.apiKey
}

func (c *Config) GetSerialNo() string {
	return c.serialNo
}

func (c *Config) GetHTTPClient() *http.Client {
	return c.httpClient
}

func (c *Config) GetCipher() secret.Cipher {
	return c.cipher
}

func (c *Config) GetCertificates() maps.Param {
	return c.certificates
}

// RequestWithSign 对发送给微信服务器的body进行SHA-256 with RSA签名, 返回*http.Request
// 参数说明:
// method: api方法类型, 如: GET、POST等
// url: api接口除去域名的绝对URL, 如: /v3/pay/transactions/jsapi
// body: 请求主体，比如支付时为支付参数，调用方需要不序列化
// 返回参数说明:
// signResult: 返回用于签名的各个参数，包括签名结果
// 签名介绍详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_0.shtml
func (c *Config) RequestWithSign(method, url string, body interface{}) (response *http.Response, err error) {
	if c.mchId == "" {
		err = ErrNoMchId
		return
	}
	if c.serialNo == "" {
		err = ErrNoSerialNo
		return
	}
	// 构造签名主体
	data, err := marshalJSON(body)
	if err != nil {
		return
	}

	method = strings.ToUpper(method) // 方法类型，转为大写
	timestamp := time.Now().Unix()   // 时间戳
	nonceStr := rands.String(32)     // 随机字符串

	source := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, url, timestamp, nonceStr, string(data))
	signature, err := rsas.SignSHA256WithRSA(c.cipher, source)
	if err != nil {
		return
	}
	// 签名头
	signatureHead := fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%d\",serial_no=\"%s\"", c.mchId, nonceStr, signature, timestamp, c.serialNo)
	request, err := http.NewRequest(method, c.domain+url, ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return
	}
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", "WECHATPAY2-SHA256-RSA2048", signatureHead))
	request.Header.Set("Content-Type", ContentTypeJSON)
	request.Header.Set("Accept", ContentTypeJSON)
	request.Header.Set("User-Agent", "Pyihe-Wechat-SDK")
	request.Header.Set("Accept-Language", "zh-CN")
	return c.httpClient.Do(request)
}

// Request 发起普通的HTTP请求
func (c *Config) Request(method, url, contentType string, data interface{}) (response *http.Response, err error) {
	body, err := marshalJSON(data)
	if err != nil {
		return
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Accept", contentType)
	request.Header.Set("User-Agent", "Pyihe-Wechat-SDK")
	request.Header.Set("Accept-Language", "zh-CN")
	return c.httpClient.Do(request)
}

// ParseWechatResponse 验证向微信服务器发送请求后从微信得到的应答签名
// 微信(验证)签名验证详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_1.shtml
func (c *Config) ParseWechatResponse(response *http.Response, dst interface{}) (requestId string, err error) {
	if response == nil {
		err = errors.New("response为空!")
		return
	}
	var body []byte
	var header = response.Header
	var code = response.StatusCode // 根据http code 判断请求是否成功

	// 获取唯一请求ID
	requestId = header.Get("Request-ID")
	// 请求已经被接收，但尚未处理，需要重复请求一遍
	if code == http.StatusAccepted {
		err = ErrRequestAgain
		return
	}
	// 请求失败
	if code != http.StatusOK && code != http.StatusNoContent {
		err = errors.NewWithCode("see errcode for detail.", errors.ErrorCode(code))
		return
	}
	if code == http.StatusOK {
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
		_ = response.Body.Close()
	}
	// 1. 验证证书序列号是否正确
	serialNo := header.Get("Wechatpay-Serial")
	publicKey, ok := c.certificates[serialNo].(*rsa.PublicKey)
	if !ok {
		err = errors.New("inconsistent Wechatpay-Serial.")
		return
	}

	// 2. 获取微信的签名结果
	wechatSign := header.Get("Wechatpay-Signature") // 微信签名
	// 3. 获取签名参数
	timestamp := header.Get("Wechatpay-Timestamp") // 时间戳
	nonceStr := header.Get("Wechatpay-Nonce")      // 随机字符串

	// 4. 构造原始的签名数据
	plainTxt := fmt.Sprintf("%v\n%v\n%v\n", timestamp, nonceStr, string(body))

	// 验证签名
	_ = c.cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8)
	err = rsas.VerifySHA256WithRSA(c.cipher, wechatSign, plainTxt)
	if err != nil {
		return
	}
	err = unmarshalJSON(body, dst)
	return
}

// ParseWechatNotify 验证微信服务器的通知，预支付、退款等请求后，微信回调的Request同样需要签名验证
// 微信（验证）签名方法详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_1.shtml
func (c *Config) ParseWechatNotify(request *http.Request, dst interface{}) (id string, err error) {
	if request == nil {
		err = ErrNoRequest
		return
	}
	if c.apiKey == "" {
		err = ErrNoApiV3Key
		return
	}

	var body []byte
	var header = request.Header

	body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}
	_ = request.Body.Close()

	// 1. 验证证书序列号是否正确
	serialNo := header.Get("Wechatpay-Serial")
	publicKey, ok := c.certificates[serialNo].(*rsa.PublicKey)
	if !ok {
		err = errors.New("inconsistent Wechatpay-Serial.")
		return
	}

	// 2. 获取微信的签名结果
	wechatSign := header.Get("Wechatpay-Signature") // 微信签名
	// 3. 获取签名参数
	timestamp := header.Get("Wechatpay-Timestamp") // 时间戳
	nonceStr := header.Get("Wechatpay-Nonce")      // 随机字符串

	// 4. 构造原始的签名数据
	plainTxt := fmt.Sprintf("%v\n%v\n%v\n", timestamp, nonceStr, string(body))

	// 验证签名
	_ = c.cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8)
	err = rsas.VerifySHA256WithRSA(c.cipher, wechatSign, plainTxt)
	if err != nil {
		return
	}

	// 签名通过的话反序列化body到结构体中
	notifyResponse := new(model.WechatNotifyResponse)
	if err = unmarshalJSON(body, &notifyResponse); err != nil {
		return
	}

	// 判断资源类型
	if notifyResponse.ResourceType != "encrypt-resource" {
		err = errors.New("错误的资源类型: " + notifyResponse.ResourceType)
		return
	}
	if notifyResponse.Resource == nil {
		err = ErrInvalidResource
		return
	}

	// 解密
	cipherText := notifyResponse.Resource.CipherText
	associateData := notifyResponse.Resource.AssociatedData
	nonce := notifyResponse.Resource.Nonce
	plainData, err := aess.DecryptAEADAES256GCM(c.cipher, c.apiKey, cipherText, associateData, nonce)
	if err != nil {
		return
	}
	err = unmarshalJSON(plainData, dst)
	return
}

// Download 下载URL对应的数据流
func (c *Config) Download(url string) (data []byte, err error) {
	if strs := strings.Split(url, c.domain); len(strs) > 1 {
		url = strs[1]
	}
	response, err := c.RequestWithSign(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	defer response.Body.Close()
	data, err = ioutil.ReadAll(response.Body)
	return
}

// VerifyHashValue 校验hash值
func (c *Config) VerifyHashValue(hashType crypto.Hash, data interface{}, hashValue string) (err error) {
	v, err := c.hasher.HashToString(data, hashType)
	if err != nil {
		return err
	}
	if strings.ToUpper(hashValue) != strings.ToUpper(v) {
		err = errors.New("HASH校验不通过!")
	}
	return
}

func unmarshalJSON(data []byte, dst interface{}) (err error) {
	if len(data) == 0 {
		return
	}
	if dst == nil {
		err = errors.New("反序列化失败: dst不能为nil!")
		return
	}
	err = json.Unmarshal(data, &dst)
	return
}

func marshalJSON(data interface{}) (bytes []byte, err error) {
	dataValue := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data).Kind()

	if dataValue.IsZero() {
		return
	}
	switch dataType {
	case reflect.String:
		bytes = []byte(dataValue.String())
	case reflect.Slice:
		if dataValue.Elem().Kind() != reflect.Uint8 {
			err = errors.New("请传入字节切片!")
			break
		}
		bytes = dataValue.Bytes()
	case reflect.Struct, reflect.Ptr:
		bytes, err = json.Marshal(data)
	default:
		err = errors.New("序列化失败, 数据类型不支持: " + dataType.String())
	}
	return
}

package service

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/pyihe/secret"

	"github.com/pyihe/wechat-sdk/v3/model"
	"github.com/pyihe/wechat-sdk/v3/pkg"
	"github.com/pyihe/wechat-sdk/v3/pkg/aess"
	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/pkg/files"
	"github.com/pyihe/wechat-sdk/v3/pkg/rsas"
)

const (
	ContentTypeJSON = "application/json"
	ContentTypeXML  = "application/xml;charset=utf-8"
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

func WithPrivateKey(file string, level secret.PKCSLevel) Option {
	return func(config *Config) {
		if err := config.merchantCipher.SetRSAPrivateKey(file, level); err != nil {
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
		serialNo := strings.ToUpper(cert.SerialNumber.Text(16))
		if err = config.wechatCipher.SetRSAPublicKey(publicKey, 0); err != nil {
			panic(err)
		}
		config.certificates.Add(serialNo, cert)
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

	// 包含商户平台证书的加密器，用于签名、签名验证、解密
	merchantCipher secret.Cipher

	// 包含微信支付平台公钥信息的加密器，用于加密上载信息
	wechatCipher secret.Cipher

	// 用于验证hash值
	hasher secret.Hasher

	// 微信平台公钥证书, key为serialNo, value为*x509.Certificate
	certificates pkg.Param
}

func NewConfig(opts ...Option) *Config {
	var c = &Config{
		domain:         "https://api.mch.weixin.qq.com",
		httpClient:     http.DefaultClient,
		merchantCipher: secret.NewCipher(),
		wechatCipher:   secret.NewCipher(),
		hasher:         secret.NewHasher(),
		certificates:   pkg.NewParam(),
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

func (c *Config) GetMerchantCipher() secret.Cipher {
	return c.merchantCipher
}

func (c *Config) GetWechatCipher() secret.Cipher {
	return c.wechatCipher
}

func (c *Config) AddCertificate(serialNo string, cert *x509.Certificate) {
	serialNo = strings.ToUpper(serialNo)
	c.certificates.Add(serialNo, cert)
}

func (c *Config) GetValidPublicKey() (serialNo string, publicKey *rsa.PublicKey) {
	now := time.Now()
	c.certificates.Range(func(key string, value interface{}) (breakOut bool) {
		data, ok := value.(*x509.Certificate)
		if !ok || data == nil {
			return
		}
		// 如果证书还没开始生效，
		if now.Before(data.NotBefore) {
			return
		}
		// 如果已经过期
		if now.After(data.NotAfter) {
			c.certificates.Delete(key)
			return
		}
		serialNo = key
		publicKey = data.PublicKey.(*rsa.PublicKey)
		breakOut = true
		return
	})
	return
}

func (c *Config) GetRSAPublicKey(serialNo string) *rsa.PublicKey {
	serialNo = strings.ToUpper(serialNo)
	data, ok := c.certificates.Get(serialNo)
	if !ok || data == nil {
		return nil
	}
	certs, ok := data.(*x509.Certificate)
	if !ok || certs == nil {
		return nil
	}

	publicKey, ok := certs.PublicKey.(*rsa.PublicKey)
	if !ok || publicKey == nil {
		return nil
	}
	return publicKey
}

func (c *Config) agent() string {
	return fmt.Sprintf("Pyihe-Wechat-SDK With GO(%s)/%s", runtime.Version(), runtime.GOOS)
}

// RequestWithSign 对发送给微信服务器的body进行SHA-256 with RSA签名, 返回*http.Request
// 参数说明:
// method: api方法类型, 如: GET、POST等
// url: api接口除去域名的绝对URL, 如: /v3/pay/transactions/jsapi
// body: 请求主体，比如支付时为支付参数，调用方需要不序列化
// 返回参数说明:
// signResult: 返回用于签名的各个参数，包括签名结果
// 签名介绍详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_0.shtml
func (c *Config) RequestWithSign(method, url string, body interface{}, headers ...string) (response *http.Response, err error) {
	if c.mchId == "" {
		err = errors.ErrNoMchId
		return
	}
	if c.serialNo == "" {
		err = errors.ErrNoSerialNo
		return
	}
	// 构造签名主体
	data, err := marshalJSON(body)
	if err != nil {
		fmt.Println(11)
		return
	}

	method = strings.ToUpper(method) // 方法类型，转为大写
	timestamp := time.Now().Unix()   // 时间戳
	nonceStr := pkg.String(32)       // 随机字符串

	source := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, url, timestamp, nonceStr, string(data))
	signature, err := rsas.SignSHA256WithRSA(c.merchantCipher, source)
	if err != nil {
		fmt.Println(22)
		return
	}
	// 签名头
	signatureHead := fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%d\",serial_no=\"%s\"", c.mchId, nonceStr, signature, timestamp, c.serialNo)
	request, err := http.NewRequest(method, c.domain+url, ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return
	}
	if headerLen := len(headers); headerLen > 0 && headerLen%2 == 0 {
		for i := 0; i < headerLen-1; i++ {
			hk := headers[i]
			hv := headers[i+1]
			request.Header.Set(hk, hv)
		}
	}
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", "WECHATPAY2-SHA256-RSA2048", signatureHead))
	request.Header.Set("Content-Type", ContentTypeJSON)
	request.Header.Set("Accept", ContentTypeJSON)
	request.Header.Set("User-Agent", c.agent())
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
	request.Header.Set("User-Agent", c.agent())
	request.Header.Set("Accept-Language", "zh-CN")
	return c.httpClient.Do(request)
}

// ParseWechatResponse 验证向微信服务器发送请求后从微信得到的应答签名
// 微信(验证)签名验证详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_1.shtml
func (c *Config) ParseWechatResponse(response *http.Response, dst interface{}) (requestId string, err error) {
	if response == nil {
		err = errors.ErrNoHttpResponse
		return
	}
	var body []byte
	var header = response.Header
	var code = response.StatusCode // 根据http code 判断请求是否成功

	// 获取唯一请求ID
	requestId = header.Get("Request-ID")
	// 读取response body
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	_ = response.Body.Close()

	// 请求失败，根据http.StatusCode返回对应的error
	// 同时将读取出来的body(如果有的话)反序列化到对应的结果中
	if code != http.StatusOK && code != http.StatusNoContent {
		if err = unmarshalJSON(body, dst); err != nil {
			return
		}
		if data, ok := dst.(model.WError); ok && data != nil {
			if err = data.Error(); err != nil {
				return
			}
		}
		err = errors.New(code)
		return
	}

	// API调用成功后的处理流程
	// 1. 验证证书序列号是否正确
	serialNo := header.Get("Wechatpay-Serial")
	publicKey := c.GetRSAPublicKey(serialNo)
	if publicKey == nil {
		err = fmt.Errorf("解析微信应答失败: Wechatpay-Serial[%s]不存在", serialNo)
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
	_ = c.merchantCipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8)
	err = rsas.VerifySHA256WithRSA(c.merchantCipher, wechatSign, plainTxt)
	if err != nil {
		return
	}
	err = unmarshalJSON(body, dst)
	return
}

// ParseWechatNotify 验证微信服务器的通知，预支付、退款等请求后，微信回调的Request同样需要签名验证
// 微信（验证）签名方法详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_1.shtml
func (c *Config) ParseWechatNotify(request *http.Request, dst interface{}) (notifyId string, err error) {
	if request == nil {
		err = errors.ErrNoHttpRequest
		return
	}
	if c.apiKey == "" {
		err = errors.ErrNoApiV3Key
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
	publicKey := c.GetRSAPublicKey(serialNo)
	if publicKey == nil {
		err = fmt.Errorf("解析微信通知失败: Wechatpay-Serial[%s]不存在", serialNo)
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
	_ = c.merchantCipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8)
	err = rsas.VerifySHA256WithRSA(c.merchantCipher, wechatSign, plainTxt)
	if err != nil {
		return
	}

	// 签名通过的话反序列化body到结构体中
	notifyResponse := new(model.WechatNotifyResponse)
	if err = unmarshalJSON(body, &notifyResponse); err != nil {
		return
	}

	notifyId = notifyResponse.Id
	// 判断资源类型
	if notifyResponse.ResourceType != "encrypt-resource" {
		err = fmt.Errorf("解析微信通知失败, 错误的资源类型: %s", notifyResponse.ResourceType)
		return
	}
	if notifyResponse.Resource == nil {
		err = errors.ErrInvalidResource
		return
	}

	// 解密
	cipherText := notifyResponse.Resource.CipherText
	associateData := notifyResponse.Resource.AssociatedData
	nonce := notifyResponse.Resource.Nonce
	plainData, err := aess.DecryptAEADAES256GCM(c.merchantCipher, c.apiKey, cipherText, associateData, nonce)
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

// UploadMedia 上传多媒体文件到微信服务器
func (c *Config) UploadMedia(url string, contentType string, fileName string, fileData []byte) (response *http.Response, err error) {
	if c.mchId == "" {
		err = errors.ErrNoMchId
		return
	}
	if c.serialNo == "" {
		err = errors.ErrNoSerialNo
		return
	}
	// 获取文件内容的sha256摘要值
	sha256Value, err := c.hasher.HashToString(fileData, crypto.SHA256)
	if err != nil {
		return
	}
	meta := pkg.NewParam()
	meta.Add("filename", fileName)
	meta.Add("sha256", sha256Value)

	// JSON序列化meta数据
	metaData, err := marshalJSON(meta)
	if err != nil {
		return
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	buildUploadBody(writer, contentType, fileName, fileData, metaData)
	if err = writer.Close(); err != nil {
		return
	}

	// 构造签名
	method := http.MethodPost      // 方法类型
	timestamp := time.Now().Unix() // 时间戳
	nonceStr := pkg.String(32)     // 随机字符串
	source := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, url, timestamp, nonceStr, string(metaData))
	signature, err := rsas.SignSHA256WithRSA(c.merchantCipher, source)
	if err != nil {
		return
	}
	// 签名头
	signatureHead := fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%d\",serial_no=\"%s\"", c.mchId, nonceStr, signature, timestamp, c.serialNo)
	// 构造请求头，这里的body为文件二进制数据
	request, err := http.NewRequest(method, c.domain+url, body)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", fmt.Sprintf("%s %s", "WECHATPAY2-SHA256-RSA2048", signatureHead))
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Accept", "*/*")
	request.Header.Set("User-Agent", c.agent())
	request.Header.Set("Accept-Language", "zh-CN")
	return c.httpClient.Do(request)
}

// VerifyHashValue 校验hash值
func (c *Config) VerifyHashValue(hashType crypto.Hash, data interface{}, hashValue string) (err error) {
	v, err := c.hasher.HashToString(data, hashType)
	if err != nil {
		return err
	}
	if strings.ToUpper(hashValue) != strings.ToUpper(v) {
		err = errors.ErrCheckHashValueFail
		return
	}
	return
}

// ImageExt 获取图片后缀名(图片格式)
func ImageExt(name string) (contentType string, err error) {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case "jpg":
		contentType = "image/jpg"
	case "bmp":
		contentType = "image/bmp"
	case "png":
		contentType = "image/png"
	default:
		err = fmt.Errorf("图片文件名必须以jpg、png、bmp为后缀: %s", ext)
	}
	return
}

// VideoExt 获取视频的后缀名(视频格式)
func VideoExt(name string) (contentType string, err error) {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case "avi":
		contentType = "video/avi"
	case "wmv":
		contentType = "video/wmv"
	case "mpeg":
		contentType = "video/mpeg"
	case "mp4":
		contentType = "video/mp4"
	case "mov":
		contentType = "video/mov"
	case "mkv":
		contentType = "video/mkv"
	case "flv":
		contentType = "video/flv"
	case "f4v":
		contentType = "video/f4v"
	case "m4v":
		contentType = "video/m4v"
	case "rmvb":
		contentType = "video/rmvb"
	default:
		err = fmt.Errorf("视频文件名只能以avi、wmv、mpeg、mp4、mov、mkv、flv、f4v、m4v、rmvb为后缀: %s", ext)
	}
	return
}

func buildUploadBody(writer *multipart.Writer, contentType, fileName string, fileContent, metaData []byte) {
	field := make(textproto.MIMEHeader)
	field.Set("Content-Disposition", "form-data; name=\"meta\";")
	field.Set("Content-Type", ContentTypeJSON)
	part, err := writer.CreatePart(field)
	if err != nil {
		return
	}
	if _, err = part.Write(metaData); err != nil {
		return
	}

	field = make(textproto.MIMEHeader)
	field.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"file\"; filename=\"%s\"", fileName))
	field.Set("Content-Type", contentType)
	part, err = writer.CreatePart(field)
	if err != nil {
		return
	}
	if _, err = part.Write(fileContent); err != nil {
		return
	}
}

func unmarshalJSON(data []byte, dst interface{}) (err error) {
	if len(data) == 0 {
		return
	}
	err = json.Unmarshal(data, &dst)
	return
}

func marshalJSON(data interface{}) (bytes []byte, err error) {
	if data == nil {
		return
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.IsZero() {
		return
	}

	dataType := reflect.TypeOf(data).Kind()
	switch dataType {
	case reflect.String:
		bytes = []byte(dataValue.String())
	case reflect.Slice:
		if dataValue.Elem().Kind() != reflect.Uint8 {
			err = errors.ErrMarshalFailInvalidDataType
			break
		}
		bytes = dataValue.Bytes()
	case reflect.Struct, reflect.Ptr, reflect.Map:
		bytes, err = json.Marshal(data)
	default:
		err = fmt.Errorf("序列化失败, 数据类型不支持: %s", dataType.String())
	}
	return
}

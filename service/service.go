package service

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/rands"
	"github.com/pyihe/secret"
	"github.com/pyihe/wechat-sdk/pkg/rsas"
	"github.com/pyihe/wechat-sdk/vars"
)

// RequestWithSign 对发送给微信服务器的body进行SHA-256 with RSA签名, 返回*http.Request
// 参数说明:
// method: api方法类型, 如: GET、POST等
// url: api接口除去域名的绝对URL, 如: /v3/pay/transactions/jsapi
// body: 请求主体，比如支付时为支付参数，调用方需要不序列化
// 返回参数说明:
// signResult: 返回用于签名的各个参数，包括签名结果
// 签名介绍详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_0.shtml
func RequestWithSign(config *Config, method, url string, body interface{}) (response *http.Response, err error) {
	if config.MchId == "" {
		err = vars.ErrNoMchId
		return
	}
	if config.Cipher == nil {
		err = vars.ErrInvalidCipher
		return
	}
	if config.SerialNo == "" {
		err = vars.ErrNoSerialNo
		return
	}
	// 构造签名主体
	var data []byte
	if body != nil {
		switch content := body.(type) {
		case []byte:
			data = content
		default:
			data, err = Marshal(body)
			if err != nil {
				return
			}
		}
	}

	method = strings.ToUpper(method) // 方法类型，转为大写
	timestamp := time.Now().Unix()   // 时间戳
	nonceStr := rands.String(32)     // 随机字符串

	source := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, url, timestamp, nonceStr, string(data))
	signature, err := rsas.SignSHA256WithRSA(config.Cipher, source)
	if err != nil {
		return
	}
	// 签名头
	signatureHead := fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%d\",serial_no=\"%s\"", config.MchId, nonceStr, signature, timestamp, config.SerialNo)
	request, err := http.NewRequest(method, config.domain+url, ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return
	}
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", "WECHATPAY2-SHA256-RSA2048", signatureHead))
	request.Header.Set("Content-Type", vars.ContentTypeJSON)
	request.Header.Set("Accept", vars.ContentTypeJSON)
	request.Header.Set("User-Agent", "Pyihe-Wechat-SDK")
	request.Header.Set("Accept-Language", "zh-CN")
	return config.HttpClient.Do(request)
}

// Request 发起普通的HTTP请求
func Request(config *Config, method, url, contentType string, body interface{}) (response *http.Response, err error) {
	data, err := Marshal(body)
	if err != nil {
		return
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Accept", contentType)
	request.Header.Set("User-Agent", "Pyihe-Wechat-SDK")
	request.Header.Set("Accept-Language", "zh-CN")
	return config.HttpClient.Do(request)
}

// VerifyResponse 验证向微信服务器发送请求后从微信得到的应答签名
// 微信(验证)签名验证详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_1.shtml
func VerifyResponse(config *Config, response *http.Response) (requestId string, body []byte, err error) {
	if response == nil {
		err = errors.New("response为空!")
		return
	}
	header := response.Header
	// 获取唯一请求ID
	requestId = header.Get("Request-ID")

	// 根据http code 判断请求是否成功
	var code = response.StatusCode
	// 请求已经被接收，但尚未处理，需要重复请求一遍
	if code == http.StatusAccepted {
		err = vars.ErrRequestAgain
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
	publicKey, ok := config.Certificates[serialNo].(*rsa.PublicKey)
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
	_ = config.Cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8)
	err = rsas.VerifySHA256WithRSA(config.Cipher, wechatSign, plainTxt)
	return
}

// VerifyRequest 验证微信服务器的通知，预支付、退款等请求后，微信回调的Request同样需要签名验证
// 微信（验证）签名方法详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_1.shtml
func VerifyRequest(config *Config, request *http.Request) (body []byte, err error) {
	header := request.Header

	body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}
	_ = request.Body.Close()

	// 1. 验证证书序列号是否正确
	serialNo := header.Get("Wechatpay-Serial")
	publicKey, ok := config.Certificates[serialNo].(*rsa.PublicKey)
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
	_ = config.Cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8)
	err = rsas.VerifySHA256WithRSA(config.Cipher, wechatSign, plainTxt)
	return
}

// Unmarshal 反序列化
func Unmarshal(data []byte, dst interface{}) (err error) {
	if len(data) > 0 {
		err = json.Unmarshal(data, &dst)
	}
	return
}

// Marshal 序列化
func Marshal(data interface{}) (bytes []byte, err error) {
	bytes, err = json.Marshal(data)
	return
}

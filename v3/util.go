package v3

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pyihe/wechat-sdk/v3/vars"

	"github.com/pyihe/go-pkg/files"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/rands"
	"github.com/pyihe/secret"
)

func (we *WeChatClient) getCertificate(serialNo string) (publicKey *rsa.PublicKey) {
	if we.publicKeys != nil {
		publicKey = we.publicKeys[serialNo]
	}
	return
}

// verifyWechatSign 验证微信服务器返回的数据签名
// 签名验证详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_1.shtml
func (we *WeChatClient) verifyWechatSign(header http.Header, reader io.ReadCloser, httpCodes ...int) (requestId string, body []byte, err error) {
	requestId = header.Get("Request-ID")

	// 1. 验证证书序列号是否正确
	serialNo := header.Get("Wechatpay-Serial")
	publicKey := we.getCertificate(serialNo)
	if publicKey == nil {
		err = errors.New("inconsistent Wechatpay-Serial.")
		return
	}

	// 2. 获取微信的签名结果
	wechatSign := header.Get("Wechatpay-Signature") // 微信签名
	// 3. 获取签名参数
	timestamp := header.Get("Wechatpay-Timestamp") // 时间戳
	nonceStr := header.Get("Wechatpay-Nonce")      // 随机字符串

	// 读取签名主体
	// 当状态码为http.StatusNoContent时无body返回
	var code int
	if len(httpCodes) > 0 {
		code = httpCodes[0]
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
			if reader != nil {
				defer reader.Close()
				body, err = ioutil.ReadAll(reader)
				if err != nil {
					return
				}
			}
		}
	}
	// 4. 构造原始的签名数据
	plainTxt := fmt.Sprintf("%v\n%v\n%v\n", timestamp, nonceStr, string(body))

	// 验证签名
	_ = we.cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel1)
	ok, err := we.cipher.RSAVerify(wechatSign, plainTxt, secret.RSASignTypePKCS1v15, crypto.SHA256)
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("签名验证失败!")
		return
	}
	return
}

// unMarshal 反序列化
func unMarshal(body []byte, dst interface{}) error {
	return json.Unmarshal(body, &dst)
}

// marshal 序列化
func marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (we *WeChatClient) decryptAEADAES256GCM(cipherText, associateData, nonce string) (plainText []byte, err error) {
	var decryptRequest = &secret.SymRequest{
		CipherData:  cipherText,
		Key:         []byte(we.apiKey),
		Type:        secret.SymTypeAES,
		ModeType:    secret.BlockModeGCM,
		PaddingType: secret.PaddingTypeNoPadding,
		AddData:     []byte(associateData),
		Nonce:       []byte(nonce),
	}
	return we.cipher.SymDecrypt(decryptRequest)
}

// SHA-256 with RSA 签名
// 请求参数说名:
// method: api方法类型, 如: GET、POST等
// abUrl: api接口除去域名的绝对URL, 如: /v3/pay/transactions/jsapi
// body: 请求主体，比如支付时为支付参数，调用方需要不序列化
// 返回参数说明:
// param: 返回用于签名的各个参数，包括签名结果
// 签名介绍详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_0.shtml
func (we *WeChatClient) signSHA256WithRSA(method, abUrl string, body interface{}) (signParam vars.Kvs, err error) {
	var data []byte
	// JSON序列化body
	if body != nil {
		switch content := body.(type) {
		case []byte:
			data = content
		default:
			data, err = json.Marshal(body)
			if err != nil {
				return nil, err
			}
		}
	}

	method = strings.ToUpper(method)
	timestamp := time.Now().Unix() // 时间戳
	noneStr := rands.String(32)    // 随机字符串

	// 签名主体, 格式为:
	//	HTTP请求方法\n
	//	API绝对URL\n
	//	请求时间戳\n
	//	请求随机字符串\n
	//	请求报文主体\n
	source := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, abUrl, timestamp, noneStr, string(data))
	signature, err := we.cipher.RSASignToString(source, secret.RSASignTypePKCS1v15, crypto.SHA256)
	if err != nil {
		return nil, err
	}

	signParam = vars.NewKvs()
	signParam.Add("method", method)
	signParam.Add("abUrl", abUrl)
	signParam.Add("timestamp", timestamp)
	signParam.Add("nonce_str", noneStr)
	signParam.Add("body", string(data))
	signParam.Add("signature", signature)
	return
}

// requestWithSign 携带签名的Http请求
// API请求发起详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_0.shtml
func (we *WeChatClient) requestWithSign(method, apiUrl string, signParam vars.Kvs, body io.Reader) (response *http.Response, err error) {
	fmt.Println(apiUrl)
	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, err
	}
	nonceStr, _ := signParam.GetString("nonce_str")
	timestamp, _ := signParam.GetInt64("timestamp")
	signature, _ := signParam.GetString("signature")
	signatureInfo := fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%d\",serial_no=\"%s\"", we.mchId, nonceStr, signature, timestamp, we.serialNo)

	request.Header.Set("Authorization", fmt.Sprintf("%s %s", "WECHATPAY2-SHA256-RSA2048", signatureInfo))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "Pyihe-Wechat-SDK")
	request.Header.Set("Accept-Language", "zh-CN")
	return we.httpClient.Do(request)
}

// request 发起普通http请求
func (we *WeChatClient) request(method, url, contentType string, body io.Reader) (response *http.Response, err error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Accept", contentType)
	request.Header.Set("User-Agent", "Pyihe-Wechat-SDK")
	request.Header.Set("Accept-Language", "zh-CN")
	return we.httpClient.Do(request)
}

// 将data写入指定路径指定的文件中
func writeToFile(path, fileName string, data []byte) error {
	//判断是否存在目标目录，如果不存在则创建
	if !strings.HasSuffix(path, "/") {
		if path == "" {
			path = "./"
		} else {
			path += "/"
		}
	}
	if err := files.MakeNewPath(path); err != nil {
		return err
	}

	f, err := os.Create(path + fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// 反序列化公钥证书
func unmarshalPublicKey(key []byte) (serialNo string, publicKey *rsa.PublicKey, err error) {
	block, _ := pem.Decode(key)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return
	}
	serialNo = strings.ToUpper(cert.SerialNumber.Text(16))
	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		err = errors.New("invalid certificate.")
		return
	}
	return
}

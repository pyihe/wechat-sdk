package wechat_sdk

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/pyihe/wechat-sdk/pkg/e"
	"github.com/pyihe/wechat-sdk/pkg/util"
)

//用于装载外部传入的请求参数
type Param map[string]interface{}

func NewParam() Param {
	return make(Param)
}

func (p Param) Get(key string) interface{} {
	if p == nil {
		return nil
	}
	return p[key]
}

func (p Param) Add(key string, value interface{}) {
	p[key] = value
}

func (p Param) Del(key string) {
	delete(p, key)
}

func (p Param) MarshalXML() (reader io.Reader, err error) {
	buffer := bytes.NewBuffer(make([]byte, 0))

	if _, err = io.WriteString(buffer, "<xml>"); err != nil {
		return
	}

	for k, v := range p {
		switch {
		case k == "detail":
			if _, err = io.WriteString(buffer, "<detail><![CDATA["); err != nil {
				return
			}
			if _, err = io.WriteString(buffer, fmt.Sprintf("%v", v)); err != nil {
				return
			}
			if _, err = io.WriteString(buffer, "]]></detail>"); err != nil {
				return
			}
		default:
			if _, err = io.WriteString(buffer, "<"+k+">"); err != nil {
				return
			}
			if err = xml.EscapeText(buffer, []byte(fmt.Sprintf("%v", v))); err != nil {
				return
			}
			if _, err = io.WriteString(buffer, "</"+k+">"); err != nil {
				return
			}
		}
	}

	if _, err = io.WriteString(buffer, "</xml>"); err != nil {
		return
	}
	return buffer, nil
}

func (p Param) SortKey() (keys []string) {
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}

func (p Param) Sign(signType string) string {
	var result string
	keys := p.SortKey()
	var signStr string
	for i, k := range keys {
		if k == "sign" {
			continue
		}
		str := ""
		if i == 0 {
			str = fmt.Sprintf("%v=%v", k, p[k])
		} else {
			str = fmt.Sprintf("&%v=%v", k, p[k])
		}
		signStr += str
	}
	signStr += fmt.Sprintf("&key=%v", defaultPayer.apiKey)
	switch signType {
	case e.SignType256:
		result = strings.ToUpper(util.SignHMACSHA256(signStr, defaultPayer.apiKey))
	case e.SignTypeMD5:
		result = strings.ToUpper(util.SignMd5(signStr))
	default:
	}
	return result
}

/*用于装载微信返回的数据*/
type resultMap map[string]string

func newResultMap() resultMap {
	return make(resultMap)
}

func (r resultMap) Add(key, value string) {
	r[key] = value
}

func (r resultMap) GetString(key string) (string, error) {
	if _, ok := r[key]; ok {
		return r[key], nil
	}
	return "", errors.New("not exist key: " + key)
}

func (r resultMap) GetInt64(key string, base int) (no int64, err error) {
	if _, ok := r[key]; ok {
		no, err = strconv.ParseInt(r[key], base, 64)
		return
	}
	return 0, errors.New("not exist key: " + key)
}

func (r resultMap) Data() map[string]string {
	return r
}

func (r resultMap) Sign(signType string) string {
	p := NewParam()
	for k, v := range r {
		p[k] = v
	}
	return p.Sign(signType)
}

//XML to Param(遍历生成)
func ParseXMLReader(reader io.Reader) resultMap {
	result := newResultMap()
	decoder := xml.NewDecoder(reader)

	var (
		t     xml.Token
		err   error
		key   string
		value string
	)

	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement: // 处理元素开始（标签）
			key = token.Name.Local
		case xml.CharData: // 处理字符数据（这里就是元素的文本）
			//这里需要注意，元素值可能包含换行符和空格，需要先去掉空格
			value = string(token)
			if newToken := strings.Replace(string(token.Copy()), " ", "", -1); newToken == "" || newToken == "\n" {
				value = ""
			}
		default:
		}
		//获取到标签和元素后，像map中添加
		if key != "xml" && key != "" && value != "" && value != "\n" {
			result.Add(key, value)
		}
	}
	return result
}

//用于像微信发送POST请求
type postRequest struct {
	Body io.Reader

	Url         string
	ContentType string
}

//向微信服务器发送POST请求
func postToWx(req *postRequest) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("have no postRequest")
	}

	response, err := http.Post(req.Url, req.ContentType, req.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("http StatusCode: " + strconv.Itoa(response.StatusCode))
	}

	return response, nil
}

//带证书的Post请求
func postToWxWithCert(req *postRequest, p12Cert *tls.Certificate) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("have no postRequest")
	}

	if p12Cert == nil {
		return nil, errors.New("need p12Cert")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{*p12Cert},
		},
		DisableCompression: true,
	}
	httpClient := http.Client{
		Transport: transport,
	}
	//发送请求
	response, err := httpClient.Post(req.Url, req.ContentType, req.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("http StatusCode: " + strconv.Itoa(response.StatusCode))
	}

	return response, err
}

package wechat_sdk

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/wechat-sdk/pkg/e"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type WeClient interface {
	SetParams(map[string]interface{})
	AddParam(key string, value interface{})
	AddParams(map[string]interface{})
	DelParam(key string)
	//GetParamByKey(string) interface{}
	//GetParams() map[string]interface{}

	GetWxOpenId(code string) (Token, error)
	GetWxUserInfo(lang string) (User, error)
	ShareH5(string) (*H5Response, error)
}

type myWeClient struct {
	sync.Mutex                       //params读写锁
	params    map[string]interface{} //用于存放请求接口时需要传入的参数
	appId     string                 //appId
	appSecret string                 //app密钥
	mchId     string                 //商户号
}

var (
	c *myWeClient
)

func NewClientWithParam(appId, appSecret, mchId string) WeClient {
	c = &myWeClient{
		params:    make(map[string]interface{}),
		appId:     appId,
		appSecret: appSecret,
		mchId:     mchId,
	}
	return c
}

func (client *myWeClient) SetParams(param map[string]interface{}) {
	client.Lock()
	if param != nil {
		client.params = param
	}
	client.Unlock()
}

func (client *myWeClient) AddParam(key string, value interface{}) {
	client.Lock()
	if client.params == nil {
		client.params = make(map[string]interface{})
	}
	client.params[key] = value
	client.Unlock()
}

func (client *myWeClient) AddParams(params map[string]interface{}) {
	client.Lock()
	if client.params == nil {
		client.params = make(map[string]interface{})
	}
	if params != nil {
		for k, v := range params {
			client.params[k] = v
		}
	}
	client.Unlock()
}

func (client *myWeClient) DelParam(key string) {
	client.Lock()
	if client.params == nil {
		return
	}
	delete(client.params, key)
	client.Unlock()
}

//func (client *myWeClient) GetParamByKey(key string) interface{} {
//	if client.params != nil {
//		if _, ok := client.params[key]; ok {
//			return client.params[key]
//		}
//	}
//	return nil
//}
//
//func (client *myWeClient) GetParams() map[string]interface{} {
//	return client.params
//}

//获取微信openId
//code: 用户同意授权后，前端获取的code
//后端可以校验微信后台设置的state
func (client *myWeClient) GetWxOpenId(code string) (token Token, err error) {
	//拉取网页授权token的api
	apiUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%v&secret=%v&code=%v&grant_type=authorization_code", client.appId, client.appSecret, code)

	response, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data *tokenInfo
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.errMsg != "" || data.errCode != 0 {
		return nil, errors.New(data.errMsg)
	}
	//获取到openId后放入client中
	params := map[string]interface{}{
		"openid":       data.openId,
		"access_token": data.authAccessToken,
	}

	client.AddParams(params)

	return data, nil
}

//获取微信用户信息
//lang: 返回国家地区版本, zh_CN 简体，zh_TW 繁体，en 英语
func (client *myWeClient) GetWxUserInfo(lang string) (User, error) {
	if client.params == nil {
		return nil, e.ErrorNilParam
	}

	openId, ok := client.params["openid"]
	if !ok {
		return nil, e.ErrorNoOpenId
	}
	accessToken, ok := client.params["access_token"]
	if !ok {
		return nil, e.ErrorNoToken
	}
	//获取用户基本信息的API url
	apiUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%v&openid=%v&lang=%v", accessToken, openId, lang)

	response, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var info *userInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	//校验是否正确返回了数据
	if info.errMsg != "" || info.errCode != 0 {
		return nil, errors.New(info.errMsg)
	}
	return info, nil
}

//微信H5页面分享
//url:待分享的页面url
func (client *myWeClient) ShareH5(url string) (*H5Response, error) {
	ticket, err := getTicketFromWx()
	if err != nil {
		return nil, err
	}

	//时间戳
	now := time.Now().Unix()

	//随机字符串
	u4, _ := uuid.NewV4()
	s := strings.Replace(fmt.Sprintf("%s", u4), "-", "", -1)
	nonceStr := s[:16]

	//签名
	t := sha1.New()
	toSignStr := fmt.Sprintf("jsapi_ticket=%v&noncestr=%v&timestamp=%vurl=%v", ticket.Ticket, nonceStr, now, url)
	io.WriteString(t, toSignStr)
	sign := fmt.Sprintf("%x", t.Sum(nil))

	resp := &H5Response{
		AppID:     client.appId,
		TimeStamp: now,
		NonceStr:  nonceStr,
		Signature: sign,
	}
	return resp, nil
}

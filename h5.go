package wechat_sdk

import (
	"encoding/json"
	"errors"
	"github.com/wechat-sdk/pkg/e"
	"io/ioutil"
	"net/http"
)

//获取ticket时的回复
type jsApiTicketInfo struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

//分享时回复给客户端的消息
type H5Response struct {
	Code      int    `json:"code"`
	Desc      string `json:"desc"`
	AppID     string `json:"app_id"`
	TimeStamp int64  `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}

//从微信拉取基础支持的access_token
func getTokenFromWX() (Token, error) {
	if c.appId == "" || c.appSecret == "" {
		return nil, e.ErrorInitClinet
	}
	tokenHost := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + c.appId + "&secret=" + c.appSecret
	response, err := http.Get(tokenHost)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var info *tokenInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	if info.errCode != 0 || info.errMsg != "" {
		return nil, errors.New(info.errMsg)
	}
	return info, nil
}

//从微信拉取h5分享用的ticket
func getTicketFromWx() (*jsApiTicketInfo, error) {
	token, err := getTokenFromWX()
	if err != nil {
		return nil, err
	}
	apiUrl := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + token.AuthAccessToken() + "&type=jsapi"
	response, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var ticketDocker *jsApiTicketInfo
	err = json.Unmarshal(body, &ticketDocker)
	if err != nil {
		return nil, err
	}
	if ticketDocker.ErrCode != 0 {
		return nil, errors.New(ticketDocker.ErrMsg)
	}
	return ticketDocker, nil
}

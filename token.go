package wechat_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Token interface {
	GetOpenId() string
	GetAuthAccessToken() string
	GetExpiresIn() int64
	GetScope() string
	GetRefreshToken() string
	RefreshAccessToken() error
	Err() string
}

type tokenInfo struct {
	OpenId          string `json:"openid"`
	AuthAccessToken string `json:"access_token"`
	ExpiresIn       int64  `json:"expires_in"`
	RefreshToken    string `json:"refresh_token"`
	Scope           string `json:"scope"`
	ErrCode         int    `json:"errcode"`
	ErrMsg          string `json:"errmsg"`
}

func (t *tokenInfo) Err() string {
	return fmt.Sprintf("code: %v msg: %v", t.ErrCode, t.ErrMsg)
}

func (t *tokenInfo) GetOpenId() string {
	return t.OpenId
}

func (t *tokenInfo) GetAuthAccessToken() string {
	return t.AuthAccessToken
}

func (t *tokenInfo) GetExpiresIn() int64 {
	return t.ExpiresIn
}

func (t *tokenInfo) GetRefreshToken() string {
	return t.RefreshToken
}

func (t *tokenInfo) GetScope() string {
	return t.Scope
}

func (t *tokenInfo) RefreshAccessToken() error {
	if c == nil {
		return errors.New("call NewClientWithParam first")
	}
	apiUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%v&grant_type=refresh_token&refresh_token=%v", c.appId, t.GetRefreshToken())

	response, err := http.Get(apiUrl)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return err
	}
	//每刷新一次token，同时更新client中的openId
	c.AddParam("openid", t.OpenId)

	return nil
}

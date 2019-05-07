package wechat_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Token interface {
	OpenId() string
	AuthAccessToken() string
	ExpiresIn() int64
	Scope() string
	GetRefreshToken() string
	RefreshAccessToken() error
}

type tokenInfo struct {
	openId          string `json:"openid"`
	authAccessToken string `json:"access_token"`
	expiresIn       int64  `json:"expires_in"`
	refreshToken    string `json:"refresh_token"`
	scope           string `json:"scope"`
	errCode         int    `json:"errcode"`
	errMsg          string `json:"errmsg"`
}

func (t *tokenInfo) OpenId() string {
	return t.openId
}

func (t *tokenInfo) AuthAccessToken() string {
	return t.authAccessToken
}

func (t *tokenInfo) ExpiresIn() int64 {
	return t.expiresIn
}

func (t *tokenInfo) GetRefreshToken() string {
	return t.refreshToken
}

func (t *tokenInfo) Scope() string {
	return t.scope
}

func (t *tokenInfo) RefreshAccessToken() error {
	if c == nil {
		return errors.New("call NewClientWithParam first")
	}
	apiUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%v&grant_type=refresh_token&refresh_token=%v", c.appId, t.refreshToken)

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
	c.AddParam("openid", t.openId)

	return nil
}

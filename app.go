package wechat_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//公众号获取普通接口授权Access_Token
func (m *myPayer) GetAppBaseAccessToken() (Param, error) {
	if err := m.checkForAccess(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", m.appId, m.secret)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result = NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//公众号获取网页授权的Access_Token
func (m *myPayer) GetAppOauthAccessToken(code string) (Param, error) {
	if code == "" {
		return nil, errors.New("code is empty")
	}
	if err := m.checkForAccess(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%v&secret=%v&code=%v&grant_type=authorization_code", m.appId, m.secret, code)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result = NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//公众号刷新网页授权Access_Token
func (m *myPayer) RefreshOauthToken(refreshToken string) (Param, error) {
	if refreshToken == "" {
		return nil, errors.New("refresh_token is empty")
	}
	if err := m.checkForAccess(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%v&grant_type=refresh_token&refresh_token=%v", m.appId, refreshToken)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result = NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//公众号拉取用户信息, lang值: zh_CN(简体中文) zh_TW(繁体中文) en(英语)
func (m *myPayer) GetAppUserInfo(oauthToken, openId, lang string) (Param, error) {
	if oauthToken == "" {
		return nil, errors.New("empty access_token")
	}
	if openId == "" {
		return nil, errors.New("empty openId")
	}
	if err := m.checkForAccess(); err != nil {
		return nil, err
	}
	if lang == "" {
		lang = "zh_CN"
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%v&openid=%v&lang=%v", oauthToken, openId, lang)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var result = NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//公众号校验授权Access_Token是否有效
func (m *myPayer) CheckOauthToken(oauthToken, openId string) (bool, error) {
	if oauthToken == "" {
		return false, errors.New("empty access_token")
	}
	if openId == "" {
		return false, errors.New("empty openId")
	}
	if err := m.checkForAccess(); err != nil {
		return false, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%v&openid=%v", oauthToken, openId)

	response, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	var result = NewParam()
	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return false, err
	}

	if msg := result.Get("errmsg"); msg == nil {
		return false, errors.New("no response")
	} else if msg.(string) != "ok" {
		return false, errors.New(msg.(string))
	}
	return true, nil
}

package dev

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

/*
	微信小程序根据code获取openId和session_key
*/

type sessionInfo struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

func (s *sessionInfo) Param(key string) (interface{}, error) {
	var err error
	switch key {
	case "openid":
		return s.Openid, err
	case "session_key":
		return s.SessionKey, err
	case "unionid":
		return s.Unionid, err
	case "errcode":
		return s.Errcode, err
	case "errmsg":
		return s.Errmsg, err
	default:
		err = errors.New(fmt.Sprintf("invalid key: %s", key))
		return nil, err
	}
}

func (s sessionInfo) ListParam() Params {
	p := make(Params)

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).IsZero() {
			tagName := strings.Split(string(t.Field(i).Tag), "\"")[1]
			p[tagName] = v.Field(i).Interface()
		}
	}
	return p
}

//通过code获取session_key
func (m *myPayer) GetSessionKeyAndOpenId(code string) (ResultParam, error) {
	if err := m.checkForAccess(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code", m.appId, m.secret, code)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result *sessionInfo
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result.Errcode != 0 {
		return nil, errors.New(result.Errmsg)
	}

	return result, nil
}

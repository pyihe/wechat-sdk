package dev

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type tokenResult struct {
	AccessToken string `json:"access_token"`
	ExpireIn    int    `json:"expire_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func (token *tokenResult) Param(key string) (interface{}, error) {
	var err error
	switch key {
	case "access_token":
		return token.AccessToken, err
	case "expire_in":
		return token.ExpireIn, err
	case "errcode":
		return token.ErrCode, err
	case "errmsg":
		return token.ErrMsg, err
	default:
		err = errors.New(fmt.Sprintf("invalid key: %s", key))
		return nil, err
	}
}

func (token tokenResult) ListParam() Params {
	p := make(Params)

	t := reflect.TypeOf(token)
	v := reflect.ValueOf(t)

	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).IsZero() {
			tagName := strings.Split(string(t.Field(i).Tag), "\"")[1]
			p[tagName] = v.Field(i).Interface()
		}
	}
	return p
}

func (m *myPayer) GetAccessTokenForMini() (ResultParam, error) {
	if err := m.checkForAccess(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", m.appId, m.secret)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result *tokenResult
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

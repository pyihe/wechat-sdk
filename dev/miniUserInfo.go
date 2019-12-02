package dev

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

type user struct {
	Gender    int        `json:"gender"`
	OpenId    string     `json:"openId"`
	NickName  string     `json:"nickName"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	Country   string     `json:"country"`
	AvatarUrl string     `json:"avatarUrl"`
	UnionId   string     `json:"unionId"`
	WaterMark *waterMark `json:"watermark"`
}

type waterMark struct {
	AppId     string `json:"appid"`
	TimeStamp int64  `json:"timestamp"`
}

func (u *user) Param(key string) (interface{}, error) {
	var err error
	switch key {
	case "gender":
		return u.Gender, err
	case "openId":
		return u.OpenId, err
	case "nickName":
		return u.NickName, err
	case "city":
		return u.City, err
	case "avatarUrl":
		return u.AvatarUrl, err
	case "province":
		return u.Province, err
	case "unionId":
		return u.UnionId, err
	case "watermark":
		return u.WaterMark, err
	case "country":
		return u.Country, err
	default:
		err = errors.New(fmt.Sprintf("invalid key: %s", key))
		return nil, err
	}
}

func (u user) ListParam() Params {
	p := make(Params)

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).IsZero() {
			tagName := strings.Split(string(t.Field(i).Tag), "\"")[1]
			p[tagName] = v.Field(i).Interface()
		}
	}
	return p
}

//小程序获取用户信息
func (m *myPayer) GetUserInfoForMini(code string, dataStr string, ivStr string) (ResultParam, error) {
	session, err := m.GetSessionKey(code)
	if err != nil {
		return nil, err
	}

	key, err := base64.StdEncoding.DecodeString(session.(*sessionInfo).SessionKey)
	if err != nil {
		return nil, err
	}

	encryptedData, err := base64.StdEncoding.DecodeString(dataStr)
	if err != nil {
		return nil, err
	}

	iv, err := base64.StdEncoding.DecodeString(ivStr)
	if err != nil {
		return nil, err
	}

	realData, err := util.AES128CBCDecrypt(encryptedData, key, iv)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nreaData = %v\n", string(realData))
	var info *user
	err = json.Unmarshal(realData, &info)
	if err != nil {
		return nil, err
	}
	if info.WaterMark == nil {
		return nil, e.ErrNoWatermark
	}
	if info.WaterMark.AppId != m.appId {
		return nil, errors.New("differ appid")
	}
	return info, nil
}



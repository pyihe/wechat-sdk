package dev

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	微信小程序获取用户电话号码
*/

type phoneInfo struct {
	PhoneNumber     string     `json:"phoneNumber"`
	PurePhoneNumber string     `json:"purePhoneNumber"`
	CountryCode     string     `json:"countryCode"`
	WaterMark       *waterMark `json:"watermark"`
}

func (phone *phoneInfo) Param(key string) (interface{}, error) {
	var err error
	switch key {
	case "phoneNumber":
		return phone.PhoneNumber, err
	case "purePhoneNumber":
		return phone.PurePhoneNumber, err
	case "countryCode":
		return phone.CountryCode, err
	case "watermark":
		return phone.WaterMark, err
	default:
		err = errors.New(fmt.Sprintf("invalid key: %s", key))
		return nil, err
	}
}

func (phone phoneInfo) ListParam() Params {
	p := make(Params)

	t := reflect.TypeOf(phone)
	v := reflect.ValueOf(phone)

	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).IsZero() {
			tagName := strings.Split(string(t.Field(i).Tag), "\"")[1]
			p[tagName] = v.Field(i).Interface()
		}
	}
	return p
}

func (m *myPayer) GetUserPhoneForMini(code string, dataStr string, ivStr string) (ResultParam, error) {
	session, err := m.GetSessionKeyAndOpenId(code)
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

	var info *phoneInfo
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

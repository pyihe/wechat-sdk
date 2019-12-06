package dev

import (
	"encoding/base64"
	"encoding/json"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	微信小程序获取用户信息
*/

//小程序获取用户信息
func (m *myPayer) GetUserInfoForMini(code string, dataStr string, ivStr string) (Param, error) {
	session, err := m.GetSessionKeyAndOpenId(code)
	if err != nil {
		return nil, err
	}
	sessionKey := session.Get("session_key")
	key, err := base64.StdEncoding.DecodeString(sessionKey.(string))
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

	var info interface{}
	err = json.Unmarshal(realData, &info)
	if err != nil {
		return nil, err
	}
	result := util.Interface2Map(info)
	if appId := result["appid"]; appId == nil || appId.(string) != m.appId {
		return nil, e.ErrAppId
	}
	return result, nil
}

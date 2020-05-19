package wechat_sdk

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

//微信小程序获取用户电话号码
func (m *myPayer) GetUserPhoneForMini(code string, dataStr string, ivStr string) (Param, error) {
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

//通过code获取session_key（微信小程序根据code获取openId和session_key）
func (m *myPayer) GetSessionKeyAndOpenId(code string) (Param, error) {
	if err := m.checkForAccess(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code", m.appId, m.secret, code)
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

	if errCode := result.Get("errcode"); errCode != nil && errCode.(float64) != 0 {
		return nil, errors.New(result.Get("errmsg").(string))
	}
	return result, nil
}

//微信小程序获取AccessToken
func (m *myPayer) GetAccessTokenForMini() (Param, error) {
	if err := m.checkForAccess(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", m.appId, m.secret)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result Param
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

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

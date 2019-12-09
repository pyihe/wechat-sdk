package dev

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

/*
	微信小程序根据code获取openId和session_key
*/

//通过code获取session_key
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

	if errCode := result.Get("errcode"); errCode != nil && errCode.(int) != 0 {
		return nil, errors.New(result.Get("errmsg").(string))
	}
	return result, nil
}

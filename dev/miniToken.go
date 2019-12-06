package dev

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	微信小程序获取AccessToken
*/

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

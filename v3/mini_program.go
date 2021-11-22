package v3

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/utils"
	"github.com/pyihe/secret"
)

// GetBaseAccessTokenByMiniProgram 微信小程序获取全局唯一的后台接口调用凭证(access_token)
// access_token有效期一般为7200秒, 需要手动刷新并且设置被动刷新机制(再次调用即可刷新)
// 返回成功实例:
// {"access_token":"ACCESS_TOKEN","expires_in":7200}
// 接口详细描述: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (we *weChatClient) GetBaseAccessTokenByMiniProgram() (Param, error) {
	if we.secret == "" {
		return nil, ErrNoSecret
	}
	var url = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", we.appId, we.secret)
	response, err := we.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result = NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err == nil {
		return result, nil
	}
	var errCode int64
	var errMsg string
	if codeData, ok := result.Get("errcode"); ok {
		errCode = int64(codeData.(float64))
	}
	if msgData, ok := result.Get("errmsg"); ok {
		errMsg = msgData.(string)
	}
	return nil, errors.NewWithCode(errMsg, errors.NewErrCode(errCode))
}

// GetOpenIdByMiniProgram 微信小程序登录验证，同时可以获取用户OpenId
// 需要传递的参数:
// jsCode: 小程序前端授权获取的js授权码
// 返回的字段有:
// openid: 用户唯一在该小程序下的唯一标示
// session_key: 会话密钥, 用于解密用户数据信息
// unionid: 用户在开放平台的唯一标示
// 接口详细介绍页面: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (we *weChatClient) GetOpenIdByMiniProgram(jsCode string) (Param, error) {
	if we.secret == "" {
		return nil, ErrNoSecret
	}
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", we.appId, we.secret, jsCode)
	response, err := we.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result = NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err == nil {
		return result, nil
	}
	var errCode int64
	var errMsg string
	if codeData, ok := result.Get("errcode"); ok {
		errCode = int64(codeData.(float64))
	}
	if msgData, ok := result.Get("errmsg"); ok {
		errMsg = msgData.(string)
	}
	return nil, errors.NewWithCode(errMsg, errors.NewErrCode(errCode))
}

// CheckMiniProgramEncryptedData 检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
// 需要传递的参数:
// accessToken: 小程序全局唯一的后台接口调用凭证
// encryptedMsgHash: 加密数据的sha256，通过Hex（Base16）编码后的字符串
// 接口详细介绍页面: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.checkEncryptedData.html
func (we *weChatClient) CheckMiniProgramEncryptedData(accessToken, encryptedMsgHash string) (bool, error) {
	var url = fmt.Sprintf("https://api.weixin.qq.com/wxa/business/checkencryptedmsg?access_token=%s", accessToken)
	var param = NewParam()
	param.Add("encrypted_msg_hash", encryptedMsgHash)

	bytesData, _ := json.Marshal(param)
	response, err := we.httpClient.Post(url, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	var result = NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return false, err
	}
	var errCode int64
	var errMsg string
	var valid bool
	if checkData, ok := result.Get("vaild"); ok {
		valid = checkData.(bool)
	}
	if codeData, ok := result.Get("errcode"); ok {
		errCode = int64(codeData.(float64))
	}
	if msgData, ok := result.Get("errmsg"); ok {
		errMsg = msgData.(string)
	}
	if errCode != 0 || errMsg != "ok" {
		return false, errors.NewWithCode(errMsg, errors.NewErrCode(errCode))
	}
	return valid, nil
}

// DecryptMiniProgramData 用于解密微信小程序的敏感加密数据，如用户信息、用户手机号码等
// 需要传递的参数有:
// code: 小程序前端的授权码
// encryptedData: 加密数据
// ivStr: iv加密向量
// 返回结果包含的字段根据加密信息的不同而不同，但都是JSON数据格式
// 手机号码的返回结果为:
// {
//    "phoneNumber": "13580006666", // 用户绑定的手机号（国外手机号会有区号）
//    "purePhoneNumber": "13580006666", // 没有区号的手机号
//    "countryCode": "86", // 区号
//    "watermark":
//    {
//        "appid":"APPID",
//        "timestamp": TIMESTAMP
//    }
// }
// 用户信息的解密结果为:
// {
//  "nickName": "Band", // 微信昵称
//  "gender": 1, // 性别
//  "language": "zh_CN", // 区域语言
//  "city": "Guangzhou", // 城市
//  "province": "Guangdong", // 省份
//  "country": "CN", // 国家代码
//  "avatarUrl": "http://wx.qlogo.cn/mmopen/vi_32/1vZvI39NWFQ9XM4LtQpFrQJ1xlgZxx3w7bQxKARol6503Iuswjjn6nIGBiaycAjAtpujxyzYsrztuuICqIM5ibXQ/0" // 头像
// }
func (we *weChatClient) DecryptMiniProgramData(code, encryptedData, ivStr string) (Param, error) {
	openData, err := we.GetOpenIdByMiniProgram(code)
	if err != nil {
		return nil, err
	}
	session, ok := openData.Get("session_key")
	if !ok {
		return nil, ErrInvalidCode
	}
	sessionKey, err := base64.StdEncoding.DecodeString(session.(string))
	if err != nil {
		return nil, err
	}
	iv, err := base64.StdEncoding.DecodeString(ivStr)
	if err != nil {
		return nil, err
	}
	var decryptRequest = &secret.SymRequest{
		CipherData:  encryptedData,
		Key:         sessionKey,
		Type:        secret.SymTypeAES,
		ModeType:    secret.BlockModeCBC,
		PaddingType: secret.PaddingTypePKCS7,
		Iv:          iv,
	}
	plainData, err := secret.NewCipher().SymDecrypt(decryptRequest)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(plainData, &data)
	if err != nil {
		return nil, err
	}
	result := utils.Interface2Map(data.(map[string]interface{}))
	return result, nil
}

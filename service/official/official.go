package official

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/maps"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// GetBaseAccessToken 公众号和小程序获取全局唯一接口调用凭证access_token
// 返回的结果包含两个字段, 分别为:
// access_token: 接口调用凭证, 有效期为2小时，需要手动刷新同时必须包含被动刷新机制(再次获取即可刷新)
// expire_in: 凭证有效时长, 单位: s(秒)
// 返回成功实例: {"access_token": "ACCESS_TOKEN", "expire_in": 7200}
// 接口详细介绍页面: https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func GetBaseAccessToken(config *service.Config) (result maps.Param, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if config.GetSecret() == "" {
		err = service.ErrNoSecret
		return
	}
	if config.GetAppId() == "" {
		err = service.ErrNoAppId
		return
	}

	var url = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", config.GetAppId(), config.GetSecret())
	response, err := config.Request(http.MethodGet, url, service.ContentTypeJSON, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result = maps.NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err == nil {
		return result, nil
	}
	errCode, _ := result.GetInt64("errcode")
	errMsg, _ := result.GetString("errmsg")
	if errMsg == "ok" {
		return result, nil
	}
	return nil, errors.NewWithCode(errMsg, errors.NewErrCode(errCode))
}

// GetOpenId 公众号授权获取用户OpenId
// 需要传递的参数为:
// grantCode: 用户授权网页后获取的授权码code
// appid, secret: 创建实例时传递，接口调用时不需要再传递了
// 返回的结果包含五个字段, 分别为:
// access_token: 网页授权access_token, 后续接口(如获取用户信息)调用凭证
// expire_in: access_token有效时间, 单位: s(秒)
// refresh_token: 用于刷新access_token的token
// openid: 用户在该公众号下的唯一标示
// scope: 用户授权的作用域
// 返回成功实例:
// {
//  "access_token":"ACCESS_TOKEN",
//  "expires_in":7200,
//  "refresh_token":"REFRESH_TOKEN",
//  "openid":"OPENID",
//  "scope":"SCOPE"
//}
// 接口详情介绍页面：https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func GetOpenId(config *service.Config, grantCode string) (result maps.Param, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if config.GetSecret() == "" {
		err = service.ErrNoSecret
		return
	}
	if config.GetAppId() == "" {
		err = service.ErrNoAppId
		return
	}
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", config.GetAppId(), config.GetSecret(), grantCode)
	response, err := config.Request(http.MethodGet, url, service.ContentTypeJSON, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result = maps.NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err == nil {
		return result, nil
	}
	errCode, _ := result.GetInt64("errcode")
	errMsg, _ := result.GetString("errmsg")
	if errMsg == "ok" {
		return result, nil
	}
	return nil, errors.NewWithCode(errMsg, errors.NewErrCode(errCode))
}

// RefreshOauthAccessToken 刷新公众号内网页授权获取的access_token
// 需要传递的参数为:
// refresh_token: 通过GetOpenIdByOfficialAccounts接口获取到的refresh_token
// appid创建实例时已传递，不需要再次传递
// 返回的结果包含5个字段:
// access_token: 网页授权access_token, 后续接口(如获取用户信息)调用凭证
// expire_in: access_token有效时间, 单位: s(秒)
// refresh_token: 用于刷新access_token的token
// openid: 用户在该公众号下的唯一标示
// scope: 用户授权的作用域
// 返回成功实例:
// {
//  "access_token":"ACCESS_TOKEN",
//  "expires_in":7200,
//  "refresh_token":"REFRESH_TOKEN",
//  "openid":"OPENID",
//  "scope":"SCOPE"
//}
// 接口详情介绍页面: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func RefreshOauthAccessToken(config *service.Config, refreshToken string) (result maps.Param, err error) {
	if config.GetAppId() == "" {
		err = service.ErrNoAppId
		return
	}
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s", config.GetAppId(), refreshToken)
	response, err := config.Request(http.MethodGet, url, service.ContentTypeJSON, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result = maps.NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err == nil {
		return result, nil
	}
	errCode, _ := result.GetInt64("errcode")
	errMsg, _ := result.GetString("errmsg")
	if errMsg == "ok" {
		return result, nil
	}
	return nil, errors.NewWithCode(errMsg, errors.NewErrCode(errCode))
}

// GetUserInfo 公众号获取用户基本信息
// 需要传递的参数为:
// accessToken: 通过GetOpenIdByOfficialAccounts接口获取到的access_token
// openId: 通过GetUserInfoByOfficialAccounts接口获取到的用户openid
// lang: 返回的国家地区语言版本, zh_CN: 简体中文, zh_TW: 繁体中文, en: 英语
// 返回结果实例如下:
// {
//  "openid": "OPENID", // 用户openid
//  "nickname": NICKNAME, // 用户微信昵称
//  "sex": 1, // 用户性别, 1: 男性, 2: 女性
//  "province":"PROVINCE", // 用户所微信所在省份
//  "city":"CITY", // 用户微信所在城市
//  "country":"COUNTRY", // 用户微信所在国家
//  "headimgurl":"https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46", // 用户微信头像
//  "privilege":[ "PRIVILEGE1" "PRIVILEGE2"     ], // 用户特权信息
//  "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL" // 用户在同一商户不同公众号的唯一标示
//}
// 接口详细介绍页面: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func GetUserInfo(accessToken, openId, lang string) (result maps.Param, err error) {
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s", accessToken, openId, lang)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result = maps.NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err == nil {
		return result, nil
	}
	errCode, _ := result.GetInt64("errcode")
	errMsg, _ := result.GetString("errmsg")
	if errMsg == "ok" {
		return result, nil
	}
	return nil, errors.NewWithCode(errMsg, errors.NewErrCode(errCode))
}

// CheckOauthAccessTokenValid 校验网页授权凭证access_token是否有效
// 需要传递的参数有:
// accessToken: 待校验的access_token
// openId: 用户openid
// 返回成功实例如下:
// { "errcode":0,"errmsg":"ok"}
// 接口详细介绍页面: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func CheckOauthAccessTokenValid(accessToken, openId string) (bool, error) {
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s", accessToken, openId)
	response, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	var result = maps.NewParam()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return false, err
	}
	errCode, _ := result.GetInt64("errcode")
	errMsg, _ := result.GetString("errmsg")
	if errMsg == "ok" {
		return true, nil
	}
	return false, errors.NewWithCode(errMsg, errors.NewErrCode(errCode))
}

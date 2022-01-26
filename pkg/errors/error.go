package errors

import (
	"net/http"
)

const (
	ErrParam ErrorCode = iota + 10000
	ErrNoAppId
	ErrNoSecret
	ErrNoMchId
	ErrNoConfig
	ErrNoCipher
	ErrNoSerialNo
	ErrNoApiV3Key
	ErrNoSDKRequest
	ErrNoCertificate
	ErrNoHttpRequest
	ErrNoHttpResponse
	ErrInvalidResource
	ErrInvalidHashType
	ErrImageFormatType
	ErrVideoFormatType
	ErrInvalidSessionKey
	ErrCheckHashValueFail
	ErrMarshalFailInvalidDataType
)

type ErrorCode int

func New(code int) ErrorCode {
	return ErrorCode(code)
}

func (code ErrorCode) Error() (err string) {
	switch code {
	case http.StatusAccepted:
		err = "Accepted(服务器已接受请求,但尚未处理): 请求尚未处理,请使用原参数重复请求一遍!"
	case http.StatusBadRequest:
		err = "Bad Request(协议或者参数非法): 请根据接口返回的详细信息检查您的程序!"
	case http.StatusUnauthorized:
		err = "Unauthorized(签名验证失败): 请检查签名参数和方法是否都符合签名算法要求!"
	case http.StatusForbidden:
		err = "Forbidden(权限异常): 请开通商户号相关权限。请联系产品或商务申请!"
	case http.StatusNotFound:
		err = "Not Found(请求的资源不存在): 请商户检查需要查询的id或者请求URL是否正确!"
	case http.StatusTooManyRequests:
		err = "Too Many Requests(请求超过频率限制): 请求未受理，请降低频率后重试!"
	case http.StatusInternalServerError:
		err = "Server Error(系统错误): 按具体接口的错误指引进行重试!"
	case http.StatusBadGateway:
		err = "Bad Gateway(服务下线，暂时不可用): 请求无法处理，请稍后重试!"
	case http.StatusServiceUnavailable:
		err = "Service Unavailable(服务不可用，过载保护): 请求无法处理，请稍后重试!"
	case ErrNoAppId:
		err = "Config缺少参数AppId!"
	case ErrNoSecret:
		err = "Config缺少参数Secret!"
	case ErrNoSerialNo:
		err = "Config缺少参数SerialNo"
	case ErrNoMchId:
		err = "Config缺少参数MchId!"
	case ErrNoApiV3Key:
		err = "Config缺少参数API v3密钥!"
	case ErrInvalidSessionKey:
		err = "未获取到微信SessionKey!"
	case ErrNoConfig:
		err = "请先初始化Config!"
	case ErrInvalidResource:
		err = "未获取到通知资源数据!"
	case ErrInvalidHashType:
		err = "不支持的Hash类型!"
	case ErrNoCertificate:
		err = "获取微信平台公钥证书失败: 证书尚未加载或已过期!"
	case ErrNoSDKRequest:
		err = "请求不能为空!"
	case ErrNoHttpRequest:
		err = "解析微信通知失败: http.Request为空!"
	case ErrNoHttpResponse:
		err = "解析微信应答失败: http.Response为空!"
	case ErrCheckHashValueFail:
		err = "Hash摘要值校验不通过!"
	case ErrMarshalFailInvalidDataType:
		err = "序列化失败: 不支持的数据类型!"
	case ErrNoCipher:
		err = "加解密/验签失败: 请先初始化Cipher!"
	case ErrParam:
		err = "参数错误: 缺少必要参数!"
	case ErrImageFormatType:
		err = "请提供正确格式的图片!"
	case ErrVideoFormatType:
		err = "请提供正确格式的视频!"
	default:
		err = "发生未知错误!"
	}
	return
}

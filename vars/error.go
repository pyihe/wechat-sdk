package vars

import "github.com/pyihe/go-pkg/errors"

var (
	ErrNoAppId           = errors.New("请提供appid!")
	ErrNoRequest         = errors.New("请求参数为空!")
	ErrNoSecret          = errors.New("请提供secret!")
	ErrNoSerialNo        = errors.New("请提供商户证书序列号!")
	ErrNoMchId           = errors.New("请提供商户号!")
	ErrNoApiV3Key        = errors.New("请提供商户API密钥!")
	ErrInvalidSessionKey = errors.New("获取session_key失败!")
	ErrRequestAgain      = errors.New("请稍后再次请求!")
	ErrInvalidCipher     = errors.New("请初始化cipher!")
	ErrInitConfig        = errors.New("请初始化Config!")
)

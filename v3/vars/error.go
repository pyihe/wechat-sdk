package vars

import "github.com/pyihe/go-pkg/errors"

var (
	ErrNoParam           = errors.New("请求的参数为空!")
	ErrNoSecret          = errors.New("WeChatClient缺少secret!")
	ErrNoSerialNo        = errors.New("WeChatClient缺少证书序列号!")
	ErrNoMchId           = errors.New("WeChatClient缺少商户号!")
	ErrNoApiV3Key        = errors.New("WeChatClient缺少API Key!")
	ErrInvalidSessionKey = errors.New("获取session_key失败!")
	ErrNotExistKey       = errors.New("key不存在!")
	ErrConvert           = errors.New("interface{}类型转换失败!")
	ErrRequestAgain      = errors.New("请稍后再次请求!")
)

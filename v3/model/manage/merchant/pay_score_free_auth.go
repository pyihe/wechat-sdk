package merchant

import (
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/model"
)

type PrePermissionRequest struct {
	ServiceId         string `json:"service_id,omitempty"`         // 服务ID
	AppId             string `json:"appid,omitempty"`              // 应用ID
	AuthorizationCode string `json:"authorization_code,omitempty"` // 授权协议号
	NotifyUrl         string `json:"notify_url,omitempty"`         // 商户接收授权回调通知的地址
}

func (p *PrePermissionRequest) Check() (err error) {
	if p.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if p.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if p.AuthorizationCode == "" {
		err = errors.New("请填写authorization_code!")
		return
	}
	return
}

type PrePermissionResponse struct {
	model.WechatError
	RequestId            string `json:"-"`                                // 请求唯一ID
	ApplyPermissionToken string `json:"apply_permission_token,omitempty"` // 预授权token
}

type QueryPermissionRequest struct {
	ServiceId         string `json:"service_id,omitempty"`         // 服务ID
	AuthorizationCode string `json:"authorization_code,omitempty"` // 授权协议号
	AppId             string `json:"appid,omitempty"`              // 应用ID
	OpenId            string `json:"openid,omitempty"`             // 用户标识
}

func (query *QueryPermissionRequest) Check() (err error) {
	if query.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if query.AuthorizationCode != "" {
		return
	}
	if query.AppId != "" && query.OpenId != "" {
		return
	}
	err = errors.New("请补全查询条件: 可以通过authorization_code或者openid查询, 参数要求参考文档!")
	return
}

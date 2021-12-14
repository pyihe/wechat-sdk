package merchant

import (
	"time"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/model"
)

/******************************商户预授权********************************************************************************/

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

/************************查询用户授权记录*********************************************************************************/

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

type QueryPermissionResponse struct {
	model.WechatError
	RequestId                string    `json:"-"`                                    // 请求唯一ID
	ServiceId                string    `json:"service_id,omitempty"`                 // 服务ID
	AppId                    string    `json:"appid,omitempty"`                      // 应用ID
	MchId                    string    `json:"mchid,omitempty"`                      // 商户号
	OpenId                   string    `json:"openid,omitempty"`                     // 用户标识
	AuthorizationCode        string    `json:"authorization_code,omitempty"`         // 授权协议号
	AuthorizationState       string    `json:"authorization_state,omitempty"`        // 授权状态
	NotifyUrl                string    `json:"notify_url,omitempty"`                 // 授权通知地址
	CancelAuthorizationTime  time.Time `json:"cancel_authorization_time,omitempty"`  // 最近一次解除授权时间
	AuthorizationSuccessTime time.Time `json:"authorization_success_time,omitempty"` // 最近一次授权成功时间
}

/************************终止授权服务************************************************************************************/

type TerminatePermissionRequest struct {
	ServiceId string `json:"service_id,omitempty"` // 服务ID
	Reason    string `json:"reason,omitempty"`     // 解除授权的原因

	AuthorizationCode string `json:"-"` // 通过授权协议号解除授权的授权协议号

	OpenId string `json:"-"`               // 通过openid解除授权的openid
	AppId  string `json:"appid,omitempty"` // 应用ID
}

func (t *TerminatePermissionRequest) Check() (err error) {
	if t.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if t.Reason == "" {
		err = errors.New("请填写reason!")
		return
	}
	if t.AuthorizationCode != "" {
		if t.OpenId != "" {
			err = errors.New("通过authorization_code解除授权时请勿填写openid!")
			return
		}
		if t.AppId != "" {
			err = errors.New("通过authorization_code解除授权时请勿填写appid!")
			return
		}
		return
	}
	if t.OpenId != "" {
		if t.AppId == "" {
			err = errors.New("通过openid解除授权时请填写appid!")
			return
		}
		if t.AuthorizationCode != "" {
			err = errors.New("通过openid解除授权时请勿填写authorization_code!")
			return
		}
		return
	}
	err = errors.New("请补全解除授权条件: 可以通过authorization_code或者openid解除授权, 参数要求参考文档!")
	return
}

type TerminatePermissionResponse struct {
	model.WechatError
	RequestId string `json:"request_id,omitempty"`
}

/****************************开启//解除授权服务回调通知*********************************************************************/

type OpenOrCloseResponse struct {
	Id                string    `json:"id,omitempty"`                  // 通知ID
	AppId             string    `json:"appid,omitempty"`               // 公众账号ID
	MchId             string    `json:"mchid,omitempty"`               // 商户号
	OutRequestNo      string    `json:"out_request_no,omitempty"`      // 商户签约单号
	ServiceId         string    `json:"service_id,omitempty"`          // 服务ID
	OpenId            string    `json:"openid,omitempty"`              // 用户标识
	UserServiceStatus string    `json:"user_service_status,omitempty"` // 回调状态
	OpenOrCloseTime   time.Time `json:"openorclose_time,omitempty"`    // 服务开启/解除授权时间
	AuthorizationCode string    `json:"authorization_code,omitempty"`  // 授权协议号
}

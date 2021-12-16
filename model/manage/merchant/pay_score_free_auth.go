package merchant

import (
	"time"

	"github.com/pyihe/wechat-sdk/vars"

	"github.com/pyihe/wechat-sdk/model/manage"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/model"
)

/******************************商户预授权********************************************************************************/

type PrePermitRequest struct {
	ServiceId         string `json:"service_id,omitempty"`         // 服务ID
	AppId             string `json:"appid,omitempty"`              // 应用ID
	AuthorizationCode string `json:"authorization_code,omitempty"` // 授权协议号
	NotifyUrl         string `json:"notify_url,omitempty"`         // 商户接收授权回调通知的地址
}

func (p *PrePermitRequest) Check() (err error) {
	if p == nil {
		err = vars.ErrNoRequest
		return
	}
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
	if query == nil {
		err = vars.ErrNoRequest
		return
	}
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
	if t == nil {
		err = vars.ErrNoRequest
		return
	}
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

/****************************开启//解除授权服务回调通知********************************************************************/

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

/*************************************订单确认回调通知********************************************************************/

type ConfirmOrderResponse struct {
}

/*************************************公共API: 创建支付分订单*************************************************************/

type CreatePayscoreOrderRequest struct {
	OutOrderNo          string                  `json:"out_order_no,omitempty"`         // 商户服务订单号
	AppId               string                  `json:"appid,omitempty"`                // 应用ID
	ServiceId           string                  `json:"service_id,omitempty"`           // 服务ID
	ServiceIntroduction string                  `json:"service_introduction,omitempty"` // 服务信息
	PostPayments        []*manage.PostPayments  `json:"post_payments,omitempty"`        // 后付费项目
	PostDiscounts       []*manage.PostDiscounts `json:"post_discounts,omitempty"`       // 后服务商户优惠
	TimeRange           *manage.TimeRange       `json:"time_range,omitempty"`           // 服务时间范围
	Location            *manage.Location        `json:"location,omitempty"`             // 服务位置
	RiskFund            *manage.RiskFund        `json:"risk_fund,omitempty"`            // 订单风险金
	Attach              string                  `json:"attach,omitempty"`               // 商户数据包
	NotifyUrl           string                  `json:"notify_url,omitempty"`           // 商户接收用户确认订单和付款成功的回调地址
	OpenId              string                  `json:"openid,omitempty"`               // 用户标识
	NeedUserConfirm     bool                    `json:"need_user_confirm"`              // 是否需要用户确认
}

func (c *CreatePayscoreOrderRequest) Check() (err error) {
	if c == nil {
		err = vars.ErrNoRequest
		return
	}
	if c.OutOrderNo == "" {
		err = errors.New("请填写out_order_no!")
		return
	}
	if c.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if c.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if c.ServiceIntroduction == "" {
		err = errors.New("请填写service_introduction!")
		return
	}
	if len(c.PostPayments) > 100 {
		err = errors.New("post_payments的上限为100!")
		return
	}
	for _, p := range c.PostPayments {
		if p.Name != "" {
			if p.Amount < 0 && p.Description == "" {
				err = errors.New("post_payments.name不为空时post_payments.amount和post_payments.description不能同时为空!")
				return
			}
		}
		if p.Count > 100 {
			err = errors.New("post_payments.count的上限为100!")
			return
		}
	}
	if len(c.PostDiscounts) > 30 {
		err = errors.New("post_discounts的数量上限为30!")
		return
	}
	for _, p := range c.PostDiscounts {
		if (p.Name != "" && p.Description == "") || (p.Name == "" && p.Description != "") {
			err = errors.New("post_discounts.name和post_discounts.description必须同时填写或者都不填写!")
			return
		}
		if p.Count > 100 {
			err = errors.New("post_discounts的上限为100!")
			return
		}
	}
	if c.TimeRange == nil {
		err = errors.New("请填写time_range!")
		return
	}
	if c.TimeRange.StartTime == "" {
		err = errors.New("请填写time_range.start_time!")
		return
	}
	if c.RiskFund == nil {
		err = errors.New("请填写risk_fund!")
		return
	}
	if c.RiskFund.Name == "" {
		err = errors.New("请填写risk_fund.name!")
		return
	}
	if c.RiskFund.Amount <= 0 {
		err = errors.New("请填写risk_fund.amount!")
		return
	}
	if c.NotifyUrl == "" {
		err = errors.New("请填写notify_url!")
		return
	}
	return
}

type PayscoreOrder struct {
	model.WechatError
	Id                  string                  `json:"-"`                              // 请求或者通知的唯一ID，如果有的话
	AppId               string                  `json:"app_id,omitempty"`               // 应用ID
	MchId               string                  `json:"mch_id,omitempty"`               // 商户号ID
	OutOrderNo          string                  `json:"out_order_no,omitempty"`         // 商户服务订单号
	ServiceId           string                  `json:"service_id,omitempty"`           // 服务ID
	ServiceIntroduction string                  `json:"service_introduction,omitempty"` // 服务信息
	State               string                  `json:"state,omitempty"`                // 订单状态
	StateDescription    string                  `json:"state_description,omitempty"`    // 订单状态说明
	TotalAmount         int64                   `json:"total_amount,omitempty"`         // 商户收款总金额
	PostPayments        []*manage.PostPayments  `json:"post_payments,omitempty"`        // 后付费项目列表
	PostDiscounts       []*manage.PostDiscounts `json:"post_discounts,omitempty"`       // 后付费优惠项目列表
	RiskFund            *manage.RiskFund        `json:"risk_fund,omitempty"`            // 订单风险金信息
	TimeRange           *manage.TimeRange       `json:"time_range,omitempty"`           // 服务时间范围
	Location            *manage.Location        `json:"location,omitempty"`             // 服务位置
	Attach              string                  `json:"attach,omitempty"`               // 商户数据包
	NotifyUrl           string                  `json:"notify_url,omitempty"`           // 商户回调地址
	OrderId             string                  `json:"order_id,omitempty"`             // 微信支付服务订单号
	Package             string                  `json:"package,omitempty"`              // 跳转微信侧小程序订单数据
	NeedCollection      bool                    `json:"need_collection,omitempty"`      // 是否需要收款
	Collection          *manage.Collection      `json:"collection,omitempty"`           // 收款信息
	OpenId              string                  `json:"openid,omitempty"`               // 用户标识
}

/*************************************公共API: 查询支付分订单*************************************************************/

type QueryPayscoreOrderRequest struct {
	OutOrderNo string `json:"out_order_no,omitempty"` // 商户服务订单号
	QueryId    string `json:"query_id,omitempty"`     // 回跳查询ID
	ServiceId  string `json:"service_id,omitempty"`   // 服务ID
	AppId      string `json:"appid,omitempty"`        // 应用ID
}

func (query *QueryPayscoreOrderRequest) Check() (err error) {
	if query == nil {
		err = vars.ErrNoRequest
		return
	}
	if query.OutOrderNo == "" && query.QueryId == "" {
		err = errors.New("out_order_no和query_id不能同时为空!")
		return
	}
	if query.OutOrderNo != "" && query.QueryId != "" {
		err = errors.New("out_order_no和query_id只能提供一项!")
		return
	}
	if query.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if query.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	return
}

/*************************************公共API: 取消支付分订单*************************************************************/

type CancelPayscoreOrderRequest struct {
	OutOrderNo string `json:"-"`                    // 商户服务订单号
	AppId      string `json:"appid,omitempty"`      // 应用ID
	ServiceId  string `json:"service_id,omitempty"` // 服务ID
	Reason     string `json:"reason,omitempty"`     // 取消原因
}

func (c *CancelPayscoreOrderRequest) Check() (err error) {
	if c == nil {
		err = vars.ErrNoRequest
		return
	}
	if c.OutOrderNo == "" {
		err = errors.New("请填写out_order_no!")
		return
	}
	if c.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if c.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if c.Reason == "" {
		err = errors.New("请填写reason!")
		return
	}
	return
}

type CancelResponse struct {
	model.WechatError
	RequestId  string `json:"request_id,omitempty"`   // 请求唯一ID
	AppId      string `json:"appid,omitempty"`        // 应用ID
	MchId      string `json:"mchid,omitempty"`        // 商户号
	OutOrderNo string `json:"out_order_no,omitempty"` // 商户订单号
	ServiceId  string `json:"service_id,omitempty"`   // 服务ID
	OrderId    string `json:"order_id,omitempty"`     // 微信支付服务订单号
}

/*************************************公共API: 修改支付分订单金额**********************************************************/

type ModifyPayscoreOrderRequest struct {
	OutOrderNo    string                  `json:"-"`                        // 商户服务订单号
	AppId         string                  `json:"appid,omitempty"`          // 应用ID
	ServiceId     string                  `json:"service_id,omitempty"`     // 服务ID
	PostPayments  []*manage.PostPayments  `json:"post_payments,omitempty"`  // 后付费项目
	PostDiscounts []*manage.PostDiscounts `json:"post_discounts,omitempty"` // 后付费商户优惠
	TotalAmount   int64                   `json:"total_amount,omitempty"`   // 总金额
	Reason        string                  `json:"reason,omitempty"`         // 修改原因
}

func (m *ModifyPayscoreOrderRequest) Check() (err error) {
	if m == nil {
		err = vars.ErrNoRequest
		return
	}
	if m.OutOrderNo == "" {
		err = errors.New("请填写out_order_no!")
		return
	}
	if m.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if m.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if len(m.PostPayments) == 0 {
		err = errors.New("请填写post_payments!")
		return
	}
	for _, p := range m.PostPayments {
		if p.Name == "" {
			err = errors.New("请填写post_payments.name!")
			return
		}
		if p.Amount < 0 {
			err = errors.New("请填写正确的post_payments.amount!")
			return
		}
		if p.Count > 100 {
			err = errors.New("post_payments.count数量限制为100!")
			return
		}
	}

	for _, p := range m.PostDiscounts {
		if (p.Name != "" && p.Description == "") || (p.Name == "" && p.Description != "") {
			err = errors.New("post_discounts.name和post_discounts.description必须同时写或者都不填写!")
			return
		}
	}
	if m.TotalAmount <= 0 {
		err = errors.New("请填写正确的total_amount!")
		return
	}
	if m.Reason == "" {
		err = errors.New("请填写reason!")
		return
	}
	return
}

/*************************************公共API: 完结支付分订单*************************************************************/

type CompletePayscoreOrderRequest struct {
	OutOrderNo    string                  `json:"-"`                        // 商户服务订单号
	AppId         string                  `json:"appid,omitempty"`          // 应用ID
	ServiceId     string                  `json:"service_id,omitempty"`     // 服务ID
	PostPayments  []*manage.PostPayments  `json:"post_payments,omitempty"`  // 后付费项目
	PostDiscounts []*manage.PostDiscounts `json:"post_discounts,omitempty"` // 后付费商户优惠
	TotalAmount   int64                   `json:"total_amount"`             // 总金额
	TimeRange     *manage.TimeRange       `json:"time_range,omitempty"`     // 服务时间段
	Location      *manage.Location        `json:"location,omitempty"`       // 服务位置
	ProfitSharing bool                    `json:"profit_sharing,omitempty"` // 微信支付服务分账标记
	GoodsTag      string                  `json:"goods_tag,omitempty"`      // 订单优惠标记
}

func (f *CompletePayscoreOrderRequest) Check() (err error) {
	if f == nil {
		err = vars.ErrNoRequest
		return
	}
	if f.OutOrderNo == "" {
		err = errors.New("请填写out_order_no!")
		return
	}
	if f.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if f.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if len(f.PostPayments) == 0 {
		err = errors.New("请填写post_payments!")
		return
	}
	for _, p := range f.PostPayments {
		if p.Name == "" {
			err = errors.New("请填写post_payments.name!")
			return
		}
		if p.Amount < 0 {
			err = errors.New("请填写正确的post_payments.amount!")
			return
		}
		if p.Count > 100 {
			err = errors.New("post_payments.count上限数量为100!")
			return
		}
	}
	for _, p := range f.PostDiscounts {
		if p.Name == "" {
			err = errors.New("请填写post_discounts.name!")
			return
		}
		if p.Description == "" {
			err = errors.New("请填写post_discounts.description!")
			return
		}
		if p.Count > 100 {
			err = errors.New("post_discounts.count数量上限为100!")
			return
		}
	}
	if f.TotalAmount < 0 {
		err = errors.New("请填写正确的total_amount!")
		return
	}
	return
}

/*************************************公共API: 商户发起催收扣款***********************************************************/

type PayscorePayRequest struct {
	OutOrderNo string `json:"-"`                    // 商户服务订单号
	AppId      string `json:"appid,omitempty"`      // 应用ID
	ServiceId  string `json:"service_id,omitempty"` // 服务ID
}

func (p *PayscorePayRequest) Check() (err error) {
	if err == nil {
		err = vars.ErrNoRequest
		return
	}
	if p.OutOrderNo == "" {
		err = errors.New("请填写out_order_no!")
		return
	}
	if p.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if p.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	return
}

type PayscorePayResponse struct {
	model.WechatError
	RequestId  string `json:"request_id,omitempty"`   // 唯一请求ID
	AppId      string `json:"appid,omitempty"`        // 应用ID
	MchId      string `json:"mchid,omitempty"`        // 商户号
	OutOrderNo string `json:"out_order_no,omitempty"` // 商户订单号
	ServiceId  string `json:"service_id,omitempty"`   // 服务ID
	OrderId    string `json:"order_id,omitempty"`     // 微信支付服务订单号
}

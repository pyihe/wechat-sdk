package payscore

import (
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
)

// PrePermitResponse 商户预授权应答
type PrePermitResponse struct {
	model.WechatError
	RequestId             string
	ApplyPermissionsToken string `json:"apply_permissions_token,omitempty"` // 预授权Token
}

// QueryPermissionsRequest 查询用户授权记录请求
type QueryPermissionsRequest struct {
	ServiceId         string `json:"service_id,omitempty"`         // 服务ID
	AuthorizationCode string `json:"authorization_code,omitempty"` // 授权协议号
	AppId             string `json:"appid,omitempty"`              // 应用ID
	OpenId            string `json:"openid,omitempty"`             // 用户标示
}

// QueryPermissionsResponse 查询用户授权记录应答（包括通过授权协议号和openid查询）
type QueryPermissionsResponse struct {
	model.WechatError
	RequestId                string
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

// TerminatePermissionRequest 解除用户授权关系请求
type TerminatePermissionRequest struct {
	ServiceId         string `json:"service_id"`      // 服务ID
	Reason            string `json:"reason"`          // 解除授权原因
	AuthorizationCode string `json:"-"`               // 授权协议号
	OpenId            string `json:"-"`               // 用户标识
	AppId             string `json:"appid,omitempty"` // 应用ID
}

// TerminatePermissionResponse 解除用户授权关系应答
type TerminatePermissionResponse struct {
	model.WechatError
	RequestId string
}

// OpenOrCloseResponse 开启/解除授权服务回调通知内容
type OpenOrCloseResponse struct {
	Id                string    // 通知ID
	AppId             string    `json:"appid,omitempty"`               // 公众号ID
	MchId             string    `json:"mchid,omitempty"`               // 商户号
	OutRequestNo      string    `json:"out_request_no,omitempty"`      // 商户签约单号
	ServiceId         string    `json:"service_id,omitempty"`          // 服务ID
	OpenId            string    `json:"openid,omitempty"`              // 用户标识
	UserServiceStatus string    `json:"user_service_status,omitempty"` // 回调状态
	OpenOrCloseTime   time.Time `json:"openorclose_time,omitempty"`    // 服务开启/解除授权时间
	AuthorizationCode string    `json:"authorization_code,omitempty"`  // 授权协议号
}

// ServiceOrder 确认订单通知参数
type ServiceOrder struct {
	model.WechatError
	Id                  string          `json:"-"`                              // 微信唯一请求ID
	AppId               string          `json:"appid,omitempty"`                // 应用ID
	MchId               string          `json:"mchid,omitempty"`                // 商户号
	OutOrderNo          string          `json:"out_order_no,omitempty"`         // 商户服务订单号
	OpenId              string          `json:"open_id,omitempty"`              // 用户标识
	ServiceId           string          `json:"service_id,omitempty"`           // 服务ID
	ServiceIntroduction string          `json:"service_introduction,omitempty"` // 服务信息
	State               string          `json:"state,omitempty"`                // 服务订单状态
	StateDescription    string          `json:"state_description,omitempty"`    // 订单状态说明
	TotalAmount         int64           `json:"total_amount,omitempty"`         // 商户收款总金额
	PostPayments        []*PostPayment  `json:"post_payments,omitempty"`        // 后付费项目
	PostDiscounts       []*PostDiscount `json:"post_discounts,omitempty"`       // 后付费商户优惠
	RiskFund            *RiskFund       `json:"risk_fund,omitempty"`            // 风险金
	TimeRange           *TimeRange      `json:"time_range,omitempty"`           // 服务时间段
	Location            *Location       `json:"location,omitempty"`             // 服务位置
	Attach              string          `json:"attach,omitempty"`               // 商户数据包
	NotifyUrl           string          `json:"notify_url,omitempty"`           // 商户回调地址
	OrderId             string          `json:"order_id,omitempty"`             // 微信支付服务订单号
	Package             string          `json:"package,omitempty"`              // 跳转微信侧小程序订单数据
	NeedCollection      bool            `json:"need_collection,omitempty"`      // 是否需要收款
	Collection          *Collection     `json:"collection,omitempty"`           // 收款信息
}

// PostPayment 后付费项目
type PostPayment struct {
	Name        string `json:"name"`                  // 付费项目名称
	Amount      int64  `json:"amount"`                // 金额
	Description string `json:"description,omitempty"` // 计费说明
	Count       int    `json:"count,omitempty"`       // 付费数量
}

// PostDiscount 后付费商户优惠
type PostDiscount struct {
	Name        string `json:"name,omitempty"`        // 优惠名称
	Description string `json:"description,omitempty"` // 优惠说明
	Amount      int64  `json:"amount,omitempty"`      // 优惠金额
}

// RiskFund 订单风险金
type RiskFund struct {
	Name        string `json:"name,omitempty"`        // 风险金名称
	Amount      int64  `json:"amount,omitempty"`      // 风险金额
	Description string `json:"description,omitempty"` // 风险说明
}

// TimeRange 服务时间段
type TimeRange struct {
	StartTime       time.Time `json:"start_time,omitempty"`        // 服务开始时间
	StartTimeRemark string    `json:"start_time_remark,omitempty"` // 服务开始时间备注
	EndTime         time.Time `json:"end_time,omitempty"`          // 服务结束时间
	EndTimeRemark   string    `json:"end_time_remark,omitempty"`   // 服务结束时间备注
}

// Location 服务位置
type Location struct {
	StartLocation string `json:"start_location,omitempty"` // 服务开始地点
	EndLocation   string `json:"end_location,omitempty"`   // 服务结束位置
}

// QueryOrderRequest 查询支付分订单
type QueryOrderRequest struct {
	OutOrderNo string `json:"out_order_no,omitempty"` // 商户服务订单号
	QueryId    string `json:"query_id,omitempty"`     // 回跳查询ID
	ServiceId  string `json:"service_id,omitempty"`   // 服务ID
	AppId      string `json:"appid,omitempty"`        // 应用ID
}

// CancelRequest 取消服务订单请求参数
type CancelRequest struct {
	OutOrderNo string `json:"-"`          // 商户服务订单号
	AppId      string `json:"appid"`      // 应用ID
	ServiceId  string `json:"service_id"` // 服务ID
	Reason     string `json:"reason"`     // 取消原因
}

// CancelResponse 取消服务订单号应答
type CancelResponse struct {
	RequestId  string `json:"-"`                      // 唯一请求ID
	AppId      string `json:"appid,omitempty"`        // 应用ID
	MchId      string `json:"mchid,omitempty"`        // 商户号
	OutOrderNo string `json:"out_order_no,omitempty"` // 商户订单号
	ServiceId  string `json:"service_id,omitempty"`   // 服务ID
	OrderId    string `json:"order_id,omitempty"`     // 微信支付服务订单号
}

// ModifyRequest 修改订单金额
type ModifyRequest struct {
	AppId         string          `json:"appid"`                    // 应用ID
	ServiceId     string          `json:"service_id"`               // 服务ID
	PostPayments  []*PostPayment  `json:"post_payments"`            // 后付费项目
	PostDiscounts []*PostDiscount `json:"post_discounts,omitempty"` // 后付费商户优惠
	TotalAmount   int64           `json:"total_amount"`             // 总金额
	Reason        string          `json:"reason"`                   // 修改原因
}

// ModifyResponse 修改订单金额应答
type ModifyResponse struct {
	model.WechatError
	RequestId           string          `json:"-"`                              // 微信唯一请求ID或者唯一通知ID
	AppId               string          `json:"appid,omitempty"`                // 应用ID
	MchId               string          `json:"mchid,omitempty"`                // 商户号
	ServiceId           string          `json:"service_id,omitempty"`           // 服务ID
	OutOrderNo          string          `json:"out_order_no,omitempty"`         // 商户服务订单号
	State               string          `json:"state,omitempty"`                // 服务订单状态
	StateDescription    string          `json:"state_description,omitempty"`    // 订单状态说明
	TotalAmount         int64           `json:"total_amount,omitempty"`         // 商户收款总金额
	ServiceIntroduction string          `json:"service_introduction,omitempty"` // 服务信息
	PostPayments        []*PostPayment  `json:"post_payments,omitempty"`        // 后付费项目
	PostDiscounts       []*PostDiscount `json:"post_discounts,omitempty"`       // 后付费商户优惠
	RiskFund            *RiskFund       `json:"risk_fund,omitempty"`            // 风险金
	TimeRange           *TimeRange      `json:"time_range,omitempty"`           // 服务时间段
	Location            *Location       `json:"location,omitempty"`             // 服务位置
	Attach              string          `json:"attach,omitempty"`               // 商户数据包
	NotifyUrl           string          `json:"notify_url,omitempty"`           // 商户回调地址
	OrderId             string          `json:"order_id,omitempty"`             // 微信支付服务订单号
	NeedCollection      bool            `json:"need_collection,omitempty"`      // 是否需要收款
	Collection          *Collection     `json:"collection,omitempty"`           // 收款信息
}

// Collection 收款信息
type Collection struct {
	State        string    `json:"state,omitempty"`         // 收款状态
	TotalAmount  int64     `json:"total_amount,omitempty"`  // 总收款金额
	PayingAmount int64     `json:"paying_amount,omitempty"` // 待收款金额
	PaidAmount   int64     `json:"paid_amount,omitempty"`   // 已收款金额
	Details      []*Detail `json:"details,omitempty"`       // 收款明细列表
}

// Detail 收款明细
type Detail struct {
	Seq           uint64    `json:"seq,omitempty"`            // 收款序号
	Amount        int64     `json:"amount,omitempty"`         // 单笔收款金额
	PaidType      string    `json:"paid_type,omitempty"`      // 收款成功渠道
	PaidTime      time.Time `json:"paid_time,omitempty"`      // 收款成功时间
	TransactionId string    `json:"transaction_id,omitempty"` // 微信支付交易单号
}

// CompleteRequest 完结订单请求参数
type CompleteRequest struct {
	AppId         string          `json:"appid"`                    // 应用ID
	ServiceId     string          `json:"service_id"`               // 服务ID
	PostPayments  []*PostPayment  `json:"post_payments"`            // 后付费项目
	PostDiscounts []*PostDiscount `json:"post_discounts,omitempty"` // 后付费商户优惠
	TotalAmount   int64           `json:"total_amount"`             // 总金额
	TimeRange     *TimeRange      `json:"time_range,omitempty"`     // 服务时间段
	Location      *Location       `json:"location,omitempty"`       // 服务位置
	ProfitSharing bool            `json:"profit_sharing,omitempty"` // 微信支付服务分账标记
	GoodsTage     string          `json:"goods_tage,omitempty"`     // 订单优惠标记
}

// CompleteResponse 完结支付分订单应答
type CompleteResponse struct {
	model.WechatError
	RequestId           string          `json:"-"`                              // 唯一请求ID
	AppId               string          `json:"appid,omitempty"`                // 应用ID
	MchId               string          `json:"mchid,omitempty"`                // 商户号
	OutOrderNo          string          `json:"out_order_no,omitempty"`         // 商户服务订单号
	ServiceId           string          `json:"service_id,omitempty"`           // 服务ID
	ServiceIntroduction string          `json:"service_introduction,omitempty"` // 服务信息
	State               string          `json:"state,omitempty"`                // 服务订单状态
	StateDescription    string          `json:"state_description,omitempty"`    // 订单状态说明
	TotalAmount         int64           `json:"total_amount,omitempty"`         // 商户收款总金额
	PostPayments        []*PostPayment  `json:"post_payments,omitempty"`        // 后付费项目
	PostDiscounts       []*PostDiscount `json:"post_discounts,omitempty"`       // 后付费商户优惠
	RiskFund            *RiskFund       `json:"risk_fund,omitempty"`            // 风险金
	TimeRange           *TimeRange      `json:"time_range,omitempty"`           // 服务时间段
	Location            *Location       `json:"location,omitempty"`             // 服务位置
	OrderId             string          `json:"order_id,omitempty"`             // 微信支付服务订单号
	NeedCollection      bool            `json:"need_collection,omitempty"`      // 是否需要收款
}

// PayOrderRequest 商户发起催收扣款请求参数
type PayOrderRequest struct {
	AppId     string `json:"appid"`      // 应用ID
	ServiceId string `json:"service_id"` // 服务ID
}

// PayOrderResponse 商户发起催收扣款应答参数
type PayOrderResponse struct {
	RequestId  string `json:"-"`                      // 唯一请求ID
	AppId      string `json:"appid,omitempty"`        // 应用ID
	MchId      string `json:"mchid,omitempty"`        // 商户号
	OutOrderNo string `json:"out_order_no,omitempty"` // 商户服务订单号
	ServiceId  string `json:"service_id,omitempty"`   // 服务ID
	OrderId    string `json:"order_id,omitempty"`     // 微信支付服务订单号
}

// SyncRequest 同步订单请求参数
type SyncRequest struct {
	AppId     string  `json:"appid"`            // 应用ID
	ServiceId string  `json:"service_id"`       // 服务ID
	Type      string  `json:"type"`             // 场景类型
	Detail    *Detail `json:"detail,omitempty"` // 内容信息详情
}

// SyncDetail 同步订单内容信息详情
type SyncDetail struct {
	PaidTime string `json:"paid_time"` // 收款成功时间
}

// SyncResponse 同步服务订单信息应答参数
type SyncResponse struct {
	RequestId           string          `json:"-"`                              // 唯一请求ID
	AppId               string          `json:"appid,omitempty"`                // 应用ID
	MchId               string          `json:"mchid,omitempty"`                // 商户号
	OutOrderNo          string          `json:"out_order_no,omitempty"`         // 商户服务订单号
	ServiceId           string          `json:"service_id,omitempty"`           // 服务ID
	ServiceIntroduction string          `json:"service_introduction,omitempty"` // 服务信息
	OpenId              string          `json:"open_id,omitempty"`              // 用户标识
	State               string          `json:"state,omitempty"`                // 服务订单状态
	StateDescription    string          `json:"state_description,omitempty"`    // 订单状态说明
	TotalAmount         int64           `json:"total_amount,omitempty"`         // 商户收款总金额
	PostPayments        []*PostPayment  `json:"post_payments,omitempty"`        // 后付费项目
	PostDiscounts       []*PostDiscount `json:"post_discounts,omitempty"`       // 后付费商户优惠
	RiskFund            *RiskFund       `json:"risk_fund,omitempty"`            // 风险金
	TimeRange           *TimeRange      `json:"time_range,omitempty"`           // 服务时间段
	Location            *Location       `json:"location,omitempty"`             // 服务位置
	Attach              string          `json:"attach,omitempty"`               // 商户数据包
	NotifyUrl           string          `json:"notify_url,omitempty"`           // 商户回调地址
	OrderId             string          `json:"order_id,omitempty"`             // 微信支付服务订单号
	NeedCollection      bool            `json:"need_collection,omitempty"`      // 是否需要收款
	Collection          *Collection     `json:"collection,omitempty"`           // 收款信息
}

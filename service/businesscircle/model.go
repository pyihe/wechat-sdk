package businesscircle

import (
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
)

// SyncPointsResponse 商圈积分同步应答参数
type SyncPointsResponse struct {
	model.WechatError
	RequestId string `json:"-"` // 唯一请求ID
}

// QueryUserAuthorizationRequest 商圈积分授权查询请求参数
type QueryUserAuthorizationRequest struct {
	SubMchId string `json:"sub_mchid,omitempty"` // 商圈商户ID
	AppId    string `json:"appid"`               // 应用ID
	OpenId   string `json:"openid"`              // 用户标识
}

// QueryUserAuthorizationResponse 商圈积分授权查询应答参数
type QueryUserAuthorizationResponse struct {
	model.WechatError
	RequestId       string    `json:"-"`                          // 唯一请求ID
	OpenId          string    `json:"openid,omitempty"`           // 用户标识
	AuthorizeState  string    `json:"authorize_state,omitempty"`  // 授权状态
	AuthorizeTime   time.Time `json:"authorize_time,omitempty"`   // 授权时间
	DeauthorizeTime time.Time `json:"deauthorize_time,omitempty"` // 取消授权时间
}

// RefundResponse 商圈退款成功通知应答参数
type RefundResponse struct {
	Id            string    `json:"-"`                        // 唯一通知ID
	MchId         string    `json:"mchid,omitempty"`          // 商户号
	MerchantName  string    `json:"merchant_name,omitempty"`  // 商圈商户名称
	ShopName      string    `json:"shop_name,omitempty"`      // 门店名称
	ShopNumber    string    `json:"shop_number,omitempty"`    // 门店编号
	AppId         string    `json:"appid,omitempty"`          // 小程序APPID
	OpenId        string    `json:"openid,omitempty"`         // 用户标识
	RefundTime    time.Time `json:"refund_time,omitempty"`    // 退款完成时间
	PayAmount     int64     `json:"pay_amount,omitempty"`     // 消费金额
	RefundAmount  int64     `json:"refund_amount,omitempty"`  // 退款金额
	TransactionId string    `json:"transaction_id,omitempty"` // 微信支付订单号
	RefundId      string    `json:"refund_id,omitempty"`      // 微信支付退款单号
}

// PaymentResponse 商圈支付结果通知参数
type PaymentResponse struct {
	Id            string    `json:"-"`                        // 唯一通知ID
	MchId         string    `json:"mchid,omitempty"`          // 商户号
	MerchantName  string    `json:"merchant_name,omitempty"`  // 商圈商户名称
	ShopName      string    `json:"shop_name,omitempty"`      // 门店名称
	ShopNumber    string    `json:"shop_number,omitempty"`    // 门店编号
	AppId         string    `json:"appid,omitempty"`          // 小程序APPID
	OpenId        string    `json:"openid,omitempty"`         // 用户标识
	TimeEnd       time.Time `json:"time_end,omitempty"`       // 交易完成时间
	Amount        int64     `json:"amount,omitempty"`         // 金额
	TransactionId string    `json:"transaction_id,omitempty"` // 微信支付订单号
	CommitTag     string    `json:"commit_tag,omitempty"`     // 手动提交积分标记
}

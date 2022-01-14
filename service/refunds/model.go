package refunds

import (
	"time"

	"github.com/pyihe/wechat-sdk/model"
)

// RefundOrder 微信支付退款API应答
type RefundOrder struct {
	Id                  string                   `json:"-"`                               // id，请求或者通知的唯一ID
	MchId               string                   `json:"mch_id,omitempty"`                // 直连商户号
	SpMchId             string                   `json:"sp_mchid,omitempty"`              // 服务商户号
	SubMchId            string                   `json:"sub_mchid,omitempty"`             // 子商户号
	RefundId            string                   `json:"refund_id,omitempty"`             // 微信支付退款单号
	OutRefundNo         string                   `json:"out_refund_no,omitempty"`         // 商户退款单号
	TransactionId       string                   `json:"transaction_id,omitempty"`        // 微信支付订单号
	OutTradeNo          string                   `json:"out_trade_no,omitempty"`          // 商户订单号
	Channel             string                   `json:"channel,omitempty"`               // 退款渠道
	UserReceivedAccount string                   `json:"user_received_account,omitempty"` // 退款入账账户
	SuccessTime         time.Time                `json:"success_time,omitempty"`          // 退款成功时间
	CreateTime          time.Time                `json:"create_time,omitempty"`           // 退款创建时间
	Status              string                   `json:"status,omitempty"`                // 退款状态
	RefundStatus        string                   `json:"refund_status,omitempty"`         // 退款状态
	FundsAccount        string                   `json:"funds_account,omitempty"`         // 资金账户
	Amount              *model.Amount            `json:"amount,omitempty"`                // 金额信息
	PromotionDetail     []*model.PromotionDetail `json:"promotion_detail,omitempty"`      // 优惠退款信息
}

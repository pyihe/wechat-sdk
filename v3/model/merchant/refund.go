package merchant

import (
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/vars"
)

type RefundRequest struct {
	TransactionId string         `json:"transaction_id,omitempty"` // 微信支付订单
	OutTradeNo    string         `json:"out_trade_no,omitempty"`   // 商户订单号
	OutRefundNo   string         `json:"out_refund_no,omitempty"`  // 商户退款单号
	Reason        string         `json:"reason,omitempty"`         // 退款原因
	NotifyUrl     string         `json:"notify_url,omitempty"`     // 退款结果回调URL
	FundsAccount  string         `json:"funds_account,omitempty"`  // 退款资金来源
	Amount        *Amount        `json:"amount,omitempty"`         // 金额信息
	GoodsDetail   []*GoodsDetail `json:"goods_detail,omitempty"`   // 退款商品
}

func (refund *RefundRequest) Check() (err error) {
	if refund.TransactionId == "" && refund.OutTradeNo == "" {
		err = errors.New("退款时微信订单号和商户订单号不能同时为空!")
		return
	}
	if refund.OutRefundNo == "" {
		err = errors.New("请填写商户退款单号!")
		return
	}
	if refund.Amount == nil {
		err = errors.New("请填写金额信息!")
		return
	}
	if err = refund.Amount.Check(vars.TradeType(0)); err != nil {
		return
	}
	for _, goods := range refund.GoodsDetail {
		if err = goods.Check(); err != nil {
			return
		}
	}
	return
}

// RefundOrder 退款订单
type RefundOrder struct {
	Id                  string             `json:"-"`                               // 申请退款，查询退款时为Request-ID, 退款通知时为通知ID
	RefundId            string             `json:"refund_id,omitempty"`             // 微信支付退款单号
	OutRefundNo         string             `json:"out_refund_no,omitempty"`         // 商户退款单号
	TransactionId       string             `json:"transaction_id,omitempty"`        // 微信支付订单号
	OutTradeNo          string             `json:"out_trade_no,omitempty"`          // 商户订单号
	Channel             string             `json:"channel,omitempty"`               // 退款渠道
	UserReceivedAccount string             `json:"user_received_account,omitempty"` // 退款入账账户
	SuccessTime         string             `json:"success_time,omitempty"`          // 退款成功时间
	CreateTime          string             `json:"create_time,omitempty"`           // 退款受理时间
	Status              string             `json:"status,omitempty"`                // 退款状态
	FundsAccount        string             `json:"funds_account,omitempty"`         // 资金账户
	Amount              *Amount            `json:"amount,omitempty"`                // 金额信息
	PromotionDetail     []*PromotionDetail `json:"promotion_detail,omitempty"`      // 优惠退款信息
}

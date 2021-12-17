package combine

import (
	"github.com/pyihe/wechat-sdk/model"
	"github.com/pyihe/wechat-sdk/vars"

	"github.com/pyihe/go-pkg/errors"
)

type RefundRequest struct {
	TransactionId string               `json:"transaction_id,omitempty"` // 微信支付订单
	OutTradeNo    string               `json:"out_trade_no,omitempty"`   // 商户订单号
	OutRefundNo   string               `json:"out_refund_no,omitempty"`  // 商户退款单号
	Reason        string               `json:"reason,omitempty"`         // 退款原因
	NotifyUrl     string               `json:"notify_url,omitempty"`     // 退款结果回调URL
	FundsAccount  string               `json:"funds_account,omitempty"`  // 退款资金来源
	Amount        *model.Amount        `json:"amount,omitempty"`         // 金额信息
	GoodsDetail   []*model.GoodsDetail `json:"goods_detail,omitempty"`   // 退款商品
}

func (refund *RefundRequest) Check() (err error) {
	if refund == nil {
		err = vars.ErrNoRequest
		return
	}
	if refund.TransactionId == "" && refund.OutTradeNo == "" {
		err = errors.New("退款时transaction_id和out_trade_no不能同时为空!")
		return
	}
	if refund.OutRefundNo == "" {
		err = errors.New("请填写out_refund_no!")
		return
	}
	if refund.Amount == nil {
		err = errors.New("请填写amount!")
		return
	}
	if refund.Amount.Refund <= 0 {
		err = errors.New("请填写amount.refund!")
		return
	}
	if refund.Amount.Total <= 0 {
		err = errors.New("请填写amount.total!")
		return
	}
	if refund.Amount.Refund > refund.Amount.Total {
		err = errors.New("退款金额不能超过订单总金额!")
		return
	}
	if refund.Amount.Currency == "" {
		err = errors.New("请填写amount.currency!")
		return
	}
	for _, f := range refund.Amount.From {
		if f.Amount <= 0 {
			err = errors.New("请填写from.amount!")
			return
		}
		if f.Account == "" {
			err = errors.New("请填写from.account!")
			return
		}
	}
	for _, goods := range refund.GoodsDetail {
		if goods.MerchantGoodsId == "" {
			err = errors.New("请填写goods_detail.merchant_goods_id!")
			return
		}
		if goods.UnitPrice <= 0 {
			err = errors.New("请填写goods_detail.unit_price!")
			return
		}
		if goods.RefundAmount <= 0 {
			err = errors.New("请填写goods_detail.refund_amount!")
			return
		}
		if goods.RefundQuantity <= 0 {
			err = errors.New("请填写goods_detail.refund_quantity!")
			return
		}
	}
	return
}

package combine

import (
	"time"

	"github.com/pyihe/wechat-sdk/model"
)

// PrepayOrder 合单支付订单
type PrepayOrder struct {
	model.WechatError
	Id                string               `json:"-"`                              // 请求唯一标示或者通知唯一ID
	CombineAppId      string               `json:"combine_appid,omitempty"`        // 合单商户appid
	CombineMchId      string               `json:"combine_mchid,omitempty"`        // 合单商户号
	CombineOutTradeNo string               `json:"combine_out_trade_no,omitempty"` // 合单商户订单号
	SceneInfo         *model.SceneInfo     `json:"scene_info,omitempty"`           // 场景信息
	SubOrders         []*SubOrderResponse  `json:"sub_orders,omitempty"`           // 子单信息
	CombinePayerInfo  *model.MerchantPayer `json:"combine_payer_info,omitempty"`   // 支付者信息
}

// SubOrderResponse 子单信息
type SubOrderResponse struct {
	MchId           string                   `json:"mchid,omitempty"`            // 子单商户号
	SubMchId        string                   `json:"sub_mchid,omitempty"`        // 二级商户号
	TradeType       string                   `json:"trade_type,omitempty"`       // 交易类型
	TradeState      string                   `json:"trade_state,omitempty"`      // 交易状态
	BankType        string                   `json:"bank_type,omitempty"`        // 付款银行
	Attach          string                   `json:"attach,omitempty"`           // 附加数据
	SuccessTime     time.Time                `json:"success_time,omitempty"`     // 支付完成时间
	TransactionId   string                   `json:"transaction_id,omitempty"`   // 微信订单号
	OutTradeNo      string                   `json:"out_trade_no,omitempty"`     // 子单商户订单号
	PromotionDetail []*model.PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能
	Amount          *model.Amount            `json:"amount,omitempty"`           // 订单金额
}

type CloseOrderResponse struct {
	model.WechatError
	RequestId string `json:"request_id,omitempty"`
}

package partner

import (
	"time"

	"github.com/pyihe/wechat-sdk/model"
)

// QueryOrderRequest 查询订单request
type QueryOrderRequest struct {
	SpMchId       string `json:"sp_mchid,omitempty"`       // 服务商户号
	SubMchId      string `json:"sub_mch_id,omitempty"`     // 子商户号
	TransactionId string `json:"transaction_id,omitempty"` // 微信支付订单号
	OutTradeNo    string `json:"out_trade_no,omitempty"`   // 商户订单号
}

// PrepayOrder 订单
type PrepayOrder struct {
	model.WechatError
	Id              string                   `json:"-"`                          // 唯一请求ID或者唯一通知ID
	SpAppId         string                   `json:"sp_appid,omitempty"`         // 服务商应用ID
	SpMchId         string                   `json:"sp_mchid,omitempty"`         // 服务商户号
	SubAppId        string                   `json:"sub_appid,omitempty"`        // 子商户应用ID
	SubMchId        string                   `json:"sub_mchid,omitempty"`        // 子商户号
	OutTradeNo      string                   `json:"out_trade_no,omitempty"`     // 商户订单号
	TransactionId   string                   `json:"transaction_id,omitempty"`   // 微信支付订单号
	TradeType       string                   `json:"trade_type,omitempty"`       // 交易类型
	TradeState      string                   `json:"trade_state,omitempty"`      // 交易状态
	TradeStateDesc  string                   `json:"trade_state_desc,omitempty"` // 交易状态描述
	BankType        string                   `json:"bank_type,omitempty"`        // 付款银行
	Attach          string                   `json:"attach,omitempty"`           // 附加数据
	SuccessTime     time.Time                `json:"success_time,omitempty"`     // 支付完成时间
	Payer           *model.PartnerPayer      `json:"payer,omitempty"`            // 支付者信息
	Amount          *model.Amount            `json:"amount,omitempty"`           // 金额信息
	SceneInfo       *model.SceneInfo         `json:"scene_info,omitempty"`       // 场景信息
	PromotionDetail []*model.PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能
}

// CloseOrderRequest 关闭订单request
type CloseOrderRequest struct {
	SpMchId    string `json:"sp_mchid"`   // 服务商户号
	SubMchId   string `json:"sub_mch_id"` // 子商户号
	OutTradeNo string `json:"-"`          // 商户订单号
}

// CloseOrderResponse 关闭应答
type CloseOrderResponse struct {
	model.WechatError
	RequestId string `json:"request_id,omitempty"`
}

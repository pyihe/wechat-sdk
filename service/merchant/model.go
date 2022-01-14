package merchant

import (
	"time"

	"github.com/pyihe/wechat-sdk/model"
)

/****************************************************《基础支付》********************************************************/

// QueryOrderRequest 订单查询request
type QueryOrderRequest struct {
	TransactionId string `json:"transaction_id,omitempty"` // 微信订单号
	OutTradeNo    string `json:"out_trade_no,omitempty"`   // 商户订单号
}

// PrepayOrder 预支付订单
type PrepayOrder struct {
	model.WechatError
	Id              string                   `json:"-"`                          // 唯一标示，可能是请求的唯一ID，也可能是通知的唯一ID
	AppId           string                   `json:"appid,omitempty"`            // 直连商户申请的公众号或者移动应用的apppid
	MchId           string                   `json:"mchid,omitempty"`            // 直连商户的商户号
	OutTradeNo      string                   `json:"out_trade_no,omitempty"`     // 商户系统内部订单号
	TransactionId   string                   `json:"transaction_id,omitempty"`   // 微信支付订单号
	TradeType       string                   `json:"trade_type,omitempty"`       // 交易类型, JSAPI:公众号支付; NATIVE:扫码支付; APP:APP支付; MICROPAY:付款码支付; MWEB:H5支付; FACEPAY:刷脸支付
	TradeState      string                   `json:"trade_state,omitempty"`      // 交易状态, SUCCESS:支付成功; REFUND:转入退款; NOTPAY:未支付; CLOSED:已关闭; REVOKED:已撤销(仅付款码支付会返回); USERPAYING:用户支付中; PAYERROR:支付失败(仅付款码支付时会返回)
	TradeStateDesc  string                   `json:"trade_state_desc,omitempty"` // 交易状态描述
	BankType        string                   `json:"bank_type,omitempty"`        // 银行类型
	Attach          string                   `json:"attach,omitempty"`           // 附加数据
	SuccessTime     time.Time                `json:"success_time,omitempty"`     // 支付完成时间
	Payer           *model.MerchantPayer     `json:"payer,omitempty"`            // 支付者信息
	Amount          *model.Amount            `json:"amount,omitempty"`           // 订单金额
	SceneInfo       *model.SceneInfo         `json:"scene_info,omitempty"`       // 场景信息
	PromotionDetail []*model.PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能
}

type CloseOrderResponse struct {
	model.WechatError
	RequestId string `json:"request_id,omitempty"`
}

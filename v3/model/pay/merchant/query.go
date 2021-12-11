package merchant

import (
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/model"
)

type QueryRequest struct {
	TransactionId string `json:"transaction_id,omitempty"` // 微信支付订单号
	OutTradeNo    string `json:"out_trade_no,omitempty"`   // 商户订单号
}

func (q *QueryRequest) Check() (err error) {
	if q.TransactionId == "" && q.OutTradeNo == "" {
		err = errors.New("订单查询transaction_id和out_trade_no不能同时为空!")
		return
	}
	return
}

// PrepayOrder 预支付订单
type PrepayOrder struct {
	Id              string                   `json:"id,omitempty"`               // 唯一标示，可能是请求的唯一ID，也可能是通知的唯一ID
	AppId           string                   `json:"appid,omitempty"`            // 直连商户申请的公众号或者移动应用的apppid
	MchId           string                   `json:"mchid,omitempty"`            // 直连商户的商户号
	OutTradeNo      string                   `json:"out_trade_no,omitempty"`     // 商户系统内部订单号
	TransactionId   string                   `json:"transaction_id,omitempty"`   // 微信支付订单号
	TradeType       string                   `json:"trade_type,omitempty"`       // 交易类型, JSAPI:公众号支付; NATIVE:扫码支付; APP:APP支付; MICROPAY:付款码支付; MWEB:H5支付; FACEPAY:刷脸支付
	TradeState      string                   `json:"trade_state,omitempty"`      // 交易状态, SUCCESS:支付成功; REFUND:转入退款; NOTPAY:未支付; CLOSED:已关闭; REVOKED:已撤销(仅付款码支付会返回); USERPAYING:用户支付中; PAYERROR:支付失败(仅付款码支付时会返回)
	TradeStateDesc  string                   `json:"trade_state_desc,omitempty"` // 交易状态描述
	BankType        string                   `json:"bank_type,omitempty"`        // 银行类型
	Attach          string                   `json:"attach,omitempty"`           // 附加数据
	SuccessTime     string                   `json:"success_time,omitempty"`     // 支付完成时间
	Payer           *model.Payer             `json:"payer,omitempty"`            // 支付者信息
	Amount          *model.Amount            `json:"amount,omitempty"`           // 订单金额
	SceneInfo       *model.SceneInfo         `json:"scene_info,omitempty"`       // 场景信息
	PromotionDetail []*model.PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能
}

package merchant

type QueryRequest struct {
	TransactionId string `json:"transaction_id,omitempty"` // 微信支付订单号
	OutTradeNo    string `json:"out_trade_no,omitempty"`   // 商户订单号
}

type PrepayOrder struct {
	Id              string             `json:"id,omitempty"`               // 唯一标示，可能是请求的唯一ID，也可能是通知的唯一ID
	AppId           string             `json:"appid,omitempty"`            // 直连商户申请的公众号或者移动应用的apppid
	MchId           string             `json:"mchid,omitempty"`            // 直连商户的商户号
	OutTradeNo      string             `json:"out_trade_no,omitempty"`     // 商户系统内部订单号
	TransactionId   string             `json:"transaction_id,omitempty"`   // 微信支付订单号
	TradeType       string             `json:"trade_type,omitempty"`       // 交易类型, JSAPI:公众号支付; NATIVE:扫码支付; APP:APP支付; MICROPAY:付款码支付; MWEB:H5支付; FACEPAY:刷脸支付
	TradeState      string             `json:"trade_state,omitempty"`      // 交易状态, SUCCESS:支付成功; REFUND:转入退款; NOTPAY:未支付; CLOSED:已关闭; REVOKED:已撤销(仅付款码支付会返回); USERPAYING:用户支付中; PAYERROR:支付失败(仅付款码支付时会返回)
	TradeStateDesc  string             `json:"trade_state_desc,omitempty"` // 交易状态描述
	BankType        string             `json:"bank_type,omitempty"`        // 银行类型
	Attach          string             `json:"attach,omitempty"`           // 附加数据
	SuccessTime     string             `json:"success_time,omitempty"`     // 支付完成时间
	Payer           *Payer             `json:"payer,omitempty"`            // 支付者信息
	Amount          *Amount            `json:"amount,omitempty"`           // 订单金额
	SceneInfo       *SceneInfo         `json:"scene_info,omitempty"`       // 场景信息
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能
}

// PromotionDetail 优惠功能
type PromotionDetail struct {
	CouponId            string         `json:"coupon_id,omitempty"`            // 券ID
	Name                string         `json:"name,omitempty"`                 // 优惠名称
	Scope               string         `json:"scope,omitempty"`                // 优惠范围
	Type                string         `json:"type,omitempty"`                 // 优惠类型
	Amount              int64          `json:"amount,omitempty"`               // 优惠券面额
	StockId             string         `json:"stock_id,omitempty"`             // 活动ID
	WechatpayContribute int64          `json:"wechatpay_contribute,omitempty"` // 微信出资
	MerchantContribute  int64          `json:"merchant_contribute,omitempty"`  // 商户出资
	OtherContribute     int64          `json:"other_contribute,omitempty"`     // 其他出资
	Currency            string         `json:"currency,omitempty"`             // 优惠币种
	GoodsDetail         []*GoodsDetail `json:"goods_detail,omitempty"`         // 单品列表

	// 退款时返回
	PromotionId  string `json:"promotion_id,omitempty"`  // 券或者立减优惠ID
	RefundAmount int64  `json:"refund_amount,omitempty"` // 优惠退款金额

}

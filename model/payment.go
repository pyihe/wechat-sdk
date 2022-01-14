package model

// PromotionDetail 优惠功能
type PromotionDetail struct {
	PromotionId         string         `json:"promotion_id,omitempty"`         // 券或者立减优惠ID
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
	RefundAmount        int64          `json:"refund_amount,omitempty"`        // 优惠退款金额
	GoodsDetail         []*GoodsDetail `json:"goods_detail,omitempty"`         // 单品列表
}

// GoodsDetail 商品详情
type GoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id,omitempty"`  // 商户侧商品编码
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` // 微信支付商品编码
	GoodsName        string `json:"goods_name,omitempty"`         // 商品名称
	Quantity         int    `json:"quantity,omitempty"`           // 用户购买的数量
	UnitPrice        int64  `json:"unit_price,omitempty"`         // 商品单价
	GoodsId          string `json:"goods_id,omitempty"`           // 商品编码
	DiscountAmount   int64  `json:"discount_amount,omitempty"`    // 商品优惠金额
	GoodsRemark      string `json:"goods_remark,omitempty"`       // 商品备注
	RefundAmount     int64  `json:"refund_amount,omitempty"`      // 商品退款金额
	RefundQuantity   int    `json:"refund_quantity,omitempty"`    // 单品的退款数量
}

// From 出资账户及金额
type From struct {
	Account string `json:"account,omitempty"` // 出资账户类型
	Amount  int64  `json:"amount,omitempty"`  // 出资金额
}

// Amount 金额信息
type Amount struct {
	Total            int64   `json:"total,omitempty"`             // 金额
	Currency         string  `json:"currency,omitempty"`          // 货币类型
	Refund           int64   `json:"refund,omitempty"`            // 退款金额
	PayerTotal       int64   `json:"payer_total,omitempty"`       // 用户支付金额
	PayerCurrency    string  `json:"payer_currency,omitempty"`    // 用户支付币种
	PayerRefund      int64   `json:"payer_refund,omitempty"`      // 用户退款金额
	SettlementRefund int64   `json:"settlement_refund,omitempty"` // 应结退款金额
	SettlementTotal  int64   `json:"settlement_total,omitempty"`  // 应结订单金额
	DiscountRefund   int64   `json:"discount_refund,omitempty"`   // 优惠退款金额
	From             []*From `json:"from,omitempty"`              // 退款出资账户及金额

	// combine
	TotalAmount int64 `json:"total_amount,omitempty"` // 标价金额
	PayerAmount int64 `json:"payer_amount,omitempty"` // 现金支付金额

	// 停车服务
	DiscountTotal int64 `json:"discount_total,omitempty"` // 折扣
}

// MerchantPayer 商户平台支付者信息
type MerchantPayer struct {
	OpenId string `json:"openid,omitempty"` // 用户OpenId
}

// PartnerPayer 服务商平台支付者信息
type PartnerPayer struct {
	SpOpenId  string `json:"sp_openid,omitempty"`  // 用户服务标示
	SubOpenId string `json:"sub_openid,omitempty"` // 用户子标示
}

// SceneInfo 场景信息
type SceneInfo struct {
	DeviceId string `json:"device_id,omitempty"` // 商户端设备号
}

// AppResponse APP支付应答
type AppResponse struct {
	WechatError        // 微信Error
	RequestId   string `json:"request_id,omitempty"` // 唯一请求ID
	PrepayId    string `json:"prepay_id,omitempty"`  // 预支付交易会话标示
}

// H5Response H5支付应答
type H5Response struct {
	WechatError        // 微信Error
	RequestId   string `json:"request_id,omitempty"` // 唯一请求ID
	H5Url       string `json:"h5_url,omitempty"`     // 支付跳转链接
}

// JSAPIResponse JSAPI支付应答
type JSAPIResponse struct {
	WechatError        // 微信Error
	RequestId   string `json:"request_id,omitempty"` // 唯一请求ID
	PrepayId    string `json:"prepay_id,omitempty"`  // 预支付交易会话标示
}

// NativeResponse native支付应答
type NativeResponse struct {
	WechatError        // 微信Error
	RequestId   string `json:"request_id,omitempty"` // 唯一请求ID
	CodeUlr     string `json:"code_ulr,omitempty"`   // 二维码链接
}

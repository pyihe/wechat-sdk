package model

// WechatError 调用微信接口返回的通用错误格式
type WechatError struct {
	Code    string       `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
	Detail  *ErrorDetail `json:"detail,omitempty"`
}

// ErrorDetail 错误详情
type ErrorDetail struct {
	Field    string `json:"field,omitempty"`
	Value    string `json:"value,omitempty"`
	Issue    string `json:"issue,omitempty"`
	Location string `json:"location,omitempty"`
}

// Response 用于商户想微信服务端发送通知
type Response struct {
	Code    string `json:"code,omitempty"`    // 返回状态码
	Message string `json:"message,omitempty"` // 返回信息
}

// WechatNotifyResponse 微信通知的回复格式
type WechatNotifyResponse struct {
	Id           string            `json:"id,omitempty"`            // 通知的唯一ID
	CreateTime   string            `json:"create_time,omitempty"`   // 通知创建时间
	EventType    string            `json:"event_type,omitempty"`    // 通知类型
	ResourceType string            `json:"resource_type,omitempty"` // 通知的资源数据类型
	Summary      string            `json:"summary,omitempty"`       // 回调摘要
	Resource     *WechatCipherData `json:"resource,omitempty"`      // 通知资源数据
}

// WechatCipherData 微信返回的加密数据
type WechatCipherData struct {
	Algorithm      string `json:"algorithm,omitempty"`       // 加密算法
	CipherText     string `json:"cipher_text,omitempty"`     // 密文
	AssociatedData string `json:"associated_data,omitempty"` // 附加数据
	OriginalType   string `json:"original_type,omitempty"`   // 原始回调类型
	Nonce          string `json:"nonce,omitempty"`           // 加密使用的随机串
}

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

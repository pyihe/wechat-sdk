package model

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

// Payer 支付者信息
type Payer struct {
	OpenId string `json:"openid,omitempty"` // 用户OpenId
}

// From 退款出资账户及金额
type From struct {
	Account string `json:"account,omitempty"` // 出资账户类型
	Amount  int64  `json:"amount,omitempty"`  // 出资金额
}

// Amount 金额信息，根据每个API参数要求填写，如果填写了API没要求的参数可能导致请求失败
type Amount struct {
	Total            int64   `json:"total,omitempty"`             // 金额
	Currency         string  `json:"currency,omitempty"`          // 货币类型
	Refund           int64   `json:"refund,omitempty"`            // 退款金额
	From             []*From `json:"from,omitempty"`              // 退款出资账户及金额
	PayerTotal       int64   `json:"payer_total,omitempty"`       // 用户支付金额
	PayerCurrency    string  `json:"payer_currency,omitempty"`    // 用户支付币种
	PayerRefund      int64   `json:"payer_refund,omitempty"`      // 用户退款金额
	SettlementRefund int64   `json:"settlement_refund,omitempty"` // 应结退款金额
	SettlementTotal  int64   `json:"settlement_total,omitempty"`  // 应结订单金额
	DiscountRefund   int64   `json:"discount_refund,omitempty"`   // 优惠退款金额

	// combine
	TotalAmount int64 `json:"total_amount,omitempty"` // 标价金额
}

// Detail 优惠功能
type Detail struct {
	CostPrice   int64          `json:"cost_price,omitempty"`   // 订单原价
	InvoiceId   string         `json:"invoice_id,omitempty"`   // 商家小票ID
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"` // 单品列表
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

// SceneInfo 场景信息
type SceneInfo struct {
	PayerClientIp string     `json:"payer_client_ip,omitempty"` // 用户终端IP
	DeviceId      string     `json:"device_id,omitempty"`       // 商户端设备号
	StoreInfo     *StoreInfo `json:"store_info,omitempty"`      // 商户门店信息
	H5Info        *H5Info    `json:"h_5_info,omitempty"`        // h5场景信息
}

// StoreInfo 门店信息
type StoreInfo struct {
	Id       string `json:"id,omitempty"`        // 门店编号
	Name     string `json:"name,omitempty"`      // 商户侧门店名称
	AreaCode string `json:"area_code,omitempty"` // 地区编码
	Address  string `json:"address,omitempty"`   // 详细的商户门店地址
}

// SettleInfo 结算信息
type SettleInfo struct {
	ProfitSharing bool  `json:"profit_sharing,omitempty"` // 是否指定分账
	SubsidyAmount int64 `json:"subsidy_amount,omitempty"` // 补差金额
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

// H5Info H5场景信息
type H5Info struct {
	Type        string `json:"type,omitempty"`         // 场景类型
	AppName     string `json:"app_name,omitempty"`     // 应用名称
	AppUrl      string `json:"app_url,omitempty"`      // 网站URL
	BundleId    string `json:"bundle_id,omitempty"`    // IOS平台BundleID
	PackageName string `json:"package_name,omitempty"` // Android平台PackageName
}

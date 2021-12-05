package vars

/*************v3预支付应答结构******************/

// PrepayResponse 预支付API应答
type PrepayResponse struct {
	RequestId string       `json:"request_id"` // 请求唯一标示
	Code      string       `json:"code"`       // 详细错误码
	Message   string       `json:"message"`    // 错误描述
	Detail    *ErrorDetail `json:"detail"`     // 错误详细描述
	PrepayId  string       `json:"prepay_id"`  // 预支付ID, JSAPI、APP支付时返回
	CodeUrl   string       `json:"code_url"`   // 二维码链接, Native支付时返回
	H5Url     string       `json:"h5_url"`     // H5支付时返回
}

// ErrorDetail 错误详细描述
type ErrorDetail struct {
	Field    string `json:"field"`    // 指示错误参数的位置。当错误参数位于请求body的JSON时，填写指向参数的JSON Pointer 。当错误参数位于请求的url或者querystring时，填写参数的变量名。
	Value    string `json:"value"`    // 错误的值
	Issue    string `json:"issue"`    // 具体错误原因
	Location string `json:"location"` // 错误位于哪儿
}

/*************v3订单查询应答结构******************/

type OrderResponse struct {
	RequestId       string             `json:"request_id"`                 // 请求唯一标示
	AppId           string             `json:"appid,omitempty"`            // 应用ID
	MchId           string             `json:"mchid,omitempty"`            // 直连商户号
	OutTradeNo      string             `json:"out_trade_no,omitempty"`     // 商户系统内部订单号
	TransactionId   string             `json:"transaction_id,omitempty"`   // 微信支付系统生成的订单号
	TradeType       string             `json:"trade_type,omitempty"`       // 交易类型: JSAPI、NATIVE、APP、MICROPAY、MWEB、FACEPAY
	TradeState      string             `json:"trade_state,omitempty"`      // 交易状态
	TradeStateDesc  string             `json:"trade_state_desc,omitempty"` // 交易状态描述
	BankType        string             `json:"bank_type,omitempty"`        // 银行类型
	Attach          string             `json:"attach,omitempty"`           // 附加数据
	SuccessTime     string             `json:"success_time,omitempty"`     // 支付完成时间
	Payer           *Payer             `json:"payer,omitempty"`            // 支付者信息
	Amount          *Amount            `json:"amount,omitempty"`           // 订单金额信息
	SceneInfo       *SceneInfo         `json:"scene_info,omitempty"`       // 支付场景描述
	PromotionDetail []*PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能，享受优惠时返回

	// 退款
	OutRefundNo         string `json:"out_refund_no,omitempty"`         // 商户退款单号
	RefundId            string `json:"refund_id,omitempty"`             // 微信退款单号
	RefundStatus        string `json:"refund_status,omitempty"`         // 退款状态
	UserReceivedAccount string `json:"user_received_account,omitempty"` // 退款入账方
}

// PromotionDetail 优惠功能，享受优惠时返回
type PromotionDetail struct {
	PromotionId         string         `json:"promotion_id,omitempty"`         // 券ID
	CouponId            string         `json:"coupon_id,omitempty"`            // 券ID
	Name                string         `json:"name,omitempty"`                 // 优惠名称
	Scope               string         `json:"scope,omitempty"`                // 优惠范围，GLOBAL: 全场代金券；SINGLE：单品优惠
	Type                string         `json:"type,omitempty"`                 // 优惠类型，CASH: 充值；NOCASH：预充值
	Amount              int64          `json:"amount,omitempty"`               // 优惠券面额
	RefundAmount        int64          `json:"refund_amount,omitempty"`        // 优惠退款金额
	StockId             string         `json:"stock_id,omitempty"`             // 活动ID
	WechatpayContribute int64          `json:"wechatpay_contribute,omitempty"` // 微信出资
	MerchantContribute  int64          `json:"merchant_contribute,omitempty"`  // 商户出资
	OtherContribute     int64          `json:"other_contribute,omitempty"`     // 其他出资
	Currency            string         `json:"currency,omitempty"`             // 优惠币种
	GoodsDetail         []*GoodsDetail `json:"goods_detail,omitempty"`         // 单品列表信息
}

/*************v3预支付异步回调应答数据结构******************/

// Notify 接收回调通知
type Notify struct {
	RequestId    string    `json:"request_id,omitempty"`    // 唯一请求标示
	Id           string    `json:"id,omitempty"`            // 通知的唯一ID
	CreateTime   string    `json:"create_time,omitempty"`   // 通知创建时间
	EventType    string    `json:"event_type,omitempty"`    // 通知类型: 如TRANSACTION.SUCCESS
	ResourceType string    `json:"resource_type,omitempty"` // 通知数据类型: 支付成功为encrypt-resource
	Resource     *Resource `json:"resource,omitempty"`      // 通知资源数据
}

// Resource 接收数据资源
type Resource struct {
	Algorithm      string `json:"algorithm,omitempty"`       // 加密算法类型
	CipherText     string `json:"cipher_text,omitempty"`     // 数据密文
	AssociatedData string `json:"associated_data,omitempty"` // 附加数据
	OriginalType   string `json:"original_type,omitempty"`   // 原始回调类型: transaction
	Nonce          string `json:"nonce,omitempty"`           // 加密使用的随机字符串
	Summary        string `json:"summary,omitempty"`         // 回调摘要
}

/*************用于回复微信服务器******************/

// NotifyResponse 用于回复微信服务器
type NotifyResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

/*************退款回复******************/

// RefundResponse 退款回复
type RefundResponse struct {
	RequestId           string             `json:"request_id,omitempty"`            // 请求唯一标示
	RefundId            string             `json:"refund_id,omitempty"`             // 微信支付退款单号
	OutRefundNo         string             `json:"out_refund_no,omitempty"`         // 商户系统内部退款单号
	TransactionId       string             `json:"transaction_id,omitempty"`        // 微信支付订单号
	OutTradeNo          string             `json:"out_trade_no,omitempty"`          // 商户侧支付订单号
	Channel             string             `json:"channel,omitempty"`               // 退款渠道,ORIGINAL：原路退款 BALANCE：退回到余额 OTHER_BALANCE：原账户异常退到其他余额账户 OTHER_BANKCARD：原银行卡异常退到其他银行卡
	UserReceivedAccount string             `json:"user_received_account,omitempty"` // 退款入账账户
	SuccessTime         string             `json:"success_time,omitempty"`          // 退款成功时间
	CreateTime          string             `json:"create_time,omitempty"`           // 退款创建时间
	Status              string             `json:"status,omitempty"`                // 退款状态
	FundsAccount        string             `json:"funds_account,omitempty"`         // 资金账户
	Amount              *Amount            `json:"amount,omitempty"`                // 退款金额信息
	PromotionDetail     []*PromotionDetail `json:"promotion_detail"`                // 优惠退款信息
}

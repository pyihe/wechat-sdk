package vars

import "github.com/pyihe/go-pkg/errors"

/*************v3下单参数******************/

// PrepayRequest 预支付Request
type PrepayRequest struct {
	PayType     TradeType   `json:"-"`                      // 支付方式: JSAPI、APP、Native、H5
	AppId       string      `json:"appid,omitempty"`        // 由微信生成的应用ID，全局唯一
	MchId       string      `json:"mchid,omitempty"`        // 直连商户的商户号，由微信支付生成并下发
	Description string      `json:"description,omitempty"`  // 商品描述
	OutTradeNo  string      `json:"out_trade_no,omitempty"` // 商户系统内部订单号
	NotifyUrl   string      `json:"notify_url,omitempty"`   // 步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 公网域名必须为https，如果是走专线接入，使用专线NAT IP或者私有回调域名可使用http
	TimeExpire  string      `json:"time_expire,omitempty"`  // 订单失效时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	Attach      string      `json:"attach,omitempty"`       // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段
	GoodsTag    string      `json:"goods_tag,omitempty"`    // 订单优惠标记
	Amount      *Amount     `json:"amount,omitempty"`       // 订单金额信息
	Payer       *Payer      `json:"payer,omitempty"`        // 支付者信息
	Detail      *Detail     `json:"detail,omitempty"`       // 优惠功能
	SceneInfo   *SceneInfo  `json:"scene_info,omitempty"`   // 场景信息
	SettleInfo  *SettleInfo `json:"settle_info,omitempty"`  // 结算信息
}

func (pre *PrepayRequest) Check() (err error) {
	if pre.PayType.Valid() == false {
		err = errors.New("unknown pay type: " + string(pre.PayType) + ".")
		return
	}
	if pre.AppId == "" {
		err = errors.New("appid is necessary param.")
		return
	}
	if pre.MchId == "" {
		err = errors.New("mchid is necessary param.")
		return
	}
	if pre.Description == "" {
		err = errors.New("description is necessary param.")
		return
	}
	if pre.OutTradeNo == "" {
		err = errors.New("out_trade_no is necessary param.")
		return
	}
	if pre.NotifyUrl == "" {
		err = errors.New("notify_url is necessary param.")
		return
	}
	if pre.Amount == nil || pre.Amount.Total <= 0 {
		err = errors.New("amount is invalid.")
		return
	} else {
		if pre.Amount.PayerTotal > 0 || pre.Amount.PayerCurrency != "" {
			err = errors.New("payer_total&payer_currency must be empty when prepay.")
			return
		}
	}
	// JSAPI 支付必须填写支付者的ID
	if pre.PayType == JSAPI {
		if pre.Payer == nil || pre.Payer.OpenId == "" {
			err = errors.New("payer is invalid.")
			return
		}
	} else {
		if pre.Payer != nil {
			err = errors.New("payer must be empty when pay type is not JSAPI.")
			return
		}
	}
	if pre.Detail != nil && len(pre.Detail.GoodsDetail) > 0 {
		for _, goods := range pre.Detail.GoodsDetail {
			if goods.MerchantGoodsId == "" {
				err = errors.New("merchant_goods_id is necessary when detail is fill in.")
				return
			}
			if goods.Quantity <= 0 {
				err = errors.New("quantity is necessary when detail is fill in.")
				return
			}
			if goods.UnitPrice <= 0 {
				err = errors.New("unit_price is necessary when detail is fill in.")
				return
			}
		}
	}
	if pre.PayType == H5 && pre.SceneInfo == nil {
		err = errors.New("scene_info is necessary for h5 pay.")
		return
	}
	if pre.SceneInfo != nil {
		if pre.SceneInfo.PayerClientIp == "" {
			err = errors.New("payer_client_ip is necessary when scene_info is fill in.")
			return
		}
		if pre.SceneInfo.StoreInfo != nil {
			if pre.SceneInfo.StoreInfo.Id == "" {
				err = errors.New("id is necessary when store info is fill in.")
				return
			}
		}
		if pre.PayType == H5 {
			if pre.SceneInfo.H5Info == nil {
				err = errors.New("scene_info.h5_info is necessary for h5 pay.")
				return
			}
			if pre.SceneInfo.H5Info.Type == "" {
				err = errors.New("scene_info.h5_info.type is necessary for h5 pay.")
				return
			}
		} else {
			if pre.SceneInfo.H5Info != nil {
				err = errors.New("h5_info must be empty when pay type is not h5.")
				return
			}
		}
	}
	return
}

// Amount 订单金额信息
type Amount struct {
	// 下单时的参数
	Total    int64  `json:"total,omitempty"`    // 订单总金额，单位为分
	Currency string `json:"currency,omitempty"` // CNY：人民币，境内商户号仅支持人民币

	// 查询订单时的回复
	PayerTotal    int64  `json:"payer_total,omitempty"`    // 用户支付金额
	PayerCurrency string `json:"payer_currency,omitempty"` // 用户支付币种

	// 退款相关
	Refund           int64   `json:"refund,omitempty"`            // 退款金额
	PayerRefund      int64   `json:"payer_refund,omitempty"`      // 退款给用户的金额
	From             []*From `json:"from,omitempty"`              // 退款需要从指定账户出资时，传递此参数指定出资金额（币种的最小单位，只能为整数）
	SettlementRefund int64   `json:"settlement_refund,omitempty"` // 应结退款金额
	SettlementTotal  int64   `json:"settlement_total,omitempty"`  // 应结订单金额
	DiscountRefund   int64   `json:"discount_refund,omitempty"`   // 优惠退款金额
}

type From struct {
	Account string `json:"account,omitempty"` // 出资账户类型
	Amount  int64  `json:"amount,omitempty"`  // 对应账户出资金额
}

// Payer 支付者信息
type Payer struct {
	OpenId string `json:"openid,omitempty"` // 用户在直连商户appid下的唯一标识
}

// Detail 优惠功能
type Detail struct {
	CostPrice   int64          `json:"cost_price,omitempty"`   // 订单原价,
	InvoiceId   string         `json:"invoice_id,omitempty"`   // 商家小票ID
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"` // 单品列表信息, 条目个数限制：【1，6000】
}

// GoodsDetail 商品信息
type GoodsDetail struct {
	// 下单时的参数
	MerchantGoodsId  string `json:"merchant_goods_id,omitempty"`  // 商户侧商品编码, 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` // 微信支付商品编码, 微信支付定义的统一商品编号（没有可不传）
	GoodsName        string `json:"goods_name,omitempty"`         // 商品名称, 商品的实际名称

	// 查询订单时返回的参数
	GoodsId        string `json:"goods_id,omitempty"`        // 商品编码
	DiscountAmount int64  `json:"discount_amount,omitempty"` // 商品优惠金额
	GoodsRemark    string `json:"goods_remark,omitempty"`    // 商品备注信息

	// 退款时的参数
	RefundAmount   int64 `json:"refund_amount,omitempty"`   // 商品退款金额
	RefundQuantity int64 `json:"refund_quantity,omitempty"` // 单品的退款数量

	// 公共部分
	Quantity  int64 `json:"quantity,omitempty"`   // 商品数量, 用户购买的数量
	UnitPrice int64 `json:"unit_price,omitempty"` // 商品单价, 商品单价，单位为分
}

// SceneInfo 场景信息
type SceneInfo struct {
	PayerClientIp string     `json:"payer_client_ip,omitempty"` // 用户终端IP, 用户的客户端IP，支持IPv4和IPv6两种格式的IP地址
	DeviceId      string     `json:"device_id,omitempty"`       // 商户端设备号, 商户端设备号（门店号或收银设备ID）
	StoreInfo     *StoreInfo `json:"store_info,omitempty"`      // 商户门店信息
	H5Info        *H5Info    `json:"h5_info,omitempty"`         // H5场景信息
}

// StoreInfo 门店信息
type StoreInfo struct {
	Id       string `json:"id,omitempty"`        // 门店编号, 商户侧门店编号
	Name     string `json:"name,omitempty"`      // 门店名称, 商户侧门店名称
	AreaCode string `json:"area_code,omitempty"` // 地区编码, 地区编码，详细请见省市区编号对照表
	Address  string `json:"address,omitempty"`   // 详细地址, 详细的商户门店地址
}

// H5Info H5场景信息
type H5Info struct {
	Type        string `json:"type,omitempty"`         // 场景类型: IOS, Android, Wap
	AppName     string `json:"app_name,omitempty"`     // 应用名称
	AppUrl      string `json:"app_url,omitempty"`      // 网站URL
	BundleId    string `json:"bundle_id,omitempty"`    // IOS平台的BundleID
	PackageName string `json:"package_name,omitempty"` // Android平台的PackageName
}

// SettleInfo 结算信息
type SettleInfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"` // 是否指定分账
}

/*************v3退款参数******************/

// RefundRequest 退款
type RefundRequest struct {
	TransactionId string         `json:"transaction_id,omitempty"` // 原始支付交易对应的微信订单号
	OutTradeNo    string         `json:"out_trade_no,omitempty"`   // 原始支付交易对应的商户订单号
	OutRefundNo   string         `json:"out_refund_no,omitempty"`  // 商户系统内部的退款单号
	Reason        string         `json:"reason,omitempty"`         // 退款原因，如果商户传入，会在下发给用户的退款消息中体现退款原因
	NotifyUrl     string         `json:"notify_url,omitempty"`     // 异步接收退款通知的回调地址, 为空时微信将会回调至商户平台上配置的回调地址
	FundsAccount  string         `json:"funds_account,omitempty"`  // 若传递次参数则使用对应的资金账户退款，否则默认使用未结算资金退款
	Amount        *Amount        `json:"amount,omitempty"`         // 订单金额信息
	GoodsDetail   []*GoodsDetail `json:"goods_detail,omitempty"`   // 退款商品
}

func (refund *RefundRequest) Check() (err error) {
	if refund.TransactionId == "" && refund.OutRefundNo == "" {
		err = errors.New("please fill in transaction_id or out_trade_no.")
		return
	}
	if refund.OutRefundNo == "" {
		err = errors.New("out_refund_no is necessary param.")
		return
	}
	if refund.Amount == nil {
		err = errors.New("amount is necessary param.")
		return
	}
	if refund.Amount.Refund <= 0 {
		err = errors.New("refund is necessary param.")
		return
	}
	if refund.Amount.Total <= 0 {
		err = errors.New("total is necessary param.")
		return
	}
	if refund.Amount.Currency == "" {
		err = errors.New("currency is necessary param.")
		return
	}
	return
}

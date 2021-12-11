package merchant

import (
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/vars"
)

// PrepayRequest 预支付request
type PrepayRequest struct {
	TradeType   vars.TradeType `json:"-"`                      // 支付方式: JSAPI、APP、Native、H5、FacePay
	AppId       string         `json:"appid,omitempty"`        // 由微信生成的应用ID，全局唯一
	MchId       string         `json:"mchid,omitempty"`        // 直连商户的商户号，由微信支付生成并下发
	Description string         `json:"description,omitempty"`  // 商品描述
	OutTradeNo  string         `json:"out_trade_no,omitempty"` // 商户系统内部订单号
	TimeExpire  string         `json:"time_expire,omitempty"`  // 订单失效时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	Attach      string         `json:"attach,omitempty"`       // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段
	NotifyUrl   string         `json:"notify_url,omitempty"`   // 步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 公网域名必须为https，如果是走专线接入，使用专线NAT IP或者私有回调域名可使用http
	GoodsTag    string         `json:"goods_tag,omitempty"`    // 订单优惠标记
	Amount      *Amount        `json:"amount,omitempty"`       // 订单金额信息
	Payer       *Payer         `json:"payer,omitempty"`        // 支付者信息
	Detail      *Detail        `json:"detail,omitempty"`       // 优惠功能
	SceneInfo   *SceneInfo     `json:"scene_info,omitempty"`   // 场景信息
	SettleInfo  *SettleInfo    `json:"settle_info,omitempty"`  // 结算信息
}

func (pre *PrepayRequest) Check() (err error) {
	if pre.TradeType.Valid() == false {
		err = errors.New("请填写正确的交易类型!")
		return
	}
	if pre.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if pre.MchId == "" {
		err = errors.New("请填写mchid!")
		return
	}
	if pre.Description == "" {
		err = errors.New("请填写description!")
		return
	}
	if pre.OutTradeNo == "" {
		err = errors.New("请填写out_trade_no!")
		return
	}
	if pre.NotifyUrl == "" {
		err = errors.New("请填写notify_url!")
		return
	}
	if pre.Amount == nil {
		err = errors.New("请填写amount!")
		return
	}
	if err = pre.Amount.Check(pre.TradeType); err != nil {
		return
	}
	if pre.TradeType == vars.JSAPI && pre.Payer == nil {
		err = errors.New("JSAPI下单请填写payer.openid!")
		return
	}
	if pre.Payer != nil {
		if err = pre.Payer.Check(); err != nil {
			return
		}
	}
	if pre.Detail != nil {
		if err = pre.Detail.Check(); err != nil {
			return
		}
	}
	if pre.SceneInfo != nil {
		if err = pre.SceneInfo.Check(); err != nil {
			return
		}
	}
	return
}

// Amount 金额信息
type Amount struct {
	Total    int64  `json:"total,omitempty"`    // 订单总金额
	Currency string `json:"currency,omitempty"` // 货币类型

	// 订单查询时返回
	PayerTotal    int64  `json:"payer_total,omitempty"`    // 用户支付金额
	PayerCurrency string `json:"payer_currency,omitempty"` // 用户支付币种

	// 退款时传递
	Refund int64   `json:"refund,omitempty"` // 退款金额
	From   []*From `json:"from,omitempty"`   // 退款出资账户及金额

	// 退款时返回
	PayerRefund      int64 `json:"payer_refund,omitempty"`      // 用户退款金额
	SettlementRefund int64 `json:"settlement_refund,omitempty"` // 应结退款金额
	SettlementTotal  int64 `json:"settlement_total,omitempty"`  // 应结订单金额
	DiscountRefund   int64 `json:"discount_refund,omitempty"`   // 优惠退款金额
}

func (a *Amount) Check(tradeType vars.TradeType) (err error) {
	switch tradeType {
	case vars.H5: // H5支付
		fallthrough
	case vars.JSAPI: // JSAPI支付
		fallthrough
	case vars.APP: // APP支付
		fallthrough
	case vars.Native: // Native支付
		fallthrough
	case vars.FacePay: // 刷脸支付（商户平台不支持）
		if a.Total <= 0 {
			err = errors.New("请填写正确的amount.total!")
			return
		}
	default: // 退款
		if a.Total <= 0 {
			err = errors.New("请填写正确的amount.total!")
			return
		}
		if a.Currency == "" {
			err = errors.New("请填写正确的amount.currency!")
			return
		}
		if a.Refund <= 0 {
			err = errors.New("请填写amount.refund!")
			return
		}
		if a.Refund > a.Total {
			err = errors.New("退款金额不能超过原订单支付金额!")
			return
		}
		for _, f := range a.From {
			if err = f.Check(); err != nil {
				return
			}
		}
	}
	return
}

// From 退款出资账户及金额
type From struct {
	Account string `json:"account,omitempty"` // 出资账户类型
	Amount  int64  `json:"amount,omitempty"`  // 出资金额
}

func (f *From) Check() (err error) {
	if f.Account == "" {
		err = errors.New("请填写from.account!")
		return
	}
	if f.Amount <= 0 {
		err = errors.New("请填写from.amount!")
		return
	}
	return
}

// Payer 支付者信息
type Payer struct {
	OpenId string `json:"open_id,omitempty"` // 支付者openid
}

func (p *Payer) Check() (err error) {
	if p.OpenId == "" {
		err = errors.New("请填写payer.openid!")
		return
	}
	return
}

// Detail 优惠功能
type Detail struct {
	CostPrice   int64          `json:"cost_price,omitempty"`   // 订单原价
	InvoiceId   string         `json:"invoice_id,omitempty"`   // 商家小票ID
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"` // 商品列表
}

func (d *Detail) Check() (err error) {
	for _, goods := range d.GoodsDetail {
		if err = goods.Check(); err != nil {
			return
		}
	}
	return
}

// GoodsDetail 单品列表
type GoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id,omitempty"`  // 商户侧商品编码
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` // 微信支付商品编码
	GoodsName        string `json:"goods_name,omitempty"`         // 商品的实际名称
	Quantity         int    `json:"quantity,omitempty"`           // 商品数量
	UnitPrice        int64  `json:"unit_price,omitempty"`         // 商品单价

	// 订单查询时返回
	GoodsId        string `json:"goods_id,omitempty"`        // 商品编码
	DiscountAmount int64  `json:"discount_amount,omitempty"` // 商品优惠金额
	GoodsRemark    string `json:"goods_remark,omitempty"`    // 商品备注信息

	// 申请退款时传递和返回
	RefundAmount   int64 `json:"refund_amount,omitempty"`   // 商品退款金额
	RefundQuantity int   `json:"refund_quantity,omitempty"` // 商品退款数量
}

func (goods *GoodsDetail) Check() (err error) {
	if goods.MerchantGoodsId == "" {
		err = errors.New("请填写goods_detail.merchant_goods_id!")
		return
	}
	if goods.Quantity <= 0 {
		err = errors.New("请填写正确的goods_detail.quantity!")
		return
	}
	if goods.UnitPrice <= 0 {
		err = errors.New("请填写正确的goods_detail.unit_price!")
		return
	}
	return
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	PayerClientIp string     `json:"payer_client_ip,omitempty"` // 用户客户端的IP
	DeviceId      string     `json:"device_id,omitempty"`       // 商户端设备号
	StoreInfo     *StoreInfo `json:"store_info,omitempty"`      // 商户门店信息
	H5Info        *H5Info    `json:"h5_info,omitempty"`         // H5场景
}

func (s *SceneInfo) Check() (err error) {
	if s.PayerClientIp == "" {
		err = errors.New("请填写scene_info.payer_client_ip!")
		return
	}
	if s.StoreInfo != nil {
		if err = s.StoreInfo.Check(); err != nil {
			return
		}
	}
	if s.H5Info != nil {
		if err = s.H5Info.Check(); err != nil {
			return
		}
	}
	return
}

// StoreInfo 商户门店信息
type StoreInfo struct {
	Id       string `json:"id,omitempty"`        // 商户侧门店编号
	Name     string `json:"name,omitempty"`      // 商户侧门店名称
	AreaCode string `json:"area_code,omitempty"` // 地区编码
	Address  string `json:"address,omitempty"`   // 详细的商户门店地址
}

func (s *StoreInfo) Check() (err error) {
	if s.Id == "" {
		err = errors.New("请填写store_info.id!")
		return
	}
	return
}

// H5Info H5场景信息
type H5Info struct {
	Type        string `json:"type,omitempty"`         // 场景类型
	AppName     string `json:"app_name,omitempty"`     // 应用名称
	AppUrl      string `json:"app_url,omitempty"`      // 网站URL
	BundleId    string `json:"bundle_id,omitempty"`    // IOS平台的BundleID
	PackageName string `json:"package_name,omitempty"` // Android平台的PackageName
}

func (h *H5Info) Check() (err error) {
	if h.Type == "" {
		err = errors.New("请填写h5_info.type!")
		return
	}
	return
}

// SettleInfo 结算信息
type SettleInfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"` // 是否指定分账
}

// PrepayResponse 预支付回复
type PrepayResponse struct {
	RequestId string `json:"request_id"`
	PrepayId  string `json:"prepay_id"`
	CodeUrl   string `json:"code_url"`
	H5Url     string `json:"h5_url"`
}

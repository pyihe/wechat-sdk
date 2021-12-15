package pay

import "github.com/pyihe/wechat-sdk/model"

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
	PayerAmount int64 `json:"payer_amount,omitempty"` // 现金支付金额
}

// Detail 优惠功能
type Detail struct {
	CostPrice   int64                `json:"cost_price,omitempty"`   // 订单原价
	InvoiceId   string               `json:"invoice_id,omitempty"`   // 商家小票ID
	GoodsDetail []*model.GoodsDetail `json:"goods_detail,omitempty"` // 单品列表
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
	ProfitSharing bool  `json:"profit_sharing"`           // 是否指定分账
	SubsidyAmount int64 `json:"subsidy_amount,omitempty"` // 补差金额
}

// H5Info H5场景信息
type H5Info struct {
	Type        string `json:"type,omitempty"`         // 场景类型
	AppName     string `json:"app_name,omitempty"`     // 应用名称
	AppUrl      string `json:"app_url,omitempty"`      // 网站URL
	BundleId    string `json:"bundle_id,omitempty"`    // IOS平台BundleID
	PackageName string `json:"package_name,omitempty"` // Android平台PackageName
}

package v3

import (
	"reflect"
	"strconv"
)

type Param map[string]interface{}

func NewParam() Param {
	return make(Param)
}

func (p Param) Add(key string, value interface{}) {
	p[key] = value
}

func (p Param) Delete(key string) (ok bool) {
	_, ok = p[key]
	delete(p, key)
	return
}

func (p Param) Get(key string) (value interface{}, ok bool) {
	value, ok = p[key]
	return
}

func (p Param) Range(fn func(key string, value interface{})) {
	for k, v := range p {
		fn(k, v)
	}
}

func (p Param) Check(fn func(key string, value interface{}) bool) (ok bool) {
	ok = true
	for k, v := range p {
		if fn(k, v) == false {
			ok = false
			break
		}
	}
	return
}

func (p Param) GetString(key string) (string, error) {
	value, ok := p.Get(key)
	if !ok {
		return "", ErrNotExistKey
	}
	return reflect.ValueOf(value).String(), nil
}

func (p Param) GetInt64(key string) (n int64, err error) {
	value, ok := p.Get(key)
	if !ok {
		return 0, ErrNotExistKey
	}
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	switch t.Kind() {
	case reflect.Bool:
		if v.Bool() == true {
			n = 1
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n = v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n = int64(v.Uint())
	case reflect.String:
		n, err = strconv.ParseInt(v.String(), 10, 64)
	case reflect.Float32, reflect.Float64:
		n = int64(v.Float())
	default:
		err = ErrConvert
	}
	return
}

/*************v3下单参数******************/

// PrepayRequest 预支付Request
type PrepayRequest struct {
	AppId       string      `json:"appid,omitempty"`        // 由微信生成的应用ID，全局唯一
	MchId       string      `json:"mchid,omitempty"`        // 直连商户的商户号，由微信支付生成并下发
	Description string      `json:"description,omitempty"`  // 商品描述
	OutTradeNo  string      `json:"out_trade_no,omitempty"` // 商户系统内部订单号
	TimeExpire  string      `json:"time_expire,omitempty"`  // 订单失效时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	Attach      string      `json:"attach,omitempty"`       // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段
	NotifyUrl   string      `json:"notify_url,omitempty"`   // 步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 公网域名必须为https，如果是走专线接入，使用专线NAT IP或者私有回调域名可使用http
	GoodsTag    string      `json:"goods_tag,omitempty"`    // 订单优惠标记
	Amount      *Amount     `json:"amount,omitempty"`       // 订单金额信息
	Payer       *Payer      `json:"payer,omitempty"`        // 支付者信息
	Detail      *Detail     `json:"detail,omitempty"`       // 优惠功能
	SceneInfo   *SceneInfo  `json:"scene_info,omitempty"`   // 场景信息
	SettleInfo  *SettleInfo `json:"settle_info,omitempty"`  // 结算信息
}

// Amount 订单金额信息
type Amount struct {
	Total    int64  `json:"total,omitempty"`    // 订单总金额，单位为分
	Currency string `json:"currency,omitempty"` // CNY：人民币，境内商户号仅支持人民币
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
	MerchantGoodsId  string `json:"merchant_goods_id,omitempty"`  // 商户侧商品编码, 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` // 微信支付商品编码, 微信支付定义的统一商品编号（没有可不传）
	GoodsName        string `json:"goods_name,omitempty"`         // 商品名称, 商品的实际名称
	Quantity         int64  `json:"quantity,omitempty"`           // 商品数量, 用户购买的数量
	UnitPrice        int64  `json:"unit_price,omitempty"`         // 商品单价, 商品单价，单位为分
}

// SceneInfo 场景信息
type SceneInfo struct {
	PayerClientIp string     `json:"payer_client_ip,omitempty"` // 用户终端IP, 用户的客户端IP，支持IPv4和IPv6两种格式的IP地址
	DeviceId      string     `json:"device_id,omitempty"`       // 商户端设备号, 商户端设备号（门店号或收银设备ID）
	StoreInfo     *StoreInfo `json:"store_info,omitempty"`      // 商户门店信息
}

// StoreInfo 门店信息
type StoreInfo struct {
	Id       string `json:"id,omitempty"`        // 门店编号, 商户侧门店编号
	Name     string `json:"name,omitempty"`      // 门店名称, 商户侧门店名称
	AreaCode string `json:"area_code,omitempty"` // 地区编码, 地区编码，详细请见省市区编号对照表
	Address  string `json:"address,omitempty"`   // 详细地址, 详细的商户门店地址
}

// SettleInfo 结算信息
type SettleInfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"` // 是否指定分账
}

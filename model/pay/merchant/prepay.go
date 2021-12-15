package merchant

import (
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/model"
	"github.com/pyihe/wechat-sdk/vars"
)

// PrepayRequest 预支付request
type PrepayRequest struct {
	TradeType   vars.TradeType    `json:"-"`                      // 支付方式: JSAPI、APP、Native、H5、FacePay
	AppId       string            `json:"appid,omitempty"`        // 由微信生成的应用ID，全局唯一
	MchId       string            `json:"mchid,omitempty"`        // 直连商户的商户号，由微信支付生成并下发
	Description string            `json:"description,omitempty"`  // 商品描述
	OutTradeNo  string            `json:"out_trade_no,omitempty"` // 商户系统内部订单号
	TimeExpire  string            `json:"time_expire,omitempty"`  // 订单失效时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒
	Attach      string            `json:"attach,omitempty"`       // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段
	NotifyUrl   string            `json:"notify_url,omitempty"`   // 步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 公网域名必须为https，如果是走专线接入，使用专线NAT IP或者私有回调域名可使用http
	GoodsTag    string            `json:"goods_tag,omitempty"`    // 订单优惠标记
	Amount      *model.Amount     `json:"amount,omitempty"`       // 订单金额信息
	Payer       *model.Payer      `json:"payer,omitempty"`        // 支付者信息
	Detail      *model.Detail     `json:"detail,omitempty"`       // 优惠功能
	SceneInfo   *model.SceneInfo  `json:"scene_info,omitempty"`   // 场景信息
	SettleInfo  *model.SettleInfo `json:"settle_info,omitempty"`  // 结算信息
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
	if pre.Amount.Total <= 0 {
		err = errors.New("请填写amount.total!")
		return
	}
	if pre.TradeType == vars.JSAPI && pre.Payer == nil {
		err = errors.New("JSAPI下单请填写payer.openid!")
		return
	}
	if pre.Payer != nil {
		if pre.Payer.OpenId == "" {
			err = errors.New("请填写payer.openid!")
			return
		}
	}
	if pre.Detail != nil {
		for _, goods := range pre.Detail.GoodsDetail {
			if goods.MerchantGoodsId == "" {
				err = errors.New("请填写goods_detail.merchant_goods_id!")
				return
			}
			if goods.Quantity <= 0 {
				err = errors.New("请填写goods_detail.quantity!")
				return
			}
			if goods.UnitPrice <= 0 {
				err = errors.New("请填写goods_detail.unit_price!")
				return
			}
		}
	}
	if pre.TradeType == vars.H5 && pre.SceneInfo == nil {
		err = errors.New("H5支付必须填写scene_info!")
		return
	}
	if pre.SceneInfo != nil {
		if pre.SceneInfo.PayerClientIp == "" {
			err = errors.New("请填写scene_info.payer_client_ip!")
			return
		}
		if pre.SceneInfo.StoreInfo != nil {
			if pre.SceneInfo.StoreInfo.Id == "" {
				err = errors.New("请填写store_info.id!")
				return
			}
		}
		if pre.TradeType == vars.H5 && pre.SceneInfo.H5Info == nil {
			err = errors.New("H5支付时必须填写scene_info.h5_info!")
			return
		}
		if pre.SceneInfo.H5Info != nil {
			if pre.SceneInfo.H5Info.Type == "" {
				err = errors.New("请填写h5_info.type!")
				return
			}
		}
	}
	return
}

// PrepayResponse 预支付回复
type PrepayResponse struct {
	model.WechatError
	RequestId string `json:"request_id"`
	PrepayId  string `json:"prepay_id"`
	CodeUrl   string `json:"code_url"`
	H5Url     string `json:"h5_url"`
}

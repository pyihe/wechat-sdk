package combine

import (
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/model"
	"github.com/pyihe/wechat-sdk/v3/vars"
)

type PrepayRequest struct {
	TradeType         vars.TradeType   `json:"-"`                              // 交易类型
	CombineAppId      string           `json:"combine_appid,omitempty"`        // 合单发起方的appid
	CombineMchId      string           `json:"combine_mchid,omitempty"`        // 合单发起方商户号
	CombineOutTradeNo string           `json:"combine_out_trade_no,omitempty"` // 合单支付总订单号
	SceneInfo         *model.SceneInfo `json:"scene_info,omitempty"`           // 支付场景描述
	SubOrders         []*SubOrder      `json:"sub_orders,omitempty"`           // 子单信息
	CombinePayerInfo  *model.Payer     `json:"combine_payer_info,omitempty"`   // 支付者信息
	TimeStart         string           `json:"time_start,omitempty"`           // 交易起始时间
	TimeExpire        string           `json:"time_expire,omitempty"`          // 交易结束时间
	NotifyUrl         string           `json:"notify_url,omitempty"`           // 通知地址
}

func (c *PrepayRequest) Check() (err error) {
	if c.TradeType.Valid() == false {
		err = errors.New("请填写正确的交易类型!")
		return
	}
	if c.CombineAppId == "" {
		err = errors.New("请填写combine_appid!")
		return
	}
	if c.CombineMchId == "" {
		err = errors.New("请填写combine_mchid!")
		return
	}
	if c.CombineOutTradeNo == "" {
		err = errors.New("请填写combine_out_trade_no!")
		return
	}
	if c.NotifyUrl == "" {
		err = errors.New("请填写notify_url!")
		return
	}
	if c.TradeType == vars.H5 && c.SceneInfo == nil {
		err = errors.New("H5支付必须填写scene_info!")
		return
	}
	if c.SceneInfo != nil {
		if c.SceneInfo.PayerClientIp == "" {
			err = errors.New("请填写scene_info.payer_client_ip!")
			return
		}
		if c.TradeType == vars.H5 {
			if c.SceneInfo.DeviceId == "" {
				err = errors.New("H5支付时必须填写scene_info.device_id!")
				return
			}
			if c.SceneInfo.H5Info == nil {
				err = errors.New("H5支付时必须填写scene_info.h5_info!")
				return
			}
			if c.SceneInfo.H5Info.Type == "" {
				err = errors.New("请填写h5_info.type!")
				return
			}
		}
	}
	if len(c.SubOrders) == 0 {
		err = errors.New("请填写sub_orders!")
		return
	}
	for _, order := range c.SubOrders {
		if order.MchId == "" {
			err = errors.New("请填写sub_orders.mchid!")
			return
		}
		if order.Attach == "" {
			err = errors.New("请填写sub_orders.attach!")
			return
		}
		if order.Amount == nil {
			err = errors.New("请填写sub_orders.amount!")
			return
		}
		if order.Amount.TotalAmount <= 0 {
			err = errors.New("请填写amount.total_amount!")
			return
		}
		if order.Amount.Currency == "" {
			err = errors.New("请填写amount.currency!")
			return
		}
		if order.OutTradeNo == "" {
			err = errors.New("请填写sub_orders.out_trade_no!")
			return
		}
		if order.Description == "" {
			err = errors.New("请填写sub_orders.description!")
			return
		}
	}
	if c.TradeType == vars.JSAPI && (c.CombinePayerInfo == nil || c.CombinePayerInfo.OpenId == "") {
		err = errors.New("请填写combine_payer_info!")
		return
	}
	return
}

// SubOrder 子单信息
type SubOrder struct {
	MchId           string                   `json:"mchid,omitempty"`            // 子单商户号
	Attach          string                   `json:"attach,omitempty"`           // 附加数据
	Amount          *model.Amount            `json:"amount,omitempty"`           // 金额信息
	OutTradeNo      string                   `json:"out_trade_no,omitempty"`     // 子单商户订单号
	GoodsTag        string                   `json:"goods_tag,omitempty"`        // 订单优惠标记
	Description     string                   `json:"description,omitempty"`      // 商品描述
	SettleInfo      *model.SettleInfo        `json:"settle_info,omitempty"`      // 结算信息
	TradeType       string                   `json:"trade_type,omitempty"`       // 交易类型
	TradeState      string                   `json:"trade_state,omitempty"`      // 交易i状态
	BankType        string                   `json:"bank_type,omitempty"`        // 付款银行
	SuccessTime     string                   `json:"success_time,omitempty"`     // 订单支付时间
	TransactionId   string                   `json:"transaction_id,omitempty"`   // 微信订单号
	PromotionDetail []*model.PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能
}

type PrepayResponse struct {
	RequestId string `json:"request_id,omitempty"`
	PrepayId  string `json:"prepay_id,omitempty"`
	CodeUrl   string `json:"code_url,omitempty"`
	H5Url     string `json:"h5_url,omitempty"`
}

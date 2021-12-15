package combine

import (
	"github.com/pyihe/wechat-sdk/model"
	"github.com/pyihe/wechat-sdk/model/pay"
)

type PrepayOrder struct {
	model.WechatError
	Id                string         `json:"-"`                              // 请求唯一ID或者通知的唯一ID
	CombineAppId      string         `json:"combine_appid,omitempty"`        // 合单商户appid
	CombineMchId      string         `json:"combine_mchid,omitempty"`        // 合单商户号
	CombineOutTradeNo string         `json:"combine_out_trade_no,omitempty"` // 合单商户订单号
	SceneInfo         *pay.SceneInfo `json:"scene_info,omitempty"`           // 场景信息
	SubOrders         []*SubOrder    `json:"sub_orders,omitempty"`           // 子单信息
	CombinePayerInfo  *pay.Payer     `json:"combine_payer_info,omitempty"`   // 支付者
}

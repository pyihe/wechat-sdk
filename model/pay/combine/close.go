package combine

import (
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/model"
)

type CloseRequest struct {
	CombineAppId      string      `json:"combine_appid,omitempty"`
	CombineOutTradeNo string      `json:"combine_out_trade_no,omitempty"`
	SubOrders         []*SubOrder `json:"sub_orders,omitempty"`
}

func (c *CloseRequest) Check() (err error) {
	if c.CombineAppId == "" {
		err = errors.New("请填写combine_appid!")
		return
	}
	if c.CombineOutTradeNo == "" {
		err = errors.New("请填写combine_out_trade_no!")
		return
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
		if order.OutTradeNo == "" {
			err = errors.New("请填写sub_orders.out_trade_no!")
			return
		}
	}
	return
}

type CloseResponse struct {
	model.WechatError
	RequestId string `json:"request_id,omitempty"`
}

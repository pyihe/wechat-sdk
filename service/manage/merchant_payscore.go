package manage

import (
	"fmt"

	"github.com/pyihe/go-pkg/errors"

	"github.com/pyihe/wechat-sdk/model/manage/merchant"
	"github.com/pyihe/wechat-sdk/service"
	"github.com/pyihe/wechat-sdk/vars"
)

/*微信支付分(公共API)*/

// CreatePayscoreOrder 创建支付分订单
// API 详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_14.shtml
func CreatePayscoreOrder(config *service.Config, request *merchant.CreatePayscoreOrderRequest) (payscoreOrder *merchant.PayscoreOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	if err = request.Check(); err != nil {
		return
	}
	response, err := service.RequestWithSign(config, "POST", "/v3/payscore/serviceorder", request)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	payscoreOrder = new(merchant.PayscoreOrder)
	payscoreOrder.Id = requestId
	err = service.Unmarshal(body, &payscoreOrder)
	return
}

// QueryPayscoreOrder 查询支付分订单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_15.shtml
func QueryPayscoreOrder(config *service.Config, request *merchant.QueryPayscoreOrderRequest) (queryResponse *merchant.PayscoreOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	if err = request.Check(); err != nil {
		return
	}
	var abUrl string
	switch {
	case request.OutOrderNo != "":
		abUrl = fmt.Sprintf("/v3/payscore/serviceorder?service_id=%s&out_order_id=%s&appid=%s", request.ServiceId, request.OutOrderNo, request.AppId)
	case request.QueryId != "":
		abUrl = fmt.Sprintf("/v3/payscore/serviceorder?service_id=%s&query_id=%s&appid=%s", request.ServiceId, request.QueryId, request.AppId)
	default:
		err = errors.New("请检查查询条件是否符合文档要求!")
		return
	}
	response, err := service.RequestWithSign(config, "GET", abUrl, request)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	queryResponse = new(merchant.PayscoreOrder)
	queryResponse.Id = requestId
	err = service.Unmarshal(body, &queryResponse)
	return
}

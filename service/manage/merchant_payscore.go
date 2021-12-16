package manage

import (
	"fmt"

	"github.com/pyihe/go-pkg/errors"

	"github.com/pyihe/wechat-sdk/model/manage/merchant"
	"github.com/pyihe/wechat-sdk/service"
)

/*微信支付分(公共API)*/

// CreatePayscoreOrder 创建支付分订单
// API 详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_14.shtml
func CreatePayscoreOrder(config *service.Config, request *merchant.CreatePayscoreOrderRequest) (payscoreOrder *merchant.PayscoreOrder, err error) {
	if err = service.CheckParam(config, request); err != nil {
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
	if err = service.CheckParam(config, request); err != nil {
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

// CancelPayscoreOrder 取消支付分订单
// API地址: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_16.shtml
func CancelPayscoreOrder(config *service.Config, request *merchant.CancelPayscoreOrderRequest) (cancelResponse *merchant.CancelResponse, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}
	var abUrl = fmt.Sprintf("/v3/payscore/serviceorder/%s/cancel", request.OutOrderNo)
	response, err := service.RequestWithSign(config, "POST", abUrl, request)
	if err != nil {
		return
	}

	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}

	cancelResponse = new(merchant.CancelResponse)
	cancelResponse.RequestId = requestId
	err = service.Unmarshal(body, &cancelResponse)
	return
}

// ModifyPayscoreOrder 修改订单金额
// API地址: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_17.shtml
func ModifyPayscoreOrder(config *service.Config, request *merchant.ModifyPayscoreOrderRequest) (modifyResponse *merchant.PayscoreOrder, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}
	var abUrl = fmt.Sprintf("/v3/payscore/serviceorder/%s/modify", request.OutOrderNo)
	response, err := service.RequestWithSign(config, "POST", abUrl, request)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	modifyResponse = new(merchant.PayscoreOrder)
	modifyResponse.Id = requestId
	err = service.Unmarshal(body, &modifyResponse)
	return
}

// CompletePayscoreOrder 完结支付分订单
// API地址: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_18.shtml
func CompletePayscoreOrder(config *service.Config, request *merchant.CompletePayscoreOrderRequest) (orderResponse *merchant.PayscoreOrder, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}
	var abUrl = fmt.Sprintf("/v3/payscore/serviceorder/%s/complete", request.OutOrderNo)
	response, err := service.RequestWithSign(config, "POST", abUrl, request)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}

	orderResponse = new(merchant.PayscoreOrder)
	orderResponse.Id = requestId
	err = service.Unmarshal(body, &orderResponse)
	return
}

// PayscorePay 商户发起催收扣款
// API地址: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_19.shtml
func PayscorePay(config *service.Config, request *merchant.PayscorePayRequest) (payResponse *merchant.PayscorePayResponse, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}

	var abUrl = fmt.Sprintf("/v3/payscore/serviceorder/%s/pay", request.OutOrderNo)
	response, err := service.RequestWithSign(config, "POST", abUrl, request)
	if err != nil {
		return
	}

	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	payResponse = new(merchant.PayscorePayResponse)
	payResponse.RequestId = requestId
	err = service.Unmarshal(body, &payResponse)
	return
}

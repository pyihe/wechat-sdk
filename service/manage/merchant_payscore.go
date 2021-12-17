package manage

import (
	"fmt"
	"net/http"

	"github.com/pyihe/wechat-sdk/pkg/aess"

	"github.com/pyihe/wechat-sdk/model"

	"github.com/pyihe/wechat-sdk/vars"

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

// SyncPayscoreOrder 同步服务订单信息
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_20.shtml
func SyncPayscoreOrder(config *service.Config, request *merchant.SyncPayscoreOrder) (payscoreOrder *merchant.PayscoreOrder, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}
	var abUrl = fmt.Sprintf("/v3/payscore/serviceorder/%s/sync", request.OutOrderNo)
	response, err := service.RequestWithSign(config, "POST", abUrl, request)
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

// PayscorePayNotify 支付成功回调通知
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_22.shtml
func PayscorePayNotify(config *service.Config, request *http.Request) (payscoreOrder *merchant.PayscoreOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.ApiKey == "" {
		err = vars.ErrNoApiV3Key
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	body, err := service.VerifyRequest(config, request)
	if err != nil {
		return
	}
	notifyResponse := new(model.WechatNotifyResponse)
	if err = service.Unmarshal(body, &notifyResponse); err != nil {
		return
	}
	if notifyResponse.ResourceType != "encrypt-resource" {
		err = errors.New("错误的资源类型: " + notifyResponse.ResourceType)
		return
	}
	if notifyResponse.Resource == nil {
		err = errors.New("未获取到通知资源数据!")
		return
	}
	cipherText := notifyResponse.Resource.CipherText
	associateData := notifyResponse.Resource.AssociatedData
	nonce := notifyResponse.Resource.Nonce
	plainText, err := aess.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
	if err != nil {
		return
	}

	payscoreOrder = new(merchant.PayscoreOrder)
	payscoreOrder.Id = notifyResponse.Id
	err = service.Unmarshal(plainText, &payscoreOrder)
	return
}

// PayscoreRefund 申请退款
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_26.shtml
func PayscoreRefund(config *service.Config, request *merchant.PayscoreRefundRequest) (refundResponse *merchant.PayscoreRefundOrder, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}
	response, err := service.RequestWithSign(config, "POST", "/v3/refund/domestic/refunds", request)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	refundResponse = new(merchant.PayscoreRefundOrder)
	refundResponse.Id = requestId
	err = service.Unmarshal(body, &refundResponse)
	return
}

// PayscoreQueryRefund 查询退款
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_27.shtml
func PayscoreQueryRefund(config *service.Config, request *merchant.PayscoreQueryRefundRequest) (refundOrder *merchant.PayscoreRefundOrder, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}
	var abUrl = fmt.Sprintf("/v3/refund/domestic/refunds/%s", request.OutRefundNo)
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	refundOrder = new(merchant.PayscoreRefundOrder)
	refundOrder.Id = requestId
	err = service.Unmarshal(body, &refundOrder)
	return
}

func PayscoreRefundNotify(config *service.Config, request *http.Request) (refundOrder *merchant.PayscoreRefundOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.ApiKey == "" {
		err = vars.ErrNoApiV3Key
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	body, err := service.VerifyRequest(config, request)
	if err != nil {
		return
	}
	notifyResponse := new(model.WechatNotifyResponse)
	if err = service.Unmarshal(body, &notifyResponse); err != nil {
		return
	}
	// 判断资源类型
	if notifyResponse.ResourceType != "encrypt-resource" {
		err = errors.New("错误的资源类型: " + notifyResponse.ResourceType)
		return
	}
	if notifyResponse.Resource == nil {
		err = errors.New("未获取到通知资源数据!")
		return
	}
	// 解密
	cipherText := notifyResponse.Resource.CipherText
	associateData := notifyResponse.Resource.AssociatedData
	nonce := notifyResponse.Resource.Nonce
	plainText, err := aess.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
	if err != nil {
		return
	}
	refundOrder = new(merchant.PayscoreRefundOrder)
	refundOrder.Id = notifyResponse.Id
	err = service.Unmarshal(plainText, &refundOrder)
	return
}

package payscore

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/service"
)

// PrePermit 商户预授权API
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_2.shtml
func PrePermit(config *service.Config, request interface{}) (permissionResponse *PrePermitResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/payscore/permissions", request)
	if err != nil {
		return
	}
	permissionResponse = new(PrePermitResponse)
	requestId, err := config.ParseWechatResponse(response, permissionResponse)
	permissionResponse.RequestId = requestId
	return
}

// QueryPermissions 查询用户授权记录
// 通过authorization_code查询API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_3.shtml
// 通过openid查询API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_5.shtml
func QueryPermissions(config *service.Config, request *QueryPermissionsRequest) (queryResponse *QueryPermissionsResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}

	if request.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}

	apiUrl := ""
	param := make(url.Values)
	param.Add("service_id", request.ServiceId)

	switch {
	case request.AuthorizationCode != "":
		apiUrl = fmt.Sprintf("/v3/payscore/permissions/authorization-code/%s?%s", request.AuthorizationCode, param.Encode())

	case request.AppId != "" && request.OpenId != "":
		param.Add("appid", request.AppId)
		apiUrl = fmt.Sprintf("/v3/payscore/permissions/openid/%s?%s", request.OpenId, param.Encode())

	default:
		err = errors.New("参数错误, 请查看文档!")
	}

	response, err := config.RequestWithSign(http.MethodGet, apiUrl, nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryPermissionsResponse)
	requestId, err := config.ParseWechatResponse(response, queryResponse)
	queryResponse.RequestId = requestId
	return
}

// TerminatePermission 解除用户授权关系
// 通过authorization_code解除API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_4.shtml
// 通过openid解除API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_6.shtml
func TerminatePermission(config *service.Config, request *TerminatePermissionRequest) (terminateResponse *TerminatePermissionResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	if request.ServiceId == "" {
		err = errors.New("请提供service_id!")
		return
	}
	if request.Reason == "" {
		err = errors.New("请提供reason!")
		return
	}
	apiUrl := ""
	switch {
	case request.AuthorizationCode != "":
		request.AppId = ""
		request.OpenId = ""
		apiUrl = fmt.Sprintf("/v3/payscore/permissions/authorization-code/%s/terminate", request.AuthorizationCode)

	case request.OpenId != "" && request.AppId != "":
		request.AuthorizationCode = ""
		apiUrl = fmt.Sprintf("/v3/payscore/permissions/openid/%s/terminate", request.OpenId)

	default:
		err = errors.New("参数错误!")
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, apiUrl, request)
	if err != nil {
		return
	}
	terminateResponse = new(TerminatePermissionResponse)
	requestId, err := config.ParseWechatResponse(response, terminateResponse)
	terminateResponse.RequestId = requestId
	return
}

// ParseOpenOrCloseNotify 解析开启/解除授权服务回调通知
// 文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_10.shtml
func ParseOpenOrCloseNotify(config *service.Config, request *http.Request) (response *OpenOrCloseResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	response = new(OpenOrCloseResponse)
	notifyId, err := config.ParseWechatNotify(request, response)
	response.Id = notifyId
	return
}

// ParseConfirmOrderNotify 解析确认订单通知内容
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_21.shtml
func ParseConfirmOrderNotify(config *service.Config, request *http.Request) (confirmResponse *ServiceOrder, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	confirmResponse = new(ServiceOrder)
	notifyId, err := config.ParseWechatNotify(request, confirmResponse)
	confirmResponse.Id = notifyId
	return
}

// CreateServiceOrder 创建支付订单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_14.shtml
func CreateServiceOrder(config *service.Config, request interface{}) (serviceOrder *ServiceOrder, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if reflect.ValueOf(request).IsZero() {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/payscore/serviceorder", request)
	if err != nil {
		return
	}
	serviceOrder = new(ServiceOrder)
	requestId, err := config.ParseWechatResponse(response, serviceOrder)
	serviceOrder.Id = requestId
	return
}

// QueryServiceOrder 查询支付分订单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_15.shtml
func QueryServiceOrder(config *service.Config, request *QueryOrderRequest) (queryResponse *ServiceOrder, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	if request.OutOrderNo == "" && request.QueryId == "" {
		err = errors.New("请填写out_order_no或者query_id!")
		return
	}
	if request.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if request.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}

	param := make(url.Values)
	param.Add("service_id", request.ServiceId)
	param.Add("appid", request.AppId)
	if request.OutOrderNo != "" {
		param.Add("out_order_no", request.OutOrderNo)
	} else {
		param.Add("query_id", request.QueryId)
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/payscore/serviceorder?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(ServiceOrder)
	requestId, err := config.ParseWechatResponse(response, queryResponse)
	queryResponse.Id = requestId
	return
}

// CancelServiceOrder 取消支付分订单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_16.shtml
func CancelServiceOrder(config *service.Config, request *CancelRequest) (cancelResponse *CancelResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	if request.OutOrderNo == "" {
		err = errors.New("请填写out_order_no!")
		return
	}
	if request.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if request.ServiceId == "" {
		err = errors.New("请填写service_id!")
		return
	}
	if request.Reason == "" {
		err = errors.New("请填写reason!")
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/payscore/serviceorder/%s/cancel", request.OutOrderNo), request)
	if err != nil {
		return
	}
	cancelResponse = new(CancelResponse)
	requestId, err := config.ParseWechatResponse(response, cancelResponse)
	cancelResponse.RequestId = requestId
	return
}

// ModifyServiceOrder 修改订单金额
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_17.shtml
func ModifyServiceOrder(config *service.Config, outOrderNo string, request *ModifyRequest) (modifyResponse *ModifyResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if outOrderNo == "" {
		err = errors.New("请提供out_order_no!")
		return
	}
	if reflect.ValueOf(request).IsZero() {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/payscore/serviceorder/%s/modify", outOrderNo), request)
	if err != nil {
		return
	}
	modifyResponse = new(ModifyResponse)
	requestId, err := config.ParseWechatResponse(response, modifyResponse)
	modifyResponse.RequestId = requestId
	return
}

// CompleteServiceOrder 完结支付分订单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_18.shtml
func CompleteServiceOrder(config *service.Config, outOrderNo string, request *CompleteRequest) (completeResponse *CompleteResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if outOrderNo == "" {
		err = errors.New("请提供out_order_no!")
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/payscore/serviceorder/%s/complete", outOrderNo), request)
	if err != nil {
		return
	}
	completeResponse = new(CompleteResponse)
	requestId, err := config.ParseWechatResponse(response, completeResponse)
	completeResponse.RequestId = requestId
	return
}

// PayServiceOrder 商户发起催收扣款
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_19.shtml
func PayServiceOrder(config *service.Config, outOrderNo string, request *PayOrderRequest) (payResponse *PayOrderResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if outOrderNo == "" {
		err = errors.New("请提供out_order_no!")
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/payscore/serviceorder/%s/pay", outOrderNo), request)
	if err != nil {
		return
	}
	payResponse = new(PayOrderResponse)
	requestId, err := config.ParseWechatResponse(response, payResponse)
	payResponse.RequestId = requestId
	return
}

// SyncServiceOrder 同步服务订单信息
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_20.shtml
func SyncServiceOrder(config *service.Config, outOrderNo string, request *SyncRequest) (syncResponse *SyncResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if outOrderNo == "" {
		err = errors.New("请提供out_order_no!")
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/payscore/serviceorder/%s/sync", outOrderNo), request)
	if err != nil {
		return
	}
	syncResponse = new(SyncResponse)
	requestId, err := config.ParseWechatResponse(response, syncResponse)
	syncResponse.RequestId = requestId
	return
}

// ParsePaymentNotify 解析支付成功回调数据
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_15.shtml
func ParsePaymentNotify(config *service.Config, request *http.Request) (payResponse *ServiceOrder, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	payResponse = new(ServiceOrder)
	notifyId, err := config.ParseWechatNotify(request, payResponse)
	payResponse.Id = notifyId
	return
}

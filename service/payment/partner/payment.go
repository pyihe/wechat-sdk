package partner

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/wechat-sdk/v3/model"
	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// JSAPI 服务商平台JSAPI支付
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_1.shtml
func JSAPI(config *service.Config, request interface{}) (jsapiResponse *model.JSAPIResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/pay/partner/transactions/jsapi"), request)
	if err != nil {
		return
	}

	jsapiResponse = new(model.JSAPIResponse)
	jsapiResponse.RequestId, err = config.ParseWechatResponse(response, jsapiResponse)
	return
}

// APP 服务商平台JSAPI支付
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_2_1.shtml
func APP(config *service.Config, request interface{}) (appResponse *model.AppResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/pay/partner/transactions/app"), request)
	if err != nil {
		return
	}
	appResponse = new(model.AppResponse)
	appResponse.RequestId, err = config.ParseWechatResponse(response, appResponse)
	return
}

// H5 服务商平台H5支付API
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_3_1.shtml
func H5(config *service.Config, request interface{}) (h5Response *model.H5Response, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/pay/partner/transactions/h5"), request)
	if err != nil {
		return
	}

	h5Response = new(model.H5Response)
	h5Response.RequestId, err = config.ParseWechatResponse(response, h5Response)
	return
}

// Native 服务商平台native支付
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_1.shtml
func Native(config *service.Config, request interface{}) (nativeResponse *model.NativeResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/pay/partner/transactions/native"), request)
	if err != nil {
		return
	}

	nativeResponse = new(model.NativeResponse)
	nativeResponse.RequestId, err = config.ParseWechatResponse(response, nativeResponse)
	return
}

// QueryOrder 服务商平台查询订单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_2.shtml
func QueryOrder(config *service.Config, request *QueryOrderRequest) (order *PrepayOrder, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	var apiUrl string
	var param = make(url.Values)
	param.Add("sp_mchid", request.SpMchId)
	param.Add("sub_mchid", request.SubMchId)

	switch {
	case request.TransactionId != "":
		apiUrl = fmt.Sprintf("/v3/pay/partner/transactions/id/%s?%s", request.OutTradeNo, param.Encode())
	case request.OutTradeNo != "":
		apiUrl = fmt.Sprintf("/v3/combine-transactions/out-trade-no/%s?%s", request.TransactionId, param.Encode())
	default:
		err = errors.ErrParam
		return
	}

	response, err := config.RequestWithSign(http.MethodGet, apiUrl, nil)
	if err != nil {
		return
	}
	order = new(PrepayOrder)
	order.Id, err = config.ParseWechatResponse(response, order)
	return
}

// CloseOrder 服务商平台关闭订单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_3.shtml
func CloseOrder(config *service.Config, request *CloseOrderRequest) (closeResponse *CloseOrderResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	param := make(url.Values)
	param.Add("sp_mchid", request.SpMchId)
	param.Add("sub_mchid", request.SubMchId)

	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/pay/partner/transactions/out-trade-no/%s/close", request.OutTradeNo), request)
	if err != nil {
		return
	}
	closeResponse = new(CloseOrderResponse)
	closeResponse.RequestId, err = config.ParseWechatResponse(response, closeResponse)
	return
}

// ParsePrepayNotify 服务商平台解析支付通知结果
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_5.shtml
func ParsePrepayNotify(config *service.Config, request *http.Request) (order *PrepayOrder, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	order = new(PrepayOrder)
	order.Id, err = config.ParseWechatNotify(request, order)
	return
}

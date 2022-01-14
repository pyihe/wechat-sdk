package combine

import (
	"fmt"
	"net/http"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/model"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// JSAPI JSAPI合单支付
// 商户平台JSAPI合单支付API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_3.shtml
// 服务商平台JSAPI合单支付API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_3.shtml
func JSAPI(config *service.Config, request interface{}) (jsapiResponse *model.JSAPIResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, "/v3/combine-transactions/jsapi", request)
	if err != nil {
		return
	}

	jsapiResponse = new(model.JSAPIResponse)
	jsapiResponse.RequestId, err = config.ParseWechatResponse(response, jsapiResponse)
	return
}

// H5 H5合单支付
// 商户平台H5合单支付API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_2.shtml
// 服务商平台H5合单支付API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_2.shtml
func H5(config *service.Config, request interface{}) (h5Response *model.H5Response, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, "/v3/combine-transactions/h5", request)
	if err != nil {
		return
	}
	h5Response = new(model.H5Response)
	h5Response.RequestId, err = config.ParseWechatResponse(response, h5Response)
	return
}

// APP APP合单支付
// 商户平台APP合单支付API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_1.shtml
// 服务商平台APP合单支付API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_1.shtml
func APP(config *service.Config, request interface{}) (appResponse *model.AppResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, "/v3/combine-transactions/app", request)
	if err != nil {
		return
	}
	appResponse = new(model.AppResponse)
	appResponse.RequestId, err = config.ParseWechatResponse(response, appResponse)
	return
}

// Native native合单支付
// 商户平台Native合单支付API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_5.shtml
// 服务商平台合单支付API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_5.shtml
func Native(config *service.Config, request interface{}) (nativeResponse *model.NativeResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, "/v3/combine-transactions/native", request)
	if err != nil {
		return
	}
	nativeResponse = new(model.NativeResponse)
	nativeResponse.RequestId, err = config.ParseWechatResponse(response, nativeResponse)
	return
}

// QueryOrder 合单查询订单API
// 商户平台合单查询订单API: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_11.shtml
// 服务商平台合单查询订单API: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_11.shtml
func QueryOrder(config *service.Config, combineOutTradeNo string) (order *PrepayOrder, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if combineOutTradeNo == "" {
		err = errors.New("请提供combine_out_trade_no!")
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/combine-transactions/out-trade-no/%s", combineOutTradeNo), nil)
	if err != nil {
		return
	}
	order = new(PrepayOrder)
	order.Id, err = config.ParseWechatResponse(response, order)
	return
}

// CloseOrder 合单关闭订单
// 商户平台合单关闭订单API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_12.shtml
// 服务商平台合单关闭订单API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_12.shtml
func CloseOrder(config *service.Config, combineOutTradeNo string, request interface{}) (closeResponse *CloseOrderResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if combineOutTradeNo == "" {
		err = errors.New("请提供combine_out_trade_no!")
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/combine-transactions/out-trade-no/%s/close", combineOutTradeNo), request)
	if err != nil {
		return
	}
	closeResponse = new(CloseOrderResponse)
	closeResponse.RequestId, err = config.ParseWechatResponse(response, closeResponse)
	return
}

// ParsePrepayNotify 解析合单支付通知结果
// 商户平台合单支付通知API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_13.shtml
// 服务商平台合单支付通知API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_13.shtml
func ParsePrepayNotify(config *service.Config, request *http.Request) (order *PrepayOrder, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	order = new(PrepayOrder)
	order.Id, err = config.ParseWechatNotify(request, order)
	return
}

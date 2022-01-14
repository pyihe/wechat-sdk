package merchant

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/model"
	"github.com/pyihe/wechat-sdk/service"
)

// JSAPI JSAPI预下单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_1.shtml
func JSAPI(config *service.Config, request interface{}) (jsapiResponse *model.JSAPIResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/pay/transactions/jsapi", request)
	if err != nil {
		return
	}

	jsapiResponse = new(model.JSAPIResponse)
	jsapiResponse.RequestId, err = config.ParseWechatResponse(response, jsapiResponse)
	return
}

// APP app预下单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_1.shtml
func APP(config *service.Config, request interface{}) (appResponse *model.AppResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/pay/transactions/app", request)
	if err != nil {
		return
	}

	appResponse = new(model.AppResponse)
	appResponse.RequestId, err = config.ParseWechatResponse(response, appResponse)
	return
}

// Native native支付预下单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_1.shtml
func Native(config *service.Config, request interface{}) (nativeResponse *model.NativeResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/pay/transactions/native", request)
	if err != nil {
		return
	}
	nativeResponse = new(model.NativeResponse)
	nativeResponse.RequestId, err = config.ParseWechatResponse(response, nativeResponse)
	return
}

// H5 h5支付预下单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_1.shtml
func H5(config *service.Config, request interface{}) (h5Response *model.H5Response, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/pay/transactions/h5", request)
	if err != nil {
		return
	}
	h5Response = new(model.H5Response)
	h5Response.RequestId, err = config.ParseWechatResponse(response, h5Response)
	return
}

// QueryOrder 商户平台查询订单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_2.shtml
func QueryOrder(config *service.Config, request *QueryOrderRequest) (order *PrepayOrder, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if config.GetMchId() == "" {
		err = service.ErrNoMchId
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}

	var apiUrl string
	var param = make(url.Values)
	param.Add("mchid", config.GetMchId())

	switch {
	case request.OutTradeNo != "":
		apiUrl = fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s?%s", request.OutTradeNo, param.Encode())
	case request.TransactionId != "":
		apiUrl = fmt.Sprintf("/v3/pay/transactions/id/%s?%s", request.TransactionId, param.Encode())
	default:
		err = errors.New("请填写transaction_id或者out_trade_no!")
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

// CloseOrder 关闭订单
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func CloseOrder(config *service.Config, outTradeNo string) (closeResponse *CloseOrderResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if outTradeNo == "" {
		err = errors.New("请提供商户订单号!")
		return
	}

	body := []byte(fmt.Sprintf("{\"mchid\": \"%s\"}", config.GetMchId()))
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s/close", outTradeNo), body)
	if err != nil {
		return
	}
	closeResponse = new(CloseOrderResponse)
	closeResponse.RequestId, err = config.ParseWechatResponse(response, closeResponse)
	return
}

// ParsePrepayNotify 处理来自微信的支付回调，包括验签、解密、回复
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_5.shtml
func ParsePrepayNotify(config *service.Config, request *http.Request) (orderResponse *PrepayOrder, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	orderResponse = new(PrepayOrder)
	orderResponse.Id, err = config.ParseWechatNotify(request, orderResponse)
	return
}

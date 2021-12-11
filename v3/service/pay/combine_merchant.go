package pay

import (
	"fmt"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/model/pay/combine"
	"github.com/pyihe/wechat-sdk/v3/service"
	"github.com/pyihe/wechat-sdk/v3/vars"
)

// CombinePrepay 合单下单，包括: JSAPI, APP, Native, H5
// API详细介绍:
// H5合单下单: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_2.shtml
// 其他合单下单: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_1.shtml
func CombinePrepay(config *service.Config, request *combine.PrepayRequest) (prepayResponse *combine.PrepayResponse, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.SerialNo == "" {
		err = vars.ErrNoSerialNo
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
	switch request.TradeType {
	case vars.JSAPI:
		abUrl = "/v3/combine-transactions/jsapi"
	case vars.APP:
		abUrl = "/v3/combine-transactions/app"
	case vars.Native:
		abUrl = "/v3/combine-transactions/native"
	case vars.H5:
		abUrl = "/v3/combine-transactions/h5"
	case vars.FacePay:
		err = errors.New("商户平台不支持刷脸支付!")
		return
	default:
		err = errors.New("未知的交易类型!")
		return
	}
	response, err := service.RequestWithSign(config, "POST", abUrl, request)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifySign(config, response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}
	prepayResponse = new(combine.PrepayResponse)
	prepayResponse.RequestId = requestId
	err = service.Unmarshal(body, &prepayResponse)
	return
}

// CombineQueryOrder 合单查询订单
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_11.shtml
func CombineQueryOrder(config *service.Config, outTradeNo string) (combineOrder *combine.PrepayOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.SerialNo == "" {
		err = vars.ErrNoSerialNo
		return
	}
	if outTradeNo == "" {
		err = errors.New("查询时合单商户号订单号不能为空!")
		return
	}
	abUrl := fmt.Sprintf("/v3/combine-transactions/out-trade-no/%s", outTradeNo)
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifySign(config, response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}
	combineOrder = new(combine.PrepayOrder)
	combineOrder.RequestId = requestId
	err = service.Unmarshal(body, &combineOrder)
	return
}

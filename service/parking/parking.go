package parking

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/service"
)

// FindParkingService 查询车牌服务开通信息
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_8_1.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_8_1.shtml
func FindParkingService(config *service.Config, request *FindRequest) (findResponse *FindResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	if request.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if request.PlateNumber == "" {
		err = errors.New("请填写plate_number!")
		return
	}
	if request.PlateColor == "" {
		err = errors.New("请填写plate_color!")
		return
	}
	if request.OpenId == "" {
		err = errors.New("请填写openid!")
		return
	}

	param := make(url.Values)
	param.Add("appid", request.AppId)
	param.Add("plate_number", request.PlateNumber)
	param.Add("plate_color", request.PlateColor)
	param.Add("openid", request.OpenId)
	if request.SubMchId != "" {
		param.Add("sub_mchid", request.SubMchId)
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/vehicle/parking/services/find?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	findResponse = new(FindResponse)
	requestId, err := config.ParseWechatResponse(response, findResponse)
	findResponse.RequestId = requestId
	return
}

// CreateParking 创建停车入场
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_8_2.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_8_2.shtml
func CreateParking(config *service.Config, request interface{}) (createResponse *CreateParkingResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if reflect.ValueOf(request).IsZero() {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/vehicle/parking/parkings", request)
	if err != nil {
		return
	}
	createResponse = new(CreateParkingResponse)
	requestId, err := config.ParseWechatResponse(response, createResponse)
	createResponse.RequestId = requestId
	return
}

// TransactionsParking 扣费受理
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_8_3.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_8_3.shtml
func TransactionsParking(config *service.Config, request interface{}) (transactionsResponse *TransactionsResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if reflect.ValueOf(request).IsZero() {
		err = service.ErrNoRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/vehicle/transactions/parking", request)
	if err != nil {
		return
	}
	transactionsResponse = new(TransactionsResponse)
	requestId, err := config.ParseWechatResponse(response, transactionsResponse)
	transactionsResponse.RequestId = requestId
	return
}

// QueryOrder 查询订单
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_8_4.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_8_4.shtml
func QueryOrder(config *service.Config, outTradeNo, subMchId string) (queryResponse *QueryResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if outTradeNo == "" {
		err = errors.New("请提供out_trade_no!")
		return
	}

	apiUrl := fmt.Sprintf("/v3/vehicle/transactions/out-trade-no/%s", outTradeNo)
	if subMchId != "" {
		param := make(url.Values)
		param.Add("sub_mchid", subMchId)
		apiUrl = fmt.Sprintf("%s?%s", apiUrl, param.Encode())
	}

	response, err := config.RequestWithSign(http.MethodGet, apiUrl, nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryResponse)
	requestId, err := config.ParseWechatResponse(response, queryResponse)
	queryResponse.RequestId = requestId
	return
}

// ParseParkingStateNotify 解析停车入场状态变更通知结果
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_8_5.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_8_5.shtml
func ParseParkingStateNotify(config *service.Config, request *http.Request) (stateResponse *ParkStateResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	stateResponse = new(ParkStateResponse)
	notifyId, err := config.ParseWechatNotify(request, stateResponse)
	stateResponse.Id = notifyId
	return
}

// ParsePaymentNotify 解析支付通知结果
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_8_6.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_8_6.shtml
func ParsePaymentNotify(config *service.Config, request *http.Request) (paymentResponse *PaymentResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	paymentResponse = new(PaymentResponse)
	notifyId, err := config.ParseWechatNotify(request, paymentResponse)
	paymentResponse.Id = notifyId
	return
}

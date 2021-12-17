package pay

import (
	"fmt"
	"net/http"

	"github.com/pyihe/wechat-sdk/pkg/aess"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/model"
	"github.com/pyihe/wechat-sdk/model/pay/combine"
	"github.com/pyihe/wechat-sdk/service"
	"github.com/pyihe/wechat-sdk/vars"
)

// CombinePrepay 合单下单，包括: JSAPI, APP, Native, H5
// API详细介绍:
// H5合单下单: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_2.shtml
// 其他合单下单: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_1.shtml
func CombinePrepay(config *service.Config, request *combine.PrepayRequest) (prepayResponse *combine.PrepayResponse, err error) {
	if err = service.CheckParam(config, request); err != nil {
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
	requestId, body, err := service.VerifyResponse(config, response)
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
	if outTradeNo == "" {
		err = errors.New("查询时合单商户号订单号不能为空!")
		return
	}
	abUrl := fmt.Sprintf("/v3/combine-transactions/out-trade-no/%s", outTradeNo)
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	combineOrder = new(combine.PrepayOrder)
	combineOrder.Id = requestId
	err = service.Unmarshal(body, &combineOrder)
	return
}

// CombineCloseOrder 合单关闭订单
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_12.shtml
func CombineCloseOrder(config *service.Config, closeRequest *combine.CloseRequest) (closeResponse *combine.CloseResponse, err error) {
	if err = service.CheckParam(config, closeRequest); err != nil {
		return
	}
	var abUrl = fmt.Sprintf("/v3/combine-transactions/out-trade-no/%s/close", closeRequest.CombineOutTradeNo)
	response, err := service.RequestWithSign(config, "POST", abUrl, closeRequest)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	closeResponse = new(combine.CloseResponse)
	closeResponse.RequestId = requestId
	err = service.Unmarshal(body, &closeResponse)
	return
}

// CombinePrepayNotify 处理合单支付的通知
// 详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_13.shtml
func CombinePrepayNotify(config *service.Config, responseWriter http.ResponseWriter, request *http.Request) (orderResponse *combine.PrepayOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	if config.ApiKey == "" {
		err = vars.ErrNoApiV3Key
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
	// 解密成功后，将明文反序列化到结构中
	orderResponse = new(combine.PrepayOrder)
	orderResponse.Id = notifyResponse.Id
	if err = service.Unmarshal(plainText, &orderResponse); err != nil {
		return
	}
	if notifyResponse.EventType == "TRANSACTION.SUCCESS" && config.CombinePrepayNotifyHandler != nil && responseWriter != nil {
		response := new(model.Response)
		response.Code = "SUCCESS"
		response.Message = "成功"
		if err = config.CombinePrepayNotifyHandler(orderResponse); err != nil {
			response.Code = "FAIL"
			response.Message = err.Error()
		}
		data, _ := service.Marshal(response)
		responseWriter.WriteHeader(http.StatusOK)
		_, _ = responseWriter.Write(data)
	}
	return
}

package pay

import (
	"fmt"
	"net/http"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/model"
	"github.com/pyihe/wechat-sdk/v3/model/merchant"
	"github.com/pyihe/wechat-sdk/v3/pkg/rsas"
	"github.com/pyihe/wechat-sdk/v3/service"
	"github.com/pyihe/wechat-sdk/v3/vars"
)

// Prepay 商户平台微信预支付下单
// API 详细介绍:
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_1.shtml (JSAPI, APP, Native, 小程序)
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_1.shtml (H5)
func Prepay(config *service.Config, prepayRequest *merchant.PrepayRequest) (prepayResponse *merchant.PrepayResponse, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.SerialNo == "" {
		err = vars.ErrNoSerialNo
		return
	}

	if config.MchId == "" {
		err = vars.ErrNoMchId
		return
	}

	if prepayRequest == nil {
		err = vars.ErrNoRequest
		return
	}
	if err = prepayRequest.Check(); err != nil {
		return
	}

	var abUrl string
	switch prepayRequest.TradeType {
	case vars.JSAPI:
		abUrl = "/v3/pay/transactions/jsapi"
	case vars.APP:
		abUrl = "/v3/pay/transactions/app"
	case vars.Native:
		abUrl = "/v3/pay/transactions/native"
	case vars.H5:
		abUrl = "/v3/pay/transactions/h5"
	case vars.FacePay:
		err = errors.New("商户平台不支持刷脸支付!")
		return
	default:
		err = errors.New("未知的交易类型!")
		return
	}

	response, err := service.RequestWithSign(config, "POST", abUrl, prepayRequest)
	if err != nil {
		return
	}

	requestId, body, err := service.VerifySign(config, response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}

	prepayResponse = new(merchant.PrepayResponse)
	prepayResponse.RequestId = requestId
	err = service.Unmarshal(body, &prepayResponse)
	return
}

// QueryOrder 订单查询
// API 详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_2.shtml
func QueryOrder(config *service.Config, queryRequest *merchant.QueryRequest) (orderResponse *merchant.PrepayOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.SerialNo == "" {
		err = vars.ErrNoSerialNo
		return
	}
	if config.MchId == "" {
		err = vars.ErrNoMchId
		return
	}
	if queryRequest == nil {
		err = vars.ErrNoRequest
		return
	}
	var abUrl string
	switch {
	case queryRequest.OutTradeNo != "":
		abUrl = fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s?mchid=%s", queryRequest.OutTradeNo, config.MchId)
	case queryRequest.TransactionId != "":
		abUrl = fmt.Sprintf("/v3/pay/transactions/id/%s?mchid=%s", queryRequest.TransactionId, config.MchId)
	default:
		err = errors.New("订单查询只支持out_trade_no或者transaction_id查询!")
		return
	}
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifySign(config, response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}
	orderResponse = new(merchant.PrepayOrder)
	orderResponse.Id = requestId
	err = service.Unmarshal(body, &requestId)
	return
}

// CloseOrder 关闭订单
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func CloseOrder(config *service.Config, outTradeNo string) (requestId string, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.SerialNo == "" {
		err = vars.ErrNoSerialNo
		return
	}

	if config.MchId == "" {
		err = vars.ErrNoMchId
		return
	}
	if outTradeNo == "" {
		err = errors.New("商户订单号不能为空!")
		return
	}

	var abUrl = fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s/close", outTradeNo)
	var body = struct {
		MchId string `json:"mchid,omitempty"`
	}{MchId: config.MchId}

	response, err := service.RequestWithSign(config, "POST", abUrl, body)
	if err != nil {
		return
	}
	requestId, _, err = service.VerifySign(config, response.Header, response.Body, response.StatusCode)
	return
}

// PrepayNotify 处理来自微信的支付回调，包括验签、解密、回复
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_5.shtml
func PrepayNotify(config *service.Config, responseWriter http.ResponseWriter, request *http.Request) (orderResponse *merchant.PrepayOrder, err error) {
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
	_, body, err := service.VerifySign(config, request.Header, request.Body)
	if err != nil {
		return
	}

	// 验签成功的话，将从request中读取出来的body反序列化到结构中
	notifyResponse := new(model.WechatNotifyResponse)
	if err = service.Unmarshal(body, &notifyResponse); err != nil {
		return
	}

	// 判断通知类型是否为支付结果通知
	if notifyResponse.EventType != "TRANSACTION.SUCCESS" {
		err = errors.New("通知类型错误: " + notifyResponse.EventType)
		return
	}

	// 判断资源类型
	if notifyResponse.ResourceType != "encrypt-resource" {
		err = errors.New("错误的资源类型: " + notifyResponse.ResourceType)
		return
	}
	// 解密
	cipherText := notifyResponse.Resource.CipherText
	associateData := notifyResponse.Resource.AssociatedData
	nonce := notifyResponse.Resource.Nonce
	plainText, err := rsas.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
	if err != nil {
		return
	}
	// 解密成功后，将明文反序列化到结构中
	orderResponse = new(merchant.PrepayOrder)
	orderResponse.Id = notifyResponse.Id
	if err = service.Unmarshal(plainText, &orderResponse); err != nil {
		return
	}
	// 如果注册了PrepayNotifyHandler, 这里将会调用，如果处理成功了，会同时给微信服务器发送成功的应答消息
	if config.PrepayNotifyHandler != nil && responseWriter != nil {
		response := new(model.Response)
		response.Code = "SUCCESS"
		response.Message = "成功"
		if err = config.PrepayNotifyHandler(orderResponse); err != nil {
			response.Code = "FAIL"
			response.Message = err.Error()
		}
		data, _ := service.Marshal(response)
		responseWriter.WriteHeader(http.StatusOK)
		_, _ = responseWriter.Write(data)
	}
	return
}

// Refund 申请退款
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml
func Refund(config *service.Config, refundRequest *merchant.RefundRequest) (refundOrder *merchant.RefundOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.SerialNo == "" {
		err = vars.ErrNoSerialNo
		return
	}

	if refundRequest == nil {
		err = vars.ErrNoRequest
		return
	}
	if err = refundRequest.Check(); err != nil {
		return
	}
	response, err := service.RequestWithSign(config, "POST", "/v3/refund/domestic/refunds", refundRequest)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifySign(config, response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}

	refundOrder = new(merchant.RefundOrder)
	refundOrder.Id = requestId
	err = service.Unmarshal(body, &refundOrder)
	return
}

// QueryRefundOrder 查询退款
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_10.shtml
func QueryRefundOrder(config *service.Config, outRefundNo string) (refundOrder *merchant.RefundOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.SerialNo == "" {
		err = vars.ErrNoSerialNo
		return
	}
	if outRefundNo == "" {
		err = errors.New("商户退款单号out_refund_no不能为空!")
		return
	}
	abUrl := fmt.Sprintf("/v3/refund/domestic/refunds/%s", outRefundNo)
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifySign(config, response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}
	refundOrder = new(merchant.RefundOrder)
	refundOrder.Id = requestId
	err = service.Unmarshal(body, &refundOrder)
	return
}

// RefundNotify 解析退款通知，如果注册了处理通知的handler同时会执行handler，并且想微信发送应答
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_11.shtml
func RefundNotify(config *service.Config, responseWriter http.ResponseWriter, request *http.Request) (refundOrder *merchant.RefundOrder, err error) {
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
	_, body, err := service.VerifySign(config, request.Header, request.Body)
	if err != nil {
		return
	}
	// 验签成功的话，将从request中读取出来的body反序列化到结构中
	notifyResponse := new(model.WechatNotifyResponse)
	if err = service.Unmarshal(body, &notifyResponse); err != nil {
		return
	}

	// 判断通知类型是否为支付结果通知
	if notifyResponse.EventType != "REFUND.SUCCESS" {
		err = errors.New("通知类型错误: " + notifyResponse.EventType)
		return
	}

	// 判断资源类型
	if notifyResponse.ResourceType != "encrypt-resource" {
		err = errors.New("错误的资源类型: " + notifyResponse.ResourceType)
		return
	}
	// 解密
	cipherText := notifyResponse.Resource.CipherText
	associateData := notifyResponse.Resource.AssociatedData
	nonce := notifyResponse.Resource.Nonce
	plainText, err := rsas.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
	if err != nil {
		return
	}
	// 解密成功后，将明文反序列化到结构中
	refundOrder = new(merchant.RefundOrder)
	refundOrder.Id = notifyResponse.Id
	if err = service.Unmarshal(plainText, &refundOrder); err != nil {
		return
	}
	if config.RefundNotifyHandler != nil && responseWriter != nil {
		response := new(model.Response)
		response.Code = "SUCCESS"
		response.Message = "成功"
		if err = config.RefundNotifyHandler(refundOrder); err != nil {
			response.Code = "FAIL"
			response.Message = err.Error()
		}
		data, _ := service.Marshal(response)
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write(data)
	}
	return
}

// TradeBill 申请交易账单
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_6.shtml
func TradeBill(config *service.Config, request *merchant.TradeBillRequest) (billResponse *merchant.BillResponse, err error) {
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
	var abUrl = fmt.Sprintf("/v3/bill/tradebill?bill_date=%s", request.BillDate)
	if request.BillType != "" {
		abUrl = fmt.Sprintf("%s&bill_type=%s", abUrl, request.BillType)
	}
	if request.TarType != "" {
		abUrl = fmt.Sprintf("%s&tar_type=%s", abUrl, request.TarType)
	}
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifySign(config, response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}
	billResponse = new(merchant.BillResponse)
	billResponse.RequestId = requestId
	err = service.Unmarshal(body, &billResponse)
	return
}

// FundFlowBill 申请资金账单
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_7.shtml
func FundFlowBill(config *service.Config, request *merchant.FundFlowRequest) (billResponse *merchant.BillResponse, err error) {
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
	abUrl := fmt.Sprintf("/v3/bill/fundflowbill?bill_date=%s", request.BillDate)
	if request.AccountType != "" {
		abUrl = fmt.Sprintf("%s&account_type=%s", abUrl, request.AccountType)
	}
	if request.TarType != "" {
		abUrl = fmt.Sprintf("%s&tar_type=%s", abUrl, request.TarType)
	}
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifySign(config, response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}
	billResponse = new(merchant.BillResponse)
	billResponse.RequestId = requestId
	err = service.Unmarshal(body, &billResponse)
	return
}

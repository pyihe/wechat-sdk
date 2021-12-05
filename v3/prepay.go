package v3

import (
	"net/http"
	"strings"

	"github.com/pyihe/go-pkg/errors"

	"github.com/pyihe/wechat-sdk/v3/vars"
)

// Prepay 预支付
// 详细文档：
// JSAPI支付: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_1.shtml
// APP支付: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_1.shtml
// Native支付: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_1.shtml
// 小程序支付: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_1.shtml
// H5支付: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_1.shtml
func (we *WeChatClient) Prepay(prePayRequest *vars.PrepayRequest) (payResponse *vars.PrepayResponse, err error) {
	if prePayRequest == nil {
		err = vars.ErrNoParam
		return
	}

	// 校验支付的必要参数
	if err = prePayRequest.Check(); err != nil {
		return
	}
	if we.serialNo == "" {
		err = vars.ErrNoSerialNo
		return
	}
	if we.mchId == "" {
		err = vars.ErrNoMchId
		return
	}

	var abUrl string
	switch prePayRequest.PayType {
	case vars.JSAPI:
		abUrl = "/v3/pay/transactions/jsapi"
	case vars.APP:
		abUrl = "/v3/pay/transactions/app"
	case vars.H5:
		abUrl = "/v3/pay/transactions/h5"
	case vars.Native:
		abUrl = "/v3/pay/transactions/native"
	}
	// 获取签名信息
	signParam, err := we.signSHA256WithRSA("POST", abUrl, prePayRequest)
	if err != nil {
		return
	}

	data, _ := signParam.GetString("body")
	response, err := we.requestWithSign("POST", we.apiDomain+abUrl, signParam, strings.NewReader(data))
	if err != nil {
		return
	}
	// 验证微信服务器返回的数据
	requestId, body, err := we.verifyWechatSign(response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}
	payResponse = new(vars.PrepayResponse)
	payResponse.RequestId = requestId
	err = unMarshal(body, &payResponse)
	return
}

// PrepayNotify 处理微信支付回调
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_5.shtml
func (we *WeChatClient) PrepayNotify(responseWriter http.ResponseWriter, request *http.Request) (*vars.OrderResponse, error) {
	if request == nil {
		return nil, errors.New("request is nil.")
	}
	// 1. 首先是验证签名
	requestId, body, err := we.verifyWechatSign(request.Header, request.Body)
	if err != nil {
		return nil, err
	}
	// 反序列化通知数据
	var notifyResource = new(vars.Notify)
	notifyResource.RequestId = requestId
	if err = unMarshal(body, &notifyResource); err != nil {
		return nil, err
	}
	if notifyResource.Resource == nil {
		return nil, errors.New("no resource.")
	}
	// 2. 其次是解密加密的通知数据
	cipherText := notifyResource.Resource.CipherText
	associateData := notifyResource.Resource.AssociatedData
	nonce := notifyResource.Resource.Nonce
	plainText, err := we.decryptAEADAES256GCM(cipherText, associateData, nonce)
	if err != nil {
		return nil, err
	}
	var order = new(vars.OrderResponse)
	if err = unMarshal(plainText, &order); err != nil {
		return nil, err
	}
	// 3. 如果调用方注册了预支付异步回调通知的handler，这里将会执行相应的handler，处理成功将会给微信服务器回复成功的消息，否则请调用方自己给微信应答
	if we.prepayNotifyFn != nil {
		err = we.prepayNotifyFn(order)
		reply := new(vars.NotifyResponse)
		reply.Code = "SUCCESS"
		reply.Message = "成功"
		if err != nil {
			reply.Code = "FAIL"
			reply.Message = err.Error()
		}
		if responseWriter != nil {
			data, _ := marshal(reply)
			responseWriter.WriteHeader(http.StatusOK)
			responseWriter.Write(data)
		}
	}
	return order, nil
}

// PrepayNotifyWithHandler 处理微信支付回调
// handler为调用者自己的业务处理逻辑，传入的参数为解密后的支付信息
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_5.shtml
func (we *WeChatClient) PrepayNotifyWithHandler(responseWriter http.ResponseWriter, request *http.Request, handler func(order *vars.OrderResponse) error) (*vars.OrderResponse, error) {
	if request == nil {
		return nil, errors.New("request is nil.")
	}
	// 1. 首先是验证签名
	requestId, body, err := we.verifyWechatSign(request.Header, request.Body)
	if err != nil {
		return nil, err
	}
	// 反序列化通知数据
	var notifyResource = new(vars.Notify)
	notifyResource.RequestId = requestId
	if err = unMarshal(body, &notifyResource); err != nil {
		return nil, err
	}
	if notifyResource.Resource == nil {
		return nil, errors.New("no resource")
	}
	// 2. 其次是解密加密的通知数据
	cipherText := notifyResource.Resource.CipherText
	associateData := notifyResource.Resource.AssociatedData
	nonce := notifyResource.Resource.Nonce
	plainText, err := we.decryptAEADAES256GCM(cipherText, associateData, nonce)
	if err != nil {
		return nil, err
	}
	var order = new(vars.OrderResponse)
	if err = unMarshal(plainText, &order); err != nil {
		return nil, err
	}
	// 3. 如果handler不为nil，将会执行handler，执行成功的话将会给微信服务端回复处理成功的消息，否则请调用方自己处理并且对异步通知进行应答
	if handler != nil {
		err = handler(order)
		reply := new(vars.NotifyResponse)
		reply.Code = "SUCCESS"
		reply.Message = "成功"
		if err != nil {
			reply.Code = "FAIL"
			reply.Message = err.Error()
		}
		if responseWriter != nil {
			data, _ := marshal(reply)
			responseWriter.WriteHeader(http.StatusOK)
			responseWriter.Write(data)
		}
	}
	return order, nil
}

package v3

import (
	"strings"

	"github.com/pyihe/wechat-sdk/v3/vars"
)

// Refund 申请退款
// 参数说明:
// 调用接口前请仔细阅读官方API文档，需要填写哪些参数，如果参数填写错误，API可能调用失败
// 申请退款API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml
func (we *WeChatClient) Refund(refundRequest *vars.RefundRequest) (refundResponse *vars.RefundResponse, err error) {
	if refundRequest == nil {
		err = vars.ErrNoParam
		return
	}

	if err = refundRequest.Check(); err != nil {
		return
	}

	if we.mchId == "" {
		err = vars.ErrNoMchId
		return
	}
	if we.serialNo == "" {
		err = vars.ErrNoSerialNo
		return
	}
	var abUrl = "/v3/refund/domestic/refunds"
	signParam, err := we.signSHA256WithRSA("POST", abUrl, refundRequest)
	if err != nil {
		return
	}
	data, _ := signParam.GetString("body")
	response, err := we.requestWithSign("POST", we.apiDomain+abUrl, signParam, strings.NewReader(data))
	if err != nil {
		return
	}
	requestId, body, err := we.verifyWechatSign(response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}
	refundResponse = new(vars.RefundResponse)
	refundResponse.RequestId = requestId
	err = unMarshal(body, &refundResponse)
	return
}

// QueryRefundOrder 查询单笔退款
func (we *WeChatClient) QueryRefundOrder(outRefundNo string) (interface{}, error) {

	return nil, nil
}

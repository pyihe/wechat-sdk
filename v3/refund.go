package v3

import (
	"strings"

	"github.com/pyihe/wechat-sdk/v3/vars"
)

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

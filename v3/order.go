package v3

import (
	"fmt"
	"strings"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/vars"
)

// QueryOrder 订单查询
// 参数说明:
// queryType: 查询类型，vars.QueryOutTradeNo, vars.QueryTransactionId, 不同类型请求的API不一样
// queryId: 需要查询的订单号, 根据查询类型填写微信服务端的transaction_id或者商户服务端的out_trade_no
// 查询接口详细介绍:
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_2.shtml
func (we *WeChatClient) QueryOrder(queryType vars.QueryType, queryId string) (queryResponse *vars.OrderResponse, err error) {
	if we.mchId == "" {
		err = vars.ErrNoMchId
		return
	}
	var abUrl string
	switch queryType {
	case vars.QueryOutTradeNo:
		abUrl = fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s?mchid=%s", queryId, we.mchId)
	case vars.QueryTransactionId:
		abUrl = fmt.Sprintf("/v3/pay/transactions/id/%s?mchid=%s", queryId, we.mchId)
	default:
		err = errors.New("unknown pay type: " + string(queryType) + ".")
		return
	}
	signParam, err := we.signSHA256WithRSA("GET", abUrl, nil)
	if err != nil {
		return
	}
	response, err := we.requestWithSign("GET", we.apiDomain+abUrl, signParam, nil)
	if err != nil {
		return
	}
	requestId, body, err := we.verifyWechatSign(response.Header, response.Body, response.StatusCode)
	if err != nil {
		return
	}
	queryResponse = new(vars.OrderResponse)
	queryResponse.RequestId = requestId
	err = unMarshal(body, &queryResponse)
	return
}

// CloseOrder 关闭订单
// 参数说明:
// outTradeNo: 商户侧订单号
// 关闭订单API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (we *WeChatClient) CloseOrder(outTradeNo string) (requestId string, err error) {
	if we.mchId == "" {
		err = vars.ErrNoMchId
		return
	}
	var abUrl = fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s/close", outTradeNo)
	var body = struct {
		MchId string `json:"mchid"`
	}{
		MchId: we.mchId,
	}
	signParam, err := we.signSHA256WithRSA("POST", abUrl, body)
	if err != nil {
		return
	}
	data, _ := signParam.GetString("body")
	response, err := we.requestWithSign("POST", we.apiDomain+abUrl, signParam, strings.NewReader(data))
	if err != nil {
		return
	}
	requestId, _, err = we.verifyWechatSign(response.Header, response.Body, response.StatusCode)
	return
}

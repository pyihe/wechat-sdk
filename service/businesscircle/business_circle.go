package businesscircle

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// SyncPoints 商圈积分同步
// 商户API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_2.shtml
// 服务商API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_6_2.shtml
func SyncPoints(config *service.Config, request interface{}) (syncResponse *SyncPointsResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if reflect.ValueOf(request).IsZero() {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/businesscircle/points/notify", request)
	if err != nil {
		return
	}

	syncResponse = new(SyncPointsResponse)
	syncResponse.RequestId, err = config.ParseWechatResponse(response, syncResponse)
	return
}

// QueryUserAuthorization 商圈积分授权查询
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_4.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_6_4.shtml
func QueryUserAuthorization(config *service.Config, request *QueryUserAuthorizationRequest) (queryResponse *QueryUserAuthorizationResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	param := make(url.Values)
	param.Add("appid", request.AppId)
	if request.SubMchId != "" {
		param.Add("sub_mchid", request.SubMchId)
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/businesscircle/user-authorizations/%s?%s", request.OpenId, param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryUserAuthorizationResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// ParseRefundNotify 解析商圈退款成功通知
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_3.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_6_3.shtml
func ParseRefundNotify(config *service.Config, request *http.Request) (refundResponse *RefundResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	refundResponse = new(RefundResponse)
	refundResponse.NotifyId, err = config.ParseWechatNotify(request, refundResponse)
	return
}

// ParsePaymentNotify 解析商圈支付结果通知
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_1.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_6_1.shtml
func ParsePaymentNotify(config *service.Config, request *http.Request) (paymentResponse *PaymentResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	paymentResponse = new(PaymentResponse)
	paymentResponse.NotifyId, err = config.ParseWechatNotify(request, paymentResponse)
	return
}

package refunds

import (
	"fmt"
	"net/http"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// Refund 申请退款
// 商户平台退款API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml
// 商户平台合单退款API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_14.shtml
// 服务商平台退款API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_9.shtml
// 服务商平台合单退款API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_14.shtml
func Refund(config *service.Config, request interface{}) (refundOrder *RefundOrder, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/refund/domestic/refunds", request)
	if err != nil {
		return
	}
	refundOrder = new(RefundOrder)
	refundOrder.Id, err = config.ParseWechatResponse(response, refundOrder)
	return
}

// QueryRefund 查询单笔退款
// 商户平台查询退款API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_10.shtml
// 商户平台合单支付查询退款API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_15.shtml
// 服务商平台查询退款API: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_10.shtml
// 服务商平台合单支付查询退款API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_15.shtml
func QueryRefund(config *service.Config, outRefundNo string) (refundOrder *RefundOrder, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/refund/domestic/refunds/%s", outRefundNo), nil)
	if err != nil {
		return
	}
	refundOrder = new(RefundOrder)
	refundOrder.Id, err = config.ParseWechatResponse(response, refundOrder)
	return
}

// ParseRefundNotify 解析退款通知结果，并返回调用方，由调用方自己根据结果处理对应的业务逻辑
// 商户平台退款通知API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_11.shtml
// 商户平台合单支付退款通知API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_16.shtml
// 服务商平台退款通知API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_11.shtml
// 服务商平台合单支付退款通知API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_16.shtml
func ParseRefundNotify(config *service.Config, request *http.Request) (refundOrder *RefundOrder, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	refundOrder = new(RefundOrder)
	refundOrder.Id, err = config.ParseWechatNotify(request, refundOrder)
	return
}

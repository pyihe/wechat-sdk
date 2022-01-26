package brand

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// CreateSharing 请求分账
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_1.shtml
func CreateSharing(config *service.Config, request interface{}) (createResponse *CreateSharingResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/brand/profitsharing/orders", request)
	if err != nil {
		return
	}
	createResponse = new(CreateSharingResponse)
	createResponse.RequestId, err = config.ParseWechatResponse(response, createResponse)
	return
}

// QuerySharing 查询分账结果
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_2.shtml
func QuerySharing(config *service.Config, subMchId, transactionId, outOrderNo string) (queryResponse *QuerySharingResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	param := make(url.Values)
	param.Add("sub_mchid", subMchId)
	param.Add("transaction_id", transactionId)
	param.Add("out_order_no", outOrderNo)

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/brand/profitsharing/orders?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QuerySharingResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// ReturnSharing 请求分账回退
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_3.shtml
func ReturnSharing(config *service.Config, request interface{}) (returnResponse *ReturnSharingResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/brand/profitsharing/returnorders", request)
	if err != nil {
		return
	}

	returnResponse = new(ReturnSharingResponse)
	returnResponse.RequestId, err = config.ParseWechatResponse(response, returnResponse)
	return
}

// QueryReturnSharing 查询分账回退结果
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_4.shtml
func QueryReturnSharing(config *service.Config, request *QueryReturnSharingRequest) (queryResponse *QueryReturnSharingResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	param := make(url.Values)
	param.Add("sub_mchid", request.SubMchId)
	param.Add("out_return_no", request.OutReturnNo)
	switch {
	case request.OrderId != "":
		param.Add("order_id", request.OrderId)
	case request.OutOrderNo != "":
		param.Add("out_order_no", request.OutOrderNo)
	default:
		err = errors.ErrParam
		return
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/brand/profitsharing/returnorders?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryReturnSharingResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// FinishSharing 完结分账
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_5.shtml
func FinishSharing(config *service.Config, request *FinishSharingRequest) (finishResponse *FinishSharingResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/brand/profitsharing/finish-order", request)
	if err != nil {
		return
	}
	finishResponse = new(FinishSharingResponse)
	finishResponse.RequestId, err = config.ParseWechatResponse(response, finishResponse)
	return
}

// QueryUnSharingAmount 查询订单剩余待分账金额
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_9.shtml
func QueryUnSharingAmount(config *service.Config, transactionId string) (queryResponse *QueryUnSharingAmountResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/brand/profitsharing/orders/%s/amounts", transactionId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryUnSharingAmountResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// QueryMaxSharingRatio 查询最大分账比例
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_10.shtml
func QueryMaxSharingRatio(config *service.Config, brandMchId string) (queryResponse *QueryMaxSharingRatioResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/brand/profitsharing/brand-configs/%s", brandMchId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryMaxSharingRatioResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// AddSharingReceiver 添加分账接收方
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_7.shtml
func AddSharingReceiver(config *service.Config, request interface{}) (addResponse *AddReceiverResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/brand/profitsharing/receivers/add", request)
	if err != nil {
		return
	}
	addResponse = new(AddReceiverResponse)
	addResponse.RequestId, err = config.ParseWechatResponse(response, addResponse)
	return
}

// DeleteSharingReceiver 删除分账接收方
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_8.shtml
func DeleteSharingReceiver(config *service.Config, request interface{}) (deleteResponse *DeleteReceiverResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/brand/profitsharing/receivers/delete", request)
	if err != nil {
		return
	}
	deleteResponse = new(DeleteReceiverResponse)
	deleteResponse.RequestId, err = config.ParseWechatResponse(response, deleteResponse)
	return
}

// ParseSharingNotify 解析分账动帐通知
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_7_6.shtml
func ParseSharingNotify(config *service.Config, request *http.Request) (notifyResponse *SharingNotifyResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	notifyResponse = new(SharingNotifyResponse)
	notifyResponse.NotifyId, err = config.ParseWechatNotify(request, notifyResponse)
	return
}

package paygiftactivity

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// CreateActivity 创建全场满额送活动
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_2.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_7_2.shtml
func CreateActivity(config *service.Config, request interface{}) (createResponse *CreateActivityResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/paygiftactivity/unique-threshold-activity", request)
	if err != nil {
		return
	}
	createResponse = new(CreateActivityResponse)
	createResponse.RequestId, err = config.ParseWechatResponse(response, createResponse)
	return
}

// QueryActivity 查询活动详情接口
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_4.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_7_4.shtml
func QueryActivity(config *service.Config, activityId string) (queryResponse *QueryActivityResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/paygiftactivity/activities/%s", activityId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryActivityResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// QueryActivityMerchant 查询活动发券商户号
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_5.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_7_5.shtml
func QueryActivityMerchant(config *service.Config, activityId string, offset, limit uint32) (queryResponse *QueryActivityMerchantResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	apiUrl := fmt.Sprintf("/v3/marketing/paygiftactivity/activities/%s/merchants", activityId)
	if limit > 0 {
		param := make(url.Values)
		param.Add("offset", fmt.Sprintf("%d", offset))
		param.Add("limit", fmt.Sprintf("%d", limit))
		apiUrl = fmt.Sprintf("%s?%s", apiUrl, param.Encode())
	}
	response, err := config.RequestWithSign(http.MethodGet, apiUrl, nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryActivityMerchantResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, response)
	return
}

// QueryActivityGoods 查询活动指定商品列表
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_6.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_7_6.shtml
func QueryActivityGoods(config *service.Config, activityId string, offset, limit uint32) (queryResponse *QueryActivityGoodsResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	apiUrl := fmt.Sprintf("/v3/marketing/paygiftactivity/activities/%s/goods", activityId)
	if limit > 0 {
		param := make(url.Values)
		param.Add("offset", fmt.Sprintf("%d", offset))
		param.Add("limit", fmt.Sprintf("%d", limit))
		apiUrl = fmt.Sprintf("%s?%s", apiUrl, param.Encode())
	}
	response, err := config.RequestWithSign(http.MethodGet, apiUrl, nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryActivityGoodsResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// TerminateActivity 终止活动
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_7.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_7_7.shtml
func TerminateActivity(config *service.Config, activityId string) (terminateResponse *TerminateActivityResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/paygiftactivity/activities/%s/terminate", activityId), nil)
	if err != nil {
		return
	}
	terminateResponse = new(TerminateActivityResponse)
	terminateResponse.RequestId, err = config.ParseWechatResponse(response, terminateResponse)
	return
}

// AddActivityMerchant 添加活动发券商户号
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_8.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_7_8.shtml
func AddActivityMerchant(config *service.Config, activityId string, request *AddActivityMerchantRequest) (addResponse *AddActivityMerchantResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/marketing/paygiftactivity/activities/%s/merchants/add", activityId), request)
	if err != nil {
		return
	}
	addResponse = new(AddActivityMerchantResponse)
	addResponse.RequestId, err = config.ParseWechatResponse(response, addResponse)
	return
}

// QueryActivityByFilter 根据一定的过滤条件查询有礼活动列表
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_9.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_7_9.shtml
func QueryActivityByFilter(config *service.Config, request *QueryActivityFilter) (queryResponse *QueryActivityFilterResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	param := make(url.Values)
	param.Add("offset", fmt.Sprintf("%d", request.Offset))
	if request.Limit == 0 {
		param.Add("limit", "20")
	} else {
		param.Add("limit", fmt.Sprintf("%d", request.Limit))
	}
	if request.ActivityName != "" {
		param.Add("activity_name", request.ActivityName)
	}
	if request.ActivityStatus != "" {
		param.Add("activity_status", request.ActivityStatus)
	}
	if request.AwardType != "" {
		param.Add("award_type", request.AwardType)
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/paygiftactivity/activities?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryActivityFilterResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// DeleteActivityMerchant 删除活动商户号
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_7_10.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_7_10.shtml
func DeleteActivityMerchant(config *service.Config, activityId string, request *DeleteActivityMerchantRequest) (delResponse *DeleteActivityMerchantResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/marketing/paygiftactivity/activities/%s/merchants/delete", activityId), request)
	if err != nil {
		return
	}
	delResponse = new(DeleteActivityMerchantResponse)
	delResponse.RequestId, err = config.ParseWechatResponse(response, delResponse)
	return
}

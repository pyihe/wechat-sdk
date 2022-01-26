package goldplan

import (
	"net/http"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// ChangeGoldPlanStatus 点金计划管理
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_1.shtml
func ChangeGoldPlanStatus(config *service.Config, request *ChangeStatusRequest) (changeResponse *ChangeStatusResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, "/v3/goldplan/merchants/changegoldplanstatus", request)
	if err != nil {
		return
	}
	changeResponse = new(ChangeStatusResponse)
	changeResponse.RequestId, err = config.ParseWechatResponse(response, changeResponse)
	return
}

// ChangeCustomPageStatus 商家小票管理
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_2.shtml
func ChangeCustomPageStatus(config *service.Config, request *ChangeCustomPageStatusRequest) (changeResponse *ChangeCustomPageStatusResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, "/v3/goldplan/merchants/changecustompagestatus", request)
	if err != nil {
		return
	}
	changeResponse = new(ChangeCustomPageStatusResponse)
	changeResponse.RequestId, err = config.ParseWechatResponse(response, changeResponse)
	return
}

// SetAdvertisingIndustryFilter 同行业过滤标签管理
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_3.shtml
func SetAdvertisingIndustryFilter(config *service.Config, request *SetIndustryFilterRequest) (setResponse *SetIndustryFilterResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, "/v3/goldplan/merchants/set-advertising-industry-filter", request)
	if err != nil {
		return
	}
	setResponse = new(SetIndustryFilterResponse)
	setResponse.RequestId, err = config.ParseWechatResponse(response, setResponse)
	return
}

// OpenAdvertisingShow 开通广告展示
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_4.shtml
func OpenAdvertisingShow(config *service.Config, request *OpenAdvertisingShowRequest) (openResponse *OpenAdvertisingShowResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	if request.SubMchId == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodPatch, "/v3/goldplan/merchants/open-advertising-show", request)
	if err != nil {
		return
	}
	openResponse = new(OpenAdvertisingShowResponse)
	openResponse.RequestId, err = config.ParseWechatResponse(response, openResponse)
	return
}

// CloseAdvertisingShow 关闭广告展示
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_5.shtml
func CloseAdvertisingShow(config *service.Config, request *CloseAdvertisingShowRequest) (closeResponse *CloseAdvertisingShowResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	if request.SubMchId == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/goldplan/merchants/close-advertising-show", request)
	if err != nil {
		return
	}
	closeResponse = new(CloseAdvertisingShowResponse)
	closeResponse.RequestId, err = config.ParseWechatResponse(response, closeResponse)
	return
}

package busifavor

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// CreateStock 创建商家券API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_1.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_1.shtml
func CreateStock(config *service.Config, request interface{}) (createResponse *CreateStockResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/busifavor/stocks", request)
	if err != nil {
		return
	}
	createResponse = new(CreateStockResponse)
	createResponse.RequestId, err = config.ParseWechatResponse(response, createResponse)
	return
}

// QueryMerchantStock 查询商家券详情
// 商户平台API: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_2.shtml
// 服务商平台API: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_2.shtml
func QueryMerchantStock(config *service.Config, stockId string) (stock *QueryMerchantStockResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if stockId == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/busifavor/stocks/%s", stockId), nil)
	if err != nil {
		return
	}
	stock = new(QueryMerchantStockResponse)
	stock.RequestId, err = config.ParseWechatResponse(response, stock)
	return
}

// UseCoupon 核销用户券
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_3.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_3.shtml
func UseCoupon(config *service.Config, request interface{}) (useResponse *UseCouponResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/busifavor/coupons/use", request)
	if err != nil {
		return
	}
	useResponse = new(UseCouponResponse)
	useResponse.RequestId, err = config.ParseWechatResponse(response, useResponse)
	return
}

// QueryUserCouponsByFilter 根据过滤条件查询用户券
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_4.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_4.shtml
func QueryUserCouponsByFilter(config *service.Config, openId string, filter *QueryUserCouponsByFilterRequest) (queryResponse *QueryUserCouponsByFilterResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if filter == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	param := make(url.Values)
	param.Add("appid", filter.AppId)
	if filter.StockId != "" {
		param.Add("stock_id", filter.StockId)
	}
	if filter.CouponState != "" {
		param.Add("coupon_state", filter.CouponState)
	}
	if filter.CreatorMerchant != "" {
		param.Add("creator_state", filter.CreatorMerchant)
	}
	if filter.BelongMerchant != "" {
		param.Add("belong_merchant", filter.BelongMerchant)
	}
	if filter.SenderMerchant != "" {
		param.Add("sender_merchant", filter.SenderMerchant)
	}
	if filter.Limit > 0 {
		param.Add("offset", fmt.Sprintf("%d", filter.Offset))
		param.Add("limit", fmt.Sprintf("%d", filter.Limit))
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/busifavor/users/%s/coupons?%s", openId, param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryUserCouponsByFilterResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// QueryUserCoupon 查询用户单张券详情
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_5.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_5.shtml
func QueryUserCoupon(config *service.Config, couponCode, appId, openId string) (queryResponse *QueryUserCouponResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/busifavor/users/%s/coupons/%s/appids/%s", openId, couponCode, appId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryUserCouponResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// UploadCouponCode 上传预存code
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_6.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_6.shtml
func UploadCouponCode(config *service.Config, stockId string, request interface{}) (uploadResponse *UploadCouponCodeResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if stockId == "" {
		err = errors.ErrParam
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/marketing/busifavor/stocks/%s/couponcodes", stockId), request)
	if err != nil {
		return
	}
	uploadResponse = new(UploadCouponCodeResponse)
	uploadResponse.RequestId, err = config.ParseWechatResponse(response, uploadResponse)
	return
}

// SetCallbacks 设置商家券事件通知URL地址
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_7.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_7.shtml
func SetCallbacks(config *service.Config, request *SetCallbackRequest) (setResponse *SetCallbackResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/busifavor/callbacks", request)
	if err != nil {
		return
	}
	setResponse = new(SetCallbackResponse)
	setResponse.RequestId, err = config.ParseWechatResponse(response, setResponse)
	return
}

// QueryCallbacks 查询商家券事件通知URL地址
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_8.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_8.shtml
func QueryCallbacks(config *service.Config, mchId string) (queryResponse *QueryCallbacksResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if mchId == "" {
		err = errors.ErrParam
		return
	}
	param := make(url.Values)
	param.Add("mchid", mchId)

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/busifavor/callbacks?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryCallbacksResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// Associate 关联订单信息
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_9.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_9.shtml
func Associate(config *service.Config, request *AssociateRequest) (assResponse *AssociateResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/busifavor/coupons/associate", request)
	if err != nil {
		return
	}
	assResponse = new(AssociateResponse)
	assResponse.RequestId, err = config.ParseWechatResponse(response, assResponse)
	return
}

// Disassociate 取消关联订单信息
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_10.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_10.shtml
func Disassociate(config *service.Config, request *DisassociateRequest) (disResponse *DisAssociateResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/busifavor/coupons/disassociate", request)
	if err != nil {
		return
	}
	disResponse = new(DisAssociateResponse)
	disResponse.RequestId, err = config.ParseWechatResponse(response, disResponse)
	return
}

// ModifyStockBudget 修改批次预算
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_11.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_11.shtml
func ModifyStockBudget(config *service.Config, stockId string, request interface{}) (budgetResponse *ModifyBudgetResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if stockId == "" {
		err = errors.ErrParam
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPatch, fmt.Sprintf("/v3/marketing/busifavor/stocks/%s/budget", stockId), request)
	if err != nil {
		return
	}
	budgetResponse = new(ModifyBudgetResponse)
	budgetResponse.RequestId, err = config.ParseWechatResponse(response, budgetResponse)
	return
}

// ModifyStock 修改商家券基本信息
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_12.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_12.shtml
func ModifyStock(config *service.Config, stockId string, request interface{}) (modifyResponse *ModifyStockResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if stockId == "" {
		err = errors.ErrParam
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPatch, fmt.Sprintf("/v3/marketing/busifavor/stocks/%s", stockId), request)
	if err != nil {
		return
	}
	modifyResponse = new(ModifyStockResponse)
	modifyResponse.RequestId, err = config.ParseWechatResponse(response, modifyResponse)
	return
}

// ReturnCoupon 申请退券
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_13.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_13.shtml
func ReturnCoupon(config *service.Config, request *ReturnRequest) (returnResponse *ReturnResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/busifavor/coupons/return", request)
	if err != nil {
		return
	}
	returnResponse = new(ReturnResponse)
	returnResponse.RequestId, err = config.ParseWechatResponse(response, returnResponse)
	return
}

// DeactivateCoupon 使券失效
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_14.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_14.shtml
func DeactivateCoupon(config *service.Config, request *DeactivateRequest) (deactivateResponse *DeactivateResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/busifavor/coupons/deactivate", request)
	if err != nil {
		return
	}
	deactivateResponse = new(DeactivateResponse)
	deactivateResponse.RequestId, err = config.ParseWechatResponse(response, deactivateResponse)
	return
}

// SubsidyPay 补差付款
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_16.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_16.shtml
func SubsidyPay(config *service.Config, request interface{}) (payResponse *SubsidyPayResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/busifavor/subsidy/pay-receipts", request)
	if err != nil {
		return
	}
	payResponse = new(SubsidyPayResponse)
	payResponse.RequestId, err = config.ParseWechatResponse(response, payResponse)
	return
}

// QuerySubsidyPayReceipt 查询营销补差付款单详情
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_18.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_2_18.shtml
func QuerySubsidyPayReceipt(config *service.Config, receiptId string) (queryResponse *QueryPayReceiptResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if receiptId == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/busifavor/subsidy/pay-receipts/%s", receiptId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryPayReceiptResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// SendCoupon 发放消费卡
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_6_1.shtml
func SendCoupon(config *service.Config, cardId string, request *SendCouponRequest) (sendResponse *SendCouponResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if cardId == "" {
		err = errors.ErrParam
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/marketing/busifavor/coupons/%s/send", cardId), request)
	if err != nil {
		return
	}
	sendResponse = new(SendCouponResponse)
	sendResponse.RequestId, err = config.ParseWechatResponse(response, sendResponse)
	return
}

// ParseReceiveCouponNotify 解析领券事件通知参数
func ParseReceiveCouponNotify(config *service.Config, request *http.Request) (recvResponse *ReceiveCouponResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	recvResponse = new(ReceiveCouponResponse)
	recvResponse.NotifyId, err = config.ParseWechatNotify(request, recvResponse)
	return
}

package apply4subject

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/pkg/rsas"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// Apply 提交申请单
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_1_1.shtml
func Apply(config *service.Config, request *ApplyRequest) (applyResponse *ApplyResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	serialNo, _ := config.GetValidPublicKey()
	if serialNo == "" {
		err = errors.ErrNoCertificate
		return
	}

	cipher := config.GetWechatCipher()
	req := request.clone()
	contactInfo := req.ContactInfo
	identityInfo := req.IdentificationInfo
	// 加密联系人信息
	{
		if contactInfo == nil {
			err = errors.ErrParam
			return
		}
		contactInfo.Name, err = rsas.EncryptOAEP(cipher, contactInfo.Name)
		if err != nil {
			return
		}
		contactInfo.Mobile, err = rsas.EncryptOAEP(cipher, contactInfo.Mobile)
		if err != nil {
			return
		}
		contactInfo.IdCardNumber, err = rsas.EncryptOAEP(cipher, contactInfo.IdCardNumber)
		if err != nil {
			return
		}
	}

	// 加密法人身份信息
	{
		if identityInfo == nil {
			err = errors.ErrParam
			return
		}
		identityInfo.IdentificationName, err = rsas.EncryptOAEP(cipher, identityInfo.IdentificationName)
		if err != nil {
			return
		}
		identityInfo.IdentificationNumber, err = rsas.EncryptOAEP(cipher, identityInfo.IdentificationNumber)
		if err != nil {
			return
		}
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/apply4subject/applyment", req, "Wechatpay-Serial", serialNo)
	if err != nil {
		return
	}
	applyResponse = new(ApplyResponse)
	applyResponse.RequestId, err = config.ParseWechatResponse(response, applyResponse)
	return
}

// CancelApplyment 撤销申请单
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_1_2.shtml
func CancelApplyment(config *service.Config, request *CancelRequest) (cancelResponse *CancelResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	var apiUrl string
	switch {
	case request.ApplymentId > 0:
		apiUrl = fmt.Sprintf("/v3/apply4subject/applyment/%d/cancel", request.ApplymentId)
	case request.BusinessCode != "":
		apiUrl = fmt.Sprintf("/v3/apply4subject/applyment/%s/cancel", request.BusinessCode)
	default:
		err = errors.ErrParam
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, apiUrl, nil)
	if err != nil {
		return
	}
	cancelResponse = new(CancelResponse)
	cancelResponse.RequestId, err = config.ParseWechatResponse(response, cancelResponse)
	return
}

// QueryApplyResult 查询申请单审核结果
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_1_3.shtml
func QueryApplyResult(config *service.Config, request *QueryApplyResultRequest) (queryResponse *QueryApplyResultResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	param := make(url.Values)
	apiUrl := "/v3/apply4subject/applyment"
	switch {
	case request.ApplymentId > 0:
		param.Add("applyment_id", fmt.Sprintf("%d", request.ApplymentId))
		apiUrl = fmt.Sprintf("%s?%s", apiUrl, param.Encode())
	case request.BusinessCode != "":
		param.Add("business_code", request.BusinessCode)
		apiUrl = fmt.Sprintf("%s?%s", apiUrl, param.Encode())
	default:
		err = errors.ErrParam
		return
	}

	response, err := config.RequestWithSign(http.MethodGet, apiUrl, nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryApplyResultResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// QueryMerchantState 获取商户开户意愿确认状态
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_1_4.shtml
func QueryMerchantState(config *service.Config, subMchId string) (queryResponse *QueryMerchantStateResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if subMchId == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/apply4subject/applyment/merchants/%s/state", subMchId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryMerchantStateResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

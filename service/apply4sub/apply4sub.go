package apply4sub

import (
	"fmt"
	"net/http"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/pkg/rsas"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// Apply 提交申请单
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_1.shtml
func Apply(config *service.Config, request *ApplyRequest) (applyResponse *ApplyResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	req := request.clone()
	if req.BusinessCode == "" {
		err = errors.ErrParam
		return
	}

	serialNo := ""
	cipher := config.GetWechatCipher()
	contactInfo := req.ContactInfo
	subjectInfo := req.SubjectInfo
	bankAccountInfo := req.BankAccountInfo

	// 获取serialNo
	serialNo, _ = config.GetValidPublicKey()
	if serialNo == "" {
		err = errors.ErrNoCertificate
		return
	}

	// 加密超级管理员相关信息
	{
		if contactInfo == nil {
			err = errors.ErrParam
			return
		}
		// 加密超级管理员的姓名
		contactInfo.ContactName, err = rsas.EncryptOAEP(cipher, contactInfo.ContactName)
		if err != nil {
			return
		}
		// 加密超级管理员证件号码
		if contactInfo.ContactIdNumber != "" {
			contactInfo.ContactIdNumber, err = rsas.EncryptOAEP(cipher, contactInfo.ContactIdNumber)
			if err != nil {
				return
			}
		}
		// 加密超级管理员手机号码
		contactInfo.MobilePhone, err = rsas.EncryptOAEP(cipher, contactInfo.MobilePhone)
		if err != nil {
			return
		}
		// 加密超级管理员邮件地址
		contactInfo.ContactEmail, err = rsas.EncryptOAEP(cipher, contactInfo.ContactEmail)
		if err != nil {
			return
		}
	}

	// 加密主体信息
	{
		if subjectInfo == nil {
			err = errors.ErrParam
			return
		}
		if subjectInfo.IdentityInfo == nil {
			err = errors.ErrParam
			return
		}
		identityInfo := request.SubjectInfo.IdentityInfo
		// 加密身份证信息
		if idCardInfo := identityInfo.IdCardInfo; idCardInfo != nil {
			idCardInfo.IdCardName, err = rsas.EncryptOAEP(cipher, idCardInfo.IdCardName)
			if err != nil {
				return
			}
			idCardInfo.IdCardNumber, err = rsas.EncryptOAEP(cipher, idCardInfo.IdCardNumber)
			if err != nil {
				return
			}
		}
		// 加密其他证件信息
		if idDocInfo := identityInfo.IdDocInfo; idDocInfo != nil {
			idDocInfo.IdDocName, err = rsas.EncryptOAEP(cipher, idDocInfo.IdDocName)
			if err != nil {
				return
			}
			idDocInfo.IdDocNumber, err = rsas.EncryptOAEP(cipher, idDocInfo.IdDocNumber)
			if err != nil {
				return
			}
		}
		if uboInfo := subjectInfo.UboInfo; uboInfo != nil {
			uboInfo.Name, err = rsas.EncryptOAEP(cipher, uboInfo.Name)
			if err != nil {
				return
			}
			uboInfo.IdNumber, err = rsas.EncryptOAEP(cipher, uboInfo.IdNumber)
			if err != nil {
				return
			}
		}
	}

	// 加密结算银行账户信息
	{
		if bankAccountInfo == nil {
			err = errors.ErrParam
			return
		}
		bankAccountInfo.AccountName, err = rsas.EncryptOAEP(cipher, bankAccountInfo.AccountName)
		if err != nil {
			return
		}
		bankAccountInfo.AccountNumber, err = rsas.EncryptOAEP(cipher, bankAccountInfo.AccountNumber)
		if err != nil {
			return
		}
	}

	response, err := config.RequestWithSign(http.MethodPost, "/v3/applyment4sub/applyment/", request, "Wechatpay-Serial", serialNo)
	if err != nil {
		return
	}
	applyResponse = new(ApplyResponse)
	applyResponse.RequestId, err = config.ParseWechatResponse(response, applyResponse)
	return
}

// QueryApplymentByApplymentId 根据applyment_id查询申请单状态
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func QueryApplymentByApplymentId(config *service.Config, applymentId uint64) (queryResponse *QueryApplymentResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/applyment4sub/applyment/applyment_id/%d", applymentId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryApplymentResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// QueryApplymentByBusinessCode 根据业务申请编号查询申请单状态
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func QueryApplymentByBusinessCode(config *service.Config, businessCode string) (queryResponse *QueryApplymentResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if businessCode == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/applyment4sub/applyment/applyment_id/%s", businessCode), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryApplymentResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// ModifySettlement 修改结算账号
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_3.shtml
func ModifySettlement(config *service.Config, request *ModifySettlementRequest) (modifyResponse *ModifySettlementResponse, err error) {
	// 校验参数
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	var serialNo string
	var cipher = config.GetWechatCipher()

	// 找到加密用的公钥信息
	serialNo, _ = config.GetValidPublicKey()
	if serialNo == "" {
		err = errors.ErrNoCertificate
		return
	}
	req := request.clone()
	req.AccountNumber, err = rsas.EncryptOAEP(cipher, request.AccountNumber)
	if err != nil {
		return
	}

	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/modify-settlement", request.SubMchId), req, "Wechatpay-Serial", serialNo)
	if err != nil {
		return
	}
	modifyResponse = new(ModifySettlementResponse)
	modifyResponse.RequestId, err = config.ParseWechatResponse(response, modifyResponse)
	return
}

// QuerySettlement 查询结算账号
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_4.shtml
func QuerySettlement(config *service.Config, subMchId string) (queryResponse *QuerySettlementResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if subMchId == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/settlement", subMchId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QuerySettlementResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

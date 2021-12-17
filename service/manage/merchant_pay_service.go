package manage

import (
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/pyihe/secret"
	"github.com/pyihe/wechat-sdk/model/manage/merchant"
	"github.com/pyihe/wechat-sdk/pkg/rsas"
	"github.com/pyihe/wechat-sdk/service"
	"github.com/pyihe/wechat-sdk/vars"
)

// GuideRegister 服务人员注册
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_4_1.shtml
func GuideRegister(config *service.Config, request merchant.GuideRegisterRequest) (guideResponse *merchant.GuideRegisterResponse, err error) {
	if err = service.CheckParam(config, &request); err != nil {
		return
	}
	if config.Cipher == nil {
		err = vars.ErrInvalidCipher
		return
	}
	config.Certificates.Range(func(key string, value interface{}) (breakOut bool) {
		publicKey, ok := value.(*rsa.PublicKey)
		if !ok {
			return
		}
		_ = config.Cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8)
		return true
	})
	request.Name, err = rsas.EncryptOAEP(config.Cipher, request.Name)
	if err != nil {
		return
	}
	request.Mobile, err = rsas.EncryptOAEP(config.Cipher, request.Mobile)
	if err != nil {
		return
	}
	response, err := service.RequestWithSign(config, "POST", "/v3/smartguide/guides", request)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	guideResponse = new(merchant.GuideRegisterResponse)
	guideResponse.RequestId = requestId
	err = service.Unmarshal(body, &guideResponse)
	return
}

// GuideAssign 服务人员分配
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_4_2.shtml
func GuideAssign(config *service.Config, request *merchant.GuideAssignRequest) (assignResponse *merchant.GuideAssignResponse, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}
	var abUrl = fmt.Sprintf("/v3/smartguide/guides/%s/assign", request.GuideId)
	response, err := service.RequestWithSign(config, "POST", abUrl, request)
	if err != nil {
		return
	}
	requestId, _, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	assignResponse = new(merchant.GuideAssignResponse)
	assignResponse.RequestId = requestId
	return
}

// GuideQuery 服务人员查询
// API 文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_4_3.shtml
func GuideQuery(config *service.Config, request merchant.GuideQueryRequest) (queryResponse *merchant.GuideQueryResponse, err error) {
	if err = service.CheckParam(config, &request); err != nil {
		return
	}
	if config.Cipher == nil {
		err = vars.ErrInvalidCipher
		return
	}
	config.Certificates.Range(func(key string, value interface{}) (breakOut bool) {
		publicKey, ok := value.(*rsa.PublicKey)
		if !ok {
			return
		}
		_ = config.Cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8)
		return true
	})
	if request.Mobile != "" {
		request.Mobile, err = rsas.EncryptOAEP(config.Cipher, request.Mobile)
		if err != nil {
			return
		}
	}
	var abUrl = fmt.Sprintf("/v3/smartguide/guides?")
	var paramUrl string
	if request.StoreId != 0 {
		paramUrl = fmt.Sprintf("%sstore_id=%d&", paramUrl, request.StoreId)
	}
	if request.UserId != "" {
		paramUrl = fmt.Sprintf("%suserid=%s&", paramUrl, request.UserId)
	}
	if request.Mobile != "" {
		paramUrl = fmt.Sprintf("%smobile=%s&", paramUrl, request.Mobile)
	}
	if request.WorkId != "" {
		paramUrl = fmt.Sprintf("%swork_id=%s&", paramUrl, request.WorkId)
	}
	if request.Offset != 0 || request.Limit != 0 {
		paramUrl = fmt.Sprintf("%slimit=%d&offset=%d&", paramUrl, request.Limit, request.Offset)
	}
	abUrl = fmt.Sprintf("%s%s", abUrl, paramUrl)
	abUrl = strings.TrimSuffix(abUrl, "?")
	abUrl = strings.TrimSuffix(abUrl, "&")
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	queryResponse = new(merchant.GuideQueryResponse)
	queryResponse.RequestId = requestId
	err = service.Unmarshal(body, queryResponse)
	return
}

// GuideUpdate 服务人员信息更新
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_4_4.shtml
func GuideUpdate(config *service.Config, request merchant.GuideUpdateRequest) (updateResponse *merchant.GuideUpdateResponse, err error) {
	if err = service.CheckParam(config, &request); err != nil {
		return
	}
	if config.Cipher == nil {
		err = vars.ErrInvalidCipher
		return
	}
	config.Certificates.Range(func(key string, value interface{}) (breakOut bool) {
		publicKey, ok := value.(*rsa.PublicKey)
		if !ok {
			return
		}
		_ = config.Cipher.SetRSAPublicKey(publicKey, secret.PKCSLevel8)
		return true
	})
	if request.Name != "" {
		request.Name, err = rsas.EncryptOAEP(config.Cipher, request.Name)
		if err != nil {
			return
		}
	}
	if request.Mobile != "" {
		request.Mobile, err = rsas.EncryptOAEP(config.Cipher, request.Mobile)
		if err != nil {
			return
		}
	}
	var abUrl = fmt.Sprintf("/v3/smartguide/guides/%s", request.GuideId)
	response, err := service.RequestWithSign(config, "PATCH", abUrl, request)
	if err != nil {
		return
	}
	requestId, _, err := service.VerifyResponse(config, response)
	updateResponse = new(merchant.GuideUpdateResponse)
	updateResponse.RequestId = requestId
	return
}

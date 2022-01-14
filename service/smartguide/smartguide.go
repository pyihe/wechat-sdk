package smartguide

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/pkg/rsas"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// Register 服务人员注册API
// 商户API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_4_1.shtml
// 服务商API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_4_1.shtml
func Register(config *service.Config, request *RegisterRequest) (registerResponse *RegisterResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	if request.Corpid == "" {
		err = errors.New("请填写corpid!")
		return
	}
	if request.UserId == "" {
		err = errors.New("请填写userid!")
		return
	}
	if request.Name == "" {
		err = errors.New("请填写name!")
		return
	}
	if request.Mobile == "" {
		err = errors.New("请填写mobile!")
		return
	}
	if request.QrCode == "" {
		err = errors.New("请填写qr_code!")
		return
	}
	if request.Avatar == "" {
		err = errors.New("请填写avatar!")
		return
	}
	cipher := config.GetCipher()
	reqClone := request.clone()
	reqClone.Name, err = rsas.EncryptOAEP(cipher, reqClone.Name)
	if err != nil {
		return
	}
	reqClone.Mobile, err = rsas.EncryptOAEP(cipher, reqClone.Mobile)
	if err != nil {
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/smartguide/guides", reqClone)
	if err != nil {
		return
	}
	registerResponse = new(RegisterResponse)
	registerResponse.RequestId, err = config.ParseWechatResponse(response, registerResponse)
	return
}

// Assign 服务人员分配
// 商户API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_4_2.shtml
// 服务商API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_4_2.shtml
func Assign(config *service.Config, request *AssignRequest) (assignResponse *AssignResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	if request.GuideId == "" {
		err = errors.New("请填写guide_id!")
		return
	}
	if request.OutTradeNo == "" {
		err = errors.New("请填写out_trade_no!")
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/smartguide/guides/%s/assign", request.GuideId), request)
	if err != nil {
		return
	}
	assignResponse = new(AssignResponse)
	assignResponse.RequestId, err = config.ParseWechatResponse(response, assignResponse)
	return
}

// Query 查询服务人员信息
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_4_3.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_4_2.shtml
func Query(config *service.Config, request *QueryRequest) (queryResponse *QueryResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}

	param := make(url.Values)
	param.Add("store_id", fmt.Sprintf("%d", request.StoreId))

	if request.UserId != "" {
		param.Add("userid", request.UserId)
	}
	if request.Mobile != "" {
		mobile := ""
		mobile, err = rsas.EncryptOAEP(config.GetCipher(), request.Mobile)
		if err != nil {
			return
		}
		param.Add("mobile", mobile)
	}
	if request.WorkId != "" {
		param.Add("work_id", request.WorkId)
	}
	if request.Limit > 0 {
		param.Add("limit", fmt.Sprintf("%d", request.Limit))
	}
	if request.Offset > 0 {
		param.Add("offset", fmt.Sprintf("%d", request.Offset))
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/smartguide/guides?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// Update 更新服务人员的信息
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_4_4.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_4_4.shtml
func Update(config *service.Config, guideId string, request *UpdateRequest) (updateResponse *UpdateResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if guideId == "" {
		err = errors.New("请提供guide_id!")
		return
	}

	var cipher = config.GetCipher()
	var cloneReq *UpdateRequest
	if request != nil && request.isZero() == false {
		cloneReq = request.clone()
		if cloneReq.Name != "" {
			cloneReq.Name, err = rsas.EncryptOAEP(cipher, cloneReq.Name)
			if err != nil {
				return
			}
		}
		if cloneReq.Mobile != "" {
			cloneReq.Mobile, err = rsas.EncryptOAEP(cipher, cloneReq.Mobile)
			if err != nil {
				return
			}
		}
	}
	response, err := config.RequestWithSign(http.MethodPatch, fmt.Sprintf("/v3/smartguide/guides/%s", guideId), cloneReq)
	if err != nil {
		return
	}
	updateResponse = new(UpdateResponse)
	updateResponse.RequestId, err = config.ParseWechatResponse(response, updateResponse)
	return
}

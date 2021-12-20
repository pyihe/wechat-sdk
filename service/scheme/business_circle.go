package scheme

import (
	"fmt"
	"net/http"

	"github.com/pyihe/go-pkg/errors"

	"github.com/pyihe/wechat-sdk/model"
	"github.com/pyihe/wechat-sdk/model/scheme"
	"github.com/pyihe/wechat-sdk/pkg/aess"
	"github.com/pyihe/wechat-sdk/service"
	"github.com/pyihe/wechat-sdk/vars"
)

// BusinessCirclePayNotify 商圈支付结果通知
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_1.shtml
func BusinessCirclePayNotify(config *service.Config, request *http.Request) (payResponse *scheme.BusinessCirclePayResponse, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	if config.ApiKey == "" {
		err = vars.ErrNoApiV3Key
		return
	}

	body, err := service.VerifyRequest(config, request)
	if err != nil {
		return
	}
	notifyResponse := new(model.WechatNotifyResponse)
	if err = service.Unmarshal(body, &notifyResponse); err != nil {
		return
	}
	// 判断资源类型
	if notifyResponse.ResourceType != "encrypt-resource" {
		err = errors.New("错误的资源类型: " + notifyResponse.ResourceType)
		return
	}
	if notifyResponse.Resource == nil {
		err = errors.New("未获取到通知资源数据!")
		return
	}
	// 解密
	cipherText := notifyResponse.Resource.CipherText
	associateData := notifyResponse.Resource.AssociatedData
	nonce := notifyResponse.Resource.Nonce
	plainText, err := aess.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
	if err != nil {
		return
	}
	payResponse = new(scheme.BusinessCirclePayResponse)
	payResponse.Id = notifyResponse.Id
	err = service.Unmarshal(plainText, &payResponse)
	return
}

// BusinessCirclePointsNotify 商圈积分同步
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_2.shtml
func BusinessCirclePointsNotify(config *service.Config, request *scheme.BusinessCirclePointsRequest) (pointsResponse *scheme.BusinessCirclePointsResponse, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}
	response, err := service.RequestWithSign(config, "POST", "/v3/businesscircle/points/notify", request)
	if err != nil {
		return
	}
	pointsResponse = new(scheme.BusinessCirclePointsResponse)
	pointsResponse.RequestId, _, err = service.VerifyResponse(config, response)
	return
}

// BusinessCircleRefundNotify 商圈退款成功通知
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_3.shtml
func BusinessCircleRefundNotify(config *service.Config, request *http.Request) (refundResponse *scheme.BusinessCircleRefundResponse, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	if config.ApiKey == "" {
		err = vars.ErrNoApiV3Key
		return
	}
	body, err := service.VerifyRequest(config, request)
	if err != nil {
		return
	}
	notifyResponse := new(model.WechatNotifyResponse)
	if err = service.Unmarshal(body, &notifyResponse); err != nil {
		return
	}
	// 判断资源类型
	if notifyResponse.ResourceType != "encrypt-resource" {
		err = errors.New("错误的资源类型: " + notifyResponse.ResourceType)
		return
	}
	if notifyResponse.Resource == nil {
		err = errors.New("未获取到通知资源数据!")
		return
	}
	// 解密
	cipherText := notifyResponse.Resource.CipherText
	associateData := notifyResponse.Resource.AssociatedData
	nonce := notifyResponse.Resource.Nonce
	plainText, err := aess.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
	if err != nil {
		return
	}
	refundResponse = new(scheme.BusinessCircleRefundResponse)
	refundResponse.Id = notifyResponse.Id
	err = service.Unmarshal(plainText, &refundResponse)
	return
}

// BusinessCircleQueryAuthorization 商圈积分授权查询
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_4.shtml
func BusinessCircleQueryAuthorization(config *service.Config, request *scheme.BusinessCircleQueryAuthorizationRequest) (queryResponse *scheme.BusinessCircleQueryAuthorizationResponse, err error) {
	if err = service.CheckParam(config, request); err != nil {
		return
	}
	var abUrl = fmt.Sprintf("/v3/businesscircle/user-authorizations/%s?appid=%s", request.OpenId, request.AppId)
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	queryResponse = new(scheme.BusinessCircleQueryAuthorizationResponse)
	queryResponse.RequestId = requestId
	err = service.Unmarshal(body, &queryResponse)
	return
}

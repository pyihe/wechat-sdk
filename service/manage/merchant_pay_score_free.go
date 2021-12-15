package manage

import (
	"fmt"
	"net/http"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/model"
	"github.com/pyihe/wechat-sdk/model/manage/merchant"
	"github.com/pyihe/wechat-sdk/pkg/rsas"
	"github.com/pyihe/wechat-sdk/service"
	"github.com/pyihe/wechat-sdk/vars"
)

/*微信支付分(免确认预授权模式)*/

// PrePermit 商户预授权
// API 详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_2.shtml
func PrePermit(config *service.Config, request *merchant.PrePermissionRequest) (permissionResponse *merchant.PrePermissionResponse, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	if err = request.Check(); err != nil {
		return
	}
	response, err := service.RequestWithSign(config, "POST", "/v3/payscore/permissions", request)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	permissionResponse = new(merchant.PrePermissionResponse)
	permissionResponse.RequestId = requestId
	err = service.Unmarshal(body, &permissionResponse)
	return
}

// QueryPermission 商户通过授权协议号或者用户openid查询与用户的授权记录
// API 详细介绍:
// 通过authorization_code，商户查询与用户授权关系: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_3.shtml
// 通过openid查询用户授权信息: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_5.shtml
func QueryPermission(config *service.Config, request *merchant.QueryPermissionRequest) (queryResponse *merchant.QueryPermissionResponse, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	if err = request.Check(); err != nil {
		return
	}
	var abUrl string
	switch {
	case request.AuthorizationCode != "":
		abUrl = fmt.Sprintf("/v3/payscore/permissions/authorization-code/%s?service_id=%s", request.AuthorizationCode, request.ServiceId)
	case request.OpenId != "" && request.AppId != "":
		abUrl = fmt.Sprintf("/v3/payscore/permissions/openid/%s?appid=%s&service_id=%s", request.OpenId, request.AppId, request.ServiceId)
	default:
		err = errors.New("请检查查询条件是否符合要求!")
		return
	}
	response, err := service.RequestWithSign(config, "GET", abUrl, nil)
	if err != nil {
		return
	}
	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	queryResponse = new(merchant.QueryPermissionResponse)
	queryResponse.RequestId = requestId
	err = service.Unmarshal(body, &queryResponse)
	return
}

// TerminatePermission 通过授权协议号或者openid解除用户授权关系
// API详细介绍:
// 通过授权协议号解除: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_4.shtml
// 通过openid解除: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_6.shtml
func TerminatePermission(config *service.Config, request *merchant.TerminatePermissionRequest) (terminateResponse *merchant.TerminatePermissionResponse, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	if err = request.Check(); err != nil {
		return
	}
	var abUrl string
	switch {
	case request.AuthorizationCode != "":
		abUrl = fmt.Sprintf("/v3/payscore/permissions/authorization-code/%s/terminate", request.AuthorizationCode)
	case request.OpenId != "" && request.AppId != "":
		abUrl = fmt.Sprintf("/v3/payscore/permissions/openid/%s/terminate", request.OpenId)
	default:
		err = errors.New("请检查解除条件是否符合要求!")
		return
	}
	response, err := service.RequestWithSign(config, "POST", abUrl, request)
	if err != nil {
		return
	}

	requestId, body, err := service.VerifyResponse(config, response)
	if err != nil {
		return
	}
	terminateResponse = new(merchant.TerminatePermissionResponse)
	terminateResponse.RequestId = requestId
	err = service.Unmarshal(body, &terminateResponse)
	return
}

// PermissionNotify 开启/解除授权服务的回调通知
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_10.shtml
func PermissionNotify(config *service.Config, responseWriter http.ResponseWriter, request *http.Request) (permissionResponse *merchant.OpenOrCloseResponse, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.ApiKey == "" {
		err = vars.ErrNoApiV3Key
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	body, err := service.VerifyRequest(config, request)
	notifyResponse := new(model.WechatNotifyResponse)
	if err = service.Unmarshal(body, &notifyResponse); err != nil {
		return
	}

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
	plainText, err := rsas.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
	if err != nil {
		return
	}

	// 将明文序列化到结果中
	permissionResponse = new(merchant.OpenOrCloseResponse)
	permissionResponse.Id = notifyResponse.Id
	if err = service.Unmarshal(plainText, &permissionResponse); err != nil {
		return
	}

	// 如果注册了PermissionNotifyHandler， 这里将会调用，同时如果handler执行成功，将会发送成功的应答数据给微信服务器
	if config.PermissionNotifyHandler != nil {
		response := new(model.Response)
		response.Code = "SUCCESS"
		response.Message = "成功"
		if err = config.PermissionNotifyHandler(permissionResponse); err != nil {
			response.Code = "FAIL"
			response.Message = err.Error()
		}
		data, _ := service.Marshal(response)
		responseWriter.WriteHeader(http.StatusOK)
		_, _ = responseWriter.Write(data)
	}
	return
}

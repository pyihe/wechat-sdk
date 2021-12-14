package manage

import (
	"github.com/pyihe/wechat-sdk/v3/model/manage/merchant"
	"github.com/pyihe/wechat-sdk/v3/service"
	"github.com/pyihe/wechat-sdk/v3/vars"
)

// PrePermission 商户预授权
// API 详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_2.shtml
func PrePermission(config *service.Config, request *merchant.PrePermissionRequest) (permissionResponse *merchant.PrePermissionResponse, err error) {
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
func QueryPermission(config *service.Config, request *merchant.QueryPermissionRequest) (data interface{}, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	return
}

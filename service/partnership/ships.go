package partnership

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// BuildPartnerShip 建立合作关系
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_5_1.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_5_1.shtml
func BuildPartnerShip(config *service.Config, idempotencyKey string, request *BuildRequest) (buildResponse *BuildResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/partnerships/build", request, "Idempotency-Key", idempotencyKey)
	if err != nil {
		return
	}
	buildResponse = new(BuildResponse)
	buildResponse.RequestId, err = config.ParseWechatResponse(response, buildResponse)
	return
}

// QueryPartnerShip 查询合作关系列表
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_5_1.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_5_1.shtml
func QueryPartnerShip(config *service.Config, request *QueryRequest) (queryResponse *QueryResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	param := make(url.Values)
	authData, _ := json.Marshal(request.AuthorizationData)
	param.Add("authorization_data", string(authData))
	if request.Partner != nil {
		partnerData, _ := json.Marshal(request.Partner)
		param.Add("partner", string(partnerData))
	}
	if request.Limit > 0 {
		param.Add("limit", fmt.Sprintf("%d", request.Limit))
		param.Add("offset", fmt.Sprintf("%d", request.Offset))
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/partnerships?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

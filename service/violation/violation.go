package violation

import (
	"fmt"
	"net/http"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// CreateViolationNotification 创建商户违规通知回调地址
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_3_1.shtml
func CreateViolationNotification(config *service.Config, notifyUrl string) (createResponse *NotificationResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	body := fmt.Sprintf(`{"notify_url": "%s"}`, notifyUrl)
	response, err := config.RequestWithSign(http.MethodPost, "/v3/merchant-risk-manage/violation-notifications", body)
	if err != nil {
		return
	}
	createResponse = new(NotificationResponse)
	createResponse.RequestId, err = config.ParseWechatResponse(response, createResponse)
	return
}

// QueryViolationNotification 查询商户违规通知回调地址
// 文档地址: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_3_2.shtml
func QueryViolationNotification(config *service.Config) (queryResponse *NotificationResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	response, err := config.RequestWithSign(http.MethodGet, "/v3/merchant-risk-manage/violation-notifications", nil)
	if err != nil {
		return
	}

	queryResponse = new(NotificationResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// ModifyViolationNotification 修改商户违规通知回调地址
// 文档地址: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_3_3.shtml
func ModifyViolationNotification(config *service.Config, notifyUrl string) (modifyResponse *NotificationResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	body := fmt.Sprintf(`{"notify_url": "%s"}`, notifyUrl)
	response, err := config.RequestWithSign(http.MethodPut, "/v3/merchant-risk-manage/violation-notifications", body)
	if err != nil {
		return
	}

	modifyResponse = new(NotificationResponse)
	modifyResponse.RequestId, err = config.ParseWechatResponse(response, modifyResponse)
	return
}

// DeleteViolationNotification 删除商户违规通知回调地址
// 文档地址: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_3_4.shtml
func DeleteViolationNotification(config *service.Config) (deleteResponse *DeleteNotificationResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	response, err := config.RequestWithSign(http.MethodDelete, "/v3/merchant-risk-manage/violation-notifications", nil)
	if err != nil {
		return
	}

	deleteResponse = new(DeleteNotificationResponse)
	deleteResponse.RequestId, err = config.ParseWechatResponse(response, deleteResponse)
	return
}

// ParseNotify 解析商户平台处置记录回调通知
// 文档地址: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_3_5.shtml
func ParseNotify(config *service.Config, request *http.Request) (notifyResponse *NotifyResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	notifyResponse = new(NotifyResponse)
	notifyResponse.NotifyId, err = config.ParseWechatNotify(request, notifyResponse)
	return
}

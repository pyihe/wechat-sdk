package complaints

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/pkg/files"
	"github.com/pyihe/wechat-sdk/v3/pkg/rsas"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// QueryComplaintList 查询投诉单列表
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_11.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_11.shtml
func QueryComplaintList(config *service.Config, request *QueryComplaintListRequest) (queryResponse *QueryComplaintListResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	param := make(url.Values)
	param.Add("begin_date", request.BeginDate)
	param.Add("end_date", request.EndDate)
	if request.Limit > 0 || request.Offset > 0 {
		param.Add("limit", fmt.Sprintf("%d", request.Limit))
		param.Add("offset", fmt.Sprintf("%d", request.Offset))
	}
	if request.ComplaintedMchId != "" {
		param.Add("complaint_mchid", request.ComplaintedMchId)
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/merchant-service/complaints-v2?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryComplaintListResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	if err != nil {
		return
	}
	cipher := config.GetMerchantCipher()
	for _, complaint := range queryResponse.Data {
		if complaint == nil || complaint.PayerPhone == "" {
			continue
		}
		var phone []byte
		phone, err = rsas.DecryptOAEP(cipher, complaint.PayerPhone)
		if err != nil {
			return
		}
		complaint.PayerPhone = string(phone)
	}
	return
}

// QueryComplaintDetail 查询投诉单详情
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_13.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_13.shtml
func QueryComplaintDetail(config *service.Config, complaintId string) (queryResponse *QueryComplaintDetailResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if complaintId == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/merchant-service/complaints-v2/%s", complaintId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryComplaintDetailResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	if err == nil && queryResponse.PayerPhone != "" {
		var phone []byte
		phone, err = rsas.DecryptOAEP(config.GetMerchantCipher(), queryResponse.PayerPhone)
		if err != nil {
			return
		}
		queryResponse.PayerPhone = string(phone)
	}
	return
}

// QueryNegotiationHistory 查询投诉协商历史
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_12.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_12.shtml
func QueryNegotiationHistory(config *service.Config, complaintId string, limit, offset uint32) (queryResponse *QueryNegotiationHistoryResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if complaintId == "" {
		err = errors.ErrParam
		return
	}
	apiUrl := fmt.Sprintf("/v3/merchant-service/complaints-v2/%s/negotiation-historys", complaintId)
	if limit > 0 || offset > 0 {
		param := make(url.Values)
		param.Add("offset", fmt.Sprintf("%d", offset))
		param.Add("limit", fmt.Sprintf("%d", limit))
		apiUrl = fmt.Sprintf("%s?%s", apiUrl, param.Encode())
	}
	response, err := config.RequestWithSign(http.MethodGet, apiUrl, nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryNegotiationHistoryResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// ParseComplaintNotify 解析投诉通知回调
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_16.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_16.shtml
func ParseComplaintNotify(config *service.Config, request *http.Request) (notifyResponse *ComplaintNotifyResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	notifyResponse = new(ComplaintNotifyResponse)
	notifyResponse.NotifyId, err = config.ParseWechatNotify(request, notifyResponse)
	return
}

// CreateNotifyUrl 创建投诉回调通知地址
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_2.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_2.shtml
func CreateNotifyUrl(config *service.Config, notifyUrl string) (createResponse *NotifyUrlResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if notifyUrl == "" {
		err = errors.ErrParam
		return
	}
	body := fmt.Sprintf(`{"url": "%s"}`, notifyUrl)
	response, err := config.RequestWithSign(http.MethodPost, "/v3/merchant-service/complaint-notifications", body)
	if err != nil {
		return
	}

	createResponse = new(NotifyUrlResponse)
	createResponse.RequestId, err = config.ParseWechatResponse(response, createResponse)
	return
}

// QueryNotifyUrl 查询投诉通知回调地址
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_3.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_3.shtml
func QueryNotifyUrl(config *service.Config) (queryResponse *NotifyUrlResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, "/v3/merchant-service/complaint-notifications", nil)
	if err != nil {
		return
	}

	queryResponse = new(NotifyUrlResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// UpdateNotifyUrl 更新投诉回调通知地址
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_4.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_4.shtml
func UpdateNotifyUrl(config *service.Config, notifyUrl string) (updateResponse *NotifyUrlResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if notifyUrl == "" {
		err = errors.ErrParam
		return
	}
	body := fmt.Sprintf(`{"url": "%s"}`, notifyUrl)
	response, err := config.RequestWithSign(http.MethodPut, "/v3/merchant-service/complaint-notifications", body)
	if err != nil {
		return
	}

	updateResponse = new(NotifyUrlResponse)
	updateResponse.RequestId, err = config.ParseWechatResponse(response, updateResponse)
	return
}

// DeleteNotifyUrl 删除投诉回调通知地址
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_5.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_5.shtml
func DeleteNotifyUrl(config *service.Config) (deleteResponse *DeleteNotifyUrlResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	response, err := config.RequestWithSign(http.MethodDelete, "/v3/merchant-service/complaint-notifications", nil)
	if err != nil {
		return
	}

	deleteResponse = new(DeleteNotifyUrlResponse)
	deleteResponse.RequestId, err = config.ParseWechatResponse(response, deleteResponse)
	return
}

// Commit 提交回复
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_14.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_14.shtml
func Commit(config *service.Config, complaintId string, request interface{}) (commitResponse *CommitResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if complaintId == "" {
		err = errors.ErrParam
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/merchant-service/complaints-v2/%s/response", complaintId), request)
	if err != nil {
		return
	}
	commitResponse = new(CommitResponse)
	commitResponse.RequestId, err = config.ParseWechatResponse(response, commitResponse)
	return
}

// Complete 反馈处理完成
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_15.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_15.shtml
func Complete(config *service.Config, complaintId, complaintMchId string) (completeResponse *CompleteResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	body := fmt.Sprintf(`{"complainted_mchid": "%s"}`, complaintMchId)
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/merchant-service/complaints-v2/%s/complete", complaintId), body)
	if err != nil {
		return
	}
	completeResponse = new(CompleteResponse)
	completeResponse.RequestId, err = config.ParseWechatResponse(response, completeResponse)
	return
}

// UploadImage 商户上传反馈图片
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_10.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_10.shtml
func UploadImage(config *service.Config, fileName string, image interface{}) (uploadResponse *UploadImageResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if fileName == "" {
		err = errors.ErrParam
		return
	}
	var content []byte
	var contentType string

	contentType, err = service.ImageExt(fileName)
	if err != nil {
		return
	}
	switch data := image.(type) {
	case string:
		content, err = ioutil.ReadFile(data)
	case []byte:
		content = data
	case io.Reader:
		content, err = ioutil.ReadAll(data)
	case io.ReadCloser:
		content, err = ioutil.ReadAll(data)
		_ = data.Close()
	default:
		err = errors.ErrImageFormatType
	}
	if err != nil {
		return
	}
	response, err := config.UploadMedia("/v3/merchant-service/images/upload", contentType, fileName, content)
	if err != nil {
		return
	}
	uploadResponse = new(UploadImageResponse)
	uploadResponse.RequestId, err = config.ParseWechatResponse(response, uploadResponse)
	return
}

// DownloadImage 图片下载
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_18.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter10_2_18.shtml
func DownloadImage(config *service.Config, mediaUrl string, imageName, imagePath string) (err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if imagePath == "" {
		imagePath = "./img"
	}
	content, err := config.Download(mediaUrl)
	if err != nil {
		return
	}
	err = files.WritToFile(imagePath, imageName, content)
	return
}

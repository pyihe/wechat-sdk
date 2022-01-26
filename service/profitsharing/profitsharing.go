package profitsharing

import (
	"crypto"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/pkg/files"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// CreateSharing 请求分账
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_1.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_1.shtml
func CreateSharing(config *service.Config, request interface{}) (sharingResponse *CreateSharingResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/profitsharing/orders", request)
	if err != nil {
		return
	}
	sharingResponse = new(CreateSharingResponse)
	sharingResponse.RequestId, err = config.ParseWechatResponse(response, sharingResponse)
	return
}

// QuerySharing 查询分账结果
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_2.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_2.shtml
func QuerySharing(config *service.Config, subMchId, transactionId, outOrderNo string) (queryResponse *QuerySharingResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	param := make(url.Values)
	param.Add("transaction_id", transactionId)
	if subMchId != "" {
		param.Add("sub_mchid", subMchId)
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/profitsharing/orders/%s?%s", outOrderNo, param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QuerySharingResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// ReturnSharing 请求分账回退
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_3.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_3.shtml
func ReturnSharing(config *service.Config, request interface{}) (returnResponse *ReturnSharingResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/profitsharing/return-orders", request)
	if err != nil {
		return
	}
	returnResponse = new(ReturnSharingResponse)
	returnResponse.RequestId, err = config.ParseWechatResponse(response, returnResponse)
	return
}

// QueryReturnSharing 查询分账回退结果
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_4.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_4.shtml
func QueryReturnSharing(config *service.Config, subMchId, outReturnNo, outOrderNo string) (queryResponse *QueryReturnSharingResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	param := make(url.Values)
	param.Add("out_order_no", outOrderNo)
	if subMchId != "" {
		param.Add("sub_mchid", subMchId)
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/profitsharing/return-orders/%s?%s", outReturnNo, param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryReturnSharingResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// Unfreeze 解冻剩余资金
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_5.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_5.shtml
func Unfreeze(config *service.Config, request interface{}) (unfreezeResponse *UnfreezeResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/profitsharing/orders/unfreeze", request)
	if err != nil {
		return
	}
	unfreezeResponse = new(UnfreezeResponse)
	unfreezeResponse.RequestId, err = config.ParseWechatResponse(response, unfreezeResponse)
	return
}

// QueryUnSharingAmount 查询订单剩余待分账金额
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_6.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_6.shtml
func QueryUnSharingAmount(config *service.Config, transactionId string) (unSplitResponse *QueryUnSharingAmountResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/profitsharing/transactions/%s/amounts", transactionId), nil)
	if err != nil {
		return
	}
	unSplitResponse = new(QueryUnSharingAmountResponse)
	unSplitResponse.RequestId, err = config.ParseWechatResponse(response, unSplitResponse)
	return
}

// QueryMaxSharingRatio 查询最大分账比例
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_7.shtml
func QueryMaxSharingRatio(config *service.Config, subMchId string) (queryResponse *QueryMaxSharingRatioResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/profitsharing/merchant-configs/%s", subMchId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryMaxSharingRatioResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// AddSharingReceiver 添加分账接收方
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_8.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_8.shtml
func AddSharingReceiver(config *service.Config, request interface{}) (addResponse *AddReceiverResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/profitsharing/receivers/add", request)
	if err != nil {
		return
	}
	addResponse = new(AddReceiverResponse)
	addResponse.RequestId, err = config.ParseWechatResponse(response, addResponse)
	return
}

// DeleteSharingReceiver 删除分账接收方
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_9.shtml
func DeleteSharingReceiver(config *service.Config, request interface{}) (deleteResponse *DeleteReceiverResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/profitsharing/receivers/delete", request)
	if err != nil {
		return
	}
	deleteResponse = new(DeleteReceiverResponse)
	deleteResponse.RequestId, err = config.ParseWechatResponse(response, deleteResponse)
	return
}

// ParseSharingNotify 解析分账动帐通知
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_10.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_10.shtml
func ParseSharingNotify(config *service.Config, request *http.Request) (notifyResponse *SharingNotifyResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	notifyResponse = new(SharingNotifyResponse)
	notifyResponse.NotifyId, err = config.ParseWechatNotify(request, notifyResponse)
	return
}

// DownloadSharingBills 下载分账账单
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_11.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_11.shtml
func DownloadSharingBills(config *service.Config, request *DownloadBillsRequest) (downloadResponse *DownloadBillResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	param := make(url.Values)
	param.Add("bill_date", request.BillDate)
	if request.SubMchId != "" {
		param.Add("sub_mchid", request.SubMchId)
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/profitsharing/bills?%s", param.Encode()), nil)
	if err != nil {
		return
	}
	downloadResponse = new(DownloadBillResponse)
	downloadResponse.RequestId, err = config.ParseWechatResponse(response, downloadResponse)
	if err != nil || downloadResponse.DownloadUrl == "" {
		return
	}

	content, err := config.Download(downloadResponse.DownloadUrl)
	if err != nil {
		return
	}

	switch downloadResponse.HashType {
	case "SHA1":
		if err = config.VerifyHashValue(crypto.SHA1, content, downloadResponse.HashValue); err != nil {
			return
		}
	default:
		err = errors.ErrInvalidHashType
		return
	}

	fileName := request.FileName
	filePath := request.FilePath
	if fileName == "" {
		fileName = fmt.Sprintf("sharing_bill_%s.xlsx", request.BillDate)
	}
	if filePath == "" {
		filePath = "./sharing"
	}
	err = files.WritToFile(filePath, fileName, content)
	return
}

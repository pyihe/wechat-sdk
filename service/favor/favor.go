package favor

import (
	"crypto"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pyihe/go-pkg/bytes"
	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/pkg/files"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// CreateStock 创建代金券批次API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_1.shtml
// 服务商平台API: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_1.shtml
func CreateStock(config *service.Config, request interface{}) (createResponse *CreateStockResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/favor/coupon-stocks", request)
	if err != nil {
		return
	}
	createResponse = new(CreateStockResponse)
	createResponse.RequestId, err = config.ParseWechatResponse(response, createResponse)
	return
}

// StartStock 激活代金券批次API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_3.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_3.shtml
func StartStock(config *service.Config, stockCreatorMchId, stockId string) (startResponse *StartStockResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	body := []byte(fmt.Sprintf("{\"stock_creator_mchid\": \"%s\"}", stockCreatorMchId))
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/marketing/favor/stocks/%s/start", stockId), body)
	if err != nil {
		return
	}
	startResponse = new(StartStockResponse)
	startResponse.RequestId, err = config.ParseWechatResponse(response, startResponse)
	return
}

// SendStock 发放代金券批次API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_2.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_2.shtml
func SendStock(config *service.Config, openId string, request *SendCouponsRequest) (sendCouponsResponse *SendStockResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if openId == "" {
		err = errors.ErrParam
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/marketing/favor/users/%s/coupons", openId), request)
	if err != nil {
		return
	}
	sendCouponsResponse = new(SendStockResponse)
	sendCouponsResponse.RequestId, err = config.ParseWechatResponse(response, sendCouponsResponse)
	return
}

// PauseStock 暂停代金券API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_13.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_13.shtml
func PauseStock(config *service.Config, stockCreatorMchId, stockId string) (pauseResponse *PauseStockResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	body := []byte(fmt.Sprintf("{\"stock_creator_mchid\": \"%s\"}", stockCreatorMchId))
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/marketing/favor/stocks/%s/pause", stockId), body)
	if err != nil {
		return
	}

	pauseResponse = new(PauseStockResponse)
	pauseResponse.RequestId, err = config.ParseWechatResponse(response, pauseResponse)
	return
}

// RestartStock 重启代金券批次API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_14.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_14.shtml
func RestartStock(config *service.Config, stockCreatorMchId, stockId string) (restartResponse *RestartStockResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	body := []byte(fmt.Sprintf("{\"stock_creator_mchid\": \"%s\"}", stockCreatorMchId))
	response, err := config.RequestWithSign(http.MethodPost, fmt.Sprintf("/v3/marketing/favor/stocks/%s/restart", stockId), body)
	if err != nil {
		return
	}

	restartResponse = new(RestartStockResponse)
	restartResponse.RequestId, err = config.ParseWechatResponse(response, restartResponse)
	return
}

// QueryStockList 条件查询批次列表API
// 商户平台API: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_4.shtml
// 服务商平台API: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_4.shtml
func QueryStockList(config *service.Config, request *QueryStockListRequest) (queryResponse *QueryStockListResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	param := make(url.Values)
	param.Add("offset", fmt.Sprintf("%d", request.Offset))
	param.Add("limit", fmt.Sprintf("%d", request.Limit))
	param.Add("stock_creator_mchid", request.StockCreatorMchId)
	if request.CreateStartTime != "" {
		param.Add("create_start_time", request.CreateStartTime)
	}
	if request.CreateEndTime != "" {
		param.Add("create_end_time", request.CreateEndTime)
	}
	if request.Status != "" {
		param.Add("status", request.Status)
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/favor/stocks?%s", param.Encode()), nil)
	if err != nil {
		return
	}

	queryResponse = new(QueryStockListResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	if err != nil {
		return
	}

	// 商户平台和服务商平台的TradeType类型不同，这里需要特殊处理
	for _, stock := range queryResponse.Data {
		useRule := stock.StockUseRule
		if useRule == nil || len(useRule.RawTradeTypes) == 0 {
			continue
		}

		switch {
		case bytes.Contain('[', useRule.RawTradeTypes): // 如果是数组的形式，则解析到[]string中
			var ts []string
			if err = json.Unmarshal(useRule.RawTradeTypes, &ts); err != nil {
				return
			}
			useRule.TradeType = ts
		default: // 否则解析到string中
			var t string
			if err = json.Unmarshal(useRule.RawTradeTypes, &t); err != nil {
				return
			}
			useRule.TradeType = append(useRule.TradeType, t)
		}
	}
	return
}

// QueryStock 查询批次详情
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_5.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_5.shtml
func QueryStock(config *service.Config, stockCreatorMchId, stockId string) (stockResponse *QueryStockResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	param := make(url.Values)
	param.Add("stock_creator_mchid", stockCreatorMchId)
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/favor/stocks/%s?%s", stockId, param.Encode()), nil)
	if err != nil {
		return
	}

	stockResponse = new(QueryStockResponse)
	stockResponse.RequestId, err = config.ParseWechatResponse(response, stockResponse)
	if err != nil {
		return
	}

	useRule := stockResponse.StockUseRule
	if useRule != nil && len(useRule.RawTradeTypes) > 0 {
		switch {
		case bytes.Contain('[', useRule.RawTradeTypes): // 如果是数组的形式，则解析到[]string中
			var ts []string
			if err = json.Unmarshal(useRule.RawTradeTypes, &ts); err != nil {
				return
			}
			useRule.TradeType = append(useRule.TradeType, ts...)
		default: // 否则解析到string中
			var t string
			if err = json.Unmarshal(useRule.RawTradeTypes, &t); err != nil {
				return
			}
			useRule.TradeType = append(useRule.TradeType, t)
		}
	}
	return
}

// QueryCoupon 查询代金券详情
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_6.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_6.shtml
func QueryCoupon(config *service.Config, couponId, appId, openId string) (queryResponse *QueryCouponResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	param := make(url.Values)
	param.Add("appid", appId)
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/favor/users/%s/coupons/%s", openId, couponId), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryCouponResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// QueryStockMerchants 查询代金券可用商户
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_7.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_7.shtml
func QueryStockMerchants(config *service.Config, request *QueryStockMerchantRequest) (queryResponse *QueryStockMerchantResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	param := make(url.Values)
	param.Add("offset", fmt.Sprintf("%d", request.Offset))
	param.Add("limit", fmt.Sprintf("%d", request.Limit))
	param.Add("stock_creator_mchid", request.StockCreatorMchId)
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/favor/stocks/%s/merchants?%s", request.StockId, param.Encode()), nil)
	if err != nil {
		return
	}

	queryResponse = new(QueryStockMerchantResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// QueryStockItems 查询代金券可用单品列表API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_8.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_8.shtml
func QueryStockItems(config *service.Config, request *QueryStockItemRequest) (queryResponse *QueryStockItemResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	param := make(url.Values)
	param.Add("offset", fmt.Sprintf("%d", request.Offset))
	param.Add("limit", fmt.Sprintf("%d", request.Limit))
	param.Add("stock_creator_mchid", request.StockCreatorMchId)
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/favor/stocks/%s/items?%s", request.StockId, param.Encode()), nil)
	if err != nil {
		return
	}

	queryResponse = new(QueryStockItemResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// QueryUserCoupons 根据商户号查询用户的券
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_9.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_9.shtml
func QueryUserCoupons(config *service.Config, request *QueryUserCouponsRequest) (queryResponse *QueryUserCouponsResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}

	param := make(url.Values)
	param.Add("appid", request.AppId)
	if request.StockId != "" {
		param.Add("stock_id", request.StockId)
	}
	if request.Status != "" {
		param.Add("status", request.Status)
	}
	if request.Limit > 0 {
		param.Add("offset", fmt.Sprintf("%d", request.Offset))
		param.Add("limit", fmt.Sprintf("%d", request.Limit))
	}

	switch {
	case request.CreatorMchId != "":
		param.Add("creator_mchid", request.CreatorMchId)
	case request.SenderMchId != "":
		param.Add("sender_mchid", request.SenderMchId)
	default:
		param.Add("available_mchid", request.AvailableMchId)
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/favor/users/%s/coupons%s", request.OpenId, param.Encode()), nil)
	if err != nil {
		return
	}
	queryResponse = new(QueryUserCouponsResponse)
	queryResponse.RequestId, err = config.ParseWechatResponse(response, queryResponse)
	return
}

// DownloadStockUseFlow 下载批次核销明细API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_10.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_10.shtml
func DownloadStockUseFlow(config *service.Config, request *DownloadRequest) (downloadResponse *DownloadResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	if request.StockId == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/favor/stocks/%s/use-flow", request.StockId), nil)
	if err != nil {
		return
	}

	downloadResponse = new(DownloadResponse)
	downloadResponse.RequestId, err = config.ParseWechatResponse(response, downloadResponse)
	if err != nil || downloadResponse.Url == "" {
		return
	}

	// 下载账单
	content, err := config.Download(downloadResponse.Url)
	if err != nil {
		return
	}
	// 校验hash值
	switch downloadResponse.HashType {
	case "SHA1":
		if err = config.VerifyHashValue(crypto.SHA1, content, downloadResponse.HashValue); err != nil {
			return
		}
	default:
		err = errors.ErrInvalidHashType
		return
	}

	// 写入文件
	fileName := request.FileName
	filePath := request.FilePath
	err = files.WritToFile(filePath, fileName, content)
	return
}

// DownloadStockRefundFlow 下载批次退款明细API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_11.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_11.shtml
func DownloadStockRefundFlow(config *service.Config, request *DownloadRequest) (downloadResponse *DownloadResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	if request.StockId == "" {
		err = errors.ErrParam
		return
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/marketing/favor/stocks/%s/refund-flow", request.StockId), nil)
	if err != nil {
		return
	}

	downloadResponse = new(DownloadResponse)
	downloadResponse.RequestId, err = config.ParseWechatResponse(response, downloadResponse)
	if err != nil || downloadResponse.Url == "" {
		return
	}

	// 下载账单
	content, err := config.Download(downloadResponse.Url)
	if err != nil {
		return
	}
	// 校验hash值
	switch downloadResponse.HashType {
	case "SHA1":
		if err = config.VerifyHashValue(crypto.SHA1, content, downloadResponse.HashValue); err != nil {
			return
		}
	default:
		err = errors.ErrInvalidHashType
		return
	}

	// 写入文件
	fileName := request.FileName
	filePath := request.FilePath
	if fileName == "" {
		fileName = fmt.Sprintf("fundflow_%s.csv", request.StockId)
	}
	if filePath == "" {
		filePath = "./stockfundflow"
	}
	err = files.WritToFile(filePath, fileName, content)
	return
}

// SetCallbacks 设置消息通知地址API
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_12.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_12.shtml
func SetCallbacks(config *service.Config, request interface{}) (settingResponse *SettingCallbacksResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if request == nil {
		err = errors.ErrNoSDKRequest
		return
	}
	response, err := config.RequestWithSign(http.MethodPost, "/v3/marketing/favor/callbacks", request)
	if err != nil {
		return
	}

	settingResponse = new(SettingCallbacksResponse)
	settingResponse.RequestId, err = config.ParseWechatResponse(response, settingResponse)
	return
}

// ParseUseNotify 解析核销事件回调通知
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_15.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_15.shtml
func ParseUseNotify(config *service.Config, request *http.Request) (useResponse *UseResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	useResponse = new(UseResponse)
	useResponse.NotifyId, err = config.ParseWechatNotify(request, useResponse)
	return
}

// UploadImage 图片上传(营销专用)
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_0_1.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_0_1.shtml
// image格式支持: 文件存储路径(/p1/p2/name.JPG), 文件二进制字节切片, 文件内容reader
func UploadImage(config *service.Config, fileName string, image interface{}) (uploadResponse *UploadImageResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}

	var content []byte     // 图片数据流
	var contentType string // 图片格式
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
	response, err := config.UploadMedia("/v3/marketing/favor/media/image-upload", contentType, fileName, content)
	if err != nil {
		return
	}
	uploadResponse = new(UploadImageResponse)
	uploadResponse.RequestId, err = config.ParseWechatResponse(response, uploadResponse)
	return
}

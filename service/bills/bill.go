package bills

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/pkg/aess"
	"github.com/pyihe/wechat-sdk/pkg/files"
	"github.com/pyihe/wechat-sdk/pkg/rsas"
	"github.com/pyihe/wechat-sdk/service"
)

// DownloadTradeBill 申请交易账单
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_6.shtml
func DownloadTradeBill(config *service.Config, request *TradeBillRequest) (billResponse *BillResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request.BillDate == "" {
		err = errors.New("请填写bill_date!")
		return
	}

	param := make(url.Values)
	param.Add("bill_date", request.BillDate)
	if request.BillType != "" {
		param.Add("bill_type", request.BillType)
	}
	if request.TarType != "" {
		param.Add("tar_type", request.TarType)
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/bill/tradebill?%s", param.Encode()), nil)
	if err != nil {
		return
	}

	billResponse = new(BillResponse)
	requestId, err := config.ParseWechatResponse(response, billResponse)
	billResponse.RequestId = requestId
	if err != nil || billResponse.DownloadUrl == "" {
		return
	}

	// 下载文件
	content, err := config.Download(billResponse.DownloadUrl)
	if err != nil {
		return
	}
	filename := request.FileName
	filePath := request.FilePath
	if filename == "" {
		filename = request.BillDate
	}
	if filePath == "" {
		filePath = "./bills"
	}
	err = files.WritToFile(filePath, filename, content)
	return
}

// DownloadFundFlowBill 申请资金账单
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_7.shtml
func DownloadFundFlowBill(config *service.Config, request *FundFlowRequest) (billResponse *BillResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request.BillDate == "" {
		err = errors.New("请填写bill_date!")
		return
	}

	param := make(url.Values)
	param.Add("bill_date", request.BillDate)
	if request.AccountType != "" {
		param.Add("account_type", request.AccountType)
	}
	if request.TarType != "" {
		param.Add("tar_type", request.TarType)
	}
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/bill/fundflowbill?%s", param.Encode()), nil)
	if err != nil {
		return
	}

	billResponse = new(BillResponse)
	requestId, err := config.ParseWechatResponse(response, billResponse)
	billResponse.RequestId = requestId
	if err != nil || billResponse.DownloadUrl == "" {
		return
	}

	// 下载文件
	content, err := config.Download(billResponse.DownloadUrl)
	if err != nil {
		return
	}
	filename := request.FileName
	filePath := request.FilePath
	if filename == "" {
		filename = request.BillDate
	}
	if filePath == "" {
		filePath = "./bills"
	}
	err = files.WritToFile(filePath, filename, content)
	return
}

// DownloadSubMerchantFundFlowBill 申请单个子商户资金账单API
// API文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_12.shtml
func DownloadSubMerchantFundFlowBill(config *service.Config, request *SubMerchantFundFlowRequest) (billResponse *SubMerchantFundFlowResponse, err error) {
	if config == nil {
		err = service.ErrInitConfig
		return
	}
	if request == nil {
		err = service.ErrNoRequest
		return
	}
	if request.SubMchId == "" {
		err = errors.New("请填写sub_mchid!")
		return
	}
	if request.BillDate == "" {
		err = errors.New("请填写bill_date!")
		return
	}
	if request.AccountType == "" {
		err = errors.New("请填写account_type!")
		return
	}
	if request.Algorithm == "" {
		request.Algorithm = "AEAD_AES_256_GCM"
	}

	param := make(url.Values)
	param.Add("sub_mchid", request.SubMchId)
	param.Add("bill_date", request.BillDate)
	param.Add("account_type", request.AccountType)
	param.Add("algorithm", request.Algorithm)
	if request.TarType == "" {
		param.Add("tar_type", request.TarType)
	}

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/bill/sub-merchant-fundflowbill?%s", param.Encode()), nil)
	if err != nil {
		return
	}

	billResponse = new(SubMerchantFundFlowResponse)
	requestId, err := config.ParseWechatResponse(response, billResponse)
	billResponse.RequestId = requestId
	if err != nil || len(billResponse.DownloadBillList) == 0 {
		return
	}

	var content []byte
	var cipher = config.GetCipher()
	// 根据序号(BillSequence)重新整理明细顺序
	var serialList = make([]*DownloadBillList, billResponse.DownloadBillCount)
	for _, list := range billResponse.DownloadBillList {
		serialList[list.BillSequence-1] = list
	}

	for _, list := range serialList {
		// 下载单个账单文件对应的数据流
		var cipherText, plainText, key []byte
		cipherText, err = config.Download(list.DownloadUrl)
		if err != nil {
			return
		}

		// 解密RSA加密后的encrypt_key
		key, err = rsas.DecryptOAEP(cipher, list.EncryptKey)
		if err != nil {
			return
		}
		plainText, err = aess.DecryptAEADAES256GCM(cipher, string(key), cipherText, "", list.Nonce)
		if err != nil {
			return
		}
		content = append(content, plainText...)
	}

	// 写入文件
	filename := request.FileName
	filePath := request.FilePath
	if filename == "" {
		filename = request.BillDate
	}
	if filePath == "" {
		filePath = "./bills"
	}
	err = files.WritToFile(filePath, filename, content)
	return
}

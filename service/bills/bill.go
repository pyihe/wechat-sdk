package bills

import (
	"crypto"
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
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_6.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_6.shtml
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
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/bill/tradebill?%s", param.Encode()), nil)
	if err != nil {
		return
	}

	billResponse = new(BillResponse)
	billResponse.RequestId, err = config.ParseWechatResponse(response, billResponse)
	if err != nil || billResponse.DownloadUrl == "" {
		return
	}

	// 下载文件
	content, err := config.Download(billResponse.DownloadUrl)
	if err != nil {
		return
	}

	// 校验hash值
	switch billResponse.HashType {
	case "SHA1":
		if err = config.VerifyHashValue(crypto.SHA1, content, billResponse.HashValue); err != nil {
			return
		}
	default:
		err = service.ErrInvalidHashType
		return
	}

	filename := request.FileName
	filePath := request.FilePath
	if filename == "" {
		filename = fmt.Sprintf("trade_bill_%s.xlsx", request.BillDate)
	}
	if filePath == "" {
		filePath = "./tradebill"
	}
	err = files.WritToFile(filePath, filename, content)
	return
}

// DownloadFundFlowBill 申请资金账单
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_7.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_7.shtml
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
	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/bill/fundflowbill?%s", param.Encode()), nil)
	if err != nil {
		return
	}

	billResponse = new(BillResponse)
	billResponse.RequestId, err = config.ParseWechatResponse(response, billResponse)
	if err != nil || billResponse.DownloadUrl == "" {
		return
	}

	// 下载文件
	content, err := config.Download(billResponse.DownloadUrl)
	if err != nil {
		return
	}

	// 校验hash值
	switch billResponse.HashType {
	case "SHA1":
		if err = config.VerifyHashValue(crypto.SHA1, content, billResponse.HashValue); err != nil {
			return
		}
	default:
		err = service.ErrInvalidHashType
		return
	}

	filename := request.FileName
	filePath := request.FilePath
	if filename == "" {
		filename = fmt.Sprintf("fund_flow_%s.xlsx", request.BillDate)
	}
	if filePath == "" {
		filePath = "./fundflow"
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

	response, err := config.RequestWithSign(http.MethodGet, fmt.Sprintf("/v3/bill/sub-merchant-fundflowbill?%s", param.Encode()), nil)
	if err != nil {
		return
	}

	billResponse = new(SubMerchantFundFlowResponse)
	billResponse.RequestId, err = config.ParseWechatResponse(response, billResponse)
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
		// 校验明文hash值
		switch list.HashType {
		case "SHA1":
			if err = config.VerifyHashValue(crypto.SHA1, plainText, list.HashValue); err != nil {
				return
			}
		default:
			err = service.ErrInvalidHashType
			return
		}

		content = append(content, plainText...)
	}

	// 写入文件
	filename := request.FileName
	filePath := request.FilePath
	if filename == "" {
		filename = fmt.Sprintf("sub_fund_flow_%s.xlsx", request.BillDate)
	}
	if filePath == "" {
		filePath = "./mchfundflow"
	}
	err = files.WritToFile(filePath, filename, content)
	return
}

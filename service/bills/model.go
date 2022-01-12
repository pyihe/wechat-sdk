package bills

import "github.com/pyihe/wechat-sdk/model"

// TradeBillRequest 申请交易账单
type TradeBillRequest struct {
	BillDate string `json:"bill_date"`           // 账单日期
	BillType string `json:"bill_type,omitempty"` // 账单类型
	TarType  string `json:"tar_type,omitempty"`  // 压缩类型
	FileName string `json:"file_name,omitempty"` // 文件存储名
	FilePath string `json:"file_path,omitempty"` // 文件存放路径
}

// FundFlowRequest 申请资金账单
type FundFlowRequest struct {
	BillDate    string `json:"bill_date"`              // 账单日期
	AccountType string `json:"account_type,omitempty"` // 资金账户类型
	TarType     string `json:"tar_type,omitempty"`     // 压缩类型
	FileName    string `json:"file_name,omitempty"`    // 文件存储名
	FilePath    string `json:"file_path,omitempty"`    // 文件存放路径
}

// BillResponse 账单申请应答
type BillResponse struct {
	model.WechatError
	RequestId   string `json:"-"`                      // 唯一请求ID
	HashType    string `json:"hash_type,omitempty"`    // 原始账单的摘要值类型
	HashValue   string `json:"hash_value,omitempty"`   // 摘要值
	DownloadUrl string `json:"download_url,omitempty"` // 账单下载地址
}

// SubMerchantFundFlowRequest 服务商平台申请单个子商户资金账单
type SubMerchantFundFlowRequest struct {
	SubMchId    string `json:"sub_mchid"`           // 子商户号
	BillDate    string `json:"bill_date"`           // 账单日期
	AccountType string `json:"account_type"`        // 资金账户类型
	Algorithm   string `json:"algorithm"`           // 加密算法
	TarType     string `json:"tar_type,omitempty"`  // 压缩格式
	FileName    string `json:"file_name,omitempty"` // 账单文件存储名
	FilePath    string `json:"file_path,omitempty"` // 账单文件存储路径
}

// SubMerchantFundFlowResponse 服务商平台申请单个子商户资金账单应答
type SubMerchantFundFlowResponse struct {
	RequestId         string              `json:"-"`
	DownloadBillCount int32               `json:"download_bill_count,omitempty"` // 下载信息总数
	DownloadBillList  []*DownloadBillList `json:"download_bill_list,omitempty"`  // 下载信息明细
}

// DownloadBillList 下载信息明细
type DownloadBillList struct {
	BillSequence int32  `json:"bill_sequence,omitempty"` // 账单文件序号, 需要将多个文件按账单序号的顺序合并为完整的资金账单文件
	DownloadUrl  string `json:"download_url,omitempty"`  // 下载地址
	EncryptKey   string `json:"encrypt_key,omitempty"`   // 加密密钥
	HashType     string `json:"hash_type,omitempty"`     // 原始账单摘要类型
	HashValue    string `json:"hash_value,omitempty"`    // 原始账单摘要值
	Nonce        string `json:"nonce,omitempty"`         // 加密账单文件使用的随机字符串
}

package profitsharing

import (
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
)

// CreateSharingResponse 请求分账应答参数
type CreateSharingResponse struct {
	model.WechatError
	SharingOrder        // 分账账单
	RequestId    string // 唯一请求ID
}

// QuerySharingResponse 查询分账结果应答参数
type QuerySharingResponse struct {
	model.WechatError
	SharingOrder        // 分账账单
	RequestId    string // 唯一请求ID
}

// ReturnSharingResponse 请求分账回退请求参数
type ReturnSharingResponse struct {
	model.WechatError
	ReturnOrder        // 分账回退账单
	RequestId   string // 唯一请求ID
}

// QueryReturnSharingResponse 查询分账回退结果API应答参数
type QueryReturnSharingResponse struct {
	model.WechatError
	ReturnOrder        // 分账回退账单
	RequestId   string // 唯一请求ID
}

// UnfreezeResponse 解冻剩余资金应答参数
type UnfreezeResponse struct {
	model.WechatError
	RequestId    string // 唯一请求ID
	SharingOrder        // 分账账单
}

// QueryUnSharingAmountResponse 查询剩余待分金额应答参数
type QueryUnSharingAmountResponse struct {
	model.WechatError
	RequestId     string // 唯一请求ID
	TransactionId string `json:"transaction_id,omitempty"`  // 微信订单号
	UnSplitAmount int64  `json:"un_split_amount,omitempty"` // 订单剩余待分金额
}

// QueryMaxSharingRatioResponse 查询最大分账比例应答参数
type QueryMaxSharingRatioResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	SubMchId  string `json:"sub_mchid,omitempty"` // 子商户号
	MaxRatio  int32  `json:"max_ratio,omitempty"` // 最大分账比例
}

// AddReceiverResponse 添加分账接收方
type AddReceiverResponse struct {
	model.WechatError
	RequestId      string // 唯一请求ID
	SubMchId       string `json:"sub_mchid,omitempty"`       // 子商户号, 服务商平台返回
	Type           string `json:"type,omitempty"`            // 分账接收方类型
	Account        string `json:"account,omitempty"`         // 分账接收方账号
	Name           string `json:"name,omitempty"`            // 分账接收方全称
	RelationType   string `json:"relation_type,omitempty"`   // 与分账方的关系类型
	CustomRelation string `json:"custom_relation,omitempty"` // 自定义的分账关系
}

// DeleteReceiverResponse 删除分账接收方
type DeleteReceiverResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	Type      string `json:"type,omitempty"`    // 分账接受方类型
	Account   string `json:"account,omitempty"` // 分账接收方账号
}

// SharingNotifyResponse 分账动帐通知参数
type SharingNotifyResponse struct {
	NotifyId      string          // 唯一通知ID
	SpMchId       string          `json:"sp_mchid,omitempty"`       // 服务商商户号: 服务商平台返回
	SubMchId      string          `json:"sub_mchid,omitempty"`      // 子商户号: 服务商平台返回
	MchId         string          `json:"mchid,omitempty"`          // 直连商户号: 商户平台返回
	TransactionId string          `json:"transaction_id,omitempty"` // 微信订单号
	OrderId       string          `json:"order_id,omitempty"`       // 微信分账/回退单号
	OutOrderNo    string          `json:"out_order_no,omitempty"`   // 商户分账/回退单号
	SuccessTime   time.Time       `json:"success_time,omitempty"`   // 成功时间
	Receiver      *NotifyReceiver `json:"receiver,omitempty"`       // 分账接收方列表
}

// DownloadBillsRequest 申请分账账单请求参数
type DownloadBillsRequest struct {
	BillDate string // 账单日期
	SubMchId string // 子商户号
	FileName string // 账单文件的存储名
	FilePath string // 账单文件的存储路径
}

// DownloadBillResponse 分账账单应答参数
type DownloadBillResponse struct {
	model.WechatError
	RequestId   string // 唯一请求ID
	HashType    string `json:"hash_type,omitempty"`    // 哈希类型
	HashValue   string `json:"hash_value,omitempty"`   // 哈希值
	DownloadUrl string `json:"download_url,omitempty"` // 账单下载地址
}

// SharingOrder 分账账单
type SharingOrder struct {
	SubMchId      string      `json:"sub_mchid,omitempty"`      // 子商户号, 服务商平台返回
	TransactionId string      `json:"transaction_id,omitempty"` // 微信订单号
	OutOrderNo    string      `json:"out_order_no,omitempty"`   // 商户分账单号
	OrderId       string      `json:"order_id,omitempty"`       // 微信分账单号
	State         string      `json:"state,omitempty"`          // 分账单状态
	Receivers     []*Receiver `json:"receivers,omitempty"`      // 分账接收方
}

// ReturnOrder 分账回退账单
type ReturnOrder struct {
	SubMchId    string    `json:"sub_mchid,omitempty"`     // 子商户号: 服务商平台返回
	OrderId     string    `json:"order_id,omitempty"`      // 微信分账单号
	OutOrderNo  string    `json:"out_order_no,omitempty"`  // 商户分账单号
	OutReturnNo string    `json:"out_return_no,omitempty"` // 商户回退单号
	ReturnId    string    `json:"return_id,omitempty"`     // 微信回退单号
	ReturnMchId string    `json:"return_mchid,omitempty"`  // 回退商户号
	Amount      int64     `json:"amount,omitempty"`        // 回退金额
	Description string    `json:"description,omitempty"`   // 回退描述
	Result      string    `json:"result,omitempty"`        // 回退结果
	FailReason  string    `json:"fail_reason,omitempty"`   // 失败原因
	CreateTime  time.Time `json:"create_time,omitempty"`   // 创建时间
	FinishTime  time.Time `json:"finish_time,omitempty"`   // 完成时间
}

// NotifyReceiver 通知里的接收方
type NotifyReceiver struct {
	Amount      int64  `json:"amount,omitempty"`      // 分账金额
	Description string `json:"description,omitempty"` // 分账描述
	Type        string `json:"type,omitempty"`        // 分账接收方类型
	Account     string `json:"account,omitempty"`     // 分账接收方账号
}

// Receiver 分账接收方
type Receiver struct {
	Amount      int64     `json:"amount,omitempty"`      // 分账金额
	Description string    `json:"description,omitempty"` // 分账描述
	Type        string    `json:"type,omitempty"`        // 分账接收方类型
	Account     string    `json:"account,omitempty"`     // 分账接收方账号
	Result      string    `json:"result,omitempty"`      // 分账结果
	FailReason  string    `json:"fail_reason,omitempty"` // 分账失败原因
	DetailId    string    `json:"detail_id,omitempty"`   // 分账明细单号
	CreateTime  time.Time `json:"create_time,omitempty"` // 分账创建时间
	FinishTime  time.Time `json:"finish_time,omitempty"` // 分账完成时间
}

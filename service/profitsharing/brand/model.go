package brand

import (
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
	"github.com/pyihe/wechat-sdk/v3/service/profitsharing"
)

// CreateSharingResponse 请求分账应答参数
type CreateSharingResponse struct {
	model.WechatError
	RequestId     string                    // 唯一请求ID
	BrandMchId    string                    `json:"brand_mchid,omitempty"`    // 品牌主商户号
	SubMchId      string                    `json:"sub_mchid,omitempty"`      // 子商户号
	TransactionId string                    `json:"transaction_id,omitempty"` // 微信订单号
	OutOrderNo    string                    `json:"out_order_no,omitempty"`   // 商户分账单号
	OrderId       string                    `json:"order_id,omitempty"`       // 微信分账单号
	Status        string                    `json:"status,omitempty"`         // 分账单状态
	Receivers     []*profitsharing.Receiver `json:"receivers,omitempty"`      // 分账接收方列表
}

// QuerySharingResponse 查询分账结果应答参数
type QuerySharingResponse struct {
	model.WechatError
	RequestId         string                    // 唯一请求ID
	SubMchId          string                    `json:"sub_mchid,omitempty"`          // 子商户号
	TransactionId     string                    `json:"transaction_id,omitempty"`     // 微信订单号
	OutOrderNo        string                    `json:"out_order_no,omitempty"`       // 商户分账单号
	OrderId           string                    `json:"order_id,omitempty"`           // 微信分账单号
	Status            string                    `json:"status,omitempty"`             // 分账单状态
	Receivers         []*profitsharing.Receiver `json:"receivers,omitempty"`          // 分账接收方列表
	FinishAmount      int64                     `json:"finish_amount,omitempty"`      // 分账完结金额
	FinishDescription string                    `json:"finish_description,omitempty"` // 分账完结描述
}

// ReturnSharingResponse 请求分账回退应答参数
type ReturnSharingResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	ReturnOrder
}

// QueryReturnSharingRequest 查询分账回退结果请求参数
type QueryReturnSharingRequest struct {
	SubMchId    string `json:"sub_mchid,omitempty"`     // 子商户号
	OutReturnNo string `json:"out_return_no,omitempty"` // 商户回退单号
	OrderId     string `json:"order_id,omitempty"`      // 微信分账单号
	OutOrderNo  string `json:"out_order_no,omitempty"`  // 商户分账单号
}

// QueryReturnSharingResponse 查询分账回退结果应答参数
type QueryReturnSharingResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	ReturnOrder
}

// FinishSharingRequest 完结分账请求参数
type FinishSharingRequest struct {
	SubMchId      string `json:"sub_mchid"`      // 子商户号
	TransactionId string `json:"transaction_id"` // 微信订单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	Description   string `json:"description"`    // 分账描述
}

// FinishSharingResponse 完结分账应答参数
type FinishSharingResponse struct {
	model.WechatError
	RequestId     string // 唯一请求ID
	SubMchId      string `json:"sub_mchid,omitempty"`      // 子商户号
	TransactionId string `json:"transaction_id,omitempty"` // 微信订单号
	OutOrderNo    string `json:"out_order_no,omitempty"`   // 商户分账单号
	OrderId       string `json:"order_id,omitempty"`       // 微信分账单号
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
	RequestId  string // 唯一请求ID
	BrandMchId string `json:"brand_mchid,omitempty"` // 品牌主商户号
	MaxRatio   int32  `json:"max_ratio,omitempty"`   // 最大分账比例
}

// AddReceiverResponse 添加分账接收方
type AddReceiverResponse struct {
	model.WechatError
	RequestId  string // 唯一请求ID
	BrandMchId string `json:"brand_mchid,omitempty"` // 品牌主商户号
	Type       string `json:"type,omitempty"`        // 分账接收方类型
	Account    string `json:"account,omitempty"`     // 分账接收方账号
}

// DeleteReceiverResponse 删除分账接收方
type DeleteReceiverResponse struct {
	model.WechatError
	RequestId  string // 唯一请求ID
	BrandMchId string `json:"brand_mchid,omitempty"` // 品牌主商户号
	Type       string `json:"type,omitempty"`        // 分账接受方类型
	Account    string `json:"account,omitempty"`     // 分账接收方账号
}

// SharingNotifyResponse 分账动帐通知参数
type SharingNotifyResponse struct {
	NotifyId      string                        `json:"id,omitempty"`             // 唯一通知ID
	SpMchId       string                        `json:"sp_mchid,omitempty"`       // 服务商商户号: 服务商平台返回
	SubMchId      string                        `json:"sub_mchid,omitempty"`      // 子商户号: 服务商平台返回
	TransactionId string                        `json:"transaction_id,omitempty"` // 微信订单号
	OrderId       string                        `json:"order_id,omitempty"`       // 微信分账/回退单号
	OutOrderNo    string                        `json:"out_order_no,omitempty"`   // 商户分账/回退单号
	SuccessTime   time.Time                     `json:"success_time,omitempty"`   // 成功时间
	Receiver      *profitsharing.NotifyReceiver `json:"receiver,omitempty"`       // 分账接收方列表
}

// ReturnOrder 分账回退账单
type ReturnOrder struct {
	SubMchId    string    `json:"sub_mchid,omitempty"`     // 子商户号
	OrderId     string    `json:"order_id,omitempty"`      // 微信分账单号
	OutOrderNo  string    `json:"out_order_no,omitempty"`  // 商户分账单号
	OutReturnNo string    `json:"out_return_no,omitempty"` // 商户回退单号
	ReturnMchId string    `json:"return_mchid,omitempty"`  // 回退商户号
	Amount      int64     `json:"amount,omitempty"`        // 回退金额
	ReturnNo    string    `json:"return_no,omitempty"`     // 微信回退单号
	Result      string    `json:"result,omitempty"`        // 回退结果
	FailReason  string    `json:"fail_reason,omitempty"`   // 失败原因
	FinishTime  time.Time `json:"finish_time,omitempty"`   // 完成时间
}

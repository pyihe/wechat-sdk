package manage

import (
	"github.com/pyihe/wechat-sdk/model"
)

// PostPayments 后付费项目
type PostPayments struct {
	Name        string `json:"name,omitempty"`  // 付费项目名称
	Amount      int64  `json:"amount"`          // 金额
	Description string `json:"description"`     // 计费说明
	Count       uint   `json:"count,omitempty"` // 付费数量
}

// PostDiscounts 后付费商户优惠
type PostDiscounts struct {
	Name        string `json:"name,omitempty"`        // 优惠名称
	Description string `json:"description,omitempty"` // 优惠说明
	Amount      int64  `json:"amount"`                // 优惠金额
	Count       uint   `json:"count,omitempty"`       // 优惠数量
}

// RiskFund 订单风险金
type RiskFund struct {
	Name        string `json:"name,omitempty"`        // 风险金名称
	Amount      int64  `json:"amount,omitempty"`      // 风险金额
	Description string `json:"description,omitempty"` // 风险说明
}

// TimeRange 服务时间段
type TimeRange struct {
	StartTime       string `json:"start_time,omitempty"`        // 服务开始时间
	StartTimeRemark string `json:"start_time_remark,omitempty"` // 服务开始时间备注
	EndTime         string `json:"end_time,omitempty"`          // 服务结束时间
	EndTimeRemark   string `json:"end_time_remark,omitempty"`   // 服务结束时间备注
}

// Location 服务位置
type Location struct {
	StartLocation string `json:"start_location,omitempty"` // 服务开始地点
	EndLocation   string `json:"end_location,omitempty"`   // 服务结束位置
}

// Collection 收款信息
type Collection struct {
	State        string    `json:"state,omitempty"`         // 收款状态
	TotalAmount  int64     `json:"total_amount,omitempty"`  // 总收款金额
	PayingAmount int64     `json:"paying_amount,omitempty"` // 待收款金额
	PaidAmount   int64     `json:"paid_amount,omitempty"`   // 已收款金额
	Details      []*Detail `json:"details,omitempty"`       // 收款明细列表
}

// Detail 收款明细列表
type Detail struct {
	Seq             int                      `json:"seq,omitempty"`              // 收款序号
	Amount          int64                    `json:"amount,omitempty"`           // 单笔收款金额
	PaidType        string                   `json:"paid_type,omitempty"`        // 收款成功渠道
	PaidTime        string                   `json:"paid_time,omitempty"`        // 收款成功时间
	TransactionId   string                   `json:"transaction_id,omitempty"`   // 微信支付交易单号
	PromotionDetail []*model.PromotionDetail `json:"promotion_detail,omitempty"` // 优惠功能
}

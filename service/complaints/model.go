package complaints

import (
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
)

// QueryComplaintListRequest 查询投诉列表请求参数
type QueryComplaintListRequest struct {
	Limit            uint32 `json:"limit,omitempty"`             // 分页大小
	Offset           uint32 `json:"offset,omitempty"`            // 分页开始位置
	BeginDate        string `json:"begin_date,omitempty"`        // 开始日期
	EndDate          string `json:"end_date,omitempty"`          // 结束日期
	ComplaintedMchId string `json:"complainted_mchid,omitempty"` // 被投诉赏花红
}

type QueryComplaintListResponse struct {
	model.WechatError
	RequestId  string       // 唯一请求ID
	Limit      uint32       `json:"limit,omitempty"`       // 分页大小
	Offset     uint32       `json:"offset,omitempty"`      // 分页起始位置
	TotalCount int32        `json:"total_count,omitempty"` // 总条数
	Data       []*Complaint `json:"data,omitempty"`        // 用户投诉信息详情
}

// QueryComplaintDetailResponse 查询投诉单详情应答参数
type QueryComplaintDetailResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	Complaint
}

// QueryNegotiationHistoryResponse 查询投诉协商历史应答参数
type QueryNegotiationHistoryResponse struct {
	model.WechatError
	RequestId  string                // 唯一请求ID
	Data       []*NegotiationHistory `json:"data,omitempty"`        // 投诉协商历史
	Limit      uint32                `json:"limit,omitempty"`       // 分页大小
	Offset     uint32                `json:"offset,omitempty"`      // 分页开始位置
	TotalCount int32                 `json:"total_count,omitempty"` // 投诉协商历史总条数
}

// ComplaintNotifyResponse 解析投诉通知回调
type ComplaintNotifyResponse struct {
	NotifyId    string // 唯一通知ID
	ComplaintId string `json:"complaint_id,omitempty"` // 投诉单号
	ActionType  string `json:"action_type,omitempty"`  // 动作类型
}

// NotifyUrlResponse 创建投诉通知回调地址应答参数
type NotifyUrlResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	MchId     string `json:"mchid,omitempty"` // 商户号
	Url       string `json:"url,omitempty"`   // 通知地址
}

// DeleteNotifyUrlResponse 删除投诉回调通知地址应答参数
type DeleteNotifyUrlResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

// CommitResponse 提交回复应答参数
type CommitResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

// CompleteResponse 反馈处理完成应答参数
type CompleteResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

// UploadImageResponse 图片上传应答参数
type UploadImageResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	MediaId   string `json:"media_id,omitempty"` // 媒体文件标识
}

// NegotiationHistory 协商历史
type NegotiationHistory struct {
	ComplaintMediaList *ComplaintMedia `json:"complaint_media,omitempty"` // 投诉资料列表
	LogId              string          `json:"log_id,omitempty"`          // 操作流水号
	Operator           string          `json:"operator,omitempty"`        // 操作人
	OperateTime        time.Time       `json:"operate_time,omitempty"`    // 操作时间
	OperateType        string          `json:"operate_type,omitempty"`    // 操作类型
	OperateDetails     string          `json:"operate_details,omitempty"` // 操作内容
	ImageList          []string        `json:"image_list,omitempty"`      // 图片凭证
}

// Complaint 单条投诉信息
type Complaint struct {
	ComplaintId           string            `json:"complaint_id,omitempty"`            // 投诉单号
	ComplaintTime         time.Time         `json:"complaint_time,omitempty"`          // 投诉时间
	ComplaintDetail       string            `json:"complaint_detail,omitempty"`        // 投诉详情
	ComplaintState        string            `json:"complaint_state,omitempty"`         // 投诉单状态
	ComplaintMchId        string            `json:"complaint_mchid,omitempty"`         // 被投诉商户号
	PayerPhone            string            `json:"payer_phone,omitempty"`             // 投诉着联系方式
	PayerOpenId           string            `json:"payer_openid,omitempty"`            // 投诉人openid
	ComplaintMediaList    []*ComplaintMedia `json:"complaint_media_list,omitempty"`    // 投诉资料列表
	ComplaintOrderInfo    []*ComplaintOrder `json:"complaint_order_info,omitempty"`    // 投诉单关联订单信息
	ComplaintFullRefunded bool              `json:"complaint_full_refunded,omitempty"` // 投诉单是否已经全额退款
	ProblemDescription    string            `json:"problem_description,omitempty"`     // 问题描述
	IncomingUserResponse  bool              `json:"incoming_user_response,omitempty"`  // 是否有待回复的用户留言
	UserComplaintTimes    int32             `json:"user_complaint_times,omitempty"`    // 用户投诉次数
}

// ComplaintMedia 投诉资料列表
type ComplaintMedia struct {
	MediaType string `json:"media_type,omitempty"` // 媒体文件业务类型
	MediaUrl  string `json:"media_url,omitempty"`  // 媒体文件请求URL
}

// ComplaintOrder 投诉关联订单信息
type ComplaintOrder struct {
	TransactionId string `json:"transaction_id,omitempty"` // 微信订单号
	OutTradeNo    string `json:"out_trade_no,omitempty"`   // 商户订单号
	Amount        int64  `json:"amount,omitempty"`         // 订单金额
}

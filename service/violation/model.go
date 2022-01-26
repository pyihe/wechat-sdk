package violation

import (
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
)

// NotificationResponse 创建商户违规通知回调地址
type NotificationResponse struct {
	model.WechatError
	RequestId string // 请求唯一ID
	NotifyUrl string `json:"notify_url,omitempty"` // 通知地址
}

// DeleteNotificationResponse 删除商户违规通知回调地址应答参数
type DeleteNotificationResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

// NotifyResponse 商户平台处置记录回调通知参数
type NotifyResponse struct {
	NotifyId          string    // 唯一通知ID
	SubMchid          string    `json:"sub_mchid,omitempty"`          // 子商户号
	CompanyName       string    `json:"company_name,omitempty"`       // 子商户公司名称
	RecordId          string    `json:"record_id,omitempty"`          // 商户违约处理通知ID
	PunishPlan        string    `json:"punish_plan,omitempty"`        // 处罚方案
	PunishTime        time.Time `json:"punish_time,omitempty"`        // 处罚时间
	PunishDescription string    `json:"punish_description,omitempty"` // 处罚方案描述信息
	RiskType          string    `json:"risk_type,omitempty"`          // 风险类型
	RiskDescription   string    `json:"risk_description,omitempty"`   // 风险描述
}

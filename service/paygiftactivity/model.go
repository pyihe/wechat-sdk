package paygiftactivity

import (
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
)

// CreateActivityResponse 创建全场满额送活动应答参数
type CreateActivityResponse struct {
	model.WechatError
	RequestId  string    // 唯一请求ID
	ActivityId string    `json:"activity_id,omitempty"` // 活动ID
	CreateTime time.Time `json:"create_time,omitempty"` // 创建时间
}

// QueryActivityResponse 查询活动详情应答参数
type QueryActivityResponse struct {
	model.WechatError
	RequestId         string            // 唯一请求ID
	ActivityId        string            `json:"activity_id,omitempty"`         // 活动ID
	ActivityType      string            `json:"activity_type,omitempty"`       // 活动类型
	ActivityBaseInfo  *ActivityBaseInfo `json:"activity_base_info,omitempty"`  // 活动基本信息
	AwardSendRule     *AwardSendRule    `json:"award_send_rule,omitempty"`     // 活动奖品发放规则
	AdvancedSetting   *AdvancedSetting  `json:"advanced_setting,omitempty"`    // 活动高级设置
	ActivityStatus    string            `json:"activity_status,omitempty"`     // 活动状态
	CreatorMerchantId string            `json:"creator_merchant_id,omitempty"` // 创建商户号
	BelongMerchantId  string            `json:"belong_merchant_id,omitempty"`  // 所属商户号
	PauseTime         time.Time         `json:"pause_time,omitempty"`          // 活动暂停时间
	RecoveryTime      time.Time         `json:"recovery_time,omitempty"`       // 活动恢复时间
	CreateTime        time.Time         `json:"create_time,omitempty"`         // 活动创建时间
	UpdateTime        time.Time         `json:"update_time,omitempty"`         // 活动创建时间
}

// QueryActivityMerchantResponse 查询活动发券商户号应答参数
type QueryActivityMerchantResponse struct {
	model.WechatError
	RequestId  string              // 唯一请求ID
	Data       []*ActivityMerchant `json:"data,omitempty"`        // 结果集
	TotalCount uint32              `json:"total_count,omitempty"` // 总数
	Offset     uint32              `json:"offset,omitempty"`      // 分页页码
	Limit      uint32              `json:"limit,omitempty"`       // 分页大小
	ActivityId string              `json:"activity_id,omitempty"` // 活动ID
}

// QueryActivityGoodsResponse 查询互动指定商品列表应答参数
type QueryActivityGoodsResponse struct {
	model.WechatError
	RequestId  string           // 唯一请求ID
	Data       []*ActivityGoods `json:"data,omitempty"`        // 结果集
	TotalCount uint32           `json:"total_count,omitempty"` // 总数
	Offset     uint32           `json:"offset,omitempty"`      // 分页页码
	Limit      uint32           `json:"limit,omitempty"`       // 分页大小
	ActivityId string           `json:"activity_id,omitempty"` // 活动ID
}

// TerminateActivityResponse 终止活动应答参数
type TerminateActivityResponse struct {
	model.WechatError
	RequestId     string    // 唯一请求ID
	TerminateTime time.Time `json:"terminate_time,omitempty"` // 活动生效时间
	ActivityId    string    `json:"activity_id,omitempty"`    // 活动ID
}

// AddActivityMerchantRequest 新增活动发券商户号
type AddActivityMerchantRequest struct {
	MerchantIdList []string `json:"merchant_id_list,omitempty"` // 发券商户号
	AddRequestNo   string   `json:"add_request_no"`             // 请求业务单据号
}

// AddActivityMerchantResponse 添加活动商户号应答参数
type AddActivityMerchantResponse struct {
	model.WechatError
	RequestId             string               // 唯一请求ID
	ActivityId            string               `json:"activity_id,omitempty"`              // 活动ID
	InvalidMerchantIdList []*InvalidMerchantId `json:"invalid_merchant_id_list,omitempty"` // 校验失败的发券商户号列表
	AddTime               time.Time            `json:"add_time,omitempty"`                 // 添加时间
}

// QueryActivityFilter 根据过滤条件查询有礼活动列表
type QueryActivityFilter struct {
	Offset         uint32 `json:"offset"`                    // 分页页码
	Limit          uint32 `json:"limit"`                     // 页码大小
	ActivityName   string `json:"activity_name,omitempty"`   // 活动名称
	ActivityStatus string `json:"activity_status,omitempty"` // 活动状态
	AwardType      string `json:"award_type,omitempty"`      // 奖品类型
}

// QueryActivityFilterResponse 根据过滤条件查询有礼活动列表
type QueryActivityFilterResponse struct {
	model.WechatError
	RequestId  string            // 唯一请求ID
	Data       []*FilterActivity `json:"data,omitempty"`        // 结果集
	TotalCount uint32            `json:"total_count,omitempty"` // 总数
	Offset     uint32            `json:"offset,omitempty"`      // 分页页码
	Limit      uint32            `json:"limit,omitempty"`       // 分页大小
}

// DeleteActivityMerchantRequest 删除活动发券商户号
type DeleteActivityMerchantRequest struct {
	MerchantIdList  []string `json:"merchant_id_list,omitempty"`  // 杀出的发券商户号
	DeleteRequestNo string   `json:"delete_request_no,omitempty"` // 请求业务单据号
}

// DeleteActivityMerchantResponse 删除活动发券商户号应答参数
type DeleteActivityMerchantResponse struct {
	model.WechatError
	RequestId  string    // 唯一请求ID
	ActivityId string    `json:"activity_id,omitempty"` // 活动ID
	DeleteTime time.Time `json:"delete_time,omitempty"` // 删除时间
}

type FilterActivity struct {
	ActivityId        string            `json:"activity_id,omitempty"`         // 活动ID
	ActivityType      string            `json:"activity_type,omitempty"`       // 活动类型
	ActivityBaseInfo  *ActivityBaseInfo `json:"activity_base_info,omitempty"`  // 活动基本信息
	AwardSendRule     *AwardSendRule    `json:"award_send_rule,omitempty"`     // 奖品发放规则
	AdvancedSetting   *AdvancedSetting  `json:"advanced_setting,omitempty"`    // 活动高级设置
	CreatorMerchantId string            `json:"creator_merchant_id,omitempty"` // 创建商户号
	BelongMerchantId  string            `json:"belong_merchant_id,omitempty"`  // 所属商户号
	CreateTime        time.Time         `json:"create_time,omitempty"`         // 活动创建时间
	UpdateTime        time.Time         `json:"update_time,omitempty"`         // 活动创建时间
}

// InvalidMerchantId 校验失败的发券商户号
type InvalidMerchantId struct {
	MchId         string `json:"mchid,omitempty"`          // 商户号
	InvalidReason string `json:"invalid_reason,omitempty"` // 无效原因
}

// ActivityGoods 多动商品
type ActivityGoods struct {
	GoodsId    string    `json:"goods_id,omitempty"`    // 商品ID
	CreateTime time.Time `json:"create_time,omitempty"` // 创建时间
	UpdateTime time.Time `json:"update_time,omitempty"` // 更新时间
}

// ActivityMerchant 活动商户
type ActivityMerchant struct {
	MchId        string    `json:"mchid,omitempty"`         // 商户号
	MerchantName string    `json:"merchant_name,omitempty"` // 商户名称
	CreateTime   time.Time `json:"create_time,omitempty"`   // 创建时间
	UpdateTime   time.Time `json:"update_time,omitempty"`   // 更新时间

}

// ActivityBaseInfo 活动基本信息
type ActivityBaseInfo struct {
	ActivityName        string            `json:"activity_name,omitempty"`         // 活动名称
	ActivitySecondTitle string            `json:"activity_second_title,omitempty"` // 活动副标题
	MerchantLogoUrl     string            `json:"merchant_logo_url,omitempty"`     // 商户Logo
	BackgroundColor     string            `json:"background_color,omitempty"`      // 背景颜色
	BeginTime           time.Time         `json:"begin_time,omitempty"`            // 活动开始时间
	EndTime             time.Time         `json:"end_time,omitempty"`              // 活动结束时间
	AvailablePeriods    *AvailablePeriods `json:"available_periods,omitempty"`     // 可用时间段
	OutRequestNo        string            `json:"out_request_no,omitempty"`        // 商户请求单号
	DeliveryPurpose     string            `json:"delivery_purpose,omitempty"`      // 投放目的
	MiniProgramsAppId   string            `json:"mini_programs_appid,omitempty"`   // 商家小程序appid
	MiniProgramsPath    string            `json:"mini_programs_path,omitempty"`    // 商家小程序path
}

// AvailablePeriods 可用时间段
type AvailablePeriods struct {
	AvailableTime    []*AvailableTime    `json:"available_time,omitempty"`     // 可用时间
	AvailableDayTime []*AvailableDayTime `json:"available_day_time,omitempty"` // 每日可用时间
}

// AvailableTime 可用时间
type AvailableTime struct {
	BeginTime time.Time `json:"begin_time,omitempty"` // 可用开始时间
	EndTime   time.Time `json:"end_time,omitempty"`   // 可用结束时间
}

// AvailableDayTime 每日可用时间
type AvailableDayTime struct {
	BeginDayTime string `json:"begin_day_time,omitempty"` // 每日可用开始时间
	EndDayTime   string `json:"end_day_time,omitempty"`   // 每日可用结束时间
}

// AwardSendRule 活动奖品发放规则
type AwardSendRule struct {
	FullSendRule *FullSendRule `json:"full_send_rule,omitempty"` // 满送活动奖品发放规则
}

// FullSendRule 满送活动奖品发放规则
type FullSendRule struct {
	TransactionAmountMinimum int32    `json:"transaction_amount_minimum,omitempty"` // 消费金额门槛
	SendContent              string   `json:"send_content,omitempty"`               // 发放内容
	AwardType                string   `json:"award_type,omitempty"`                 // 奖品类型
	AwardList                []*Award `json:"award_list,omitempty"`                 // 奖品基本信息列表
	MerchantOption           string   `json:"merchant_option,omitempty"`            // 发券商户号选项
	MerchantIdList           []string `json:"merchant_id_list,omitempty"`           // 发券商户号
}

// Award 奖品基本信息列表
type Award struct {
	StockId          string `json:"stock_id,omitempty"`           // 批次ID
	OriginalImageUrl string `json:"original_image_url,omitempty"` // 奖品原始图(大图)
	ThumbnailUrl     string `json:"thumbnail_url,omitempty"`      // 奖品缩略图
}

// AdvancedSetting 活动高级设置
type AdvancedSetting struct {
	DeliveryUserCategory string   `json:"delivery_user_category,omitempty"` // 投放用户类别
	MerchantMemberAppId  string   `json:"merchant_member_appid,omitempty"`  // 商家会员appid
	GoodsTags            []string `json:"goods_tags,omitempty"`             // 订单优惠标记
}

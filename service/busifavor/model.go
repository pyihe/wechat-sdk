package busifavor

import (
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
)

// CreateStockResponse 创建商家券应答参数
type CreateStockResponse struct {
	model.WechatError
	RequestId  string    `json:"-"`                     // 唯一请求ID
	StockId    string    `json:"stock_id,omitempty"`    // 微信为每个商家券批次分配的唯一ID
	CreateTime time.Time `json:"create_time,omitempty"` // 创建时间
}

// QueryMerchantStockResponse 查询商家券详情应答参数
type QueryMerchantStockResponse struct {
	RequestId string
	model.WechatError
	MerchantStock
}

// UseCouponResponse 核销用户券API应答参数
type UseCouponResponse struct {
	model.WechatError
	RequestId        string
	StockId          string    `json:"stock_id,omitempty"`           // 批次号
	OpenId           string    `json:"openid,omitempty"`             // 用户标识
	WechatpayUseTime time.Time `json:"wechatpay_use_time,omitempty"` // 系统核销券成功的时间
}

// QueryUserCouponsByFilterRequest 根据过滤条件查询用户券
type QueryUserCouponsByFilterRequest struct {
	AppId           string // 公众账号ID
	StockId         string // 批次号
	CouponState     string // 券状态
	CreatorMerchant string // 创建批次的商户号
	BelongMerchant  string // 批次归属商户号
	SenderMerchant  string // 批次发放商户号
	Offset          uint32 // 分页页码
	Limit           uint32 // 分页大小
}

// QueryUserCouponsByFilterResponse 根据过滤条件查询用户券API应答参数
type QueryUserCouponsByFilterResponse struct {
	model.WechatError
	RequestId  string
	Data       []*UserCoupon `json:"data,omitempty"`        // 结果集
	TotalCount int32         `json:"total_count,omitempty"` // 总数量
	Offset     uint32        `json:"offset,omitempty"`      // 分页页码
	Limit      uint32        `json:"limit,omitempty"`       // 分页大小
}

// QueryUserCouponResponse 查询用户单张券详情应答参数
type QueryUserCouponResponse struct {
	model.WechatError
	RequestId string
	UserCoupon
}

// UploadCouponCodeResponse 上传预存code应答参数
type UploadCouponCodeResponse struct {
	model.WechatError
	RequestId      string
	StockId        string      `json:"stock_id,omitempty"`        // 批次号
	TotalCount     uint64      `json:"total_count,omitempty"`     // 去重后上传code总数
	SuccessCount   uint64      `json:"success_count,omitempty"`   // 上传成功code个数
	SuccessCodes   []string    `json:"success_codes,omitempty"`   // 上传成功的code列表
	SuccessTime    time.Time   `json:"success_time,omitempty"`    // 上传成功时间
	FailCount      uint64      `json:"fail_count,omitempty"`      // 上传失败code个数
	FailCodes      []*FailCode `json:"fail_codes,omitempty"`      // 上传失败的code及原因
	ExistCodes     []string    `json:"exist_codes,omitempty"`     // 已经存在的code列表
	DuplicateCodes []string    `json:"duplicate_codes,omitempty"` // 本次请求中重复的code列表
}

// SetCallbackRequest 设置商家券事件通知地址
type SetCallbackRequest struct {
	MchId     string `json:"mchid,omitempty"` // 商户号
	NotifyUrl string `json:"notify_url"`      // 通知URL地址
}

// SetCallbackResponse 设置商家券事件通知地址API应答参数
type SetCallbackResponse struct {
	model.WechatError
	RequestId  string    // 唯一请求ID
	UpdateTime time.Time `json:"update_time,omitempty"` // 修改时间
	NotifyUrl  string    `json:"notify_url,omitempty"`  // 通知URL地址
	MchId      string    `json:"mchid,omitempty"`       // 商户号
}

type QueryCallbacksResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	NotifyUrl string `json:"notify_url,omitempty"` // 通知URL地址
	MchId     string `json:"mchid,omitempty"`      // 商户号
}

// AssociateRequest 关联订单信息API请求参数
type AssociateRequest struct {
	StockId      string `json:"stock_id"`       // 批次号
	CouponCode   string `json:"coupon_code"`    // 券code
	OutTradeNo   string `json:"out_trade_no"`   // 关联的商户订单号
	OutRequestNo string `json:"out_request_no"` // 商户请求单号
}

// AssociateResponse 关联订单信息API应答参数
type AssociateResponse struct {
	model.WechatError
	RequestId              string    // 唯一请求ID
	WechatpayAssociateTime time.Time `json:"wechatpay_associate_time,omitempty"` // 关联成功时间
}

// DisassociateRequest 取消关联订单信息APi请求参数
type DisassociateRequest struct {
	StockId      string `json:"stock_id"`       // 批次号
	CouponCode   string `json:"coupon_code"`    // 券code
	OutTradeNo   string `json:"out_trade_no"`   // 关联的商户订单号
	OutRequestNo string `json:"out_request_no"` // 商户请求单号
}

// DisAssociateResponse 取消关联订单信息API应答参数
type DisAssociateResponse struct {
	model.WechatError
	RequestId                 string    // 唯一请求ID
	WechatpayDisassociateTime time.Time `json:"wechatpay_disassociate_time,omitempty"` // 取消关联成功的时间
}

// ModifyBudgetResponse 修改批次预算
type ModifyBudgetResponse struct {
	model.WechatError
	RequestId       string // 唯一请求ID
	MaxCoupons      int32  `json:"max_coupons,omitempty"`        // 批次最大发放个数
	MaxCouponsByDay int32  `json:"max_coupons_by_day,omitempty"` // 当前单天发放上限个数
}

// ModifyStockResponse 修改商家券基本信息
type ModifyStockResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

// ReturnRequest 申请退券API请求参数
type ReturnRequest struct {
	CouponCode      string `json:"coupon_code"`       // 券code
	StockId         string `json:"stock_id"`          // 批次号
	ReturnRequestNo string `json:"return_request_no"` // 退券请求单据号
}

// ReturnResponse 申请退券API应答参数
type ReturnResponse struct {
	model.WechatError
	RequestId           string    // 唯一请求ID
	WechatpayReturnTime time.Time `json:"wechatpay_return_time,omitempty"` // 微信退券成功的时间
}

// DeactivateRequest 使券失效
type DeactivateRequest struct {
	CouponCode          string `json:"coupon_code"`                 // 券code
	StockId             string `json:"stock_id"`                    // 批次号
	DeactivateRequestNo string `json:"deactivate_request_no"`       // 失效请求单据号
	DeactivateReason    string `json:"deactivate_reason,omitempty"` // 失效原因
}

// DeactivateResponse 使券失效API应答参数
type DeactivateResponse struct {
	model.WechatError
	RequestId               string    // 唯一请求ID
	WechatpayDeactivateTime time.Time `json:"wechatpay_deactivate_time,omitempty"` // 券成功失效的时间
}

// SubsidyPayResponse 补差付款API应答参数
type SubsidyPayResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	PayReceipt
}

// QueryPayReceiptResponse 查询补差付款单详情应答参数
type QueryPayReceiptResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	PayReceipt
}

// SendCouponRequest 发放消费卡请求参数
type SendCouponRequest struct {
	AppId        string `json:"appid"`          // 消费卡所属appid
	OpenId       string `json:"open_id"`        // 用户openid
	OutRequestNo string `json:"out_request_no"` // 商户单据号
	SendTime     string `json:"send_time"`      // 请求发卡时间
}

// SendCouponResponse 发放消费卡应答参数
type SendCouponResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	CardCode  string `json:"card_code,omitempty"` // 消费卡code
}

// ReceiveCouponResponse 领券事件通知参数
type ReceiveCouponResponse struct {
	NotifyId     string      // 唯一通知ID
	EventType    string      `json:"event_type,omitempty"`    // 事件类型
	CouponCode   string      `json:"coupon_code,omitempty"`   // 券code
	StockId      string      `json:"stock_id,omitempty"`      // 批次号
	SendTime     time.Time   `json:"send_time,omitempty"`     // 发放时间
	OpenId       string      `json:"openid,omitempty"`        // 用户标识
	SendChannel  string      `json:"send_channel,omitempty"`  // 发放渠道
	SendMerchant string      `json:"send_merchant,omitempty"` // 发券商户号
	AttachInfo   *AttachInfo `json:"attach_info,omitempty"`   // 附加信息
}

// AttachInfo 领券回调通知附加信息
type AttachInfo struct {
	TransactionId   string `json:"transaction_id,omitempty"`    // 交易订单编号
	ActCode         string `json:"act_code,omitempty"`          // 支付有礼活动编号或者营销馆活动ID
	HallCode        string `json:"hall_code,omitempty"`         // 营销馆ID
	HallBelongMchId int32  `json:"hall_belong_mchid,omitempty"` // 营销馆所属商户号
	CardId          string `json:"card_id,omitempty"`           // 会员卡ID
	Code            string `json:"code,omitempty"`              // 会员卡code
	ActivityId      string `json:"activity_id,omitempty"`       // 会员活动ID
}

// PayReceipt 补差付款票据
type PayReceipt struct {
	SubsidyReceiptId string    `json:"subsidy_receipt_id,omitempty"` // 补差付款单号
	StockId          string    `json:"stock_id,omitempty"`           // 商家券批次号
	CouponCode       string    `json:"coupon_code,omitempty"`        // 券code
	TransactionId    string    `json:"transaction_id,omitempty"`     // 微信支付订单号
	PayerMerchant    string    `json:"payer_merchant,omitempty"`     // 营销补差扣款商户号
	PayeeMerchant    string    `json:"payee_merchant,omitempty"`     // 营销补差入账商户号
	Amount           int64     `json:"amount,omitempty"`             // 补差付款金额
	Description      string    `json:"description,omitempty"`        // 补差付款描述
	Status           string    `json:"status,omitempty"`             // 补差付款单据状态
	FailReason       string    `json:"fail_reason,omitempty"`        // 补差付款失败原因
	SuccessTime      time.Time `json:"success_time,omitempty"`       // 补差付款成功时间
	OutSubsidyNo     string    `json:"out_subsidy_no,omitempty"`     // 业务请求唯一单号
	CreateTime       time.Time `json:"create_time,omitempty"`        // 补差付款发起时间
}

// FailCode code上传失败参数
type FailCode struct {
	CouponCode string `json:"coupon_code,omitempty"` // 上传失败的券code
	Code       string `json:"code,omitempty"`        // 上传失败错误码
	Message    string `json:"message,omitempty"`     // 上传失败错误信息
}

// UserCoupon 用户券
type UserCoupon struct {
	BelongMerchant      string              `json:"belong_merchant,omitempty"`        // 批次归属商户号
	StockName           string              `json:"stock_name,omitempty"`             // 商家券批次名称
	Comment             string              `json:"comment,omitempty"`                // 批次备注
	GoodsName           string              `json:"goods_name,omitempty"`             // 适用商品范围
	StockType           string              `json:"stock_type,omitempty"`             // 批次类型
	Transferable        bool                `json:"transferable,omitempty"`           // 是否允许转赠
	Shareable           bool                `json:"shareable,omitempty"`              // 是否允许分享领券链接
	CouponState         string              `json:"coupon_state,omitempty"`           // 券状态
	DisplayPatternInfo  *DisplayPatternInfo `json:"display_pattern_info,omitempty"`   // 样式信息
	CouponUseRule       *CouponUseRule      `json:"coupon_use_rule,omitempty"`        // 券核销规则
	CustomEntrance      *CustomEntrance     `json:"custom_entrance,omitempty"`        // 自定义入口
	CouponCode          string              `json:"coupon_code,omitempty"`            // 券code
	StockId             string              `json:"stock_id,omitempty"`               // 批次号
	AvailableStartTime  time.Time           `json:"available_start_time,omitempty"`   // 券可使用开始时间
	ExpireTime          time.Time           `json:"expire_time,omitempty"`            // 券过期时间
	ReceiveTime         time.Time           `json:"receive_time,omitempty"`           // 领券时间
	SendRequestNo       string              `json:"send_request_no,omitempty"`        // 发券请求单号
	UseRequestNo        string              `json:"use_request_no,omitempty"`         // 核销请求单号
	AssociateOutTradeNo string              `json:"associate_out_trade_no,omitempty"` // 关联的商户订单号
	UseTime             time.Time           `json:"use_time,omitempty"`               // 券核销时间
}

// MerchantStock 商家券
type MerchantStock struct {
	StockName            string                `json:"stock_name,omitempty"`             // 商家券批次名称
	BelongMerchant       string                `json:"belong_merchant,omitempty"`        // 批次归属商户号
	Comment              string                `json:"comment,omitempty"`                // 批次备注
	GoodsName            string                `json:"goods_name,omitempty"`             // 适用商品范围
	StockType            string                `json:"stock_type,omitempty"`             // 批次类型
	CouponUseRule        *CouponUseRule        `json:"coupon_use_rule,omitempty"`        // 核销规则
	StockSendRule        *StockSendRule        `json:"stock_send_rule,omitempty"`        // 发放规则
	CustomEntrance       *CustomEntrance       `json:"custom_entrance,omitempty"`        // 自定义入口
	DisplayPatternInfo   *DisplayPatternInfo   `json:"display_pattern_info,omitempty"`   // 样式信息
	StockState           string                `json:"stock_state,omitempty"`            // 批次状态
	CouponCodeMode       string                `json:"coupon_code_mode,omitempty"`       // 券code模式
	StockId              string                `json:"stock_id,omitempty"`               // 批次号
	CouponCodeCount      *CouponCodeCount      `json:"coupon_code_count,omitempty"`      // 券code数量
	NotifyConfig         *NotifyConfig         `json:"notify_config,omitempty"`          // 事件通知配置
	SendCountInformation *SendCountInformation `json:"send_count_information,omitempty"` // 批次发放情况
}

// CouponUseRule 核销规则
type CouponUseRule struct {
	CouponAvailableTime *CouponAvailableTime `json:"coupon_available_time,omitempty"` // 券可核销时间
	FixedNormalCoupon   *FixedNormalCoupon   `json:"fixed_normal_coupon,omitempty"`   // 固定面额满减券使用规则
	DiscountCoupon      *DiscountCoupon      `json:"discount_coupon,omitempty"`       // 折扣券使用规则
	ExchangeCoupon      *ExchangeCoupon      `json:"exchange_coupon,omitempty"`       // 换购券使用规则
	UseMethod           string               `json:"use_method,omitempty"`            // 核销方式
	MiniProgramsAppId   string               `json:"mini_programs_appid,omitempty"`   // 小程序Appid
	MiniProgramsPath    string               `json:"mini_programs_path,omitempty"`    // 小程序path
}

// CouponAvailableTime 券可核销时间
type CouponAvailableTime struct {
	AvailableBeginTime       time.Time                  `json:"available_begin_time,omitempty"`        // 开始时间
	AvailableEndTime         time.Time                  `json:"available_end_time,omitempty"`          // 结束时间
	AvailableDayAfterReceive int32                      `json:"available_day_after_receive,omitempty"` // 生效后N天内有效
	AvailableWeek            *AvailableWeek             `json:"available_week,omitempty"`              // 固定周期有效时间段
	IrregularyAvaliableTime  []*IrregularyAvaliableTime `json:"irregulary_avaliable_time,omitempty"`   // 无规律的有效时间段
	WaitDaysAfterReceive     int32                      `json:"wait_days_after_receive,omitempty"`     // 领取后N天开始生效
}

// AvailableWeek 固定周期有效时间段
type AvailableWeek struct {
	WeekDay          []time.Weekday      `json:"week_day,omitempty"`           // 可用星期数
	AvailableDayTime []*AvailableDayTime `json:"available_day_time,omitempty"` // 当天可用时间段
}

// AvailableDayTime 当天可用时间段
type AvailableDayTime struct {
	BeginTime int64 `json:"begin_time,omitempty"` // 当天可用开始时间
	EndTime   int64 `json:"end_time,omitempty"`   // 当天可用结束时间
}

// IrregularyAvaliableTime 无规律的有效时间段
type IrregularyAvaliableTime struct {
	BeginTime time.Time `json:"begin_time,omitempty"` // 开始时间
	EndTime   time.Time `json:"end_time,omitempty"`   // 结束时间
}

// FixedNormalCoupon 固定面额满减券使用规则
type FixedNormalCoupon struct {
	DiscountAmount     int64 `json:"discount_amount,omitempty"`     // 优惠金额
	TransactionMinimum int64 `json:"transaction_minimum,omitempty"` // 消费门槛
}

// DiscountCoupon 折扣券使用规则
type DiscountCoupon struct {
	DiscountPercent    int32 `json:"discount_percent,omitempty"`    // 折扣百分比
	TransactionMinimum int64 `json:"transaction_minimum,omitempty"` // 消费门槛
}

// ExchangeCoupon 换购券使用规则
type ExchangeCoupon struct {
	ExchangePrice      int64 `json:"exchange_price,omitempty"`      // 单品换购价
	TransactionMinimum int64 `json:"transaction_minimum,omitempty"` // 消费门槛
}

// StockSendRule 发放规则
type StockSendRule struct {
	MaxAmount          int64 `json:"max_amount,omitempty"`           // 批次总预算
	MaxCoupons         int32 `json:"max_coupons,omitempty"`          // 批次最大发放个数
	MaxCouponsPerUser  int32 `json:"max_coupons_per_user,omitempty"` // 用户最大可领个数
	MaxAmountByDay     int32 `json:"max_amount_by_day,omitempty"`    // 单日发放上限金额
	MaxCouponsByDay    int32 `json:"max_coupons_by_day,omitempty"`   // 单日发放上限个数
	NaturalPersonLimit bool  `json:"natural_person_limit,omitempty"` // 是否开启自然人限制
	PreventApiAbuse    bool  `json:"prevent_api_abuse,omitempty"`    // 可疑账号拦截
	Transferable       bool  `json:"transferable,omitempty"`         // 是否允许转赠
	Shareable          bool  `json:"shareable,omitempty"`            // 是否允许分享链接
}

// CustomEntrance 自定义入口
type CustomEntrance struct {
	MiniProgramsInfo *MiniProgramsInfo `json:"mini_programs_info,omitempty"` // 小程序入口
	AppId            string            `json:"appid,omitempty"`              // 商户公众号appid
	HallId           string            `json:"hall_id,omitempty"`            // 营销馆ID
	StoreId          string            `json:"store_id,omitempty"`           // 可用门店ID
	CodeDisplayMode  string            `json:"code_display_mode,omitempty"`  // code展示模式
}

// MiniProgramsInfo 小程序入口
type MiniProgramsInfo struct {
	MiniProgramsAppId string `json:"mini_programs_appid,omitempty"` // 商家小程序appid
	MiniProgramsPath  string `json:"mini_programs_path,omitempty"`  // 商家小程序path
	EntranceWords     string `json:"entrance_words,omitempty"`      // 入口文案
	GuidingWords      string `json:"guiding_words,omitempty"`       // 引导文案
}

// DisplayPatternInfo 样式信息
type DisplayPatternInfo struct {
	Description     string      `json:"description,omitempty"`       // 使用须知
	MerchantLogoUrl string      `json:"merchant_logo_url,omitempty"` // 商户logo
	MerchantName    string      `json:"merchant_name,omitempty"`     // 商户名称
	BackgroundColor string      `json:"background_color,omitempty"`  // 背景颜色
	CouponImageUrl  string      `json:"coupon_image_url,omitempty"`  // 券详情图片
	FinderInfo      *FinderInfo `json:"finder_info,omitempty"`       // 视频号相关信息
}

// FinderInfo 视频号相关信息
type FinderInfo struct {
	FinderId                 string `json:"finder_id,omitempty"`                    // 视频号ID
	FinderVideoId            string `json:"finder_video_id,omitempty"`              // 视频号视频ID
	FinderVideoCoverImageUrl string `json:"finder_video_cover_image_url,omitempty"` // 视频号封面图
}

// CouponCodeCount 券code数量
type CouponCodeCount struct {
	TotalCount     uint64 `json:"total_count,omitempty"`     // 该批次总共已上传的code总数
	AvailableCount uint64 `json:"available_count,omitempty"` // 该批次当前可用的code数
}

// NotifyConfig 事件通知配置
type NotifyConfig struct {
	NotifyAppId string `json:"notify_appid,omitempty"` // 事件通知appid
}

// SendCountInformation 批次发放情况
type SendCountInformation struct {
	TotalSendNum    uint64 `json:"total_send_num,omitempty"`    // 已发放券张数
	TotalSendAmount uint64 `json:"total_send_amount,omitempty"` // 已发放该券金额
	TodaySendNum    uint64 `json:"today_send_num,omitempty"`    // 单天已发放券张数
	TodaySendAmount uint64 `json:"today_send_amount,omitempty"` // 单天已发放券金额
}

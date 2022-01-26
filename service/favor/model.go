package favor

import (
	"encoding/json"
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
)

// CreateStockResponse 创建代金券批次API应答参数
type CreateStockResponse struct {
	model.WechatError
	RequestId  string    `json:"-"`                     // 唯一请求ID
	StockId    string    `json:"stock_id,omitempty"`    // 批次号
	CreateTime time.Time `json:"create_time,omitempty"` // 创建时间
}

// StartStockResponse 激活代金券批次API应答参数
type StartStockResponse struct {
	model.WechatError
	RequestId string    `json:"-"`                    // 唯一请求ID
	StockId   string    `json:"stock_id,omitempty"`   // 批次号
	StartTime time.Time `json:"start_time,omitempty"` // 生效时间
}

// SendCouponsRequest 发放代金券批次API请求参数
type SendCouponsRequest struct {
	StockId           string `json:"stock_id"`                 // 批次号
	OutRequestNo      string `json:"out_request_no"`           // 商户单据号
	AppId             string `json:"appid"`                    // 公众账号ID
	StockCreatorMchId string `json:"stock_creator_mchid"`      // 创建批次的商户ID
	CouponValue       uint64 `json:"coupon_value,omitempty"`   // 指定面额发券: 面额
	CouponMinimum     uint64 `json:"coupon_minimum,omitempty"` // 指定面额发券: 券门槛
}

// SendStockResponse 发放代金券批次API应答参数
type SendStockResponse struct {
	model.WechatError
	RequestId string `json:"-"`                   // 唯一请求ID
	CouponId  string `json:"coupon_id,omitempty"` // 代金券ID
}

// PauseStockResponse 暂停代金券批次应答参数
type PauseStockResponse struct {
	model.WechatError
	RequestId string `json:"-"`                    // 唯一请求ID
	PauseTime string `json:"pause_time,omitempty"` // 暂停时间
	StockId   string `json:"stock_id,omitempty"`   // 批次号
}

// RestartStockResponse 重启代金券批次应答参数
type RestartStockResponse struct {
	model.WechatError
	RequestId   string    `json:"-"`                      // 唯一请求ID
	StockId     string    `json:"stock_id,omitempty"`     // 批次号
	RestartTime time.Time `json:"restart_time,omitempty"` // 生效时间
}

// QueryStockListRequest 条件查询批次列表
type QueryStockListRequest struct {
	Offset            uint32 `json:"offset"`                      // 分页页码
	Limit             uint32 `json:"limit"`                       // 分页大小
	StockCreatorMchId string `json:"stock_creator_mchid"`         // 创建批次的商户号
	CreateStartTime   string `json:"create_start_time,omitempty"` // 起始时间
	CreateEndTime     string `json:"create_end_time,omitempty"`   // 终止时间
	Status            string `json:"status,omitempty"`            // 批次状态
}

// QueryStockListResponse 条件查询批次列表应答参数
type QueryStockListResponse struct {
	model.WechatError
	RequestId  string   `json:"-"`                     // 唯一请求ID
	TotalCount int64    `json:"total_count,omitempty"` // 批次总数
	Limit      uint32   `json:"limit,omitempty"`       // 分页大小
	Offset     uint32   `json:"offset,omitempty"`      // 分页页码
	Data       []*Stock `json:"data,omitempty"`        // 批次详情
}

// QueryStockResponse 查询批次详情请求参数
type QueryStockResponse struct {
	model.WechatError
	RequestId string `json:"-"` // 唯一请求ID
	Stock            // 代金券批次信息
}

// QueryCouponResponse 查询代金券详情
type QueryCouponResponse struct {
	model.WechatError
	NoCash                  bool               `json:"no_cash,omitempty"`                   // 是否无资金流
	SingleItem              bool               `json:"singleitem,omitempty"`                // 是否单品优惠
	RequestId               string             `json:"-"`                                   // 唯一请求ID
	StockCreatorMchId       string             `json:"stock_creator_mchid,omitempty"`       // 创建批次的商户号
	StockId                 string             `json:"stock_id,omitempty"`                  // 批次号
	CouponId                string             `json:"coupon_id,omitempty"`                 // 代金券ID
	CouponName              string             `json:"coupon_name,omitempty"`               // 代金券名称
	Status                  string             `json:"status,omitempty"`                    // 代金券状态
	Description             string             `json:"description,omitempty"`               // 使用说明
	CouponType              string             `json:"coupon_type,omitempty"`               // 券类型
	CreateTime              time.Time          `json:"create_time,omitempty"`               // 领券时间
	AvailableBeginTime      time.Time          `json:"available_begin_time,omitempty"`      // 可用开始时间
	AvailableEndTime        time.Time          `json:"available_end_time,omitempty"`        // 可用结束时间
	CutToMessage            *CutToMessage      `json:"cut_to_message,omitempty"`            // 单品优惠特定信息
	NormalCouponInformation *FixedNormalCoupon `json:"normal_coupon_information,omitempty"` // 满减券信息
}

// QueryStockMerchantRequest 查询代金券可用商户请求参数
type QueryStockMerchantRequest struct {
	Offset            uint32 `json:"offset,omitempty"`              // 分页页码
	Limit             uint32 `json:"limit,omitempty"`               // 分页大小
	StockCreatorMchId string `json:"stock_creator_mchid,omitempty"` // 创建批次的商户号
	StockId           string `json:"stock_id,omitempty"`            // 批次号
}

// QueryStockMerchantResponse 查询代金券可用商户应答参数
type QueryStockMerchantResponse struct {
	model.WechatError
	RequestId   string   `json:"-"`                     // 唯一请求ID
	TotalCount  uint32   `json:"total_count,omitempty"` // 可用商户总数量
	MerchantIds []string `json:"data,omitempty"`        // 可用商户列表
	Offset      uint32   `json:"offset,omitempty"`      // 分页页码
	Limit       uint32   `json:"limit,omitempty"`       // 分页大小
	StockId     string   `json:"stock_id,omitempty"`    // 批次号
}

// QueryStockItemRequest 查询代金券可用单品请求参数
type QueryStockItemRequest struct {
	Offset            uint32 `json:"offset,omitempty"`              // 分页页码
	Limit             uint32 `json:"limit,omitempty"`               // 分页大小
	StockCreatorMchId string `json:"stock_creator_mchid,omitempty"` // 创建批次的商户号
	StockId           string `json:"stock_id,omitempty"`            // 批次号
}

// QueryStockItemResponse 查询代金券可用单品应答参数
type QueryStockItemResponse struct {
	model.WechatError
	RequestId  string   `json:"-"`                     // 唯一请求ID
	TotalCount uint32   `json:"total_count,omitempty"` // 可用商户总数量
	Items      []string `json:"data,omitempty"`        // 可用单品列表
	Offset     uint32   `json:"offset,omitempty"`      // 分页页码
	Limit      uint32   `json:"limit,omitempty"`       // 分页大小
	StockId    string   `json:"stock_id,omitempty"`    // 批次号
}

// QueryUserCouponsRequest 根据商户号查用户的券API请求参数
type QueryUserCouponsRequest struct {
	OpenId         string `json:"openid"`                    // 用户标识
	AppId          string `json:"appid"`                     // 公众账号ID
	StockId        string `json:"stock_id,omitempty"`        // 批次ID
	Status         string `json:"status,omitempty"`          // 券状态
	CreatorMchId   string `json:"creator_mchid,omitempty"`   // 创建批次的商户ID
	SenderMchId    string `json:"sender_mchid,omitempty"`    // 批次发放商户号
	AvailableMchId string `json:"available_mchid,omitempty"` // 可用商户号
	Offset         uint32 `json:"offset,omitempty"`          // 分页页码
	Limit          uint32 `json:"limit,omitempty"`           // 分页大小
}

// QueryUserCouponsResponse 根据商户号查用户的券API应答参数
type QueryUserCouponsResponse struct {
	model.WechatError
	RequestId  string    `json:"-"`                     // 唯一请求ID
	TotalCount uint32    `json:"total_count,omitempty"` // 查询结果总数
	Limit      uint32    `json:"limit,omitempty"`       // 分页大小
	Offset     uint32    `json:"offset,omitempty"`      // 分页页码
	Coupons    []*Coupon `json:"data,omitempty"`        // 结果集
}

// DownloadRequest 明细下载请求参数
type DownloadRequest struct {
	StockId  string `json:"stock_id,omitempty"`  // 批次ID
	FileName string `json:"file_name,omitempty"` // 明细文件名
	FilePath string `json:"file_path,omitempty"` // 文件存放路径
}

// DownloadResponse 明细下载应答参数
type DownloadResponse struct {
	model.WechatError
	RequestId string `json:"-"`                    // 唯一请求ID
	Url       string `json:"url,omitempty"`        // 下载连接
	HashValue string `json:"hash_value,omitempty"` // 安全校验码
	HashType  string `json:"hash_type,omitempty"`  // 哈希算法类型
}

// SettingCallbacksResponse 设置消息通知地址
type SettingCallbacksResponse struct {
	model.WechatError
	RequestId  string    `json:"-"`                     // 唯一请求ID
	UpdateTime time.Time `json:"update_time,omitempty"` // 修改时间
	NotifyUrl  string    `json:"notify_url,omitempty"`  // 通知地址
}

// UseResponse 核销事件回调通知参数
type UseResponse struct {
	NotifyId                string              // 唯一通知ID
	NoCash                  bool                `json:"no_cash,omitempty"`                   // 是否无资金流
	SingleItem              bool                `json:"singleitem,omitempty"`                // 是否单品优惠
	StockCreatorMchId       string              `json:"stock_creator_mchid,omitempty"`       // 创建批次的商户号
	StockId                 string              `json:"stock_id,omitempty"`                  // 批次号
	CouponId                string              `json:"coupon_id,omitempty"`                 // 代金券ID
	CouponName              string              `json:"coupon_name,omitempty"`               // 代金券名称
	Status                  string              `json:"status,omitempty"`                    // 代金券状态
	Description             string              `json:"description,omitempty"`               // 使用说明
	CouponType              string              `json:"coupon_type,omitempty"`               // 券类型
	CreateTime              time.Time           `json:"create_time,omitempty"`               // 领券时间
	AvailableBeginTime      time.Time           `json:"available_begin_time,omitempty"`      // 可用开始时间
	AvailableEndTime        time.Time           `json:"available_end_time,omitempty"`        // 可用结束时间
	SingleDiscountOff       *SingleDiscountOff  `json:"single_discount_off,omitempty"`       // 单品优惠特定信息
	DiscountTo              *DiscountTo         `json:"discount_to,omitempty"`               // 减至优惠特定信息
	NormalCouponInformation *FixedNormalCoupon  `json:"normal_coupon_information,omitempty"` // 普通满减券信息
	ConsumeInformation      *ConsumeInformation `json:"consume_information,omitempty"`       // 实扣代金券信息
}

// UploadImageResponse 上传图片
type UploadImageResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	MediaUrl  string `json:"media_url,omitempty"` // 媒体文件URL地址
}

// Stock 代金券批次详情
type Stock struct {
	SingleItem         bool          `json:"singleitem,omitempty"`           // 是否单品优惠
	NoCash             bool          `json:"no_cash,omitempty"`              // 是否无资金流
	StockId            string        `json:"stock_id,omitempty"`             // 批次号
	StockCreatorMchId  string        `json:"stock_creator_mchid,omitempty"`  // 创建批次的商户号
	StockName          string        `json:"stock_name,omitempty"`           // 批次名称
	Status             string        `json:"status,omitempty"`               // 批次状态
	StockType          string        `json:"stock_type,omitempty"`           // 批次类型
	Description        string        `json:"description,omitempty"`          // 使用说明
	DistributedCoupons uint32        `json:"distributed_coupons,omitempty"`  // 已发券数量
	CreateTime         time.Time     `json:"create_time,omitempty"`          // 创建时间
	AvailableBeginTime time.Time     `json:"available_begin_time,omitempty"` // 可用开始时间
	AvailableEndTime   time.Time     `json:"available_end_time,omitempty"`   // 可用结束时间
	StartTime          time.Time     `json:"start_time,omitempty"`           // 激活批次的时间
	StopTime           time.Time     `json:"stop_time,omitempty"`            // 终止批次的时间
	StockUseRule       *StockUseRule `json:"stock_use_rule,omitempty"`       // 满减券批次使用规则
	CutToMessage       *CutToMessage `json:"cut_to_message,omitempty"`       // 减至批次特定信息
}

// StockUseRule 满减批次使用规则
type StockUseRule struct {
	CombineUse        bool               `json:"combine_use,omitempty"`          // 是否可叠加其他优惠
	MaxCouponsPerUser uint32             `json:"max_coupons_per_user,omitempty"` // 单个用户可领个数
	CouponType        string             `json:"coupon_type,omitempty"`          // 券类型
	MaxCoupons        uint64             `json:"max_coupons,omitempty"`          // 最大发券数
	MaxAmount         uint64             `json:"max_amount,omitempty"`           // 总预算
	MaxAmountByDay    uint64             `json:"max_amount_by_day,omitempty"`    // 单天发放上限金额
	RawTradeTypes     json.RawMessage    `json:"trade_type,omitempty"`           // 用于接收支付方式的序列化数据
	TradeType         []string           `json:"-"`                              // 支付方式
	GoodsTag          []string           `json:"goods_tag,omitempty"`            // 订单优惠标记
	FixedNormalCoupon *FixedNormalCoupon `json:"fixed_normal_coupon,omitempty"`  // 固定面额批次特定信息
}

// FixedNormalCoupon 固定面额批次特定信息
type FixedNormalCoupon struct {
	CouponAmount       uint64 `json:"coupon_amount,omitempty"`       // 面额
	TransactionMinimum uint64 `json:"transaction_minimum,omitempty"` // 门槛
}

// CutToMessage 减至批次特定信息
type CutToMessage struct {
	SinglePriceMax uint64 `json:"single_price_max,omitempty"` // 可用优惠的商品最高单价
	CutToPrice     uint64 `json:"cut_to_price,omitempty"`     // 减至后的优惠单价
}

// Coupon 优惠券
type Coupon struct {
	NoCash                  bool                `json:"no_cash,omitempty"`                   // 是否无资金流
	SingleItem              bool                `json:"singleitem,omitempty"`                // 是否单品优惠
	StockCreatorMchId       string              `json:"stock_creator_mchid,omitempty"`       // 创建批次的商户号
	StockId                 string              `json:"stock_id,omitempty"`                  // 批次号
	CouponId                string              `json:"coupon_id,omitempty"`                 // 代金券ID
	CouponName              string              `json:"coupon_name,omitempty"`               // 代金券名称
	Status                  string              `json:"status,omitempty"`                    // 代金券状态
	Description             string              `json:"description,omitempty"`               // 使用说明
	CouponType              string              `json:"coupon_type,omitempty"`               // 券类型
	CreateTime              time.Time           `json:"create_time,omitempty"`               // 领券时间
	AvailableBeginTime      time.Time           `json:"available_begin_time,omitempty"`      // 可用开始时间
	AvailableEndTime        time.Time           `json:"available_end_time,omitempty"`        // 可用结束时间
	CutToMessage            *CutToMessage       `json:"cut_to_message,omitempty"`            // 单品优惠特定信息
	NormalCouponInformation *FixedNormalCoupon  `json:"normal_coupon_information,omitempty"` // 满减券信息
	ConsumeInformation      *ConsumeInformation `json:"consume_information,omitempty"`       // 已实扣代金券信息
}

// ConsumeInformation 已实扣代金券核销信息
type ConsumeInformation struct {
	ConsumeMchId  string         `json:"consume_mchid,omitempty"`  // 核销商户号
	TransactionId string         `json:"transaction_id,omitempty"` // 支付单号
	ConsumeTime   time.Time      `json:"consume_time,omitempty"`   // 核销时间
	GoodsDetail   []*GoodsDetail `json:"goods_detail,omitempty"`   // 单品信息
}

// GoodsDetail 单品信息
type GoodsDetail struct {
	Quantity       uint32 `json:"quantity,omitempty"`        // 商品数量
	GoodsId        string `json:"goods_id,omitempty"`        // 商品编码
	Price          int64  `json:"price,omitempty"`           // 商品价格
	DiscountAmount int64  `json:"discount_amount,omitempty"` // 优惠金额
}

// SingleDiscountOff 单品优惠特定信息
type SingleDiscountOff struct {
	SinglePriceMax int64 `json:"single_price_max,omitempty"`
}

// DiscountTo 减至优惠特定信心
type DiscountTo struct {
	CutToPrice int64 `json:"cut_to_price,omitempty"` // 减至后优惠单价
	MaxPrice   int64 `json:"max_price,omitempty"`    // 最高价格
}

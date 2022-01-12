package parking

import (
	"time"

	"github.com/pyihe/wechat-sdk/model"
)

// FindRequest 查询车牌服务开通信息请求参数
type FindRequest struct {
	AppId       string `json:"appid,omitempty"`        // 应用ID
	SubMchId    string `json:"sub_mchid,omitempty"`    // 子商户号
	PlateNumber string `json:"plate_number,omitempty"` // 车牌号
	PlateColor  string `json:"plate_color,omitempty"`  // 车牌颜色
	OpenId      string `json:"openid,omitempty"`       // 用户标识
}

// FindResponse 查询车牌服务开通信息应答参数
type FindResponse struct {
	model.WechatError
	RequestId       string    `json:"-"`                           // 唯一请求ID
	PlateNumber     string    `json:"plate_number,omitempty"`      // 车牌号
	PlateColor      string    `json:"plate_color,omitempty"`       // 车牌颜色
	OpenId          string    `json:"openid,omitempty"`            // 用户标识
	ServiceOpenTime time.Time `json:"service_open_time,omitempty"` // 车牌服务开通时间
	ServiceState    string    `json:"service_state,omitempty"`     // 车牌服务开通壮体啊
}

// CreateParkingResponse 创建停车入场应答参数
type CreateParkingResponse struct {
	model.WechatError
	RequestId    string    `json:"-"`                        // 唯一请求ID
	Id           string    `json:"id,omitempty"`             // 停车入场ID
	OutParkingNo string    `json:"out_parking_no,omitempty"` // 商户入场ID
	PlateNumber  string    `json:"plate_number,omitempty"`   // 车牌号
	PlateColor   string    `json:"plate_color,omitempty"`    // 车牌颜色
	StartTime    time.Time `json:"start_time,omitempty"`     // 入场时间
	ParkingName  string    `json:"parking_name,omitempty"`   // 停车场名称
	FreeDuration int32     `json:"free_duration,omitempty"`  // 免费时长
	State        string    `json:"state,omitempty"`          // 停车入场状态
	BlockReason  string    `json:"block_reason,omitempty"`   // 不可用状态描述
}

// TransactionsResponse 扣费受理应答参数
type TransactionsResponse struct {
	model.WechatError
	RequestId             string                   `json:"-"`                                 // 唯一请求ID
	AppId                 string                   `json:"appid,omitempty"`                   // 应用ID
	SubAppId              string                   `json:"sub_appid,omitempty"`               // 子商户应用ID
	SpMchId               string                   `json:"sp_mchid,omitempty"`                // 商户号
	SubMchId              string                   `json:"sub_mchid,omitempty"`               // 子商户号
	Description           string                   `json:"description,omitempty"`             // 服务描述
	CreateTime            time.Time                `json:"create_time,omitempty"`             // 订单创建时间
	OutTradeNo            string                   `json:"out_trade_no,omitempty"`            // 商户订单号
	TransactionId         string                   `json:"transaction_id,omitempty"`          // 微信支付订单号
	TradeState            string                   `json:"trade_state,omitempty"`             // 交易状态
	TradeStateDescription string                   `json:"trade_state_description,omitempty"` // 交易状态描述
	SuccessTime           time.Time                `json:"success_time,omitempty"`            // 支付完成时间
	BankType              string                   `json:"bank_type,omitempty"`               // 付款银行
	UserRepaid            string                   `json:"user_repaid,omitempty"`             // 用户是否已还款
	Attach                string                   `json:"attach,omitempty"`                  // 附加数据
	TradeScene            string                   `json:"trade_scene,omitempty"`             // 交易场景
	ParkingInfo           *ParkInfo                `json:"parking_info,omitempty"`            // 停车场景信息
	Payer                 *Payer                   `json:"payer,omitempty"`                   // 支付者信息
	Amount                *model.Amount            `json:"amount,omitempty"`                  // 订单金额信息
	PromotionDetail       []*model.PromotionDetail `json:"promotion_detail,omitempty"`        // 优惠信息
}

// Payer 支付者信息
type Payer struct {
	OpenId    string `json:"openid,omitempty"`     // 用户标识
	SubOpenId string `json:"sub_openid,omitempty"` // 用户在sub_appid下的标识
}

// ParkInfo 停车场景信息
type ParkInfo struct {
	ParkingId        string    `json:"parking_id,omitempty"`        // 停车入场ID
	PlateNumber      string    `json:"plate_number,omitempty"`      // 车牌号
	PlateColor       string    `json:"plate_color,omitempty"`       // 车牌颜色
	StartTime        time.Time `json:"start_time,omitempty"`        // 入场时间
	EndTime          time.Time `json:"end_time,omitempty"`          // 出场时间
	ParkingName      string    `json:"parking_name,omitempty"`      // 停车场名称
	ChargingDuration int32     `json:"charging_duration,omitempty"` // 计费时长
	DeviceId         string    `json:"device_id,omitempty"`         // 停车场设备ID
}

// QueryResponse 查询订单应答参数
type QueryResponse struct {
	model.WechatError
	RequestId             string                   `json:"-"`                                 // 唯一请求ID
	AppId                 string                   `json:"appid,omitempty"`                   // 应用ID
	SubAppId              string                   `json:"sub_appid,omitempty"`               // 子商户应用ID
	SpMchId               string                   `json:"sp_mchid,omitempty"`                // 商户号
	SubMchId              string                   `json:"sub_mchid,omitempty"`               // 子商户号
	Description           string                   `json:"description,omitempty"`             // 服务描述
	CreateTime            time.Time                `json:"create_time,omitempty"`             // 订单创建时间
	OutTradeNo            string                   `json:"out_trade_no,omitempty"`            // 商户订单号
	TransactionId         string                   `json:"transaction_id,omitempty"`          // 微信支付订单号
	TradeState            string                   `json:"trade_state,omitempty"`             // 交易状态
	TradeStateDescription string                   `json:"trade_state_description,omitempty"` // 交易状态描述
	SuccessTime           time.Time                `json:"success_time,omitempty"`            // 支付完成时间
	BankType              string                   `json:"bank_type,omitempty"`               // 付款银行
	UserRepaid            string                   `json:"user_repaid,omitempty"`             // 用户是否已还款
	Attach                string                   `json:"attach,omitempty"`                  // 附加数据
	TradeScene            string                   `json:"trade_scene,omitempty"`             // 交易场景
	ParkingInfo           *ParkInfo                `json:"parking_info,omitempty"`            // 停车场景信息
	Payer                 *Payer                   `json:"payer,omitempty"`                   // 支付者信息
	Amount                *model.Amount            `json:"amount,omitempty"`                  // 订单金额信息
	PromotionDetail       []*model.PromotionDetail `json:"promotion_detail,omitempty"`        // 优惠信息
}

// ParkStateResponse 停车入场状态变更通知参数
type ParkStateResponse struct {
	Id                      string    `json:"id,omitempty"`                        // 唯一通知ID
	SpMchId                 string    `json:"sp_mchid,omitempty"`                  // 商户号
	SubMchId                string    `json:"sub_mchid,omitempty"`                 // 子商户号
	ParkingId               string    `json:"parking_id,omitempty"`                // 停车入场ID
	OutParkingNo            string    `json:"out_parking_no,omitempty"`            // 商户入场ID
	PlateNumber             string    `json:"plate_number,omitempty"`              // 车牌号
	PlateColor              string    `json:"plate_color,omitempty"`               // 车牌颜色
	StartTime               time.Time `json:"start_time,omitempty"`                // 入场时间
	ParkingName             string    `json:"parking_name,omitempty"`              // 停车场名称
	FreeDuration            int32     `json:"free_duration,omitempty"`             // 免费停车时长
	BlockedStateDescription string    `json:"blocked_state_description,omitempty"` // 不可用状态描述
	StateUpdateTime         time.Time `json:"state_update_time,omitempty"`         // 状态变更时间
}

// PaymentResponse 支付结果通知参数
type PaymentResponse struct {
	Id                    string                   `json:"-"`                                 // 唯一通知ID
	AppId                 string                   `json:"appid,omitempty"`                   // 应用ID
	SpMchId               string                   `json:"sp_mchid,omitempty"`                // 商户号
	SubAppId              string                   `json:"sub_appid,omitempty"`               // 子商户应用ID
	SubMchId              string                   `json:"sub_mchid,omitempty"`               // 子商户号
	OutTradeNo            string                   `json:"out_trade_no,omitempty"`            // 商户
	TransactionId         string                   `json:"transaction_id,omitempty"`          // 微信支付订单号
	Description           string                   `json:"description,omitempty"`             // 服务描述
	CreateTime            time.Time                `json:"create_time,omitempty"`             // 订单创建时间
	TradeState            string                   `json:"trade_state,omitempty"`             // 交易状态
	TradeStateDescription string                   `json:"trade_state_description,omitempty"` // 交易状态描述
	SuccessTime           time.Time                `json:"success_time,omitempty"`            // 支付完成时间
	BankType              string                   `json:"bank_type,omitempty"`               // 付款银行
	Attach                string                   `json:"attach,omitempty"`                  // 附加数据
	UserRepaid            string                   `json:"user_repaid,omitempty"`             // 用户是否已还款
	TradeScene            string                   `json:"trade_scene,omitempty"`             // 交易场景
	ParkingInfo           *ParkInfo                `json:"parking_info,omitempty"`            // 停车场景信息
	Payer                 *Payer                   `json:"payer,omitempty"`                   // 支付者信息
	Amount                *model.Amount            `json:"amount,omitempty"`                  // 订单金额信息
	PromotionDetail       []*model.PromotionDetail `json:"promotion_detail,omitempty"`        // 优惠信息
}

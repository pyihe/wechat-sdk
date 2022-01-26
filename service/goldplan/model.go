package goldplan

import "github.com/pyihe/wechat-sdk/v3/model"

// ChangeStatusRequest 点金计划管理API请求参数
type ChangeStatusRequest struct {
	SubMchId      string `json:"sub_mchid"`      // 特约商户号
	OperationType string `json:"operation_type"` // 操作类型
}

// ChangeStatusResponse 点金计划管理API应答参数
type ChangeStatusResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	SubMchId  string `json:"sub_mchid,omitempty"` // 特约商户号
}

// ChangeCustomPageStatusRequest 商家小票管理API请求参数
type ChangeCustomPageStatusRequest struct {
	SubMchId      string `json:"sub_mchid"`      // 特约商户号
	OperationType string `json:"operation_type"` // 操作类型
}

// ChangeCustomPageStatusResponse 商家小票管理API应答参数
type ChangeCustomPageStatusResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	SubMchId  string `json:"sub_mchid,omitempty"` // 特约商户号
}

// SetIndustryFilterRequest 同行业过滤标签管理请求参数
type SetIndustryFilterRequest struct {
	SubMchId                   string   `json:"sub_mchid"`                    // 特约商户号
	AdvertisingIndustryFilters []string `json:"advertising_industry_filters"` // 同业过滤标签值
}

// SetIndustryFilterResponse 同行业过滤标签管理应答参数
type SetIndustryFilterResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

// OpenAdvertisingShowRequest 开通广告展示请求参数
type OpenAdvertisingShowRequest struct {
	SubMchId                   string   `json:"sub_mchid"`                              // 特约商户号
	AdvertisingIndustryFilters []string `json:"advertising_industry_filters,omitempty"` // 同业过滤标签值
}

// OpenAdvertisingShowResponse 开通广告展示应答参数
type OpenAdvertisingShowResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

// CloseAdvertisingShowRequest 关闭广告展示请求参数
type CloseAdvertisingShowRequest struct {
	SubMchId string `json:"sub_mchid"` // 特约商户号
}

// CloseAdvertisingShowResponse 关闭广告展示应答参数
type CloseAdvertisingShowResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

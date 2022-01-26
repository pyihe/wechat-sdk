package partnership

import (
	"time"

	"github.com/pyihe/wechat-sdk/v3/model"
)

// BuildRequest 建立合作关系
type BuildRequest struct {
	Partner        *Partner        `json:"partner"`         // 合作方信息
	AuthorizedData *AuthorizedData `json:"authorized_data"` // 被授权数据
}

// BuildResponse 建立合作关系应答参数
type BuildResponse struct {
	model.WechatError
	RequestId      string               // 唯一请求ID
	State          string               `json:"state,omitempty"`           // 合作状态
	Partner        *Partner             `json:"partner,omitempty"`         // 合作方信息
	AuthorizedData *BuildAuthorizedData `json:"authorized_data,omitempty"` // 被授权数据
	BuildTime      time.Time            `json:"build_time,omitempty"`      // 建立合作关系事件
	CreateTime     time.Time            `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     time.Time            `json:"update_time,omitempty"`     // 更新时间
}

// QueryRequest 查询合作关系列表请求参数
type QueryRequest struct {
	Partner           *Partner        `json:"partner,omitempty"`  // 合作方信息
	AuthorizationData *AuthorizedData `json:"authorization_data"` // 被授权数据
	Limit             uint64          `json:"limit,omitempty"`    // 分页大小
	Offset            uint64          `json:"offset,omitempty"`   // 分页页码
}

// QueryResponse 查询合作关系应答参数
type QueryResponse struct {
	model.WechatError
	RequestId  string  // 唯一请求ID
	Data       []*Ship `json:"data,omitempty"`        // 合作关系结果集
	Limit      uint64  `json:"limit,omitempty"`       // 分页大小
	Offset     uint64  `json:"offset,omitempty"`      // 分页页码
	TotalCount uint64  `json:"total_count,omitempty"` // 总数量
}

type Ship struct {
	Partner        *Partner        `json:"partner,omitempty"`         // 合作方信息
	AuthorizedData *AuthorizedData `json:"authorized_data,omitempty"` // 被授权数据
	BuildTime      time.Time       `json:"build_time,omitempty"`      // 建立合作关系时间
	TerminateTime  time.Time       `json:"terminate_time,omitempty"`  // 终止合作关系时间
	CreateTime     time.Time       `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     time.Time       `json:"update_time,omitempty"`     // 更新时间
}

// Partner 合作方信息
type Partner struct {
	Type       string `json:"type"`                  // 合作方类别
	AppId      string `json:"appid,omitempty"`       // 合作方appid
	MerchantId string `json:"merchant_id,omitempty"` // 合作方商户ID
}

// AuthorizedData 被授权数据(请求参数)
type AuthorizedData struct {
	BusinessType string `json:"business_type"`      // 授权业务类别
	StockId      string `json:"stock_id,omitempty"` // 授权批次ID
}

// BuildAuthorizedData 被授权数据(应答参数)
type BuildAuthorizedData struct {
	BusinessType string   `json:"business_type"`       // 授权业务类别
	StockId      string   `json:"stock_id,omitempty"`  // 授权批次ID
	Scenarios    []string `json:"scenarios,omitempty"` // 授权场景
}

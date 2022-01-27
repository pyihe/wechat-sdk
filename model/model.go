package model

import (
	"fmt"
)

type WError interface {
	Error() error
}

// WechatHeader 微信Header
type WechatHeader struct {
	RequestId string // 唯一请求ID
	NotifyId  string // 唯一通知ID
}

// WechatError 调用微信接口返回的通用错误格式
type WechatError struct {
	Code    interface{} `json:"code,omitempty"`    // 详细错误码
	Message interface{} `json:"message,omitempty"` // 错误描述
	Detail  ErrorDetail `json:"detail,omitempty"`  // 错误详细信息
}

func (w WechatError) Error() error {
	if w.Code == nil &&
		w.Message == nil &&
		w.Detail.Field == nil &&
		w.Detail.Value == nil &&
		w.Detail.Issue == nil &&
		w.Detail.Location == nil {
		return nil
	}
	return fmt.Errorf(`{"code":"%v" "message":"%v" "field":"%v" "value":"%v" "issue":"%v" "location":"%v"}`, w.Code, w.Message, w.Detail.Field, w.Detail.Value, w.Detail.Issue, w.Detail.Location)
}

// ErrorDetail 错误详情
type ErrorDetail struct {
	Field    interface{} `json:"field,omitempty"`    // 错误参数的位置
	Value    interface{} `json:"value,omitempty"`    // 错误的值
	Issue    interface{} `json:"issue,omitempty"`    // 具体错误的原因
	Location interface{} `json:"location,omitempty"` // 出错的位置
}

// WechatNotifyResponse 微信通知的回复格式
type WechatNotifyResponse struct {
	Id           string            `json:"id,omitempty"`            // 通知的唯一ID
	CreateTime   string            `json:"create_time,omitempty"`   // 通知创建时间
	EventType    string            `json:"event_type,omitempty"`    // 通知类型
	ResourceType string            `json:"resource_type,omitempty"` // 通知的资源数据类型
	Summary      string            `json:"summary,omitempty"`       // 回调摘要
	Resource     *WechatCipherData `json:"resource,omitempty"`      // 通知资源数据
}

// WechatCipherData 微信返回的加密数据
type WechatCipherData struct {
	Algorithm      string `json:"algorithm,omitempty"`       // 加密算法
	CipherText     string `json:"ciphertext,omitempty"`      // 密文
	AssociatedData string `json:"associated_data,omitempty"` // 附加数据
	OriginalType   string `json:"original_type,omitempty"`   // 原始回调类型
	Nonce          string `json:"nonce,omitempty"`           // 加密使用的随机串
}

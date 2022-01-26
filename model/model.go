package model

import (
	"fmt"
)

// WechatHeader 微信Header
type WechatHeader struct {
	RequestId string // 唯一请求ID
	NotifyId  string // 唯一通知ID
}

// WechatError 调用微信接口返回的通用错误格式
type WechatError struct {
	Code    string      `json:"code,omitempty"`    // 详细错误码
	Message string      `json:"message,omitempty"` // 错误描述
	Detail  ErrorDetail `json:"detail,omitempty"`  // 错误详细信息
}

func (w WechatError) Error() error {
	if len(w.Code) == 0 &&
		len(w.Message) == 0 &&
		len(w.Detail.Field) == 0 &&
		len(w.Detail.Value) == 0 &&
		len(w.Detail.Issue) == 0 &&
		len(w.Detail.Location) == 0 {
		return nil
	}
	return fmt.Errorf(`{"code":"%s" "message":"%s" "field":"%s" "value":"%s" "issue":"%s" "location":"%s"}`, w.Code, w.Message, w.Detail.Field, w.Detail.Value, w.Detail.Issue, w.Detail.Location)
}

// ErrorDetail 错误详情
type ErrorDetail struct {
	Field    string `json:"field,omitempty"`    // 错误参数的位置
	Value    string `json:"value,omitempty"`    // 错误的值
	Issue    string `json:"issue,omitempty"`    // 具体错误的原因
	Location string `json:"location,omitempty"` // 出错的位置
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

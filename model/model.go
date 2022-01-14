package model

// WechatError 调用微信接口返回的通用错误格式
type WechatError struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Detail  ErrorDetail `json:"detail,omitempty"`
}

func (w WechatError) IsZero() bool {
	if len(w.Code) != 0 {
		return false
	}
	if len(w.Message) != 0 {
		return false
	}
	return w.Detail.IsZero()
}

// ErrorDetail 错误详情
type ErrorDetail struct {
	Field    string `json:"field,omitempty"`
	Value    string `json:"value,omitempty"`
	Issue    string `json:"issue,omitempty"`
	Location string `json:"location,omitempty"`
}

func (e ErrorDetail) IsZero() bool {
	if len(e.Field) != 0 {
		return false
	}
	if len(e.Value) != 0 {
		return false
	}
	if len(e.Issue) != 0 {
		return false
	}
	if len(e.Location) != 0 {
		return false
	}
	return true
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

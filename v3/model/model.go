package model

// Response 用于商户想微信服务端发送通知
type Response struct {
	Code    string `json:"code,omitempty"`    // 返回状态码
	Message string `json:"message,omitempty"` // 返回信息
}

type WechatNotifyResponse struct {
	Id           string            `json:"id,omitempty"`            // 通知的唯一ID
	CreateTime   string            `json:"create_time,omitempty"`   // 通知创建时间
	EventType    string            `json:"event_type,omitempty"`    // 通知类型
	ResourceType string            `json:"resource_type,omitempty"` // 通知的资源数据类型
	Summary      string            `json:"summary,omitempty"`       // 回调摘要
	Resource     *WechatCipherData `json:"resource,omitempty"`      // 通知资源数据
}

type WechatCipherData struct {
	Algorithm      string `json:"algorithm,omitempty"`       // 加密算法
	CipherText     string `json:"cipher_text,omitempty"`     // 密文
	AssociatedData string `json:"associated_data,omitempty"` // 附加数据
	OriginalType   string `json:"original_type,omitempty"`   // 原始回调类型
	Nonce          string `json:"nonce,omitempty"`           // 加密使用的随机串
}

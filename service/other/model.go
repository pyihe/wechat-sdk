package other

import "github.com/pyihe/wechat-sdk/v3/model"

// UploadResponse 图片上传应答参数
type UploadResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
	MediaId   string `json:"media_id,omitempty"` // 媒体文件标识ID
}

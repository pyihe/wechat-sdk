package other

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pyihe/wechat-sdk/v3/pkg/errors"
	"github.com/pyihe/wechat-sdk/v3/service"
)

// UploadImage 图片上传
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter2_1_1.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter2_1_1.shtml
func UploadImage(config *service.Config, fileName string, image interface{}) (uploadResponse *UploadResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	if fileName == "" {
		err = fmt.Errorf("请提供图片文件名")
		return
	}
	var content []byte
	var contentType string

	contentType, err = service.VideoExt(fileName)
	if err != nil {
		return
	}
	switch data := image.(type) {
	case string:
		content, err = ioutil.ReadFile(data)
	case []byte:
		content = data
	case io.Reader:
		content, err = ioutil.ReadAll(data)
	case io.ReadCloser:
		content, err = ioutil.ReadAll(data)
		_ = data.Close()
	default:
		err = errors.ErrImageFormatType
	}
	if err != nil {
		return
	}
	response, err := config.UploadMedia("/v3/merchant/media/upload", contentType, fileName, content)
	if err != nil {
		return
	}
	uploadResponse = new(UploadResponse)
	uploadResponse.RequestId, err = config.ParseWechatResponse(response, uploadResponse)
	return
}

// UploadVideo 视频上传
// 商户平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter2_1_2.shtml
// 服务商平台文档: https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter2_1_2.shtml
func UploadVideo(config *service.Config, fileName string, video interface{}) (uploadResponse *UploadResponse, err error) {
	if config == nil {
		err = errors.ErrNoConfig
		return
	}
	var content []byte
	var contentType string

	contentType, err = service.ImageExt(fileName)
	if err != nil {
		return
	}
	switch data := video.(type) {
	case string:
		content, err = ioutil.ReadFile(data)
	case []byte:
		content = data
	case io.Reader:
		content, err = ioutil.ReadAll(data)
	case io.ReadCloser:
		content, err = ioutil.ReadAll(data)
		_ = data.Close()
	default:
		err = errors.ErrVideoFormatType
	}

	response, err := config.UploadMedia("/v3/merchant/media/video_upload", contentType, fileName, content)
	if err != nil {
		return
	}
	uploadResponse = new(UploadResponse)
	uploadResponse.RequestId, err = config.ParseWechatResponse(response, uploadResponse)
	return
}

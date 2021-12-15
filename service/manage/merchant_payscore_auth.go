package manage

import (
	"net/http"

	"github.com/pyihe/wechat-sdk/model/manage/merchant"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/model"
	"github.com/pyihe/wechat-sdk/pkg/rsas"
	"github.com/pyihe/wechat-sdk/service"
	"github.com/pyihe/wechat-sdk/vars"
)

/*微信支付分(需确认模式)*/

// ConfirmOrderNotify 确认订单回调通知
// API详细介绍: https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter6_1_21.shtml
func ConfirmOrderNotify(config *service.Config, w http.ResponseWriter, request *http.Request) (confirmResponse *merchant.PayscoreOrder, err error) {
	if config == nil {
		err = vars.ErrInitConfig
		return
	}
	if config.ApiKey == "" {
		err = vars.ErrNoApiV3Key
		return
	}
	if request == nil {
		err = vars.ErrNoRequest
		return
	}
	body, err := service.VerifyRequest(config, request)
	if err != nil {
		return
	}
	// 验证签名
	notifyResponse := new(model.WechatNotifyResponse)
	if err = service.Unmarshal(body, &notifyResponse); err != nil {
		return
	}
	// 判断资源类型
	if notifyResponse.ResourceType != "encrypt-resource" {
		err = errors.New("错误的资源类型: " + notifyResponse.ResourceType)
		return
	}
	if notifyResponse.Resource == nil {
		err = errors.New("未获取到通知资源数据!")
		return
	}
	// 解密
	cipherText := notifyResponse.Resource.CipherText
	associateData := notifyResponse.Resource.AssociatedData
	nonce := notifyResponse.Resource.Nonce
	plainText, err := rsas.DecryptAEADAES256GCM(config.Cipher, config.ApiKey, cipherText, associateData, nonce)
	if err != nil {
		return
	}

	confirmResponse = new(merchant.PayscoreOrder)
	confirmResponse.Id = notifyResponse.Id
	if err = service.Unmarshal(plainText, &confirmResponse); err != nil {
		return
	}
	if config.PayScoreConfirmNotifyHandler != nil && w != nil {
		response := new(model.Response)
		response.Code = "SUCCESS"
		response.Message = "成功"
		if err = config.PayScoreConfirmNotifyHandler(confirmResponse); err != nil {
			response.Code = "FAIL"
			response.Message = err.Error()
		}
		data, _ := service.Marshal(response)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}
	return
}

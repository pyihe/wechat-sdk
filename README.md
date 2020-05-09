# wechat-sdk
##### 功能列表
| 方法名  |  解释 |  备注 |
| :---- | :----| :----|
| GetUserPhoneForMini | 小程序获取电话号码 | 客户端调用微信接口获取加密信息时不能在回调中再次调用登陆接口, 否则会让session_key失效 |
| GetSessionKeyAndOpenId | 通过小程序授权code获取session_key和用户openid |
| GetAccessTokenForMini | 小程序获取AccessToken | |
| GetUserInfoForMini | 小程序获取用户基本信息 | |
| CloseOrder |  关闭订单 |   |
| DownloadBill | 下载对账单  |   |
| DownloadComment | 拉取订单评论 |  |
| DownloadFundFlow | 下载资金账单 |  |
| GetPublicKey | 获取RSA加密公钥 |  |
| RefundOrder | 申请退款 |  |
| RefundNotify | 解析微信退款回调内容, 主要是对req_info解密 |   |
| RefundQuery | 退款查询 |  |
| Report | 向微信发送接口调用结果的报告, 包括接口调用时间 |  |
| ReverseOrder | 撤销订单 |  |
| Transfers | 企业付款到用户零钱 |  |
| TransfersQuery | 查询企业付款到用户零钱的结果 |  |
| TransferBank | 企业付款到银行卡 | 未测试 |
| TransferBankQuery | 查询企业付款到银行卡的结果 | 未测试 |
| UnifiedMicro | 扫码下单 | 如果需要传detail参数, 参数格式参照[单品优惠活动detail字段列表说明](https://pay.weixin.qq.com/wiki/doc/api/danpin.php?chapter=9_102&index=2) |
| UnifiedOrder | 统一下单: H5/APP/MWEB/NATIVE | 返回给前端的唤起支付参数中, package = prepay_id=xxxxxxx |
| UnifiedQuery | 下单结果查询 |  |
| GetAppBaseAccessToken | 公众号开发获取基础接口调用Access_Token | 用于调用其他接口 |
| GetAppOauthAccessToken | 公众号开发通过code获取网页授权Access_Token | 此Access_Token用于拉取用户信息, code只能用一次 |
| RefreshOauthToken | 公众号开发刷新网页授权Access_Token ||
| GetAppUserInfo | 公众号开发拉取用户信息 | 返回结果参考[微信文档](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html) |
| CheckOauthToken | 公众号开发校验网页授权Access_Token是否有效 ||

**Notice: 所有请求接口都不需要加入appid/mch_id/key/secret/sign参数**

##### 如何使用: 
```
package main

import (
	"fmt"
	dev "github.com/hong008/wechat-sdk"
)

func main() {
	var appId, mchId, apiKey, apiSecret string

	client := dev.NewPayer(dev.WithAppId(appId), dev.WithMchId(mchId), dev.WithApiKey(apiKey), dev.WithSecret(apiSecret))

	//unified order
	param := dev.NewParam()
	param.Add("nonce_str", "yourNonceStr")
	param.Add("body", "yourBody")
	param.Add("out_trade_no", "yourOutTradeNo")
	param.Add("total_fee", 1)
	param.Add("spbill_create_ip", "yourIp")
	param.Add("notify_url", "yourUrl")
	param.Add("trade_type", "JSAPI")
	result, err := client.UnifiedOrder(param)
	if err != nil {
		handleErr(err)
	}
    appId, _ := result.GetString("apppid")
    prepayId, _ := result.GetString("prepay_id")
    param = dev.NewParam()
    param.Add("appId", appId)
    param.Add("timeStamp", time.Now().Unix())
    param.Add("nonceStr", "nonceStr")
    param.Add("package", "prepay_id="+prepayId)
    param.Add("signType", "MD5")
    //use to evoke wechat pay 
    sign := param.Sign("MD5")


    //download bill
	param = dev.NewParam()
	param.Add("nonce_str", "yourNonceStr")
	param.Add("bill_date", "yourDate")
	param.Add("bill_type", "ALL")
	param.Add("tar_type", "GZIP")
	err := client.DownloadBill(param, "./bill")
	if err != nil {
		handleErr(err)
	}
    

    //get phone for mini program user
    result, err := client.GetUserPhoneForMini("code", "encryptedData", "iv")
    if err != nil {
    	handleErr(err)
    }
    var phone string
    if countryCode := result.Get("countryCode"); countryCode != nil && countryCode.(string) == "86" {
    	purePhone := result.Get("purePhoneNumber")
    	phone = purePhone.(string)
    } else {
    	phoneNumber := result.Get("phoneNumber")
    	phone = phoneNumber.(string)
    }
    fmt.Printf("user phone is %s\n", phone)
}

```

```
package main

import (
	"fmt"
	"net/http"
	
	dev "github.com/hong008/wechat-sdk"
)

func main() {
	var appId, mchId, apiKey, apiSecret string

	client := dev.NewPayer(dev.WithAppId(appId), dev.WithMchId(mchId), dev.WithApiKey(apiKey), dev.WithSecret(apiSecret))

	//handle refund notify
	http.HandleFunc("/refund_notify", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		result, err := client.RefundNotify(request.Body)
		if err != nil {
			handleErr(err)
		}
		fmt.Printf("RefundNotify Result = %v\n", result.Data())
	})
	http.ListenAndServe(":8810", nil)
}

```

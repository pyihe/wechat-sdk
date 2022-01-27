# wechat-sdk

### Update

更新至V3版本微信API

### 如何在项目中引用

#### V2版本微信API

```go
// get go mod
go get github.com/pyihe/wechat-sdk@v1.0.0

// import in project
import "github.com/pyihe/wechat-sdk"
```

[v2版本API文档](https://github.com/pyihe/wechat-sdk/blob/v2/README.md)

#### V3版本微信API

```go
// get go mod
go get github.com/pyihe/wechat-sdk/v3@v3.0.0

// import in project
import "github.com/pyihe/wechat-sdk/v3"
```

### Usage

**NOTE:**

1. 考虑v3版本的微信API参数较多且存在很多嵌套的结构体, 对于只需要body参数的API, 本package统统使用interface{}作为函数形参，请调用者自己构造可序列化的参数体(
   结构体、map或者序列化好了的json字符串或者字节切片[]byte)!
2. 对于非interface{}的形参, 调用者只需要按照函数规定传参即可!
3. 对于不同功能但应答参数相似的API(如预支付API和支付查询API等), 本package为了避免重复声明接收API应答参数的结构体, 最终使用了公共的结构体, 调用者处理API返回结果时,
   请严格参考 [微信官方文档](https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml) 忽略掉文档没有的参数!
4. 对于微信异步回调回来的通知, 本package会将通知结果反序列化至对应的应答结构体, 请调用者根据结果处理自己的业务逻辑, 并在处理完成后一定告知微信服务器!

```go
package main

import (
	"github.com/pyihe/wechat-sdk/v3/service"
	"github.com/pyihe/wechat-sdk/v3/service/mini"
	"github.com/pyihe/wechat-sdk/v3/service/official"
	"github.com/pyihe/wechat-sdk/v3/service/payment/merchant"
)

// WechatConfig 微信相关参数配置
var WechatConfig struct {
	AppId          string `json:"appid"`       // 应用ID
	MchId          string `json:"mchid"`       // 商户号
	SerialNo       string `json:"serial_no"`   // 商户证书序列号
	Apikey         string `json:"apikey"`      // API v3 Key
	PrivateKeyPath string `json:"private_key"` // 商户平台私钥文件
	PublicKeyPath  string `json:"public_key"`  // 商户微信支付平台公钥文件
}

// PayBody 支付请求参数
type PayBody struct {
	AppId       string  `json:"appid"`        // 应用ID
	MchId       string  `json:"mchid"`        // 商户号
	Description string  `json:"description"`  // 商品描述
	OutTradeNo  string  `json:"out_trade_no"` // 商户订单号
	NotifyUrl   string  `json:"notify_url"`   // 支付异步回调通知
	Amount      *Amount `json:"amount"`       // 订单金额信息
	Payer       *Payer  `json:"payer"`        // 支付者
}

type Amount struct {
	Total    int64  `json:"total"`              // 订单总金额
	Currency string `json:"currency,omitempty"` // 货币类型
}

type Payer struct {
	OpenId string `json:"openid"` // 用户标识
}

func handleErr(err error) {
	//...
}

func main() {
	var opts = []service.Option{
		service.WithAppId(WechatConfig.AppId),
		service.WithMchId(WechatConfig.MchId),
		service.WithApiV3Key(WechatConfig.Apikey),
		service.WithSerialNo(WechatConfig.SerialNo),
	}
	var srvConfig = service.NewConfig(opts...)

	var param = &PayBody{
		AppId:       "wxdace645e0bc2cXXX",
		MchId:       "1900006XXX",
		Description: "Image形象店-深圳腾大-QQ公仔",
		OutTradeNo:  "1217752501201407033233368318",
		NotifyUrl:   "https://weixin.qq.com/",
		Amount: &Amount{
			Total:    1,
			Currency: "CNY",
		},
		Payer: &Payer{OpenId: "o4GgauInH_RCEdvrrNGrntXDuXXX"},
	}
	// 普通商户支付
	merchantResponse, err := merchant.JSAPI(srvConfig, param)
	if err != nil {
		handleErr(err)
	}
	if err = merchantResponse.Error(); err != nil {
		handleErr(err)
	}
	// 小程序获取用户openid
	miniData, err := mini.GetOpenId(srvConfig, "your jsCode")
	if err != nil {
		handleErr(err)
	}
	// 公众号获取用户openid
	officialData, err := official.GetOpenId(srvConfig, "your grant code")
	if err != nil {
		handleErr(err)
	}
}
```

### TODO

- [x] [账单申请及下载(商户、服务商)](https://github.com/pyihe/wechat-sdk/tree/master/service/bills)
- [x] [智慧商圈(商户、服务商)](https://github.com/pyihe/wechat-sdk/tree/master/service/businesscircle)
- [x] [证书下载](https://github.com/pyihe/wechat-sdk/tree/master/service/certificate)
- [x] [合单支付(商户、服务商)](https://github.com/pyihe/wechat-sdk/tree/master/service/combine)
- [x] [代金券(商户、服务商)](https://github.com/pyihe/wechat-sdk/tree/master/service/favor)
- [x] 基础支付( [商户](https://github.com/pyihe/wechat-sdk/tree/master/service/merchant)
  、[服务商](https://github.com/pyihe/wechat-sdk/tree/master/service/partner))
- [x] [小程序](https://github.com/pyihe/wechat-sdk/tree/master/service/mini)
- [x] [微信公众号](https://github.com/pyihe/wechat-sdk/tree/master/service/official)
- [x] [支付分停车服务(商户、服务商)](https://github.com/pyihe/wechat-sdk/tree/master/service/parking)
- [x] [支付分(商户)](https://github.com/pyihe/wechat-sdk/tree/master/service/payscore)
- [x] [退款(商户、服务商)](https://github.com/pyihe/wechat-sdk/tree/master/service/refunds)
- [x] [支付即服务(商户、服务商)](https://github.com/pyihe/wechat-sdk/tree/master/service/smartguide)
- [x] [商家券(商户、服务商)]()
- [x] [委托营销(商户、服务商)]()
- [x] [消费卡(商户)]()
- [x] [分账(商户、服务商)]()
- [x] [图片上传(营销专用)]()
- [x] [支付有礼(商户、服务商)]()
- [x] [消费者投诉2.0(商户、服务商)]()
- [x] [其他能力(图片上传、视频上传)]()
- [x] [特约商户进件(服务商)]()
- [x] [点金计划(服务商)]()
- [x] [商户开户意愿确认(服务商)]()
- [x] [商户违规通知(服务商)]()
- [x] [连锁品牌分账(服务商)]()
- [ ] 电商收付通(服务商)
- [ ] **付款码支付(官方尚未升级)**
- [ ] **现金红包(官方尚未升级)**
- [ ] **付款(官方尚未升级)**
- [ ] **海关报关(官方尚未升级)**

### 致谢

感谢[Jetbrains开源开发许可证](https://www.jetbrains.com/zh-cn/community/opensource/#support) 提供的免费开发工具支持!

<img src="https://github.com/pyihe/wechat-sdk/blob/master/jetbrains.png" width="50" height="50"/>

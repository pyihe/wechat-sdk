## 《商家券》相关功能

|Name|Function|
|:----|:----|
|创建商家券|[CreateStock](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L15)|
|查询商家券详情|[QueryMerchantStock](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L36)|
|核销用户券|[UseCoupon](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L57)|
|根据过滤条件查询用户券|[QueryUserCouponsByFilter](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L78)|
|查询用户单张券详情|[QueryUserCoupon](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L120)|
|上传预存code|[UploadCouponCode](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L138)|
|设置商家券事件通知URL地址|[SetCallbacks](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L163)|
|查询商家券事件通知URL地址|[QueryCallbacks](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L184)|
|关联订单信息|[Associate](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L208)|
|取消关联订单信息|[Disassociate](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L229)|
|修改批次预算|[ModifyStockBudget](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L250)|
|修改商家券基本信息|[ModifyStock](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L275)|
|申请退券|[ReturnCoupon](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L300)|
|使券失效|[DeactivateCoupon](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L321)|
|补差付款|[SubsidyPay](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L342)|
|查询营销补差付款单详情|[QuerySubsidyPayReceipt](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L363)|
|发放消费卡|[SendCoupon](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L383)|
|解析领券事件通知|[ParseReceiveCouponNotify](https://github.com/pyihe/wechat-sdk/blob/master/service/busifavor/busifavor.go#L406)|
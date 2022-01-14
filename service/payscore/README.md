## 《支付分》相关功能

|Name|Function|
|:----|:----|
|商户预授权|[PrePermit](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L15)|
|查询用户授权记录|[QueryPermissions](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L36)|
|解除用户授权关系|[TerminatePermission](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L79)|
|解析开启/解除授权服务回调通知|[ParseOpenOrCloseNotify](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L122)|
|解析确认订单通知内容|[ParseConfirmOrderNotify](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L134)|
|创建支付订单|[CreateServiceOrder](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L146)|
|查询支付分订单|[QueryServiceOrder](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L166)|
|取消支付分订单|[CancelServiceOrder](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L208)|
|修改订单金额|[ModifyServiceOrder](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L245)|
|完结支付分订单|[CompleteServiceOrder](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L269)|
|商户发起催收扣款|[PayServiceOrder](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L293)|
|同步服务订单信息|[SyncServiceOrder](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L317)|
|解析支付成功回调数据|[ParsePaymentNotify](https://github.com/pyihe/wechat-sdk/blob/master/service/payscore/payscore.go#L341)|
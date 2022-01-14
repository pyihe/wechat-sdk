## 《微信支付分停车服务》相关功能

|Name|Function|
|:----|:----|
|查询车牌服务开通信息|[FindParkingService](https://github.com/pyihe/wechat-sdk/blob/master/service/parking/parking.go#L16)|
|创建停车入场|[CreateParking](https://github.com/pyihe/wechat-sdk/blob/master/service/parking/parking.go#L63)|
|扣费受理|[TransactionsParking](https://github.com/pyihe/wechat-sdk/blob/master/service/parking/parking.go#L84)|
|查询订单|[QueryOrder](https://github.com/pyihe/wechat-sdk/blob/master/service/parking/parking.go#L105)|
|解析停车入场状态变更通知结果|[ParseParkingStateNotify](https://github.com/pyihe/wechat-sdk/blob/master/service/parking/parking.go#L134)|
|解析支付通知结果|[ParsePaymentNotify](https://github.com/pyihe/wechat-sdk/blob/master/service/parking/parking.go#L147)|
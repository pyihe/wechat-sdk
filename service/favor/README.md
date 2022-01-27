## 《代金券》相关功能

|Name|Function|
|:----|:----|
|创建代金券批次|[CreateStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L21)|
|激活代金券批次|[StartStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L42)|
|发放代金券批次|[SendStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L61)|
|暂停代金券|[PauseStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L86)|
|重启代金券批次|[RestartStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L106)|
|条件查询批次列表|[QueryStockList](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L125)|
|查询批次详情|[QueryStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L187)|
|查询代金券详情|[QueryCoupon](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L228)|
|查询代金券可用商户|[QueryStockMerchants](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L247)|
|查询代金券可用单品列表|[QueryStockItems](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L274)|
|根据商户号查询用户的券|[QueryUserCoupons](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L301)|
|下载批次核销明细|[DownloadStockUseFlow](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L345)|
|下载批次退款明细|[DownloadStockRefundFlow](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L395)|
|设置消息通知地址|[SetCallbacks](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L451)|
|解析核销事件回调通知|[ParseUseNotify](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L473)|
|图片上传(营销专用)|[UploadImage](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L487)|
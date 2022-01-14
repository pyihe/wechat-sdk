## 《代金券》相关功能

|Name|Function|
|:----|:----|
|创建代金券批次|[CreateStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L22)|
|激活代金券批次|[StartStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L43)|
|发放代金券批次|[SendStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L70)|
|暂停代金券|[PauseStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L95)|
|重启代金券批次|[RestartStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L123)|
|条件查询批次列表|[QueryStockList](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L150)|
|查询批次详情|[QueryStock](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L220)|
|查询代金券详情|[QueryCoupon](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L269)|
|查询代金券可用商户|[QueryStockMerchants](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L300)|
|查询代金券可用单品列表|[QueryStockItems](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L338)|
|根据商户号查询用户的券|[QueryUserCoupons](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L376)|
|下载批次核销明细|[DownloadStockUseFlow](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L431)|
|下载批次退款明细|[DownloadStockRefundFlow](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L481)|
|设置消息通知地址|[SettingCallbacks](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L537)|
|解析核销事件回调通知|[ParseUseNotify](https://github.com/pyihe/wechat-sdk/blob/master/service/favor/favor.go#L559)|
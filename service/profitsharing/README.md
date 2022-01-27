## 《分账》相关功能

|Name|Function|
|:----|:----|
|请求分账|[CreateSharing](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L17)|
|查询分账结果|[QuerySharing](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L38)|
|请求分账回退|[ReturnSharing](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L61)|
|查询分账回退结果|[QueryReturnSharing](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L82)|
|解冻剩余资金|[Unfreeze](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L105)|
|查询订单剩余待分账金额|[QueryUnSharingAmount](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L126)|
|查询最大分账比例|[QueryMaxSharingRatio](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L143)|
|添加分账接收方|[AddSharingReceiver](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L161)|
|删除分账接收方|[DeleteSharingReceiver](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L182)|
|解析分账动帐通知|[ParseSharingNotify](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L203)|
|下载分账账单|[DownloadSharingBills](https://github.com/pyihe/wechat-sdk/blob/master/service/profitsharing/profitsharing.go#L216)|
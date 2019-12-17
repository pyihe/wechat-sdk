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
| UnifiedMicro | 扫码下单 |  |
| UnifiedOrder | 统一下单: H5/APP/MWEB/NATIVE |  |
| UnifiedQuery | 下单结果查询 |  |

##### 如何使用: 
```


```


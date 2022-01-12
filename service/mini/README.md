## 《小程序》相关功能

|API|Function|
|:---------|:-----------|
|获取基础调用access_token|[GetBaseAccessToken](https://github.com/pyihe/wechat-sdk/blob/v3/v3/service/mini_program/mini.go#L22)|
|小程序登录授权时获取用户openid和session_key|[GetOpenId](https://github.com/pyihe/wechat-sdk/blob/v3/v3/service/mini_program/mini.go#L63)|
|校验加密信息是否由微信生成(只支持手机号加密数据且只能检测最近3天加密的数据)|[CheckEncryptData](https://github.com/pyihe/wechat-sdk/blob/v3/v3/service/mini_program/mini.go#L97)|
|解密小程序的敏感数据，如用户信息、手机号码等|[DecryptOpenData](https://github.com/pyihe/wechat-sdk/blob/v3/v3/service/mini_program/mini.go#L153)|
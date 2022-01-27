### 《公众号》相关功能

|API|Function|
|:---------|:-----------|
|获取基础调用access_token|[GetBaseAccessToken](https://github.com/pyihe/wechat-sdk/blob/master/service/official/official.go#L19)|
|获取用户openid以及网页授权access_token|[GetOpenId](https://github.com/pyihe/wechat-sdk/blob/master/service/official/official.go#L72)|
|刷新网页授权access_token|[RefreshOauthAccessToken](https://github.com/pyihe/wechat-sdk/blob/master/service/official/official.go#L124)|
|获取用户信息|[GetUserInfo](https://github.com/pyihe/wechat-sdk/blob/master/service/official/official.go#L167)|
|校验网页授权access_token是否有效|[CheckOauthAccessTokenValid](https://github.com/pyihe/wechat-sdk/blob/master/service/official/official.go#L195)|
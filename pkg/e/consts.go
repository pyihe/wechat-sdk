package e

const (
	SignTypeMD5        = "MD5"
	SignType256        = "HMAC-SHA256"
	UnifiedOrderApiUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	QueryOrderApiUrl   = "https://api.mch.weixin.qq.com/pay/orderquery"

	H5     = "MWEB"   //H5支付
	App    = "APP"    //APP支付
	JSAPI  = "JSAPI"  //JSAPI支付
	Native = "NATIVE" //Native(扫码)支付
)

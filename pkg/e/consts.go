package e

const (
	PostContentType    = "application/xml;charset=utf-8"
	PublicKey          = "public.pem" //企业转账到银行卡的AES公钥保存文件名
	SignTypeMD5        = "MD5"
	SignType256        = "HMAC-SHA256"
	UnifiedOrderApiUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	QueryOrderApiUrl   = "https://api.mch.weixin.qq.com/pay/orderquery"

	H5     = "MWEB"   //H5支付
	App    = "APP"    //APP支付
	JSAPI  = "JSAPI"  //JSAPI支付
	Native = "NATIVE" //Native(扫码)支付
)

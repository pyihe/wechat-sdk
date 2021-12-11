package vars

const (
	_BeginTradeType TradeType = iota // begin
	JSAPI                            // JSAPI支付
	H5                               // H5支付
	Native                           // Native支付
	APP                              // App支付
	FacePay                          // 刷脸支付
	_EndTradeType                    // end
)

const (
	_BeginPlatform Platform = iota // begin
	Merchant                       // 服务商平台
	Partner                        // 商户平台
	_EndPlatform                   // end
)

const (
	ContentTypeJSON = "application/json"
	PostContentType = "application/xml;charset=utf-8"
)

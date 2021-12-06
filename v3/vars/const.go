package vars

const (
	JSAPI  TradeType = "JSAPI"
	H5     TradeType = "H5"
	Native TradeType = "Native"
	APP    TradeType = "APP"

	QueryOutTradeNo    QueryType = "out-trade-no"
	QueryTransactionId QueryType = "transactions"

	// 内部API类型
	apiTypePrepay              = "prepay"               // 预支付
	apiTypePrepayNotify        = "prepay-notify"        // 支付通知
	apiTypeOrderQuery          = "order-query"          // 订单查询
	apiTypeOrderClose          = "order-close"          // 关闭订单
	apiTypeRefund              = "refund"               // 查询单笔退款订单
	apiTypeRefundQuery         = "refund_query"         // 退款查询
	apiTypeCertificateDownload = "certificate-download" // 证书下载

)

const (
	ContentTypeJSON = "application/json"
	PostContentType = "application/xml;charset=utf-8"
)

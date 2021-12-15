package vars

// Platform 平台类型: 服务商平台、商户平台
type Platform int

func (p Platform) Valid() (ok bool) {
	return p > _BeginPlatform && p < _EndPlatform
}

// TradeType 交易类型
type TradeType int

func (t TradeType) Valid() (ok bool) {
	return t > _BeginTradeType && t < _EndTradeType
}

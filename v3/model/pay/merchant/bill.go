package merchant

import "github.com/pyihe/go-pkg/errors"

type TradeBillRequest struct {
	BillDate string `json:"bill_date,omitempty"` // 账单日期，格式为YYYY-MM-DD
	BillType string `json:"bill_type,omitempty"` // 账单类型，ALL: 返回当日所有订单信息; SUCCESS: 返回当日成功支付的订单; REFUND: 返回当日退款订单
	TarType  string `json:"tar_type,omitempty"`  // 压缩类型
}

func (t *TradeBillRequest) Check() (err error) {
	if t.BillDate == "" {
		err = errors.New("请填写账单日期!")
		return
	}
	return
}

type FundFlowRequest struct {
	BillDate    string `json:"bill_date,omitempty"`    // 账单日期
	AccountType string `json:"account_type,omitempty"` // 资金账户类型
	TarType     string `json:"tar_type,omitempty"`     // 压缩类型
}

func (f *FundFlowRequest) Check() (err error) {
	if f.BillDate == "" {
		err = errors.New("请填写账单日期!")
		return
	}
	return
}

type BillResponse struct {
	RequestId   string `json:"-"`                      // 唯一请求ID
	HashType    string `json:"hash_type,omitempty"`    // 原始账单的摘要值类型
	HashValue   string `json:"hash_value,omitempty"`   // 摘要值
	DownloadUrl string `json:"download_url,omitempty"` // 账单下载地址
}

package bills

import (
	"fmt"
	"testing"

	"github.com/pyihe/wechat-sdk/v3/service/tests"
)

func TestDownloadTradeBill(t *testing.T) {
	request := &TradeBillRequest{
		BillDate: "2021-11-30",
		BillType: "ALL",
		FileName: "bill.xlsx",
		FilePath: "./files",
	}
	response, err := DownloadTradeBill(tests.Config, request)
	if err != nil {
		t.Logf("err: %v\n", err)
		return
	}
	t.Logf("response: %+v\n", *response)
}

func TestDownloadFundFlowBill(t *testing.T) {
	request := &FundFlowRequest{
		BillDate:    "2021-11-30",
		AccountType: "BASIC",
		FileName:    "fund_flow.xlsx",
		FilePath:    "./files",
	}
	response, err := DownloadFundFlowBill(tests.Config, request)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	t.Logf("response: %+v\n", *response)
}

func TestDownloadSubMerchantFundFlowBill(t *testing.T) {
	request := &SubMerchantFundFlowRequest{
		SubMchId:    "19000000001",
		BillDate:    "2021-12-31",
		AccountType: "BASIC",
		Algorithm:   "AEAD_AES_256_GCM",
		FileName:    "fund_flow.xlsx",
		FilePath:    "./files",
	}
	response, err := DownloadSubMerchantFundFlowBill(tests.Config, request)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	t.Logf("response: %+v\n", *response)
}

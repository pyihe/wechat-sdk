package v3

import (
	"fmt"
	"testing"

	"github.com/pyihe/wechat-sdk/v3/vars"

	"github.com/pyihe/go-pkg/rands"
)

var (
	appId       = "wx7224e0425e3b8655"
	mchId       = "1491203582"
	serialNo    = "1BACB027A363F9975BD335B62EA65229E68C4BD1"
	apiKey      = "r3WoXhjKdcdsQFcQ55y5WrfRIbi8AO3l"
	privatePath = "./pem/apiclient_key.pem"
	openId      = "oVqjJ0jDoXwuNCkgWDmSMbI15eEQ"
	publicPath  = "./pem/public_key_2026_11_07.pem"
)

var (
	opts = []Options{
		WithMchId(mchId),
		WithV3Key(apiKey),
		WithSerialNo(serialNo),
		WithPrivateKey(privatePath),
		WithPublicKey(publicPath),
	}
	weClient = NewWechatPayer(appId, opts...)
)

func TestWeChatClient_Prepay(t *testing.T) {
	outTradeNo := rands.String(32)
	t.Logf("订单号为: %v\n", outTradeNo)
	request := &vars.PrepayRequest{
		TradeType:   vars.JSAPI,
		AppId:       appId,
		MchId:       mchId,
		Description: "API TEST",
		OutTradeNo:  outTradeNo,
		NotifyUrl:   "https://8654-125-69-40-247.ngrok.io/notify",
		Amount: &vars.Amount{
			Total:    1,
			Currency: "CNY",
		},
		Payer: &vars.Payer{OpenId: openId},
	}
	result, err := weClient.Prepay(request)
	if err != nil {
		t.Logf("err: %v", err)
		return
	}
	t.Logf("%+v\n", *result)
}

func TestWeChatClient_QueryOrder(t *testing.T) {
	var queryId = "10rRFGJAPvaW3t9CVBCcxSKKXOT7zprc"
	queryResponse, err := weClient.QueryOrder(vars.QueryOutTradeNo, queryId)
	if err != nil {
		t.Logf("err: %v\n", err)
		return
	}
	t.Logf("%+v\n", *queryResponse)
}

func TestWeChatClient_Close(t *testing.T) {
	var outTradeNO = "10rRFGJAPvaW3t9CVBCcxSKKXOT7zprc"
	requestId, err := weClient.CloseOrder(outTradeNO)
	if err != nil {
		t.Logf("err: %v\n", err)
		return
	}
	t.Logf("requestId: %v\n", requestId)
}

func TestWeChatClient_Refund(t *testing.T) {
	var outTradeNO = "QMzaFGhvhGsxM9GiJg5HBPXZuSyS9dkH"
	var outRefundNo = rands.String(32)
	t.Logf("outRefundNo: %v\n", outRefundNo)
	var request = &vars.RefundRequest{
		OutTradeNo:  outTradeNO,
		OutRefundNo: outRefundNo,
		Amount: &vars.Amount{
			Total:    1,
			Refund:   1,
			Currency: "CNY",
		},
	}
	result, err := weClient.Refund(request)
	if err != nil {
		t.Logf("err: %v\n", err)
		return
	}
	t.Logf("%+v\n", *result)
}

func TestWeChatClient_DownloadCertificates(t *testing.T) {
	param, err := weClient.DownloadCertificates("")
	fmt.Println(param)
	fmt.Println(err)
}

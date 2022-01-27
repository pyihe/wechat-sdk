package apply4sub

import (
	"github.com/pyihe/wechat-sdk/v3/service"
	"testing"
)

var (
	config     *service.Config
	mchId      = "**********"
	appId      = "******************"
	apiV3Key   = "********************************"
	serialNo   = "****************************************"
	publicKey  = "./pem/public_key.pem"
	privateKey = "./pem/apiclient_key.pem"
)

func init() {
	opts := []service.Option{
		service.WithAppId(appId),
		service.WithMchId(mchId),
		service.WithApiV3Key(apiV3Key),
		service.WithPrivateKey(privateKey),
		service.WithPublicKey(publicKey),
		service.WithSerialNo(serialNo),
	}
	config = service.NewConfig(opts...)
}

func TestApply(t *testing.T) {
	request := &ApplyRequest{
		BusinessCode: "1900013511_10000",
		ContactInfo: &ContactInfo{
			ContactName:     "Kevin",
			ContactIdNumber: "12345678901x",
			OpenId:          "oUpF8uMuAJO_M2pxb1Q9zNjWeS6o",
			MobilePhone:     "12345678901",
			ContactEmail:    "xxx@test.com",
		},
		SubjectInfo: &SubjectInfo{
			SubjectType: "SUBJECT_TYPE_ENTERPRISE",
			IdentityInfo: &IdentityInfo{
				IdDocType: "IDENTIFICATION_TYPE_IDCARD",
				Owner:     true,
			},
		},
		BusinessInfo: &BusinessInfo{
			MerchantShortname: "张三餐饮店",
			ServicePhone:      "0758XXXXX",
			SalesInfo: &SalesInfo{
				SalesScenesType: []string{"SALES_SCENES_STORE"},
			},
		},
		SettlementInfo: &SettlementInfo{
			SettlementId:      "719",
			QualificationType: "餐饮",
		},
		BankAccountInfo: &BankAccountInfo{
			BankAccountType: "BANK_ACCOUNT_TYPE_CORPORATE",
			AccountName:     "李四",
			AccountBank:     "工商银行",
			BankAddressCode: "110000",
			BankBranchId:    "402713354941",
			AccountNumber:   "11111111111111111111111111111111",
		},
	}
	response, err := Apply(config, request)
	if err != nil {
		t.Logf("apply err: %v\n", err)
		return
	}
	if err = response.Error(); err != nil {
		t.Logf("response err: %v\n", err)
		return
	}
	t.Logf("response: %+v\n", *response)
}

func TestQueryApplyment(t *testing.T) {
	request := &QueryApplymentRequest{
		//ApplymentId:  2000001234567890,
		BusinessCode: "1900013511_10000",
	}
	response, err := QueryApplyment(config, request)
	if err != nil {
		t.Logf("apply err: %v\n", err)
		return
	}
	if err = response.Error(); err != nil {
		t.Logf("response err: %v\n", err)
		return
	}
	t.Logf("response: %+v\n", *response)
}

func TestModifySettlement(t *testing.T) {
	request := &ModifySettlementRequest{
		SubMchId:        "xxxxxxxx",
		AccountType:     "ACCOUNT_TYPE_BUSINESS",
		AccountBank:     "工商银行",
		BankAddressCode: "110000",
		BankName:        "施秉县农村信用合作联社城关信用社",
		BankBranchId:    "402713354941",
		AccountNumber:   "123456789111199992229999",
	}
	response, err := ModifySettlement(config, request)
	if err != nil {
		t.Logf("apply err: %v\n", err)
		return
	}
	if err = response.Error(); err != nil {
		t.Logf("response err: %v\n", err)
		return
	}
	t.Logf("response: %+v\n", *response)
}

func TestQuerySettlement(t *testing.T) {
	response, err := QuerySettlement(config, "")
	if err != nil {
		t.Logf("apply err: %v\n", err)
		return
	}
	if err = response.Error(); err != nil {
		t.Logf("response err: %v\n", err)
		return
	}
	t.Logf("response: %+v\n", *response)
}

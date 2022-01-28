package apply4subject

import (
	"testing"

	"github.com/pyihe/wechat-sdk/v3/service/tests"
)

func TestApply(t *testing.T) {
	request := &ApplyRequest{
		ChannelId:    "",
		BusinessCode: "1111111111",
		ContactInfo: &ContactInfo{
			Name:         "李三",
			Mobile:       "12345678911",
			IdCardNumber: "1111111111111111111111111111111",
		},
		SubjectInfo: &SubjectInfo{
			SubjectType: "SUBJECT_TYPE_ENTERPRISE",
		},
		IdentificationInfo: &IdentificationInfo{
			IdentificationType:      "IDENTIFICATION_TYPE_IDCARD",
			IdentificationName:      "李三",
			IdentificationNumber:    "1111111111111111111111111111111",
			IdentificationValidDate: "[\"1970-01-01\",\"forever\"]",
			IdentificationFrontCopy: "0P3ng6KTIW4-Q_l2FjKLZuhHjBWoMAjmVtCz7ScmhEIThCaV-4BBgVwtNkCHO_XXqK5dE5YdOmFJBZR9FwczhJehHhAZN6BKXQPcs-VvdSo",
		},
	}
	response, err := Apply(tests.Config, request)
	if err != nil {
		t.Logf("err: %v\n", err)
		return
	}

	t.Logf("response: %+v\n", *response)
}

func TestCancelApplyment(t *testing.T) {
	request := &CancelRequest{
		//ApplymentId:  20000011111,
		BusinessCode: "1900013511_10000",
	}
	response, err := CancelApplyment(tests.Config, request)
	if err != nil {
		t.Logf("err: %v\n", err)
		return
	}

	t.Logf("response: %+v\n", *response)
}

func TestQueryApplyResult(t *testing.T) {
	request := &QueryApplyResultRequest{
		ApplymentId: 20000011111,
		//BusinessCode: "1900013511_10000",
	}
	response, err := QueryApplyResult(tests.Config, request)
	if err != nil {
		t.Logf("err: %v\n", err)
		return
	}

	t.Logf("response: %+v\n", *response)
}

func TestQueryMerchantState(t *testing.T) {
	response, err := QueryMerchantState(tests.Config, "*************")
	if err != nil {
		t.Logf("err: %v\n", err)
		return
	}

	t.Logf("response: %+v\n", *response)
}

package wechat_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pyihe/secret"
	"github.com/pyihe/util/certs"
	"github.com/pyihe/util/utils"
	"github.com/pyihe/wechat-sdk/pkg"
)

func (m *myPayer) ProfitSharing(param Param, p12CertPath string, multiTag bool) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	//读取证书
	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		shareMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "transaction_id", "out_order_no", "receivers"}
		shareOptionalParam = []string{"sign_type"}
		signType           = pkg.SignType256
	)

	if t, ok := param["sign_type"]; ok {
		if t.(string) != pkg.SignType256 {
			return nil, errors.New("only support HMAC-SHA256")
		}
	}

	for _, v := range shareMustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}
	for key := range param {
		if !utils.Contain(shareMustParam, key) && !utils.Contain(shareOptionalParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var url = "https://api.mch.weixin.qq.com/secapi/pay/profitsharing"
	if multiTag {
		url = "https://api.mch.weixin.qq.com/secapi/pay/multiprofitsharing"
	}
	var request = &postRequest{
		Body:        reader,
		Url:         url,
		ContentType: pkg.PostContentType,
	}
	response, err := postToWxWithCert(request, cert)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return nil, errors.New(returnMsg)
	}
	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return nil, errors.New(errDes)
	}
	sign = result.Sign(signType)
	if resultSign, _ := result.GetString("sign"); resultSign != sign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) QueryProfitSharing(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	param.Add("mch_id", m.mchId)

	var (
		mustParam     = []string{"mch_id", "nonce_str", "sign", "transaction_id", "out_order_no"}
		optionalParam = []string{"sign_type"}
		signType      = pkg.SignType256
	)
	if t, ok := param["sign_type"]; ok {
		if t.(string) != pkg.SignType256 {
			return nil, errors.New("only support HMAC-SHA256")
		}
	}
	for _, v := range mustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}

	for key := range param {
		if !utils.Contain(mustParam, key) && !utils.Contain(optionalParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/profitsharingquery",
		ContentType: pkg.PostContentType,
	}
	response, err := postToWx(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return nil, errors.New(returnMsg)
	}

	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return nil, errors.New(errDes)
	}
	sign = result.Sign(signType)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) AddProfitSharingReceiver(param Param) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		mustParam     = []string{"mch_id", "appid", "nonce_str", "sign", "receiver"}
		optionalParam = []string{"sign_type"}
		signType      = pkg.SignType256
	)
	if t, ok := param["sign_type"]; ok {
		if t.(string) != pkg.SignType256 {
			return nil, errors.New("only support HMAC-SHA256")
		}
	}
	for _, v := range mustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}

	for key := range param {
		if !utils.Contain(mustParam, key) && !utils.Contain(optionalParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/profitsharingaddreceiver",
		ContentType: pkg.PostContentType,
	}
	response, err := postToWx(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return nil, errors.New(returnMsg)
	}

	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return nil, errors.New(errDes)
	}
	sign = result.Sign(signType)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) RemoveProfitSharingReceiver(param Param) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		mustParam     = []string{"mch_id", "appid", "nonce_str", "sign", "receiver"}
		optionalParam = []string{"sign_type"}
		signType      = pkg.SignType256
	)
	if t, ok := param["sign_type"]; ok {
		if t.(string) != pkg.SignType256 {
			return nil, errors.New("only support HMAC-SHA256")
		}
	}
	for _, v := range mustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}

	for key := range param {
		if !utils.Contain(mustParam, key) && !utils.Contain(optionalParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/profitsharingremovereceiver",
		ContentType: pkg.PostContentType,
	}
	response, err := postToWx(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return nil, errors.New(returnMsg)
	}

	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return nil, errors.New(errDes)
	}
	sign = result.Sign(signType)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) FinishProfitSharing(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	//读取证书
	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		mustParam     = []string{"mch_id", "appid", "nonce_str", "sign", "transaction_id", "out_order_no", "description"}
		optionalParam = []string{"sign_type"}
		signType      = pkg.SignType256
	)

	if t, ok := param["sign_type"]; ok {
		if t.(string) != pkg.SignType256 {
			return nil, errors.New("only support HMAC-SHA256")
		}
	}

	for _, v := range mustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}
	for key := range param {
		if !utils.Contain(mustParam, key) && !utils.Contain(optionalParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/secapi/pay/profitsharingfinish",
		ContentType: pkg.PostContentType,
	}
	response, err := postToWxWithCert(request, cert)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return nil, errors.New(returnMsg)
	}
	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return nil, errors.New(errDes)
	}
	sign = result.Sign(signType)
	if resultSign, _ := result.GetString("sign"); resultSign != sign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) ReturnProfitSharing(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	//读取证书
	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		mustParam     = []string{"mch_id", "appid", "nonce_str", "sign", "out_return_no", "return_account_type", "return_account", "return_amount", "description"}
		oneParam      = []string{"order_id", "out_order_no"}
		optionalParam = []string{"sign_type"}
		signType      = pkg.SignType256
	)

	if t, ok := param["sign_type"]; ok {
		if t.(string) != pkg.SignType256 {
			return nil, errors.New("only support HMAC-SHA256")
		}
	}

	var count = 0
	for _, k := range oneParam {
		if v := param.Get(k); v != nil {
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("need order number: order_id or out_order_no")
	} else if count > 1 {
		return nil, errors.New("just one order number: order_id or out_order_no")
	}

	for _, v := range mustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}
	for key := range param {
		if !utils.Contain(mustParam, key) && !utils.Contain(optionalParam, key) && !utils.Contain(oneParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/secapi/pay/profitsharingreturn",
		ContentType: pkg.PostContentType,
	}
	response, err := postToWxWithCert(request, cert)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return nil, errors.New(returnMsg)
	}
	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return nil, errors.New(errDes)
	}
	sign = result.Sign(signType)
	if resultSign, _ := result.GetString("sign"); resultSign != sign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) QueryProfitSharingReturn(param Param) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		mustParam     = []string{"mch_id", "appid", "nonce_str", "sign", "out_return_no"}
		oneParam      = []string{"order_id", "out_order_no"}
		optionalParam = []string{"sign_type"}
		signType      = pkg.SignType256
	)
	if t, ok := param["sign_type"]; ok {
		if t.(string) != pkg.SignType256 {
			return nil, errors.New("only support HMAC-SHA256")
		}
	}

	var count = 0
	for _, k := range oneParam {
		if v := param.Get(k); v != nil {
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("need order number: order_id or out_order_no")
	} else if count > 1 {
		return nil, errors.New("just one order number: order_id or out_order_no")
	}

	for _, v := range mustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}

	for key := range param {
		if !utils.Contain(mustParam, key) && !utils.Contain(optionalParam, key) && !utils.Contain(oneParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/profitsharingreturnquery",
		ContentType: pkg.PostContentType,
	}
	response, err := postToWx(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return nil, errors.New(returnMsg)
	}

	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return nil, errors.New(errDes)
	}
	sign = result.Sign(signType)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

func (m *myPayer) ProfitSharingNotify(body io.Reader) (ResultParam, error) {
	if len(m.apiV3Key) == 0 {
		return nil, errors.New("no api v3 key")
	}
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var data *profitSharingResult
	if err = json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("read from body fail")
	}

	var p = newResultMap()
	p.Add("id", data.Id)
	p.Add("create_time", data.CreateTime)
	p.Add("event_type", data.EventType)
	p.Add("summary", data.Summary)
	p.Add("resource_type", data.ResourceType)
	p.Add("algorithm", data.Resource.Algorithm)
	p.Add("original_type", data.Resource.OriginalType)
	p.Add("ciphertext", data.Resource.Ciphertext)
	p.Add("associated_data", data.Resource.AssociatedData)
	p.Add("nonce", data.Resource.Nonce)

	//对密文执行解密
	c := secret.NewCipher()
	c.SetGCMNonce([]byte(data.Resource.Nonce))
	request := &secret.SymRequest{
		CipherData: data.Resource.Ciphertext,
		Key:        []byte(m.apiV3Key),
		ModeType:   secret.BlockModeGCM,
		AddData:    []byte(data.Resource.AssociatedData),
		Type:       secret.SymTypeAES,
	}
	plainBytes, err := c.SymDecrypt(request)
	if err != nil {
		return nil, err
	}
	var plainData *plain
	if err = json.Unmarshal(plainBytes, &plainData); err != nil {
		return nil, err
	}
	p.Add("mchid", plainData.Mchid)
	p.Add("sp_mchid", plainData.SpMchid)
	p.Add("sub_mchid", plainData.SubMchid)
	p.Add("transaction_id", plainData.TransactionId)
	p.Add("order_id", plainData.OrderId)
	p.Add("out_order_no", plainData.OutOrderNo)
	p.Add("type", plainData.Receiver.Type)
	p.Add("account", plainData.Receiver.Account)
	p.Add("amount", fmt.Sprintf("%v", plainData.Receiver.Amount))
	p.Add("description", plainData.Receiver.Description)
	p.Add("success_time", plainData.Receiver.SuccessTime)
	return p, nil
}

type profitSharingResult struct {
	Id           string    `json:"id"`
	CreateTime   string    `json:"create_time"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	ResourceType string    `json:"resource_type"`
	Resource     *resource `json:"resource"`
}

type resource struct {
	Algorithm      string `json:"algorithm"`
	OriginalType   string `json:"original_type"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type plain struct {
	Mchid         string    `json:"mchid"`
	SpMchid       string    `json:"sp_mchid"`
	SubMchid      string    `json:"sub_mchid"`
	TransactionId string    `json:"transaction_id"`
	OrderId       string    `json:"order_id"`
	OutOrderNo    string    `json:"out_order_no"`
	Receiver      *receiver `json:"receiver"`
}

type receiver struct {
	Type        string `json:"type"`
	Account     string `json:"account"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	SuccessTime string `json:"success_time"`
}

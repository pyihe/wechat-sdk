package wechat_sdk

import (
	"bytes"
	"crypto"
	"errors"
	"fmt"

	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/pyihe/secret"
	"github.com/pyihe/util/certs"
	"github.com/pyihe/util/files"
	"github.com/pyihe/util/utils"
	"github.com/pyihe/wechat-sdk/pkg"
)

//统一下单
func (m *myPayer) UnifiedOrder(param Param) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//获取交易类型和签名类型
	var (
		unifiedMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "body", "out_trade_no", "total_fee", "spbill_create_ip", "notify_url", "trade_type"}
		unifiedOptionalParam = []string{"device_info", "sign_type", "detail", "attach", "fee_type", "time_start", "time_expire", "goods_tag", "limit_pay", "receipt", "openid", "product_id", "scene_info", "profit_sharing"}
		tradeType            string
		signType             = pkg.SignTypeMD5 //默认MD5签名方式
	)

	if t, ok := param["trade_type"]; ok {
		tradeType = t.(string)
	} else {
		return nil, pkg.ErrTradeType
	}
	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	switch tradeType {
	case pkg.H5:
		//H5支付必须要传scene_info参数
		if sceneInfo := param.Get("scene_info"); sceneInfo == nil || sceneInfo.(string) == "" {
			return nil, errors.New("H5 pay need param scene_info")
		}
	case pkg.App:
		//App支付不需要product_id, openid, scene_info参数
		if _, ok := param["product_id"]; ok {
			return nil, errors.New("APP pay no need product_id")
		}
		if _, ok := param["openid"]; ok {
			return nil, errors.New("APP pay no need openid")
		}
		if _, ok := param["scene_info"]; ok {
			return nil, errors.New("APP pay no need scene_info")
		}
	case pkg.JSAPI:
		//JSAPI支付必须传openid参数
		if openId, ok := param["openid"]; !ok || openId.(string) == "" {
			return nil, pkg.ErrOpenId
		}
	case pkg.Native:
	default:
		return nil, errors.New("invalid trade_type")
	}
	//这里校验是否包含必传的参数
	for _, v := range unifiedMustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need " + v)
		}
	}
	//这里校验是否包含不必要的参数
	for key := range param {
		if !utils.Contain(unifiedMustParam, key) && !utils.Contain(unifiedOptionalParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(signType)
	//将签名添加到需要发送的参数里
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/unifiedorder",
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
	return result, err
}

//订单查询
func (m *myPayer) UnifiedQuery(param Param) (ResultParam, error) {
	if param == nil {
		return nil, errors.New("param is empty")
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		signType           = pkg.SignTypeMD5 //此处默认MD5
		queryMustParam     = []string{"appid", "mch_id", "nonce_str", "sign"}
		queryOneParam      = []string{"transaction_id", "out_trade_no"}
		queryOptionalParam = []string{"sign_type"}
	)
	//校验订单号
	var count = 0
	for _, k := range queryOneParam {
		if v := param.Get(k); v != nil {
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("need order number: transaction_id or out_trade_no")
	} else if count > 1 {
		return nil, errors.New("just one order number: transaction_id or out_trade_no")
	}

	//校验其他参数
	for _, k := range queryMustParam {
		if k == "sign" {
			continue
		}
		if param.Get(k) == nil {
			return nil, errors.New("need param: " + k)
		}
	}

	for k := range param {
		if k == "sign" {
			continue
		}
		if k == "sign_type" {
			signType = param[k].(string)
		}
		if !utils.Contain(queryMustParam, k) && !utils.Contain(queryOptionalParam, k) && !utils.Contain(queryOneParam, k) {
			return nil, errors.New("no need param: " + k)
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
		Url:         "https://api.mch.weixin.qq.com/pay/orderquery",
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

//扫码下单
func (m *myPayer) UnifiedMicro(param Param) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//获取交易类型和签名类型
	var (
		microMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "body", "out_trade_no", "total_fee", "spbill_create_ip", "auth_code"}
		microOptionalParam = []string{"device_info", "sign_type", "detail", "attach", "fee_type", "goods_tag", "limit_pay", "time_start", "time_expire", "receipt", "scene_info", "profit_sharing"}
		signType           = pkg.SignTypeMD5 //默认MD5签名方式
	)

	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	for _, v := range microMustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}

	for key := range param {
		if !utils.Contain(microMustParam, key) && !utils.Contain(microOptionalParam, key) {
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
		Url:         "https://api.mch.weixin.qq.com/pay/micropay",
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

//关闭订单
func (m *myPayer) CloseOrder(param Param) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		signType           = pkg.SignTypeMD5
		closeMustParam     = []string{"appid", "mch_id", "out_trade_no", "nonce_str", "sign"}
		closeOptionalParam = []string{"sign_type"}
	)

	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	for _, v := range closeMustParam {
		if v == "sign" {
			continue
		}
		if _, ok := param[v]; !ok {
			return nil, errors.New("need param: " + v)
		}
	}

	for key := range param {
		if !utils.Contain(closeMustParam, key) && !utils.Contain(closeOptionalParam, key) {
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
		Url:         "https://api.mch.weixin.qq.com/pay/closeorder",
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

//申请退款
func (m *myPayer) RefundOrder(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, errors.New("param is empty")
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
		signType         = pkg.SignTypeMD5
		refundOneParams  = []string{"transaction_id", "out_trade_no"}
		refundMustParams = []string{"appid", "mch_id", "nonce_str", "sign", "out_refund_no", "total_fee", "refund_fee"}
	)

	//校验订单号
	var count = 0
	for _, k := range refundOneParams {
		if v := param.Get(k); v != nil {
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("need order number: transaction_id or out_trade_no")
	} else if count > 1 {
		return nil, errors.New("just one order number: transaction_id or out_trade_no")
	}

	for _, k := range refundMustParams {
		if k == "sign" {
			continue
		}
		if param.Get(k) == nil {
			return nil, errors.New("need param: " + k)
		}
	}

	var refundOptionalParams = []string{"sign_type", "refund_fee_type", "refund_desc", "refund_account", "notify_url"}
	for k := range param {
		if k == "sign_type" {
			signType = param[k].(string)
		}
		if !utils.Contain(refundMustParams, k) && !utils.Contain(refundOptionalParams, k) && !utils.Contain(refundOneParams, k) {
			return nil, errors.New("no need param: " + k)
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
		Url:         "https://api.mch.weixin.qq.com/secapi/pay/refund",
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

//查询退款
func (m *myPayer) RefundQuery(param Param) (ResultParam, error) {
	if param == nil {
		return nil, errors.New("param is empty")
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		signType             = pkg.SignTypeMD5
		count                = 0
		refundQueryOneParam  = []string{"transaction_id", "out_trade_no", "out_refund_no", "refund_id"}
		refundQueryMustParam = []string{"appid", "mch_id", "nonce_str", "sign"}
	)

	for _, k := range refundQueryOneParam {
		if v := param.Get(k); v != nil {
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("need one param of refund_id/out_refund_no/transaction_id/out_trade_no")
	} else if count > 1 {
		return nil, errors.New("more than one param refund_id/out_refund_no/transaction_id/out_trade_no")
	}

	for _, k := range refundQueryMustParam {
		if k == "sign" {
			continue
		}
		if param.Get(k) == nil {
			return nil, errors.New("need param: " + k)
		}
	}

	var refundQueryOptionalParam = []string{"sign_type", "offset"}
	for k := range param {
		if k == "sign_type" {
			signType = param[k].(string)
		}
		if !utils.Contain(refundQueryMustParam, k) && !utils.Contain(refundQueryOptionalParam, k) && !utils.Contain(refundQueryOneParam, k) {
			return nil, errors.New("no need param: " + k)
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
		Url:         "https://api.mch.weixin.qq.com/pay/refundquery",
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
	if resultSign, _ := result.GetString("sign"); resultSign != sign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

//处理退款通知
//对结果中对req_info执行解密：
func (m *myPayer) RefundNotify(body io.Reader) (ResultParam, error) {
	result := ParseXMLReader(body)
	if len(result) == 0 {
		return nil, errors.New("reader has nothing")
	}

	var reqInfoStr string
	if reqInfoStr, _ = result.GetString("req_info"); reqInfoStr == "" {
		return nil, errors.New("wx response without req_info")
	}

	//1. 对加密串A做base64解码，得到加密串B
	//2. 对商户key做md5，得到32位小写key*
	keyStr, _ := secret.NewHasher().HashToString(m.apiKey, crypto.MD5)
	md5Key := strings.ToLower(keyStr)
	//3. 用key*对加密串B做AES-256-ECB解密（PKCS7Padding）
	cipher := secret.NewCipher()
	decryptRequest := &secret.SymRequest{
		PlainData:   nil,
		CipherData:  reqInfoStr,
		Key:         []byte(md5Key),
		Type:        secret.SymTypeAES,
		ModeType:    secret.BlockModeECB,
		PaddingType: secret.PaddingTypePKCS7,
		AddData:     nil,
	}
	realData, err := cipher.SymDecrypt(decryptRequest)
	if err != nil {
		return nil, err
	}
	xmlReader := strings.NewReader(string(realData))
	reqData := ParseXMLReader(xmlReader)
	for k, v := range reqData {
		result[k] = v
	}
	delete(result, "req_info")
	return result, nil
}

//撤销订单
func (m *myPayer) ReverseOrder(param Param, certPath string) (ResultParam, error) {
	if param == nil {
		return nil, errors.New("param is empty")
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	//读取证书
	cert, err := certs.P12ToPem(certPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//校验订单号
	var (
		signType             = pkg.SignTypeMD5 //此处默认MD5
		reverseMustParam     = []string{"appid", "mch_id", "nonce_str", "sign"}
		reverseOneParam      = []string{"transaction_id", "out_trade_no"}
		reverseOptionalParam = []string{"sign_type"}
	)
	var count = 0
	for _, k := range reverseOneParam {
		if v := param.Get(k); v != nil {
			count++
		}
	}
	if count == 0 {
		return nil, errors.New("need order number: transaction_id or out_trade_no")
	} else if count > 1 {
		return nil, errors.New("just one order number: transaction_id or out_trade_no")
	}

	//校验其他参数
	for _, k := range reverseMustParam {
		if k == "sign" {
			continue
		}
		if param.Get(k) == nil {
			return nil, errors.New("need param: " + k)
		}
	}

	for k := range param {
		if k == "sign" {
			continue
		}
		if k == "sign_type" {
			signType = param[k].(string)
		}
		if !utils.Contain(reverseMustParam, k) && !utils.Contain(reverseOptionalParam, k) && !utils.Contain(reverseOneParam, k) {
			return nil, errors.New("no need param: " + k)
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
		Url:         "https://api.mch.weixin.qq.com/secapi/pay/reverse",
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
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, pkg.ErrCheckSign
	}

	return result, nil
}

//企业支付,包括企业付款到零钱、查询企业付款到零钱、企业付款到银行卡
func (m *myPayer) Transfers(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("mch_appid", m.appId)
	param.Add("mchid", m.mchId)

	var transMustParam = []string{"mch_appid", "mchid", "nonce_str", "sign", "partner_trade_no", "openid", "check_name", "amount", "desc", "spbill_create_ip"}
	for _, k := range transMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, errors.New("need param: " + k)
		}
	}

	var transOptionalParam = []string{"device_info", "re_user_name"}
	for k := range param {
		if !utils.Contain(transMustParam, k) && !utils.Contain(transOptionalParam, k) {
			return nil, errors.New("no need param: " + k)
		}
	}

	sign := param.Sign(pkg.SignTypeMD5)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers",
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
	return result, nil
}

//企业付款到银行卡
func (m *myPayer) TransferBank(param Param, p12CertPath string, publicKeyPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("mch_id", m.mchId)

	var bankMustParam = []string{"mch_id", "partner_trade_no", "nonce_str", "sign", "enc_bank_no", "enc_true_name", "bank_code", "amount"}
	for _, k := range bankMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, errors.New("need param: " + k)
		}
	}

	var bankOptionalParam = []string{"desc"}
	for k := range param {
		if !utils.Contain(bankMustParam, k) && !utils.Contain(bankOptionalParam, k) {
			return nil, errors.New("no need param: " + k)
		}
	}

	bankCard := param.Get("enc_bank_no").(string)
	bankName := param.Get("enc_true_name").(string)

	cipher := secret.NewCipher()
	if err = cipher.SetRSAPublicKey(publicKeyPath, secret.PKCSLevel1); err != nil {
		return nil, err
	}
	encryptBankCard, err := cipher.RSAEncryptToString(bankCard, secret.RSAEncryptTypeOAEP, nil)
	if err != nil {
		return nil, err
	}
	encryptBankName, err := cipher.RSAEncryptToString(bankName, secret.RSAEncryptTypeOAEP, nil)
	if err != nil {
		return nil, err
	}
	param.Add("enc_bank_no", encryptBankCard)
	param.Add("enc_true_name", encryptBankName)

	sign := param.Sign(pkg.SignTypeMD5)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank",
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
	sign = result.Sign(pkg.SignTypeMD5)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

//查询转账到零钱
func (m *myPayer) TransfersQuery(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var queryTransferMustParam = []string{"nonce_str", "sign", "partner_trade_no", "mch_id", "appid"}
	for _, k := range queryTransferMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, errors.New("need param: " + k)
		}
	}

	for k := range param {
		if !utils.Contain(queryTransferMustParam, k) {
			return nil, errors.New("no need param: " + k)
		}
	}
	sign := param.Sign(pkg.SignTypeMD5)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo",
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
	return result, nil
}

//查询企业付款到银行卡
func (m *myPayer) TransferBankQuery(param Param, p12CertPath string) (ResultParam, error) {
	if param == nil {
		return nil, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}

	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return nil, err
	}

	param.Add("mch_id", m.mchId)

	var mustParam = []string{"mch_id", "partner_trade_no", "nonce_str", "sign"}
	for _, k := range mustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return nil, errors.New("need param: " + k)
		}
	}

	for k := range param {
		if !utils.Contain(mustParam, k) {
			return nil, errors.New("no need param " + k)
		}
	}

	sign := param.Sign(pkg.SignTypeMD5)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/mmpaysptrans/query_bank",
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
	sign = result.Sign(pkg.SignTypeMD5)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, pkg.ErrCheckSign
	}
	return result, nil
}

//交易保障
func (m *myPayer) Report(param Param) error {
	if param == nil {
		return pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	var (
		signType            = pkg.SignTypeMD5
		reportMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "interface_url", "execute_time", "return_code", "return_msg", "result_code", "user_ip"}
		reportOptionalParam = []string{"device_info", "sign_type", "err_code", "err_code_des", "out_trade_no", "time"}
		reportMicroParam    = []string{"appid", "mch_id", "nonce_str", "sign", "interface_url", "trades", "user_ip"}
		reportMicroOptional = []string{"device_info"}
	)
	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	if v := param.Get("trade"); v != nil {
		for _, k := range reportMicroParam {
			if k == "sign" {
				continue
			}
			if _, ok := param[k]; !ok {
				return errors.New("need param: " + k)
			}
		}
		for key := range param {
			if !utils.Contain(reportMicroParam, key) && !utils.Contain(reportMicroOptional, key) {
				return errors.New("no need param: " + key)
			}
		}
	} else {
		for _, k := range reportMustParam {
			if k == "sign" {
				continue
			}
			if _, ok := param[k]; !ok {
				return errors.New("need param: " + k)
			}
		}
		for key := range param {
			if !utils.Contain(reportMustParam, key) && !utils.Contain(reportOptionalParam, key) {
				return errors.New("no need param: " + key)
			}
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/payitil/report",
		ContentType: pkg.PostContentType,
	}
	response, err := postToWx(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return errors.New(returnMsg)
	}

	if resultCode, err := result.GetString("result_code"); err == nil && resultCode != "SUCCESS" {
		return errors.New("report fail")
	}
	return nil
}

//下载对账单
func (m *myPayer) DownloadBill(param Param, path string) error {
	if param == nil {
		return pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//校验参数
	var downloadMustParam = []string{"appid", "mch_id", "nonce_str", "sign", "bill_date"}
	for _, k := range downloadMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return errors.New("need " + k)
		}
	}

	//校验多余的参数
	var downloadOptionalParam = []string{"bill_type", "tar_type"}
	var tarType string
	for k := range param {
		if !utils.Contain(downloadMustParam, k) && !utils.Contain(downloadOptionalParam, k) {
			return errors.New("no need param: " + k)
		}
		if k == "tar_type" {
			tarType = param.Get(k).(string)
		}
	}

	//签名
	sign := param.Sign(pkg.SignTypeMD5)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/downloadbill",
		ContentType: pkg.PostContentType,
	}

	response, err := postToWx(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	result := ParseXMLReader(bytes.NewReader(content))
	if _, err := result.GetString("return_code"); err == nil {
		returnMsg, _ := result.GetString("return_msg")
		return errors.New(returnMsg)
	}

	if tarType != "" {
		//微信传过来的为gzip压缩了的内容，需要解压
		content, err = utils.UnGZIP(content)
		if err != nil {
			return err
		}
	}

	if !strings.HasSuffix(path, "/") {
		if path == "" {
			path = "./"
		} else {
			path += "/"
		}
	}
	if err = files.MakeNewPath(path); err != nil {
		return err
	}

	//将结果转换为excel文件
	var fileName = param.Get("bill_date").(string) + ".xlsx"
	var sheetName = "ALL" //以查询日期为sheet名
	if t := param.Get("bill_type"); t != nil {
		sheetName = t.(string)
	}

	var billFile *excelize.File
	billFile, _ = excelize.OpenFile(path + fileName)
	if billFile == nil {
		billFile = excelize.NewFile()
		billFile.SetSheetName("Sheet1", sheetName)
	} else {
		billFile.NewSheet(sheetName)
	}

	allData := strings.Replace(string(content), "`", "", -1) //替换掉所有掉参数值前的`符号

	//取订单数据:根据微信返回的结果进行字符串分割操作
	data := strings.Split(allData, "总交易单数")
	orders := strings.Split(data[0], "\n")
	for i, o := range orders {
		if strings.Replace(o, " ", "", -1) == "" {
			continue
		}
		axis := "A" + strconv.Itoa(i+1)
		singleOrder := strings.Split(o, ",")
		billFile.SetSheetRow(sheetName, axis, &singleOrder)
	}
	statis := "总交易单数" + data[1]
	statisArray := strings.Split(statis, "\n")
	for i, s := range statisArray {
		axis := "A" + strconv.Itoa(len(orders)+i+1)
		titles := strings.Split(s, ",")
		billFile.SetSheetRow(sheetName, axis, &titles)
	}

	err = billFile.SaveAs(fileName)
	if err != nil {
		return err
	}
	return os.Rename("./"+fileName, path+fileName)
}

//下载评论
func (m *myPayer) DownloadComment(param Param, p12CertPath string, path string) (offset uint64, err error) {
	if param == nil {
		return 0, pkg.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return 0, err
	}

	//读取证书
	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return 0, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//校验签名方式
	var (
		signType                 = pkg.SignType256
		downCommentMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "begin_time", "end_time", "offset"}
		downCommentOptionalParam = []string{"sign_type", "limit"}
	)

	if _, ok := param["sign_type"]; ok {
		signType = param["sign_type"].(string)
		if signType != pkg.SignType256 {
			return 0, errors.New("download comment only support HMAC-SHA256")
		}
	}
	param.Add("sign_type", signType)

	for _, k := range downCommentMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return 0, errors.New("need param: " + k)
		}
	}

	//校验不必要的参数
	for k := range param {
		if !utils.Contain(downCommentMustParam, k) && !utils.Contain(downCommentOptionalParam, k) {
			return 0, errors.New("no need param: " + k)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return 0, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment",
		ContentType: pkg.PostContentType,
	}

	response, err := postToWxWithCert(request, cert)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	result := ParseXMLReader(bytes.NewReader(content))
	if returnCode, err := result.GetString("return_code"); err == nil && returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return 0, errors.New(returnMsg)
	}
	if resultCode, err := result.GetString("result_code"); err == nil && resultCode != "SUCCESS" {
		errMsg, _ := result.GetString("err_code_des")
		return 0, errors.New(errMsg)
	}
	//判断是否存在目标目录，如果不存在则创建
	if !strings.HasSuffix(path, "/") {
		if path == "" {
			path = "./"
		} else {
			path += "/"
		}
	}
	if err = files.MakeNewPath(path); err != nil {
		return 0, err
	}

	//将结果转换为excel文件，并存放到指定目录
	var fileName = "comment.xlsx"
	var sheetName = "comment" + fmt.Sprintf("%v", param.Get("offset"))
	var commentFile *excelize.File
	commentFile, _ = excelize.OpenFile(path + fileName)
	if commentFile == nil {
		commentFile = excelize.NewFile()
		commentFile.SetSheetName("Sheet1", sheetName)
	} else {
		commentFile.NewSheet(sheetName)
	}

	allData := strings.Replace(string(content), "`", "", -1)
	data := strings.Split(allData, "\n")
	commentFile.SetSheetRow(sheetName, "A1", &[]string{"评论时间", "支付订单号", "评论星级", "评论内容"})
	for i, d := range data {
		if i == 0 {
			//读取微信返回的offset
			offset, err = strconv.ParseUint(d, 10, 64)
			if err != nil {
				return offset, err
			}
			continue
		}
		axis := "A" + strconv.Itoa(i+1)
		singleData := strings.Split(d, ",")
		commentFile.SetSheetRow(sheetName, axis, &singleData)
	}

	err = commentFile.SaveAs(fileName)
	if err != nil {
		return 0, err
	}
	err = os.Rename("./"+fileName, path+fileName)
	return offset, err
}

//下载资金账单
func (m *myPayer) DownloadFundFlow(param Param, p12CertPath string, path string) error {
	if param == nil {
		return errors.New("param is empty")
	}
	if err := m.checkForPay(); err != nil {
		return err
	}
	//读取证书
	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//校验签名方式
	var (
		signType              = pkg.SignType256
		fundFlowMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "bill_date", "account_type"}
		fundFlowOptionalParam = []string{"sign_type", "tar_type"}
	)
	if _, ok := param["sign_type"]; ok {
		signType = param["sign_type"].(string)
		if signType != pkg.SignType256 {
			return errors.New("download fund flow only support HMAC-SHA256")
		}
	}

	//校验必须的参数
	for _, k := range fundFlowMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return errors.New("need param: " + k)
		}
	}

	var tarType string
	//校验是否有不必要的参数
	for k := range param {
		if !utils.Contain(fundFlowMustParam, k) && !utils.Contain(fundFlowOptionalParam, k) {
			return errors.New("no need param: " + k)
		}
		if k == "tar_type" {
			tarType = param.Get(k).(string)
		}
	}

	sign := param.Sign(signType)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/downloadfundflow",
		ContentType: pkg.PostContentType,
	}

	response, err := postToWxWithCert(request, cert)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	result := ParseXMLReader(bytes.NewReader(content))
	if returnCode, err := result.GetString("return_code"); err == nil && returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return errors.New(returnMsg)
	}
	if resultCode, err := result.GetString("result_code"); err == nil && resultCode != "SUCCESS" {
		errMsg, _ := result.GetString("err_code_des")
		return errors.New(errMsg)
	}

	if tarType != "" {
		//需要解压
		content, err = utils.UnGZIP(content)
		if err != nil {
			return err
		}
	}

	//判断是否存在目标目录，如果不存在则创建
	if !strings.HasSuffix(path, "/") {
		if path == "" {
			path = "./"
		} else {
			path += "/"
		}
	}
	if err = files.MakeNewPath(path); err != nil {
		return err
	}

	//将结果转换为excel文件，并存放到指定目录
	var fileName = param.Get("bill_date").(string) + ".xlsx"
	var sheetName = "Basic" //以账户类型为sheet名
	if t := param.Get("account_type"); t != nil {
		sheetName = t.(string)
	}

	//判断是否已经存在excel文件，如果存在直接增加sheet页，否则先创建文件再增加sheet页
	var billFile *excelize.File
	billFile, _ = excelize.OpenFile(path + fileName)
	if billFile == nil {
		billFile = excelize.NewFile()
		billFile.SetSheetName("Sheet1", sheetName)
	} else {
		billFile.NewSheet(sheetName)
	}

	allData := strings.Replace(string(content), "`", "", -1) //替换掉所有掉参数值前的`符号

	//取订单数据:根据微信返回的结果进行字符串分割操作
	data := strings.Split(allData, "资金流水总笔数")
	orders := strings.Split(data[0], "\n")
	for i, o := range orders {
		if strings.Replace(o, " ", "", -1) == "" {
			continue
		}
		axis := "A" + strconv.Itoa(i+1)
		singleOrder := strings.Split(o, ",")
		billFile.SetSheetRow(sheetName, axis, &singleOrder)
	}
	statis := "资金流水总笔数" + data[1]
	statisArray := strings.Split(statis, "\n")
	for i, s := range statisArray {
		axis := "A" + strconv.Itoa(len(orders)+i+1)
		titles := strings.Split(s, ",")
		billFile.SetSheetRow(sheetName, axis, &titles)
	}

	err = billFile.SaveAs(fileName)
	if err != nil {
		return err
	}
	return os.Rename("./"+fileName, path+fileName)
}

//获取RSA加密公钥
func (m *myPayer) GetPublicKey(p12CertPath string, targetPath string) error {
	if m.mchId == "" {
		return pkg.ErrParams
	}

	cert, err := certs.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return err
	}

	param := NewParam()
	//u4 := uuid.NewV4()
	//s := strings.Replace(fmt.Sprintf("%s", u4), "-", "", -1)
	nonceStr := fmt.Sprintf("%d", time.Now().UnixNano())

	param.Add("mch_id", m.mchId)
	param.Add("nonce_str", nonceStr)
	param.Add("sign_type", pkg.SignTypeMD5)
	param.Add("sign", param.Sign(pkg.SignTypeMD5))

	reader, err := param.MarshalXML()
	if err != nil {
		return err
	}
	var request = &postRequest{
		Body:        reader,
		Url:         "https://fraud.mch.weixin.qq.com/risk/getpublickey",
		ContentType: pkg.PostContentType,
	}
	response, err := postToWxWithCert(request, cert)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result := ParseXMLReader(response.Body)
	if returnCode, _ := result.GetString("return_code"); returnCode != "SUCCESS" {
		returnMsg, _ := result.GetString("return_msg")
		return errors.New(returnMsg)
	}
	if resultCode, _ := result.GetString("result_code"); resultCode != "SUCCESS" {
		errDes, _ := result.GetString("err_code_des")
		return errors.New(errDes)
	}
	keyValue, err := result.GetString("pub_key")
	if err != nil {
		return err
	}
	//判断是否存在目标目录，如果不存在则创建
	if !strings.HasSuffix(targetPath, "/") {
		if targetPath == "" {
			targetPath = "./"
		} else {
			targetPath += "/"
		}
	}
	if err = files.MakeNewPath(targetPath); err != nil {
		return err
	}

	f, err := os.Create(targetPath + pkg.PublicKey)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(keyValue))
	if err != nil {
		return err
	}
	return nil
}

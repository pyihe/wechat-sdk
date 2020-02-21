package wechat_sdk

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	下载资金账单
*/

func (m *myPayer) DownloadFundFlow(param Param, p12CertPath string, path string) error {
	if param == nil {
		return errors.New("param is empty")
	}
	if err := m.checkForPay(); err != nil {
		return err
	}
	//读取证书
	cert, err := util.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//校验签名方式
	var (
		signType              = e.SignType256
		fundFlowMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "bill_date", "account_type"}
		fundFlowOptionalParam = []string{"sign_type", "tar_type"}
	)
	if _, ok := param["sign_type"]; ok {
		signType = param["sign_type"].(string)
		if signType != e.SignType256 {
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
		if !util.HaveInArray(fundFlowMustParam, k) && !util.HaveInArray(fundFlowOptionalParam, k) {
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
		ContentType: e.PostContentType,
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
		content, err = util.UnGZIP(content)
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
	if err = util.MakeNewPath(path); err != nil {
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

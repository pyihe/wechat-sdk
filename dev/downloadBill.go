package dev

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
	下载对账单
*/

var (
	downloadMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "bill_date"}
	downloadOptionalParam = []string{"bill_type", "tar_type"}
)

const downloadApiUrl = "https://api.mch.weixin.qq.com/pay/downloadbill"

func (m *myPayer) DownloadBill(param Param, path string) error {
	if param == nil {
		return e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//校验参数
	for _, k := range downloadMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return errors.New("need " + k)
		}
	}

	//校验多余的参数
	var tarType string
	for k := range param {
		if !util.HaveInArray(downloadMustParam, k) && !util.HaveInArray(downloadOptionalParam, k) {
			return errors.New("no need param: " + k)
		}
		if k == "tar_type" {
			tarType = param.Get(k).(string)
		}
	}

	//签名
	sign := param.Sign(e.SignTypeMD5)
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         downloadApiUrl,
		ContentType: "application/xml;charset=utf-8",
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
		buffer := new(bytes.Buffer)
		err = binary.Write(buffer, binary.LittleEndian, content)
		if err != nil {
			return err
		}
		zipReader, err := gzip.NewReader(buffer)
		if err != nil {
			return err
		}
		defer zipReader.Close()
		content, err = ioutil.ReadAll(zipReader)
		if err != nil {
			return err
		}
	}

	//将结果转换为excel文件
	var fileName = param.Get("bill_date").(string) + ".xlsx"
	var sheetName = "ALL" //以查询日期为sheet名
	if t := param.Get("bill_type"); t != nil {
		sheetName = t.(string)
	}

	billFile := excelize.NewFile()
	billFile.NewSheet(sheetName)
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
	billFile.DeleteSheet("sheet1")
	err = billFile.SaveAs(fileName)
	if err != nil {
		return err
	}
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
	return os.Rename("./"+fileName, path+fileName)
}

package dev

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

func (m *myPayer) DownloadComment(param Param, p12CertPath string, path string) (offset uint64, err error) {
	if param == nil {
		return 0, e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return 0, err
	}

	//读取证书
	cert, err := util.P12ToPem(p12CertPath, m.mchId)
	if err != nil {
		return 0, err
	}

	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//校验签名方式
	var signType = e.SignType256
	if _, ok := param["sign_type"]; ok {
		signType = param["sign_type"].(string)
		if signType != e.SignType256 {
			return 0, errors.New("download comment only support HMAC-SHA256")
		}
	}
	param.Add("sign_type", signType)

	var downCommentMustParam = []string{"appid", "mch_id", "nonce_str", "sign", "begin_time", "end_time", "offset"}
	for _, k := range downCommentMustParam {
		if k == "sign" {
			continue
		}
		if _, ok := param[k]; !ok {
			return 0, errors.New("need param: " + k)
		}
	}

	//校验不必要的参数
	var downCommentOptionalParam = []string{"sign_type", "limit"}
	for k := range param {
		if !util.HaveInArray(downCommentMustParam, k) && !util.HaveInArray(downCommentOptionalParam, k) {
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
		ContentType: e.PostContentType,
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
	if err = util.MakeNewPath(path); err != nil {
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

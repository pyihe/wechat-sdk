package wechat_sdk

import (
	"errors"
	"github.com/hong008/wechat-sdk/pkg/e"
	"github.com/hong008/wechat-sdk/pkg/util"
)

/*
	统一下单
*/

//统一下单
func (m *myPayer) UnifiedOrder(param Param) (ResultParam, error) {
	if param == nil {
		return nil, e.ErrParams
	}
	if err := m.checkForPay(); err != nil {
		return nil, err
	}
	param.Add("appid", m.appId)
	param.Add("mch_id", m.mchId)

	//获取交易类型和签名类型
	var (
		unifiedMustParam     = []string{"appid", "mch_id", "nonce_str", "sign", "body", "out_trade_no", "total_fee", "spbill_create_ip", "notify_url", "trade_type"}
		unifiedOptionalParam = []string{"device_info", "sign_type", "detail", "attach", "fee_type", "time_start", "time_expire", "goods_tag", "limit_pay", "receipt", "openid", "product_id", "scene_info"}
		tradeType            string
		signType             = e.SignTypeMD5 //默认MD5签名方式
	)

	if t, ok := param["trade_type"]; ok {
		tradeType = t.(string)
	} else {
		return nil, e.ErrTradeType
	}
	if t, ok := param["sign_type"]; ok {
		signType = t.(string)
	}

	switch tradeType {
	case e.H5:
		//H5支付必须要传scene_info参数
		if sceneInfo := param.Get("scene_info"); sceneInfo == nil || sceneInfo.(string) == "" {
			return nil, errors.New("H5 pay need param scene_info")
		}
	case e.App:
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
	case e.JSAPI:
		//JSAPI支付必须传openid参数
		if openId, ok := param["openid"]; !ok || openId.(string) == "" {
			return nil, e.ErrOpenId
		}
	case e.Native:
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
		if !util.HaveInArray(unifiedMustParam, key) && !util.HaveInArray(unifiedOptionalParam, key) {
			return nil, errors.New("no need param: " + key)
		}
	}

	sign := param.Sign(m.apiKey, signType)
	//将签名添加到需要发送的参数里
	param.Add("sign", sign)

	reader, err := param.MarshalXML()
	if err != nil {
		return nil, err
	}

	var request = &postRequest{
		Body:        reader,
		Url:         "https://api.mch.weixin.qq.com/pay/unifiedorder",
		ContentType: e.PostContentType,
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
	sign = result.Sign(m.apiKey, signType)
	if wxSign, _ := result.GetString("sign"); sign != wxSign {
		return nil, e.ErrCheckSign
	}
	return result, err
}

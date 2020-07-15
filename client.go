package wechat_sdk

import (
	"errors"
	"io"
	"sync"
)

type ResultParam interface {
	//获取整型数据, 如支付/退款金额
	GetInt64(key string, base int) (value int64, err error)
	//获取字符串数据, 如订单号
	GetString(key string) (value string, err error)
	//
	Data() map[string]string
}

type WePayer interface {
	//支付相关
	//统一下单
	UnifiedOrder(param Param) (ResultParam, error)
	//扫码下单
	UnifiedMicro(param Param) (ResultParam, error)
	//撤销订单
	ReverseOrder(param Param, p12CertPath string) (ResultParam, error)
	//查询订单
	UnifiedQuery(param Param) (ResultParam, error)
	//关闭订单
	CloseOrder(param Param) (ResultParam, error)
	//退款
	RefundOrder(param Param, p12CertPath string) (ResultParam, error)
	//退款查询
	RefundQuery(param Param) (ResultParam, error)
	//解析退款通知, 结果将不会返回req_info
	RefundNotify(body io.Reader) (ResultParam, error)
	//下单对账单
	DownloadBill(param Param, fileSavePath string) error
	//下载资金账单
	DownloadFundFlow(param Param, p12CertPath string, fileSavePath string) error
	//拉取订单评论数据
	DownloadComment(param Param, p12CertPath string, fileSavePath string) (offset uint64, err error)
	//交易保障
	Report(param Param) error

	//
	//企业付款
	//付款到零钱
	Transfers(param Param, p12CertPath string) (ResultParam, error)
	//查询企业付款到零钱
	TransfersQuery(param Param, p12CertPath string) (ResultParam, error)
	//从微信获取RSA加密的公钥
	GetPublicKey(p12CertPath string, targetPath string) error
	//TODO 待验证 企业付款到银行卡
	TransferBank(param Param, p12CertPath string, publicKeyPath string) (ResultParam, error)
	//TODO 待验证 查询企业付款到银行卡
	TransferBankQuery(param Param, p12CertPath string) (ResultParam, error)

	//分账相关
	//申请单次分账or多次分账
	ProfitSharing(param Param, p12CertPath string, multiTag bool) (ResultParam, error)
	//查询申请分账的结果
	QueryProfitSharing(param Param, p12CertPath string) (ResultParam, error)
	//添加分账接收方
	AddProfitSharingReceiver(param Param) (ResultParam, error)
	//删除分账接收方
	RemoveProfitSharingReceiver(param Param) (ResultParam, error)
	//完结分账
	FinishProfitSharing(param Param, p12CertPath string) (ResultParam, error)
	//分账回退
	ReturnProfitSharing(param Param, p12CertPath string) (ResultParam, error)
	//回退结果查询
	QueryProfitSharingReturn(param Param) (ResultParam, error)
	//分账动帐通知(此处无视返回结果的层级关系，对需要的字段直接使用Get方法获取对应的结果)
	ProfitSharingNotify(body io.Reader) (ResultParam, error)

	//红包相关
	//发放红包
	SendRedPack(param Param, p12CertPath string) (ResultParam, error)
	//发放裂变红包
	SendGroupRedPack(param Param, p12CertPath string) (ResultParam, error)
	//查询红包发放记录
	GetRedPackRecords(param Param, p12CertPath string) (ResultParam, error)

	//小程序相关
	//获取授权access_token
	GetAccessTokenForMini() (Param, error) //获取小程序接口凭证，使用者自己保存token，过期重新获取
	//获取微信信息
	GetUserInfoForMini(code, encryptedData, ivData string) (Param, error)
	//获取微信手机号码
	GetUserPhoneForMini(code string, encryptedData string, ivData string) (Param, error)
	//获取session_key
	GetSessionKeyAndOpenId(code string) (Param, error)

	//公众号相关
	GetAppBaseAccessToken() (Param, error)
	GetAppOauthAccessToken(code string) (Param, error)
	RefreshOauthToken(refreshToken string) (Param, error)
	GetAppUserInfo(oauthToken, openId, lang string) (Param, error)
	CheckOauthToken(oauthToken, openId string) (bool, error)
}

type option func(*myPayer)

var (
	defaultPayer *myPayer
)

type myPayer struct {
	once     sync.Once
	appId    string //appid
	mchId    string //mchid
	secret   string //secret用于获取token
	apiKey   string //用于支付
	apiV3Key string //api v3 key
}

//不向微信发送接口请求report
func NewPayer(options ...option) WePayer {
	defaultPayer = &myPayer{}
	defaultPayer.once.Do(func() {
		for _, option := range options {
			option(defaultPayer)
		}
	})
	return defaultPayer
}

func WithAppId(appId string) option {
	return func(payer *myPayer) {
		payer.appId = appId
	}
}

func WithMchId(mchId string) option {
	return func(payer *myPayer) {
		payer.mchId = mchId
	}
}

func WithSecret(secret string) option {
	return func(payer *myPayer) {
		payer.secret = secret
	}
}

func WithApiKey(key string) option {
	return func(payer *myPayer) {
		payer.apiKey = key
	}
}

func WithApiV3Key(v3Key string) option {
	return func(payer *myPayer) {
		payer.apiV3Key = v3Key
	}
}

func (m *myPayer) checkForPay() error {
	if m.appId == "" {
		return errors.New("need appid")
	}
	if m.mchId == "" {
		return errors.New("need mch_id")
	}
	if m.apiKey == "" {
		return errors.New("need api key")
	}
	return nil
}

func (m *myPayer) checkForAccess() error {
	if m.appId == "" {
		return errors.New("need appid")
	}
	if m.secret == "" {
		return errors.New("need secret")
	}
	return nil
}

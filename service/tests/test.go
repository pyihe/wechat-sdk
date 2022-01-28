package tests

import "github.com/pyihe/wechat-sdk/v3/service"

var (
	Config     *service.Config
	mchId      = "**********"
	appId      = "******************"
	apiV3Key   = "********************************"
	serialNo   = "****************************************"
	publicKey  = "../test/pem/public_key.pem"
	privateKey = "../test/pem/apiclient_key.pem"
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
	Config = service.NewConfig(opts...)
}

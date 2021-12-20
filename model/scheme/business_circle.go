package scheme

import (
	"time"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/vars"
)

/*************************************智慧商圈: 支付结果通知**************************************************************/

type BusinessCirclePayResponse struct {
	Id            string    `json:"id,omitempty"`             // 通知ID
	MchId         string    `json:"mchid,omitempty"`          // 商户号
	MerchantName  string    `json:"merchant_name,omitempty"`  // 商圈商户名称
	ShopName      string    `json:"shop_name,omitempty"`      // 门店名称
	ShopNumber    string    `json:"shop_number,omitempty"`    // 门店编号
	AppId         string    `json:"appid,omitempty"`          // 小程序appid
	OpenId        string    `json:"openid,omitempty"`         // 用户标识
	TimeEnd       time.Time `json:"time_end,omitempty"`       // 交易完成时间
	Amount        int64     `json:"amount,omitempty"`         // 用户实际消费金额
	TransactionId string    `json:"transaction_id,omitempty"` // 微信支付订单号
	CommitTag     string    `json:"commit_tag,omitempty"`     // 手动提交积分标记
}

/*************************************智慧商圈: 商圈积分同步**************************************************************/

type BusinessCirclePointsRequest struct {
	TransactionId    string `json:"transaction_id"`              // 微信订单号
	AppId            string `json:"appid"`                       // 小程序appid
	OpenId           string `json:"openid"`                      // 用户标识
	EarnPoints       bool   `json:"earn_points"`                 // 是否活得积分
	IncreasedPoints  int64  `json:"increased_points"`            // 订单新增积分值
	PointsUpdateTime string `json:"points_update_time"`          // 积分更新时间
	NoPointsRemarks  string `json:"no_points_remarks,omitempty"` // 未活得积分的备注信息
	TotalPoints      int64  `json:"total_points,omitempty"`      // 当前顾客积分总额
}

func (b *BusinessCirclePointsRequest) Check() (err error) {
	if b == nil {
		err = vars.ErrNoRequest
		return
	}
	if b.TransactionId == "" {
		err = errors.New("请填写transaction_id!")
		return
	}
	if b.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if b.OpenId == "" {
		err = errors.New("请填写openid!")
		return
	}
	if b.IncreasedPoints == 0 {
		err = errors.New("请填写increased_points!")
		return
	}
	if b.PointsUpdateTime == "" {
		err = errors.New("请填写points_update_time!")
		return
	}
	return
}

type BusinessCirclePointsResponse struct {
	RequestId string `json:"request_id,omitempty"`
}

/*************************************智慧商圈: 退款成功通知**************************************************************/

type BusinessCircleRefundResponse struct {
	Id            string    `json:"id,omitempty"`             // 通知ID
	MchId         string    `json:"mchid,omitempty"`          // 商户号
	MerchantName  string    `json:"merchant_name,omitempty"`  // 商圈商户名称
	ShopName      string    `json:"shop_name,omitempty"`      // 门店名称
	ShopNumber    string    `json:"shop_number,omitempty"`    // 门店编号
	AppId         string    `json:"appid,omitempty"`          // 小程序appid
	OpenId        string    `json:"openid,omitempty"`         // 用户标识
	RefundTime    time.Time `json:"refund_time,omitempty"`    // 退款完成时间
	PayAmount     int64     `json:"pay_amount,omitempty"`     // 用户实际消费金额
	RefundAmount  int64     `json:"refund_amount,omitempty"`  // 用户退款金额
	TransactionId string    `json:"transaction_id,omitempty"` // 微信支付订单号
	RefundId      string    `json:"refund_id,omitempty"`      // 微信支付退款单号
}

/*************************************智慧商圈: 商圈积分授权查询***********************************************************/

type BusinessCircleQueryAuthorizationRequest struct {
	AppId  string `json:"-"` // 小程序appid
	OpenId string `json:"-"` // 顾客openid
}

func (b *BusinessCircleQueryAuthorizationRequest) Check() (err error) {
	if b == nil {
		err = vars.ErrNoRequest
		return
	}
	if b.AppId == "" {
		err = errors.New("请填写appid!")
		return
	}
	if b.OpenId == "" {
		err = errors.New("请填写openid!")
		return
	}
	return
}

type BusinessCircleQueryAuthorizationResponse struct {
	RequestId       string    `json:"-"`                          // 请求的唯一ID
	OpenId          string    `json:"openid,omitempty"`           // 顾客openid
	AuthorizeState  string    `json:"authorize_state,omitempty"`  // 授权状态
	AuthorizeTime   time.Time `json:"authorize_time,omitempty"`   // 授权时间
	DeauthorizeTime time.Time `json:"deauthorize_time,omitempty"` // 取消授权时间
}

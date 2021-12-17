package merchant

import (
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/wechat-sdk/model"
	"github.com/pyihe/wechat-sdk/vars"
)

/*************************************支付即服务: 服务人员注册*************************************************************/

type GuideRegisterRequest struct {
	Corpid      string `json:"corpid,omitempty"`       // 企业ID
	StoreId     int32  `json:"store_id,omitempty"`     // 门店ID
	UserId      string `json:"userid,omitempty"`       // 企业微信的员工ID
	Name        string `json:"name,omitempty"`         // 企业微信的员工姓名
	Mobile      string `json:"mobile,omitempty"`       // 手机号码
	QrCode      string `json:"qr_code,omitempty"`      // 员工个人二维码
	Avatar      string `json:"avatar,omitempty"`       // 头像URL
	GroupQrCode string `json:"group_qrcode,omitempty"` // 群二维码URL
}

func (guide *GuideRegisterRequest) Check() (err error) {
	if guide == nil {
		err = vars.ErrNoRequest
		return
	}
	if guide.Corpid == "" {
		err = errors.New("请填写corpid!")
		return
	}
	if guide.StoreId == 0 {
		err = errors.New("请填写store_id!")
		return
	}
	if guide.UserId == "" {
		err = errors.New("请填写userid!")
		return
	}
	if guide.Name == "" {
		err = errors.New("请填写name!")
		return
	}
	if guide.Mobile == "" {
		err = errors.New("请填写mobile!")
		return
	}
	if guide.QrCode == "" {
		err = errors.New("请填写qr_code!")
		return
	}
	if guide.Avatar == "" {
		err = errors.New("请填写avatar!")
		return
	}
	return
}

type GuideRegisterResponse struct {
	model.WechatError
	RequestId string `json:"-"`                  // 唯一请求ID
	GuideId   string `json:"guide_id,omitempty"` // 服务人员ID
}

/*************************************支付即服务: 服务人员分配*************************************************************/

type GuideAssignRequest struct {
	GuideId    string `json:"-"`                      // 服务人员ID
	OutTradeNo string `json:"out_trade_no,omitempty"` // 商户订单号
}

func (guide *GuideAssignRequest) Check() (err error) {
	if guide == nil {
		err = vars.ErrNoRequest
		return
	}
	if guide.GuideId == "" {
		err = errors.New("请填写guide_id!")
		return
	}
	if guide.OutTradeNo == "" {
		err = errors.New("请填写out_trade_no!")
		return
	}
	return
}

type GuideAssignResponse struct {
	RequestId string
}

/*************************************支付即服务: 服务人员查询*************************************************************/

type GuideQueryRequest struct {
	StoreId int32  `json:"store_id,omitempty"` // 门店ID
	UserId  string `json:"userid,omitempty"`   // 企业微信的员工ID
	Mobile  string `json:"mobile,omitempty"`   // 手机号码
	WorkId  string `json:"work_id,omitempty"`  // 工号
	Limit   int32  `json:"limit,omitempty"`    // 最大资源条数
	Offset  int32  `json:"offset,omitempty"`   // 请求资源起始位置
}

func (guide *GuideQueryRequest) Check() (err error) {
	if guide == nil {
		err = vars.ErrNoRequest
		return
	}
	if guide.StoreId == 0 {
		err = errors.New("请填写store_id!")
		return
	}
	if guide.Limit > 10 {
		err = errors.New("limit上限为10!")
		return
	}
	return
}

type GuideQueryResponse struct {
	RequestId  string
	Data       []*GuidePeople `json:"data,omitempty"`        // 服务人员列表
	TotalCount int32          `json:"total_count,omitempty"` // 服务人员数量
	Limit      int32          `json:"limit,omitempty"`       // 最大资源条数
	Offset     int32          `json:"offset,omitempty"`      // 请求资源起始位置
}

type GuidePeople struct {
	GuideId string `json:"guide_id,omitempty"` // 服务人员ID
	StoreId int32  `json:"store_id,omitempty"` // 门店ID
	Name    string `json:"name,omitempty"`     // 服务人员姓名
	Mobile  string `json:"mobile,omitempty"`   // 服务人员手机号码
	UserId  string `json:"userid,omitempty"`   // 微信企业员工的ID
	WorkId  string `json:"work_id,omitempty"`  // 工号
}

/*************************************支付即服务: 服务人员信息更新*********************************************************/

type GuideUpdateRequest struct {
	GuideId     string `json:"-"`                      // 服务人员ID
	Name        string `json:"name,omitempty"`         // 服务人员姓名
	Mobile      string `json:"mobile,omitempty"`       // 服务人员手机号码
	QrCode      string `json:"qr_code,omitempty"`      // 服务人员二维码
	Avatar      string `json:"avatar,omitempty"`       // 服务人员头像
	GroupQrCode string `json:"group_qrcode,omitempty"` // 群二维码URL
}

func (guide *GuideUpdateRequest) Check() (err error) {
	if guide == nil {
		err = vars.ErrNoRequest
		return
	}
	if guide.GuideId == "" {
		err = errors.New("请填写guide_id!")
		return
	}
	return
}

type GuideUpdateResponse struct {
	RequestId string
}

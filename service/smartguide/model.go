package smartguide

import "github.com/pyihe/wechat-sdk/v3/model"

// RegisterRequest 服务人员注册请求
type RegisterRequest struct {
	SubMchId    string `json:"sub_mchid,omitempty"`    // 子商户ID
	Corpid      string `json:"corpid"`                 // 企业ID
	StoreId     int64  `json:"store_id"`               // 门店ID
	UserId      string `json:"userid"`                 // 企业微信的员工ID
	Name        string `json:"name"`                   // 企业微信的员工姓名
	Mobile      string `json:"mobile"`                 // 手机号码
	QrCode      string `json:"qr_code"`                // 员工个人二维码
	Avatar      string `json:"avatar"`                 // 头像URL
	GroupQrCode string `json:"group_qrcode,omitempty"` // 群二维码URL
}

func (r *RegisterRequest) clone() *RegisterRequest {
	return &RegisterRequest{
		SubMchId:    r.SubMchId,
		Corpid:      r.Corpid,
		StoreId:     r.StoreId,
		UserId:      r.UserId,
		Name:        r.Name,
		Mobile:      r.Mobile,
		QrCode:      r.QrCode,
		Avatar:      r.Avatar,
		GroupQrCode: r.GroupQrCode,
	}
}

// RegisterResponse 服务人员注册应答
type RegisterResponse struct {
	model.WechatError
	RequestId string
	GuideId   string `json:"guide_id,omitempty"` // 服务人员ID
}

// AssignRequest 服务人员分配request
type AssignRequest struct {
	GuideId    string `json:"-"`
	SubMchId   string `json:"sub_mchid,omitempty"`
	OutTradeNo string `json:"out_trade_no"`
}

// AssignResponse 服务人员分配response
type AssignResponse struct {
	model.WechatError
	RequestId string
}

// QueryRequest 查询服务人员
type QueryRequest struct {
	SubMchId string `json:"sub_mchid,omitempty"` // 子商户号
	StoreId  int64  `json:"store_id,omitempty"`  // 门店ID
	UserId   string `json:"userid,omitempty"`    // 企业微信的员工ID
	Mobile   string `json:"mobile,omitempty"`    // 手机号码
	WorkId   string `json:"work_id,omitempty"`   // 工号
	Limit    int    `json:"limit,omitempty"`     // 最大资源条数
	Offset   int    `json:"offset,omitempty"`    // 请求资源起始位置
}

// QueryResponse 查询服务人员
type QueryResponse struct {
	model.WechatError
	RequestId  string    `json:"-"`
	Data       []*Worker `json:"data,omitempty"`        // 服务人员列表
	TotalCount int       `json:"total_count,omitempty"` // 服务人员数量
	Limit      int       `json:"limit,omitempty"`       // 最大资源条数
	Offset     int       `json:"offset,omitempty"`      // 请求资源起始位置
}

// Worker 服务人员信息
type Worker struct {
	GuideId string `json:"guide_id,omitempty"` // 服务人员ID
	StoreId int64  `json:"store_id,omitempty"` // 门店ID
	Name    string `json:"name,omitempty"`     // 服务人员姓名
	Mobile  string `json:"mobile,omitempty"`   // 服务人员手机号
	UserId  string `json:"userid,omitempty"`   // 企业温馨的员工ID
	WorkId  string `json:"work_id,omitempty"`  // 工号
}

// UpdateRequest 服务人员信息更新request
type UpdateRequest struct {
	SubMchId    string `json:"sub_mchid,omitempty"`    // 子商户号
	Name        string `json:"name,omitempty"`         // 服务人员姓名
	Mobile      string `json:"mobile,omitempty"`       // 服务人员手机号码
	QrCode      string `json:"qr_code,omitempty"`      // 服务人员二维码URL
	Avatar      string `json:"avatar,omitempty"`       // 服务人员头像URL
	GroupQrCode string `json:"group_qrcode,omitempty"` // 群二维码URL
}

func (u *UpdateRequest) isZero() bool {
	if len(u.SubMchId) > 0 {
		return false
	}
	if len(u.Name) > 0 {
		return false
	}
	if len(u.Mobile) > 0 {
		return false
	}
	if len(u.QrCode) > 0 {
		return false
	}
	if len(u.Avatar) > 0 {
		return false
	}
	if len(u.GroupQrCode) > 0 {
		return false
	}
	return true
}

func (u *UpdateRequest) clone() *UpdateRequest {
	return &UpdateRequest{
		SubMchId:    u.SubMchId,
		Name:        u.Name,
		Mobile:      u.Mobile,
		QrCode:      u.QrCode,
		Avatar:      u.Avatar,
		GroupQrCode: u.GroupQrCode,
	}
}

// UpdateResponse 更新服务人员信息应答
type UpdateResponse struct {
	model.WechatError
	RequestId string `json:"-"`
}

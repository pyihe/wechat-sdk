package wechat_sdk

type User interface {
	OpenId() string
	NickName() string
	Sex() int
	City() string
	Country() string
	HeadUrl() string
	Unionid() string
}

/*
annotation
*/

type userInfo struct {
	openId   string `json:"openid"`
	nickName string `json:"nickname"`
	sex      int    `json:"sex"`
	city     string `json:"CITY"`
	country  string `json:"COUNTRY"`
	headUrl  string `json:"headimgurl"`
	unionid  string `json:"unionid"`
	errMsg   string `json:"errmsg"`
	errCode  int    `json:"errcode"`
}

func (u *userInfo) OpenId() string {
	return u.openId
}

func (u *userInfo) NickName() string {
	return u.nickName
}

func (u *userInfo) Sex() int {
	return u.sex
}

func (u *userInfo) City() string {
	return u.city
}

func (u *userInfo) Country() string {
	return u.country
}

func (u *userInfo) HeadUrl() string {
	return u.country
}

func (u *userInfo) Unionid() string {
	return u.unionid
}

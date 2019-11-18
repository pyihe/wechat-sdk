package wechat_sdk

type User interface {
	GetOpenId() string
	GetNickName() string
	GetSex() int
	GetProvince() string
	GetCity() string
	GetCountry() string
	GetHeadUrl() string
	GetUnionid() string
}

/*
annotation
*/

type userInfo struct {
	OpenId   string `json:"openid"`
	NickName string `json:"nickname"`
	Sex      int    `json:"sex"`
	Province string `json:"PROVINCE"`
	City     string `json:"CITY"`
	Country  string `json:"COUNTRY"`
	HeadUrl  string `json:"headimgurl"`
	Unionid  string `json:"unionid"`
	ErrMsg   string `json:"errmsg"`
	ErrCode  int    `json:"errcode"`
}

func (u *userInfo) GetOpenId() string {
	return u.OpenId
}

func (u *userInfo) GetNickName() string {
	return u.NickName
}

func (u *userInfo) GetSex() int {
	return u.Sex
}

func (u *userInfo) GetProvince() string {
	return u.Province
}

func (u *userInfo) GetCity() string {
	return u.City
}

func (u *userInfo) GetCountry() string {
	return u.Country
}

func (u *userInfo) GetHeadUrl() string {
	return u.HeadUrl
}

func (u *userInfo) GetUnionid() string {
	return u.Unionid
}

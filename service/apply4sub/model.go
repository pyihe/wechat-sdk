package apply4sub

import (
	"github.com/pyihe/go-pkg/clone"
	"github.com/pyihe/wechat-sdk/v3/model"
)

// ApplyRequest 提交申请单请求参数
type ApplyRequest struct {
	BusinessCode    string           `json:"business_code"`           // 业务申请编号
	ContactInfo     *ContactInfo     `json:"contact_info"`            // 超级管理员信息
	SubjectInfo     *SubjectInfo     `json:"subject_info"`            // 主体资料
	BusinessInfo    *BusinessInfo    `json:"business_info"`           // 经营资料
	SettlementInfo  *SettlementInfo  `json:"settlement_info"`         // 结算规则
	BankAccountInfo *BankAccountInfo `json:"bank_account_info"`       // 结算银行账户
	AdditionInfo    *AdditionInfo    `json:"addition_info,omitempty"` // 补充材料
}

func (apply *ApplyRequest) clone() *ApplyRequest {
	data := clone.DeepClone(apply)
	req, ok := data.(*ApplyRequest)
	if !ok || req == nil {
		return nil
	}
	return req
}

// ApplyResponse 提交申请单应答参数
type ApplyResponse struct {
	model.WechatError
	RequestId   string // 唯一请求ID
	ApplymentId uint64 `json:"applyment_id,omitempty"` // 微信支付申请单号
}

// QueryApplymentRequest 查询申请单状态请求参数
type QueryApplymentRequest struct {
	ApplymentId  uint64 // 申请单号
	BusinessCode string // 业务申请编号
}

// QueryApplymentResponse 查询申请单状态应答参数
type QueryApplymentResponse struct {
	model.WechatError
	RequestId         string         // 唯一请求ID
	BusinessCode      string         `json:"business_code,omitempty"`       // 业务申请编号
	ApplymentId       uint64         `json:"applyment_id,omitempty"`        // 微信支付申请单号
	SubMchId          string         `json:"sub_mchid,omitempty"`           // 特约商户号
	SignUrl           string         `json:"sign_url,omitempty"`            // 超级管理员签约链接
	ApplymentState    string         `json:"applyment_state,omitempty"`     // 申请单状态
	ApplymentStateMsg string         `json:"applyment_state_msg,omitempty"` // 申请状态描述
	AuditDetail       []*AuditDetail `json:"audit_detail,omitempty"`        // 驳回原因详情
}

// ModifySettlementRequest 修改结算账号请求参数
type ModifySettlementRequest struct {
	SubMchId        string `json:"-"`                           // 特约商户/二级商户号
	AccountType     string `json:"account_type,omitempty"`      // 账户类型
	AccountBank     string `json:"account_bank,omitempty"`      // 开户银行
	BankAddressCode string `json:"bank_address_code,omitempty"` // 开户银行省市编码
	BankName        string `json:"bank_name,omitempty"`         // 开户银行全称(含支行)
	BankBranchId    string `json:"bank_branch_id,omitempty"`    // 开户银行联行号
	AccountNumber   string `json:"account_number,omitempty"`    // 银行账号
}

func (m *ModifySettlementRequest) clone() *ModifySettlementRequest {
	return &ModifySettlementRequest{
		SubMchId:        m.SubMchId,
		AccountType:     m.AccountType,
		AccountBank:     m.AccountBank,
		BankAddressCode: m.BankAddressCode,
		BankName:        m.BankName,
		BankBranchId:    m.BankBranchId,
		AccountNumber:   m.AccountNumber,
	}
}

// ModifySettlementResponse 修改结算账号应答参数
type ModifySettlementResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

// QuerySettlementResponse 查询结算账号应答参数
type QuerySettlementResponse struct {
	model.WechatError
	RequestId        string // 唯一请求ID
	AccountType      string `json:"account_type,omitempty"`       // 账户类型
	AccountBank      string `json:"account_bank,omitempty"`       // 开户银行
	BankName         string `json:"bank_name,omitempty"`          // 银行名称(含支行)
	BankBranchId     string `json:"bank_branch_id,omitempty"`     // 开户银行联行号
	AccountNumber    string `json:"account_number,omitempty"`     // 银行账号
	VerifyResult     string `json:"verify_result,omitempty"`      // 汇款验证结果
	VerifyFailReason string `json:"verify_fail_reason,omitempty"` // 汇款验证失败原因
}

// AuditDetail 驳回原因详情
type AuditDetail struct {
	Field        string `json:"field,omitempty"`         // 字段名
	FieldName    string `json:"field_name,omitempty"`    // 字段名称
	RejectReason string `json:"reject_reason,omitempty"` // 驳回原因
}

// ContactInfo 超级管理员信息
type ContactInfo struct {
	ContactName     string `json:"contact_name"`                // 超级管理员姓名, 需要加密处理
	ContactIdNumber string `json:"contact_id_number,omitempty"` // 超级管理员身份证件号, 需要加密处理
	OpenId          string `json:"open_id,omitempty"`           // 超级管理员openid
	MobilePhone     string `json:"mobile_phone"`                // 超级管理员联系手机, 需要加密
	ContactEmail    string `json:"contact_email"`               // 超级管理员邮箱, 需要加密处理
}

// SubjectInfo 主体资料
type SubjectInfo struct {
	SubjectType           string               `json:"subject_type"`                      // 主体类型
	BusinessLicenseInfo   *BusinessLicenseInfo `json:"business_license_info,omitempty"`   // 营业执照
	CertificateInfo       *CertificateInfo     `json:"certificate_info,omitempty"`        // 登记证书
	OrganizationInfo      *OrganizationInfo    `json:"organization_info,omitempty"`       // 组织机构代码证
	CertificateLetterCopy string               `json:"certificate_letter_copy,omitempty"` // 单位证明函照片
	IdentityInfo          *IdentityInfo        `json:"identity_info"`                     // 经营者/法人身份证件
	UboInfo               *UboInfo             `json:"ubo_info,omitempty"`                // 最终受益人信息
}

// BusinessLicenseInfo 营业执照
type BusinessLicenseInfo struct {
	LicenseCopy   string `json:"license_copy"`   // 营业执照照片
	LicenseNumber string `json:"license_number"` // 注册号/统一社会信用代码
	MerchantName  string `json:"merchant_name"`  // 商户名称
	LegalPerson   string `json:"legal_person"`   // 个体户经营者/法人姓名
}

// CertificateInfo 登记证书
type CertificateInfo struct {
	CertCopy       string `json:"cert_copy"`       // 登记证书照片
	CertType       string `json:"cert_type"`       // 登记证书类型
	CertNumber     string `json:"cert_number"`     // 证书号
	MerchantName   string `json:"merchant_name"`   // 商户名称
	CompanyAddress string `json:"company_address"` // 注册地址
	LegalPerson    string `json:"legal_person"`    // 法人姓名
	PeriodBegin    string `json:"period_begin"`    // 有效期限开始日期
	PeriodEnd      string `json:"period_end"`      // 有效期结束日期
}

// OrganizationInfo 组织机构代码证
type OrganizationInfo struct {
	OrganizationCopy string `json:"organization_copy"` // 组织机构代码证照片
	OrganizationCode string `json:"organization_code"` // 组织机构代码
	OrgPeriodBegin   string `json:"org_period_begin"`  // 组织机构代码证有效期开始日期
	OrgPeriodEnd     string `json:"org_period_end"`    // 组织机构代码证有效期结束日期
}

// IdentityInfo 经营者/法人身份证件
type IdentityInfo struct {
	IdDocType  string      `json:"id_doc_type"`            // 证件类型
	IdCardInfo *IdCardInfo `json:"id_card_info,omitempty"` // 身份证信息
	IdDocInfo  *IdDocInfo  `json:"id_doc_info,omitempty"`  // 其他类型证件信息
	Owner      bool        `json:"owner"`                  // 经营者/法人是否为受益人
}

// IdCardInfo 身份证信息
type IdCardInfo struct {
	IdCardCopy      string `json:"id_card_copy"`      // 身份证人像面照片, 需要加密
	IdCardNational  string `json:"id_card_national"`  // 身份证国徽面照片，需要加密
	IdCardName      string `json:"id_card_name"`      // 身份证姓名，需要加密
	IdCardNumber    string `json:"id_card_number"`    // 身份证号码，需要加密
	CardPeriodBegin string `json:"card_period_begin"` // 身份证有效期开始时间
	CardPeriodEnd   string `json:"card_period_end"`   // 身份证有效期结束日期
}

// IdDocInfo 其他类型证件信息
type IdDocInfo struct {
	IdDocCopy      string `json:"id_doc_copy"`      // 证件照片
	IdDocName      string `json:"id_doc_name"`      // 证件姓名，需要加密
	IdDocNumber    string `json:"id_doc_number"`    // 证件号码，需要加密
	DocPeriodBegin string `json:"doc_period_begin"` // 证件有效期开始时间
	DocPeriodEnd   string `json:"doc_period_end"`   // 证件有效期结束日期
}

// UboInfo 最终受益人信息
type UboInfo struct {
	IdType         string `json:"id_type"`                    // 证件类型
	IdCardCopy     string `json:"id_card_copy,omitempty"`     // 身份证人像面照片
	IdCardNational string `json:"id_card_national,omitempty"` // 身份证国徽面照片
	IdDocCopy      string `json:"id_doc_copy,omitempty"`      // 证件照片
	Name           string `json:"name"`                       // 受益人姓名，需要加密
	IdNumber       string `json:"id_number"`                  // 证件号码，需要加密
	IdPeriodBegin  string `json:"id_period_begin"`            // 证件有效期开始时间
	IdPeriodEnd    string `json:"id_period_end"`              // 证件有效期结束时间
}

// BusinessInfo 经营资料
type BusinessInfo struct {
	MerchantShortname string     `json:"merchant_shortname"` // 商户简称
	ServicePhone      string     `json:"service_phone"`      // 客服电话
	SalesInfo         *SalesInfo `json:"sales_info"`         // 经营场景
}

// SalesInfo 经营场景
type SalesInfo struct {
	SalesScenesType []string         `json:"sales_scenes_type"`           // 经营场景类型
	BizStoreInfo    *BizStoreInfo    `json:"biz_store_info,omitempty"`    // 线下门店场景
	MpInfo          *MpInfo          `json:"mp_info,omitempty"`           // 公众号场景
	MiniProgramInfo *MiniProgramInfo `json:"mini_program_info,omitempty"` // 小程序场景
	AppInfo         *AppInfo         `json:"app_info,omitempty"`          // APP场景
	WebInfo         *WebInfo         `json:"web_info,omitempty"`          // 互联网网站场景
	WeworkInfo      *WeworkInfo      `json:"wework_info,omitempty"`       // 企业微信场景
}

// BizStoreInfo 线下门店场景
type BizStoreInfo struct {
	BizStoreName     string   `json:"biz_store_name"`          // 门店名称
	BizAddressCode   string   `json:"biz_address_code"`        // 门店省市编码
	BizStoreAddress  string   `json:"biz_store_address"`       // 门店地址
	StoreEntrancePic []string `json:"store_entrance_pic"`      // 门店门头照片
	IndoorPic        []string `json:"indoor_pic"`              // 店内环境照片
	BizSubAppid      string   `json:"biz_sub_appid,omitempty"` // 线下场所对应的商家APPID
}

// MpInfo 公众号场景
type MpInfo struct {
	MpAppId    string   `json:"mp_appid,omitempty"`     // 服务商公众号APPID
	MpSubAppId string   `json:"mp_sub_appid,omitempty"` // 商家公众号APPID
	MpPics     []string `json:"mp_pics,omitempty"`      // 公众号页面截图
}

// MiniProgramInfo 小程序场景
type MiniProgramInfo struct {
	MiniProgramAppId    string   `json:"mini_program_appid,omitempty"`     // 服务商小程序APPID
	MiniProgramSubAppId string   `json:"mini_program_sub_appid,omitempty"` // 商家小程序APPID
	MiniProgramPics     []string `json:"mini_program_pics,omitempty"`      // 小程序截图
}

// AppInfo APP场景
type AppInfo struct {
	AppAppId    string   `json:"app_appid,omitempty"`     // 服务商应用APPID
	AppSubAppId string   `json:"app_sub_appid,omitempty"` // 商家应用APPID
	AppPics     []string `json:"app_pics,omitempty"`      // APP截图
}

// WebInfo 互联网网站场景
type WebInfo struct {
	Domain           string `json:"domain"`                      // 互联网网站域名
	WebAuthorisation string `json:"web_authorisation,omitempty"` // 网站授权函
	WebAppId         string `json:"web_appid,omitempty"`         // 互联网网站对应的商家APPID
}

// WeworkInfo 企业微信场景信息
type WeworkInfo struct {
	SubCorpId  string   `json:"sub_corp_id"`           // 商家企业微信CorpID
	WeworkPics []string `json:"wework_pics,omitempty"` // 企业微信页面截图
}

// SettlementInfo 结算规则
type SettlementInfo struct {
	SettlementId        string   `json:"settlement_id"`                  // 入驻结算规则ID
	QualificationType   string   `json:"qualification_type"`             // 所属行业
	Qualifications      []string `json:"qualifications,omitempty"`       // 特殊资质图片
	ActivitiesId        string   `json:"activities_id,omitempty"`        // 优惠费率活动ID
	ActivitiesRate      string   `json:"activities_rate,omitempty"`      // 优惠费率活动值
	ActivitiesAdditions []string `json:"activities_additions,omitempty"` // 优惠费率活动补充材料
}

// BankAccountInfo 结算银行账户
type BankAccountInfo struct {
	BankAccountType string `json:"bank_account_type"`        // 账户类型
	AccountName     string `json:"account_name"`             // 开户名称，需要加密
	AccountBank     string `json:"account_bank"`             // 开户银行
	BankAddressCode string `json:"bank_address_code"`        // 开户银行省市编码
	BankBranchId    string `json:"bank_branch_id,omitempty"` // 开户银行联行号
	BankName        string `json:"bank_name,omitempty"`      // 开户银行全称（含支行）
	AccountNumber   string `json:"account_number"`           // 银行账号
}

// AdditionInfo 补充材料
type AdditionInfo struct {
	LegalPersonCommitment string   `json:"legal_person_commitment,omitempty"` // 法人开户承诺函
	LegalPersonVideo      string   `json:"legal_person_video,omitempty"`      // 法人开户意愿视频
	BusinessAdditionPics  []string `json:"business_addition_pics,omitempty"`  // 补充材料
	BusinessAdditionMsg   string   `json:"business_addition_msg,omitempty"`   // 补充说明
}

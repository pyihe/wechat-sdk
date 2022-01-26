package apply4subject

import (
	"github.com/pyihe/go-pkg/clone"
	"github.com/pyihe/wechat-sdk/v3/model"
)

// ApplyRequest 提交申请单请求参数
type ApplyRequest struct {
	ChannelId          string              `json:"channel_id,omitempty"`    // 渠道商户号
	BusinessCode       string              `json:"business_code"`           // 业务申请编号
	ContactInfo        *ContactInfo        `json:"contact_info"`            // 联系人信息
	SubjectInfo        *SubjectInfo        `json:"subject_info"`            // 主体信息
	IdentificationInfo *IdentificationInfo `json:"identification_info"`     // 法人身份信息
	AdditionInfo       *AdditionInfo       `json:"addition_info,omitempty"` // 补充材料
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
	RequestId string // 唯一请求ID
}

// CancelRequest 撤销申请单请求参数
type CancelRequest struct {
	ApplymentId  uint64 // 申请单编号
	BusinessCode string // 业务申请编号
}

// CancelResponse 撤销申请单应答参数
type CancelResponse struct {
	model.WechatError
	RequestId string // 唯一请求ID
}

// QueryApplyResultRequest 查询申请单审核结果请求参数
type QueryApplyResultRequest struct {
	ApplymentId  uint64 // 申请单编号
	BusinessCode string // 业务申请编号
}

// QueryApplyResultResponse 查询申请单审核结果应答参数
type QueryApplyResultResponse struct {
	model.WechatError
	RequestId      string // 唯一请求ID
	ApplymentState string `json:"applyment_state,omitempty"` // 申请状态
	QrcodeData     string `json:"qrcode_data,omitempty"`     // 二维码图片
	RejectParam    string `json:"reject_param,omitempty"`    // 驳回参数
	RejectReason   string `json:"reject_reason,omitempty"`   // 驳回原因
}

// QueryMerchantStateResponse 查询商户开户意愿确认状态应答参数
type QueryMerchantStateResponse struct {
	model.WechatError
	RequestId      string // 唯一请求ID
	AuthorizeState string `json:"authorize_state,omitempty"` // 授权状态
}

// ContactInfo 联系人信息
type ContactInfo struct {
	Name         string `json:"name"`           // 联系人姓名，需要加密
	Mobile       string `json:"mobile"`         // 联系人手机号，需要加密
	IdCardNumber string `json:"id_card_number"` // 联系人身份证号码，需要加密
}

// SubjectInfo 主体信息
type SubjectInfo struct {
	SubjectType          string                  `json:"subject_type"`                     // 主体类型
	BusinessLicenceInfo  *BusinessLicenceInfo    `json:"business_licence_info,omitempty"`  // 企业营业执照信息
	CertificateInfo      *CertificateInfo        `json:"certificate_info,omitempty"`       // 登记证书信息
	CompanyProveCopy     string                  `json:"company_prove_copy,omitempty"`     // 单位证明函照片
	AssistProveInfo      *AssistProveInfo        `json:"assist_prove_info,omitempty"`      // 辅助证明材料信息
	SpecialOperationList []*SpecialOperationList `json:"special_operation_list,omitempty"` // 经营许可证
}

// BusinessLicenceInfo 营业执照
type BusinessLicenceInfo struct {
	LicenceNumber    string `json:"licence_number"`     // 营业执照注册号
	LicenceCopy      string `json:"licence_copy"`       // 营业执照照片
	MerchantName     string `json:"merchant_name"`      // 商户名称
	LegalPerson      string `json:"legal_person"`       // 法人姓名
	CompanyAddress   string `json:"company_address"`    // 注册地址
	LicenceValidDate string `json:"licence_valid_date"` // 营业执照有效日期
}

// CertificateInfo 登记证书信息
type CertificateInfo struct {
	CertType       string `json:"cert_type"`       // 证书类型
	CertNumber     string `json:"cert_number"`     // 证书编号
	CertCopy       string `json:"cert_copy"`       // 证书照片
	MerchantName   string `json:"merchant_name"`   // 商户名称
	LegalPerson    string `json:"legal_person"`    // 法人姓名
	CompanyAddress string `json:"company_address"` // 注册地址
	CertValidDate  string `json:"cert_valid_date"` // 证书有效期
}

// AssistProveInfo 辅助证明材料信息
type AssistProveInfo struct {
	MicroBizType     string `json:"micro_biz_type"`     // 小微经营类型
	StoreName        string `json:"store_name"`         // 门店名称
	StoreAddressCode string `json:"store_address_code"` // 门店省市编码
	StoreAddress     string `json:"store_address"`      // 门店地址
	StoreHeaderCopy  string `json:"store_header_copy"`  // 门店门头照片
	StoreIndoorCopy  string `json:"store_indoor_copy"`  // 店内环境照片
}

// SpecialOperationList 经营许可证
type SpecialOperationList struct {
	CategoryId        uint32   `json:"category_id"`                   // 行业ID
	OperationCopyList []string `json:"operation_copy_list,omitempty"` // 行业经营许可证资质照片
}

// IdentificationInfo 法人身份信息
type IdentificationInfo struct {
	IdentificationType      string `json:"identification_type"`                // 法人证件类型
	IdentificationName      string `json:"identification_name"`                // 证件姓名
	IdentificationNumber    string `json:"identification_number"`              // 证件号码
	IdentificationValidDate string `json:"identification_valid_date"`          // 证件有效日期
	IdentificationFrontCopy string `json:"identification_front_copy"`          // 证件正面照片
	IdentificationBackCopy  string `json:"identification_back_copy,omitempty"` // 证件反面照片
}

// AdditionInfo 附加材料
type AdditionInfo struct {
	ConfirmMchidList []string `json:"confirm_mchid_list,omitempty"` // 待确认商户号列表
}

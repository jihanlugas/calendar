package request

type PageCompanypaymentmethod struct {
	Paging
	CompanyID         string `json:"companyId" form:"companyId" query:"companyId"`
	PaymentmethodID   string `json:"paymentmethodId" form:"paymentmethodId" query:"paymentmethodId"`
	CompanyName       string `json:"companyName" form:"companyName" query:"companyName"`
	PaymentmethodName string `json:"paymentmethodName" form:"paymentmethodName" query:"paymentmethodName"`
	CreateName        string `json:"createName" form:"createName" query:"createName"`
	UpdateName        string `json:"updateName" form:"updateName" query:"updateName"`
	Preloads          string `json:"preloads" form:"preloads" query:"preloads"`
}

type CreateCompanypaymentmethod struct {
	CompanyID       string `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	PaymentmethodID string `json:"paymentmethodId" form:"paymentmethodId" query:"paymentmethodId" validate:"required"`
}

type UpdateCompanypaymentmethod struct {
	PaymentmethodID string `json:"paymentmethodId" form:"paymentmethodId" query:"paymentmethodId" validate:"required"`
}

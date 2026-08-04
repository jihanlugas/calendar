package request

type PageCustomer struct {
	Paging
	CompanyID   string `json:"companyId" form:"companyId" query:"companyId"`
	Name        string `json:"name" form:"name" query:"name" search:"name"`
	Email       string `json:"email" form:"email" query:"email" search:"email"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" query:"phoneNumber" search:"phone_number"`
	Address     string `json:"address" form:"address" query:"address" search:"address"`
	CompanyName string `json:"companyName" form:"companyName" query:"companyName" search:"company_name"`
	CreateName  string `json:"createName" form:"createName" query:"createName" search:"create_name"`
	Search      string `json:"search" form:"search" query:"search"`
	Preloads    string `json:"preloads" form:"preloads" query:"preloads"`
}

type CreateCustomer struct {
	CompanyID   string `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Email       string `json:"email" form:"email" query:"email"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" query:"phoneNumber"`
	Address     string `json:"address" form:"address" query:"address"`
}

type UpdateCustomer struct {
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Email       string `json:"email" form:"email" query:"email"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" query:"phoneNumber"`
	Address     string `json:"address" form:"address" query:"address"`
}

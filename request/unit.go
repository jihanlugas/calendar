package request

type PageUnit struct {
	Paging
	CompanyID    string `json:"companyId" form:"companyId" query:"companyId"`
	PropertyID   string `json:"propertyId" form:"propertyId" query:"propertyId"`
	Name         string `json:"name" form:"name" query:"name"`
	Description  string `json:"description" form:"description" query:"description"`
	CompanyName  string `json:"companyName" form:"companyName" query:"companyName"`
	PropertyName string `json:"propertyName" form:"propertyName" query:"propertyName"`
	CreateName   string `json:"createName" form:"createName" query:"createName"`
	Preloads     string `json:"preloads" form:"preloads" query:"preloads"`
}

type CreateUnit struct {
	CompanyID   string `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	PropertyID  string `json:"propertyId" form:"propertyId" query:"propertyId" validate:"required"`
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
}

type UpdateUnit struct {
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
}

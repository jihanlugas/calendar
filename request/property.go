package request

type PageProperty struct {
	Paging
	CompanyID   string `json:"companyId" form:"companyId" query:"companyId"`
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`
	CompanyName string `json:"companyName" form:"companyName" query:"companyName"`
	CreateName  string `json:"createName" form:"createName" query:"createName"`
	Preloads    string `json:"preloads" form:"preloads" query:"preloads"`
}

type CreateProperty struct {
	CompanyID   string                           `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	Name        string                           `json:"name" form:"name" query:"name" validate:"required"`
	Description string                           `json:"description" form:"description" query:"description" validate:""`
	Units       []CreatePropertyPropertytimeline `json:"units" form:"units" query:"units" validate:"required"`
	//CreatePropertytimeline
}

type UpdateProperty struct {
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
}

type CreatePropertyPropertytimeline struct {
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
}


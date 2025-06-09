package request

type PageProduct struct {
	Paging
	CompanyID   string `json:"companyId" form:"companyId" query:"companyId"`
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`
	Price       int64  `json:"price" form:"price" query:"price"`
	CompanyName string `json:"companyName" form:"companyName" query:"companyName"`
	CreateName  string `json:"createName" form:"createName" query:"createName"`
	Preloads    string `json:"preloads" form:"preloads" query:"preloads"`
}

type CreateProduct struct {
	CompanyID   string `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
	Price       int64  `json:"price" form:"price" query:"price" validate:"required"`
}

type UpdateProduct struct {
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
	Price       int64  `json:"price" form:"price" query:"price" validate:"required"`
}

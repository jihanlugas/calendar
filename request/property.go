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
	CompanyID      string                           `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	Name           string                           `json:"name" form:"name" query:"name" validate:"required"`
	Description    string                           `json:"description" form:"description" query:"description" validate:""`
	Price          int64                            `json:"price" form:"price" query:"price" validate:"required"`
	Propertygroups []CreatePropertyPropertytimeline `json:"propertygroups" form:"propertygroups" query:"propertygroups" validate:"required"`
	//CreatePropertytimeline
}

type UpdateProperty struct {
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
	Price       int64  `json:"price" form:"price" query:"price" validate:"required"`
}

type CreatePropertyPropertytimeline struct {
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
}

//type CreatePropertytimeline struct {
//	DefaultStartDtValue int            `json:"defaultStartDtValue" form:"defaultStartDtValue" query:"defaultStartDtValue" validate:"required"`
//	DefaultStartDtUnit  model.TimeUnit `json:"defaultStartDtUnit" form:"defaultStartDtUnit" query:"defaultStartDtUnit" validate:"required"`
//	DefaultEndDtValue   int            `json:"defaultEndDtValue" form:"defaultEndDtValue" query:"defaultEndDtValue" validate:"required"`
//	DefaultEndDtUnit    model.TimeUnit `json:"defaultEndDtUnit" form:"defaultEndDtUnit" query:"defaultEndDtUnit" validate:"required"`
//	MinZoomTimelineHour int            `json:"minZoomTimelineHour" form:"minZoomTimelineHour" query:"minZoomTimelineHour" validate:"required"`
//	MaxZoomTimelineHour int            `json:"maxZoomTimelineHour" form:"maxZoomTimelineHour" query:"maxZoomTimelineHour" validate:"required"`
//	DragSnapMin         int            `json:"dragSnapMin" form:"dragSnapMin" query:"dragSnapMin" validate:"required"`
//}

package request

import (
	"time"

	"github.com/jihanlugas/calendar/model"
)

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
	Units          []CreatePropertyPropertytimeline `json:"units" form:"units" query:"units" validate:"required"`
	Propertyprices []CreatePropertyPropertyprice    `json:"propertyprices" form:"propertyprices" query:"propertyprices" validate:"required"`
}

type CreatePropertyPropertyprice struct {
	Price     int64            `json:"price" validate:"required,min=0"`
	Weekdays  model.Int32Array `json:"weekdays" validate:"required,min=1,dive,oneof=0 1 2 3 4 5 6"`
	Priority  int              `json:"priority" validate:"required"`
	StartTime *time.Time       `json:"startTime" validate:""`
	EndTime   *time.Time       `json:"endTime" validate:""`
}

type UpdateProperty struct {
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
}

type CreatePropertyPropertytimeline struct {
	Name        string `json:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" form:"description" query:"description" validate:""`
}

type SortPropertyPrice struct {
	Propertyprices []PropertyPrice `json:"propertyprices" form:"propertyprices" query:"propertyprices" validate:"required"`
}

type PropertyPrice struct {
	ID       string `json:"id" form:"id" query:"id" validate:"required"`
	Priority int    `json:"priority" validate:"required"`
}

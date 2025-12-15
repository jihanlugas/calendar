package request

import (
	"time"

	"github.com/jihanlugas/calendar/constant"
)

type PageEvent struct {
	Paging
	CompanyID    string `json:"companyId" form:"companyId" query:"companyId"`
	PropertyID   string `json:"propertyId" form:"propertyId" query:"propertyId"`
	UnitID       string `json:"unitId" form:"unitId" query:"unitId" validate:""`
	Name         string `json:"name" form:"name" query:"name"`
	Description  string `json:"description" form:"description" query:"description"`
	CompanyName  string `json:"companyName" form:"companyName" query:"companyName"`
	PropertyName string `json:"propertyName" form:"propertyName" query:"propertyName"`
	UnitName     string `json:"unitName" form:"unitName" query:"unitName"`
	CreateName   string `json:"createName" form:"createName" query:"createName"`
	Preloads     string `json:"preloads" form:"preloads" query:"preloads"`
}

type TimelineEvent struct {
	CompanyID  string    `json:"companyId" form:"companyId" query:"companyId" validate:""`
	PropertyID string    `json:"propertyId" form:"propertyId" query:"propertyId" validate:""`
	UnitID     string    `json:"unitId" form:"unitId" query:"unitId" validate:""`
	StartDt    time.Time `json:"startDt" form:"startDt" query:"startDt" validate:"required"`
	EndDt      time.Time `json:"endDt" form:"endDt" query:"endDt" validate:"required"`
	Preloads   string    `json:"preloads" form:"preloads" query:"preloads" validate:""`
}

type CreateEvent struct {
	CompanyID   string               `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	PropertyID  string               `json:"propertyId" form:"propertyId" query:"propertyId" validate:"required"`
	UnitID      string               `json:"unitId" form:"unitId" query:"unitId" validate:"required"`
	Name        string               `json:"name" form:"name" query:"name" validate:"required"`
	Description string               `json:"description" form:"description" query:"description" validate:""`
	StartDt     time.Time            `json:"startDt" form:"startDt" query:"startDt" validate:"required"`
	EndDt       time.Time            `json:"endDt" form:"endDt" query:"endDt" validate:"required"`
	Status      constant.EventStatus `json:"status" form:"status" query:"status" validate:"required"`
	Price       int64                `json:"price" form:"price" query:"price" validate:""`
}

type UpdateEvent struct {
	UnitID      string               `json:"unitId" form:"unitId" query:"unitId" validate:"required"`
	Name        string               `json:"name" form:"name" query:"name" validate:"required"`
	Description string               `json:"description" form:"description" query:"description" validate:""`
	StartDt     time.Time            `json:"startDt" form:"startDt" query:"startDt" validate:"required"`
	EndDt       time.Time            `json:"endDt" form:"endDt" query:"endDt" validate:"required"`
	Status      constant.EventStatus `json:"status" form:"status" query:"status" validate:"required"`
}

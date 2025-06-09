package request

import "time"

type PageEvent struct {
	Paging
	CompanyID         string `json:"companyId" form:"companyId" query:"companyId"`
	PropertyID        string `json:"propertyId" form:"propertyId" query:"propertyId"`
	PropertygroupID   string `json:"propertygroupId" form:"propertygroupId" query:"propertygroupId" validate:""`
	Name              string `json:"name" form:"name" query:"name"`
	Description       string `json:"description" form:"description" query:"description"`
	CompanyName       string `json:"companyName" form:"companyName" query:"companyName"`
	PropertyName      string `json:"propertyName" form:"propertyName" query:"propertyName"`
	PropertygroupName string `json:"propertygroupName" form:"propertygroupName" query:"propertygroupName"`
	CreateName        string `json:"createName" form:"createName" query:"createName"`
	Preloads          string `json:"preloads" form:"preloads" query:"preloads"`
}

type TimelineEvent struct {
	CompanyID       string    `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	PropertyID      string    `json:"propertyId" form:"propertyId" query:"propertyId" validate:"required"`
	PropertygroupID string    `json:"propertygroupId" form:"propertygroupId" query:"propertygroupId" validate:""`
	StartDt         time.Time `json:"startDt" form:"startDt" query:"startDt" validate:"required"`
	EndDt           time.Time `json:"endDt" form:"endDt" query:"endDt" validate:"required"`
	Preloads        string    `json:"preloads" form:"preloads" query:"preloads" validate:""`
}

type CreateEvent struct {
	CompanyID       string    `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	PropertyID      string    `json:"propertyId" form:"propertyId" query:"propertyId" validate:"required"`
	PropertygroupID string    `json:"propertygroupId" form:"propertygroupId" query:"propertygroupId" validate:"required"`
	Name            string    `json:"name" form:"name" query:"name" validate:"required"`
	Description     string    `json:"description" form:"description" query:"description" validate:""`
	StartDt         time.Time `json:"startDt" form:"startDt" query:"startDt" validate:"required"`
	EndDt           time.Time `json:"endDt" form:"endDt" query:"endDt" validate:"required"`
}

type UpdateEvent struct {
	PropertygroupID string    `json:"propertygroupId" form:"propertygroupId" query:"propertygroupId" validate:"required"`
	Name            string    `json:"name" form:"name" query:"name" validate:"required"`
	Description     string    `json:"description" form:"description" query:"description" validate:""`
	StartDt         time.Time `json:"startDt" form:"startDt" query:"startDt" validate:"required"`
	EndDt           time.Time `json:"endDt" form:"endDt" query:"endDt" validate:"required"`
}

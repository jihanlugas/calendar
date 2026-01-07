package request

import (
	"time"

	"github.com/jihanlugas/calendar/model"
)

type CreatePropertyprice struct {
	CompanyID  string           `json:"companyId" validate:"required"`
	PropertyID string           `json:"propertyId" validate:"required"`
	Price      int64            `json:"price" validate:"required,min=0"`
	Weekdays   model.Int32Array `json:"weekdays" validate:"required,min=1,dive,oneof=0 1 2 3 4 5 6"`
	StartTime  *time.Time       `json:"startTime" validate:""`
	EndTime    *time.Time       `json:"endTime" validate:""`
}

type UpdatePropertyprice struct {
	Price     int64            `json:"price" validate:"required,min=0"`
	Weekdays  model.Int32Array `json:"weekdays" validate:"required,min=1,dive,oneof=0 1 2 3 4 5 6"`
	StartTime *time.Time       `json:"startTime" validate:""`
	EndTime   *time.Time       `json:"endTime" validate:""`
}

type GetPrice struct {
	PropertyID string    `json:"propertyId" validate:"required"`
	StartDt    time.Time `json:"startDt" validate:"required"`
	EndDt      time.Time `json:"endDt" validate:"required"`
}

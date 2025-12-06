package request

import "time"

type GetPrice struct {
	PropertyID string    `json:"propertyId" validate:"required"`
	StartDt    time.Time `json:"startDt" validate:"required"`
	EndDt      time.Time `json:"endDt" validate:"required"`
}

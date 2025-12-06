package propertyprice

import (
	"errors"
	"fmt"
	"time"

	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"gorm.io/gorm"
)

type Repository interface {
	Name() string
	GetTableById(conn *gorm.DB, id string, preloads ...string) (tPropertyprice model.Propertyprice, err error)
	GetViewById(conn *gorm.DB, id string, preloads ...string) (vPropertyprice model.PropertypriceView, err error)
	Create(conn *gorm.DB, tPropertyprice model.Propertyprice) error
	Creates(conn *gorm.DB, tPropertyprices []model.Propertyprice) error
	Update(conn *gorm.DB, tPropertyprice model.Propertyprice) error
	Save(conn *gorm.DB, tPropertyprice model.Propertyprice) error
	Delete(conn *gorm.DB, tPropertyprice model.Propertyprice) error
	GetPrice(conn *gorm.DB, req request.GetPrice) (price int64, err error)
}

type repository struct {
}

func (r repository) Name() string {
	return "propertyprice"
}

func (r repository) GetTableById(conn *gorm.DB, id string, preloads ...string) (tPropertyprice model.Propertyprice, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&tPropertyprice).Error
	return tPropertyprice, err
}

func (r repository) GetViewById(conn *gorm.DB, id string, preloads ...string) (vPropertyprice model.PropertypriceView, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&vPropertyprice).Error
	return vPropertyprice, err
}

func (r repository) Create(conn *gorm.DB, tPropertyprice model.Propertyprice) error {
	return conn.Create(&tPropertyprice).Error
}

func (r repository) Creates(conn *gorm.DB, tPropertyprices []model.Propertyprice) error {
	return conn.Create(&tPropertyprices).Error
}

func (r repository) Update(conn *gorm.DB, tPropertyprice model.Propertyprice) error {
	return conn.Model(&tPropertyprice).Updates(&tPropertyprice).Error
}

func (r repository) Save(conn *gorm.DB, tPropertyprice model.Propertyprice) error {
	return conn.Save(&tPropertyprice).Error
}

func (r repository) Delete(conn *gorm.DB, tPropertyprice model.Propertyprice) error {
	return conn.Delete(&tPropertyprice).Error
}

func (r repository) GetPrice(conn *gorm.DB, req request.GetPrice) (int64, error) {
	var prices []model.Propertyprice

	if err := conn.Where("property_id = ?", req.PropertyID).Order("priority DESC").Find(&prices).Error; err != nil {
		return 0, err
	}
	if len(prices) == 0 {
		return 0, errors.New("no price configuration found for property")
	}

	totalPrice := int64(0)
	current := req.StartDt

	for current.Before(req.EndDt) {

		nextHour := current.Add(time.Hour)
		if nextHour.After(req.EndDt) {
			nextHour = req.EndDt
		}

		weekday := int32(current.Weekday())

		var selected *model.Propertyprice

		for _, price := range prices {
			if !contains(price.Weekdays, weekday) {
				continue
			}

			// ===== FULL DAY MODE =====
			if price.StartTime == nil || price.EndTime == nil {
				// langsung match tanpa cek jam
				selected = &price
				break
			}

			// ===== TIME SLOT MODE =====
			rStart := time.Date(0, 1, 1, price.StartTime.Hour(), price.StartTime.Minute(), 0, 0, time.UTC)
			rEnd := time.Date(0, 1, 1, price.EndTime.Hour(), price.EndTime.Minute(), 0, 0, time.UTC)
			now := time.Date(0, 1, 1, current.Hour(), current.Minute(), 0, 0, time.UTC)

			if (now.Equal(rStart) || now.After(rStart)) && now.Before(rEnd) {
				selected = &price
				break
			}
		}

		if selected == nil {
			return 0, fmt.Errorf("no matching price for %s", current)
		}

		// harga/jam
		totalPrice += selected.Price
		current = nextHour
	}

	return totalPrice, nil
}

func contains(arr []int32, v int32) bool {
	for _, n := range arr {
		if n == v {
			return true
		}
	}
	return false
}

func NewRepository() Repository {
	return &repository{}
}

package event

import (
	"fmt"
	"strings"

	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"gorm.io/gorm"
)

type Repository interface {
	Name() string
	GetTableById(conn *gorm.DB, id string, preloads ...string) (tEvent model.Event, err error)
	GetViewById(conn *gorm.DB, id string, preloads ...string) (vEvent model.EventView, err error)
	Create(conn *gorm.DB, tEvent model.Event) error
	Creates(conn *gorm.DB, tEvents []model.Event) error
	Update(conn *gorm.DB, tEvent model.Event) error
	Save(conn *gorm.DB, tEvent model.Event) error
	Delete(conn *gorm.DB, tEvent model.Event) error
	Page(conn *gorm.DB, req request.PageEvent) (vEvents []model.EventView, count int64, err error)
	Timeline(conn *gorm.DB, req request.TimelineEvent) (vEvents []model.EventView, err error)
}

type repository struct {
}

func (r repository) Name() string {
	return "event"
}

func (r repository) GetTableById(conn *gorm.DB, id string, preloads ...string) (tEvent model.Event, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&tEvent).Error
	return tEvent, err
}

func (r repository) GetViewById(conn *gorm.DB, id string, preloads ...string) (vEvent model.EventView, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&vEvent).Error
	return vEvent, err
}

func (r repository) Create(conn *gorm.DB, tEvent model.Event) error {
	return conn.Create(&tEvent).Error
}

func (r repository) Creates(conn *gorm.DB, tEvents []model.Event) error {
	return conn.Create(&tEvents).Error
}

func (r repository) Update(conn *gorm.DB, tEvent model.Event) error {
	return conn.Model(&tEvent).Updates(&tEvent).Error
}

func (r repository) Save(conn *gorm.DB, tEvent model.Event) error {
	return conn.Save(&tEvent).Error
}

func (r repository) Delete(conn *gorm.DB, tEvent model.Event) error {
	return conn.Delete(&tEvent).Error
}

func (r repository) DeleteByOrderId(conn *gorm.DB, id string) error {
	return conn.Where("order_id = ? ", id).Delete(&model.Event{}).Error
}

func (r repository) Page(conn *gorm.DB, req request.PageEvent) (vEvents []model.EventView, count int64, err error) {
	query := conn.Model(&vEvents)

	// preloads
	if req.Preloads != "" {
		preloads := strings.Split(req.Preloads, ",")
		for _, preload := range preloads {
			query = query.Preload(preload)
		}
	}

	// query
	if req.CompanyID != "" {
		query = query.Where("company_id = ?", req.CompanyID)
	}
	if req.PropertyID != "" {
		query = query.Where("property_id = ?", req.PropertyID)
	}
	if req.UnitID != "" {
		query = query.Where("unit_id = ?", req.UnitID)
	}
	if req.Name != "" {
		query = query.Where("name ILIKE ?", "%"+req.Name+"%")
	}
	if req.Description != "" {
		query = query.Where("description ILIKE ?", "%"+req.Description+"%")
	}
	if req.CompanyName != "" {
		query = query.Where("company_name ILIKE ?", "%"+req.CompanyName+"%")
	}
	if req.PropertyName != "" {
		query = query.Where("property_name ILIKE ?", "%"+req.PropertyName+"%")
	}
	if req.UnitName != "" {
		query = query.Where("unit_name ILIKE ?", "%"+req.UnitName+"%")
	}
	if req.CreateName != "" {
		query = query.Where("create_name ILIKE ?", "%"+req.CreateName+"%")
	}

	err = query.Count(&count).Error
	if err != nil {
		return vEvents, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "name", "asc"))
	}

	if req.Limit >= 0 {
		query = query.Offset((req.GetPage() - 1) * req.GetLimit()).Limit(req.GetLimit())
	}

	err = query.Find(&vEvents).Error
	if err != nil {
		return vEvents, count, err
	}

	return vEvents, count, err
}

func (r repository) Timeline(conn *gorm.DB, req request.TimelineEvent) (vEvents []model.EventView, err error) {
	query := conn.Model(&vEvents)

	// preloads
	if req.Preloads != "" {
		preloads := strings.Split(req.Preloads, ",")
		for _, preload := range preloads {
			query = query.Preload(preload)
		}
	}

	// query
	query = query.Where("end_dt >= ?", req.StartDt)
	query = query.Where("start_dt <= ?", req.EndDt)
	if req.CompanyID != "" {
		query = query.Where("company_id = ?", req.CompanyID)
	}
	if req.PropertyID != "" {
		query = query.Where("property_id = ?", req.PropertyID)
	}
	if req.UnitID != "" {
		query = query.Where("unit_id = ?", req.UnitID)
	}

	err = query.Find(&vEvents).Error
	if err != nil {
		return vEvents, err
	}

	return vEvents, err
}

func NewRepository() Repository {
	return &repository{}
}

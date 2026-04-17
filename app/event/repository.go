package event

import (
	"fmt"
	"strings"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"gorm.io/gorm"
)

type Repository interface {
	base.Repository[model.Event, model.EventView]
	Page(conn *gorm.DB, req request.PageEvent) (vEvents []model.EventView, count int64, err error)
	Timeline(conn *gorm.DB, req request.TimelineEvent) (vEvents []model.EventView, err error)
}

type repository struct {
	base.Repository[model.Event, model.EventView]
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
	return &repository{
		Repository: base.NewRepository[model.Event, model.EventView]("event"),
	}
}

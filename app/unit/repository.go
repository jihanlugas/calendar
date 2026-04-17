package unit

import (
	"fmt"
	"strings"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"gorm.io/gorm"
)

type Repository interface {
	base.Repository[model.Unit, model.UnitView]
	Page(conn *gorm.DB, req request.PageUnit) (vUnits []model.UnitView, count int64, err error)
}

type repository struct {
	base.Repository[model.Unit, model.UnitView]
}

func (r repository) Page(conn *gorm.DB, req request.PageUnit) (vUnits []model.UnitView, count int64, err error) {
	query := conn.Model(&vUnits)

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
	if req.CreateName != "" {
		query = query.Where("create_name ILIKE ?", "%"+req.CreateName+"%")
	}

	err = query.Count(&count).Error
	if err != nil {
		return vUnits, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "name", "asc"))
	}

	if req.Limit >= 0 {
		query = query.Offset((req.GetPage() - 1) * req.GetLimit()).Limit(req.GetLimit())
	}

	err = query.Find(&vUnits).Error
	if err != nil {
		return vUnits, count, err
	}

	return vUnits, count, err
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Unit, model.UnitView]("unit"),
	}
}

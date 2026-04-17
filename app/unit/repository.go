package unit

import (
	"fmt"
	"strings"

	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"gorm.io/gorm"
)

type Repository interface {
	Name() string
	GetTableById(conn *gorm.DB, id string, preloads ...string) (tUnit model.Unit, err error)
	GetViewById(conn *gorm.DB, id string, preloads ...string) (vUnit model.UnitView, err error)
	Create(conn *gorm.DB, tUnit model.Unit) error
	Creates(conn *gorm.DB, tUnits []model.Unit) error
	Update(conn *gorm.DB, tUnit model.Unit) error
	Save(conn *gorm.DB, tUnit model.Unit) error
	Delete(conn *gorm.DB, tUnit model.Unit) error
	Page(conn *gorm.DB, req request.PageUnit) (vUnits []model.UnitView, count int64, err error)
}

type repository struct {
}

func (r repository) Name() string {
	return "unit"
}

func (r repository) GetTableById(conn *gorm.DB, id string, preloads ...string) (tUnit model.Unit, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&tUnit).Error
	return tUnit, err
}

func (r repository) GetViewById(conn *gorm.DB, id string, preloads ...string) (vUnit model.UnitView, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&vUnit).Error
	return vUnit, err
}

func (r repository) Create(conn *gorm.DB, tUnit model.Unit) error {
	return conn.Create(&tUnit).Error
}

func (r repository) Creates(conn *gorm.DB, tUnits []model.Unit) error {
	return conn.Create(&tUnits).Error
}

func (r repository) Update(conn *gorm.DB, tUnit model.Unit) error {
	return conn.Model(&tUnit).Updates(&tUnit).Error
}

func (r repository) Save(conn *gorm.DB, tUnit model.Unit) error {
	return conn.Save(&tUnit).Error
}

func (r repository) Delete(conn *gorm.DB, tUnit model.Unit) error {
	return conn.Delete(&tUnit).Error
}

func (r repository) DeleteByOrderId(conn *gorm.DB, id string) error {
	return conn.Where("order_id = ? ", id).Delete(&model.Unit{}).Error
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
	return &repository{}
}

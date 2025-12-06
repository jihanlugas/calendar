package property

import (
	"fmt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"gorm.io/gorm"
	"strings"
)

type Repository interface {
	Name() string
	GetTableById(conn *gorm.DB, id string, preloads ...string) (tProperty model.Property, err error)
	GetViewById(conn *gorm.DB, id string, preloads ...string) (vProperty model.PropertyView, err error)
	Create(conn *gorm.DB, tProperty model.Property) error
	Creates(conn *gorm.DB, tProperties []model.Property) error
	Update(conn *gorm.DB, tProperty model.Property) error
	Save(conn *gorm.DB, tProperty model.Property) error
	Delete(conn *gorm.DB, tProperty model.Property) error
	Page(conn *gorm.DB, req request.PageProperty) (vProperties []model.PropertyView, count int64, err error)
}

type repository struct {
}

func (r repository) Name() string {
	return "property"
}

func (r repository) GetTableById(conn *gorm.DB, id string, preloads ...string) (tProperty model.Property, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}

	err = conn.Where("id = ? ", id).First(&tProperty).Error
	return tProperty, err
}

func (r repository) GetViewById(conn *gorm.DB, id string, preloads ...string) (vProperty model.PropertyView, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&vProperty).Error
	return vProperty, err
}

func (r repository) Create(conn *gorm.DB, tProperty model.Property) error {
	return conn.Create(&tProperty).Error
}

func (r repository) Creates(conn *gorm.DB, tProperties []model.Property) error {
	return conn.Create(&tProperties).Error
}

func (r repository) Update(conn *gorm.DB, tProperty model.Property) error {
	return conn.Model(&tProperty).Updates(&tProperty).Error
}

func (r repository) Save(conn *gorm.DB, tProperty model.Property) error {
	return conn.Save(&tProperty).Error
}

func (r repository) Delete(conn *gorm.DB, tProperty model.Property) error {
	return conn.Delete(&tProperty).Error
}

func (r repository) DeleteByOrderId(conn *gorm.DB, id string) error {
	return conn.Where("order_id = ? ", id).Delete(&model.Property{}).Error
}

func (r repository) Page(conn *gorm.DB, req request.PageProperty) (vProperties []model.PropertyView, count int64, err error) {
	query := conn.Model(&vProperties)

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
	if req.Name != "" {
		query = query.Where("name ILIKE ?", "%"+req.Name+"%")
	}
	if req.Description != "" {
		query = query.Where("description ILIKE ?", "%"+req.Description+"%")
	}
	if req.CompanyName != "" {
		query = query.Where("company_name ILIKE ?", "%"+req.CompanyName+"%")
	}
	if req.CreateName != "" {
		query = query.Where("create_name ILIKE ?", "%"+req.CreateName+"%")
	}

	err = query.Count(&count).Error
	if err != nil {
		return vProperties, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "name", "asc"))
	}

	if req.Limit >= 0 {
		query = query.Offset((req.GetPage() - 1) * req.GetLimit()).Limit(req.GetLimit())
	}

	err = query.Find(&vProperties).Error
	if err != nil {
		return vProperties, count, err
	}

	return vProperties, count, err
}

func NewRepository() Repository {
	return &repository{}
}

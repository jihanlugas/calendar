package propertygroup

import (
	"fmt"
	"strings"

	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"gorm.io/gorm"
)

type Repository interface {
	Name() string
	GetTableById(conn *gorm.DB, id string, preloads ...string) (tPropertygroup model.Propertygroup, err error)
	GetViewById(conn *gorm.DB, id string, preloads ...string) (vPropertygroup model.PropertygroupView, err error)
	Create(conn *gorm.DB, tPropertygroup model.Propertygroup) error
	Creates(conn *gorm.DB, tPropertygroups []model.Propertygroup) error
	Update(conn *gorm.DB, tPropertygroup model.Propertygroup) error
	Save(conn *gorm.DB, tPropertygroup model.Propertygroup) error
	Delete(conn *gorm.DB, tPropertygroup model.Propertygroup) error
	Page(conn *gorm.DB, req request.PagePropertygroup) (vPropertygroups []model.PropertygroupView, count int64, err error)
}

type repository struct {
}

func (r repository) Name() string {
	return "propertygroup"
}

func (r repository) GetTableById(conn *gorm.DB, id string, preloads ...string) (tPropertygroup model.Propertygroup, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&tPropertygroup).Error
	return tPropertygroup, err
}

func (r repository) GetViewById(conn *gorm.DB, id string, preloads ...string) (vPropertygroup model.PropertygroupView, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&vPropertygroup).Error
	return vPropertygroup, err
}

func (r repository) Create(conn *gorm.DB, tPropertygroup model.Propertygroup) error {
	return conn.Create(&tPropertygroup).Error
}

func (r repository) Creates(conn *gorm.DB, tPropertygroups []model.Propertygroup) error {
	return conn.Create(&tPropertygroups).Error
}

func (r repository) Update(conn *gorm.DB, tPropertygroup model.Propertygroup) error {
	return conn.Model(&tPropertygroup).Updates(&tPropertygroup).Error
}

func (r repository) Save(conn *gorm.DB, tPropertygroup model.Propertygroup) error {
	return conn.Save(&tPropertygroup).Error
}

func (r repository) Delete(conn *gorm.DB, tPropertygroup model.Propertygroup) error {
	return conn.Delete(&tPropertygroup).Error
}

func (r repository) DeleteByOrderId(conn *gorm.DB, id string) error {
	return conn.Where("order_id = ? ", id).Delete(&model.Propertygroup{}).Error
}

func (r repository) Page(conn *gorm.DB, req request.PagePropertygroup) (vPropertygroups []model.PropertygroupView, count int64, err error) {
	query := conn.Model(&vPropertygroups)

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
		return vPropertygroups, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "name", "asc"))
	}

	if req.Limit >= 0 {
		query = query.Offset((req.GetPage() - 1) * req.GetLimit()).Limit(req.GetLimit())
	}

	err = query.Find(&vPropertygroups).Error
	if err != nil {
		return vPropertygroups, count, err
	}

	return vPropertygroups, count, err
}

func NewRepository() Repository {
	return &repository{}
}

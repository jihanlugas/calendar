package company

import (
	"fmt"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"gorm.io/gorm"
)

type Repository interface {
	base.Repository[model.Company, model.CompanyView]
	Page(conn *gorm.DB, req request.PageCompany) (vCompanies []model.CompanyView, count int64, err error)
}

type repository struct {
	base.Repository[model.Company, model.CompanyView]
}

func (r repository) Page(conn *gorm.DB, req request.PageCompany) (vCompanies []model.CompanyView, count int64, err error) {
	query := conn.Model(&vCompanies)
	if req.Name != "" {
		query = query.Where("name ILIKE ?", "%"+req.Name+"%")
	}
	if req.Description != "" {
		query = query.Where("description ILIKE ?", "%"+req.Description+"%")
	}
	if req.CreateName != "" {
		query = query.Where("create_name ILIKE ?", "%"+req.CreateName+"%")
	}

	err = query.Count(&count).Error
	if err != nil {
		return vCompanies, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "name", "asc"))
	}
	if req.Limit >= 0 {
		query = query.Offset((req.GetPage() - 1) * req.GetLimit()).Limit(req.GetLimit())
	}

	err = query.Find(&vCompanies).Error
	if err != nil {
		return vCompanies, count, err
	}

	return vCompanies, count, err
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Company, model.CompanyView]("company"),
	}
}

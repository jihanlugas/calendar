package customer

import (
	"fmt"
	"strings"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/utils"
	"gorm.io/gorm"
)

type Repository interface {
	base.Repository[model.Customer, model.CustomerView]
	Page(conn *gorm.DB, req request.PageCustomer) (vCustomers []model.CustomerView, count int64, err error)
}

type repository struct {
	base.Repository[model.Customer, model.CustomerView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Customer, model.CustomerView]("customer"),
	}
}

func (r repository) Page(conn *gorm.DB, req request.PageCustomer) (vCustomers []model.CustomerView, count int64, err error) {
	query := conn.Model(&vCustomers)

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
	if req.Email != "" {
		query = query.Where("email ILIKE ?", "%"+req.Email+"%")
	}
	if req.PhoneNumber != "" {
		query = query.Where("no_hp ILIKE ?", "%"+utils.FormatPhoneTo62(req.PhoneNumber)+"%")
	}
	if req.Address != "" {
		query = query.Where("address ILIKE ?", "%"+req.Address+"%")
	}
	if req.CompanyName != "" {
		query = query.Where("company_name ILIKE ?", "%"+req.CompanyName+"%")
	}
	if req.CreateName != "" {
		query = query.Where("create_name ILIKE ?", "%"+req.CreateName+"%")
	}

	query = base.ApplyGlobalSearch(query, req.Search, req)

	err = query.Count(&count).Error
	if err != nil {
		return vCustomers, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "name", "asc"))
	}

	if req.Limit >= 0 {
		query = query.Offset((req.GetPage() - 1) * req.GetLimit()).Limit(req.GetLimit())
	}

	err = query.Find(&vCustomers).Error
	if err != nil {
		return vCustomers, count, err
	}

	return vCustomers, count, err

}

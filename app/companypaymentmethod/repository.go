package companypaymentmethod

import (
	"fmt"
	"strings"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"gorm.io/gorm"
)

type Repository interface {
	base.Repository[model.Companypaymentmethod, model.CompanypaymentmethodView]
	Page(conn *gorm.DB, req request.PageCompanypaymentmethod) (vCompanypaymentmethods []model.CompanypaymentmethodView, count int64, err error)
}

type repository struct {
	base.Repository[model.Companypaymentmethod, model.CompanypaymentmethodView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Companypaymentmethod, model.CompanypaymentmethodView]("companypaymentmethod"),
	}
}

func (r repository) Page(conn *gorm.DB, req request.PageCompanypaymentmethod) (vCompanypaymentmethods []model.CompanypaymentmethodView, count int64, err error) {
	query := conn.Model(&vCompanypaymentmethods)

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
	if req.PaymentmethodID != "" {
		query = query.Where("paymentmethod_id = ?", req.PaymentmethodID)
	}
	if req.CompanyName != "" {
		query = query.Where("company_name ILIKE ?", "%"+req.CompanyName+"%")
	}
	if req.PaymentmethodName != "" {
		query = query.Where("paymentmethod_name ILIKE ?", "%"+req.PaymentmethodName+"%")
	}
	if req.CreateName != "" {
		query = query.Where("create_name ILIKE ?", "%"+req.CreateName+"%")
	}
	if req.UpdateName != "" {
		query = query.Where("update_name ILIKE ?", "%"+req.UpdateName+"%")
	}

	err = query.Count(&count).Error
	if err != nil {
		return
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "create_dt", "asc"))
	}

	if req.Limit >= 0 {
		query = query.Offset((req.GetPage() - 1) * req.GetLimit()).Limit(req.GetLimit())
	}

	err = query.Find(&vCompanypaymentmethods).Error
	if err != nil {
		return
	}

	return
}

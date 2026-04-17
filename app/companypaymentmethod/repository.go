package companypaymentmethod

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Companypaymentmethod, model.CompanypaymentmethodView]
}

type repository struct {
	base.Repository[model.Companypaymentmethod, model.CompanypaymentmethodView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Companypaymentmethod, model.CompanypaymentmethodView]("companypaymentmethod"),
	}
}

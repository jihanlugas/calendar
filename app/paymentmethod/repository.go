package paymentmethod

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Paymentmethod, model.PaymentmethodView]
}

type repository struct {
	base.Repository[model.Paymentmethod, model.PaymentmethodView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Paymentmethod, model.PaymentmethodView]("paymentmethod"),
	}
}

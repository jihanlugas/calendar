package orderpayment

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Orderpayment, model.OrderpaymentView]
}

type repository struct {
	base.Repository[model.Orderpayment, model.OrderpaymentView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Orderpayment, model.OrderpaymentView]("orderpayment"),
	}
}

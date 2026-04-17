package orderdiscount

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Orderdiscount, model.OrderdiscountView]
}

type repository struct {
	base.Repository[model.Orderdiscount, model.OrderdiscountView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Orderdiscount, model.OrderdiscountView]("orderdiscount"),
	}
}

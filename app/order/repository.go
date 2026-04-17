package order

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Order, model.OrderView]
}

type repository struct {
	base.Repository[model.Order, model.OrderView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Order, model.OrderView]("order"),
	}
}

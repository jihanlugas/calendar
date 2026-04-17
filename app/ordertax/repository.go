package ordertax

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Ordertax, model.OrdertaxView]
}

type repository struct {
	base.Repository[model.Ordertax, model.OrdertaxView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Ordertax, model.OrdertaxView]("ordertax"),
	}
}

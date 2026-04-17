package orderevent

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Orderevent, model.OrdereventView]
}

type repository struct {
	base.Repository[model.Orderevent, model.OrdereventView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Orderevent, model.OrdereventView]("orderevent"),
	}
}

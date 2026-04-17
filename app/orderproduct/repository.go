package orderproduct

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Orderproduct, model.OrderproductView]
}

type repository struct {
	base.Repository[model.Orderproduct, model.OrderproductView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Orderproduct, model.OrderproductView]("orderproduct"),
	}
}

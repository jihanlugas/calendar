package discount

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Discount, model.DiscountView]
}

type repository struct {
	base.Repository[model.Discount, model.DiscountView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Discount, model.DiscountView]("discount"),
	}
}

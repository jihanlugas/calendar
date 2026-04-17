package tax

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Tax, model.TaxView]
}

type repository struct {
	base.Repository[model.Tax, model.TaxView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Tax, model.TaxView]("tax"),
	}
}

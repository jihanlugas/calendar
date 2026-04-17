package propertytimeline

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Repository interface {
	base.Repository[model.Propertytimeline, model.PropertytimelineView]
}

type repository struct {
	base.Repository[model.Propertytimeline, model.PropertytimelineView]
}

func NewRepository() Repository {
	return &repository{
		Repository: base.NewRepository[model.Propertytimeline, model.PropertytimelineView]("propertytimeline"),
	}
}

package propertytimeline

import (
	"github.com/jihanlugas/calendar/model"
	"gorm.io/gorm"
)

type Repository interface {
	Name() string
	GetTableById(conn *gorm.DB, id string, preloads ...string) (tPropertytimeline model.Propertytimeline, err error)
	GetViewById(conn *gorm.DB, id string, preloads ...string) (vPropertytimeline model.PropertytimelineView, err error)
	Create(conn *gorm.DB, tPropertytimeline model.Propertytimeline) error
	Update(conn *gorm.DB, tPropertytimeline model.Propertytimeline) error
	Save(conn *gorm.DB, tPropertytimeline model.Propertytimeline) error
	Delete(conn *gorm.DB, tPropertytimeline model.Propertytimeline) error
}

type repository struct {
}

func (r repository) Name() string {
	return "propertytimeline"
}

func (r repository) GetTableById(conn *gorm.DB, id string, preloads ...string) (tPropertytimeline model.Propertytimeline, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&tPropertytimeline).Error
	return tPropertytimeline, err
}

func (r repository) GetViewById(conn *gorm.DB, id string, preloads ...string) (vPropertytimeline model.PropertytimelineView, err error) {
	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}
	err = conn.Where("id = ? ", id).First(&vPropertytimeline).Error
	return vPropertytimeline, err
}

func (r repository) Create(conn *gorm.DB, tPropertytimeline model.Propertytimeline) error {
	return conn.Create(&tPropertytimeline).Error
}

func (r repository) Update(conn *gorm.DB, tPropertytimeline model.Propertytimeline) error {
	return conn.Model(&tPropertytimeline).Updates(&tPropertytimeline).Error
}

func (r repository) Save(conn *gorm.DB, tPropertytimeline model.Propertytimeline) error {
	return conn.Save(&tPropertytimeline).Error
}

func (r repository) Delete(conn *gorm.DB, tPropertytimeline model.Propertytimeline) error {
	return conn.Delete(&tPropertytimeline).Error
}

func (r repository) DeleteByOrderId(conn *gorm.DB, id string) error {
	return conn.Where("order_id = ? ", id).Delete(&model.Propertytimeline{}).Error
}

func NewRepository() Repository {
	return &repository{}
}

package model

import (
	"fmt"
	"time"

	"github.com/jihanlugas/calendar/config"
	"github.com/jihanlugas/calendar/utils"
	"gorm.io/gorm"
)

func (m *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	if m.ID == "" {
		m.ID = utils.GetUniqueID()
	}

	if m.CreateDt.IsZero() {
		m.CreateDt = now
	}
	if m.UpdateDt.IsZero() {
		m.UpdateDt = now
	}
	return
}

func (m *Unit) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	m.UpdateDt = now
	return
}

func (m *UnitView) AfterFind(tx *gorm.DB) (err error) {
	if m.PhotoID != "" {
		m.PhotoUrl = fmt.Sprintf("%s/%s", config.FileBaseUrl, m.PhotoUrl)
	}
	return
}

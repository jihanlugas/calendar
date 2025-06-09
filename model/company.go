package model

import (
	"fmt"
	"github.com/jihanlugas/calendar/config"
	"github.com/jihanlugas/calendar/utils"
	"gorm.io/gorm"
	"time"
)

func (m *Company) BeforeCreate(tx *gorm.DB) (err error) {
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

func (m *Company) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	m.UpdateDt = now
	return
}

func (m *CompanyView) AfterFind(tx *gorm.DB) (err error) {
	if m.PhotoID != "" {
		m.PhotoUrl = fmt.Sprintf("%s/%s", config.FileBaseUrl, m.PhotoUrl)
	}
	return
}

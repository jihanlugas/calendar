package model

import (
	"time"

	"github.com/jihanlugas/calendar/utils"
	"gorm.io/gorm"
)

func (m *Propertyprice) BeforeCreate(tx *gorm.DB) (err error) {
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

func (m *Propertyprice) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	m.UpdateDt = now
	return
}

func (m *PropertypriceView) AfterFind(tx *gorm.DB) (err error) {
	if m.StartTime != nil {
		m.StartTimeFormated = m.StartTime.Format("15:04")
	}
	if m.EndTime != nil {
		m.EndTimeFormated = m.EndTime.Format("15:04")
	}
	return
}

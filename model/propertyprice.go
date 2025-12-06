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

//func (m *PropertypriceView) AfterFind(tx *gorm.DB) (err error) {
//	if m.PhotoID != "" {
//		m.PhotoUrl = fmt.Sprintf("%s/%s", config.FileBaseUrl, m.PhotoUrl)
//	}
//	return
//}

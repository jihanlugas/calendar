package model

import (
	"time"

	"github.com/jihanlugas/calendar/constant"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TimeUnit string

const (
	TimeUnitHour TimeUnit = "hour"
	TimeUnitDay  TimeUnit = "day"
	TimeUnitWeek TimeUnit = "week"
)

type Photo struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	ClientName  string         `gorm:"not null" json:"clientName"`
	ServerName  string         `gorm:"not null" json:"serverName"`
	RefTable    string         `gorm:"not null" json:"refTable"`
	Ext         string         `gorm:"not null" json:"ext"`
	PhotoPath   string         `gorm:"not null" json:"photoPath"`
	PhotoSize   int64          `gorm:"not null" json:"photoSize"`
	PhotoWidth  int64          `gorm:"not null" json:"photoWidth"`
	PhotoHeight int64          `gorm:"not null" json:"photoHeight"`
	CreateBy    string         `gorm:"not null" json:"createBy"`
	CreateDt    time.Time      `gorm:"not null" json:"createDt"`
	DeleteDt    gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Photoinc struct {
	ID        string `gorm:"primaryKey" json:"id"`
	RefTable  string `gorm:"not null" json:"refTable"`
	FolderInc int64  `gorm:"not null" json:"folderInc"`
	Folder    string `gorm:"not null" json:"folder"`
	Running   int64  `gorm:"not null" json:"running"`
}

type User struct {
	ID                string         `gorm:"primaryKey"`
	Role              string         `gorm:"not null"`
	Email             string         `gorm:"not null"`
	Username          string         `gorm:"not null"`
	PhoneNumber       string         `gorm:"not null"`
	Address           string         `gorm:"not null"`
	Fullname          string         `gorm:"not null"`
	Passwd            string         `gorm:"not null"`
	PassVersion       int            `gorm:"not null"`
	IsActive          bool           `gorm:"not null"`
	PhotoID           string         `gorm:"not null"`
	LastLoginDt       *time.Time     `gorm:"null"`
	BirthDt           *time.Time     `gorm:"null"`
	BirthPlace        string         `gorm:"not null"`
	AccountVerifiedDt *time.Time     `gorm:"null"`
	CreateBy          string         `gorm:"not null"`
	CreateDt          time.Time      `gorm:"not null"`
	UpdateBy          string         `gorm:"not null"`
	UpdateDt          time.Time      `gorm:"not null"`
	DeleteDt          gorm.DeletedAt `gorm:"null"`
}

type Company struct {
	ID          string         `gorm:"primaryKey"`
	Name        string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	Email       string         `gorm:"not null"`
	PhoneNumber string         `gorm:"not null"`
	Address     string         `gorm:"not null"`
	PhotoID     string         `gorm:"not null"`
	CreateBy    string         `gorm:"not null"`
	CreateDt    time.Time      `gorm:"not null"`
	UpdateBy    string         `gorm:"not null"`
	UpdateDt    time.Time      `gorm:"not null"`
	DeleteDt    gorm.DeletedAt `gorm:"null"`
}

type Usercompany struct {
	ID               string         `gorm:"primaryKey"`
	UserID           string         `gorm:"not null"`
	CompanyID        string         `gorm:"not null"`
	IsDefaultCompany bool           `gorm:"not null"`
	IsCreator        bool           `gorm:"not null"`
	CreateBy         string         `gorm:"not null"`
	CreateDt         time.Time      `gorm:"not null"`
	UpdateBy         string         `gorm:"not null"`
	UpdateDt         time.Time      `gorm:"not null"`
	DeleteDt         gorm.DeletedAt `gorm:"null"`
}

type Property struct {
	ID          string         `gorm:"primaryKey"`
	CompanyID   string         `gorm:"not null"`
	Name        string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	PhotoID     string         `gorm:"not null"`
	OpenTime    *time.Time     `gorm:"null"`
	CloseTime   *time.Time     `gorm:"null"`
	CreateBy    string         `gorm:"not null"`
	CreateDt    time.Time      `gorm:"not null"`
	UpdateBy    string         `gorm:"not null"`
	UpdateDt    time.Time      `gorm:"not null"`
	DeleteDt    gorm.DeletedAt `gorm:"null"`
}

type Propertyprice struct {
	ID         string         `gorm:"primaryKey"`
	CompanyID  string         `gorm:"not null"`
	PropertyID string         `gorm:"not null"`
	Priority   int            `gorm:"not null"`
	Weekdays   pq.Int32Array  `gorm:"type:int[];not null"`
	StartTime  *time.Time     `gorm:"null"`
	EndTime    *time.Time     `gorm:"null"`
	Price      int64          `gorm:"not null"`
	CreateBy   string         `gorm:"not null"`
	CreateDt   time.Time      `gorm:"not null"`
	UpdateBy   string         `gorm:"not null"`
	UpdateDt   time.Time      `gorm:"not null"`
	DeleteDt   gorm.DeletedAt `gorm:"null"`
}

type Propertytimeline struct {
	ID                  string         `gorm:"primaryKey"`
	DefaultStartDtValue int            `gorm:"not null;default:6"`
	DefaultStartDtUnit  TimeUnit       `gorm:"not null;default:Hour"`
	DefaultEndDtValue   int            `gorm:"not null;default:6"`
	DefaultEndDtUnit    TimeUnit       `gorm:"not null;default:Hour"`
	MinZoomTimelineHour int            `gorm:"not null;default:6"`
	MaxZoomTimelineHour int            `gorm:"not null;default:24"`
	DragSnapMin         int            `gorm:"not null;default:30"`
	CreateBy            string         `gorm:"not null"`
	CreateDt            time.Time      `gorm:"not null"`
	UpdateBy            string         `gorm:"not null"`
	UpdateDt            time.Time      `gorm:"not null"`
	DeleteDt            gorm.DeletedAt `gorm:"null"`
}

type Propertygroup struct {
	ID          string         `gorm:"primaryKey"`
	CompanyID   string         `gorm:"not null"`
	PropertyID  string         `gorm:"not null"`
	Name        string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	PhotoID     string         `gorm:"not null"`
	CreateBy    string         `gorm:"not null"`
	CreateDt    time.Time      `gorm:"not null"`
	UpdateBy    string         `gorm:"not null"`
	UpdateDt    time.Time      `gorm:"not null"`
	DeleteDt    gorm.DeletedAt `gorm:"null"`
}

type Event struct {
	ID              string               `gorm:"primaryKey" json:"id"`
	CompanyID       string               `gorm:"not null" json:"company_id"`
	PropertyID      string               `gorm:"not null" json:"property_id"`
	PropertygroupID string               `gorm:"not null" json:"propertygroup_id"`
	Name            string               `gorm:"not null" json:"name"`
	Description     string               `gorm:"not null" json:"description"`
	StartDt         time.Time            `gorm:"not null" json:"start_dt"`
	EndDt           time.Time            `gorm:"not null" json:"end_dt"`
	Status          constant.EventStatus `gorm:"not null" json:"status"`
	Price           int64                `gorm:"not null" json:"price"`
	CreateBy        string               `gorm:"not null" json:"create_by"`
	CreateDt        time.Time            `gorm:"not null" json:"create_dt"`
	UpdateBy        string               `gorm:"not null" json:"update_by"`
	UpdateDt        time.Time            `gorm:"not null" json:"update_dt"`
	DeleteDt        gorm.DeletedAt       `gorm:"null" json:"delete_dt"`
}

type Product struct {
	ID          string         `gorm:"primaryKey"`
	CompanyID   string         `gorm:"not null"`
	Name        string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	Price       int64          `gorm:"not null"`
	PhotoID     string         `gorm:"not null"`
	CreateBy    string         `gorm:"not null"`
	CreateDt    time.Time      `gorm:"not null"`
	UpdateBy    string         `gorm:"not null"`
	UpdateDt    time.Time      `gorm:"not null"`
	DeleteDt    gorm.DeletedAt `gorm:"null"`
}

type Transaction struct {
	ID           string         `gorm:"primaryKey"`
	CompanyID    string         `gorm:"not null"`
	TotalEvent   int64          `gorm:"not null"`
	TotalProduct int64          `gorm:"not null"`
	Total        int64          `gorm:"not null"`
	Paid         bool           `gorm:"not null"`
	CreateBy     string         `gorm:"not null"`
	CreateDt     time.Time      `gorm:"not null"`
	UpdateBy     string         `gorm:"not null"`
	UpdateDt     time.Time      `gorm:"not null"`
	DeleteDt     gorm.DeletedAt `gorm:"null"`
}

type Transactionevent struct {
	ID            string         `gorm:"primaryKey"`
	CompanyID     string         `gorm:"not null"`
	TransactionID string         `gorm:"not null"`
	EventID       string         `gorm:"not null"`
	Price         int64          `gorm:"not null"`
	Paid          bool           `gorm:"not null"`
	CreateBy      string         `gorm:"not null"`
	CreateDt      time.Time      `gorm:"not null"`
	UpdateBy      string         `gorm:"not null"`
	UpdateDt      time.Time      `gorm:"not null"`
	DeleteDt      gorm.DeletedAt `gorm:"null"`
}

type Transactionproduct struct {
	ID            string         `gorm:"primaryKey"`
	CompanyID     string         `gorm:"not null"`
	TransactionID string         `gorm:"not null"`
	EventID       string         `gorm:"not null"`
	ProductID     string         `gorm:"not null"`
	ProductName   string         `gorm:"not null"`
	Price         int64          `gorm:"not null"`
	Paid          bool           `gorm:"not null"`
	CreateBy      string         `gorm:"not null"`
	CreateDt      time.Time      `gorm:"not null"`
	UpdateBy      string         `gorm:"not null"`
	UpdateDt      time.Time      `gorm:"not null"`
	DeleteDt      gorm.DeletedAt `gorm:"null"`
}

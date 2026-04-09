package model

import (
	"time"

	"github.com/jihanlugas/calendar/constant"
	"gorm.io/gorm"
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

type Paymentmethod struct {
	ID       string         `gorm:"primaryKey"`
	Name     string         `gorm:"not null"`
	CreateBy string         `gorm:"not null"`
	CreateDt time.Time      `gorm:"not null"`
	UpdateBy string         `gorm:"not null"`
	UpdateDt time.Time      `gorm:"not null"`
	DeleteDt gorm.DeletedAt `gorm:"null"`
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

type Companypaymentmethod struct {
	ID              string         `gorm:"primaryKey"`
	CompanyID       string         `gorm:"not null"`
	PaymentmethodID string         `gorm:"not null"`
	CreateBy        string         `gorm:"not null"`
	CreateDt        time.Time      `gorm:"not null"`
	UpdateBy        string         `gorm:"not null"`
	UpdateDt        time.Time      `gorm:"not null"`
	DeleteDt        gorm.DeletedAt `gorm:"null"`
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
	Weekdays   Int32Array     `gorm:"type:int[];not null"`
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
	MinZoomTimelineHour int            `gorm:"not null;default:6"`
	MaxZoomTimelineHour int            `gorm:"not null;default:24"`
	DragSnapMin         int            `gorm:"not null;default:30"`
	CreateBy            string         `gorm:"not null"`
	CreateDt            time.Time      `gorm:"not null"`
	UpdateBy            string         `gorm:"not null"`
	UpdateDt            time.Time      `gorm:"not null"`
	DeleteDt            gorm.DeletedAt `gorm:"null"`
}

type Unit struct {
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
	ID           string               `gorm:"primaryKey" json:"id"`
	CompanyID    string               `gorm:"not null" json:"company_id"`
	PropertyID   string               `gorm:"not null" json:"property_id"`
	UnitID       string               `gorm:"not null" json:"unit_id"`
	OrdereventID string               `gorm:"not null" json:"orderevent_id"`
	Name         string               `gorm:"not null" json:"name"`
	Description  string               `gorm:"not null" json:"description"`
	StartDt      time.Time            `gorm:"not null" json:"start_dt"`
	EndDt        time.Time            `gorm:"not null" json:"end_dt"`
	Status       constant.EventStatus `gorm:"not null" json:"status"`
	CreateBy     string               `gorm:"not null" json:"create_by"`
	CreateDt     time.Time            `gorm:"not null" json:"create_dt"`
	UpdateBy     string               `gorm:"not null" json:"update_by"`
	UpdateDt     time.Time            `gorm:"not null" json:"update_dt"`
	DeleteDt     gorm.DeletedAt       `gorm:"null" json:"delete_dt"`
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

type Tax struct {
	ID          string           `gorm:"primaryKey"`
	CompanyID   string           `gorm:"not null"`
	Name        string           `gorm:"not null"`
	Type        constant.TaxType `gorm:"not null"` // percentage or fixed
	Value       int64            `gorm:"not null"`
	IsInclusive bool             `gorm:"not null"`
	IsActive    bool             `gorm:"not null"`
	CreateBy    string           `gorm:"not null"`
	CreateDt    time.Time        `gorm:"not null"`
	UpdateBy    string           `gorm:"not null"`
	UpdateDt    time.Time        `gorm:"not null"`
	DeleteDt    gorm.DeletedAt   `gorm:"null"`
}

type Discount struct {
	ID          string                `gorm:"primaryKey"`
	CompanyID   string                `gorm:"not null"`
	Name        string                `gorm:"not null"`
	Description string                `gorm:"not null"`
	Type        constant.DiscountType `gorm:"not null"` // percentage or fixed
	Value       int64                 `gorm:"not null"`
	MaxDiscount *int64                `gorm:"not null"`
	MinPurchase *int64                `gorm:"not null"`
	StartDt     *time.Time            `gorm:"null"`
	EndDt       *time.Time            `gorm:"null"`
	UsageLimit  *int                  `gorm:"not null"`
	Used        *int                  `gorm:"not null"`
	IsActive    bool                  `gorm:"not null"`
	CreateBy    string                `gorm:"not null"`
	CreateDt    time.Time             `gorm:"not null"`
	UpdateBy    string                `gorm:"not null"`
	UpdateDt    time.Time             `gorm:"not null"`
	DeleteDt    gorm.DeletedAt        `gorm:"null"`
}

type Order struct {
	ID        string         `gorm:"primaryKey"`
	CompanyID string         `gorm:"not null"`
	Tax       int64          `gorm:"not null"`
	Discount  int64          `gorm:"not null"`
	Rounding  int64          `gorm:"not null"`
	Subtotal  int64          `gorm:"not null"` // OrderEvent + OrderProduct
	Total     int64          `gorm:"not null"` // Subtotal + Tax - Discount
	Payment   int64          `gorm:"not null"`
	CreateBy  string         `gorm:"not null"`
	CreateDt  time.Time      `gorm:"not null"`
	UpdateBy  string         `gorm:"not null"`
	UpdateDt  time.Time      `gorm:"not null"`
	DeleteDt  gorm.DeletedAt `gorm:"null"`
}

type Orderevent struct {
	ID        string         `gorm:"primaryKey"`
	CompanyID string         `gorm:"not null"`
	OrderID   string         `gorm:"not null"`
	EventID   string         `gorm:"not null"`
	Total     int64          `gorm:"not null"`
	CreateBy  string         `gorm:"not null"`
	CreateDt  time.Time      `gorm:"not null"`
	UpdateBy  string         `gorm:"not null"`
	UpdateDt  time.Time      `gorm:"not null"`
	DeleteDt  gorm.DeletedAt `gorm:"null"`
}

type Orderproduct struct {
	ID        string         `gorm:"primaryKey"`
	CompanyID string         `gorm:"not null"`
	OrderID   string         `gorm:"not null"`
	ProductID string         `gorm:"not null"`
	Qty       int64          `gorm:"not null"`
	Price     int64          `gorm:"not null"`
	Total     int64          `gorm:"not null"`
	CreateBy  string         `gorm:"not null"`
	CreateDt  time.Time      `gorm:"not null"`
	UpdateBy  string         `gorm:"not null"`
	UpdateDt  time.Time      `gorm:"not null"`
	DeleteDt  gorm.DeletedAt `gorm:"null"`
}

type Ordertax struct {
	ID        string         `gorm:"primaryKey"`
	CompanyID string         `gorm:"not null"`
	OrderID   string         `gorm:"not null"`
	TaxID     string         `gorm:"not null"`
	Total     int64          `gorm:"not null"`
	CreateBy  string         `gorm:"not null"`
	CreateDt  time.Time      `gorm:"not null"`
	UpdateBy  string         `gorm:"not null"`
	UpdateDt  time.Time      `gorm:"not null"`
	DeleteDt  gorm.DeletedAt `gorm:"null"`
}

type Orderdiscount struct {
	ID         string         `gorm:"primaryKey"`
	CompanyID  string         `gorm:"not null"`
	OrderID    string         `gorm:"not null"`
	DiscountID string         `gorm:"not null"`
	Total      int64          `gorm:"not null"`
	CreateBy   string         `gorm:"not null"`
	CreateDt   time.Time      `gorm:"not null"`
	UpdateBy   string         `gorm:"not null"`
	UpdateDt   time.Time      `gorm:"not null"`
	DeleteDt   gorm.DeletedAt `gorm:"null"`
}

type Orderpayment struct {
	ID              string         `gorm:"primaryKey"`
	CompanyID       string         `gorm:"not null"`
	OrderID         string         `gorm:"not null"`
	PaymentmethodID string         `gorm:"not null"`
	Total           int64          `gorm:"not null"`
	CreateBy        string         `gorm:"not null"`
	CreateDt        time.Time      `gorm:"not null"`
	UpdateBy        string         `gorm:"not null"`
	UpdateDt        time.Time      `gorm:"not null"`
	DeleteDt        gorm.DeletedAt `gorm:"null"`
}

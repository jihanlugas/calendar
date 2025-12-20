package model

import (
	"time"

	"github.com/jihanlugas/calendar/constant"
	"gorm.io/gorm"
)

type UserView struct {
	ID                string         `json:"id"`
	CompanyID         string         `json:"companyId"`
	UsercompanyID     string         `json:"usercompanyId"`
	Role              string         `json:"role"`
	Email             string         `json:"email"`
	Username          string         `json:"username"`
	PhoneNumber       string         `json:"phoneNumber"`
	Address           string         `json:"address"`
	Fullname          string         `json:"fullname"`
	Passwd            string         `json:"-"`
	PassVersion       int            `json:"passVersion"`
	IsActive          bool           `json:"isActive"`
	PhotoID           string         `json:"photoId"`
	PhotoUrl          string         `json:"photoUrl"`
	LastLoginDt       *time.Time     `json:"lastLoginDt"`
	BirthDt           *time.Time     `json:"birthDt"`
	BirthPlace        string         `json:"birthPlace"`
	AccountVerifiedDt *time.Time     `json:"accountVerifiedDt"`
	CreateBy          string         `json:"createBy"`
	CreateDt          time.Time      `json:"createDt"`
	UpdateBy          string         `json:"updateBy"`
	UpdateDt          time.Time      `json:"updateDt"`
	DeleteDt          gorm.DeletedAt `json:"deleteDt"`
	CreateName        string         `json:"createName"`
	UpdateName        string         `json:"updateName"`

	Company       *CompanyView      `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Usercompanies []UsercompanyView `json:"usercompanies,omitempty" gorm:"foreignKey:UserID"`
}

func (UserView) TableName() string {
	return VIEW_USER
}

type CompanyView struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Email       string         `json:"email"`
	PhoneNumber string         `json:"phoneNumber"`
	Address     string         `json:"address"`
	PhotoID     string         `json:"photoId"`
	PhotoUrl    string         `json:"photoUrl"`
	CreateBy    string         `json:"createBy"`
	CreateDt    time.Time      `json:"createDt"`
	UpdateBy    string         `json:"updateBy"`
	UpdateDt    time.Time      `json:"updateDt"`
	DeleteDt    gorm.DeletedAt `json:"deleteDt"`
	CreateName  string         `json:"createName"`
	UpdateName  string         `json:"updateName"`

	Users      []UserView     `json:"users" gorm:"foreignKey:CompanyID"`
	Properties []PropertyView `json:"properties" gorm:"foreignKey:CompanyID"`
	Units      []UnitView     `json:"units" gorm:"foreignKey:CompanyID"`
	Products   []ProductView  `json:"products" gorm:"foreignKey:CompanyID"`
	Events     []EventView    `json:"events" gorm:"foreignKey:CompanyID"`
}

func (CompanyView) TableName() string {
	return VIEW_COMPANY
}

type UsercompanyView struct {
	ID               string         `json:"id"`
	UserID           string         `json:"userId"`
	CompanyID        string         `json:"companyId"`
	IsDefaultCompany bool           `json:"isDefaultCompany"`
	IsCreator        bool           `json:"isCreator"`
	CreateBy         string         `json:"createBy"`
	CreateDt         time.Time      `json:"createDt"`
	UpdateBy         string         `json:"updateBy"`
	UpdateDt         time.Time      `json:"updateDt"`
	DeleteDt         gorm.DeletedAt `json:"deleteDt"`
	UserName         string         `json:"userName"`
	CompanyName      string         `json:"companyName"`
	CreateName       string         `json:"createName"`
	UpdateName       string         `json:"updateName"`

	User    *UserView    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Company *CompanyView `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
}

func (UsercompanyView) TableName() string {
	return VIEW_USERCOMPANY
}

type PropertyView struct {
	ID          string         `json:"id"`
	CompanyID   string         `json:"companyId"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	PhotoID     string         `json:"photoId"`
	PhotoUrl    string         `json:"photoUrl"`
	CreateBy    string         `json:"createBy"`
	CreateDt    time.Time      `json:"createDt"`
	UpdateBy    string         `json:"updateBy"`
	UpdateDt    time.Time      `json:"updateDt"`
	DeleteDt    gorm.DeletedAt `json:"deleteDt"`
	CompanyName string         `json:"companyName"`
	CreateName  string         `json:"createName"`
	UpdateName  string         `json:"updateName"`

	Company          *CompanyView          `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Propertytimeline *PropertytimelineView `json:"propertytimeline,omitempty" gorm:"foreignKey:ID"`
	Units            []UnitView            `json:"units" gorm:"foreignKey:PropertyID"`
	Events           []EventView           `json:"events" gorm:"foreignKey:PropertyID"`
	Propertyprices   []PropertypriceView   `json:"propertyprices" gorm:"foreignKey:PropertyID"`
}

func (PropertyView) TableName() string {
	return VIEW_PROPERTY
}

type PropertypriceView struct {
	ID           string         `json:"id"`
	CompanyID    string         `json:"companyId"`
	PropertyID   string         `json:"propertyId"`
	Priority     int            `json:"priority"`
	Weekdays     Int32Array     `json:"weekdays" gorm:"type:int[]"`
	StartTime    *time.Time     `json:"startTime"`
	EndTime      *time.Time     `json:"endTime"`
	Price        int64          `json:"price"`
	CreateBy     string         `json:"createBy"`
	CreateDt     time.Time      `json:"createDt"`
	UpdateBy     string         `json:"updateBy"`
	UpdateDt     time.Time      `json:"updateDt"`
	DeleteDt     gorm.DeletedAt `json:"deleteDt"`
	CompanyName  string         `json:"companyName"`
	PropertyName string         `json:"propertyName"`
	CreateName   string         `json:"createName"`
	UpdateName   string         `json:"updateName"`

	Company  *CompanyView  `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Property *PropertyView `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
}

func (PropertypriceView) TableName() string {
	return VIEW_PROPERTYPRICE
}

type PropertytimelineView struct {
	ID                  string         `json:"id"`
	MinZoomTimelineHour int            `json:"minZoomTimelineHour"`
	MaxZoomTimelineHour int            `json:"maxZoomTimelineHour"`
	DragSnapMin         int            `json:"dragSnapMin"`
	CreateBy            string         `json:"createBy"`
	CreateDt            time.Time      `json:"createDt"`
	UpdateBy            string         `json:"updateBy"`
	UpdateDt            time.Time      `json:"updateDt"`
	DeleteDt            gorm.DeletedAt `json:"deleteDt"`
	CreateName          string         `json:"createName"`
	UpdateName          string         `json:"updateName"`
}

func (PropertytimelineView) TableName() string {
	return VIEW_PROPERTYTIMELINE
}

type UnitView struct {
	ID           string         `json:"id"`
	CompanyID    string         `json:"companyId"`
	PropertyID   string         `json:"propertyId"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	PhotoID      string         `json:"photoId"`
	PhotoUrl     string         `json:"photoUrl"`
	CreateBy     string         `json:"createBy"`
	CreateDt     time.Time      `json:"createDt"`
	UpdateBy     string         `json:"updateBy"`
	UpdateDt     time.Time      `json:"updateDt"`
	DeleteDt     gorm.DeletedAt `json:"deleteDt"`
	CompanyName  string         `json:"companyName"`
	PropertyName string         `json:"propertyName"`
	CreateName   string         `json:"createName"`
	UpdateName   string         `json:"updateName"`

	Company  *CompanyView  `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Property *PropertyView `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
	Events   []EventView   `json:"events,omitempty" gorm:"foreignKey:UnitID"`
}

func (UnitView) TableName() string {
	return VIEW_UNIT
}

type EventView struct {
	ID           string               `json:"id"`
	CompanyID    string               `json:"companyId"`
	PropertyID   string               `json:"propertyId"`
	UnitID       string               `json:"unitId"`
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	StartDt      time.Time            `json:"startDt"`
	EndDt        time.Time            `json:"endDt"`
	Status       constant.EventStatus `json:"status"`
	Price        int64                `json:"price"`
	CreateBy     string               `json:"createBy"`
	CreateDt     time.Time            `json:"createDt"`
	UpdateBy     string               `json:"updateBy"`
	UpdateDt     time.Time            `json:"updateDt"`
	DeleteDt     gorm.DeletedAt       `json:"deleteDt"`
	CompanyName  string               `json:"companyName"`
	PropertyName string               `json:"propertyName"`
	UnitName     string               `json:"unitName"`
	CreateName   string               `json:"createName"`
	UpdateName   string               `json:"updateName"`

	Company  *CompanyView  `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Property *PropertyView `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
	Unit     *UnitView     `json:"unit,omitempty"`
}

func (EventView) TableName() string {
	return VIEW_EVENT
}

type ProductView struct {
	ID          string         `json:"id"`
	CompanyID   string         `json:"companyId"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int64          `json:"price"`
	PhotoID     string         `json:"photoId"`
	PhotoUrl    string         `json:"photoUrl"`
	OpenTime    *time.Time     `json:"openTime"`
	CloseTime   *time.Time     `json:"closeTime"`
	CreateBy    string         `json:"createBy"`
	CreateDt    time.Time      `json:"createDt"`
	UpdateBy    string         `json:"updateBy"`
	UpdateDt    time.Time      `json:"updateDt"`
	DeleteDt    gorm.DeletedAt `json:"deleteDt"`
	CompanyName string         `json:"companyName"`
	CreateName  string         `json:"createName"`
	UpdateName  string         `json:"updateName"`

	Company *CompanyView `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
}

func (ProductView) TableName() string {
	return VIEW_PRODUCT
}

type TransactionView struct {
	ID           string         `json:"id"`
	CompanyID    string         `json:"companyId"`
	TotalEvent   int64          `json:"totalEvent"`
	TotalProduct int64          `json:"totalProduct"`
	Total        int64          `json:"total"`
	Paid         bool           `json:"paid"`
	CreateBy     string         `json:"createBy"`
	CreateDt     time.Time      `json:"createDt"`
	UpdateBy     string         `json:"updateBy"`
	UpdateDt     time.Time      `json:"updateDt"`
	DeleteDt     gorm.DeletedAt `json:"deleteDt"`
	CompanyName  string         `json:"companyName"`
	CreateName   string         `json:"createName"`
	UpdateName   string         `json:"updateName"`

	Company             *CompanyView             `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Transactionevents   []TransactioneventView   `json:"transactionevents,omitempty" gorm:"foreignKey:TransactionID"`
	Transactionproducts []TransactionproductView `json:"transactionproducts,omitempty" gorm:"foreignKey:TransactionID"`
}

func (TransactionView) TableName() string {
	return VIEW_TRANSACTION
}

type TransactioneventView struct {
	ID            string         `json:"id"`
	CompanyID     string         `json:"companyId"`
	TransactionID string         `json:"transactionId"`
	EventID       string         `json:"eventId"`
	Price         int64          `json:"price"`
	Paid          bool           `json:"paid"`
	CreateBy      string         `json:"createBy"`
	CreateDt      time.Time      `json:"createDt"`
	UpdateBy      string         `json:"updateBy"`
	UpdateDt      time.Time      `json:"updateDt"`
	DeleteDt      gorm.DeletedAt `json:"deleteDt"`
	CompanyName   string         `json:"companyName"`
	CreateName    string         `json:"createName"`
	UpdateName    string         `json:"updateName"`

	Company     *CompanyView     `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Transaction *TransactionView `json:"transaction,omitempty" gorm:"foreignKey:TransactionID"`
	Event       *EventView       `json:"event,omitempty" gorm:"foreignKey:EventID"`
}

func (TransactioneventView) TableName() string {
	return VIEW_TRANSACTIONEVENT
}

type TransactionproductView struct {
	ID            string         `json:"id"`
	CompanyID     string         `json:"companyId"`
	TransactionID string         `json:"transactionId"`
	EventID       string         `json:"eventId"`
	ProductID     string         `json:"productId"`
	ProductName   string         `json:"productName"`
	Price         int64          `json:"price"`
	Paid          bool           `json:"paid"`
	CreateBy      string         `json:"createBy"`
	CreateDt      time.Time      `json:"createDt"`
	UpdateBy      string         `json:"updateBy"`
	UpdateDt      time.Time      `json:"updateDt"`
	DeleteDt      gorm.DeletedAt `json:"deleteDt"`
	CompanyName   string         `json:"companyName"`
	CreateName    string         `json:"createName"`
	UpdateName    string         `json:"updateName"`

	Company     *CompanyView     `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Transaction *TransactionView `json:"transaction,omitempty" gorm:"foreignKey:TransactionID"`
	Event       *EventView       `json:"event,omitempty" gorm:"foreignKey:EventID"`
	Product     *ProductView     `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (TransactionproductView) TableName() string {
	return VIEW_TRANSACTIONPRODUCT
}

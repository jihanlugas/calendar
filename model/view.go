package model

import (
	"time"

	"github.com/jihanlugas/calendar/constant"
	"gorm.io/gorm"
)

type PaymentmethodView struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	CreateBy string         `json:"createBy"`
	CreateDt time.Time      `json:"createDt"`
	UpdateBy string         `json:"updateBy"`
	UpdateDt time.Time      `json:"updateDt"`
	DeleteDt gorm.DeletedAt `json:"deleteDt"`
}

func (PaymentmethodView) TableName() string {
	return VIEW_PAYMENTMETHOD
}

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

	Users                 []UserView                 `json:"users" gorm:"foreignKey:CompanyID"`
	Properties            []PropertyView             `json:"properties" gorm:"foreignKey:CompanyID"`
	Units                 []UnitView                 `json:"units" gorm:"foreignKey:CompanyID"`
	Products              []ProductView              `json:"products" gorm:"foreignKey:CompanyID"`
	Events                []EventView                `json:"events" gorm:"foreignKey:CompanyID"`
	Companypaymentmethods []CompanypaymentmethodView `json:"companypaymentmethods" gorm:"foreignKey:CompanyID"`
}

func (CompanyView) TableName() string {
	return VIEW_COMPANY
}

type CompanypaymentmethodView struct {
	ID                string         `json:"id"`
	CompanyID         string         `json:"companyId"`
	PaymentmethodID   string         `json:"paymentmethodId"`
	CreateBy          string         `json:"createBy"`
	CreateDt          time.Time      `json:"createDt"`
	UpdateBy          string         `json:"updateBy"`
	UpdateDt          time.Time      `json:"updateDt"`
	DeleteDt          gorm.DeletedAt `json:"deleteDt"`
	CompanyName       string         `json:"companyName"`
	PaymentmethodName string         `json:"paymentmethodName"`
	CreateName        string         `json:"createName"`
	UpdateName        string         `json:"updateName"`

	Company       *CompanyView       `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Paymentmethod *PaymentmethodView `json:"paymentmethod,omitempty" gorm:"foreignKey:PaymentmethodID"`
}

func (CompanypaymentmethodView) TableName() string {
	return VIEW_COMPANYPAYMENTMETHOD
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
	ID                string         `json:"id"`
	CompanyID         string         `json:"companyId"`
	PropertyID        string         `json:"propertyId"`
	Priority          int            `json:"priority"`
	Weekdays          Int32Array     `json:"weekdays" gorm:"type:int[]"`
	StartTime         *time.Time     `json:"startTime"`
	EndTime           *time.Time     `json:"endTime"`
	Price             int64          `json:"price"`
	CreateBy          string         `json:"createBy"`
	CreateDt          time.Time      `json:"createDt"`
	UpdateBy          string         `json:"updateBy"`
	UpdateDt          time.Time      `json:"updateDt"`
	DeleteDt          gorm.DeletedAt `json:"deleteDt"`
	CompanyName       string         `json:"companyName"`
	PropertyName      string         `json:"propertyName"`
	CreateName        string         `json:"createName"`
	UpdateName        string         `json:"updateName"`
	StartTimeFormated string         `json:"startTimeFormatted"`
	EndTimeFormated   string         `json:"endTimeFormatted"`

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
	OrdereventID string               `json:"ordereventId"`
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	StartDt      time.Time            `json:"startDt"`
	EndDt        time.Time            `json:"endDt"`
	Status       constant.EventStatus `json:"status"`
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

	Company    *CompanyView    `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Property   *PropertyView   `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
	Unit       *UnitView       `json:"unit,omitempty"`
	Orderevent *OrdereventView `json:"orderevent,omitempty" gorm:"foreignKey:OrdereventID"`
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

type TaxView struct {
	ID          string           `json:"id"`
	CompanyID   string           `json:"companyId"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Type        constant.TaxType `json:"type"`
	Value       int64            `json:"value"`
	CreateBy    string           `json:"createBy"`
	CreateDt    time.Time        `json:"createDt"`
	UpdateBy    string           `json:"updateBy"`
	UpdateDt    time.Time        `json:"updateDt"`
	DeleteDt    gorm.DeletedAt   `json:"deleteDt"`
	CompanyName string           `json:"companyName"`
	CreateName  string           `json:"createName"`
	UpdateName  string           `json:"updateName"`

	Company *CompanyView `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
}

func (TaxView) TableName() string {
	return VIEW_TAX
}

type DiscountView struct {
	ID          string                `json:"id"`
	CompanyID   string                `json:"companyId"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Type        constant.DiscountType `json:"type"`
	Value       int64                 `json:"value"`
	CreateBy    string                `json:"createBy"`
	CreateDt    time.Time             `json:"createDt"`
	UpdateBy    string                `json:"updateBy"`
	UpdateDt    time.Time             `json:"updateDt"`
	DeleteDt    gorm.DeletedAt        `json:"deleteDt"`
	CompanyName string                `json:"companyName"`
	CreateName  string                `json:"createName"`
	UpdateName  string                `json:"updateName"`

	Company *CompanyView `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
}

func (DiscountView) TableName() string {
	return VIEW_DISCOUNT
}

type OrderView struct {
	ID          string         `json:"id"`
	CompanyID   string         `json:"companyId"`
	Tax         int64          `json:"tax"`
	Discount    int64          `json:"discount"`
	Rounding    int64          `json:"rounding"`
	Subtotal    int64          `json:"subtotal"`
	Total       int64          `json:"total"`
	Payment     int64          `json:"payment"`
	CreateBy    string         `json:"createBy"`
	CreateDt    time.Time      `json:"createDt"`
	UpdateBy    string         `json:"updateBy"`
	UpdateDt    time.Time      `json:"updateDt"`
	DeleteDt    gorm.DeletedAt `json:"deleteDt"`
	CompanyName string         `json:"companyName"`
	CreateName  string         `json:"createName"`
	UpdateName  string         `json:"updateName"`

	Company        *CompanyView        `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Orderevents    []OrdereventView    `json:"orderevents,omitempty" gorm:"foreignKey:OrderID"`
	Orderproducts  []OrderproductView  `json:"orderproducts,omitempty" gorm:"foreignKey:OrderID"`
	Ordertaxes     []OrdertaxView      `json:"ordertaxes,omitempty" gorm:"foreignKey:OrderID"`
	Orderdiscounts []OrderdiscountView `json:"orderdiscounts,omitempty" gorm:"foreignKey:OrderID"`
	Orderpayments  []OrderpaymentView  `json:"orderpayments,omitempty" gorm:"foreignKey:OrderID"`
}

func (OrderView) TableName() string {
	return VIEW_ORDER
}

type OrdereventView struct {
	ID         string         `json:"id"`
	CompanyID  string         `json:"companyId"`
	OrderID    string         `json:"orderId"`
	EventID    string         `json:"eventId"`
	CreateBy   string         `json:"createBy"`
	CreateDt   time.Time      `json:"createDt"`
	UpdateBy   string         `json:"updateBy"`
	UpdateDt   time.Time      `json:"updateDt"`
	DeleteDt   gorm.DeletedAt `json:"deleteDt"`
	OrderName  string         `json:"orderName"`
	EventName  string         `json:"eventName"`
	CreateName string         `json:"createName"`
	UpdateName string         `json:"updateName"`

	Company *CompanyView `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Order   *OrderView   `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Event   *EventView   `json:"event,omitempty" gorm:"foreignKey:EventID"`
}

func (OrdereventView) TableName() string {
	return VIEW_ORDEREVENT
}

type OrderproductView struct {
	ID          string         `json:"id"`
	CompanyID   string         `json:"companyId"`
	OrderID     string         `json:"orderId"`
	ProductID   string         `json:"productId"`
	Quantity    int            `json:"quantity"`
	Price       int64          `json:"price"`
	CreateBy    string         `json:"createBy"`
	CreateDt    time.Time      `json:"createDt"`
	UpdateBy    string         `json:"updateBy"`
	UpdateDt    time.Time      `json:"updateDt"`
	DeleteDt    gorm.DeletedAt `json:"deleteDt"`
	OrderName   string         `json:"orderName"`
	ProductName string         `json:"productName"`
	CreateName  string         `json:"createName"`
	UpdateName  string         `json:"updateName"`

	Company *CompanyView `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Order   *OrderView   `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product *ProductView `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (OrderproductView) TableName() string {
	return VIEW_ORDERPRODUCT
}

type OrdertaxView struct {
	ID         string         `json:"id"`
	CompanyID  string         `json:"companyId"`
	OrderID    string         `json:"orderId"`
	TaxID      string         `json:"taxId"`
	Total      int64          `json:"total"`
	CreateBy   string         `json:"createBy"`
	CreateDt   time.Time      `json:"createDt"`
	UpdateBy   string         `json:"updateBy"`
	UpdateDt   time.Time      `json:"updateDt"`
	DeleteDt   gorm.DeletedAt `json:"deleteDt"`
	OrderName  string         `json:"orderName"`
	TaxName    string         `json:"taxName"`
	CreateName string         `json:"createName"`
	UpdateName string         `json:"updateName"`

	Company *CompanyView `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Order   *OrderView   `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Tax     *TaxView     `json:"tax,omitempty" gorm:"foreignKey:TaxID"`
}

func (OrdertaxView) TableName() string {
	return VIEW_ORDERTAX
}

type OrderdiscountView struct {
	ID           string         `json:"id"`
	CompanyID    string         `json:"companyId"`
	OrderID      string         `json:"orderId"`
	DiscountID   string         `json:"discountId"`
	Total        int64          `json:"total"`
	CreateBy     string         `json:"createBy"`
	CreateDt     time.Time      `json:"createDt"`
	UpdateBy     string         `json:"updateBy"`
	UpdateDt     time.Time      `json:"updateDt"`
	DeleteDt     gorm.DeletedAt `json:"deleteDt"`
	OrderName    string         `json:"orderName"`
	DiscountName string         `json:"discountName"`
	CreateName   string         `json:"createName"`
	UpdateName   string         `json:"updateName"`

	Company  *CompanyView  `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Order    *OrderView    `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Discount *DiscountView `json:"discount,omitempty" gorm:"foreignKey:DiscountID"`
}

func (OrderdiscountView) TableName() string {
	return VIEW_ORDERDISCOUNT
}

type OrderpaymentView struct {
	ID                string         `json:"id"`
	CompanyID         string         `json:"companyId"`
	OrderID           string         `json:"orderId"`
	PaymentmethodID   string         `json:"paymentmethodId"`
	Total             int64          `json:"total"`
	CreateBy          string         `json:"createBy"`
	CreateDt          time.Time      `json:"createDt"`
	UpdateBy          string         `json:"updateBy"`
	UpdateDt          time.Time      `json:"updateDt"`
	DeleteDt          gorm.DeletedAt `json:"deleteDt"`
	OrderName         string         `json:"orderName"`
	PaymentmethodName string         `json:"paymentmethodName"`
	CreateName        string         `json:"createName"`
	UpdateName        string         `json:"updateName"`

	Company       *CompanyView       `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Order         *OrderView         `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Paymentmethod *PaymentmethodView `json:"paymentmethod,omitempty" gorm:"foreignKey:PaymentmethodID"`
}

func (OrderpaymentView) TableName() string {
	return VIEW_ORDERPAYMENT
}

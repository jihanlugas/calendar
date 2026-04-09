package model

import "time"

const (
	VIEW_PAYMENTMETHOD        = "paymentmethods_view"
	VIEW_USER                 = "users_view"
	VIEW_COMPANY              = "companies_view"
	VIEW_COMPANYPAYMENTMETHOD = "companypaymentmethods_view"
	VIEW_USERCOMPANY          = "usercompanies_view"
	VIEW_PROPERTY             = "properties_view"
	VIEW_PROPERTYPRICE        = "propertyprices_view"
	VIEW_PROPERTYTIMELINE     = "propertytimelines_view"
	VIEW_UNIT                 = "units_view"
	VIEW_EVENT                = "events_view"
	VIEW_PRODUCT              = "products_view"
	VIEW_TAX                  = "taxes_view"
	VIEW_DISCOUNT             = "discounts_view"
	VIEW_ORDER                = "orders_view"
	VIEW_ORDEREVENT           = "orderevents_view"
	VIEW_ORDERPRODUCT         = "orderproducts_view"
	VIEW_ORDERTAX             = "ordertaxes_view"
	VIEW_ORDERDISCOUNT        = "orderdiscounts_view"
	VIEW_ORDERPAYMENT         = "orderpayments_view"
)

type UserLogin struct {
	ExpiredDt     time.Time `json:"expiredDt"`
	UserID        string    `json:"userId"`
	PassVersion   int       `json:"passVersion"`
	CompanyID     string    `json:"companyId"`
	Role          string    `json:"role"`
	UsercompanyID string    `json:"usercompanyId"`
}

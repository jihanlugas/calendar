package model

import "time"

const (
	VIEW_USER               = "users_view"
	VIEW_COMPANY            = "companies_view"
	VIEW_USERCOMPANY        = "usercompanies_view"
	VIEW_PROPERTY           = "properties_view"
	VIEW_PROPERTYTIMELINE   = "propertytimelines_view"
	VIEW_PROPERTYGROUP      = "propertygroups_view"
	VIEW_EVENT              = "events_view"
	VIEW_PRODUCT            = "products_view"
	VIEW_TRANSACTION        = "transactions_view"
	VIEW_TRANSACTIONEVENT   = "transactionevents_view"
	VIEW_TRANSACTIONPRODUCT = "transactionproducts_view"
)

type UserLogin struct {
	ExpiredDt     time.Time `json:"expiredDt"`
	UserID        string    `json:"userId"`
	PassVersion   int       `json:"passVersion"`
	CompanyID     string    `json:"companyId"`
	Role          string    `json:"role"`
	UsercompanyID string    `json:"usercompanyId"`
}

var PreloadCompany = []string{
	"Users",
	"Properties",
	"Propertygroups",
	"Products",
	"Events",
}

var PreloadUser = []string{
	"Company",
	"Company.Properties",
	"Company.Properties.Propertygroups",
	"Company.Properties.Propertytimeline",
	"Usercompanies",
	"Usercompanies.Company",
}

var PreloadUsercompany = []string{
	"Company",
	"User",
}

var PreloadProperty = []string{
	"Company",
	"Propertytimeline",
	"Propertygroups",
	"Events",
}

var PreloadPropertygroup = []string{
	"Company",
	"Property",
	"Events",
}

var PreloadEvent = []string{
	"Company",
	"Property",
	"Propertygroup",
}

var PreloadProduct = []string{
	"Company",
}

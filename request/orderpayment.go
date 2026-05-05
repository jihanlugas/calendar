package request

type CreateOrderpayment struct {
	CompanyID              string `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	OrderID                string `json:"orderId" form:"orderId" query:"orderId" validate:"required"`
	CompanypaymentmethodID string `json:"companypaymentmethodId" form:"companypaymentmethodId" query:"companypaymentmethodId" validate:"required"`
	Name                   string `json:"name" form:"name" query:"name" validate:"required"`
	Total                  int64  `json:"total" form:"total" query:"total" validate:"required"`
}

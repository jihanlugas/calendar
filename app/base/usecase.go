package base

import (
	"errors"

	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/response"
	"gorm.io/gorm"
)

type Usecase interface {
	GetConnection() *gorm.DB
	RequireCompanyIDAllowed(loginUser jwt.UserLogin, companyID string) (err error)
}

type usecase struct{}

func NewUsecase() Usecase {
	return &usecase{}
}

func (u *usecase) GetConnection() *gorm.DB {
	return db.GetGlobalConnection()
}

func (u *usecase) RequireCompanyIDAllowed(loginUser jwt.UserLogin, companyID string) (err error) {
	if loginUser.Role != constant.RoleAdmin {
		if loginUser.CompanyID != companyID {
			return errors.New(response.ErrorHandlerIDOR)
		}
	}

	return err
}

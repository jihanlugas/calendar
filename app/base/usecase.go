package base

import (
	"errors"

	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/response"
	"gorm.io/gorm"
)

type Usecase interface {
	WithConn() (*gorm.DB, func())
	RequireCompanyIDAllowed(loginUser jwt.UserLogin, companyID string) error
}

type usecase struct{}

func NewUsecase() Usecase {
	return &usecase{}
}

func (u *usecase) WithConn() (*gorm.DB, func()) {
	return db.GetConnection()
}

func (u *usecase) RequireCompanyIDAllowed(loginUser jwt.UserLogin, companyID string) error {
	if jwt.IsSaveCompanyIDOR(loginUser, companyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}
	return nil
}

package orderpayment

import (
	"fmt"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/app/companypaymentmethod"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
)

type Usecase interface {
	Create(loginUser jwt.UserLogin, req request.CreateOrderpayment) (err error)
}

type usecase struct {
	baseUsecase                    base.Usecase
	repository                     Repository
	companypaymentmethodRepository companypaymentmethod.Repository
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateOrderpayment) (err error) {
	conn := u.baseUsecase.GetConnection()

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID)
	if err != nil {
		return err
	}

	tCompanypaymentmethod, err := u.companypaymentmethodRepository.GetTableById(conn, req.CompanypaymentmethodID)
	if err != nil {
		return err
	}

	tx := conn.Begin()

	tOrderpayment := model.Orderpayment{
		CompanyID:              req.CompanyID,
		OrderID:                req.OrderID,
		CompanypaymentmethodID: req.CompanypaymentmethodID,
		PaymentmethodID:        tCompanypaymentmethod.PaymentmethodID,
		Name:                   req.Name,
		Total:                  req.Total,
		CreateBy:               loginUser.UserID,
	}

	err = u.repository.Create(tx, tOrderpayment)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return err
}

func NewUsecase(baseUsecase base.Usecase, repository Repository, companypaymentmethodRepository companypaymentmethod.Repository) Usecase {
	return &usecase{
		baseUsecase:                    baseUsecase,
		repository:                     repository,
		companypaymentmethodRepository: companypaymentmethodRepository,
	}
}

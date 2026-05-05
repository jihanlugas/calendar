package orderpayment

import (
	"errors"
	"fmt"

	"github.com/jihanlugas/calendar/app/companypaymentmethod"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
)

type Usecase interface {
	Create(loginUser jwt.UserLogin, req request.CreateOrderpayment) error
}

type usecase struct {
	repository                     Repository
	companypaymentmethodRepository companypaymentmethod.Repository
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateOrderpayment) error {
	var err error
	var tOrderpayment model.Orderpayment
	var tCompanypaymentmethod model.Companypaymentmethod

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tCompanypaymentmethod, err = u.companypaymentmethodRepository.GetTableById(conn, req.CompanypaymentmethodID)
	if err != nil {
		return err
	}

	tx := conn.Begin()

	tOrderpayment = model.Orderpayment{
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
		return err
	}

	return err
}

func NewUsecase(repository Repository, companypaymentmethodRepository companypaymentmethod.Repository) Usecase {
	return &usecase{
		repository:                     repository,
		companypaymentmethodRepository: companypaymentmethodRepository,
	}
}

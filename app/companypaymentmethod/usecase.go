package companypaymentmethod

import (
	"errors"
	"fmt"

	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
)

type Usecase interface {
	Page(loginUser jwt.UserLogin, req request.PageCompanypaymentmethod) (vCompanypaymentmethods []model.CompanypaymentmethodView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vCompanypaymentmethod model.CompanypaymentmethodView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateCompanypaymentmethod) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdateCompanypaymentmethod) error
	Delete(loginUser jwt.UserLogin, id string) error
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageCompanypaymentmethod) (vCompanypaymentmethods []model.CompanypaymentmethodView, count int64, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return vCompanypaymentmethods, count, errors.New(response.ErrorHandlerIDOR)
	}

	vCompanypaymentmethods, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vCompanypaymentmethods, count, err
	}

	return vCompanypaymentmethods, count, err
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vCompanypaymentmethod model.CompanypaymentmethodView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	vCompanypaymentmethod, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vCompanypaymentmethod, err
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vCompanypaymentmethod.CompanyID) {
		return vCompanypaymentmethod, errors.New(response.ErrorHandlerIDOR)
	}

	return vCompanypaymentmethod, err
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateCompanypaymentmethod) error {
	var err error
	var tCompanypaymentmethod model.Companypaymentmethod

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tCompanypaymentmethod = model.Companypaymentmethod{
		CompanyID:       req.CompanyID,
		PaymentmethodID: req.PaymentmethodID,
		CreateBy:        loginUser.UserID,
		UpdateBy:        loginUser.UserID,
	}

	err = u.repository.Create(tx, tCompanypaymentmethod)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateCompanypaymentmethod) error {
	var err error
	var tCompanypaymentmethod model.Companypaymentmethod

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tCompanypaymentmethod, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tCompanypaymentmethod.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tCompanypaymentmethod.PaymentmethodID = req.PaymentmethodID
	tCompanypaymentmethod.UpdateBy = loginUser.UserID
	err = u.repository.Save(tx, tCompanypaymentmethod)
	if err != nil {
		return fmt.Errorf("failed to update %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Delete(loginUser jwt.UserLogin, id string) error {
	var err error
	var tCompanypaymentmethod model.Companypaymentmethod

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tCompanypaymentmethod, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tCompanypaymentmethod.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	err = u.repository.Delete(tx, tCompanypaymentmethod)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

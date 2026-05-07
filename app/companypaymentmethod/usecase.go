package companypaymentmethod

import (
	"fmt"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
)

type Usecase interface {
	Page(loginUser jwt.UserLogin, req request.PageCompanypaymentmethod) (vCompanypaymentmethods []model.CompanypaymentmethodView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vCompanypaymentmethod model.CompanypaymentmethodView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateCompanypaymentmethod) (err error)
	Update(loginUser jwt.UserLogin, id string, req request.UpdateCompanypaymentmethod) (err error)
	Delete(loginUser jwt.UserLogin, id string) (err error)
}

type usecase struct {
	baseUsecase base.Usecase
	repository  Repository
}

func NewUsecase(baseUsecase base.Usecase, repository Repository) Usecase {
	return &usecase{
		baseUsecase: baseUsecase,
		repository:  repository,
	}
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageCompanypaymentmethod) (vCompanypaymentmethods []model.CompanypaymentmethodView, count int64, err error) {
	conn, closeConn := u.baseUsecase.WithConn()
	defer closeConn()

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID)
	if err != nil {
		return vCompanypaymentmethods, count, err
	}

	vCompanypaymentmethods, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vCompanypaymentmethods, count, err
	}

	return vCompanypaymentmethods, count, nil
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vCompanypaymentmethod model.CompanypaymentmethodView, err error) {
	conn, closeConn := u.baseUsecase.WithConn()
	defer closeConn()

	vCompanypaymentmethod, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vCompanypaymentmethod, err
	}

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, vCompanypaymentmethod.CompanyID)
	if err != nil {
		return vCompanypaymentmethod, err
	}

	return vCompanypaymentmethod, nil
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateCompanypaymentmethod) (err error) {

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID)
	if err != nil {
		return err
	}

	conn, closeConn := u.baseUsecase.WithConn()
	defer closeConn()

	tCompanypaymentmethod := model.Companypaymentmethod{
		CompanyID:       req.CompanyID,
		PaymentmethodID: req.PaymentmethodID,
		CreateBy:        loginUser.UserID,
		UpdateBy:        loginUser.UserID,
	}

	tx := conn.Begin()

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

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateCompanypaymentmethod) (err error) {
	conn, closeConn := u.baseUsecase.WithConn()
	defer closeConn()

	tCompanypaymentmethod, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, tCompanypaymentmethod.CompanyID)
	if err != nil {
		return err
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

func (u usecase) Delete(loginUser jwt.UserLogin, id string) (err error) {
	conn, closeConn := u.baseUsecase.WithConn()
	defer closeConn()

	tCompanypaymentmethod, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, tCompanypaymentmethod.CompanyID)
	if err != nil {
		return err
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

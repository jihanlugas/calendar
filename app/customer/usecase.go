package customer

import (
	"fmt"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/utils"
)

type Usecase interface {
	Page(loginUser jwt.UserLogin, req request.PageCustomer) (vCustomers []model.CustomerView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vCustomer model.CustomerView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateCustomer) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdateCustomer) error
	Delete(loginUser jwt.UserLogin, id string) error
}

type usecase struct {
	baseUsecase base.Usecase
	repository  Repository
}

func (u *usecase) Page(loginUser jwt.UserLogin, req request.PageCustomer) (vCustomers []model.CustomerView, count int64, err error) {
	conn := u.baseUsecase.GetConnection()

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID)
	if err != nil {
		return vCustomers, count, err
	}

	vCustomers, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vCustomers, count, err
	}

	return vCustomers, count, err
}

func (u *usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vCustomer model.CustomerView, err error) {
	conn := u.baseUsecase.GetConnection()

	vCustomer, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vCustomer, fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, vCustomer.CompanyID)
	if err != nil {
		return vCustomer, err
	}

	return vCustomer, err
}

func (u *usecase) Create(loginUser jwt.UserLogin, req request.CreateCustomer) (err error) {
	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID)
	if err != nil {
		return err
	}

	conn := u.baseUsecase.GetConnection()

	tx := conn.Begin()

	tCustomer := model.Customer{
		ID:          utils.GetUniqueID(),
		CompanyID:   req.CompanyID,
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: utils.FormatPhoneTo62(req.PhoneNumber),
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	err = u.repository.Create(tx, tCustomer)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u *usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateCustomer) (err error) {
	conn := u.baseUsecase.GetConnection()

	tCustomer, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, tCustomer.CompanyID)
	if err != nil {
		return err
	}

	tx := conn.Begin()

	tCustomer.Name = req.Name
	tCustomer.Email = req.Email
	tCustomer.PhoneNumber = utils.FormatPhoneTo62(req.PhoneNumber)
	tCustomer.UpdateBy = loginUser.UserID
	err = u.repository.Save(tx, tCustomer)
	if err != nil {
		return fmt.Errorf("failed to update %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u *usecase) Delete(loginUser jwt.UserLogin, id string) (err error) {
	conn := u.baseUsecase.GetConnection()

	tCustomer, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, tCustomer.CompanyID)
	if err != nil {
		return err
	}

	tx := conn.Begin()

	err = u.repository.Delete(tx, tCustomer)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func NewUsecase(baseUsecase base.Usecase, repository Repository) Usecase {
	return &usecase{
		baseUsecase: baseUsecase,
		repository:  repository,
	}
}

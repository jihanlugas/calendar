package product

import (
	"errors"
	"fmt"

	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
	"github.com/jihanlugas/calendar/utils"
)

type Usecase interface {
	Page(loginUser jwt.UserLogin, req request.PageProduct) (vProducts []model.ProductView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProduct model.ProductView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateProduct) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdateProduct) error
	Delete(loginUser jwt.UserLogin, id string) error
}

type usecase struct {
	repository Repository
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageProduct) (vProducts []model.ProductView, count int64, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return vProducts, count, errors.New(response.ErrorHandlerIDOR)
	}

	vProducts, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vProducts, count, err
	}

	return vProducts, count, err
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProduct model.ProductView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	vProduct, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vProduct, fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vProduct.CompanyID) {
		return vProduct, errors.New(response.ErrorHandlerIDOR)
	}

	return vProduct, err
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateProduct) error {
	var err error
	var tProduct model.Product

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tProduct = model.Product{
		ID:          utils.GetUniqueID(),
		CompanyID:   req.CompanyID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	err = u.repository.Create(tx, tProduct)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateProduct) error {
	var err error
	var tProduct model.Product

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tProduct, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tProduct.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tProduct.Name = req.Name
	tProduct.Description = req.Description
	tProduct.Price = req.Price
	tProduct.UpdateBy = loginUser.UserID
	err = u.repository.Save(tx, tProduct)
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
	var tProduct model.Product

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tProduct, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tProduct.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	err = u.repository.Delete(tx, tProduct)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func NewUsecase(repository Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

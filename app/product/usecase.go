package product

import (
	"fmt"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/utils"
)

type Usecase interface {
	Page(loginUser jwt.UserLogin, req request.PageProduct) (vProducts []model.ProductView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProduct model.ProductView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateProduct) (err error)
	Update(loginUser jwt.UserLogin, id string, req request.UpdateProduct) (err error)
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

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageProduct) (vProducts []model.ProductView, count int64, err error) {
	conn := u.baseUsecase.GetConnection()

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID)
	if err != nil {
		return vProducts, count, err
	}

	vProducts, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vProducts, count, err
	}

	return vProducts, count, nil
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProduct model.ProductView, err error) {
	conn := u.baseUsecase.GetConnection()

	vProduct, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vProduct, fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, vProduct.CompanyID); err != nil {
		return vProduct, err
	}

	return vProduct, nil
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateProduct) (err error) {
	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID)
	if err != nil {
		return err
	}

	conn := u.baseUsecase.GetConnection()

	tx := conn.Begin()

	tProduct := model.Product{
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

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateProduct) (err error) {
	conn := u.baseUsecase.GetConnection()

	tProduct, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, tProduct.CompanyID); err != nil {
		return err
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

func (u usecase) Delete(loginUser jwt.UserLogin, id string) (err error) {
	conn := u.baseUsecase.GetConnection()

	tProduct, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, tProduct.CompanyID); err != nil {
		return err
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

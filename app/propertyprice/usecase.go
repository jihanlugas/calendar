package propertyprice

import (
	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
)

type Usecase interface {
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProperty model.PropertypriceView, err error)
	Create(loginUser jwt.UserLogin, req request.CreatePropertyprice) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdatePropertyprice) error
	Delete(loginUser jwt.UserLogin, id string) error
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

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vPropertyprice model.PropertypriceView, err error) {
	conn := u.baseUsecase.GetConnection()

	vPropertyprice, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vPropertyprice, err
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, vPropertyprice.CompanyID); err != nil {
		return vPropertyprice, err
	}

	return vPropertyprice, nil
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreatePropertyprice) error {
	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID); err != nil {
		return err
	}

	conn := u.baseUsecase.GetConnection()

	countPropertyprices, err := u.repository.CountByPropertyID(conn, req.PropertyID)
	if err != nil {
		return err
	}

	tPropertyprice := model.Propertyprice{
		CompanyID:  req.CompanyID,
		PropertyID: req.PropertyID,
		Price:      req.Price,
		Weekdays:   req.Weekdays,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Priority:   int(countPropertyprices + 1),
		CreateBy:   loginUser.UserID,
		UpdateBy:   loginUser.UserID,
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := u.repository.Create(tx, tPropertyprice); err != nil {
		_ = tx.Rollback().Error
		return err
	}

	return tx.Commit().Error
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdatePropertyprice) error {
	conn := u.baseUsecase.GetConnection()

	tPropertyprice, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return err
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, tPropertyprice.CompanyID); err != nil {
		return err
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	tPropertyprice.Price = req.Price
	tPropertyprice.Weekdays = req.Weekdays
	tPropertyprice.StartTime = req.StartTime
	tPropertyprice.EndTime = req.EndTime
	if err := u.repository.Save(tx, tPropertyprice); err != nil {
		_ = tx.Rollback().Error
		return err
	}

	return tx.Commit().Error
}

func (u usecase) Delete(loginUser jwt.UserLogin, id string) error {
	conn := u.baseUsecase.GetConnection()

	tPropertyprice, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return err
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, tPropertyprice.CompanyID); err != nil {
		return err
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := u.repository.Delete(tx, tPropertyprice); err != nil {
		_ = tx.Rollback().Error
		return err
	}

	return tx.Commit().Error
}

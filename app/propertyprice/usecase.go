package propertyprice

import (
	"errors"

	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
)

type Usecase interface {
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProperty model.PropertypriceView, err error)
	Create(loginUser jwt.UserLogin, req request.CreatePropertyprice) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdatePropertyprice) error
	Delete(loginUser jwt.UserLogin, id string) error
}

type usecase struct {
	repository Repository
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vPropertyprice model.PropertypriceView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	vPropertyprice, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vPropertyprice, err
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vPropertyprice.CompanyID) {
		return vPropertyprice, errors.New(response.ErrorHandlerIDOR)
	}

	return vPropertyprice, err
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreatePropertyprice) error {
	var err error
	var tPropertyprice model.Propertyprice

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	countPropertyprices, err := u.repository.CountByPropertyID(tx, req.PropertyID)
	if err != nil {
		return err
	}

	tPropertyprice = model.Propertyprice{
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

	err = u.repository.Create(tx, tPropertyprice)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdatePropertyprice) error {
	var err error
	var tPropertyprice model.Propertyprice

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	tPropertyprice, err = u.repository.GetTableById(tx, id)
	if err != nil {
		return err
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tPropertyprice.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tPropertyprice.Price = req.Price
	tPropertyprice.Weekdays = req.Weekdays
	tPropertyprice.StartTime = req.StartTime
	tPropertyprice.EndTime = req.EndTime

	err = u.repository.Update(tx, tPropertyprice)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (u usecase) Delete(loginUser jwt.UserLogin, id string) error {
	var err error
	var tPropertyprice model.Propertyprice

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	tPropertyprice, err = u.repository.GetTableById(tx, id)
	if err != nil {
		return err
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tPropertyprice.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	err = u.repository.Delete(tx, tPropertyprice)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func NewUsecase(repository Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

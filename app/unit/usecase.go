package unit

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
	Page(loginUser jwt.UserLogin, req request.PageUnit) (vUnits []model.UnitView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vUnit model.UnitView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateUnit) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdateUnit) error
	Delete(loginUser jwt.UserLogin, id string) error
}

type usecase struct {
	repository Repository
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageUnit) (vUnits []model.UnitView, count int64, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return vUnits, count, errors.New(response.ErrorHandlerIDOR)
	}

	vUnits, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vUnits, count, err
	}

	return vUnits, count, err
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vUnit model.UnitView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	vUnit, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vUnit, errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vUnit.CompanyID) {
		return vUnit, errors.New(response.ErrorHandlerIDOR)
	}

	return vUnit, err
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateUnit) error {
	var err error
	var tUnit model.Unit

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tUnit = model.Unit{
		ID:          utils.GetUniqueID(),
		CompanyID:   req.CompanyID,
		PropertyID:  req.PropertyID,
		Name:        req.Name,
		Description: req.Description,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	err = u.repository.Create(tx, tUnit)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create %s: %v", u.repository.Name(), err))
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateUnit) error {
	var err error
	var tUnit model.Unit

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tUnit, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tUnit.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tUnit.Name = req.Name
	tUnit.Description = req.Description
	tUnit.UpdateBy = loginUser.UserID
	err = u.repository.Save(tx, tUnit)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to update %s: %v", u.repository.Name(), err))
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Delete(loginUser jwt.UserLogin, id string) error {
	var err error
	var tUnit model.Unit

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tUnit, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tUnit.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	err = u.repository.Delete(tx, tUnit)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to delete %s: %v", u.repository.Name(), err))
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

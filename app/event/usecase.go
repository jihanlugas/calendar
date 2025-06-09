package event

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
	Timeline(loginUser jwt.UserLogin, req request.TimelineEvent) (vEvents []model.EventView, err error)
	Page(loginUser jwt.UserLogin, req request.PageEvent) (vEvents []model.EventView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vEvent model.EventView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateEvent) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdateEvent) error
	Delete(loginUser jwt.UserLogin, id string) error
}

type usecase struct {
	repository Repository
}

func (u usecase) Timeline(loginUser jwt.UserLogin, req request.TimelineEvent) (vEvents []model.EventView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return vEvents, errors.New(response.ErrorHandlerIDOR)
	}

	vEvents, err = u.repository.Timeline(conn, req)
	if err != nil {
		return vEvents, err
	}

	return vEvents, err
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageEvent) (vEvents []model.EventView, count int64, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return vEvents, count, errors.New(response.ErrorHandlerIDOR)
	}

	vEvents, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vEvents, count, err
	}

	return vEvents, count, err
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vEvent model.EventView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	vEvent, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vEvent, errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vEvent.CompanyID) {
		return vEvent, errors.New(response.ErrorHandlerIDOR)
	}

	return vEvent, err
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateEvent) error {
	var err error
	var tEvent model.Event

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tEvent = model.Event{
		ID:              utils.GetUniqueID(),
		CompanyID:       req.CompanyID,
		PropertyID:      req.PropertyID,
		PropertygroupID: req.PropertygroupID,
		Name:            req.Name,
		Description:     req.Description,
		StartDt:         req.StartDt,
		EndDt:           req.EndDt,
		CreateBy:        loginUser.UserID,
		UpdateBy:        loginUser.UserID,
	}

	err = u.repository.Create(tx, tEvent)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create %s: %v", u.repository.Name(), err))
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateEvent) error {
	var err error
	var tEvent model.Event

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tEvent, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tEvent.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tEvent.PropertygroupID = req.PropertygroupID
	tEvent.Name = req.Name
	tEvent.Description = req.Description
	tEvent.StartDt = req.StartDt
	tEvent.EndDt = req.EndDt
	tEvent.UpdateBy = loginUser.UserID
	err = u.repository.Save(tx, tEvent)
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
	var tEvent model.Event

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tEvent, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tEvent.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	err = u.repository.Delete(tx, tEvent)
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

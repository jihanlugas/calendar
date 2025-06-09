package propertygroup

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
	Page(loginUser jwt.UserLogin, req request.PagePropertygroup) (vPropertygroups []model.PropertygroupView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vPropertygroup model.PropertygroupView, err error)
	Create(loginUser jwt.UserLogin, req request.CreatePropertygroup) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdatePropertygroup) error
	Delete(loginUser jwt.UserLogin, id string) error
}

type usecase struct {
	repository Repository
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PagePropertygroup) (vPropertygroups []model.PropertygroupView, count int64, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return vPropertygroups, count, errors.New(response.ErrorHandlerIDOR)
	}

	vPropertygroups, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vPropertygroups, count, err
	}

	return vPropertygroups, count, err
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vPropertygroup model.PropertygroupView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	vPropertygroup, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vPropertygroup, errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vPropertygroup.CompanyID) {
		return vPropertygroup, errors.New(response.ErrorHandlerIDOR)
	}

	return vPropertygroup, err
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreatePropertygroup) error {
	var err error
	var tPropertygroup model.Propertygroup

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tPropertygroup = model.Propertygroup{
		ID:          utils.GetUniqueID(),
		CompanyID:   req.CompanyID,
		PropertyID:  req.PropertyID,
		Name:        req.Name,
		Description: req.Description,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	err = u.repository.Create(tx, tPropertygroup)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create %s: %v", u.repository.Name(), err))
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdatePropertygroup) error {
	var err error
	var tPropertygroup model.Propertygroup

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tPropertygroup, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tPropertygroup.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tPropertygroup.Name = req.Name
	tPropertygroup.Description = req.Description
	tPropertygroup.UpdateBy = loginUser.UserID
	err = u.repository.Save(tx, tPropertygroup)
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
	var tPropertygroup model.Propertygroup

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tPropertygroup, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tPropertygroup.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	err = u.repository.Delete(tx, tPropertygroup)
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

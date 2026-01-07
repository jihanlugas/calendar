package property

import (
	"errors"
	"fmt"

	"github.com/jihanlugas/calendar/app/propertyprice"
	"github.com/jihanlugas/calendar/app/propertytimeline"
	"github.com/jihanlugas/calendar/app/unit"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
	"github.com/jihanlugas/calendar/utils"
)

type Usecase interface {
	Page(loginUser jwt.UserLogin, req request.PageProperty) (vProperties []model.PropertyView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProperty model.PropertyView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateProperty) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdateProperty) error
	Delete(loginUser jwt.UserLogin, id string) error
	GetPrice(req request.GetPrice) (price int64, err error)
}

type usecase struct {
	repository                 Repository
	repositoryPropertytimeline propertytimeline.Repository
	repositoryUnit             unit.Repository
	repositoryPropertyprice    propertyprice.Repository
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageProperty) (vProperties []model.PropertyView, count int64, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return vProperties, count, errors.New(response.ErrorHandlerIDOR)
	}

	vProperties, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vProperties, count, err
	}

	return vProperties, count, err
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProperty model.PropertyView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	vProperty, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vProperty, fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vProperty.CompanyID) {
		return vProperty, errors.New(response.ErrorHandlerIDOR)
	}

	return vProperty, err
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateProperty) error {
	var err error
	var tProperty model.Property
	var tPropertytimeline model.Propertytimeline
	var tUnits []model.Unit
	var tPropertyprices []model.Propertyprice

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tProperty = model.Property{
		ID:          utils.GetUniqueID(),
		CompanyID:   req.CompanyID,
		Name:        req.Name,
		Description: req.Description,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	err = u.repository.Create(tx, tProperty)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repository.Name(), err)
	}

	tPropertytimeline = model.Propertytimeline{
		ID:       tProperty.ID,
		CreateBy: loginUser.UserID,
		UpdateBy: loginUser.UserID,
	}

	err = u.repositoryPropertytimeline.Create(tx, tPropertytimeline)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repositoryPropertytimeline.Name(), err)
	}

	for _, unit := range req.Units {
		tUnit := model.Unit{
			CompanyID:   req.CompanyID,
			PropertyID:  tProperty.ID,
			Name:        unit.Name,
			Description: unit.Description,
			PhotoID:     "",
			CreateBy:    loginUser.UserID,
			UpdateBy:    loginUser.UserID,
		}

		tUnits = append(tUnits, tUnit)
	}

	err = u.repositoryUnit.Creates(tx, tUnits)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repositoryPropertytimeline.Name(), err)
	}

	for index, propertyprice := range req.Propertyprices {
		tPropertyprice := model.Propertyprice{
			PropertyID: tProperty.ID,
			Price:      propertyprice.Price,
			Weekdays:   propertyprice.Weekdays,
			Priority:   len(req.Propertyprices) - index,
			CreateBy:   loginUser.UserID,
			UpdateBy:   loginUser.UserID,
		}

		tPropertyprices = append(tPropertyprices, tPropertyprice)
	}

	err = u.repositoryPropertyprice.Creates(tx, tPropertyprices)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repositoryPropertyprice.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateProperty) error {
	var err error
	var tProperty model.Property

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tProperty, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tProperty.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	tProperty.Name = req.Name
	tProperty.Description = req.Description
	tProperty.UpdateBy = loginUser.UserID
	err = u.repository.Save(tx, tProperty)
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
	var tProperty model.Property

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tProperty, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, tProperty.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	err = u.repository.Delete(tx, tProperty)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) GetPrice(req request.GetPrice) (price int64, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	price, err = u.repositoryPropertyprice.GetPrice(conn, req)
	if err != nil {
		return price, err
	}

	return
}

func NewUsecase(repository Repository, repositoryPropertytimeline propertytimeline.Repository, repositoryUnit unit.Repository, repositoryPropertyprice propertyprice.Repository) Usecase {
	return &usecase{
		repository:                 repository,
		repositoryPropertytimeline: repositoryPropertytimeline,
		repositoryUnit:             repositoryUnit,
		repositoryPropertyprice:    repositoryPropertyprice,
	}
}

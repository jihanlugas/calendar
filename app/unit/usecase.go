package unit

import (
	"fmt"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
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
	baseUsecase base.Usecase
	repository  Repository
}

func NewUsecase(baseUsecase base.Usecase, repository Repository) Usecase {
	return &usecase{
		baseUsecase: baseUsecase,
		repository:  repository,
	}
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageUnit) (vUnits []model.UnitView, count int64, err error) {
	conn := u.baseUsecase.GetConnection()

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID)
	if err != nil {
		return vUnits, count, err
	}

	vUnits, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vUnits, count, err
	}

	return vUnits, count, nil
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vUnit model.UnitView, err error) {
	conn := u.baseUsecase.GetConnection()

	vUnit, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vUnit, fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, vUnit.CompanyID); err != nil {
		return vUnit, err
	}

	return vUnit, nil
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateUnit) error {
	var err error

	if err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID); err != nil {
		return err
	}

	conn := u.baseUsecase.GetConnection()

	tUnit := model.Unit{
		ID:          utils.GetUniqueID(),
		CompanyID:   req.CompanyID,
		PropertyID:  req.PropertyID,
		Name:        req.Name,
		Description: req.Description,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := u.repository.Create(tx, tUnit); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to create %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateUnit) error {
	conn := u.baseUsecase.GetConnection()

	tUnit, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, tUnit.CompanyID); err != nil {
		return err
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	tUnit.Name = req.Name
	tUnit.Description = req.Description
	tUnit.UpdateBy = loginUser.UserID
	if err := u.repository.Save(tx, tUnit); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to update %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (u usecase) Delete(loginUser jwt.UserLogin, id string) error {
	conn := u.baseUsecase.GetConnection()

	tUnit, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, tUnit.CompanyID); err != nil {
		return err
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := u.repository.Delete(tx, tUnit); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to delete %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

package property

import (
	"fmt"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/app/propertyprice"
	"github.com/jihanlugas/calendar/app/propertytimeline"
	"github.com/jihanlugas/calendar/app/unit"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/utils"
)

type Usecase interface {
	Page(loginUser jwt.UserLogin, req request.PageProperty) (vProperties []model.PropertyView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProperty model.PropertyView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateProperty) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdateProperty) error
	Delete(loginUser jwt.UserLogin, id string) error
	GetPrice(req request.GetPrice) (price int64, err error)
	SortPropertyPrice(loginUser jwt.UserLogin, id string, req request.SortPropertyPrice) error
}

type usecase struct {
	baseUsecase                base.Usecase
	repository                 Repository
	repositoryPropertytimeline propertytimeline.Repository
	repositoryUnit             unit.Repository
	repositoryPropertyprice    propertyprice.Repository
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageProperty) (vProperties []model.PropertyView, count int64, err error) {
	conn := u.baseUsecase.GetConnection()

	err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID)
	if err != nil {
		return vProperties, count, err
	}

	vProperties, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vProperties, count, err
	}

	return vProperties, count, nil
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vProperty model.PropertyView, err error) {
	conn := u.baseUsecase.GetConnection()

	vProperty, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vProperty, fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, vProperty.CompanyID); err != nil {
		return vProperty, err
	}

	return vProperty, nil
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateProperty) error {
	var err error

	if err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID); err != nil {
		return err
	}

	conn := u.baseUsecase.GetConnection()

	tProperty := model.Property{
		ID:          utils.GetUniqueID(),
		CompanyID:   req.CompanyID,
		Name:        req.Name,
		Description: req.Description,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := u.repository.Create(tx, tProperty); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to create %s: %v", u.repository.Name(), err)
	}

	tPropertytimeline := model.Propertytimeline{
		ID:       tProperty.ID,
		CreateBy: loginUser.UserID,
		UpdateBy: loginUser.UserID,
	}

	if err := u.repositoryPropertytimeline.Create(tx, tPropertytimeline); err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repositoryPropertytimeline.Name(), err)
	}

	var tUnits []model.Unit
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

	if err := u.repositoryUnit.Creates(tx, tUnits); err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repositoryUnit.Name(), err)
	}

	var tPropertyprices []model.Propertyprice
	for index, propertyprice := range req.Propertyprices {
		tPropertyprice := model.Propertyprice{
			CompanyID:  req.CompanyID,
			PropertyID: tProperty.ID,
			StartTime:  propertyprice.StartTime,
			EndTime:    propertyprice.EndTime,
			Price:      propertyprice.Price,
			Weekdays:   propertyprice.Weekdays,
			Priority:   len(req.Propertyprices) - index,
			CreateBy:   loginUser.UserID,
			UpdateBy:   loginUser.UserID,
		}
		tPropertyprices = append(tPropertyprices, tPropertyprice)
	}

	if err := u.repositoryPropertyprice.Creates(tx, tPropertyprices); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to create %s: %v", u.repositoryPropertytimeline.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateProperty) error {
	conn := u.baseUsecase.GetConnection()

	tProperty, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, tProperty.CompanyID); err != nil {
		return err
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	tProperty.Name = req.Name
	tProperty.Description = req.Description
	tProperty.UpdateBy = loginUser.UserID
	if err := u.repository.Save(tx, tProperty); err != nil {
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

	tProperty, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, tProperty.CompanyID); err != nil {
		return err
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
	conn := u.baseUsecase.GetConnection()

	price, err = u.repositoryPropertyprice.GetPrice(conn, req)
	if err != nil {
		return price, err
	}

	return
}

func (u usecase) SortPropertyPrice(loginUser jwt.UserLogin, id string, req request.SortPropertyPrice) error {
	conn := u.baseUsecase.GetConnection()

	tProperty, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, tProperty.CompanyID); err != nil {
		return err
	}

	tx := conn.Begin()

	for _, propertyprice := range req.Propertyprices {
		tPropertyprice, err := u.repositoryPropertyprice.GetTableById(tx, propertyprice.ID)
		if err != nil {
			return fmt.Errorf("failed to get %s: %v", u.repositoryPropertyprice.Name(), err)
		}

		err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, tPropertyprice.CompanyID)
		if err != nil {
			return err
		}

		tPropertyprice.Priority = propertyprice.Priority
		err = u.repositoryPropertyprice.Save(tx, tPropertyprice)
		if err != nil {
			return fmt.Errorf("failed to update %s: %v", u.repositoryPropertyprice.Name(), err)
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func NewUsecase(baseUsecase base.Usecase, repository Repository, repositoryPropertytimeline propertytimeline.Repository, repositoryUnit unit.Repository, repositoryPropertyprice propertyprice.Repository) Usecase {
	return &usecase{
		baseUsecase:                baseUsecase,
		repository:                 repository,
		repositoryPropertytimeline: repositoryPropertytimeline,
		repositoryUnit:             repositoryUnit,
		repositoryPropertyprice:    repositoryPropertyprice,
	}
}

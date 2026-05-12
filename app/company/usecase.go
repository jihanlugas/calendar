package company

import (
	"errors"
	"fmt"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/app/usercompany"
	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/utils"
)

type Usecase interface {
	Update(loginUser jwt.UserLogin, id string, req request.UpdateCompany) (err error)
}

type usecase struct {
	baseUsecase           base.Usecase
	repository            Repository
	repositoryUsercompany usercompany.Repository
}

func NewUsecase(baseUsecase base.Usecase, repository Repository, repositoryUsercompany usercompany.Repository) Usecase {
	return &usecase{
		baseUsecase:           baseUsecase,
		repository:            repository,
		repositoryUsercompany: repositoryUsercompany,
	}
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateCompany) (err error) {
	conn := u.baseUsecase.GetConnection()

	tCompany, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	switch loginUser.Role {
	case constant.RoleAdmin:
	case constant.RoleUser:
		return errors.New("role not allowed")
	case constant.RoleUseradmin:
		vUsercompany, err := u.repositoryUsercompany.GetViewByUserIdAndCompanyId(conn, loginUser.UserID, loginUser.CompanyID)
		if err != nil {
			return fmt.Errorf("failed to get %s: %v", u.repositoryUsercompany.Name(), err)
		}

		err = u.baseUsecase.RequireCompanyIDAllowed(loginUser, vUsercompany.CompanyID)
		if err != nil {
			return err
		}
	default:
		return errors.New("not allowed")
	}

	tx := conn.Begin()

	tCompany.Name = req.Name
	tCompany.Description = req.Description
	tCompany.Email = req.Email
	tCompany.Address = req.Address
	tCompany.PhoneNumber = utils.FormatPhoneTo62(req.PhoneNumber)
	tCompany.UpdateBy = loginUser.UserID
	err = u.repository.Save(tx, tCompany)
	if err != nil {
		return fmt.Errorf("failed to update %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return err

}

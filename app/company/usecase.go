package company

import (
	"errors"
	"fmt"
	"github.com/jihanlugas/calendar/app/usercompany"
	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
	"github.com/jihanlugas/calendar/utils"
)

type Usecase interface {
	Update(loginUser jwt.UserLogin, id string, req request.UpdateCompany) error
}

type usecase struct {
	repository            Repository
	repositoryUsercompany usercompany.Repository
}

func NewUsecase(repository Repository, repositoryUsercompany usercompany.Repository) Usecase {
	return &usecase{
		repository:            repository,
		repositoryUsercompany: repositoryUsercompany,
	}
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateCompany) error {
	var err error
	var tCompany model.Company

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tCompany, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get %s: %v", u.repository.Name(), err))
	}

	switch loginUser.Role {
	case constant.RoleAdmin:
	case constant.RoleUser:
		return errors.New("role not allowed")
	case constant.RoleUseradmin:
		vUsercompany, err := u.repositoryUsercompany.GetViewByUserIdAndCompanyId(conn, loginUser.UserID, loginUser.CompanyID)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to get %s: %v", u.repositoryUsercompany.Name(), err))
		}

		if jwt.IsSaveCompanyIDOR(loginUser, vUsercompany.CompanyID) {
			return errors.New(response.ErrorHandlerIDOR)
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
		return errors.New(fmt.Sprintf("failed to update %s: %v", u.repository.Name(), err))
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err

}

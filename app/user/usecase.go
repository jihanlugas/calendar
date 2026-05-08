package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/app/usercompany"
	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/cryption"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/utils"
	"gorm.io/gorm"
)

type Usecase interface {
	Page(loginUser jwt.UserLogin, req request.PageUser) (vUsers []model.UserView, count int64, err error)
	GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vUser model.UserView, err error)
	Create(loginUser jwt.UserLogin, req request.CreateUser) error
	Update(loginUser jwt.UserLogin, id string, req request.UpdateUser) error
	ChangePassword(loginUser jwt.UserLogin, req request.ChangePassword) error
	Delete(loginUser jwt.UserLogin, id string) error
}

type usecase struct {
	baseUsecase       base.Usecase
	repository        Repository
	repositoryCompany usercompany.Repository
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageUser) (vUsers []model.UserView, count int64, err error) {
	conn := u.baseUsecase.GetConnection()

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, req.CompanyID); err != nil {
		return vUsers, count, err
	}

	vUsers, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vUsers, count, err
	}

	return vUsers, count, nil
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vUser model.UserView, err error) {
	conn := u.baseUsecase.GetConnection()

	vUser, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vUser, fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	vUsercompany, err := u.repositoryCompany.GetViewByUserIdAndCompanyId(conn, vUser.ID, loginUser.CompanyID)
	if err != nil {
		return vUser, fmt.Errorf("failed to get %s: %v", u.repositoryCompany.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, vUsercompany.CompanyID); err != nil {
		return vUser, err
	}

	return vUser, nil
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateUser) error {
	conn := u.baseUsecase.GetConnection()

	now := time.Now()

	encodePasswd, err := cryption.EncryptAES64(req.Passwd)
	if err != nil {
		return fmt.Errorf("failed to encode password: %v", err)
	}

	tUser := model.User{
		ID:                utils.GetUniqueID(),
		Role:              constant.RoleUser,
		Email:             req.Email,
		Username:          req.Username,
		PhoneNumber:       utils.FormatPhoneTo62(req.PhoneNumber),
		Address:           req.Address,
		Fullname:          req.Fullname,
		Passwd:            encodePasswd,
		PassVersion:       1,
		IsActive:          true,
		PhotoID:           "",
		LastLoginDt:       nil,
		BirthDt:           req.BirthDt,
		BirthPlace:        req.BirthPlace,
		AccountVerifiedDt: &now,
		CreateBy:          loginUser.UserID,
		UpdateBy:          loginUser.UserID,
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := u.repository.Create(tx, tUser); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to create %s: %v", u.repository.Name(), err)
	}

	tUsercompany := model.Usercompany{
		ID:               utils.GetUniqueID(),
		UserID:           tUser.ID,
		CompanyID:        loginUser.CompanyID,
		IsDefaultCompany: true,
		IsCreator:        false,
		CreateBy:         loginUser.UserID,
		UpdateBy:         loginUser.UserID,
	}
	if err := u.repositoryCompany.Create(tx, tUsercompany); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to create %s: %v", u.repositoryCompany.Name(), err)
	}

	return tx.Commit().Error
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateUser) error {
	conn := u.baseUsecase.GetConnection()

	tUser, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	vUsercompany, err := u.repositoryCompany.GetViewByUserIdAndCompanyId(conn, tUser.ID, loginUser.CompanyID)
	if err != nil {
		return fmt.Errorf("failed to update %s: %v", u.repositoryCompany.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, vUsercompany.CompanyID); err != nil {
		return err
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if tUser.Email != req.Email {
		_, err = u.repository.GetByEmail(tx, req.Email)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("failed to update %s: %v", u.repository.Name(), err)
			}
		} else {
			return errors.New("email already exist")
		}
	}

	if tUser.PhoneNumber != utils.FormatPhoneTo62(req.PhoneNumber) {
		_, err = u.repository.GetByPhoneNumber(tx, utils.FormatPhoneTo62(req.PhoneNumber))
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("failed to update %s: %v", u.repository.Name(), err)
			}
		} else {
			return errors.New("phone number already exist")
		}
	}

	tUser.Fullname = req.Fullname
	tUser.Email = req.Email
	tUser.PhoneNumber = utils.FormatPhoneTo62(req.PhoneNumber)
	tUser.Username = req.Username
	tUser.Address = req.Address
	tUser.BirthDt = req.BirthDt
	tUser.BirthPlace = req.BirthPlace
	tUser.UpdateBy = loginUser.UserID
	if err := u.repository.Save(tx, tUser); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to update %s: %v", u.repository.Name(), err)
	}

	return tx.Commit().Error
}

func (u usecase) ChangePassword(loginUser jwt.UserLogin, req request.ChangePassword) error {
	conn := u.baseUsecase.GetConnection()

	tUser, err := u.repository.GetTableById(conn, loginUser.UserID)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := cryption.CheckAES64(req.CurrentPasswd, tUser.Passwd); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("invalid current password")
	}

	encodePasswd, err := cryption.EncryptAES64(req.Passwd)
	if err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to encode password: %v", err)
	}

	tUser.Passwd = encodePasswd
	tUser.PassVersion += 1
	tUser.UpdateBy = loginUser.UserID
	if err := u.repository.Save(tx, tUser); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to update password: %v", err)
	}

	return tx.Commit().Error
}

func (u usecase) Delete(loginUser jwt.UserLogin, id string) error {
	conn := u.baseUsecase.GetConnection()

	tUser, err := u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	vUsercompany, err := u.repositoryCompany.GetViewByUserIdAndCompanyId(conn, tUser.ID, loginUser.CompanyID)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repositoryCompany.Name(), err)
	}

	if err := u.baseUsecase.RequireCompanyIDAllowed(loginUser, vUsercompany.CompanyID); err != nil {
		return err
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := u.repository.Delete(tx, tUser); err != nil {
		_ = tx.Rollback().Error
		return fmt.Errorf("failed to delete %s: %v", u.repository.Name(), err)
	}

	return tx.Commit().Error
}

func NewUsecase(baseUsecase base.Usecase, repository Repository, repositoryCompany usercompany.Repository) Usecase {
	return &usecase{
		baseUsecase:       baseUsecase,
		repository:        repository,
		repositoryCompany: repositoryCompany,
	}
}

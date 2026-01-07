package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/jihanlugas/calendar/app/usercompany"
	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/cryption"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
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
	repository        Repository
	repositoryCompany usercompany.Repository
}

func (u usecase) Page(loginUser jwt.UserLogin, req request.PageUser) (vUsers []model.UserView, count int64, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return vUsers, count, errors.New(response.ErrorHandlerIDOR)
	}

	vUsers, count, err = u.repository.Page(conn, req)
	if err != nil {
		return vUsers, count, err
	}

	return vUsers, count, err
}

func (u usecase) GetById(loginUser jwt.UserLogin, id string, preloads ...string) (vUser model.UserView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	vUser, err = u.repository.GetViewById(conn, id, preloads...)
	if err != nil {
		return vUser, fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	vUsercompany, err := u.repositoryCompany.GetViewByUserIdAndCompanyId(conn, vUser.ID, loginUser.CompanyID)
	if err != nil {
		return vUser, fmt.Errorf("failed to get %s: %v", u.repositoryCompany.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vUsercompany.CompanyID) {
		return vUser, errors.New(response.ErrorHandlerIDOR)
	}

	return vUser, err
}

func (u usecase) Create(loginUser jwt.UserLogin, req request.CreateUser) error {
	var err error
	var tUser model.User

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	now := time.Now()

	encodePasswd, err := cryption.EncryptAES64(req.Passwd)
	if err != nil {
		return fmt.Errorf("failed to encode password: %v", err)
	}

	tUser = model.User{
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

	err = u.repository.Create(tx, tUser)
	if err != nil {
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
	err = u.repositoryCompany.Create(tx, tUsercompany)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", u.repositoryCompany.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Update(loginUser jwt.UserLogin, id string, req request.UpdateUser) error {
	var err error
	var tUser model.User

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tUser, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	vUsercompany, err := u.repositoryCompany.GetViewByUserIdAndCompanyId(conn, tUser.ID, loginUser.CompanyID)
	if err != nil {
		return fmt.Errorf("failed to update %s: %v", u.repositoryCompany.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vUsercompany.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

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
	err = u.repository.Save(tx, tUser)
	if err != nil {
		return fmt.Errorf("failed to update %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) ChangePassword(loginUser jwt.UserLogin, req request.ChangePassword) error {
	var err error
	var tUser model.User

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tUser, err = u.repository.GetTableById(conn, loginUser.UserID)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	tx := conn.Begin()

	err = cryption.CheckAES64(req.CurrentPasswd, tUser.Passwd)
	if err != nil {
		return fmt.Errorf("invalid current password")
	}

	encodePasswd, err := cryption.EncryptAES64(req.Passwd)
	if err != nil {
		return fmt.Errorf("failed to encode password: %v", err)
	}

	tUser.Passwd = encodePasswd
	tUser.PassVersion += 1
	tUser.UpdateBy = loginUser.UserID
	err = u.repository.Save(tx, tUser)
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecase) Delete(loginUser jwt.UserLogin, id string) error {
	var err error
	var tUser model.User

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tUser, err = u.repository.GetTableById(conn, id)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	vUsercompany, err := u.repositoryCompany.GetViewByUserIdAndCompanyId(conn, tUser.ID, loginUser.CompanyID)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repositoryCompany.Name(), err)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, vUsercompany.CompanyID) {
		return errors.New(response.ErrorHandlerIDOR)
	}

	tx := conn.Begin()

	err = u.repository.Delete(tx, tUser)
	if err != nil {
		return fmt.Errorf("failed to get %s: %v", u.repository.Name(), err)
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func NewUsecase(repository Repository, repositoryCompany usercompany.Repository) Usecase {
	return &usecase{
		repository:        repository,
		repositoryCompany: repositoryCompany,
	}
}

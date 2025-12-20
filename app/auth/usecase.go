package auth

import (
	"errors"
	"time"

	"github.com/jihanlugas/calendar/app/company"
	"github.com/jihanlugas/calendar/app/user"
	"github.com/jihanlugas/calendar/app/usercompany"
	"github.com/jihanlugas/calendar/config"
	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/cryption"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/utils"
)

type Usecase interface {
	SignIn(req request.Signin) (token string, userLogin jwt.UserLogin, err error)
	RefreshToken(userLogin jwt.UserLogin) (token string, err error)
	Init(userLogin jwt.UserLogin) (vUser model.UserView, err error)
}

type usecase struct {
	userRepository        user.Repository
	companyRepository     company.Repository
	usercompanyRepository usercompany.Repository
}

func (u usecase) SignIn(req request.Signin) (token string, userLogin jwt.UserLogin, err error) {

	var tUser model.User
	var tCompany model.Company
	var tUsercompany model.Usercompany

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if utils.IsValidEmail(req.Username) {
		tUser, err = u.userRepository.GetByEmail(conn, req.Username)
	} else {
		tUser, err = u.userRepository.GetByUsername(conn, req.Username)
	}

	if err != nil {
		return "", userLogin, err
	}

	err = cryption.CheckAES64(req.Passwd, tUser.Passwd)
	if err != nil {
		return "", userLogin, errors.New("invalid username or password")
	}

	if !tUser.IsActive {
		return "", userLogin, errors.New("user not active")
	}

	if tUser.Role != constant.RoleAdmin {
		tUsercompany, err = u.usercompanyRepository.GetCompanyDefaultByUserId(conn, tUser.ID)
		if err != nil {
			return "", userLogin, errors.New("usercompany not found : " + err.Error())
		}

		tCompany, err = u.companyRepository.GetTableById(conn, tUsercompany.CompanyID)
		if err != nil {
			return "", userLogin, errors.New("company not found : " + err.Error())
		}
	}

	now := time.Now()
	tx := conn.Begin()

	tUser.LastLoginDt = &now
	tUser.UpdateBy = tUser.ID
	err = u.userRepository.Update(tx, model.User{
		ID:          tUser.ID,
		LastLoginDt: &now,
		UpdateBy:    tUser.ID,
	})
	if err != nil {
		return "", userLogin, err
	}

	err = tx.Commit().Error
	if err != nil {
		return "", userLogin, err
	}

	expiredAt := time.Now().Add(time.Minute * time.Duration(config.AuthTokenExpiredMinute))
	userLogin.ExpiredDt = expiredAt
	userLogin.UserID = tUser.ID
	userLogin.Role = tUser.Role
	userLogin.PassVersion = tUser.PassVersion
	userLogin.CompanyID = tCompany.ID
	userLogin.UsercompanyID = tUsercompany.ID
	token, err = jwt.CreateToken(userLogin)
	if err != nil {
		return "", userLogin, err
	}

	return token, userLogin, err
}

func (u usecase) RefreshToken(userLogin jwt.UserLogin) (token string, err error) {
	userLogin.ExpiredDt = time.Now().Add(time.Minute * time.Duration(config.AuthTokenExpiredMinute))

	token, err = jwt.CreateToken(userLogin)
	if err != nil {
		return token, err
	}

	return token, err
}

func (u usecase) Init(userLogin jwt.UserLogin) (vUser model.UserView, err error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	userPreloads := []string{"Company", "Company.Properties", "Company.Properties.Propertytimeline", "Company.Properties.Units", "Usercompanies", "Usercompanies.Company", "Usercompanies.User"}
	vUser, err = u.userRepository.GetViewById(conn, userLogin.UserID, userPreloads...)
	if err != nil {
		return vUser, err
	}

	return vUser, err
}

func NewUsecase(userRepository user.Repository, companyRepository company.Repository, usercompanyRepository usercompany.Repository) Usecase {
	return usecase{
		userRepository:        userRepository,
		companyRepository:     companyRepository,
		usercompanyRepository: usercompanyRepository,
	}
}

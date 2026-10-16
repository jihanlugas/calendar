package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/app/company"
	"github.com/jihanlugas/calendar/app/user"
	"github.com/jihanlugas/calendar/app/usercompany"
	"github.com/jihanlugas/calendar/app/userprovider"
	"github.com/jihanlugas/calendar/config"
	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/cryption"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
	"github.com/jihanlugas/calendar/utils"
	"gorm.io/gorm"
)

type Usecase interface {
	SignIn(req request.Signin) (token string, userLogin jwt.UserLogin, err error)
	RefreshToken(userLogin jwt.UserLogin) (token string, err error)
	Init(userLogin jwt.UserLogin) (vUser model.UserView, err error)
	GoogleCallback(callback response.GoogleCallback) (token string, userLogin jwt.UserLogin, err error)
	GoogleLinkCallback(loginUser jwt.UserLogin, callback response.GoogleCallback) (err error)
	GoogleUnlink(loginUser jwt.UserLogin) (err error)
}

type usecase struct {
	baseUsecase            base.Usecase
	userRepository         user.Repository
	companyRepository      company.Repository
	usercompanyRepository  usercompany.Repository
	userproviderRepository userprovider.Repository
}

func (u usecase) SignIn(req request.Signin) (token string, userLogin jwt.UserLogin, err error) {

	var tUser model.User
	var tCompany model.Company
	var tUsercompany model.Usercompany

	conn := u.baseUsecase.GetConnection()

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
		return "", userLogin, fmt.Errorf("failed to commit transaction: %w", err)
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
	conn := u.baseUsecase.GetConnection()

	userPreloads := []string{"Company", "Company.Properties", "Company.Properties.Propertytimeline", "Company.Properties.Units", "Usercompanies", "Usercompanies.Company", "Usercompanies.User"}
	vUser, err = u.userRepository.GetViewById(conn, userLogin.UserID, userPreloads...)
	if err != nil {
		return vUser, err
	}

	return vUser, err
}

func (u usecase) GoogleCallback(callback response.GoogleCallback) (token string, userLogin jwt.UserLogin, err error) {
	var tUserprovider model.Userprovider
	var tCompany model.Company
	var tUsercompany model.Usercompany

	conn := u.baseUsecase.GetConnection()

	tUserprovider, err = u.userproviderRepository.GetTableByProviderNameAndEmail(conn, constant.OauthProviderGoogle, callback.Email, "User")
	if err != nil {
		err = errors.New("akun tidak terdaftar")
		return token, userLogin, err
	}

	if tUserprovider.User == nil {
		err = errors.New("akun tidak valid")
		return token, userLogin, err
	}

	if !tUserprovider.User.IsActive {
		return token, userLogin, errors.New("user not active")
	}

	if tUserprovider.User.Role != constant.RoleAdmin {
		tUsercompany, err = u.usercompanyRepository.GetCompanyDefaultByUserId(conn, tUserprovider.User.ID)
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

	tUserprovider.User.LastLoginDt = &now
	tUserprovider.User.UpdateBy = tUserprovider.User.ID
	err = u.userRepository.Update(tx, model.User{
		ID:          tUserprovider.User.ID,
		LastLoginDt: &now,
		UpdateBy:    tUserprovider.User.ID,
	})
	if err != nil {
		return token, userLogin, err
	}

	expiredAt := time.Now().Add(time.Minute * time.Duration(config.AuthTokenExpiredMinute))
	userLogin.ExpiredDt = expiredAt
	userLogin.UserID = tUserprovider.User.ID
	userLogin.Role = tUserprovider.User.Role
	userLogin.PassVersion = tUserprovider.User.PassVersion
	userLogin.CompanyID = tCompany.ID
	userLogin.UsercompanyID = tUsercompany.ID
	token, err = jwt.CreateToken(userLogin)
	if err != nil {
		return "", userLogin, err
	}

	err = tx.Commit().Error
	if err != nil {
		return token, userLogin, err
	}

	return token, userLogin, err
}

func (u usecase) GoogleLinkCallback(loginUser jwt.UserLogin, callback response.GoogleCallback) (err error) {
	var tUserprovider model.Userprovider

	conn := u.baseUsecase.GetConnection()

	tUserprovider, err = u.userproviderRepository.GetTableByProviderNameAndEmail(conn, constant.OauthProviderGoogle, callback.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if tUserprovider.ID != "" {
		return errors.New("email sudah terdaftar")
	}

	tx := conn.Begin()

	err = u.userproviderRepository.Create(tx, model.Userprovider{
		UserID:         loginUser.UserID,
		ProviderName:   constant.OauthProviderGoogle,
		ProviderUserID: callback.ID,
		Email:          callback.Email,
		Fullname:       callback.Name,
		CreateBy:       loginUser.UserID,
		UpdateBy:       loginUser.UserID,
	})
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err

}

func (u usecase) GoogleUnlink(loginUser jwt.UserLogin) (err error) {
	var tUserprovider model.Userprovider

	conn := u.baseUsecase.GetConnection()

	tUserprovider, err = u.userproviderRepository.GetTableByProviderNameAndUserId(conn, "google", loginUser.UserID)
	if err != nil {
		return err
	}

	tx := conn.Begin()

	err = u.userproviderRepository.Delete(tx, tUserprovider)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err

}

func NewUsecase(baseUsecase base.Usecase, userRepository user.Repository, companyRepository company.Repository, usercompanyRepository usercompany.Repository) Usecase {
	return &usecase{
		baseUsecase:           baseUsecase,
		userRepository:        userRepository,
		companyRepository:     companyRepository,
		usercompanyRepository: usercompanyRepository,
	}
}

package photo

import (
	"errors"
	"fmt"

	"github.com/jihanlugas/calendar/app/base"
	"github.com/jihanlugas/calendar/model"
)

type Usecase interface {
	GetById(id string) (tPhoto model.Photo, err error)
	Upload() (tPhoto model.Photo, err error)
}

type usecase struct {
	baseUsecase base.Usecase
	repository  Repository
}

func (u usecase) GetById(id string) (tPhoto model.Photo, err error) {
	conn, closeConn := u.baseUsecase.WithConn()
	defer closeConn()

	tPhoto, err = u.repository.GetById(conn, id)
	if err != nil {
		return tPhoto, errors.New(fmt.Sprint("failed to get order: ", err))
	}

	return tPhoto, nil
}

func (u usecase) Upload() (tPhoto model.Photo, err error) {
	return tPhoto, nil
}

func NewUsecase(baseUsecase base.Usecase, repository Repository) Usecase {
	return &usecase{
		baseUsecase: baseUsecase,
		repository:  repository,
	}
}

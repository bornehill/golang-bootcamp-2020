package repository

import "github.com/bornehill/golang-bootcamp-2020/domain/model"

type CentreRepository interface {
	GetAll() (*[]*model.Centre, error)
}

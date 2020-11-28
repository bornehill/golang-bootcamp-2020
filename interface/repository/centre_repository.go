package repository

import (
	"github.com/bornehill/golang-bootcamp-2020/domain/model"
	"github.com/bornehill/golang-bootcamp-2020/usecase/repository"
)

type centreRepository struct {
	centres *[]*model.Centre
}

func OpenCentreRepository(centres *[]*model.Centre) repository.CentreRepository {
	return &centreRepository{centres}
}

func (cr *centreRepository) GetAll() (*[]*model.Centre, error) {
	return cr.centres, nil
}

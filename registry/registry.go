package registry

import (
	"github.com/bornehill/golang-bootcamp-2020/infrastructure/datastore"
	"github.com/bornehill/golang-bootcamp-2020/interface/controller"
)

type registry struct {
	db *datastore.DbRepository
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *datastore.DbRepository) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Centre: r.NewCentreController(),
	}
}

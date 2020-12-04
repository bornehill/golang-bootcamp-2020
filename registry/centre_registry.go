package registry

import (
	"github.com/bornehill/golang-bootcamp-2020/interface/controller"
	ip "github.com/bornehill/golang-bootcamp-2020/interface/presenter"
	"github.com/bornehill/golang-bootcamp-2020/usecase/repository"
	"github.com/bornehill/golang-bootcamp-2020/usecase/interactor"
	up "github.com/bornehill/golang-bootcamp-2020/usecase/presenter"
)

func (r *registry) NewCentreController() controller.CentreController {
	return controller.NewCentreController(r.NewCentreInteractor())
}

func (r *registry) NewCentreInteractor() interactor.CentreInteractor {
	return interactor.NewCentreInteractor(r.NewCentreRepository(), r.NewCentrePresenter())
}

func (r *registry) NewCentreRepository() repository.CentreRepository {
	return r.db.Centres
}

func (r *registry) NewCentrePresenter() up.CentrePresenter {
	return ip.NewCentrePresenter()
}

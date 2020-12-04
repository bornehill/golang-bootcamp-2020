package interactor

import (
	"github.com/bornehill/golang-bootcamp-2020/usecase/presenter"
	"github.com/bornehill/golang-bootcamp-2020/usecase/repository"
)

type centreInteractor struct {
	CentrePresenter presenter.CentrePresenter
	CentreRepository repository.CentreRepository
}

type CentreInteractor interface {
	Get() ([]byte, error)
}

func NewCentreInteractor(r repository.CentreRepository, p presenter.CentrePresenter) CentreInteractor {
	return &centreInteractor{p, r}
}

func (ci *centreInteractor) Get() ([]byte, error) {
	data, err := ci.CentreRepository.GetAll()
	if(err != nil) {
		return nil, err
	}
	
	b, err := ci.CentrePresenter.JsonResponse(data)		
	
	if(err == nil) {
		return b, nil
	}

	return nil, err
}

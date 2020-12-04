package controller

import "github.com/bornehill/golang-bootcamp-2020/usecase/interactor"

type centreController struct {
	CentreInteractor interactor.CentreInteractor
}

type CentreController interface {
	GetCentres() ([]byte, error)
}

func NewCentreController(us interactor.CentreInteractor) CentreController {
	return &centreController{us}
}

func (cc *centreController) GetCentres() ([]byte, error) {
	return cc.CentreInteractor.Get()
}

package presenter

import "github.com/bornehill/golang-bootcamp-2020/domain/model"

type CentrePresenter interface {
	JsonResponse(data *[]*model.Centre) ([]byte, error)
}

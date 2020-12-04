package presenter

import (
	"encoding/json"

	"github.com/bornehill/golang-bootcamp-2020/domain/model"
	"github.com/bornehill/golang-bootcamp-2020/usecase/presenter"
)

type centrePresenter struct{}

func NewCentrePresenter() presenter.CentrePresenter {
	return &centrePresenter{}
}

func (cp *centrePresenter) JsonResponse(data *[]*model.Centre) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return b, nil
}

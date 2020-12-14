package presenter

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/bornehill/golang-bootcamp-2020/domain/model"
	mocks "github.com/bornehill/golang-bootcamp-2020/mocks/usecase/presenter"
	"github.com/bornehill/golang-bootcamp-2020/usecase/presenter"
)

func TestNewCentrePresenter(t *testing.T) {
	tests := []struct {
		name string
		want presenter.CentrePresenter
	}{
		{
			name: "Create an AppController success",
			want: &centrePresenter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCentrePresenter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCentrePresenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_centrePresenter_JsonResponse(t *testing.T) {
	testData, testBytes := testData()

	tests := []struct {
		name    string
		cp      *mocks.CentrePresenter
		args    *[]*model.Centre
		want    []byte
		wantErr bool
		err     error
	}{
		{
			name:    "Should parse data",
			cp:      &mocks.CentrePresenter{},
			args:    testData,
			want:    testBytes,
			wantErr: false,
			err:     nil,
		},
		{
			name:    "Should throws error",
			cp:      &mocks.CentrePresenter{},
			args:    nil,
			want:    nil,
			wantErr: true,
			err:     errors.New("sample error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cp.On("JsonResponse", tt.args).Return(tt.want, tt.err)

			got, err := tt.cp.JsonResponse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("centrePresenter.JsonResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("centrePresenter.JsonResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testData() (*[]*model.Centre, []byte) {
	ret := make([]*model.Centre, 0, 0)

	for id := 0; id <= 2; id++ {
		centre := &model.Centre{
			Id:       id,
			Capacity: 100,
			Openness: 10,
			Name:     "Item" + strconv.Itoa(id),
			Address:  "address_" + strconv.Itoa(id),
			Email:    "email@test" + strconv.Itoa(id),
			Phone:    "999-99-999-9" + strconv.Itoa(id),
			Line:     "gym",
		}
		ret = append(ret, centre)
	}

	b, _ := json.Marshal(ret)

	return &ret, b
}

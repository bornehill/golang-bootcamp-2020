package interactor

import (
	"errors"
	"reflect"
	"testing"

	mockinteractor "github.com/bornehill/golang-bootcamp-2020/mocks/usecase/interactor"
	mockpresenter "github.com/bornehill/golang-bootcamp-2020/mocks/usecase/presenter"
	mockrepository "github.com/bornehill/golang-bootcamp-2020/mocks/usecase/repository"
	"github.com/bornehill/golang-bootcamp-2020/usecase/presenter"
	"github.com/bornehill/golang-bootcamp-2020/usecase/repository"
)

func TestNewCentreInteractor(t *testing.T) {
	type args struct {
		r repository.CentreRepository
		p presenter.CentrePresenter
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should create a CentreInteractor",
			args: args{r: &mockrepository.CentreRepository{}, p: &mockpresenter.CentrePresenter{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCentreInteractor(tt.args.r, tt.args.p)

			if _, ok := got.(CentreInteractor); !ok {
				t.Errorf("NewCentreInteractor() = %v, want CentreInteractor", got)
			}
		})
	}
}

func Test_centreInteractor_Get(t *testing.T) {
	tests := []struct {
		name    string
		ci      *mockinteractor.CentreInteractor
		want    []byte
		err     error
		wantErr bool
	}{
		{
			name:    "Should call Get",
			ci:      &mockinteractor.CentreInteractor{},
			want:    []byte{0},
			err:     nil,
			wantErr: false,
		},
		{
			name:    "Should not call Get",
			ci:      &mockinteractor.CentreInteractor{},
			want:    nil,
			err:     errors.New("sample error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt.ci.On("Get").Return(tt.want, tt.err)
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.ci.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("centreInteractor.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("centreInteractor.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

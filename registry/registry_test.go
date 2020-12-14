package registry

import (
	"testing"

	"github.com/bornehill/golang-bootcamp-2020/infrastructure/datastore"
	"github.com/bornehill/golang-bootcamp-2020/interface/controller"
	mockregistry "github.com/bornehill/golang-bootcamp-2020/mocks/registry"
	mockdata "github.com/bornehill/golang-bootcamp-2020/mocks/usecase/repository"
)

func TestNewRegistry(t *testing.T) {
	type args struct {
		db *datastore.DbRepository
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should create a Registry",
			args: args{db: &datastore.DbRepository{Centres: &mockdata.CentreRepository{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRegistry(tt.args.db)

			if _, ok := got.(Registry); !ok {
				t.Errorf("NewRegistry() = %v, want Registry", got)
			}
		})
	}
}

func Test_registry_NewAppController(t *testing.T) {
	tests := []struct {
		name string
		r    *mockregistry.Registry
	}{
		{
			name: "Should create an AppController",
			r:    &mockregistry.Registry{},
		},
	}
	for _, tt := range tests {
		tt.r.On("NewAppController").Return(controller.AppController{})

		t.Run(tt.name, func(t *testing.T) {
			var got interface{} = tt.r.NewAppController()

			if _, ok := got.(controller.AppController); !ok {
				t.Errorf("registry.NewAppController() = %v, want AppController", got)
			}
		})
	}
}

package router

import (
	"net/http"
	"testing"

	"github.com/bornehill/golang-bootcamp-2020/interface/controller"
	mocks "github.com/bornehill/golang-bootcamp-2020/mocks/infrastructure/router"
)

func TestNewAppRouter(t *testing.T) {
	type args struct {
		ac controller.AppController
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should create a AppRouter",
			args: args{ac: controller.AppController{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAppRouter(tt.args.ac)
			if _, ok := got.(AppRouter); !ok {
				t.Errorf("NewAppRouter() = %v, want AppRouter", got)
			}
		})
	}
}

func Test_appRouter_GetCentres(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		ar   *mocks.AppRouter
		args args
	}{
		{
			name: "Should call GetCentres",
			ar:   &mocks.AppRouter{},
			args: args{w: nil, r: &http.Request{}},
		},
	}
	for _, tt := range tests {
		tt.ar.On("GetCentres", tt.args.w, tt.args.r).Return()

		t.Run(tt.name, func(t *testing.T) {
			tt.ar.GetCentres(tt.args.w, tt.args.r)
		})
	}
}

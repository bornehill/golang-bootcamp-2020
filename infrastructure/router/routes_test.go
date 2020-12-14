package router

import (
	"testing"

	"github.com/bornehill/golang-bootcamp-2020/interface/controller"
	"github.com/gorilla/mux"
)

func TestNewRouter(t *testing.T) {
	type args struct {
		prefix string
		ac     controller.AppController
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should create a mux Router",
			args: args{prefix: "test", ac: controller.AppController{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got interface{} = NewRouter(tt.args.prefix, tt.args.ac)

			if _, ok := got.(*mux.Router); !ok {
				t.Errorf("NewRouter() = %v, want mux Router", got)
			}
		})
	}
}

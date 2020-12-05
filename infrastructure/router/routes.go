package router

import (
	"fmt"
	"net/http"

	"github.com/bornehill/golang-bootcamp-2020/interface/controller"

	"github.com/gorilla/mux"
)

func NewRouter(prefix string, c controller.AppController) *mux.Router {
	ar := NewAppRouter(c)
	r := mux.NewRouter()

	api := r.PathPrefix(prefix).Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api booking-time")
	})
	api.HandleFunc("/centres", ar.GetCentres).Methods(http.MethodGet)

	return r
}

package router

import (
	"fmt"
	"net/http"

	"github.com/bornehill/golang-bootcamp-2020/interface/controller"

	"github.com/gorilla/mux"
)

var app controller.AppController

func NewRouter(prefix string, c controller.AppController) *mux.Router {
	app = c
	r := mux.NewRouter()

	api := r.PathPrefix(prefix).Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api booking-time")
	})
	api.HandleFunc("/centres", getCentres).Methods(http.MethodGet)

	return r
}


func getCentres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	b, err := app.Centre.GetCentres()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

package router

import (
	"net/http"

	"github.com/bornehill/golang-bootcamp-2020/interface/controller"
)

type appRouter struct {
	App controller.AppController
}

type AppRouter interface {
	GetCentres(w http.ResponseWriter, r *http.Request)
}

func NewAppRouter(ac controller.AppController) AppRouter {
	return &appRouter{ac}
}

func (ar *appRouter) GetCentres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	b, err := ar.App.Centre.GetCentres()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

package router

import (
	"fmt"
	"net/http"

	"github.com/bornehill/golang-bootcamp-2020/interface/controller"

	"github.com/gorilla/mux"
)

/*NewRouter Create new Mux Router
<param name="prefix">API URL Prefix</param>
<param name="ac">App Controller</param> */
func NewRouter(prefix string, ac controller.AppController) *mux.Router {
	ar := NewAppRouter(ac)
	r := mux.NewRouter()

	api := r.PathPrefix(prefix).Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api booking-time")
	})
	api.HandleFunc("/centres", ar.GetCentres).Methods(http.MethodGet)

	return r
}

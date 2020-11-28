package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bornehill/golang-bootcamp-2020/config"
	"github.com/bornehill/golang-bootcamp-2020/infrastructure/datastore"
	
	"github.com/gorilla/mux"
)

var (
	db *datastore.DbRepository
)

func main() {
	config.ReadConfig()
	db = datastore.OpenDb()

	r := mux.NewRouter()
	api := r.PathPrefix(config.Settings.Server.Prefix).Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api booking-time")
	})
	api.HandleFunc("/centres", getCentres).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":" + config.Settings.Server.Port, r))
}

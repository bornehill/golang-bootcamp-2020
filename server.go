package main

import (
	"log"
	"net/http"

	"github.com/bornehill/golang-bootcamp-2020/config"
	"github.com/bornehill/golang-bootcamp-2020/infrastructure/datastore"
	"github.com/bornehill/golang-bootcamp-2020/infrastructure/router"
	"github.com/bornehill/golang-bootcamp-2020/registry"
)

func main() {
	config.ReadConfig()
	db := datastore.OpenDb()

	reg := registry.NewRegistry(db)
	r := router.NewRouter(config.Settings.Server.Prefix, reg.NewAppController())

	log.Fatal(http.ListenAndServe(":" + config.Settings.Server.Port, r))
}

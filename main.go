package main

import (
	"log"
	"net/http"

	"github.com/trqvel/rest-to-do/api"
	"github.com/trqvel/rest-to-do/db"
)

func main() {

	db.Init()

	router := api.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}

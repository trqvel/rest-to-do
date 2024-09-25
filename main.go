package main

import (
	"log"
	"net/http"

	"github.com/trqvel/rest-to-do/api"
	"github.com/trqvel/rest-to-do/db"
)

func main() {
	// Инициализация базы данных
	db.Init()

	// Настройка маршрутов
	router := api.SetupRoutes()

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8080", router))
}

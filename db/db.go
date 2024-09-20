package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Init() {
	dsn := "user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	var err error
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}
	fmt.Println("Успешное подключение к базе данных")
}

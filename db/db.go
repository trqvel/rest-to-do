package db

import (
	"database/sql"
	"fmt"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var Conn *gorm.DB

func Init() {
	
	dsn := "file:rest-to-do.db?cache=shared"
	var err error

	sqlDB, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	Conn, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка инициализации GORM: ", err)
	}
	fmt.Println("Успешное подключение к базе данных")
}

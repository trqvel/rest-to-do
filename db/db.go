package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Драйвер для GORM
	"gorm.io/gorm"
	_ "modernc.org/sqlite" // Импортируем modernc.org/sqlite для использования без CGo
)

var Conn *gorm.DB

func Init() {
	// Подключение к базе данных SQLite без CGo
	dsn := "file:rest-to-do.db?cache=shared"
	var err error

	// Используем стандартный интерфейс database/sql с драйвером modernc.org/sqlite
	sqlDB, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	// Используем GORM с уже открытым подключением
	Conn, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка инициализации GORM: ", err)
	}
	fmt.Println("Успешное подключение к базе данных")
}

package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// переменная, через которую мы будем работать с БД
var DB *gorm.DB

func InitDB() {
	// в dsn вводим данные, которые мы указали при создании контейнера
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5433 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}

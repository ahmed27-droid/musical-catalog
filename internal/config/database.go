package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=arhimed2734 dbname=musical-catalog port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("не удалось подключиться к БД: " + err.Error())
	}
	return db
}

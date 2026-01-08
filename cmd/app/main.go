package main

import "musical-catalog/internal/config"

func main() {

	db := config.SetupDatabase()

	db.AutoMigrate()

}

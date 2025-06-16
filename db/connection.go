package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	log.Println("Trying to connect to DB")
	dsn := "host=localhost user=gorm password=gorm dbname=books port=5432 sslmode=disable"
	dbConnection, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Failed to connect to db")
		log.Fatal(err)
		panic(err)
	}
	log.Println("Connected to DB!")
	// log.Println("DB migration in progress")
	// dbConnection.AutoMigrate(&models.Books{})
	return dbConnection
}

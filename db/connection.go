package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	dsn := buildConnectionString()
	dbConnection, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Failed to connect to db")
		log.Fatal(err)
		panic(err)
	}
	log.Println("Connected to DB!")
	return dbConnection
}

func buildConnectionString() string {
	pg_user := os.Getenv("PG_USER")
	pg_password := os.Getenv("PG_PASSWORD")
	pg_host := os.Getenv("PG_HOST")
	pg_db := os.Getenv("PG_DB")
	pg_port := os.Getenv("PG_PORT")
	pg_ssl_mode := os.Getenv("PG_SSL_MODE")
	log.Println("Trying to connect to DB")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", pg_host, pg_user, pg_password, pg_db, pg_port, pg_ssl_mode)
	return dsn
}

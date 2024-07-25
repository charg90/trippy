package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() error {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Log environment variables for debugging
	fmt.Printf("Connecting to database with DSN: host=%s user=%s password=%s dbname=%s port=%s sslmode=disable\n",
		host, user, password, dbname, port)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	fmt.Println("Database connected")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

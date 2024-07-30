package main

import (
	"log"
	"trippy/config"
	"trippy/db"
	"trippy/db/models"
)

func init() {
	config.LoadEnv()
	db.ConnectDb()
}

func main() {
	log.Println("Starting database migration...")

	err := db.DB.AutoMigrate(&models.User{}, &models.ActivityType{}, &models.Trip{}, &models.Activity{}, &models.Plan{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migration completed successfully.")
}

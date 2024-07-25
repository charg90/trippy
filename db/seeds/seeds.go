package main

import (
	"log"
	"trippy/config"
	"trippy/db"
	"trippy/db/models"

	"github.com/google/uuid"
)

func init() {
	config.LoadEnv()
	db.ConnectDb()
}

func main() {
	if db.DB == nil {
		log.Fatal("Database not initialized")
	}

	log.Println("Starting seeding process")

	commonActivities := []models.ActivityType{
		{ID: uuid.New(), Type: "Eating at Restaurants", Image: "🍽️"},
		{ID: uuid.New(), Type: "Buying Presents", Image: "🎁"},
		{ID: uuid.New(), Type: "Going on Excursions", Image: "🗺️"},
		{ID: uuid.New(), Type: "Visiting Museums", Image: "🏛️"},
		{ID: uuid.New(), Type: "Hiking", Image: "🥾"},
		{ID: uuid.New(), Type: "Shopping", Image: "🛍️"},
		{ID: uuid.New(), Type: "Sightseeing", Image: "👀"},
		{ID: uuid.New(), Type: "Attending Concerts", Image: "🎶"},
		{ID: uuid.New(), Type: "Watching Movies", Image: "🎬"},
		{ID: uuid.New(), Type: "Going to the Beach", Image: "🏖️"},
		{ID: uuid.New(), Type: "Camping", Image: "🏕️"},
		{ID: uuid.New(), Type: "Visiting Parks", Image: "🌳"},
		{ID: uuid.New(), Type: "Visiting Amusement Parks", Image: "🎢"},
		{ID: uuid.New(), Type: "Cycling", Image: "🚴‍♂️"},
		{ID: uuid.New(), Type: "Visiting Historical Sites", Image: "🏰"},
		{ID: uuid.New(), Type: "Picnicking", Image: "🥪"},
		{ID: uuid.New(), Type: "Attending Festivals", Image: "🎉"},
	}

	for _, activity := range commonActivities {
		log.Printf("Creating activity: %v", activity)
		if err := db.DB.Create(&activity).Error; err != nil {
			log.Fatalf("failed to create activity: %v", err)
		}
	}

	log.Println("Seeding process completed")
}

package database

import (
	"log"
	"travel-risk-assessment/models"
)

func Migrate() {
	Instance.AutoMigrate(&models.User{}, &models.Photo{}, &models.Symptom{}, &models.Disease{}, &models.TravelHistory{})
	log.Println("Database Migration Completed!")
}

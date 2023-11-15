package database

import (
	"log"
	"travel-risk-assessment/models"
)

func Migrate() {
	Instance.AutoMigrate(&models.User{}, &models.Photo{}, &models.Symptom{})
	log.Println("Database Migration Completed!")
}

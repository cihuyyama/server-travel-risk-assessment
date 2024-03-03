package database

import (
	"log"
	"net/http"
	"travel-risk-assessment/models"

	"github.com/gin-gonic/gin"
)

func Migrate() {
	Instance.AutoMigrate(
		&models.User{},
		&models.Photo{},
		&models.Symptom{},
		&models.Disease{},
		&models.TravelHistory{},
		&models.MedicalHistory{},
		&models.Endemicity{},
		&models.Prevention{},
		&models.Treatment{},
		&models.DrugConflict{},
		&models.LogHistory{},
		&models.MedicalScoreRisk{},
		&models.TravelScoreRisk{})
	log.Println("Database Migration Completed!")
}

func ManualMigrate(context *gin.Context) {
	Instance.Migrator().CreateTable(
		&models.User{},
		&models.Photo{},
		&models.Symptom{},
		&models.Disease{},
		&models.TravelHistory{},
		&models.MedicalHistory{},
		&models.Endemicity{},
		&models.Prevention{},
		&models.Treatment{},
		&models.DrugConflict{},
		&models.LogHistory{},
		&models.MedicalScoreRisk{},
		&models.TravelScoreRisk{})
	log.Println("Database Migration Completed!")
	context.JSON(http.StatusCreated, gin.H{"message": "Database Migration Completed!"})
}

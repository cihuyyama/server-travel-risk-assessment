package controllers

import (
	"net/http"
	"travel-risk-assessment/app"
	"travel-risk-assessment/database"
	"travel-risk-assessment/models"

	"github.com/gin-gonic/gin"
)

func CountDiseaseSymptom(diseaseID uint, symptomsID []uint) (int, error) {

	var count int64
	if err := database.Instance.Model(&models.Symptom{}).
		Joins("JOIN disease_symptoms ON symptoms.id = disease_symptoms.symptom_id").
		Where("disease_symptoms.disease_id = ? AND symptoms.id IN (?)", diseaseID, symptomsID).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func PostTravelList(context *gin.Context) {
	var arr app.PostTravelForm
	if err := context.ShouldBindJSON(&arr); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var diseases []models.Disease
	if err := database.Instance.Find(&diseases).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	var result []app.PostTravelResponse

	for i := 1; i <= len(diseases); i++ {
		count, err := CountDiseaseSymptom(uint(i), arr.Symptoms)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		count = count * 100 / len(arr.Symptoms)
		result = append(result, app.PostTravelResponse{
			DiseaseId:   diseases[i-1].ID,
			DiseaseName: diseases[i-1].DiseaseName,
			Percentage:  count,
		})
	}

	context.JSON(http.StatusOK, gin.H{"data": result, "status": "success"})
}

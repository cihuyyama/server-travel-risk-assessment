package controllers

import (
	"net/http"
	"travel-risk-assessment/database"
	"travel-risk-assessment/models"

	"github.com/gin-gonic/gin"
)

func GetPreTravelProps(context *gin.Context) {
	var endemic models.Endemicity
	if err := database.Instance.
		Preload("DiseaseEndemic").
		Preload("DiseaseEndemic.Treatment").
		Preload("DiseaseEndemic.Prevention").
		Where("id = ?", context.Param("id")).
		First(&endemic).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Endemic tidak ditemukan", "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": endemic, "status": "success"})
}

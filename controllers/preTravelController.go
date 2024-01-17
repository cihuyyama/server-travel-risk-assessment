package controllers

import (
	"net/http"
	"travel-risk-assessment/app"
	"travel-risk-assessment/database"
	"travel-risk-assessment/models"

	"github.com/gin-gonic/gin"
)

func GetPreTravelProps(context *gin.Context) {
	var disease models.Disease
	if err := database.Instance.Where("id = ?", context.Param("id")).First(&disease).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Penyakit tidak ditemukan", "status": "error"})
		return
	}

	var endemic models.Endemicity
	if err := database.Instance.Where("id = ?", context.Param("id")).First(&endemic).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Endemic tidak ditemukan", "status": "error"})
		return
	}

	var preventions []models.Prevention
	if err := database.Instance.Table("preventions").Where("preventions.disease_id = ?", context.Param("id")).Scan(&preventions).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	var preTravelProps = app.PreTravelProps{
		DiseaseID:   disease.ID,
		DiseaseName: disease.DiseaseName,
		DiseaseDesc: disease.DiseaseDesc,
		Province:    endemic.Province,
		RiskLevel:   endemic.RiskLevel,
		Prevention:  preventions,
	}

	context.JSON(http.StatusOK, gin.H{"data": preTravelProps, "status": "success"})
}

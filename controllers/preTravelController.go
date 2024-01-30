package controllers

import (
	"net/http"
	"strings"
	"travel-risk-assessment/database"
	"travel-risk-assessment/helpers"
	"travel-risk-assessment/models"

	"github.com/gin-gonic/gin"
)

func GetPreTravelProps(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No bearer", "status": "error"})
		return
	}

	tokenString = parts[1]

	claims, err := helpers.ParseToken(tokenString)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		context.Abort()
		return
	}
	var user models.User
	if err := database.Instance.Where("id = ?", claims.ID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Pengguna tidak ditemukan", "status": "error"})
		return
	}

	var medicalHistory models.MedicalHistory
	if err := database.Instance.Where("id = ?", user.ID).First(&medicalHistory).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

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

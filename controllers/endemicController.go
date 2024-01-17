package controllers

import (
	"net/http"
	"travel-risk-assessment/app"
	"travel-risk-assessment/database"
	"travel-risk-assessment/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateEndemic(context *gin.Context) {
	var endemicForm app.EndemicForm
	if err := context.ShouldBindJSON(&endemicForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(endemicForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	endemic := models.Endemicity{
		DiseaseID: endemicForm.DiseaseID,
		Province:  endemicForm.Province,
		RiskLevel: endemicForm.RiskLevel,
	}

	if err := database.Instance.Create(&endemic).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Endemic berhasil ditambahkan", "status": "success"})
}

func GetAllEndemics(context *gin.Context) {
	var endemics []models.Endemicity
	if err := database.Instance.Find(&endemics).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": endemics, "status": "success"})
}

func GetEndemicByID(context *gin.Context) {
	var endemic models.Endemicity
	if err := database.Instance.Where("id = ?", context.Param("id")).First(&endemic).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Endemic tidak ditemukan", "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": endemic, "status": "success"})
}

func UpdateEndemic(context *gin.Context) {
	var endemic models.Endemicity
	if err := database.Instance.Where("id = ?", context.Param("id")).First(&endemic).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Endemic tidak ditemukan", "status": "error"})
		return
	}

	var endemicForm app.EndemicForm
	if err := context.ShouldBindJSON(&endemicForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(endemicForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	endemic.DiseaseID = endemicForm.DiseaseID
	endemic.Province = endemicForm.Province
	endemic.RiskLevel = endemicForm.RiskLevel

	if err := database.Instance.Save(&endemic).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Endemic berhasil diupdate", "status": "success"})
}

func DeleteEndemic(context *gin.Context) {
	var endemic models.Endemicity
	if err := database.Instance.Where("id = ?", context.Param("id")).First(&endemic).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Endemic tidak ditemukan", "status": "error"})
		return
	}

	if err := database.Instance.Delete(&endemic).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Endemic berhasil dihapus", "status": "success"})
}

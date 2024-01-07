package controllers

import (
	"net/http"
	"travel-risk-assessment/app"
	"travel-risk-assessment/database"
	"travel-risk-assessment/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateDisease(context *gin.Context) {
	var diseaseFormCreate app.DiseaseFormCreate
	if err := context.ShouldBindJSON(&diseaseFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(diseaseFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	disease := models.Disease{
		DiseaseName: diseaseFormCreate.DiseaseName,
		DiseaseDesc: diseaseFormCreate.DiseaseDesc,
	}

	if err := database.Instance.Create(&disease).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Penyakit berhasil ditambahkan", "status": "success"})
}

func GetAllDiseases(context *gin.Context) {
	var diseases []models.Disease
	if err := database.Instance.Find(&diseases).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": diseases, "status": "success"})
}

func GetDiseaseByID(context *gin.Context) {
	var disease models.Disease
	if err := database.Instance.Where("id = ?", context.Param("id")).First(&disease).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Penyakit tidak ditemukan", "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": disease, "status": "success"})
}

func UpdateDiseaseByID(context *gin.Context) {
	var diseaseFormCreate app.DiseaseFormCreate
	if err := context.ShouldBindJSON(&diseaseFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(diseaseFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	disease := models.Disease{
		DiseaseName: diseaseFormCreate.DiseaseName,
		DiseaseDesc: diseaseFormCreate.DiseaseDesc,
	}

	if err := database.Instance.Where("id = ?", context.Param("id")).Updates(&disease).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Penyakit tidak ditemukan", "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Penyakit berhasil diupdate", "status": "success"})
}

func DeleteDiseaseByID(context *gin.Context) {
	var disease models.Disease
	if err := database.Instance.Where("id = ?", context.Param("id")).Delete(&disease).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Penyakit tidak ditemukan", "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Penyakit berhasil dihapus", "status": "success"})
}

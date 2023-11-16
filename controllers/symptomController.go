package controllers

import (
	"net/http"
	"strconv"
	"travel-risk-assessment/app"
	"travel-risk-assessment/database"
	"travel-risk-assessment/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateSymptom(context *gin.Context) {
	var symptomFormCreate app.SymptomFormCreate
	if err := context.ShouldBindJSON(&symptomFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(symptomFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	// tokenString := context.GetHeader("Authorization")
	// claims, err := helpers.ParseToken(tokenString)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	context.Abort()
	// 	return
	// }

	symptom := models.Symptom{
		SymptomName: symptomFormCreate.SymptomName,
		SymptomChar: symptomFormCreate.SymptomChar,
	}

	if err := database.Instance.Create(&symptom).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Gejala berhasil ditambahkan", "status": "success"})
}

func GetAllSymptoms(context *gin.Context) {
	var symptoms []models.Symptom

	if err := database.Instance.Find(&symptoms).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": symptoms, "status": "success"})
}

func GetSymptomByID(context *gin.Context) {
	id := context.Param("id")

	var symptomModel models.Symptom
	if err := database.Instance.Where("id = ?", id).First(&symptomModel).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Symptom tidak ditemukan", "status": "error"})
		return
	}

	var symptom app.SymptomResult
	if err := database.Instance.Table("symptoms").Select("symptoms.id, symptoms.symptom_name, symptoms.symptom_char").Where("symptoms.id = ?", id).Scan(&symptom).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": symptom, "status": "success"})
}

func UpdateSymptom(context *gin.Context) {
	symptomID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID foto tidak valid", "status": "error"})
		return
	}
	var symptomFormUpdate app.SymptomFormUpdate
	if err := context.ShouldBindJSON(&symptomFormUpdate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(symptomFormUpdate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	var symptom models.Symptom

	if err := database.Instance.Where("id = ?", symptomID).First(&symptom).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Symptom tidak ditemukan", "status": "error"})
		return
	}

	symptom.SymptomName = symptomFormUpdate.SymptomName
	symptom.SymptomChar = symptomFormUpdate.SymptomChar
	if err := database.Instance.Save(&symptom).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Symptom berhasil diupdate", "status": "success"})

}

func DeleteSymptom(context *gin.Context) {
	symptomID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID foto tidak valid", "status": "error"})
		return
	}

	var symptom models.Symptom

	if err := database.Instance.Where("id = ?", symptomID).First(&symptom).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Symptom tidak ditemukan", "status": "error"})
		return
	}

	if err := database.Instance.Delete(&symptom).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Symptom berhasil dihapus", "status": "success"})
}

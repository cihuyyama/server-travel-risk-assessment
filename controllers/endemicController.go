package controllers

import (
	"fmt"
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
		Province:  endemicForm.Province,
		RiskLevel: endemicForm.RiskLevel,
		RiskScore: endemicForm.RiskScore,
	}

	if err := database.Instance.Create(&endemic).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Endemic berhasil ditambahkan", "status": "success"})
}

func GetAllEndemics(context *gin.Context) {
	var endemics []models.Endemicity
	if err := database.Instance.
		Preload("DiseaseEndemic").
		Preload("DiseaseEndemic.Treatment").
		Preload("DiseaseEndemic.Prevention").
		Preload("DiseaseEndemic.DiseaseSymptom").
		Find(&endemics).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": endemics, "status": "success"})
}

func GetEndemicByID(context *gin.Context) {
	var endemic models.Endemicity
	if err := database.Instance.
		Preload("DiseaseEndemic").
		Preload("DiseaseEndemic.Treatment").
		Preload("DiseaseEndemic.Prevention").
		Preload("DiseaseEndemic.DiseaseSymptom").
		Where("id = ?", context.Param("id")).First(&endemic).Error; err != nil {
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

	endemic.Province = endemicForm.Province
	endemic.RiskLevel = endemicForm.RiskLevel
	endemic.RiskScore = endemicForm.RiskScore

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

func AppendDisease(context *gin.Context) {
	var appender app.AppendDiseaseForm
	if err := context.ShouldBindJSON(&appender); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(appender); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	var endemic models.Endemicity
	if err := database.Instance.Where("id = ?", appender.EndemicID).First(&endemic).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Endemic tidak ditemukan", "status": "error"})
		return
	}

	var disease models.Disease
	if err := database.Instance.Where("id = ?", appender.DiseaseID).First(&disease).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Penyakit tidak ditemukan", "status": "error"})
		return
	}

	if err := database.Instance.Model(&endemic).Association("DiseaseEndemic").Append(&disease); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	CreateLogHistory(endemic.Province, disease.DiseaseName, "ADD")

	message := fmt.Sprintf("Penyakit %s berhasil ditambahkan ke endemis %s", disease.DiseaseName, endemic.Province)
	context.JSON(http.StatusOK, gin.H{"message": message, "status": "success"})
}

func UnappendDisease(context *gin.Context) {
	var appender app.AppendDiseaseForm
	if err := context.ShouldBindJSON(&appender); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(appender); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	var endemic models.Endemicity
	if err := database.Instance.Where("id = ?", appender.EndemicID).First(&endemic).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Endemic tidak ditemukan", "status": "error"})
		return
	}

	var disease models.Disease
	if err := database.Instance.Where("id = ?", appender.DiseaseID).First(&disease).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Penyakit tidak ditemukan", "status": "error"})
		return
	}

	if err := database.Instance.Model(&endemic).Association("DiseaseEndemic").Delete(&disease); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	CreateLogHistory(endemic.Province, disease.DiseaseName, "DELETE")

	message := fmt.Sprintf("Penyakit %s berhasil dihapus dari endemis %s", disease.DiseaseName, endemic.Province)
	context.JSON(http.StatusOK, gin.H{"message": message, "status": "success"})
}

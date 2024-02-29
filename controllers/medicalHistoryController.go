package controllers

import (
	"net/http"
	"strings"
	"travel-risk-assessment/app"
	"travel-risk-assessment/database"
	"travel-risk-assessment/helpers"
	"travel-risk-assessment/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateMedicalHistory(context *gin.Context) {
	var medicalHistoryFormCreate app.MedicalHistoryForm
	if err := context.ShouldBindJSON(&medicalHistoryFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(medicalHistoryFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

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

	medicalHistory := models.MedicalHistory{
		UserID:               user.ID,
		Age:                  medicalHistoryFormCreate.Age,
		PreexistingCondition: medicalHistoryFormCreate.PreexistingCondition,
		CurrentMedication:    medicalHistoryFormCreate.CurrentMedication,
		Allergies:            medicalHistoryFormCreate.Allergies,
		PreviousVaccination:  medicalHistoryFormCreate.PreviousVaccination,
		Pregnant:             medicalHistoryFormCreate.Pregnant,
		VaccineBcg:           medicalHistoryFormCreate.VaccineBcg,
		VaccineHepatitis:     medicalHistoryFormCreate.VaccineHepatitis,
		VaccineDengue:        medicalHistoryFormCreate.VaccineDengue,
	}

	if err := database.Instance.Create(&medicalHistory).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Riwayat medis berhasil ditambahkan", "status": "success"})
}

func GetAllMedicalHistory(context *gin.Context) {
	var medicalHistory []models.MedicalHistory
	if err := database.Instance.Find(&medicalHistory).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": medicalHistory, "status": "success"})
}

func GetMedicalHistoryByID(context *gin.Context) {
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
	if err := database.Instance.Where("user_id = ?", user.ID).First(&medicalHistory).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": medicalHistory, "status": "success"})
}

func UpdateMedicalHistory(context *gin.Context) {
	var medicalHistoryFormUpdate app.MedicalHistoryForm
	if err := context.ShouldBindJSON(&medicalHistoryFormUpdate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(medicalHistoryFormUpdate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	tokenString := context.GetHeader("Authorization")
	// Split the "Authorization" header to remove the "Bearer " prefix
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No bearer", "status": "error"})
		return
	}

	// Get the token from the split
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
	if err := database.Instance.Where("id = ? AND user_id = ?", context.Param("id"), user.ID).First(&medicalHistory).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	medicalHistory.Age = medicalHistoryFormUpdate.Age
	medicalHistory.PreexistingCondition = medicalHistoryFormUpdate.PreexistingCondition
	medicalHistory.CurrentMedication = medicalHistoryFormUpdate.CurrentMedication
	medicalHistory.Allergies = medicalHistoryFormUpdate.Allergies
	medicalHistory.PreviousVaccination = medicalHistoryFormUpdate.PreviousVaccination
	medicalHistory.Pregnant = medicalHistoryFormUpdate.Pregnant
	medicalHistory.VaccineBcg = medicalHistoryFormUpdate.VaccineBcg
	medicalHistory.VaccineHepatitis = medicalHistoryFormUpdate.VaccineHepatitis
	medicalHistory.VaccineDengue = medicalHistoryFormUpdate.VaccineDengue

	if err := database.Instance.Save(&medicalHistory).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Riwayat medis berhasil diupdate", "status": "success"})
}

func DeleteMedicalHistory(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	// Split the "Authorization" header to remove the "Bearer " prefix
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No bearer", "status": "error"})
		return
	}

	// Get the token from the split
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
	if err := database.Instance.Where("id = ? AND user_id = ?", context.Param("id"), user.ID).First(&medicalHistory).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	if err := database.Instance.Delete(&medicalHistory).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Riwayat medis berhasil dihapus", "status": "success"})
}

func CreateMedicalScoreRisk(context *gin.Context) {
	var medicalScoreRisk app.MedicalScoreRiskForm
	if err := context.ShouldBindJSON(&medicalScoreRisk); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(medicalScoreRisk); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

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

	totalScore := medicalScoreRisk.Age + medicalScoreRisk.PreexistingCondition + medicalScoreRisk.CurrentMedication + medicalScoreRisk.Allergies + medicalScoreRisk.PreviousVaccination + medicalScoreRisk.Pregnant

	category := ""
	if totalScore > 60 {
		category = "Tinggi"
	} else if totalScore > 20 {
		category = "Medium"
	} else if totalScore > 10 {
		category = "Rendah"
	} else {
		category = "Tidak ada Resiko"
	}

	medicalScoreRiskModel := models.MedicalScoreRisk{
		UserID:               user.ID,
		Age:                  medicalScoreRisk.Age,
		PreexistingCondition: medicalScoreRisk.PreexistingCondition,
		CurrentMedication:    medicalScoreRisk.CurrentMedication,
		Allergies:            medicalScoreRisk.Allergies,
		PreviousVaccination:  medicalScoreRisk.PreviousVaccination,
		Pregnant:             medicalScoreRisk.Pregnant,
		TotalScore:           totalScore,
		Categories:           category,
	}

	if err := database.Instance.Create(&medicalScoreRiskModel).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Score medis berhasil ditambahkan", "status": "success"})
}

func GetAllMedicalScoreRisk(context *gin.Context) {
	var medicalScoreRisk []models.MedicalScoreRisk
	if err := database.Instance.Find(&medicalScoreRisk).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": medicalScoreRisk, "status": "success"})
}

func GetMedicalScoreRiskByID(context *gin.Context) {
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

	var medicalScoreRisk models.MedicalScoreRisk
	if err := database.Instance.Where("user_id = ?", user.ID).First(&medicalScoreRisk).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": medicalScoreRisk, "status": "success"})
}

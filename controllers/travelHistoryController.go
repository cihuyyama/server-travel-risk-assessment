package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"travel-risk-assessment/app"
	"travel-risk-assessment/database"
	"travel-risk-assessment/helpers"
	"travel-risk-assessment/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateTravelHistory(context *gin.Context) {
	var travelFormCreate app.TravelHistoryForm
	if err := context.ShouldBindJSON(&travelFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(travelFormCreate); err != nil {
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

	travel := models.TravelHistory{
		UserID:        user.ID,
		City:          travelFormCreate.City,
		Province:      travelFormCreate.Province,
		Duration:      travelFormCreate.Duration,
		TravelPurpose: travelFormCreate.TravelPurpose,
	}

	if err := database.Instance.Create(&travel).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Travel History berhasil ditambahkan", "status": "success"})
}

func GetAllHistory(context *gin.Context) {
	var travel []models.TravelHistory
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

	if err := database.Instance.Table("travel_histories").Select("*").Where("travel_histories.user_id = ?", claims.ID).Scan(&travel).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"data": "asjkdnaks", "message": err.Error(), "status": "error"})
		return
	}

	if len(travel) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"data": []string{}, "message": "Travel History tidak ditemukan", "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": travel, "status": "success"})
}

func GetTravelHistoryByID(context *gin.Context) {
	id := context.Param("id")
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

	var travel models.TravelHistory
	if err := database.Instance.Where("id = ? AND user_id = ?", id, claims.ID).First(&travel).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"data": []string{}, "message": "history tidak ditemukan", "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": travel, "status": "success"})
}

func UpdateTravelhistory(context *gin.Context) {
	travelID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID tidak valid", "status": "error"})
		return
	}
	var travelForm app.TravelHistoryForm
	if err := context.ShouldBindJSON(&travelForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(travelForm); err != nil {
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

	var travel models.TravelHistory

	if err := database.Instance.Where("id = ? AND user_id = ?", travelID, user.ID).First(&travel).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "history tidak ditemukan", "status": "error"})
		return
	}

	travel.City = travelForm.City
	travel.Province = travelForm.Province
	travel.Duration = travelForm.Duration
	travel.TravelPurpose = travelForm.TravelPurpose
	travel.UpdatedAt = time.Now()
	if err := database.Instance.Save(&travel).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "history berhasil diupdate", "status": "success"})

}

func DeleteTravelHistory(context *gin.Context) {
	travelID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID tidak valid", "status": "error"})
		return
	}

	var travel models.TravelHistory

	if err := database.Instance.Where("id = ?", travelID).First(&travel).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "history tidak ditemukan", "status": "error"})
		return
	}

	if err := database.Instance.Delete(&travel).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "history berhasil dihapus", "status": "success"})
}

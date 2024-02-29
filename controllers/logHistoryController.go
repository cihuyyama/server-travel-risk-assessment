package controllers

import (
	"fmt"
	"net/http"
	"time"
	"travel-risk-assessment/app"
	"travel-risk-assessment/database"
	"travel-risk-assessment/models"

	"github.com/gin-gonic/gin"
)

func CreateLogHistory(province string, diseaseName string, updatedType string) {
	logHistory := models.LogHistory{
		Province:    province,
		DiseaseName: diseaseName,
		UpdatedType: updatedType,
	}

	if err := database.Instance.Create(&logHistory).Error; err != nil {
		fmt.Println(err)
		return
	}
}

func GetAllLogHistory(context *gin.Context) {
	var logHistory []models.LogHistory

	if err := database.Instance.Order("created_at DESC").Find(&logHistory).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	logResults := make([]app.LogResult, len(logHistory))
	for i, history := range logHistory {
		createdDateTime := history.CreatedAt.In(loc)
		createdDate := createdDateTime.Format("02-01-2006")
		createdTime := createdDateTime.Format("15:04:05")

		logResults[i] = app.LogResult{
			ID:          history.ID,
			Province:    history.Province,
			DiseaseName: history.DiseaseName,
			CreatedDate: createdDate,
			CreatedTime: createdTime,
			UpdatedType: history.UpdatedType,
		}
	}

	context.JSON(http.StatusOK, gin.H{"data": logResults, "status": "success"})
}

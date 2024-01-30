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
	if err := database.Instance.
		Preload("Treatment").
		Preload("Prevention").
		Preload("DiseaseSymptom").
		Find(&diseases).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": diseases, "status": "success"})
}

func GetDiseaseByID(context *gin.Context) {
	var disease models.Disease
	if err := database.Instance.
		Preload("Treatment").
		Preload("Prevention").
		Preload("DiseaseSymptom").
		Where("id = ?", context.Param("id")).
		First(&disease).Error; err != nil {
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

func CreateTreatment(context *gin.Context) {
	var treatmentFormCreate app.TreatmentForm
	if err := context.ShouldBindJSON(&treatmentFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(treatmentFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	treatment := models.Treatment{
		DiseaseID:   treatmentFormCreate.DiseaseID,
		Title:       treatmentFormCreate.Title,
		Description: treatmentFormCreate.Description,
	}

	if err := database.Instance.Create(&treatment).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Treatment berhasil ditambahkan", "status": "success"})
}

func GetAllTreatments(context *gin.Context) {
	var treatments []models.Treatment
	if err := database.Instance.Find(&treatments).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": treatments, "status": "success"})
}

func GetTreatmentByID(context *gin.Context) {
	var treatment models.Treatment
	if err := database.Instance.Where("id = ?", context.Param("id")).First(&treatment).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Treatment tidak ditemukan", "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": treatment, "status": "success"})
}

func UpdateTreatmentByID(context *gin.Context) {
	var treatmentFormCreate app.TreatmentForm
	if err := context.ShouldBindJSON(&treatmentFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(treatmentFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	treatment := models.Treatment{
		DiseaseID:   treatmentFormCreate.DiseaseID,
		Title:       treatmentFormCreate.Title,
		Description: treatmentFormCreate.Description,
	}

	if err := database.Instance.Where("id = ?", context.Param("id")).Updates(&treatment).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Treatment tidak ditemukan", "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Treatment berhasil diupdate", "status": "success"})
}

func DeleteTreatmentByID(context *gin.Context) {
	var treatment models.Treatment
	if err := database.Instance.Where("id = ?", context.Param("id")).Delete(&treatment).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Treatment tidak ditemukan", "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Treatment berhasil dihapus", "status": "success"})
}

func CreatePrevention(context *gin.Context) {
	var preventionFormCreate app.PreventionForm
	if err := context.ShouldBindJSON(&preventionFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(preventionFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	prevention := models.Prevention{
		DiseaseID:   preventionFormCreate.DiseaseID,
		Title:       preventionFormCreate.Title,
		Description: preventionFormCreate.Description,
	}

	if err := database.Instance.Create(&prevention).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Prevention berhasil ditambahkan", "status": "success"})
}

func GetAllPreventions(context *gin.Context) {
	var preventions []models.Prevention
	if err := database.Instance.Find(&preventions).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": preventions, "status": "success"})
}

func GetPreventionByID(context *gin.Context) {
	var prevention models.Prevention
	if err := database.Instance.Where("id = ?", context.Param("id")).First(&prevention).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Prevention tidak ditemukan", "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": prevention, "status": "success"})
}

func UpdatePreventionByID(context *gin.Context) {
	var preventionFormCreate app.PreventionForm
	if err := context.ShouldBindJSON(&preventionFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(preventionFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	prevention := models.Prevention{
		DiseaseID:   preventionFormCreate.DiseaseID,
		Title:       preventionFormCreate.Title,
		Description: preventionFormCreate.Description,
	}

	if err := database.Instance.Where("id = ?", context.Param("id")).Updates(&prevention).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Prevention tidak ditemukan", "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Prevention berhasil diupdate", "status": "success"})
}

func DeletePreventionByID(context *gin.Context) {
	var prevention models.Prevention
	if err := database.Instance.Where("id = ?", context.Param("id")).Delete(&prevention).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Prevention tidak ditemukan", "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Prevention berhasil dihapus", "status": "success"})
}

func CreateDrugConflict(context *gin.Context) {
	var drugConflictFormCreate app.DrugConflict
	if err := context.ShouldBindJSON(&drugConflictFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(drugConflictFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	drugConflict := models.DrugConflict{
		DiseaseID: drugConflictFormCreate.DiseaseID,
		DrugName:  drugConflictFormCreate.DrugName,
		DrugDesc:  drugConflictFormCreate.DrugDesc,
	}

	if err := database.Instance.Create(&drugConflict).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Drug Conflict berhasil ditambahkan", "status": "success"})
}

func GetAllDrugConflicts(context *gin.Context) {
	var drugConflicts []models.DrugConflict
	if err := database.Instance.Find(&drugConflicts).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": drugConflicts, "status": "success"})
}

func GetDrugConflictByID(context *gin.Context) {
	var drugConflict models.DrugConflict
	if err := database.Instance.Where("id = ?", context.Param("id")).First(&drugConflict).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Drug Conflict tidak ditemukan", "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": drugConflict, "status": "success"})
}

func UpdateDrugConflictByID(context *gin.Context) {
	var drugConflictFormCreate app.DrugConflict
	if err := context.ShouldBindJSON(&drugConflictFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(drugConflictFormCreate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	drugConflict := models.DrugConflict{
		DiseaseID: drugConflictFormCreate.DiseaseID,
		DrugName:  drugConflictFormCreate.DrugName,
		DrugDesc:  drugConflictFormCreate.DrugDesc,
	}

	if err := database.Instance.Where("id = ?", context.Param("id")).Updates(&drugConflict).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Drug Conflict tidak ditemukan", "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Drug Conflict berhasil diupdate", "status": "success"})
}

func DeleteDrugConflictByID(context *gin.Context) {
	var drugConflict models.DrugConflict
	if err := database.Instance.Where("id = ?", context.Param("id")).Delete(&drugConflict).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Drug Conflict tidak ditemukan", "status": "error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Drug Conflict berhasil dihapus", "status": "success"})
}

func AppendSymptomToDisease(context *gin.Context) {
	var appender app.DiseaseSymptomForm
	if err := context.ShouldBindJSON(&appender); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(appender); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	var symptom models.Symptom
	if err := database.Instance.Where("id = ?", appender.SymptomID).First(&symptom).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Gejala tidak ditemukan", "status": "error"})
		return
	}

	var disease models.Disease
	if err := database.Instance.Where("id = ?", appender.DiseaseID).First(&disease).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Penyakit tidak ditemukan", "status": "error"})
		return
	}

	if err := database.Instance.Model(&disease).Association("DiseaseSymptom").Append(&symptom); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	message := fmt.Sprintf("Berhasil menambahkan gejala %s dalam Penyakit %s", symptom.SymptomName, disease.DiseaseName)
	context.JSON(http.StatusOK, gin.H{"message": message, "status": "success"})
}

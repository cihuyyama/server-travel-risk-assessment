package router

import (
	"travel-risk-assessment/controllers"
	"travel-risk-assessment/database"
	"travel-risk-assessment/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	{
		public.POST("/users/login", controllers.Login)
		public.POST("/users/register", controllers.CreateUser)
		public.GET("/migrate", database.ManualMigrate)
	}

	protected := r.Group("/api")
	protected.Use(middlewares.Authenticate())
	{
		protected.GET("/users", controllers.GetUserByToken)
		protected.PUT("/users", controllers.UpdateUser)
		protected.DELETE("/users", controllers.DeleteUser)
		protected.GET("/photos", controllers.GetAllPhotos)
		protected.GET("/photos/:id", controllers.GetPhotoByID)
		protected.POST("/photos", controllers.CreatePhoto)
		protected.PUT("/photos/:id", controllers.UpdatePhoto)
		protected.DELETE("/photos/:id", controllers.DeletePhoto)
	}

	protectedSymptom := r.Group("/api")
	protectedSymptom.Use(middlewares.Authenticate())
	{
		protectedSymptom.POST("/symptoms", controllers.CreateSymptom)
		protectedSymptom.GET("/symptoms", controllers.GetAllSymptoms)
		protectedSymptom.GET("/symptoms/:id", controllers.GetSymptomByID)
		protectedSymptom.PUT("/symptoms/:id", controllers.UpdateSymptom)
		protectedSymptom.DELETE("/symptoms/:id", controllers.DeleteSymptom)
	}

	protectedTravel := r.Group("/api")
	protectedTravel.Use(middlewares.Authenticate())
	{
		protectedTravel.POST("/travels", controllers.CreateTravelHistory)
		protectedTravel.GET("/travels", controllers.GetAllHistory)
		protectedTravel.GET("/travels/:id", controllers.GetTravelHistoryByID)
		protectedTravel.PUT("/travels/:id", controllers.UpdateTravelhistory)
		protectedTravel.DELETE("/travels/:id", controllers.DeleteTravelHistory)
	}

	protectedMedical := r.Group("/api")
	protectedMedical.Use(middlewares.Authenticate())
	{
		protectedMedical.POST("/medicals", controllers.CreateMedicalHistory)
		protectedMedical.GET("/medicals", controllers.GetAllMedicalHistory)
		protectedMedical.GET("/medicals/", controllers.GetMedicalHistoryByID)
		protectedMedical.PUT("/medicals/:id", controllers.UpdateMedicalHistory)
		protectedMedical.DELETE("/medicals/:id", controllers.DeleteMedicalHistory)
	}

	protectedDisease := r.Group("/api")
	protectedDisease.Use(middlewares.Authenticate())
	{
		protectedDisease.POST("/diseases", controllers.CreateDisease)
		protectedDisease.GET("/diseases", controllers.GetAllDiseases)
		protectedDisease.GET("/diseases/:id", controllers.GetDiseaseByID)
		protectedDisease.PUT("/diseases/:id", controllers.UpdateDiseaseByID)
		protectedDisease.DELETE("/diseases/:id", controllers.DeleteDiseaseByID)
		protectedDisease.POST("/diseases/append", controllers.AppendDisease)

	}

	protectedEndemic := r.Group("/api")
	protectedEndemic.Use(middlewares.Authenticate())
	{
		protectedEndemic.POST("/endemics", controllers.CreateEndemic)
		protectedEndemic.GET("/endemics", controllers.GetAllEndemics)
		protectedEndemic.GET("/endemics/:id", controllers.GetEndemicByID)
		protectedEndemic.PUT("/endemics/:id", controllers.UpdateEndemic)
		protectedEndemic.DELETE("/endemics/:id", controllers.DeleteEndemic)
	}

	protectedTreatment := r.Group("/api")
	protectedTreatment.Use(middlewares.Authenticate())
	{
		protectedTreatment.POST("/treatments", controllers.CreateTreatment)
		protectedTreatment.GET("/treatments", controllers.GetAllTreatments)
		protectedTreatment.GET("/treatments/:id", controllers.GetTreatmentByID)
		protectedTreatment.PUT("/treatments/:id", controllers.UpdateTreatmentByID)
		protectedTreatment.DELETE("/treatments/:id", controllers.DeleteTreatmentByID)
	}

	protectedPrevention := r.Group("/api")
	protectedPrevention.Use(middlewares.Authenticate())
	{
		protectedPrevention.POST("/preventions", controllers.CreatePrevention)
		protectedPrevention.GET("/preventions", controllers.GetAllPreventions)
		protectedPrevention.GET("/preventions/:id", controllers.GetPreventionByID)
		protectedPrevention.PUT("/preventions/:id", controllers.UpdatePreventionByID)
		protectedPrevention.DELETE("/preventions/:id", controllers.DeletePreventionByID)
	}

	protectedPreTravel := r.Group("/api")
	protectedPreTravel.Use(middlewares.Authenticate())
	{
		protectedPreTravel.GET("/pretravel/:id", controllers.GetPreTravelProps)
	}

	return r
}

package router

import (
	"travel-risk-assessment/controllers"
	"travel-risk-assessment/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	{
		public.POST("/users/login", controllers.Login)
		public.POST("/users/register", controllers.CreateUser)
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

	return r
}

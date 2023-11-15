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
		public.POST("/symptoms", controllers.CreateSymptom)
		public.GET("/symptoms", controllers.GetAllSymptoms)
		public.GET("/symptoms/:id", controllers.GetSymptomByID)
		public.PUT("/symptoms/:id", controllers.UpdateSymptom)
		public.DELETE("/symptoms/:id", controllers.DeleteSymptom)
	}

	protected := r.Group("/api")
	protected.Use(middlewares.Authenticate())
	{
		protected.GET("/users/:id", controllers.GetUserByID)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)
		protected.GET("/photos", controllers.GetAllPhotos)
		protected.GET("/photos/:id", controllers.GetPhotoByID)
		protected.POST("/photos", controllers.CreatePhoto)
		protected.PUT("/photos/:id", controllers.UpdatePhoto)
		protected.DELETE("/photos/:id", controllers.DeletePhoto)
	}

	return r
}

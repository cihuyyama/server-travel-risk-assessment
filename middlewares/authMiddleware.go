package middlewares

import (
	"net/http"
	"strings"
	"travel-risk-assessment/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			context.JSON(http.StatusBadRequest, gin.H{"message": "No bearer", "status": "error"})
			context.Abort()
			return
		}

		tokenString = parts[1]

		if tokenString == "" {
			context.JSON(401, gin.H{"message": "request does not contain an access token", "status": "error"})
			context.Abort()
			return
		}
		err := helpers.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"message": err.Error(), "status": "error"})
			context.Abort()
			return
		}
		context.Next()
	}
}

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

func CreateUser(context *gin.Context) {
	var userFormRegister app.UserFormRegister
	if err := context.ShouldBindJSON(&userFormRegister); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		context.Abort()
		return
	}

	if _, err := govalidator.ValidateStruct(userFormRegister); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}
	var user models.User

	if len(userFormRegister.Password) < 6 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Password minimal 6 karakter", "status": "error"})
		context.Abort()
		return
	}

	if err := database.Instance.Where("email = ?", userFormRegister.Email).First(&user).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists", "status": "error"})
		context.Abort()
		return
	}

	if err := database.Instance.Where("username = ?", userFormRegister.Username).First(&user).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Username sudah terdaftar", "status": "error"})
		context.Abort()
		return
	}

	user = models.User{
		Username: userFormRegister.Username,
		Email:    userFormRegister.Email,
		Password: userFormRegister.Password,
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": record.Error.Error(), "status": "error"})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Berhasil Membuat Akun", "status": "success"})
}

func Login(context *gin.Context) {
	var userFormLogin app.UserFormLogin
	if err := context.ShouldBindJSON(&userFormLogin); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(userFormLogin); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	var user models.User
	if err := database.Instance.Where("email = ?", userFormLogin.Email).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email atau password salah", "status": "error"})
		return
	}

	if err := user.CheckPassword(userFormLogin.Password); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email atau password salah", "status": "error"})
		return
	}

	token, err := helpers.GenerateJWT(user.ID, user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating token", "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "token": token, "status": "success"})
}

func GetUserByToken(context *gin.Context) {
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

	var userResult app.UserResult
	userResult.ID = user.ID
	userResult.Username = user.Username
	userResult.Email = user.Email
	userResult.CreatedAt = user.CreatedAt.String()
	userResult.UpdatedAt = user.UpdatedAt.String()

	context.JSON(http.StatusOK, gin.H{"data": userResult, "status": "success"})
}

func UpdateUser(context *gin.Context) {
	var userFormUpdate app.UserFormUpdate
	if err := context.ShouldBindJSON(&userFormUpdate); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	if _, err := govalidator.ValidateStruct(userFormUpdate); err != nil {
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

	var user models.User
	if len(userFormUpdate.Password) < 6 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Password minimal 6 karakter", "status": "error"})
		context.Abort()
		return
	}

	if err := database.Instance.Where("email = ? AND id != ?", userFormUpdate.Email, claims.ID).First(&user).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email sudah terdaftar", "status": "error"})
		context.Abort()
		return
	}

	if err := database.Instance.Where("username = ? AND id != ?", userFormUpdate.Username, claims.ID).First(&user).Error; err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Username sudah terdaftar", "status": "error"})
		context.Abort()
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		context.Abort()
		return
	}

	if err := database.Instance.Where("id = ?", claims.ID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Pengguna tidak ditemukan", "status": "error"})
		return
	}

	if err := database.Instance.First(&user, claims.ID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Pengguna tidak ditemukan", "status": "error"})
		return
	}

	user.Username = userFormUpdate.Username
	user.Email = userFormUpdate.Email
	if userFormUpdate.Password != "" {
		if err := user.HashPassword(userFormUpdate.Password); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
			context.Abort()
			return
		}
	}

	if err := database.Instance.Save(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Berhasil mengupdate pengguna", "status": "success"})
}

func DeleteUser(context *gin.Context) {
	var user models.User

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

	if err := database.Instance.Where("id = ?", claims.ID).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Pengguna tidak ditemukan", "status": "error"})
		return
	}

	if err := database.Instance.First(&user, claims.ID).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Pengguna tidak ditemukan", "status": "error"})
		return
	}

	if err := database.Instance.Delete(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Berhasil menghapus pengguna", "status": "success"})
}

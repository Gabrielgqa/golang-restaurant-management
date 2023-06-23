package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func Signup() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	return true, ""
}

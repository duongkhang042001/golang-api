package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {

	authGroup := router.Group("/auth")

	// Tuyến đường đăng nhập
	authGroup.GET("/login", LoginHandler)

	// Tuyến đường đăng ký
	authGroup.GET("/register", RegisterHandler)
}

func LoginHandler(ctx *gin.Context) {
	// Xử lý đăng nhập

	// Trả về kết quả trong đối tượng JSON
}

func RegisterHandler(ctx *gin.Context) {
	// Xử lý đăng ký
	// Trả về kết quả trong đối tượng JSON
}

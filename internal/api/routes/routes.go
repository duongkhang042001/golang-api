package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(app *gin.Engine, db *gorm.DB) {
	RegisterAuthRoutes(app)
	RegisterUserRoutes(app)
}

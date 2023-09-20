package main

import (
	"core-api/internal/api/configs"
	"core-api/internal/api/middlewares"
	"core-api/internal/api/routes"

	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var cors = middlewares.CorsMiddleware()
var request = middlewares.RequestIdMiddleware()

func init() {
	db = configs.SetupDatabaseConnection()
}

func main() {
	defer configs.CloseDatabaseConnection(db)

	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

	if os.Getenv("APP_ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	if os.Getenv("APP_ENV") == "DEBUG" {
		gin.SetMode(gin.DebugMode)
	}

	if os.Getenv("APP_ENV") == "TEST" {
		gin.SetMode(gin.TestMode)
	}

	app := gin.Default()

	app.ForwardedByClientIP = true

	app.SetTrustedProxies([]string{"127.0.0.1"})

	app.Use(cors)

	app.Use(request)

	routes.SetupRoutes(app)

	app.Run()
}

// https://github.com/araujo88/golang-rest-api-template/

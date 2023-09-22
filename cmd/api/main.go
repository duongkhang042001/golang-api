package main

import (
	"core-api/internal/api/routes"
	"core-api/pkg/database"
	"core-api/pkg/middlewares"

	"os"

	"github.com/gin-gonic/gin"
)

var cors = middlewares.CorsMiddleware()
var request = middlewares.RequestIdMiddleware()

func main() {

	database.SetupDatabaseConnection()

	defer database.CloseDatabaseConnection()

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

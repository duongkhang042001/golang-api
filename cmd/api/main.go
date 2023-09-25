package main

import (
	"core-api/internal/api/routes"
	"core-api/pkg/database"
	"core-api/pkg/middlewares"
	"core-api/pkg/redis"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	redis.InitRedis()

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

	app.Use(middlewares.CorsMiddleware())

	app.Use(middlewares.RequestIdMiddleware())

	// app.Use(middlewares.JwtMiddleware())

	routes.SetupRoutes(app)

	app.Run()
}

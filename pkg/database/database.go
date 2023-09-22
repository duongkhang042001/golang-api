package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

func SetupDatabaseConnection() *gorm.DB {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Failed to load env file!")
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")
	dbPort := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database!")
	}

	database.AutoMigrate()

	connection = database

	return connection
}

func CloseDatabaseConnection() {
	dbSQL, err := connection.DB()
	if err != nil {
		panic("Failed to close connection from database!")
	}
	dbSQL.Close()
}

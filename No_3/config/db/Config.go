package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func InitialMigration() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Errorf("Error loading .env file %v", errEnv.Error())
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPass, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Errorf(err.Error())
		panic("Cannot connect to DB")
	}

	for _, model := range RegisterModels() {
		err = DB.AutoMigrate(model.Model)

		if err != nil {
			log.Errorf(err.Error())
		}
	}

	log.Info("Database migrated successfully.")

}

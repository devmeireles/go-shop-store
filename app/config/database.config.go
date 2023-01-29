package config

import (
	"fmt"
	"log"
	"os"

	"github.com/devmeireles/go-shop-store/app/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	var (
		db  *gorm.DB
		err error
	)

	if os.Getenv("ENVIRONMENT") == "test" {
		db, err = gorm.Open(sqlite.Open("../../db_test.db"), &gorm.Config{})
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		var (
			DbHost     = os.Getenv("DB_HOST")
			DbUser     = os.Getenv("DB_USER")
			DbName     = os.Getenv("DB_NAME")
			DbPassword = os.Getenv("DB_PASSWORD")
			DbPort     = os.Getenv("DB_PORT")
		)

		connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable port=5432 password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  connectionString,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("This is the error:", err)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(
		models.Product{},
	)

	DB = Dbinstance{
		Db: db,
	}
}

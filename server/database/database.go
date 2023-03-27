package database

import (
	"log"

	"github.com/whicencer/golang-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func Connect() {
	dbURL := "postgres://postgres:denielpuffmeister@127.0.0.1:5432/postgres"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("Failed to connect to database, \n", err)
	}
	log.Println("Connected")

	err = db.AutoMigrate(&models.Book{}, &models.User{})

	DB = DBInstance{
		Db: db,
	}
}
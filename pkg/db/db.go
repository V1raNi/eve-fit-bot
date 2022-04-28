package db

import (
	"eve-fit-bot/pkg/models"
	"eve-fit-bot/pkg/util"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(config util.Config) *gorm.DB {
	dbURL := config.ConnectionString

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Fit{})

	return db
}

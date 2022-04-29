package db

import (
	"eve-fit-bot/pkg/models"
	"eve-fit-bot/pkg/util"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(config util.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Fit{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

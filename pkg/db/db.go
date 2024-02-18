package db

import (
	"log"

	"github.com/14jasimmtp/auth-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Connection(url string) Handler {
	db, err := gorm.Open(postgres.Open(url),&gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return Handler{db}
}

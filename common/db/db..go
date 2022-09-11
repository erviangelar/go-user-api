package db

import (
	"log"

	"github.com/erviangelar/go-user-api/common/config"
	"github.com/erviangelar/go-user-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var users = []models.User{
	{Name: "user", Username: "user", Role: "user", Password: "$2a$14$pGJf5uGp6F8jTkhYspfoUe4hAGfDgGfbz99KFwJ9Xv8JKtVE1eXpO"},
	{Name: "admin", Username: "admin", Role: "admin", Password: "$2a$14$AFxmldcPxurFLsay/fNtE.NPXUmgregh1VsNUHBmbuLe1m3wby9pO"},
}

func Init(config *config.Configurations) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBConn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	db.Create(&users)
	return db
}

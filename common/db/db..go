package db

import (
	"log"

	utilCache "github.com/erviangelar/go-user-api/common/cache"
	"github.com/erviangelar/go-user-api/common/config"
	cache "github.com/erviangelar/go-user-api/common/helper"
	"github.com/erviangelar/go-user-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbs struct {
	DB    *gorm.DB
	Cache *utilCache.RedisClient
}

var users = []models.User{
	{Name: "user", Username: "user", Role: []string{"user"}, Password: "$2a$14$pGJf5uGp6F8jTkhYspfoUe4hAGfDgGfbz99KFwJ9Xv8JKtVE1eXpO"},
	{Name: "admin", Username: "admin", Role: []string{"admin"}, Password: "$2a$14$AFxmldcPxurFLsay/fNtE.NPXUmgregh1VsNUHBmbuLe1m3wby9pO"},
}

func Init(config *config.Configurations) (*Dbs, error) {
	db, err := gorm.Open(postgres.Open(config.DBConn), &gorm.Config{})
	redisCache, err := cache.New(config)

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	db.Create(&users)
	return &Dbs{
		DB:    db,
		Cache: redisCache,
	}, nil
}

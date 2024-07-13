package database

import (
	"context"
	"fmt"
	"log"

	"github.com/takumi616/ielts-vocabularies-api/infrastructures"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	Db *gorm.DB
}

type Test struct {
	gorm.Model
	Message string
}

func Open(ctx context.Context, pgConfig *infrastructures.PgConfig) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pgConfig.Host, pgConfig.Port, pgConfig.User, pgConfig.PassWord, pgConfig.DbName, pgConfig.Sslmode)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to open postgresql: %v", err)
	}

	db.AutoMigrate(&Test{})
	db.Create(
		&Test{
			Message: "test record",
		},
	)
}

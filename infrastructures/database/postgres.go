package database

import (
	"context"
	"fmt"
	"log"

	"github.com/takumi616/go-restapi/infrastructures"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	Db *gorm.DB
}

func Open(ctx context.Context, pgConfig *infrastructures.PgConfig) (*Gorm, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pgConfig.Host, pgConfig.Port, pgConfig.User, pgConfig.PassWord, pgConfig.DbName, pgConfig.Sslmode)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to open postgresql: %v", err)
		return nil, err
	}

	db.AutoMigrate(&Vocabulary{})
	return &Gorm{Db: db}, nil
}

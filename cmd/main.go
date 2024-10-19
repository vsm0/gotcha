package main

import (
	"github.com/vsm0/gotcha/model"

	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	file := "gacha.db"
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}

	// create or migrate tables
	err = db.AutoMigrate(
		&model.Account{},
		&model.InventoryItem{},
		&model.Item{},
		&model.Reward{},
		&model.Event{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
}

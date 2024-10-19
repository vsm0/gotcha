package main

import (
	"github.com/vsm0/gotcha/model"
	"github.com/vsm0/gotcha/state"

	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	file := "gacha.db"
	dial := sqlite.Open(file)
	conf := gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db, err := gorm.Open(dial, &conf)
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

	if err := state.NewApp(db).Run(); err != nil {
		panic(err)
	}
}

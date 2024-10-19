package test

import (
	"github.com/vsm0/gotcha/model"

	"crypto/sha256"
	"math/rand"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var rng = rand.New(rand.NewSource(1))

func getDb(t *testing.T) *gorm.DB {
	file := "gacha.db"
	dial := sqlite.Open(file)
	conf := gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db, err := gorm.Open(dial, &conf)
	if err != nil {
		t.Fatalf("DB Connection failure: %v", err)
	}

	return db
}

func migrate(t *testing.T, db *gorm.DB) {
	// create or migrate tables
	err := db.AutoMigrate(
		&model.Account{},
		&model.InventoryItem{},
		&model.Item{},
		&model.Reward{},
		&model.Event{},
	)
	if err != nil {
		t.Fatalf("Migration failure: %v", err)
	}
}

func uid() uint {
	return uint(rng.Uint32())
}

func hash(s string) string {
	sha := sha256.New()
	sha.Write([]byte(s))
	return string(sha.Sum(nil))
}

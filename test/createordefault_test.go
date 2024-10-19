package test

import (
	"github.com/vsm0/gotcha/model"

	"math/rand"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestOpenOrDefault(t *testing.T) {
	file := "gacha.db"
	dial := sqlite.Open(file)
	conf := gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db, err := gorm.Open(dial, &conf)
	if err != nil {
		t.Fatalf("DB Connection failure: %v", err)
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
		t.Fatalf("Migration failure: %v", err)
	}

	src := rand.NewSource(1)
	rng := rand.New(src)
	id := func() uint {
		return uint(rng.Uint32())
	}

	accounts := []*model.Account{
		{Id: id()},
		{Id: id()},
		{Id: id()},
		{Id: id()},
	}

	for _, v := range accounts {
		t.Logf("Create Account: %v", v.Id)
	}

	if res := db.Create(accounts); res.Error != nil {
		t.Fatalf("Creation failure: %v", res.Error)
	}

	items := []*model.Item{
		{Id: id(), Name: "Primagen"},
		{Id: id(), Name: "Magic Dust"},
		{Id: id(), Name: "Hard Bow"},
		{Id: id(), Name: "Macca"},
	}

	for _, v := range items {
		t.Logf("Create Item: %v", v.Name)
	}

	if res := db.Create(items); res.Error != nil {
		t.Fatalf("Creation failure: %v", res.Error)
	}

	inventoryItems := []*model.InventoryItem{
		{AccountId: accounts[0].Id, ItemId: items[0].Id, ItemCount: 160},
		{AccountId: accounts[0].Id, ItemId: items[1].Id, ItemCount: 420},
		{AccountId: accounts[1].Id, ItemId: items[2].Id, ItemCount: 3},
	}

	for _, v := range inventoryItems {
		t.Logf("Create Inventory Item: %v:%v", v.AccountId, v.ItemId)
	}

	if res := db.Create(inventoryItems); res.Error != nil {
		t.Fatalf("Creation failure: %v", res.Error)
	}

	rewards := []*model.Reward{
		{Id: id(), ItemId: items[3].Id, ItemCountRequired: 20, DropId: items[2].Id, DropRate: 70, DropCount: 1},
		{Id: id(), ItemId: items[1].Id, ItemCountRequired: 50, DropId: items[0].Id, DropRate: 85, DropCount: 10},
	}

	for _, v := range rewards {
		t.Logf("Create Reward: %v", v.Id)
	}

	if res := db.Create(rewards); res.Error != nil {
		t.Fatalf("Creation failure: %v", res.Error)
	}

	events := []*model.Event{
		{Id: id(), Name: "Hard Bow Event", Start: time.Now(), End: time.Now(), RewardId: rewards[0].Id},
		{Id: id(), Name: "Primagen Event", Start: time.Now(), End: time.Now(), RewardId: rewards[1].Id},
	}

	for _, v := range events {
		t.Logf("Create Event: %v", v.Name)
	}

	if res := db.Create(events); res.Error != nil {
		t.Fatalf("Creation failure: %v", res.Error)
	}
}

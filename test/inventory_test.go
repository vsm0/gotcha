package test

import (
	"github.com/vsm0/gotcha/model"

	"testing"

	"gorm.io/gorm"
)

func getAccountItem(t *testing.T) (*gorm.DB, *model.Account, *model.Item) {
	db := getDb(t)
	migrate(t, db)

	a := &model.Account{
		Username: "Anon",
	}
	db.First(a)

	i := &model.Item{
		Name: "Macca",
	}
	db.First(i)

	return db, a, i
}

func TestCreateInventoryItem(t *testing.T) {
	TestCreateAccount(t)
	TestCreateItem(t)

	db, a, i := getAccountItem(t)

	item := &model.InventoryItem{
		AccountId: a.Id,
		ItemId: i.Id,
		ItemCount: 5,
	}

	if res := db.Create(item); res.Error != nil {
		t.Logf("%v", res.Error)
	}
}

func TestUpdateInventoryItem(t *testing.T) {
	db, a, i := getAccountItem(t)

	item := &model.InventoryItem{
		AccountId: a.Id,
		ItemId: i.Id,
	}
	db.First(item)

	// simulate removing all items
	item.ItemCount = 0

	// delete item if empty
	if item.ItemCount == 0 {
		// dont actually delete, so the next test passes
	}
}

func TestDeleteInventoryItem(t *testing.T) {
	db, a, i := getAccountItem(t)

	if res := db.Delete(&model.InventoryItem{}, "account_id LIKE ? AND item_id LIKE ?", a.Id, i.Id); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}

	// cleanup for other tests; should not run in actual code
	TestDeleteAccount(t)
	TestDeleteItem(t)
}

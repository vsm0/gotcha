package test

import (
	"github.com/vsm0/gotcha/model"

	"testing"
)

func TestDropDb(t *testing.T) {
	db := getDb(t)

	tables := []interface{}{
		&model.Account{},
		&model.Item{},
		&model.InventoryItem{},
		&model.Reward{},
		&model.Event{},
	}

	if err := db.Migrator().DropTable(tables...); err != nil {
		t.Fatalf("%v", err)
	}
}

package test

import (
	"github.com/vsm0/gotcha/model"

	"testing"
)

func TestCreateItem(t *testing.T) {
	db := getDb(t)
	migrate(t, db)

	item := &model.Item{
		Name: "Macca",
	}

	if res := db.Create(item); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
}

func TestQueryItem(t *testing.T) {
	db := getDb(t)
	migrate(t, db)

	item := &model.Item{
		Name: "Macca",
	}

	if res := db.First(item); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
}

func TestDeleteItem(t *testing.T) {
	db := getDb(t)
	migrate(t, db)

	if res := db.Delete(&model.Item{}, "name LIKE ?", "Macca"); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
}

package test

import (
	"github.com/vsm0/gotcha/model"

	"testing"
)

func TestCreateReward(t *testing.T) {
	// create items
	db := getDb(t)
	migrate(t, db)

	items := []model.Item{
		{Name: "Macca"},
		{Name: "Human Organs"},
	}

	if res := db.Create(items); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}

	// create reward
	reward := &model.Reward{
		ItemId: items[0].Id,
		ItemCountRequired: 5,
	}
	
	if res := db.Create(reward); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
} 

func TestDeleteReward(t *testing.T) {
	// delete reward
	db := getDb(t)
	migrate(t, db)

	item := &model.Item{
		Name: "Macca",
	}
	db.First(item)

	if res := db.Delete(&model.Reward{}, "item_id LIKE ?", item.Id); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}

	// delete test items
	if res := db.Delete(&model.Item{}, "name LIKE ?", "Macca"); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
	if res := db.Delete(&model.Item{}, "name LIKE ?", "Human Organs"); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
}

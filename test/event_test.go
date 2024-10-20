package test

import (
	"github.com/vsm0/gotcha/model"

	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {
	TestCreateReward(t)

	db := getDb(t)
	migrate(t, db)

	item := &model.Item{
		Name: "Macca",
	}
	db.First(item)

	reward := &model.Reward{
		ItemId: item.Id,
	}
	db.First(reward)

	now := time.Now()

	event := &model.Event{
		Name: "Human Trafficking Event",
		Start: now,
		End: now.Add(time.Hour * 24 * 10), // ends in ten days
		RewardId: reward.Id,
	}

	if res := db.Create(event); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
}

func TestDeleteEvent(t *testing.T) {
	TestDeleteReward(t)

	db := getDb(t)
	migrate(t, db)

	event := &model.Event{
		Name: "Human Trafficking Event",
	}
	db.First(event)

	if res := db.Delete(event); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
}

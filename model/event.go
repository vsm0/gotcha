package model

import "time"

type Event struct {
	Id uint `gorm:"primaryKey"`
	Name string
	Start time.Time
	End time.Time
	RewardId uint `gorm:"foreignKey:RewardId"`
}

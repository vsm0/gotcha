package model

type Item struct {
	Id uint `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

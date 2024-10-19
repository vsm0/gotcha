package model

type Account struct {
	Id uint `gorm:"primaryKey"` // autoincrements by default
	Username string `gorm:"unique"`
	Password string
}

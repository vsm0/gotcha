package model

type Reward struct {
	Id uint `gorm:"primaryKey"`

	// item requirements
	ItemId uint `gorm:"foreignKey:ItemId"`
	ItemCountRequired uint

	// item drops
	DropId uint `gorm:"foreignKey:ItemId"`
	DropRate uint
	DropCount uint
}

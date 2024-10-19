package model

type InventoryItem struct {
	// account and item id form a composite unique constraint
	AccountId uint `gorm:"foreignKey:AccountId;index:inventory_id,unique"`
	ItemId uint `gorm:"foreignKey:ItemId;index:inventory_id,unique"`
	ItemCount uint
}

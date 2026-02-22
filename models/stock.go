package models

type Stock struct {
	ID      uint   `gorm:"primaryKey"`
	Quantit uint64 `gorm:"not null"`
	ItemId  uint   `gorm:"not null"`
}

func (Stock) TableName() string {
	return "stocks"
}

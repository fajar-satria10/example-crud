package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model          // adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

func (Item) TableName() string {
	return "items"
}

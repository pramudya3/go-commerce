package domain

import (
	"time"
)

type (
	Product struct {
		ID         uint64       `json:"id" gorm:"primary_key;index"`
		CategoryID uint64       `json:"category_id"`
		Name       string       `json:"name"`
		Desciption string       `json:"description"`
		Stock      int          `json:"stock"`
		Price      int          `json:"price"`
		CreatedAt  time.Time    `json:"created_at"`
		UpdatedAt  *time.Time   `json:"updated_at"`
		CartDetail []CartDetail `gorm:"foreignKey:ProductID"`
	}

	CreateProductRequest struct {
		CategoryID  uint64 `json:"category_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Stock       int    `json:"stock"`
		Price       int    `json:"price"`
	}

	ProductResponse struct {
		ID         uint64     `json:"id" gorm:"primary_key;index"`
		CategoryID uint64     `json:"category_id"`
		Name       string     `json:"name"`
		Desciption string     `json:"description"`
		Stock      int        `json:"stock"`
		Price      int        `json:"price"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  *time.Time `json:"updated_at"`
	}
)

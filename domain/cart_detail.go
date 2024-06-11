package domain

import "time"

type (
	CartDetail struct {
		ID        uint64     `json:"id" gorm:"primary_key"`
		CartID    uint64     `json:"cart_id"`
		ProductID uint64     `json:"product_id"`
		Quantity  int        `json:"quantity"`
		Subtotal  int        `json:"subtotal"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	CreateCartDetail struct {
		ProductID uint64 `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}

	UpdateCartDetail struct {
		ID        uint64    `json:"id"`
		CartID    uint64    `json:"cart_id"`
		ProductID *uint64   `json:"product_id"`
		Quantity  int       `json:"quantity"`
		Subtotal  *int      `json:"subtotal"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	DeleteCartDetail struct {
		ID uint64 `json:"id"`
	}
)

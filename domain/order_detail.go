package domain

import "time"

type (
	OrderDetail struct {
		ID        uint64     `json:"id" gorm:"primary_key"`
		OrderID   uint64     `json:"order_id"`
		ProductID uint64     `json:"product_id"`
		Quantity  int        `json:"quantity"`
		Subtotal  int        `json:"subtotal"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	CreateOrderDetail struct {
		ProductID uint64 `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}
)

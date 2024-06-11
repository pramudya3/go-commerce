package domain

import "time"

type (
	Order struct {
		ID              uint64     `json:"id" gorm:"primary_key"`
		UserID          uint64     `json:"user_id"`
		PaymentID       uint64     `json:"payment_id"`
		TotalPrice      int        `json:"total_price"`
		Status          string     `json:"status"`
		ShippingAddress string     `json:"shipping_address"`
		CreatedAt       time.Time  `json:"created_at"`
		UpdatedAt       *time.Time `json:"updated_at"`
	}

	CreateOrder struct {
		PaymentID       uint64 `json:"payment_id"`
		ShippingAddress string `json:"shipping_address"`
	}
)

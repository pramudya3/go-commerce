package domain

import (
	"time"
)

const (
	StatusPaymentPending      = "Pending"
	StatusPaymentSuccessfully = "Successfully"
	StatusPaymentFailed       = "Failed"
)

type (
	Payment struct {
		ID            uint64     `json:"id" gorm:"primaryKey"`
		UserID        uint64     `json:"user_id"`
		CartID        uint64     `json:"cart_id"`
		PaymentMethod string     `json:"payment_method"`
		Status        string     `json:"status"`
		TotalPrice    int        `json:"total_price"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     *time.Time `json:"updated_at"`

		Cart []Cart `gorm:"foreignKey:PaymentID"`
	}

	PaymentRequest struct {
		PaymentMethod string `json:"payment_method"`
	}

	ResponsePayment struct {
		ID            uint64     `json:"id" gorm:"primaryKey"`
		UserID        uint64     `json:"user_id"`
		CartID        uint64     `json:"cart_id"`
		PaymentMethod string     `json:"payment_method"`
		Status        string     `json:"status"`
		TotalPrice    int        `json:"total_price"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     *time.Time `json:"updated_at"`
	}
)

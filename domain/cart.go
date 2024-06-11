package domain

import "time"

const (
	StatusOnCart    = "On Cart"
	StatusOnOrder   = "On Order"
	StatusOnPayment = "On Payment"
)

// /my-cart
// /my-order

type (
	Cart struct {
		ID              uint64     `json:"id" gorm:"primary_key"`
		UserID          uint64     `json:"user_id"`
		PaymentID       *uint64    `json:"payment_id"`
		Status          string     `json:"status"`
		ShippingAddress *string    `json:"shipping_address"`
		TotalPrice      int        `json:"total_price"`
		CreatedAt       time.Time  `json:"created_at"`
		UpdatedAt       *time.Time `json:"updated_at"`

		CartDetail []CartDetail `gorm:"foreignKey:CartID"`
	}

	CartUpdateRequest struct {
		UserID          uint64     `json:"user_id"`
		PaymentID       *uint64    `json:"payment_id"`
		Status          *string    `json:"status"`
		ShippingAddress string     `json:"shipping_address"`
		TotalPrice      int        `json:"total_price"`
		UpdatedAt       *time.Time `json:"updated_at"`
	}

	CartResponse struct {
		ID         uint64     `json:"id" gorm:"primary_key"`
		UserID     uint64     `json:"user_id"`
		PaymentID  uint64     `json:"payment_id"`
		Status     string     `json:"status"`
		TotalPrice int        `json:"total_price"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  *time.Time `json:"updated_at"`
	}

	CartResponseWithDetails struct {
		Cart        *CartResponse `json:"cart"`
		CartDetails []*CartDetail `json:"cart_details"`
	}

	CheckoutRequest struct {
		ShippingAddress string   `json:"shipping_address"`
		CartDetailIDs   []uint64 `json:"cart_detail_ids"`
	}
)

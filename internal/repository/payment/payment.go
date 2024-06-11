package payment

import (
	"context"
	"go-commerce/domain"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) error
	UpdatePayment(ctx context.Context, payment *domain.Payment) error
	FetchPayment(ctx context.Context, userID uint64) ([]*domain.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func (p *paymentRepository) FetchPayment(ctx context.Context, userID uint64) ([]*domain.Payment, error) {
	payments := []*domain.Payment{}
	if err := p.db.Where("user_id = ?", userID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (p *paymentRepository) CreatePayment(ctx context.Context, payment *domain.Payment) error {
	if err := p.db.Create(&payment).Error; err != nil {
		return err
	}
	return nil
}

func (p *paymentRepository) UpdatePayment(ctx context.Context, payment *domain.Payment) error {
	if err := p.db.Where("id = ?", payment.ID).Save(&payment).Error; err != nil {
		return err
	}
	return nil
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

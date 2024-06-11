package payment

import (
	"context"
	"go-commerce/domain"
	"go-commerce/internal/repository/payment"
	"go-commerce/internal/usecase/cart"
	"go-commerce/pkg/config"
	"go-commerce/pkg/utils"
	"time"
)

type PaymentUsecase interface {
	CreatePayment(ctx context.Context, userID uint64, payment *domain.PaymentRequest) (*domain.ResponsePayment, error)
	UpdatePayment(ctx context.Context, payment *domain.Payment) error
	FetchPayment(ctx context.Context, userID uint64) ([]*domain.ResponsePayment, error)
}

type paymentUsecase struct {
	cartUc      cart.CartUsecase
	paymentRepo payment.PaymentRepository
	timeout     time.Duration
}

func (p *paymentUsecase) FetchPayment(ctx context.Context, userID uint64) ([]*domain.ResponsePayment, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	payments, err := p.paymentRepo.FetchPayment(ctx, userID)
	if err != nil {
		return nil, err
	}

	data := []*domain.ResponsePayment{}
	utils.CopyJsonStruct(payments, &data)

	return data, nil
}

func (p *paymentUsecase) CreatePayment(ctx context.Context, userID uint64, payment *domain.PaymentRequest) (*domain.ResponsePayment, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	myOrder, err := p.cartUc.MyOrder(ctx, userID)
	if err != nil {
		return nil, err
	}

	newPayment := &domain.Payment{
		UserID:        userID,
		CartID:        myOrder.Cart.ID,
		PaymentMethod: payment.PaymentMethod,
		Status:        domain.StatusPaymentPending,
		TotalPrice:    myOrder.Cart.TotalPrice,
		CreatedAt:     time.Now(),
	}

	if err := p.paymentRepo.CreatePayment(ctx, newPayment); err != nil {
		return nil, err
	}

	status := domain.StatusOnPayment
	now := time.Now()
	if err := p.cartUc.UpdateCartByCartID(ctx, myOrder.Cart.ID, &domain.CartUpdateRequest{
		UserID:     myOrder.Cart.UserID,
		PaymentID:  &newPayment.ID,
		Status:     &status,
		TotalPrice: newPayment.TotalPrice,
		UpdatedAt:  &now,
	}); err != nil {
		return nil, err
	}

	return &domain.ResponsePayment{
		ID:            newPayment.ID,
		UserID:        newPayment.UserID,
		CartID:        newPayment.CartID,
		PaymentMethod: newPayment.PaymentMethod,
		Status:        newPayment.Status,
		TotalPrice:    newPayment.TotalPrice,
		CreatedAt:     newPayment.CreatedAt,
		UpdatedAt:     newPayment.UpdatedAt,
	}, nil
}

func (p *paymentUsecase) UpdatePayment(ctx context.Context, payment *domain.Payment) error {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	now := time.Now()
	payment.UpdatedAt = &now

	return p.paymentRepo.UpdatePayment(ctx, payment)
}

func NewPaymentUsecase(cartUc cart.CartUsecase, paymentRepo payment.PaymentRepository, cfg *config.Config) PaymentUsecase {
	return &paymentUsecase{
		cartUc:      cartUc,
		paymentRepo: paymentRepo,
		timeout:     time.Duration(cfg.TimeoutCtx) * time.Second,
	}
}

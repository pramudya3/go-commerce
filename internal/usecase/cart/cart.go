package cart

import (
	"context"
	"errors"
	"fmt"
	"go-commerce/domain"
	"go-commerce/internal/repository/cart"
	"go-commerce/internal/usecase/product"
	"go-commerce/pkg/config"
	"go-commerce/pkg/utils"
	"time"

	"gorm.io/gorm"
)

type CartUsecase interface {
	Create(ctx context.Context, cart *domain.Cart) error
	FetchByUserID(ctx context.Context, userID uint64) ([]*domain.CartResponse, error)
	FindByStatusAndUserID(ctx context.Context, status string, userID uint64) (*domain.Cart, error)
	UpdateCartByCartID(ctx context.Context, cartID uint64, cart *domain.CartUpdateRequest) error

	Checkout(ctx context.Context, userID uint64, checkout *domain.CheckoutRequest) error
	MyCart(ctx context.Context, userID uint64) (*domain.CartResponseWithDetails, error)
	MyOrder(ctx context.Context, userID uint64) (*domain.CartResponseWithDetails, error)

	AddToCart(ctx context.Context, userID uint64, cartDetail *domain.CreateCartDetail) error
	Delete(ctx context.Context, cartDetailID uint64) error
	FetchByCartID(ctx context.Context, cartID uint64) ([]*domain.CartDetail, error)
	FetchCartDetailByIDs(ctx context.Context, cartDetailIDs any) ([]*domain.CartDetail, error)
	UpdateCartIDByIDs(ctx context.Context, cartDetailIDs any, cartDetail *domain.CartDetail) error
	FindByID(ctx context.Context, cartDetailID uint64) (*domain.CartDetail, error)
}

type cartUsecase struct {
	cartRepo  cart.CartRepository
	timeout   time.Duration
	productUc product.ProductUsecase
}

func (c *cartUsecase) MyOrder(ctx context.Context, userID uint64) (*domain.CartResponseWithDetails, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	cart, err := c.cartRepo.FindByStatusAndUserID(ctx, domain.StatusOnOrder, userID)
	if err != nil {
		return nil, err
	}

	cartDetails, err := c.cartRepo.FetchByCartID(ctx, cart.ID)
	if err != nil {
		return nil, err
	}

	data := &domain.CartResponse{}
	utils.CopyJsonStruct(cart, &data)

	return &domain.CartResponseWithDetails{
		Cart:        data,
		CartDetails: cartDetails,
	}, nil
}

func (c *cartUsecase) MyCart(ctx context.Context, userID uint64) (*domain.CartResponseWithDetails, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	cart, err := c.cartRepo.FindByStatusAndUserID(ctx, domain.StatusOnCart, userID)
	if err != nil {
		return nil, err
	}

	cartDetails, err := c.cartRepo.FetchByCartID(ctx, cart.ID)
	if err != nil {
		return nil, err
	}

	data := &domain.CartResponse{}
	utils.CopyJsonStruct(cart, &data)

	return &domain.CartResponseWithDetails{
		Cart:        data,
		CartDetails: cartDetails,
	}, nil
}

func (c *cartUsecase) Checkout(ctx context.Context, userID uint64, checkout *domain.CheckoutRequest) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	cart, err := c.cartRepo.FindByStatusAndUserID(ctx, domain.StatusOnOrder, userID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			if err = c.cartRepo.Create(ctx, &domain.Cart{
				UserID:    userID,
				Status:    domain.StatusOnOrder,
				CreatedAt: time.Now(),
			}); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	cartDetails, err := c.cartRepo.FetchCartDetailByIDs(ctx, checkout.CartDetailIDs)
	if err != nil {
		return err
	}

	var totalPrice int
	for _, cd := range cartDetails {
		fmt.Println(cd)
		totalPrice += cd.Subtotal
	}

	if err := c.cartRepo.UpdateCartIDByIDs(ctx, checkout.CartDetailIDs, &domain.CartDetail{CartID: cart.ID}); err != nil {
		return err
	}

	cart.TotalPrice = totalPrice
	cart.ShippingAddress = &checkout.ShippingAddress
	now := time.Now()
	cart.UpdatedAt = &now
	if err := c.cartRepo.UpdateByCartID(ctx, cart); err != nil {
		return err
	}

	return nil
}

func (c *cartUsecase) Create(ctx context.Context, cart *domain.Cart) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	cart.CreatedAt = time.Now()

	return c.cartRepo.Create(ctx, cart)
}

func (c *cartUsecase) FetchByUserID(ctx context.Context, userID uint64) ([]*domain.CartResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	carts, err := c.cartRepo.FetchByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	data := []*domain.CartResponse{}
	utils.CopyJsonStruct(carts, &data)

	return data, nil
}

func (c *cartUsecase) FindByStatusAndUserID(ctx context.Context, status string, userID uint64) (*domain.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return c.cartRepo.FindByStatusAndUserID(ctx, status, userID)
}

func (c *cartUsecase) UpdateCartByCartID(ctx context.Context, cartID uint64, cart *domain.CartUpdateRequest) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	now := time.Now()
	cartUpdate := &domain.Cart{
		ID:              cartID,
		UserID:          cart.UserID,
		PaymentID:       cart.PaymentID,
		Status:          *cart.Status,
		ShippingAddress: &cart.ShippingAddress,
		TotalPrice:      cart.TotalPrice,
		UpdatedAt:       &now,
	}

	return c.cartRepo.UpdateByCartID(ctx, cartUpdate)
}

func NewCartUsecase(cartRepo cart.CartRepository, productUc product.ProductUsecase, cfg *config.Config) CartUsecase {
	return &cartUsecase{
		cartRepo:  cartRepo,
		timeout:   time.Duration(cfg.TimeoutCtx) * time.Second,
		productUc: productUc,
	}
}

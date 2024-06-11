package cart

import (
	"context"
	"go-commerce/domain"

	"gorm.io/gorm"
)

type CartRepository interface {
	Create(ctx context.Context, cart *domain.Cart) error
	FetchByUserID(ctx context.Context, userID uint64) ([]*domain.Cart, error)
	FindByStatusAndUserID(ctx context.Context, status string, userID uint64) (*domain.Cart, error)
	UpdateByCartID(ctx context.Context, cart *domain.Cart) error
	// UpdateCartByCartID(ctx context.Context, cartID uint64, cart *domain.CartUpdateRequest) error

	AddToCart(ctx context.Context, cartDetail *domain.CartDetail) error
	Delete(ctx context.Context, cartDetailID uint64) error
	FetchByCartID(ctx context.Context, cartID uint64) ([]*domain.CartDetail, error)
	FetchCartDetailByIDs(ctx context.Context, cartDetailIDs any) ([]*domain.CartDetail, error)
	UpdateCartDetailByID(ctx context.Context, cartDetail *domain.CartDetail) error
	UpdateCartIDByIDs(ctx context.Context, cartDetailIDs any, cartDetail *domain.CartDetail) error
	FindByID(ctx context.Context, cartDetailID uint64) (*domain.CartDetail, error)
}

type cartRepository struct {
	db *gorm.DB
}

func (c *cartRepository) UpdateByCartID(ctx context.Context, cart *domain.Cart) error {
	if err := c.db.Where("id = ?", cart.ID).Save(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) Create(ctx context.Context, cart *domain.Cart) error {
	if err := c.db.Create(cart).Error; err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) FetchByUserID(ctx context.Context, userID uint64) ([]*domain.Cart, error) {
	carts := []*domain.Cart{}
	if err := c.db.Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (c *cartRepository) FindByStatusAndUserID(ctx context.Context, status string, userID uint64) (*domain.Cart, error) {
	cart := &domain.Cart{}
	if err := c.db.Where("status = ? AND user_id = ?", status, userID).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		db: db,
	}
}

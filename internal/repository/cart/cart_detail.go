package cart

import (
	"context"
	"go-commerce/domain"
)

func (c *cartRepository) FindByID(ctx context.Context, cartDetailID uint64) (*domain.CartDetail, error) {
	cartDetail := &domain.CartDetail{}
	if err := c.db.Where("id = ?", cartDetailID).First(&cartDetail).Error; err != nil {
		return nil, err
	}
	return cartDetail, nil
}

func (c *cartRepository) UpdateCartIDByIDs(ctx context.Context, cartDetailIDs any, cartDetail *domain.CartDetail) error {
	if err := c.db.Model(&cartDetail).Where("id IN ?", cartDetailIDs).Update("cart_id", cartDetail.CartID).Error; err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) FetchCartDetailByIDs(ctx context.Context, cartDetailIDs any) ([]*domain.CartDetail, error) {
	detailCarts := []*domain.CartDetail{}
	if err := c.db.Where("id IN ?", cartDetailIDs).Find(&detailCarts).Error; err != nil {
		return nil, err
	}
	return detailCarts, nil
}

func (c *cartRepository) UpdateCartDetailByID(ctx context.Context, cartDetail *domain.CartDetail) error {
	if err := c.db.Where("id = ?", cartDetail.ID).Save(cartDetail).Error; err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) FetchByCartID(ctx context.Context, cartID uint64) ([]*domain.CartDetail, error) {
	cartDetails := []*domain.CartDetail{}
	if err := c.db.Where("cart_id = ?", cartID).Find(&cartDetails).Error; err != nil {
		return nil, err
	}

	return cartDetails, nil
}

func (c *cartRepository) AddToCart(ctx context.Context, cartDetail *domain.CartDetail) error {
	if err := c.db.Create(cartDetail).Error; err != nil {
		return err
	}
	return nil
}

func (c *cartRepository) Delete(ctx context.Context, cartDetailID uint64) error {
	cartDetail := &domain.CartDetail{}
	if err := c.db.Where("id = ?", cartDetailID).Delete(&cartDetail).Error; err != nil {
		return err
	}
	return nil
}

package cart

import (
	"context"
	"errors"
	"go-commerce/domain"

	"go-commerce/pkg/utils"
	"time"

	"gorm.io/gorm"
)

func (c *cartUsecase) FindByID(ctx context.Context, cartDetailID uint64) (*domain.CartDetail, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return c.cartRepo.FindByID(ctx, cartDetailID)
}

func (c *cartUsecase) UpdateCartIDByIDs(ctx context.Context, cartDetailIDs any, cartDetail *domain.CartDetail) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return c.cartRepo.UpdateCartIDByIDs(ctx, cartDetailIDs, cartDetail)
}

func (c *cartUsecase) FetchCartDetailByIDs(ctx context.Context, cartDetailIDs any) ([]*domain.CartDetail, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return c.cartRepo.FetchCartDetailByIDs(ctx, cartDetailIDs)
}

func (c *cartUsecase) UpdateCartDetailByID(ctx context.Context, cartDetail *domain.UpdateCartDetail) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	now := time.Now()
	cartDetailUpdate := &domain.CartDetail{
		ID:        cartDetail.ID,
		CartID:    cartDetail.CartID,
		ProductID: *cartDetail.ProductID,
		Quantity:  cartDetail.Quantity,
		Subtotal:  *cartDetail.Subtotal,
		UpdatedAt: &now,
	}

	return c.cartRepo.UpdateCartDetailByID(ctx, cartDetailUpdate)
}

func (c *cartUsecase) AddToCart(ctx context.Context, userID uint64, cartDetail *domain.CreateCartDetail) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	// use cart with status "On Cart",
	// if not found, create new cart
	cart, err := c.cartRepo.FindByStatusAndUserID(ctx, domain.StatusOnCart, userID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			err = c.cartRepo.Create(ctx, &domain.Cart{})
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	product, err := c.productUc.GetProductByID(ctx, cartDetail.ProductID)
	if err != nil {
		return err
	}

	if product.Stock < cartDetail.Quantity {
		return &utils.CustomError{Message: "Limited product stock"}
	}

	product.Stock -= cartDetail.Quantity
	if err := c.productUc.Update(ctx, &domain.Product{
		ID:         product.ID,
		CategoryID: product.CategoryID,
		Name:       product.Name,
		Desciption: product.Desciption,
		Stock:      product.Stock,
		Price:      product.Price,
	}); err != nil {
		return err
	}

	subTotal := product.Price * cartDetail.Quantity

	newCartDetail := &domain.CartDetail{
		CartID:    cart.ID,
		ProductID: cartDetail.ProductID,
		Quantity:  cartDetail.Quantity,
		Subtotal:  subTotal,
		CreatedAt: time.Now(),
	}

	return c.cartRepo.AddToCart(ctx, newCartDetail)
}

func (c *cartUsecase) Delete(ctx context.Context, cartDetailID uint64) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	cart, err := c.cartRepo.FindByID(ctx, cartDetailID)
	if err != nil {
		return err
	}

	product, err := c.productUc.GetProductByID(ctx, uint64(cart.ProductID))
	if err != nil {
		return err
	}

	now := time.Now()
	if err := c.productUc.Update(ctx, &domain.Product{
		ID:         product.ID,
		CategoryID: product.CategoryID,
		Name:       product.Name,
		Desciption: product.Desciption,
		Stock:      product.Stock + cart.Quantity,
		Price:      product.Price,
		UpdatedAt:  &now,
	}); err != nil {
		return err
	}

	return c.cartRepo.Delete(ctx, cartDetailID)
}

func (c *cartUsecase) FetchByCartID(ctx context.Context, cartID uint64) ([]*domain.CartDetail, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	return c.cartRepo.FetchByCartID(ctx, cartID)
}

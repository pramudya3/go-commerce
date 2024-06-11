package product

import (
	"context"
	"go-commerce/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	GetProductByCategory(ctx context.Context, categoryId uint64) ([]*domain.Product, error)
	GetProductByID(ctx context.Context, productID uint64) (*domain.Product, error)
	FetchProduct(ctx context.Context) ([]*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) FetchProduct(ctx context.Context) ([]*domain.Product, error) {
	products := []*domain.Product{}
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepository) Update(ctx context.Context, product *domain.Product) error {
	if err := p.db.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p *productRepository) GetProductByID(ctx context.Context, productID uint64) (*domain.Product, error) {
	product := &domain.Product{}
	if err := p.db.Where("id = ?", productID).First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepository) Create(ctx context.Context, product *domain.Product) error {
	if err := p.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *productRepository) GetProductByCategory(ctx context.Context, categoryId uint64) ([]*domain.Product, error) {
	products := []*domain.Product{}

	if err := p.db.Where("category_id = ?", categoryId).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

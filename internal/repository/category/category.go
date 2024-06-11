package category

import (
	"context"
	"go-commerce/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	FetchCategories(ctx context.Context) ([]*domain.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func (c *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	if err := c.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (c *categoryRepository) FetchCategories(ctx context.Context) ([]*domain.Category, error) {
	categories := []*domain.Category{}
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

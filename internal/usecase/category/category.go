package category

import (
	"context"
	"go-commerce/domain"
	repo "go-commerce/internal/repository/category"
	"go-commerce/pkg/config"
	"go-commerce/pkg/utils"
	"time"
)

type CategoryUsecase interface {
	Create(ctx context.Context, category *domain.CreateCategory) error
	FetchCategories(ctx context.Context) ([]*domain.CategoryResponse, error)
}

type categoryUsecase struct {
	categoryRepository repo.CategoryRepository
	timeout            time.Duration
}

func (c *categoryUsecase) Create(ctx context.Context, category *domain.CreateCategory) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	newCategory := &domain.Category{
		Name: category.Name,
	}

	return c.categoryRepository.Create(ctx, newCategory)
}

func (c *categoryUsecase) FetchCategories(ctx context.Context) ([]*domain.CategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	categories, err := c.categoryRepository.FetchCategories(ctx)
	if err != nil {
		return nil, err
	}

	datas := []*domain.CategoryResponse{}
	utils.CopyJsonStruct(categories, &datas)

	return datas, nil
}

func NewCategoryUsecase(categoryRepo repo.CategoryRepository, cfg *config.Config) CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: categoryRepo,
		timeout:            time.Duration(cfg.TimeoutCtx) * time.Second,
	}
}

package product

import (
	"context"
	"go-commerce/domain"
	repo "go-commerce/internal/repository/product"
	"go-commerce/pkg/config"
	"go-commerce/pkg/utils"
	"time"
)

type ProductUsecase interface {
	Create(ctx context.Context, product *domain.CreateProductRequest) error
	GetProductByCategory(ctx context.Context, categoryId uint64) ([]*domain.ProductResponse, error)
	GetProductByID(ctx context.Context, productID uint64) (*domain.ProductResponse, error)
	FetchProduct(ctx context.Context) ([]*domain.ProductResponse, error)
	Update(ctx context.Context, product *domain.Product) error
}

type productUsecase struct {
	productRepo repo.ProductRepository
	timeout     time.Duration
}

func (p *productUsecase) FetchProduct(ctx context.Context) ([]*domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	products, err := p.productRepo.FetchProduct(ctx)
	if err != nil {
		return nil, err
	}

	data := []*domain.ProductResponse{}
	utils.CopyJsonStruct(products, &data)

	return data, nil
}

func (p *productUsecase) Update(ctx context.Context, product *domain.Product) error {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	now := time.Now()
	product.UpdatedAt = &now

	return p.productRepo.Update(ctx, product)
}

func (p *productUsecase) GetProductByID(ctx context.Context, productID uint64) (*domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	product, err := p.productRepo.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	data := &domain.ProductResponse{}
	utils.CopyJsonStruct(product, &data)

	return data, nil
}

func (p *productUsecase) Create(ctx context.Context, product *domain.CreateProductRequest) error {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	newProduct := &domain.Product{
		CategoryID: product.CategoryID,
		Name:       product.Name,
		Desciption: product.Description,
		Stock:      product.Stock,
		Price:      product.Price,
		CreatedAt:  time.Now(),
	}

	return p.productRepo.Create(ctx, newProduct)
}

func (p *productUsecase) GetProductByCategory(ctx context.Context, categoryId uint64) ([]*domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	products, err := p.productRepo.GetProductByCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	datas := []*domain.ProductResponse{}
	utils.CopyJsonStruct(products, &datas)

	return datas, nil
}

func NewProductUsecase(productRepo repo.ProductRepository, cfg *config.Config) ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
		timeout:     time.Duration(cfg.TimeoutCtx) * time.Second,
	}
}

package product

import (
	productRepo "go-commerce/internal/repository/product"
	productUc "go-commerce/internal/usecase/product"
	"go-commerce/pkg/config"
	"go-commerce/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(g *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	productRepo := productRepo.NewProductRepository(db)
	productUsecase := productUc.NewProductUsecase(productRepo, cfg)
	productHandler := NewProductHandler(productUsecase)

	auth := middleware.Auth()

	route := g.Group("/product")
	{
		route.POST("/", auth, productHandler.CreateProduct)
		route.GET("/", productHandler.FetchProduct)
		route.GET("/category/:id", productHandler.GetProductByCategoryID)
	}
}

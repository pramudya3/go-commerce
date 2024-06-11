package cart

import (
	repo "go-commerce/internal/repository/cart"
	productRepo "go-commerce/internal/repository/product"
	uc "go-commerce/internal/usecase/cart"
	productUc "go-commerce/internal/usecase/product"
	"go-commerce/pkg/config"
	"go-commerce/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CartRoutes(g *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	productRepo := productRepo.NewProductRepository(db)
	cartRepo := repo.NewCartRepository(db)

	productUc := productUc.NewProductUsecase(productRepo, cfg)
	cartUc := uc.NewCartUsecase(cartRepo, productUc, cfg)

	cart := NewCartHandler(cartUc, productUc)

	auth := middleware.Auth()

	route := g.Group("/cart")
	{
		route.GET("/", auth, cart.MyCart)
		route.POST("/", auth, cart.AddToCart)
		route.DELETE("/detail/:id", auth, cart.DeleteCartDetail)
		route.GET("/all", auth, cart.FindCartByID)
	}

	otherRoute := g.Group("/")
	{
		otherRoute.POST("/checkout", auth, cart.Chekcout)
		otherRoute.GET("/order", auth, cart.MyOrder)
	}
}

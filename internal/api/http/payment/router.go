package payment

import (
	"go-commerce/internal/repository/cart"
	"go-commerce/internal/repository/payment"
	"go-commerce/internal/repository/product"
	cartUc "go-commerce/internal/usecase/cart"
	uc "go-commerce/internal/usecase/payment"
	productUc "go-commerce/internal/usecase/product"
	"go-commerce/pkg/config"
	"go-commerce/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentRoutes(g *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	paymentRepo := payment.NewPaymentRepository(db)
	cartRepo := cart.NewCartRepository(db)
	productRepo := product.NewProductRepository(db)

	productUc := productUc.NewProductUsecase(productRepo, cfg)
	cartUc := cartUc.NewCartUsecase(cartRepo, productUc, cfg)
	paymentUc := uc.NewPaymentUsecase(cartUc, paymentRepo, cfg)

	paymentHandler := NewPaymentHandler(paymentUc, cartUc)

	auth := middleware.Auth()

	route := g.Group("/payment")
	{
		route.POST("/", auth, paymentHandler.CreatePayment)
		route.GET("/", auth, paymentHandler.PaymentInfo)
	}
}

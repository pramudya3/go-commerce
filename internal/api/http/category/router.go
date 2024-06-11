package category

import (
	categoryRepo "go-commerce/internal/repository/category"
	categoryUC "go-commerce/internal/usecase/category"
	"go-commerce/pkg/config"
	"go-commerce/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoutes(g *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	categoryRepo := categoryRepo.NewCategoryRepository(db)
	categoryUc := categoryUC.NewCategoryUsecase(categoryRepo, cfg)
	categoryHandler := NewCategoryHandler(categoryUc)

	auth := middleware.Auth()

	route := g.Group("/category")
	{
		route.POST("/", auth, categoryHandler.Create)
		route.GET("/", categoryHandler.FetchCategory)
	}
}

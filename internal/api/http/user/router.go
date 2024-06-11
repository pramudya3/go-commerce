package http

import (
	repo "go-commerce/internal/repository/user"
	uc "go-commerce/internal/usecase/user"
	"go-commerce/pkg/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(g *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	userRepo := repo.NewUserRepository(db)
	userUsecase := uc.NewUserUsecase(userRepo, cfg)
	userHandler := NewUserHandler(userUsecase)

	route := g.Group("/user")
	{
		route.POST("/register", userHandler.Register)
		route.POST("/login", userHandler.Login)
	}
}

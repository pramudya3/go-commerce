package server

import (
	"fmt"
	"go-commerce/domain"
	"go-commerce/internal/api/http/cart"
	category "go-commerce/internal/api/http/category"
	"go-commerce/internal/api/http/payment"
	product "go-commerce/internal/api/http/product"
	user "go-commerce/internal/api/http/user"

	"go-commerce/pkg/config"
	database "go-commerce/pkg/database/gorm"
	rds "go-commerce/pkg/database/redis"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	gin   *gin.Engine
	cfg   *config.Config
	db    *gorm.DB
	redis *redis.Client
}

func NewServer() *Server {
	cfg := config.LoadConfig()
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalln("failed to init database")
	}

	db.AutoMigrate(&domain.User{}, &domain.Category{}, &domain.Product{}, &domain.Cart{}, &domain.CartDetail{}, &domain.Payment{})

	rdb, err := rds.NewRedis(cfg)
	if err != nil {
		log.Fatalln("failed to init redis")
	}

	return &Server{
		gin:   gin.Default(),
		cfg:   cfg,
		db:    db,
		redis: rdb,
	}
}

func (s *Server) Start() {
	s.initRoutes()

	log.Println("HTTP server is listening at: ", s.cfg.ServerAddr)
	log.Fatalln(s.gin.Run(fmt.Sprintf("%v", s.cfg.ServerAddr)))
}

func (s *Server) initRoutes() {
	v1 := s.gin.Group("/api/v1")
	user.UserRoutes(v1, s.db, s.cfg)
	product.ProductRoutes(v1, s.db, s.cfg)
	category.CategoryRoutes(v1, s.db, s.cfg)
	cart.CartRoutes(v1, s.db, s.cfg)
	payment.PaymentRoutes(v1, s.db, s.cfg)
}

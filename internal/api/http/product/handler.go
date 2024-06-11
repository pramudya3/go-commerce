package product

import (
	"errors"
	"go-commerce/domain"
	productUc "go-commerce/internal/usecase/product"
	"go-commerce/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productHandler struct {
	productUsecase productUc.ProductUsecase
}

func NewProductHandler(productUsecase productUc.ProductUsecase) *productHandler {
	return &productHandler{
		productUsecase: productUsecase,
	}
}

func (p *productHandler) CreateProduct(c *gin.Context) {
	req := &domain.CreateProductRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, err.Error()))
		return
	}

	if err := p.productUsecase.Create(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusCreated)
}

func (p *productHandler) FetchProduct(c *gin.Context) {
	products, err := p.productUsecase.FetchProduct(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(http.StatusOK, products))
}

func (p *productHandler) GetProductByCategoryID(c *gin.Context) {
	categoryIdStr := c.Param("id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, err.Error()))
		return
	}

	products, err := p.productUsecase.GetProductByCategory(c, uint64(categoryId))
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, utils.ResponseFailed(http.StatusNotFound, "Category not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(http.StatusOK, products))
}

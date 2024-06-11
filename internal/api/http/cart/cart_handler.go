package cart

import (
	"errors"
	"go-commerce/domain"
	"go-commerce/internal/usecase/cart"
	"go-commerce/internal/usecase/product"
	"go-commerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type cartHandler struct {
	cartUc    cart.CartUsecase
	productUc product.ProductUsecase
}

func NewCartHandler(cartUc cart.CartUsecase, productUc product.ProductUsecase) *cartHandler {
	return &cartHandler{
		cartUc:    cartUc,
		productUc: productUc,
	}
}

func (h *cartHandler) MyCart(c *gin.Context) {
	id := c.GetFloat64("user_id")

	myCart, err := h.cartUc.MyCart(c, uint64(id))
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, utils.ResponseFailed(http.StatusNotFound, "Cart not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(http.StatusOK, myCart))
}

func (h *cartHandler) MyOrder(c *gin.Context) {
	id := c.GetFloat64("user_id")

	myCart, err := h.cartUc.MyOrder(c, uint64(id))
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, utils.ResponseFailed(http.StatusNotFound, "Order not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(http.StatusOK, myCart))
}

func (h *cartHandler) Chekcout(c *gin.Context) {
	id := c.GetFloat64("user_id")
	req := &domain.CheckoutRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.cartUc.Checkout(c, uint64(id), req); err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, utils.ResponseFailed(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}
}

func (h *cartHandler) FindCartByID(c *gin.Context) {
	id := c.GetFloat64("user_id")
	cart, err := h.cartUc.FetchByUserID(c, uint64(id))
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, utils.ResponseFailed(http.StatusNotFound, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseFailed(http.StatusOK, cart))
}

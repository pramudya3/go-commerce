package cart

import (
	"errors"
	"go-commerce/domain"
	"go-commerce/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *cartHandler) AddToCart(c *gin.Context) {
	id := c.GetFloat64("user_id")

	req := &domain.CreateCartDetail{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.cartUc.AddToCart(c, uint64(id), req); err != nil {
		custErr, ok := err.(*utils.CustomError)
		if ok {
			c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, custErr.Message))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *cartHandler) DeleteCartDetail(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.cartUc.Delete(c, uint64(id)); err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, utils.ResponseFailed(http.StatusNotFound, "Product not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

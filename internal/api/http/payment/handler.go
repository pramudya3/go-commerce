package payment

import (
	"go-commerce/domain"
	"go-commerce/internal/usecase/cart"
	"go-commerce/internal/usecase/payment"
	"go-commerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentUc payment.PaymentUsecase
	cartUc    cart.CartUsecase
}

func NewPaymentHandler(paymentUc payment.PaymentUsecase, cartUc cart.CartUsecase) *paymentHandler {
	return &paymentHandler{
		paymentUc: paymentUc,
		cartUc:    cartUc,
	}
}

func (h *paymentHandler) CreatePayment(c *gin.Context) {
	userID := c.GetFloat64("user_id")

	req := &domain.PaymentRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseSuccess(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := h.paymentUc.CreatePayment(c, uint64(userID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseSuccess(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(http.StatusOK, res))
}

func (h *paymentHandler) PaymentInfo(c *gin.Context) {
	id := c.GetFloat64("user_id")

	payments, err := h.paymentUc.FetchPayment(c, uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseSuccess(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(http.StatusOK, payments))
}

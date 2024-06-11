package category

import (
	"go-commerce/domain"
	categoryUc "go-commerce/internal/usecase/category"
	"go-commerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryUsecase categoryUc.CategoryUsecase
}

func NewCategoryHandler(categoryUc categoryUc.CategoryUsecase) *categoryHandler {
	return &categoryHandler{
		categoryUsecase: categoryUc,
	}
}

func (h *categoryHandler) Create(c *gin.Context) {
	req := &domain.CreateCategory{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.categoryUsecase.Create(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error))
		return
	}

	c.Status(http.StatusCreated)
}

func (h *categoryHandler) FetchCategory(c *gin.Context) {
	categories, err := h.categoryUsecase.FetchCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(http.StatusOK, categories))
}

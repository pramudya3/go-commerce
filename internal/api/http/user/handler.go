package http

import (
	"errors"
	"go-commerce/domain"
	uc "go-commerce/internal/usecase/user"
	"go-commerce/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type userHandler struct {
	usecase uc.UserUsecase
}

func NewUserHandler(userUsecase uc.UserUsecase) *userHandler {
	return &userHandler{
		usecase: userUsecase,
	}
}

func (u *userHandler) Register(c *gin.Context) {
	req := &domain.UserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, err.Error()))
		return
	}

	if err := u.usecase.Register(c, req); err != nil {
		pgErr := &pgconn.PgError{}
		if errors.As(err, &pgErr) {
			// catch unique constraint violation error
			if pgErr.Code == "23505" {
				c.JSON(http.StatusConflict, utils.ResponseFailed(http.StatusConflict, "Email already exist"))
				return
			}
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusCreated)
}

func (u *userHandler) Login(c *gin.Context) {
	login := &domain.UserLogin{}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := u.usecase.Login(c, login)
	if err != nil {
		custErr, ok := err.(*utils.CustomError)
		if ok {
			c.JSON(http.StatusBadRequest, utils.ResponseFailed(http.StatusBadRequest, custErr.Message))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(http.StatusOK, res))
}

func (u *userHandler) Logout(c *gin.Context) {
	id := c.GetInt("user_id")

	if err := u.usecase.Logout(c, uint64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (u *userHandler) RefreshToken(c *gin.Context) {
	id := c.GetInt("user_id")

	res, err := u.usecase.RefreshToken(c, uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseFailed(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess(http.StatusOK, res))
}

func (u *userHandler) ExtractToken(c *gin.Context) {
	id := c.GetFloat64("user_id")

	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
	})
}

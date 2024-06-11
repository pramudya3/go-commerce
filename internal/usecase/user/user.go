package user_usecase

import (
	"context"
	"go-commerce/domain"
	repo "go-commerce/internal/repository/user"
	"go-commerce/pkg/config"
	"go-commerce/pkg/token"
	"go-commerce/pkg/utils"
	"time"
)

type UserUsecase interface {
	Register(ctx context.Context, user *domain.UserRequest) error
	Login(ctx context.Context, user *domain.UserLogin) (*domain.UserToken, error)
	Logout(ctx context.Context, id uint64) error
	RefreshToken(ctx context.Context, id uint64) (*domain.UserToken, error)
	FindUserByID(ctx context.Context, id uint64) (*domain.UserResponse, error)
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userUsecase struct {
	userRepo repo.UserRepository
	timeout  time.Duration
}

func NewUserUsecase(userRepo repo.UserRepository, cfg *config.Config) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		timeout:  time.Duration(cfg.TimeoutCtx) * time.Second,
	}
}

func (u *userUsecase) Register(ctx context.Context, user *domain.UserRequest) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	user.HashPassword()
	newUser := &domain.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return u.userRepo.Create(ctx, newUser)
}

func (u *userUsecase) Login(ctx context.Context, user *domain.UserLogin) (*domain.UserToken, error) {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	userExist, err := u.userRepo.FindUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, &utils.CustomError{
			Message: "Email not found",
		}
	}

	if ok := userExist.CompareHashPassword(user.Password); !ok {
		return nil, &utils.CustomError{
			Message: "Wrong password",
		}
	}

	accessToken := token.GenerateAccessToken(userExist.ID)
	refreshToken := token.GenerateRefreshToken(userExist.ID)

	return &domain.UserToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *userUsecase) Logout(ctx context.Context, id uint64) error {
	// clear user token from redis

	return nil
}

func (u *userUsecase) RefreshToken(ctx context.Context, id uint64) (*domain.UserToken, error) {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	// todo: generate new access_token from refresh token
	// todo: store it on redis

	return nil, nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, id uint64, user *domain.UserRequest) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	user.HashPassword()
	userUpdate := &domain.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		UpdatedAt: time.Now(),
	}

	return u.userRepo.Update(ctx, userUpdate)
}

func (u *userUsecase) FindUserByID(ctx context.Context, id uint64) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	user, err := u.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}, nil
}

func (u *userUsecase) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepo.FindUserByEmail(ctx, email)
}

package user

import (
	"context"
	"go-commerce/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindUserByID(ctx context.Context, id uint64) (*domain.User, error)
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindUserByID(ctx context.Context, id uint64) (*domain.User, error) {
	user := &domain.User{}
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Update(ctx context.Context, user *domain.User) error {
	if err := u.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

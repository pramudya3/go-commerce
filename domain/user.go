package domain

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		ID        uint64    `json:"id" gorm:"not null;primary_key"`
		Name      string    `json:"name"`
		Email     string    `json:"email" gorm:"unique;not null;index:idx_user_email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Cart      []Cart    `gorm:"foreignKey:UserID"`
		Payment   []Payment `gorm:"foreignKey:UserID"`
	}

	// for create or update user
	UserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserResponse struct {
		ID        uint64     `json:"id"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	UserToken struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	ReqRefreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}
)

// hashing password from user request
func (u *UserRequest) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}

// comparing password from hashed password (from database) with plain password (from request)
func (u *User) CompareHashPassword(plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword)); err != nil {
		return false
	}

	return true
}

func (u *User) AccessTokenKey() string {
	return fmt.Sprintf("x-access-token:%d", u.ID)
}

func (u *User) RefreshTokenKey() string {
	return fmt.Sprintf("x-refresh-token:%d", u.ID)
}

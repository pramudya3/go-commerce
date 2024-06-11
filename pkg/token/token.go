package token

import (
	"errors"
	"fmt"
	"go-commerce/pkg/config"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AccessTokenExpired  = 1 // hour
	RefreshTokenExpired = 7 // day
)

type (
	JwtCustomClaims struct {
		jwt.RegisteredClaims
		ID float64 `json:"id"`
	}

	JwtToken struct {
		ID float64 `json:"id"`
		// Roles string  `json:"roles"`
	}
)

func GenerateAccessToken(id uint64) string {
	cfg := config.LoadConfig()
	claims := &JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * AccessTokenExpired)},
		},
		ID: float64(id),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(cfg.TokenSecret))
	if err != nil {
		log.Fatalf("failed generate access token, err: %v", err)
		return ""
	}

	return token
}

func GenerateRefreshToken(id uint64) string {
	cfg := config.LoadConfig()
	claims := &JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().AddDate(0, 0, RefreshTokenExpired)},
		},
		ID: float64(id),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(cfg.TokenSecret))
	if err != nil {
		log.Fatalf("failed generate refresh token, err: %v", err)
		return ""
	}

	return token
}

func ValidateToken(jwtToken string) (*JwtToken, error) {
	cfg := config.LoadConfig()
	cleanJWT := strings.ReplaceAll(jwtToken, "Bearer ", "")

	token, err := jwt.Parse(cleanJWT, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.TokenSecret), nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed claims jwt token")
	}

	id, ok := claims["id"].(float64)
	if !ok {
		fmt.Println(reflect.TypeOf(claims["id"]))
		return nil, fmt.Errorf("failed to get id from claims")
	}

	return &JwtToken{
		ID: float64(id),
	}, nil
}

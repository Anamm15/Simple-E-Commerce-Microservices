package jwt

import (
	"errors"
	"fmt"
	"time"

	"api-gateway/internal/helper/constant"
	"api-gateway/internal/helper/enum"

	"github.com/golang-jwt/jwt/v5"
)

// Definisi durasi token (bisa dipindah ke config/env)
const tokenDuration = 24 * time.Hour

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// 1. Custom Claims Struct
// Menggunakan struct daripada MapClaims agar Type-Safe dan mencegah typo saat akses key.
type JWTCustomClaims struct {
	UserID string           `json:"user_id"`
	Email  string           `json:"email"`
	Role   enum.AccountRole `json:"role"`
	jwt.RegisteredClaims
}

func ValidateJWT(tokenString string) (*JWTCustomClaims, error) {
	secretKey := constant.GetJWTSecret()
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, err
	}

	// Validasi Claims dan Token
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

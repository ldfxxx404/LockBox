package models

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UserID  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
	jwt.RegisteredClaims
}

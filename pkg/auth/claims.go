package auth

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BaseClaims struct {
	UUID uuid.UUID
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

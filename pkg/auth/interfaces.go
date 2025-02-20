package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthExtendInterface interface {
	CreateClaims(baseClaims BaseClaims) CustomClaims
	CreateToken(claims CustomClaims) (string, error)
	ParseToken(tokenString string) (*CustomClaims, error)
	GetToken(c *gin.Context) string
	SetToken(c *gin.Context, token string, maxAge int)
	ClearToken(c *gin.Context)
	RefreshToken(claims *CustomClaims, token string) (*CustomClaims, string)
}

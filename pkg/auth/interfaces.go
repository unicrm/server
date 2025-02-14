package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/pkg/auth/internal"
)

type AuthExtendInterface interface {
	CreateClaims(baseClaims internal.BaseClaims) internal.CustomClaims
	CreateToken(claims internal.CustomClaims) (string, error)
	ParseToken(tokenString string) (*internal.CustomClaims, error)
	GetToken(c *gin.Context) string
	SetToken(c *gin.Context, token string, maxAge int)
	ClearToken(c *gin.Context)
	RefreshToken(claims *internal.CustomClaims, token string) (*internal.CustomClaims, string)
}

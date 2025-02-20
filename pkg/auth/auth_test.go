package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var (
	auth  *AuthExtend
	token string
)

func TestMain(m *testing.M) {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	m.Run()
}

func TestCreateToken(t *testing.T) {
	InitAuth(JWT{SigningKey: "ssssssssssss", Issuer: "unicrm"})
	var err error
	var claims CustomClaims
	token, claims, err = LoginToken(uuid.UUID{})
	zap.L().Info("CreateToken", zap.Any("claims", claims))
	assert.Nil(t, err)
}

func TestAuth(t *testing.T) {
	// 模拟请求上下文
	header := http.Header{"Cookie": []string{"x-token=" + token}}
	header.Add("x-token", token)
	request := &http.Request{Header: header}
	context := gin.Context{Request: request}
	// 验证token是否正确
	assert.Equal(t, token, auth.GetToken(&context))
}

func TestMaddleware(t *testing.T) {
	// 模拟请求上下文
	header := http.Header{"Cookie": []string{"x-token=" + token}}
	header.Add("x-token", token)
	request := &http.Request{Header: header}
	context := gin.Context{Request: request}
	// 初始化缓存
	AUTH_BLACK_CACHE = cache.New(5*time.Minute, 10*time.Minute)
	// 测试中间件
	handlerFunc := auth.JWTAuthMiddleware()
	handlerFunc(&context)
}

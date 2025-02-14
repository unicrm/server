package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/initialize"
	"github.com/unicrm/server/pkg/auth"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var token string

func TestMain(m *testing.M) {
	// 日志初始化
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	// 初始化配置文件
	gin.SetMode(gin.TestMode)
	initialize.LoadConfig("../../../config.debug.yaml")
	router = initialize.Routers()
	m.Run()
}

func TestRouters(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

func TestLogin(t *testing.T) {
	auth.InitAuth(globals.UNICRM_CONFIG.JWT)
	auth.AUTH_BLACK_CACHE = cache.New(5*time.Minute, 10*time.Minute)
	auth.SetBlacklist([]string{})
	var err error
	token, _, err = auth.LoginToken(uuid.New())
	if err != nil {
		zap.L().Panic(err.Error())
	}
}

func TestAuth(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/jwt", nil)
	req.Header.Set("x-token", token)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/initialize"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	initialize.LoadConfig("../config.debug.yaml")
	globals.UNICRM_LOGGER, _ = zap.NewDevelopment()
	m.Run()
}

func TestRouters(t *testing.T) {
	router := initialize.Routers()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

const (
	SuccessCode = 0
	ErrorCode   = -1
)

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		ErrorCode,
		nil,
		message,
	})
}

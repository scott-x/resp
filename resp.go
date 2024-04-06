package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// success
func Success[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// error instance
func ErrorBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"data": "bad request",
	})
}

func ErrorNotAuthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"data": "not authorized",
	})
}

func ErrorForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"code": 403,
		"data": "content forbidden",
	})
}

func ErrorNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code": 404,
		"data": "not found",
	})
}

func ErrorServerInternal(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"data": "internal error",
	})
}

// custom error: suggest starting from 2000
func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": message,
	})
}

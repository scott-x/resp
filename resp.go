package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_BAD_REQUEST       = "bad request"
	ERROR_NOT_ATHORIZED     = "permission not allowed"
	ERROR_CONTENT_FORBIDDEN = "content forbidden"
	ERROR_NOT_FOUND         = "not found"
	ERROR_INTERNAL_ERROR    = "internal error"
)

// gin.H equals to map[string]interface{}
func genMap(code int, T any) map[string]any {
	m := make(map[string]any)
	m["code"] = code
	m["data"] = T
	return m
}

// success
func Success[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, genMap(200, data))
}

// error instance
func ErrorBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, genMap(400, ERROR_BAD_REQUEST))
}

func ErrorNotAuthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, genMap(401, ERROR_NOT_ATHORIZED))
}

func ErrorForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, genMap(403, ERROR_CONTENT_FORBIDDEN))
}

func ErrorNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, genMap(404, ERROR_NOT_FOUND))
}

func ErrorServerInternal(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, genMap(500, ERROR_INTERNAL_ERROR))
}

// custom error: suggest starting from 2000
func Error[T any](c *gin.Context, code int, data T) {
	c.JSON(http.StatusOK, genMap(code, data))
}

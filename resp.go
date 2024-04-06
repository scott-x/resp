package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// success
func Success[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": ResponseCode(SUCCESS),
		"data": data,
	})
}

// return "success" message
func SuccessInfo(c *gin.Context) {
	Success[string](c, "success")
}

// error
func Error(c *gin.Context, code ResponseCode, message string) {
	if code < 2000 {
		panic("invalid error return code, as promised it should be started from 2001")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

// error instance
func ErrorInternal(c *gin.Context) {
	Error(c, ResponseCode(INTERNAL_ERROR), "internal error")
}

func ErrorBadRequest(c *gin.Context) {
	Error(c, ResponseCode(BAD_REQUEST), "internal error")
}

func ErrorNotAuthorized(c *gin.Context) {
	Error(c, ResponseCode(NotAuthorized), "not authorized")
}

package resp

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// success
func Success[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": int(ResponseCode(SUCCESS)),
		"data": data,
	})
}

// return "success" message
func SuccessMessage(c *gin.Context) {
	Success[string](c, "success")
}

// error
func customError(c *gin.Context, code ResponseCode, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": int(code),
		"msg":  message,
	})
}

// error instance
func ErrorInternal(c *gin.Context) {
	customError(c, ResponseCode(ERROR_INTERNAL), "internal error")
}

func ErrorBadRequest(c *gin.Context) {
	customError(c, ResponseCode(ERROR_BAD_REQUEST), "bad request")
}

func ErrorNotAuthorized(c *gin.Context) {
	customError(c, ResponseCode(ERROR_NOT_AUTHORIZED), "not authorized")
}

func Error(c *gin.Context, code int, message string) {
	if code < int(ResponseCode(ERROR_NOT_AUTHORIZED)) {
		log.Panicf("invalid error of the return code, as promised it should be no less than %d", int(ResponseCode(ERROR_NOT_AUTHORIZED)))
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

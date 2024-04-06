package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/scott-x/resp/mycode"
	"log"
	"net/http"
)

// success
func Success[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": int(mycode.ResponseCode(mycode.SUCCESS)),
		"data": data,
	})
}

// return "success" message
func SuccessMessage(c *gin.Context) {
	Success[string](c, "success")
}

// error
func customError(c *gin.Context, code mycode.ResponseCode, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": int(code),
		"msg":  message,
	})
}

// error instance
func ErrorInternal(c *gin.Context) {
	customError(c, mycode.ResponseCode(mycode.ERROR_INTERNAL), "internal error")
}

func ErrorBadRequest(c *gin.Context) {
	customError(c, mycode.ResponseCode(mycode.ERROR_BAD_REQUEST), "bad request")
}

func ErrorNotAuthorized(c *gin.Context) {
	customError(c, mycode.ResponseCode(mycode.ERROR_NOT_AUTHORIZED), "not authorized")
}

func Error(c *gin.Context, code int, message string) {
	minRequirement := int(mycode.ResponseCode(mycode.ERROR_NOT_AUTHORIZED))
	if code < minRequirement {
		log.Panicf("invalid error of the return mycode, as promised it should be no less than %d", minRequirement)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

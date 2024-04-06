# resp
golang http response encapsulation base on gin

### API
- `func Success[T any](c *gin.Context, data T)`
- `func Error(c *gin.Context, code int, message string)`
- `func ErrorBadRequest(c *gin.Context)`
- `func ErrorNotAuthorized(c *gin.Context)`
- `func ErrorForbidden(c *gin.Context)`
- `func ErrorNotFound(c *gin.Context)`
- `func ErrorServerInternal(c *gin.Context)`


### demo

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scott-x/resp"
)

func main() {
	route := gin.Default()

	route.GET("/test", func(c *gin.Context) {
		resp.SuccessMessage(c)
	})
	
	route.GET("/test1", func(c *gin.Context) {
		resp.ErrorInternal(c)
	})

	route.GET("/test2", func(c *gin.Context) {
		resp.ErrorBadRequest(c)
	})

	route.GET("/test3", func(c *gin.Context) {
		resp.ErrorNotAuthorized(c)
	})

	route.GET("/test4", func(c *gin.Context) {
		resp.Error(c, 2001, "username, password does not match")
	})

	route.GET("/test5", func(c *gin.Context) {
		type Output struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		o := Output{
			Name: "Scott",
			Age:  18,
		}
		resp.Success(c, o)
	})

	route.GET("/test6", func(c *gin.Context) {
		resp.Success(c, "success")
	})

	route.Run(":8989")
}
```
# resp
golang http response encapsulation base on gin

### API

- `func Success[T any](c *gin.Context, data T)`: success response and the return code is 2000
- `func Error(c *gin.Context, code int, message string)`: error response, as promised the return code should be stared from 2001
- `func ErrorInternal(c *gin.Context) `
- `func ErrorBadRequest(c *gin.Context)`
- `func ErrorNotAuthorized(c *gin.Context)`


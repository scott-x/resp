# resp
golang http response encapsulation base on gin

### API
- `func Success[T any](c *gin.Context, data T)`: success response and the return code is 2000
- `func SuccessMessage(c *gin.Context)`: return success message
- `func ErrorInternal(c *gin.Context) `: return code 2001
- `func ErrorBadRequest(c *gin.Context)`: return code 2002
- `func ErrorNotAuthorized(c *gin.Context)`:return code 2003
- `func Error(c *gin.Context, code int, message string)`: ***code must be no less than 2003, otherwise it will panic***
package resp

type ResponseCode int

const (
	SUCCESS ResponseCode = 2000 + iota
	ERROR_INTERNAL
	ERROR_BAD_REQUEST
	ERROR_NOT_AUTHORIZED
)

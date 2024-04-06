package resp

type ResponseCode int

const (
	SUCCESS ResponseCode = 2000 + iota
	INTERNAL_ERROR
	BAD_REQUEST
	NotAuthorized
)

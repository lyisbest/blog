package utils

type blogErrorInterface interface {
	Code() int
	Message() string
}

type BlogError struct {
	ErrorCode    int
	ErrorMessage string
}

func (e BlogError) Code() int {
	return e.ErrorCode
}

func (e BlogError) Message() string {
	return e.ErrorMessage
}

func (e BlogError) Error() string {
	return "ErrorCode: " + string(rune(e.Code())) + ", ErrorMessage: " + e.Message()
}

var QueryError = BlogError{ErrorCode: -1, ErrorMessage: "query failed!"}

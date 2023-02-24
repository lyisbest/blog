package utils

import "net/http"

type blogErrorInterface interface {
	Code() int
	Message() string
}

type BlogError struct {
	ErrorCode    int
	ErrorMessage string
}

var InnerErrorCode int = http.StatusInternalServerError

func (e BlogError) Code() int {
	return e.ErrorCode
}

func (e BlogError) Message() string {
	return e.ErrorMessage
}

func (e BlogError) Error() string {
	return "ErrorCode: " + string(rune(e.Code())) + ", ErrorMessage: " + e.Message()
}

func ResolveError(e interface{}) (int, string) {
	err, ok := e.(error)
	if !ok {
		return 0, ""
	}
	blogError, ok1 := err.(BlogError)
	if !ok1 {
		return InnerErrorCode, err.Error()
	}
	return blogError.Code(), blogError.Message()
}

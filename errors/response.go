package errors

import (
	"fmt"
	"net/http"
)

// Response ...
type Response struct {
	Code       int
	HTTPCode   int
	Message    string
	ErrorCause error
}

func (e *Response) Error() string {
	return fmt.Sprintf("[%d] %s - %v", e.Code, e.Message, e.ErrorCause)
}

// Error Codes
var (
	UserSignUpDetailsInvalid = Response{
		Code:     400,
		HTTPCode: http.StatusBadRequest,
		Message:  "UserSignUpDetailsInvalid",
	}
)

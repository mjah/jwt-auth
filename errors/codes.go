package errors

import (
	"fmt"
	"net/http"
)

// ErrorCode ...
type ErrorCode struct {
	HTTPStatus int
	Code       int
	Title      string
	Errors     error
}

// Error Codes
var (
	DatabaseConnectionFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       001,
		Title:      "DatabaseConnectionError",
	}
	PasswordGenerationFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       002,
		Title:      "PasswordGenerationFailed",
	}
	SignUpDetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       101,
		Title:      "UserSignUpDetailsInvalid",
	}
	SignUpDetailsValidationFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       102,
		Title:      "SignUpDetailsValidationFailed",
	}
)

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("[%d] %s - %v", e.Code, e.Title, e.Errors)
}

// New ...
func New(errorCode ErrorCode, errors error) *ErrorCode {
	return &ErrorCode{
		HTTPStatus: errorCode.HTTPStatus,
		Code:       errorCode.Code,
		Title:      errorCode.Title,
		Errors:     errors,
	}
}

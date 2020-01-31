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
		Code:       1001,
		Title:      "DatabaseConnectionError",
	}
	DatabaseQueryFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       1002,
		Title:      "DatabaseQueryFailed",
	}
	PasswordGenerationFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       2001,
		Title:      "PasswordGenerationFailed",
	}
	SignUpDetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3001,
		Title:      "UserSignUpDetailsInvalid",
	}
	SignUpDetailsValidationFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3002,
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

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
		Title:      "DatabaseConnectionFailed",
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
	PasswordCheckFailed = ErrorCode{
		HTTPStatus: http.StatusUnauthorized,
		Code:       2002,
		Title:      "PasswordCheckFailed",
	}
	AccessTokenIssueFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       2101,
		Title:      "AccessTokenIssueFailed",
	}
	RefreshTokenIssueFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       2103,
		Title:      "RefreshTokenIssueFailed",
	}
	AuthorizationBearerTokenEmpty = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       2102,
		Title:      "AuthorizationBearerTokenEmpty",
	}
	SignUpDetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3001,
		Title:      "SignUpDetailsInvalid",
	}
	SignUpDetailsValidationFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3002,
		Title:      "SignUpDetailsValidationFailed",
	}
	EmailAlreadyExists = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3003,
		Title:      "EmailAlreadyExists",
	}
	UsernameAlreadyExists = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3004,
		Title:      "UsernameAlreadyExists",
	}
	EmailAndUsernameAlreadyExists = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3005,
		Title:      "EmailAndUsernameAlreadyExists",
	}
	DefaultRoleDoesNotExist = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       3006,
		Title:      "DefaultRoleDoesNotExist",
	}
	SignInDetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3101,
		Title:      "SignInDetailsInvalid",
	}
	SignInDetailsValidationFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3102,
		Title:      "SignInDetailsValidationFailed",
	}
	EmailDoesNotExist = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3103,
		Title:      "EmailDoesNotExist",
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

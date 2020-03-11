// Package errors implements custom error handling.
package errors

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// ErrorCode holds the error code information.
type ErrorCode struct {
	HTTPStatus int
	Code       int
	Title      string
	Errors     error
}

// Error Codes
var (
	// 202 Accepted
	MessageQueueFailed = ErrorCode{
		HTTPStatus: http.StatusAccepted,
		Code:       202001,
		Title:      "MessageQueueFailed",
	}

	// 400 Bad Request
	DetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       400001,
		Title:      "DetailsInvalid",
	}
	AuthorizationBearerTokenEmpty = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       400002,
		Title:      "AuthorizationBearerTokenEmpty",
	}

	// 401 Unauthorized
	PasswordInvalid = ErrorCode{
		HTTPStatus: http.StatusUnauthorized,
		Code:       401001,
		Title:      "PasswordInvalid",
	}
	JWTTokenInvalid = ErrorCode{
		HTTPStatus: http.StatusUnauthorized,
		Code:       401002,
		Title:      "JWTTokenInvalid",
	}
	RefreshTokenIsRevoked = ErrorCode{
		HTTPStatus: http.StatusUnauthorized,
		Code:       401003,
		Title:      "RefreshTokenIsRevoked",
	}
	UserIsNotActive = ErrorCode{
		HTTPStatus: http.StatusUnauthorized,
		Code:       401004,
		Title:      "UserIsNotActive",
	}
	UUIDTokenDoesNotMatch = ErrorCode{
		HTTPStatus: http.StatusUnauthorized,
		Code:       401005,
		Title:      "UUIDTokenDoesNotMatch",
	}
	UUIDTokenExpired = ErrorCode{
		HTTPStatus: http.StatusUnauthorized,
		Code:       401006,
		Title:      "UUIDTokenExpired",
	}

	// 404 Not Found
	EmailDoesNotExist = ErrorCode{
		HTTPStatus: http.StatusNotFound,
		Code:       404001,
		Title:      "EmailDoesNotExist",
	}
	UserDoesNotExist = ErrorCode{
		HTTPStatus: http.StatusNotFound,
		Code:       404002,
		Title:      "UserDoesNotExist",
	}

	// 409 Conflict
	EmailAlreadyExists = ErrorCode{
		HTTPStatus: http.StatusConflict,
		Code:       409001,
		Title:      "EmailAlreadyExists",
	}
	UsernameAlreadyExists = ErrorCode{
		HTTPStatus: http.StatusConflict,
		Code:       409002,
		Title:      "UsernameAlreadyExists",
	}
	EmailAndUsernameAlreadyExists = ErrorCode{
		HTTPStatus: http.StatusConflict,
		Code:       409003,
		Title:      "EmailAndUsernameAlreadyExists",
	}
	EmailAlreadyConfirmed = ErrorCode{
		HTTPStatus: http.StatusConflict,
		Code:       409004,
		Title:      "EmailAlreadyConfirmed",
	}

	// 500 Internal Server Error
	DatabaseConnectionFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       500001,
		Title:      "DatabaseConnectionFailed",
	}
	DatabaseQueryFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       500002,
		Title:      "DatabaseQueryFailed",
	}
	PasswordGenerationFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       500003,
		Title:      "PasswordGenerationFailed",
	}
	AccessTokenIssueFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       500004,
		Title:      "AccessTokenIssueFailed",
	}
	RefreshTokenIssueFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       500005,
		Title:      "RefreshTokenIssueFailed",
	}
	DefaultRoleAssignFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       500006,
		Title:      "DefaultRoleAssignFailed",
	}
)

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("[%d] %s - %v", e.Code, e.Title, e.Errors)
}

// OmitDetailsInProd hides the error details when in production.
func (e *ErrorCode) OmitDetailsInProd() *ErrorCode {
	if viper.GetString("environment") == "production" {
		e.Errors = nil
	}
	return e
}

// New allows for an error code to be created.
func New(errorCode ErrorCode, errors error) *ErrorCode {
	return &ErrorCode{
		HTTPStatus: errorCode.HTTPStatus,
		Code:       errorCode.Code,
		Title:      errorCode.Title,
		Errors:     errors,
	}
}

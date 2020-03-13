// Package errors implements custom error handling.
package errors

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// ErrorCode holds the error code information.
type ErrorCode struct {
	HTTPStatus  int
	Code        int
	Title       string
	Description string
	Errors      error
}

// Error Codes
var (
	// 202 Accepted
	MessageQueueFailed = ErrorCode{
		HTTPStatus:  http.StatusAccepted,
		Code:        202001,
		Title:       "MessageQueueFailed",
		Description: "Message queue failed.",
	}

	// 400 Bad Request
	DetailsInvalid = ErrorCode{
		HTTPStatus:  http.StatusBadRequest,
		Code:        400001,
		Title:       "DetailsInvalid",
		Description: "Details invalid.",
	}
	AuthorizationBearerTokenEmpty = ErrorCode{
		HTTPStatus:  http.StatusBadRequest,
		Code:        400002,
		Title:       "AuthorizationBearerTokenEmpty",
		Description: "Authorization bearer token empty.",
	}

	// 401 Unauthorized
	PasswordInvalid = ErrorCode{
		HTTPStatus:  http.StatusUnauthorized,
		Code:        401001,
		Title:       "PasswordInvalid",
		Description: "Password invalid.",
	}
	JWTTokenInvalid = ErrorCode{
		HTTPStatus:  http.StatusUnauthorized,
		Code:        401002,
		Title:       "JWTTokenInvalid",
		Description: "Token invalid.",
	}
	RefreshTokenIsRevoked = ErrorCode{
		HTTPStatus:  http.StatusUnauthorized,
		Code:        401003,
		Title:       "RefreshTokenIsRevoked",
		Description: "Refresh token is revoked.",
	}
	UserIsNotActive = ErrorCode{
		HTTPStatus:  http.StatusUnauthorized,
		Code:        401004,
		Title:       "UserIsNotActive",
		Description: "User is not active.",
	}
	UUIDTokenDoesNotMatch = ErrorCode{
		HTTPStatus:  http.StatusUnauthorized,
		Code:        401005,
		Title:       "UUIDTokenDoesNotMatch",
		Description: "Token does not match.",
	}
	UUIDTokenExpired = ErrorCode{
		HTTPStatus:  http.StatusUnauthorized,
		Code:        401006,
		Title:       "UUIDTokenExpired",
		Description: "Token expired.",
	}

	// 404 Not Found
	EmailDoesNotExist = ErrorCode{
		HTTPStatus:  http.StatusNotFound,
		Code:        404001,
		Title:       "EmailDoesNotExist",
		Description: "Email does not exist.",
	}
	UserDoesNotExist = ErrorCode{
		HTTPStatus:  http.StatusNotFound,
		Code:        404002,
		Title:       "UserDoesNotExist",
		Description: "User does not exist.",
	}

	// 409 Conflict
	EmailAlreadyExists = ErrorCode{
		HTTPStatus:  http.StatusConflict,
		Code:        409001,
		Title:       "EmailAlreadyExists",
		Description: "Email already exists.",
	}
	UsernameAlreadyExists = ErrorCode{
		HTTPStatus:  http.StatusConflict,
		Code:        409002,
		Title:       "UsernameAlreadyExists",
		Description: "Username already exists.",
	}
	EmailAndUsernameAlreadyExists = ErrorCode{
		HTTPStatus:  http.StatusConflict,
		Code:        409003,
		Title:       "EmailAndUsernameAlreadyExists",
		Description: "Email and username already exists.",
	}
	EmailAlreadyConfirmed = ErrorCode{
		HTTPStatus:  http.StatusConflict,
		Code:        409004,
		Title:       "EmailAlreadyConfirmed",
		Description: "Email already confirmed.",
	}

	// 500 Internal Server Error
	DatabaseConnectionFailed = ErrorCode{
		HTTPStatus:  http.StatusInternalServerError,
		Code:        500001,
		Title:       "DatabaseConnectionFailed",
		Description: "Database connection failed.",
	}
	DatabaseQueryFailed = ErrorCode{
		HTTPStatus:  http.StatusInternalServerError,
		Code:        500002,
		Title:       "DatabaseQueryFailed",
		Description: "Database query failed.",
	}
	PasswordGenerationFailed = ErrorCode{
		HTTPStatus:  http.StatusInternalServerError,
		Code:        500003,
		Title:       "PasswordGenerationFailed",
		Description: "Password generation failed.",
	}
	AccessTokenIssueFailed = ErrorCode{
		HTTPStatus:  http.StatusInternalServerError,
		Code:        500004,
		Title:       "AccessTokenIssueFailed",
		Description: "Access token issue failed.",
	}
	RefreshTokenIssueFailed = ErrorCode{
		HTTPStatus:  http.StatusInternalServerError,
		Code:        500005,
		Title:       "RefreshTokenIssueFailed",
		Description: "Refresh token issue failed.",
	}
	DefaultRoleAssignFailed = ErrorCode{
		HTTPStatus:  http.StatusInternalServerError,
		Code:        500006,
		Title:       "DefaultRoleAssignFailed",
		Description: "Default role assign failed.",
	}
)

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("[%d] %s - %v", e.Code, e.Title, e.Errors)
}

// OmitDetailsInProd hides the error details when in production.
func (e *ErrorCode) OmitDetailsInProd() *ErrorCode {
	if viper.GetString("environment") == "production" && e.Code != 400001 {
		e.Errors = nil
	}
	return e
}

// New allows for an error code to be created.
func New(errorCode ErrorCode, errors error) *ErrorCode {
	return &ErrorCode{
		HTTPStatus:  errorCode.HTTPStatus,
		Code:        errorCode.Code,
		Title:       errorCode.Title,
		Description: errorCode.Description,
		Errors:      errors,
	}
}

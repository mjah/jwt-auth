package errors

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
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
	AccessTokenClaimsParseFailed = ErrorCode{
		HTTPStatus: http.StatusUnauthorized,
		Code:       2104,
		Title:      "AccessTokenClaimsParseFailed",
	}
	RefreshTokenClaimsParseFailed = ErrorCode{
		HTTPStatus: http.StatusUnauthorized,
		Code:       2105,
		Title:      "RefreshTokenClaimsParseFailed",
	}
	TokenValidationFailed = ErrorCode{
		HTTPStatus: http.StatusInternalServerError,
		Code:       2106,
		Title:      "TokenValidationFailed",
	}
	AuthorizationBearerTokenEmpty = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       2102,
		Title:      "AuthorizationBearerTokenEmpty",
	}
	RefreshTokenIsRevoked = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       2107,
		Title:      "RefreshTokenIsRevoked",
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
	UserDoesNotExist = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3104,
		Title:      "UserDoesNotExist",
	}
	UserIsNotActive = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3105,
		Title:      "UserIsNotActive",
	}
	UpdateDetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3106,
		Title:      "UpdateDetailsInvalid",
	}
	UpdateDetailsValidationFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3107,
		Title:      "UpdateDetailsValidationFailed",
	}
	ResetPasswordDetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3108,
		Title:      "ResetPasswordDetailsInvalid",
	}
	ResetPasswordDetailsValidationFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3109,
		Title:      "ResetPasswordDetailsValidationFailed",
	}
	SendResetPasswordEmailDetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3110,
		Title:      "SendResetPasswordEmailDetailsInvalid",
	}
	SendResetPasswordEmailDetailsValidationFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3111,
		Title:      "SendResetPasswordEmailDetailsValidationFailed",
	}
	ConfirmDetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3201,
		Title:      "ConfirmDetailsInvalid",
	}
	ConfirmDetailsValidationFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3202,
		Title:      "ConfirmDetailsValidationFailed",
	}
	SendConfirmEmailDetailsInvalid = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3210,
		Title:      "SendResetPasswordEmailDetailsInvalid",
	}
	SendConfirmEmailDetailsValidationFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3211,
		Title:      "SendResetPasswordEmailDetailsValidationFailed",
	}
	UserAlreadyConfirmed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3203,
		Title:      "UserAlreadyConfirmed",
	}
	ConfirmTokenDoesNotMatch = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3204,
		Title:      "ConfirmTokenDoesNotMatch",
	}
	ConfirmTokenExpired = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3205,
		Title:      "ConfirmTokenExpired",
	}
	ResetPasswordTokenDoesNotMatch = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3301,
		Title:      "ResetPasswordTokenDoesNotMatch",
	}
	ResetPasswordTokenExpired = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3302,
		Title:      "ResetPasswordTokenExpired",
	}
	MessageQueueFailed = ErrorCode{
		HTTPStatus: http.StatusBadRequest,
		Code:       3303,
		Title:      "MessageQueueFailed",
	}
)

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("[%d] %s - %v", e.Code, e.Title, e.Errors)
}

// OmitDetailsInProd ...
func (e *ErrorCode) OmitDetailsInProd() *ErrorCode {
	if viper.GetString("environment") == "production" {
		e.Errors = nil
	}
	return e
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

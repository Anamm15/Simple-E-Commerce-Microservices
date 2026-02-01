package constant

import "errors"

const (
	// Authentication Success
	MsgLoginSuccess    = "Login successful"
	MsgRegisterSuccess = "Account registered successfully"
	MsgLogoutSuccess   = "Successfully logged out"
	MsgTokenRefreshed  = "Token refreshed successfully"

	// Authentication Errors
	MsgInvalidCredentials = "Invalid email or password"
	MsgTokenMissing       = "Authorization token is required"
	MsgTokenInvalid       = "Invalid or expired token"
	MsgTokenExpired       = "Your session has expired, please login again"
	MsgEmailAlreadyExists = "Email address is already registered"

	// Authorization Errors
	MsgInsufficientPermissions = "You do not have the required role to access this resource"
)

var (
	ErrInvalidCredentials = errors.New(MsgInvalidCredentials)
	ErrTokenMissing       = errors.New(MsgTokenMissing)
	ErrTokenInvalid       = errors.New(MsgTokenInvalid)
	ErrTokenExpired       = errors.New(MsgTokenExpired)
	ErrEmailAlreadyExists = errors.New(MsgEmailAlreadyExists)
)

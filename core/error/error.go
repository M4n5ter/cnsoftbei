package error

import "fmt"

type Code = int

type DatabaseError struct {
	Message string
	Code    Code
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error: %s (code: %d)", e.Message, e.Code)
}

const (
	// Database errors
	ErrDatabaseConnection Code = 1000
	ErrDatabaseQuery      Code = 1001
	ErrDatabaseNotFound   Code = 1002
	ErrDatabaseDuplicate  Code = 1003

	// Validation errors
	ErrValidation Code = 2000

	// Authentication errors
	ErrAuthentication Code = 3000
	ErrUnauthorized   Code = 3001
	ErrForbidden      Code = 3002
	ErrTokenExpired   Code = 3003
	ErrTokenInvalid   Code = 3004
	ErrTokenMalformed Code = 3005

	// Internal errors
	ErrInternal                   Code = 4000
	ErrNotFound                   Code = 4001
	ErrUnknown                    Code = 4002
	ErrPanic                      Code = 4003
	ErrTimeout                    Code = 4004
	ErrCanceled                   Code = 4005
	ErrFailed                     Code = 4006
	ErrNotImplemented             Code = 4007
	ErrNotSupported               Code = 4008
	ErrNotReady                   Code = 4009
	ErrAlreadyExists              Code = 4010
	ErrNotAvailable               Code = 4011
	ErrTooManyRequests            Code = 4012
	ErrServiceUnavailable         Code = 4013
	ErrGatewayTimeout             Code = 4014
	ErrBadGateway                 Code = 4015
	ErrUnavailableForLegalReasons Code = 4016
	ErrInternalServerError        Code = 4017
	ErrNotAcceptable              Code = 4018

	// External errors
	ErrExternal               Code = 5000
	ErrExternalService        Code = 5001
	ErrExternalUnavailable    Code = 5002
	ErrExternalTimeout        Code = 5003
	ErrExternalFailed         Code = 5004
	ErrExternalNotImplemented Code = 5005
	ErrExternalNotSupported   Code = 5006
	ErrExternalNotReady       Code = 5007
	ErrExternalAlreadyExists  Code = 5008
)

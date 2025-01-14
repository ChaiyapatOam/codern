package domain

import (
	"errors"
	"fmt"
)

const (
	ErrInternal   = 1
	ErrRoute      = 2
	ErrFileSystem = 3

	ErrLoggingError    = 1000
	ErrAuthHeader      = 1010
	ErrValidation      = 1011
	ErrBodyParser      = 1012
	ErrBodyValidator   = 1013
	ErrQueryParser     = 1014
	ErrQueryValidator  = 1015
	ErrParamsParser    = 1016
	ErrParamsValidator = 1017

	ErrSessionPrefix     = 2000
	ErrSignatureMismatch = 2001
	ErrInvalidSession    = 2002
	ErrSessionExpired    = 2003
	ErrInvalidEmail      = 2010
	ErrDupEmail          = 2011
	ErrUserData          = 2022
	ErrUserPassword      = 2023

	ErrWorkspaceNotFound   = 3000
	ErrWorkspaceNoPerm     = 3001
	ErrWorkspacePermFailed = 3002
	ErrCreateSubmission    = 3010
	ErrGetSubmission       = 3011
	ErrGetAssignment       = 3012
)

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func NewErrorWithData(code int, message string, data interface{}) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewErrorf(code int, message string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(message, args...),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("(code %d) %s", e.Code, e.Message)
}

func HasErrorCode(err error, code int) bool {
	var domainError *Error
	if errors.As(err, &domainError) {
		return domainError.Code == code
	}
	return false
}

type ValidationError struct {
	namespace string
	Field     string `json:"field"`
	Type      string `json:"type"`
}

func NewValidationError(namespace string, field string, keyType string) *ValidationError {
	return &ValidationError{
		namespace: namespace,
		Field:     field,
		Type:      keyType,
	}
}

func (e *ValidationError) Code() int {
	return ErrValidation
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("payload: %s %s %s", e.namespace, e.Field, e.Type)
}

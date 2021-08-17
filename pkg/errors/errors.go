package errors

import (
	"errors"
	"fmt"
	libErrors "github.com/pkg/errors"
)

const (
	// NoType error
	NoType = ErrorType(iota)
	// BadRequest error
	BadRequest
	// NotFound error
	NotFound
	AliTTSToken
)

type ErrorType uint

type customError struct {
	errorType     ErrorType
	originalError error
	context       errorContext
}

type errorContext struct {
	Field   string
	Message string
}

// New 创建一个customError
func (errorType ErrorType) New(msg string) error {
	return customError{errorType: errorType, originalError: errors.New(msg)}
}

// Newf 用formatted message创建一个customError
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

// Wrap 创建一个wrapped error
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrapf(err, msg)
}

// Wrapf 用formatted message创建一个customError
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: libErrors.Wrapf(err, msg, args...)}
}

// Error 返回customError的错误信息
func (custom customError) Error() string {
	return custom.originalError.Error()
}

func (custom customError) Cause() error {
	return custom.originalError
}

// New 创建一个 no type error
func New(msg string) error {
	return customError{
		errorType:     NoType,
		originalError: errors.New(msg),
	}
}

// Newf 用 formatted message 创建一个no type error
func Newf(msg string, args ...interface{}) error {
	return customError{errorType: NoType, originalError: errors.New(fmt.Sprintf(msg, args...))}
}

// Cause 返回原始错误
func Cause(err error) error {
	return libErrors.Cause(err)
}

func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := libErrors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			context:       customErr.context,
		}
	}

	return customError{errorType: NoType, originalError: wrappedError}
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(customError); ok {
		return customError{errorType: customErr.errorType, originalError: customErr.originalError, context: context}
	}

	return customError{errorType: NoType, originalError: err, context: context}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if customErr, ok := err.(customError); ok || customErr.context != emptyContext {

		return map[string]string{"field": customErr.context.Field, "message": customErr.context.Message}
	}

	return nil
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}

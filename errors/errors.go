package errors

import (
	stdErrors "errors"
	"net/http"

	"github.com/nongod/common/code"
)

// New creates a custom error with additional attributes.
func New(message string, opts ...OptionFunc) error {
	err := &customError{
		message:    message,
		attributes: make(map[attributeKey]any),
	}

	return err.with(opts...)
}

// Wrap wraps an existing error with additional attributes or a new message.
func Wrap(err error, opts ...OptionFunc) error {
	if err == nil {
		return (&customError{}).with(opts...)
	}

	wrappedErr := &customError{
		cause:      err,
		message:    err.Error(),
		attributes: make(map[attributeKey]any),
	}

	return wrappedErr.with(opts...)
}

// Cause returns the root cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//	type wrapper interface {
//	    Unwrap() error
//	}
//
// If the error does not implement Unwrap, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	for err != nil {
		cause := stdErrors.Unwrap(err)
		if cause == nil {
			break
		}
		err = cause
	}

	return err
}

// FindErrorCode retrieves the error code from an error implementing the code.Provider interface.
// If the error does not implement the interface, it returns NilCode.
func FindErrorCode(err error) code.Code {
	if errWithCode, ok := err.(code.Provider); ok {
		return errWithCode.Code()
	}

	return code.NilCode
}

// GetHTTPStatus retrieves the HTTP status code from an error implementing the httpStatusCoder interface.
// If the error does not implement the interface, it returns http.StatusInternalServerError.
func GetHTTPStatus(err error) int {
	if errWithHTTPStatus, ok := err.(httpStatusProvider); ok {
		return errWithHTTPStatus.httpStatus()
	}

	return http.StatusInternalServerError
}

var _ error = (*customError)(nil)
var _ code.Provider = (*customError)(nil)
var _ httpStatusProvider = (*customError)(nil)

// customError is a custom error implementation with support for attributes like error code and HTTP status.
type customError struct {
	cause      error
	message    string
	attributes map[attributeKey]any // Additional attributes (e.g., error code, HTTP status code).
}

// Error returns the error message.
func (e *customError) Error() string {
	return e.message
}

// Unwrap retrieves the underlying error cause, if any.
func (e *customError) Unwrap() error {
	return e.cause
}

// with applies a list of OptionFunc to the customError, modifying its attributes or message.
func (e *customError) with(opts ...OptionFunc) *customError {
	for _, opt := range opts {
		opt(e)
	}

	return e
}

// setCode sets the error code attribute.
func (e *customError) setCode(code code.Code) {
	e.attributes[attributeKeyCode] = code
}

// Code retrieves the error code. If none is set, it returns NilCode.
func (e *customError) Code() code.Code {
	if errorCode, ok := e.attributes[attributeKeyCode].(code.Code); ok {
		return errorCode
	}

	return code.NilCode
}

// setHTTPStatus sets the HTTP status code.
func (e *customError) setHTTPStatus(httpStatus int) {
	e.attributes[attributeKeyHTTPStatus] = httpStatus
}

// httpStatus retrieves the HTTP status code. If none is set, it returns http.StatusInternalServerError.
func (e *customError) httpStatus() int {
	if httpStatus, ok := e.attributes[attributeKeyHTTPStatus].(int); ok {
		return httpStatus
	}

	return http.StatusInternalServerError
}

// httpStatusCoder is an interface for errors with an HTTP status code.
type httpStatusProvider interface {
	httpStatus() int
}

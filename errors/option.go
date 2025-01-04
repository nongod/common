package errors

import "github.com/nongod/common/code"

// OptionFunc defines a function that modifies a customError.
type OptionFunc func(*customError)

// WithMessage sets a custom message for the error.
func WithMessage(message string) OptionFunc {
	return func(e *customError) {
		e.message = message
	}
}

// WithCode sets a custom error code for the error.
func WithCode(code code.Code) OptionFunc {
	return func(e *customError) {
		e.setCode(code)
	}
}

// WithHTTPStatus sets an HTTP status code for the error.
func WithHTTPStatus(httpStatus int) OptionFunc {
	return func(e *customError) {
		e.setHTTPStatus(httpStatus)
	}
}

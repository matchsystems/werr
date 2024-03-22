package werr

import (
	"errors"
)

type UnwrapErr interface {
	Unwrap() error
}

var _ UnwrapErr = (*wrapError)(nil)

type wrapError struct {
	caller   string
	err      error
	funcName string
	msg      string
}

func newError(err error, msg string) error {
	sourceCaller, funcName := caller(defaultCallerSkip)

	return wrapError{
		caller:   sourceCaller,
		err:      err,
		msg:      msg,
		funcName: funcName,
	}
}

func Message(err error) string {
	var wErr wrapError
	if errors.As(err, &wErr) {
		return wErr.msg
	}

	return ""
}

// UnwrapMessageFunc recursively walks through wrapped errors and returns messages.
// If the input error (err) is not a wrapped error, it is returned empty string.
// Example:
//
//  1. returns latest message:
//
//     UnwrapMessageFunc(err, func(msg string) bool) {
//     return true
//     }
//
//  2. returns latest not empty message:
//
//     UnwrapMessageFunc(err, func(msg string) bool) {
//     return len(msg) != 0
//     }
func UnwrapMessageFunc(err error, fn func(msg string) bool) string {
	var msg string

	var wErr wrapError
	for errors.As(err, &wErr) {
		if fn(Message(err)) {
			msg = Message(err)
		}
		err = wErr.Unwrap()
	}

	return msg
}

// Unwrap recursively traverses the wrapped errors and returns the innermost non-wrapped error.
// If the input error (err) is not a wrapped error, it is returned unchanged.
func Unwrap(err error) error {
	var wErr wrapError
	for errors.As(err, &wErr) {
		err = wErr.Unwrap()
	}

	return err
}

// UnwrapAll recursively traverses the wrapped errors and returns the innermost non-wrapped error.
// If the input error (err) is not a wrapped error, it is returned unchanged.
func UnwrapAll(err error) error {
	u, ok := err.(UnwrapErr)
	if ok {
		return UnwrapAll(u.Unwrap())
	}

	return err
}

func (e wrapError) Error() string {
	return ErrorStackMarshaler(e.caller, e.err, e.funcName, e.msg)
}

func (e wrapError) Message() string {
	return e.msg
}

func (e wrapError) Unwrap() error {
	return e.err
}

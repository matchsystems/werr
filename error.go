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

// Unwrap recursively traverses the wrapped errors and returns the innermost non-wrapped error.
// If the input error (err) is not a wrapped error, it is returned unchanged.
func Unwrap(err error) error {
	var wErr wrapError
	for errors.As(err, &wErr) {
		err = wErr.Unwrap()
	}

	return err
}

func (e wrapError) Error() string {
	if e.msg != "" {
		e.msg = "\t" + e.msg
	}

	return e.caller + "\t" + e.funcName + e.msg + "\n" + e.err.Error()
}

func (e wrapError) Unwrap() error {
	return e.err
}

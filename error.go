package werr

import (
	"errors"

	"gitlab.com/matchsystems-golang/stacktrace"
)

type UnwrapErr interface {
	Unwrap() error
}

var _ UnwrapErr = (*wrapError)(nil)

type wrapError struct {
	err    error
	msg    string
	frames stacktrace.Frames
}

const defaultCallerSkip = 4

func newError(err error, msg string) error {
	return wrapError{
		err:    err,
		msg:    msg,
		frames: stacktrace.GetStacktrace(defaultCallerSkip, 1),
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

// UnwrapAll recursively traverses the wrapped errors and returns the innermost non-wrapped error.
// If the input error (err) is not a wrapped error, it is returned unchanged.
func UnwrapAll(err error) error {
	u, ok := err.(UnwrapErr) //nolint:errorlint // unwrap
	if ok {
		return UnwrapAll(u.Unwrap())
	}

	return err
}

func (e wrapError) Error() string {
	return ErrorStackMarshaler(
		e.err,
		e.msg,
		e.frames,
	)
}

func (e wrapError) Unwrap() error {
	return e.err
}

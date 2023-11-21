package werr

import (
	"fmt"
)

// Wrap takes an error and returns a new wrapped error.
// If the input error (err) is nil, the function returns nil.
// Otherwise, it creates a new wrapped error using the input error
// and an empty message text.
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	return newError(err, "")
}

// Wrapf takes an error, a format string, and optional arguments, and returns a new wrapped error.
// If the input error (err) is nil, the function returns nil.
// Otherwise, it creates a new wrapped error using the input error
// and formats the message text based on the provided format and arguments.
func Wrapf(err error, format string, a ...any) error {
	if err == nil {
		return nil
	}

	return newError(err, fmt.Sprintf(format, a...))
}

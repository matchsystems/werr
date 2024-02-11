package werr

import (
	"strings"

	"github.com/matchsystems/stacktrace"
)

//nolint:gochecknoglobals // for custom setting
var (
	// ErrorStackMarshaler extract the stack from err.
	ErrorStackMarshaler ErrorStackMarshalerFn = DefaultErrorStackMarshaler
)

type ErrorStackMarshalerFn func(err error, msg string, frames stacktrace.Frames) string

func DefaultErrorStackMarshaler(err error, msg string, frames stacktrace.Frames) string {
	var result string

	if len(msg) > 0 {
		result = msg + "\n"
	}

	result += Unwrap(err).Error()

	if len(frames) > 0 {
		result = result + "\n" + strings.Join(frames.Pretty(), "\n")
	}

	return result
}

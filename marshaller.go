package werr

import (
	"strconv"
)

//nolint:gochecknoglobals // for custom setting
var (
	// ErrorStackMarshaler extract the stack from err.
	ErrorStackMarshaler ErrorStackMarshalerFn = DefaultErrorStackMarshaler
)

type ErrorStackMarshalerFn func(err error, caller, funcName, msg string, line int) string

func DefaultErrorStackMarshaler(err error, caller, funcName, msg string, line int) string {
	if msg != "" {
		msg = "\t" + msg
	}

	return caller + ":" + strconv.Itoa(line) + "\t" + funcName + msg + "\n" + err.Error()
}

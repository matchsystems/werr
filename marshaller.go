package werr

//nolint:gochecknoglobals // for custom setting
var (
	// ErrorStackMarshaler extract the stack from err.
	ErrorStackMarshaler ErrorStackMarshalerFn = DefaultErrorStackMarshaler
)

type ErrorStackMarshalerFn func(caller string, err error, funcName, msg string) string

func DefaultErrorStackMarshaler(caller string, err error, funcName, msg string) string {
	if msg != "" {
		msg = "\t" + msg
	}

	return caller + "\t" + funcName + msg + "\n" + err.Error()
}

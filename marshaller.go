package werr

var (
	// ErrorStackMarshaler extract the stack from err
	ErrorStackMarshaler func(caller string, err error, funcName, msg string) string = func(caller string, err error, funcName, msg string) string {
		if msg != "" {
			msg = "\t" + msg
		}

		return caller + "\t" + funcName + msg + "\n" + err.Error()
	}
)

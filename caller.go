package werr

import (
	"path"
	"runtime"
	"strings"
)

const defaultCallerSkip = 3

func caller(skip int) (string, string, int) {
	// make caller func invisible
	if skip < 1 {
		skip = 1
	}

	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", "", -1
	}

	fnName := runtime.FuncForPC(pc).Name()
	ix := strings.LastIndex(fnName, ".")
	sourceCaller := fnName[0:ix] + "/" + path.Base(file)

	if len(fnName) > ix {
		return sourceCaller, fnName[ix+1:] + "()", line
	}

	return sourceCaller, "", line
}

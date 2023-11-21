package werr

import (
	"path"
	"runtime"
	"strconv"
	"strings"
)

const defaultCallerSkip = 3

func caller(skip int) (string, string) {
	// make caller func invisible
	if skip < 1 {
		skip = 1
	}

	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", ""
	}

	fnName := runtime.FuncForPC(pc).Name()
	ix := strings.LastIndex(fnName, ".")
	sourceCaller := fnName[0:ix] + "/" + path.Base(file) + ":" + strconv.Itoa(line)

	if len(fnName) > ix {
		return sourceCaller, fnName[ix+1:] + "()"
	}

	return sourceCaller, ""
}

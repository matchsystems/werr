package werr_test

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	"github.com/matchsystems/stacktrace"
	"github.com/stretchr/testify/require"

	"github.com/matchsystems/werr"
)

func callerErr() error {
	return werr.Wrapf(fmt.Errorf("hello world"), "description = %d", 1)
}

func caller() []uintptr {
	pcs := make([]uintptr, 10)
	runtime.Callers(0, pcs)

	return pcs
}

func nested() []uintptr {
	return caller()
}

func TestDefaultErrorStackMarshaler(t *testing.T) {
	t.Parallel()

	t.Run("deep pretty print", func(t *testing.T) {
		t.Parallel()

		pcs := nested()
		require.NotEmpty(t, pcs)
		frames := stacktrace.ExtractFrames(pcs)
		require.NotEmpty(t, frames)
		callersFrames := werr.DefaultErrorStackMarshaler(errors.New("test"), "hello", frames)

		require.Equal(t, `hello
test
werr/marshaller_test.go:21#caller
werr/marshaller_test.go:27#nested
werr/marshaller_test.go:36#TestDefaultErrorStackMarshaler.func1`, callersFrames)
	})

	t.Run("pretty print", func(t *testing.T) {
		t.Parallel()

		pcs := caller()
		require.NotEmpty(t, pcs)
		frames := stacktrace.ExtractFrames(pcs)
		require.NotEmpty(t, frames)
		callersFrames := werr.DefaultErrorStackMarshaler(errors.New("test"), "hello", frames)

		require.Equal(t, `hello
test
werr/marshaller_test.go:21#caller
werr/marshaller_test.go:52#TestDefaultErrorStackMarshaler.func2`, callersFrames)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		err := callerErr()

		require.Equal(t, `description = 1
hello world
werr/marshaller_test.go:16#callerErr`, err.Error())
	})
}

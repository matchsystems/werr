package werr_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/matchsystems/werr"
)

func TestDefaultErrorStackMarshaler(t *testing.T) {
	t.Parallel()

	caller := "main.go"
	err := errors.New("example error")
	funcName := "TestFunction"
	msg := "custom message"
	line := 123

	result := werr.DefaultErrorStackMarshaler(err, caller, funcName, msg, line)

	require.Contains(t, result, caller)
	require.Contains(t, result, funcName)
	require.Contains(t, result, msg)
	require.Contains(t, result, err.Error())
	require.Contains(t, result, "123")
}

package werr_test

import (
	"errors"
	"testing"

	"github.com/matchsystems/werr"
	"github.com/stretchr/testify/require"
)

func TestDefaultErrorStackMarshaler(t *testing.T) {
	t.Parallel()

	caller := "main.go"
	err := errors.New("example error")
	funcName := "TestFunction"
	msg := "custom message"

	result := werr.DefaultErrorStackMarshaler(caller, err, funcName, msg)

	require.Contains(t, result, caller)
	require.Contains(t, result, funcName)
	require.Contains(t, result, msg)
	require.Contains(t, result, err.Error())
}

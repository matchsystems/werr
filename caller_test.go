package werr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func c() (string, string) { return caller(0) }
func b() (string, string) { return c() }
func a() (string, string) { return b() }

var varCaller, varFunc = caller(0)

func TestCaller(t *testing.T) {
	t.Parallel()

	t.Run("when called in a function", func(t *testing.T) {
		t.Parallel()

		sourceCaller, funcName := a()
		require.Equal(t, "github.com/matchsystems/werr/caller_test.go:9", sourceCaller)
		require.Equal(t, "c()", funcName)
	})

	t.Run("when called from outside", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, "github.com/matchsystems/werr/caller_test.go:13", varCaller)
		require.Equal(t, "init()", varFunc)
	})
}

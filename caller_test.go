package werr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func c() (string, string, int) { return caller(0) }
func b() (string, string, int) { return c() }
func a() (string, string, int) { return b() }

var varCaller, varFunc, varLine = caller(0)

func TestCaller(t *testing.T) {
	t.Parallel()

	t.Run("when called in a function", func(t *testing.T) {
		t.Parallel()

		sourceCaller, funcName, line := a()
		require.Equal(t, "github.com/matchsystems/werr/caller_test.go", sourceCaller)
		require.Equal(t, "c()", funcName)
		require.Equal(t, 9, line)
	})

	t.Run("when called from outside", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, "github.com/matchsystems/werr/caller_test.go", varCaller)
		require.Equal(t, "init()", varFunc)
		require.Equal(t, 13, varLine)
	})
}

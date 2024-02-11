package werr

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func nestedThirdErr() error  { return fmt.Errorf("example nested error") }
func nestedSecondErr() error { return Wrap(nestedThirdErr()) }
func nestedErr() error       { return Wrap(nestedSecondErr()) }

func TestWrapError_Error(t *testing.T) {
	t.Parallel()

	t.Run("with message", func(t *testing.T) {
		t.Parallel()
		wrappedErr := Wrapf(errors.New("original error"), "additional message")
		require.Equal(t, `additional message
original error
werr/error_test.go:20#TestWrapError_Error.func1`, wrappedErr.Error())
	})

	t.Run("when wrap chain", func(t *testing.T) {
		t.Parallel()
		subWrappedErr := Wrap(errors.New("original error"))
		wrappedErr := Wrapf(subWrappedErr, "additional message")
		require.Equal(t, `additional message
original error
werr/error_test.go:29#TestWrapError_Error.func2
werr/error_test.go:28#TestWrapError_Error.func2`, wrappedErr.Error())
	})

	t.Run("without message", func(t *testing.T) {
		t.Parallel()
		wrappedErr := wrapError{err: errors.New("original error")}
		require.Equal(t, `original error`, wrappedErr.Error())
	})
}

func TestUnwrap(t *testing.T) {
	t.Parallel()

	t.Run("when wrap chain", func(t *testing.T) {
		t.Parallel()

		err1 := errors.New("original error")
		err2 := fmt.Errorf("fmt wrap: %w", err1)
		err3 := Wrap(err2)
		err4 := Wrap(err3)

		require.EqualError(t, err2, Unwrap(err4).Error())
	})

	t.Run("unwrap all", func(t *testing.T) {
		t.Parallel()

		err1 := errors.New("original error")
		err2 := fmt.Errorf("fmt wrap: %w", err1)
		err3 := Wrap(err2)
		err4 := Wrap(err3)

		require.EqualError(t, err1, UnwrapAll(err4).Error())
	})

	t.Run("nested trace", func(t *testing.T) {
		t.Parallel()
		err := nestedErr()
		require.Equal(t, `example nested error
werr/error_test.go:13#nestedErr
werr/error_test.go:12#nestedSecondErr`, err.Error())
	})

	t.Run("when nil", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, Unwrap(nil))
	})

	t.Run("when without wrap", func(t *testing.T) {
		t.Parallel()
		errWithoutUnwrap := errors.New("error without Unwrap")
		require.ErrorIs(t, Unwrap(errWithoutUnwrap), errWithoutUnwrap)
	})
}

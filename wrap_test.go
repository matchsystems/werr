package werr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	t.Parallel()

	t.Run("with error", func(t *testing.T) {
		t.Parallel()
		originalErr := errors.New("original error")
		wrappedErr := Wrap(originalErr)

		// Ensure that the wrapped error is of type error
		require.IsType(t, wrapError{}, wrappedErr)

		// Ensure that the wrapped error contains the original error
		require.ErrorIs(t, originalErr, wrappedErr)
	})

	t.Run("when nil", func(t *testing.T) {
		t.Parallel()
		// Ensure that wrapping a nil error results in nil
		require.NoError(t, Wrap(nil))
	})

	t.Run("check skip count is sufficient", func(t *testing.T) {
		t.Parallel()
		originalErr := errors.New("original error")
		wrappedErr := Wrap(originalErr)

		// Ensure the skip count 3 is enough
		require.Equal(t, "github.com/matchsystems/werr.TestWrap/wrap_test.go:34\tfunc3()\noriginal error", wrappedErr.Error())
	})
}

func TestWrapf(t *testing.T) {
	t.Parallel()

	t.Run("with error", func(t *testing.T) {
		t.Parallel()
		originalErr := errors.New("original error")
		wrappedErr := Wrapf(originalErr, "additional message: %s", "some details")

		// Ensure that the wrapped error is of type *wrapError
		require.IsType(t, wrapError{}, wrappedErr)

		// Ensure that the wrapped error contains the original error
		require.ErrorIs(t, wrappedErr, originalErr)

		// Ensure that the wrapped error contains the additional message
		require.Contains(t, wrappedErr.Error(), "additional message: some details")
	})

	t.Run("when nil", func(t *testing.T) {
		t.Parallel()
		wrappedErr := Wrapf(nil, "additional message: %s", "some details")

		// Ensure that wrapping a nil error results in nil
		require.NoError(t, wrappedErr, "Expected wrapped error to be nil")
	})
}

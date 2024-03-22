package werr

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func nestedThirdErr() error  { return errors.New("example nested error") }
func nestedSecondErr() error { return Wrap(nestedThirdErr()) }
func nestedErr() error       { return Wrap(nestedSecondErr()) }

func TestWrapError_Error(t *testing.T) {
	t.Parallel()

	err := errors.New("original error")

	t.Run("with message", func(t *testing.T) {
		t.Parallel()
		wrappedErr := wrapError{
			caller:   "caller",
			err:      err,
			funcName: "function",
			msg:      "additional message",
		}

		expected := "caller\tfunction\tadditional message\noriginal error"
		require.Equal(t, expected, wrappedErr.Error())
	})

	t.Run("when wrap chain", func(t *testing.T) {
		t.Parallel()
		subWrappedErr := wrapError{
			caller:   "subCaller",
			err:      err,
			funcName: "subFunction",
			msg:      "",
		}
		wrappedErr := wrapError{
			caller:   "caller",
			err:      subWrappedErr,
			funcName: "function",
			msg:      "additional message",
		}

		expected := "caller\tfunction\tadditional message\nsubCaller\tsubFunction\noriginal error"
		require.Equal(t, expected, wrappedErr.Error())
	})

	t.Run("without message", func(t *testing.T) {
		t.Parallel()
		wrappedErr := wrapError{
			caller:   "caller",
			err:      err,
			funcName: "function",
		}

		expected := "caller\tfunction\noriginal error"
		require.Equal(t, expected, wrappedErr.Error())
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

		require.Equal(t, "github.com/matchsystems/werr/error_test.go:13\tnestedErr()\ngithub.com/matchsystems/werr/error_test.go:12\tnestedSecondErr()\nexample nested error", nestedErr().Error()) //nolint: lll // test
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

func TestMessage(t *testing.T) {
	t.Parallel()

	err := errors.New("foo")

	wErr1Msg := "bar"
	wErr1 := Wrapf(err, wErr1Msg)

	wErr2Msg := "baz"
	wErr2 := Wrapf(wErr1, wErr2Msg)

	wErr3 := Wrap(wErr1)

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		msg := Message(err)
		require.Empty(t, msg)

		msg = Message(wErr1)
		require.Equal(t, wErr1Msg, msg)

		msg = Message(wErr2)
		require.Equal(t, wErr2Msg, msg)
	})

	t.Run("without message", func(t *testing.T) {
		t.Parallel()

		msg := Message(wErr3)
		require.Empty(t, msg)
	})
}

func TestUnwrapFunc(t *testing.T) {
	t.Parallel()

	err := errors.New("foo")

	t.Run("latest with message", func(t *testing.T) {
		t.Parallel()

		wErr1Msg := "bar"
		wErr1 := Wrapf(err, wErr1Msg)

		wErr2Msg := "baz"
		wErr2 := Wrapf(wErr1, wErr2Msg)

		var msg string
		UnwrapFunc(wErr2, func(err error) {
			if len(Message(err)) != 0 {
				msg = Message(err)
			}
		})
		require.Equal(t, wErr1Msg, msg)
	})

	t.Run("latest without message", func(t *testing.T) {
		t.Parallel()

		wErr1 := Wrap(err)

		wErr2Msg := "baz"
		wErr2 := Wrapf(wErr1, wErr2Msg)

		var msg string
		UnwrapFunc(wErr2, func(err error) {
			msg = Message(err)
		})
		require.Empty(t, msg)
	})

	t.Run("latest not empty message", func(t *testing.T) {
		t.Parallel()

		wErr1 := Wrap(err)

		wErr2Msg := "baz"
		wErr2 := Wrapf(wErr1, wErr2Msg)

		var msg string
		UnwrapFunc(wErr1, func(err error) {
			if len(Message(err)) != 0 {
				msg = Message(err)
			}
		})
		require.Empty(t, msg)

		UnwrapFunc(wErr2, func(err error) {
			if len(Message(err)) != 0 {
				msg = Message(err)
			}
		})
		require.Equal(t, wErr2Msg, msg)
	})
}

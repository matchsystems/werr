package benchmark

import (
	"errors"
	"fmt"
	"testing"

	"github.com/joomcode/errorx"

	"github.com/matchsystems/werr"
)

var errSink error

func BenchmarkSimpleError10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		errSink = function0(10, createSimpleError)
	}
	consumeResult(errSink)
}

func BenchmarkWrapError10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		errSink = function0(10, createWrapError)
	}
	consumeResult(errSink)
}

func BenchmarkWrapMsgError10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		errSink = function0(10, createWrapMsgError)
	}
	consumeResult(errSink)
}

func BenchmarkErrorxError10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		errSink = function0(10, createErrorxError)
	}
	consumeResult(errSink)
}

func BenchmarkSimpleError100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		errSink = function0(100, createSimpleError)
	}
	consumeResult(errSink)
}

func BenchmarkWrapError100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		errSink = function0(100, createWrapError)
	}
	consumeResult(errSink)
}

func BenchmarkWrapMsgError100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		errSink = function0(100, createWrapMsgError)
	}
	consumeResult(errSink)
}

func BenchmarkErrorxError100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		errSink = function0(100, createErrorxError)
	}
	consumeResult(errSink)
}

func BenchmarkSimpleErrorPrint100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := function0(100, createSimpleError)
		emulateErrorPrint(err)
		errSink = err
	}
	consumeResult(errSink)
}

func BenchmarkWrapErrorPrint100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := function0(100, createWrapError)
		emulateErrorPrint(err)
		errSink = err
	}
	consumeResult(errSink)
}

func BenchmarkWrapMsgErrorPrint100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := function0(100, createWrapMsgError)
		emulateErrorPrint(err)
		errSink = err
	}
	consumeResult(errSink)
}

func BenchmarkErrorxErrorPrint100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := function0(100, createErrorxError)
		emulateErrorPrint(err)
		errSink = err
	}
	consumeResult(errSink)
}

func createSimpleError() error {
	return errors.New("benchmark")
}

func createWrapError() error {
	return werr.Wrap(errors.New("benchmark"))
}

func createWrapMsgError() error {
	return werr.Wrapf(errors.New("benchmark"), "benchmark")
}

var ErrorX = errorx.NewNamespace("errorx.benchmark").NewType("stack_trace")

func createErrorxError() error {
	return ErrorX.New("benchmark")
}

func function0(depth int, generate func() error) error {
	if depth == 0 {
		return generate()
	}

	switch depth % 3 {
	case 0:
		return function1(depth-1, generate)
	case 1:
		return function2(depth-1, generate)
	default:
		return function3(depth-1, generate)
	}
}

func function1(depth int, generate func() error) error {
	if depth == 0 {
		return generate()
	}

	return function4(depth-1, generate)
}

func function2(depth int, generate func() error) error {
	if depth == 0 {
		return generate()
	}

	return function4(depth-1, generate)
}

func function3(depth int, generate func() error) error {
	if depth == 0 {
		return generate()
	}

	return function4(depth-1, generate)
}

func function4(depth int, generate func() error) error {
	switch depth {
	case 0:
		return generate()
	default:
		return function0(depth-1, generate)
	}
}

type sinkError struct {
	value int
}

func (sinkError) Error() string {
	return ""
}

// Perform error formatting and consume the result to disallow optimizations against output.
func emulateErrorPrint(err error) {
	output := fmt.Sprintf("%+v", err)
	if len(output) > 10000 && output[1000:1004] == "DOOM" {
		panic("this was not supposed to happen")
	}
}

// Consume error with a possible side effect to disallow optimizations against err.
func consumeResult(err error) {
	if e, ok := err.(sinkError); ok && e.value == 1 { //nolint:errorlint // casting
		panic("this was not supposed to happen")
	}
}

// A public function to discourage optimizations against errSink variable.
func ExportSink() error {
	return errSink
}

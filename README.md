## Highlights

The **werr** library provides efficient error wrapping capabilities for Go. It is designed with a focus on performance
and enables the recording of functions where an error occurred. Here are some key features of the library:

* Recording information about functions that triggered the error for easy debugging.
* Support for custom error messages to make errors more informative.
* Very high performance compared to other error-wrapping libraries.

## Introduction

At first, you start with the basic:

```go
return err
```

Later, when the hunt for the elusive error becomes challenging, you enhance it with a description:

```go
return fmt.Errorf("uh oh! something terrible happened: %w", err)
```

As your needs grow beyond mere descriptions and you seek traces, you may resort to integrating external libraries:

```go
return errorx.Decorate(err, "this could be so much better")
```

Feeling the inconvenience and slowdown, you finally discover the simplicity of:

```go
return werr.Wrap(err)
```

# werr

To use **werr** in your Go application, simply import it in your code:

```go
import "github.com/matchsystems/werr"
```

## Features

* **Error Creation**: Create errors just like before with `errors.New("error message")` and use them seamlessly.
* **Error Wrapping**: Wrap an existing error using `werr.Wrap(err)`.
* **Custom Messages**: Wrap errors with custom messages using `werr.Wrapf(err, "my error")`
  or `werr.Wrapf(err, "custom error message: %s", "details")`.
* **Error Wrapping**: Wrap existing errors using `werr.Wrap(err)`.
* **Error Unwrapping**: Unwrap an error using `werr.Unwrap(err)`
* **Informative Errors**: Retrieve detailed error information, including function call stack, for better debugging.

## Error check

If you need to check an error, use standard tools like `errors.Is(err, sql.ErrNoRows)`.

## Example

```go
package main

import (
	"errors"
	"fmt"

	"github.com/matchsystems/werr"
)

var errExample = errors.New("find me")

func main() {
	err := example()
	if errors.Is(err, errExample) { // error checking
		fmt.Printf("trace: \n%v\n", err)               // error printing
		fmt.Printf("\nunwrap: %v\n", werr.Unwrap(err)) // error unwrapping
	}
}

func example() error {
	return werr.Wrap(example2()) // possible without text
}

func example2() error {
	return werr.Wrapf(example3(), "without if") // possible without 'if'
}

func example3() error {
	if err := newError(); err != nil {
		return werr.Wrapf(err, "wow error!")
	}

	return nil
}

func newError() error {
	return errExample
}
```

#### Result

```
trace: 
main/main.go:21 example()
main/main.go:25 example2()      without if
main/main.go:30 example3()      wow error!
find me

unwrap: find me
```

### Stack traces benchmark

As performance is obviously an issue, some measurements are in order. The benchmark is provided with the library. In all
of benchmark cases, a very simple code is called that does nothing but grows a number of frames and immediately returns
an error.

Result sample, MacBook Air M1 @ 3.2GHz:

| name                          |     runs | ns/op | note                                         |
|-------------------------------|---------:|------:|----------------------------------------------|
| BenchmarkSimpleError10        | 37410418 | 28.29 | simple error, 10 frames deep                 |
| BenchmarkWrapError10          |  1919391 | 621.7 | same with wrap error                         |
| BenchmarkWrapMsgError10       |  1782106 | 672.8 | same with message                            |
| BenchmarkErrorxError10        |   967269 |  1260 | errorx library, same frames                  |
|                               |          |       |                                              |
| BenchmarkSimpleError100       |  1897574 | 631.7 | simple error, 100 frames deep                |
| BenchmarkWrapError100         |   909345 |  1259 | same with wrap error                         |
| BenchmarkWrapMsgError100      |   867218 |  1310 | same with message                            |
| BenchmarkErrorxError100       |   309862 |  3855 | errorx library, same frames                  |
|                               |          |       |                                              |
| BenchmarkSimpleErrorPrint100  |  1721605 | 697.2 | simple error, 100 frames deep, format output |
| BenchmarkWrapErrorPrint100    |   759574 |  1482 | same with wrap error                         |
| BenchmarkWrapMsgErrorPrint100 |   715376 |  1555 | same with  message                           |
| BenchmarkErrorxErrorPrint100  |    37346 | 32493 | errorx library, same frames                  |

Key takeaways:

* With deep enough call stack, trace capture brings **10x slowdown**
* This is an absolute **worst case measurement, no-op function**; in a real life, much more time is spent doing actual
  work
* Then again, in real life code invocation does not always result in error, so the overhead is proportional to the % of
  error returns
* Still, it pays to omit stack trace collection when it would be of no use
* It is actually **much more expensive to format** an error with a stack trace than to create it, roughly **another 10x
  **
* Compared to the most naive approach to stack trace collection, error creation it is **100x** cheaper with werr
* Therefore, it is totally OK to create an error with a stack trace that would then be handled and not printed to log
* Realistically, stack trace overhead is only painful either if a code is very hot (called a lot and returns errors
  often) or if an error is used as a control flow mechanism and does not constitute an actual problem; in both cases,
  stack trace should be omitted

## More

Portions of the description and benchmark were adapted from the project [errorx](https://github.com/joomcode/errorx)
module github.com/matchsystems/werr/benchmark

go 1.21.4

require (
	github.com/go-errors/errors v1.5.1
	github.com/joomcode/errorx v1.1.1
	github.com/matchsystems/werr v0.0.0-00000000000000-000000000000
)

require github.com/matchsystems/stacktrace v0.0.0-20240211125017-574c181c27b5 // indirect

replace github.com/matchsystems/werr => ../

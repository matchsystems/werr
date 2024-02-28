.PHONY: *

lint:
	golangci-lint run

test:
	CGO_ENABLED=1 go test -race -v ./...
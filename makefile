.PHONY: build test lint clean dev

build:
	go build -o main cmd/app/main.go

test:
	go test -v ./...

lint:
	golangci-lint run

clean:
	go clean
	rm -f main

dev:
	CompileDaemon \
		--build="go build -o main cmd/app/main.go" \
		--command=./main \
		--exclude="*.md,*.txt" \
		--pattern="(.+\.go|.+\.yaml)$$" \
		--graceful-kill \
		--color=true
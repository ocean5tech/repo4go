build:
	@go build -o bin/hello

run: build
	@./bin/hello

test:
	@go test -v ./...

build:
	@go build -o bin/go-repositories-cloner

run: build
	@./bin/go-repositories-cloner

test:
	@go test -v ./...
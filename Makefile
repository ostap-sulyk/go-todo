.PHONY: build run clean

build:
	@go build -o bin/go-todo ./cmd/server/main.go

run: build
	@./bin/go-todo

clean:
	@rm -rf bin
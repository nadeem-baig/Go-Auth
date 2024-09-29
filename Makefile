build:
	@go build -o bin/dist cmd/main.go

run: build
	@./bin/dist

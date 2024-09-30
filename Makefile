build:
	@go build -o bin/dist cmd/main.go

run: build
	@./bin/dist

# Run tests
test:
	@go test ./... -v

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

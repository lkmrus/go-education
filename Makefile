migration-up:
	@echo "Running migration up"
	@go run migrations/auto.go
start:
	@echo "Starting server"
	@go build -o ./cmd/go_build_demo_app_cmd ./cmd && go run cmd/main.go
test:
	@echo "Running tests"
	@go test -v ./...
test-coverage:
	@echo "Running tests with coverage"
	@go test -cover ./...

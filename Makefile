# Makefile
.PHONY: build test run clean docker-build

APP_NAME=honeypotter
BUILD_DIR=build

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) cmd/honeypotter/main.go

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

run:
	go run cmd/honeypotter/main.go

clean:
	rm -rf $(BUILD_DIR)

docker-build:
	docker build -t $(APP_NAME) -f deployments/docker/Dockerfile .

lint:
	golangci-lint run

migrate-up:
	migrate -path internal/infrastructure/database/migrations -database "postgres://user:password@localhost/dbname?sslmode=disable" up

migrate-down:
	migrate -path internal/infrastructure/database/migrations -database "postgres://user:password@localhost/dbname?sslmode=disable" down
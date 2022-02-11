BIN := "./bin/cr"
REPO=github.com/pls87/creative-rotation

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X '${REPO}/cmd/commands.Release=develop' -X '${REPO}/cmd/commands.BuildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S)' -X '${REPO}/cmd/commands.GitHash=$(GIT_HASH)'

lint:
	golangci-lint run ./...

test:
	go test -v -race -count 100 ./...

build: build-img-api build-img-stats build-img-migrations

run: build
	./scripts/run-api-with-tool.sh

run-integration-test: build build-img-integration
	./scripts/run-integration-test.sh

build-local:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd

run-local: build-local
	$(BIN) server --config ./configs/sample.toml & $(BIN) update_stats --config ./configs/sample.toml

run-database-rabbit: build-img-api build-img-migrations
	./scripts/run-database-rabbit.sh

run-api-local: build
	$(BIN) server --config ./configs/sample.toml

run-stats-updater-local: build
	$(BIN) update_stats --config ./configs/sample.toml

build-img-migrations:
	docker build --no-cache -t cr:migrations migrations/.

build-img-api:
	docker build -t cr:api -f build/api/Dockerfile .

build-img-stats:
	docker build -t cr:stats -f build/stats/Dockerfile .

build-img-integration:
	docker build --no-cache -t cr:integration-tests -f build/integration/Dockerfile .

.PHONY: build run build-local run-local test lint run-api run-stats-updater
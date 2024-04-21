BUILD_DIR=build
SQLC_DIR=server/service/repo/db/sqlcimpl/sqlc
MODULE=mobile-telemetry

deps:
	go mod tidy
.PHONY: deps

$(BUILD_DIR)/server: deps
	go build -o $@ $(MODULE)/server

build: $(BUILD_DIR)/server

ctags:
	ctags -R
.PHONY: ctags

lint:
	golangci-lint run
.PHONY: lint

sqlc:
	cd $(SQLC_DIR) && sqlc generate
.PHONY: sqlc

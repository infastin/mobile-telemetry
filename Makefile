BUILD_DIR=build
MSGP_DIR=server/service/repo/db/badgerimpl/queries
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

msgp:
	go generate ./$(MSGP_DIR)
.PHONY: msgp

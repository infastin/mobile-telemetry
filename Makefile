BUILD_DIR=build
ENT_DIR=server/service/repo/db/entimpl/ent
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

entdesc:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./$(ENT_DIR)/schema
.PHONY: entdesc

entgen:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./$(ENT_DIR)/schema
.PHONY: entgen

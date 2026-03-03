BIN_DIR := $(CURDIR)/bin
GO := go

# Versions
AIR_VERSION := v1.64.5
GOLANGCI_LINT_VERSION := v2.10.1
PROTOC_VERSION := 34.0
PROTOC_GEN_GO_VERSION := latest
PROTOC_GEN_GO_GRPC_VERSION := latest

# OS detection
ifeq ($(OS),Windows_NT)
	EXT := .exe
	PROTOC_FILE := protoc-$(PROTOC_VERSION)-win64.zip
	PROTOC_URL := https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_FILE)
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		EXT :=
		PROTOC_FILE := protoc-$(PROTOC_VERSION)-linux-x86_64.zip
		PROTOC_URL := https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_FILE)
	endif
	ifeq ($(UNAME_S),Darwin)
		EXT :=
		PROTOC_FILE := protoc-$(PROTOC_VERSION)-osx-x86_64.zip
		PROTOC_URL := https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_FILE)
	endif
endif

# Binaries
AIR := $(BIN_DIR)/air$(EXT)
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint$(EXT)
PROTOC := $(BIN_DIR)/protoc$(EXT)
PROTOC_GEN_GO := $(BIN_DIR)/protoc-gen-go$(EXT)
PROTOC_GEN_GO_GRPC := $(BIN_DIR)/protoc-gen-go-grpc$(EXT)

TOOLS := $(AIR) $(GOLANGCI_LINT) $(PROTOC) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC)

# Commands
.PHONY: all
all: build

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(AIR): | $(BIN_DIR)
	GOBIN=$(BIN_DIR) $(GO) install github.com/air-verse/air@$(AIR_VERSION)

$(GOLANGCI_LINT): | $(BIN_DIR)
	GOBIN=$(BIN_DIR) $(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

$(PROTOC_GEN_GO): | $(BIN_DIR)
	GOBIN=$(BIN_DIR) $(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)

$(PROTOC_GEN_GO_GRPC): | $(BIN_DIR)
	GOBIN=$(BIN_DIR) $(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

$(PROTOC): | $(BIN_DIR)
	@echo "Downloading protoc $(PROTOC_VERSION)..."
	@curl -L -o $(BIN_DIR)/$(PROTOC_FILE) $(PROTOC_URL)
	@unzip -o -j $(BIN_DIR)/$(PROTOC_FILE) "bin/protoc$(EXT)" -d $(BIN_DIR)
	@rm $(BIN_DIR)/$(PROTOC_FILE)

.PHONY: tools
tools: $(TOOLS)

.PHONY: check-tools
check-tools:
	@for tool in $(TOOLS); do \
		if [ ! -f $$tool ]; then \
			echo "Missing $$tool. Installing tools..."; \
			$(MAKE) tools; \
			break; \
		fi \
	done

.PHONY: dev
dev: check-tools
	$(AIR)

.PHONY: lint
lint: check-tools
	$(GOLANGCI_LINT) run ./...

.PHONY: build
build:
	mkdir -p bin
	$(GO) build -o bin/app ./cmd/app

.PHONY: proto-gen
proto-gen: check-tools
	@echo "Generating protobuf code..."
	@$(PROTOC) \
		--proto_path=protobuf protobuf/orders.proto \
		--go_out=apps/common/genproto/orders \
		--go_opt=paths=source_relative \
		--go-grpc_out=apps/common/genproto/orders \
		--go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go="$(PROTOC_GEN_GO)" \
		--plugin=protoc-gen-go-grpc="$(PROTOC_GEN_GO_GRPC)"

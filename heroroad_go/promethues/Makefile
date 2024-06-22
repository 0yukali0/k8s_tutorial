GO := go
GOFMT := gofmt
GOCI_LINT := golangci-lint
BUILD_DIR := build
PING_EXE := ping.exe

.PHONY: build
build: $(BUILD_DIR)
$(BUILD_DIR): go.mod go.sum $(shell find pkg)
	@echo "building ping project"
	@mkdir "$(BUILD_DIR)"
	"$(GO)" build \
	-o "$(BUILD_DIR)/$(PING_EXE)" \
	./pkg/cmd/

.PHONY: lint
lint:
	"$(GOCI_LINT)" run ./...

.PHONY: fmt
fmt:
	"$(GOFMT)" -w -l $(shell find pkg)

.PHONY: clean
clean: 
	@echo "cleaning up caches and output"
	@rm -rf "${BUILD_DIR}"
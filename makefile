# Go parameters
GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")
BUILD_DIR := build

# Target operating systems and architectures
TARGETS := linux/amd64 windows/amd64 darwin/amd64

# Generate build targets
build_targets = $(addsuffix /main,$(TARGETS))
build_targets := $(addprefix $(BUILD_DIR)/,$(build_targets))

# Docker image name
DOCKER_IMAGE := graphservice

# Default target to install dependencies, format code, and run all tests
.PHONY: all
all: install-deps format test build

# Install dependencies
.PHONY: install-deps
install-deps:
	@echo "Installing dependencies..."
	go mod tidy

# Format code
.PHONY: format
format:
	@echo "Formatting code..."
	go fmt ./...

# Build executables for multiple operating systems and architectures
.PHONY: build
build: $(build_targets)

$(BUILD_DIR)/%/main: $(GO_FILES)
	@echo "Building $@..."
	@mkdir -p $(dir $@)
	GOOS=$(word 1,$(subst /, ,$*)) GOARCH=$(word 2,$(subst /, ,$*)) CGO_ENABLED=0 go build -o $@ ./cmd

# Run test cases
.PHONY: test
test:
	@echo "Running test cases..."
	go test ./...

# Run performance test cases
.PHONY: perf-test
perf-test:
	@echo "Running performance test cases..."
	go test -bench=.

# Build Docker image
.PHONY: docker
docker: $(BUILD_DIR)/linux/amd64/main
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Clean build directory
.PHONY: clean
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)

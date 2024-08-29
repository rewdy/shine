BINARY_NAME=shine
DIST_DIR=dist
SRC_DIR=./cmd/shine

# Build the binary (default)
build:
	@echo "\033[34mMaking single local build: \033[33m./$(DIST_DIR)/$(BINARY_NAME)\033[34m..\033[0m"
	go build -o $(DIST_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Build for all targets
build-releases: build-macos-amd build-macos-arm build-linux build-windows

build-macos-amd:
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-macos-amd64 $(SRC_DIR)

build-macos-arm:
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME)-macos-arm64 $(SRC_DIR)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64 $(SRC_DIR)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe $(SRC_DIR)

dev:
	air -c .air.toml

# Run tests
test:
	go test ./...

# Format the code
fmt:
	go fmt ./...

# Clean up generated files
clean:
	rm -rf $(DIST_DIR) tmp

# Install the binary
install:
	go install $(SRC_DIR)

# Run the application
run:
	go run $(SRC_DIR)/main.go

.PHONY: build build-releases build-macos-amd build-macos-arm build-linux build-windows dev test fmt clean install run
BINARY_NAME=xorfiles-go-one
BUILD_DIR=./bin

all: run

build:
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) -tags netgo .

clean:
	@rm -f $(BUILD_DIR)/$(BINARY_NAME)

run: build
	@cd $(BUILD_DIR) && ./$(BINARY_NAME)

.PHONY: all build clean run

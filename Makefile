BINARY_NAME=example.exe
BUILD_DIR=builder
SOURCES=$(wildcard *.go)
GO=go
MSYS_PATH=D:/msys64
GCC_PATH=D:/msys64/mingw64/bin/gcc.exe
GO_BIN_PATH=D:/golang/bin/go.exe
PATH := $(GO_BIN_PATH):$(PATH)

export MSYS2_PATH_TYPE=inherit
export MSYSTEM=MINGW64
export CHERE_INVOKING=1


build:
	@if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)
	$(GO) build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCES)
	$(GCC_PATH) -v

	$(MSYS_PATH)/usr/bin/bash.exe -l -c 'GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)'

run:
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	rmdir /S /Q $(BUILD_DIR)

.PHONY: build clean run

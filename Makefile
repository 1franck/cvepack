MAIN_CMD = ./cmd/cvepack.go
MAIN_BINARY_NAME = cvepack
MAIN_BINARY_BIN = ./bin/$(MAIN_BINARY_NAME)

DB_COMPILER_CMD = ./cmd/compile-advdb.go
DB_COMPILER_BINARY_NAME = compile-advdb
DB_COMPILER_BINARY_BIN = ./bin/$(DB_COMPILER_BINARY_NAME)

GO_ENV_DARWIN_64 = GOARCH=amd64 GOOS=darwin
GO_ENV_LINUX_64 = GOARCH=amd64 GOOS=linux
GO_ENV_WINDOWS_64 = GOARCH=amd64 GOOS=windows

GO_PROD_FLAGS = -ldflags "-s -w"

VERSION = $(shell grep -oE 'VERSION = "[^"]+"' internal/const.go | cut -d'"' -f2)

RUN_OS_BIN =
ifeq ($(OS),Windows_NT)
	RUN_OS_BIN := windows.exe
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		RUN_OS_BIN := linux
	endif
	ifeq ($(UNAME_S),Darwin)
		RUN_OS_BIN := darwin
	endif
endif

.DEFAULT_GOAL := run

build:
	$(GO_ENV_DARWIN_64) go build -o $(MAIN_BINARY_BIN)-darwin $(MAIN_CMD)
	$(GO_ENV_LINUX_64) go build -o $(MAIN_BINARY_BIN)-linux $(MAIN_CMD)
	$(GO_ENV_WINDOWS_64) go build -o $(MAIN_BINARY_BIN)-windows.exe $(MAIN_CMD)
	$(GO_ENV_DARWIN_64) go build -o $(DB_COMPILER_BINARY_BIN)-darwin $(DB_COMPILER_CMD)
	$(GO_ENV_LINUX_64) go build -o $(DB_COMPILER_BINARY_BIN)-linux $(DB_COMPILER_CMD)
	$(GO_ENV_WINDOWS_64) go build -o $(DB_COMPILER_BINARY_BIN)-windows.exe $(DB_COMPILER_CMD)

build-prod:
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o $(MAIN_BINARY_BIN)-darwin $(MAIN_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o $(MAIN_BINARY_BIN)-linux $(MAIN_CMD)
	$(GO_ENV_WINDOWS_64) go build $(GO_PROD_FLAGS) -o $(MAIN_BINARY_BIN)-windows.exe $(MAIN_CMD)
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BINARY_BIN)-darwin $(DB_COMPILER_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BINARY_BIN)-linux $(DB_COMPILER_CMD)
	$(GO_ENV_WINDOWS_64) go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BINARY_BIN)-windows.exe $(DB_COMPILER_CMD)

run: build
	$(MAIN_BINARY_BIN)-$(RUN_OS_BIN)

clean:
	go clean
	rm $(MAIN_BINARY_BIN)-darwin
	rm $(MAIN_BINARY_BIN)-linux
	rm $(MAIN_BINARY_BIN)-windows.exe
	rm $(DB_COMPILER_BINARY_BIN)-darwin
	rm $(DB_COMPILER_BINARY_BIN)-linux
	rm $(DB_COMPILER_BINARY_BIN)-windows.exe

version:
	@echo $(VERSION)
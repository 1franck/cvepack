#cvepack bin
MAIN_CMD = ./cmd/cvepack.go
MAIN_BIN = cvepack
MAIN_BIN_PATH = ./bin/$(MAIN_BIN)

#compile-advdb bin
DB_COMPILER_CMD = ./cmd/compile-advdb.go
DB_COMPILER_BIN = compile-advdb
DB_COMPILER_BIN_PATH = ./bin/$(DB_COMPILER_BIN)

#upload-compiled-advdb bin
UPLOAD_DB_CMD = ./cmd/upload-compiled-advdb.go
UPLOAD_DB_BIN = upload-compiled-advdb
UPLOAD_DB_BIN_PATH = ./bin/$(UPLOAD_DB_BIN)

GO_ENV_DARWIN_64 = GOARCH=amd64 GOOS=darwin
GO_ENV_LINUX_64 = GOARCH=amd64 GOOS=linux
GO_ENV_WIN_64 = GOARCH=amd64 GOOS=windows

GO_PROD_FLAGS = -ldflags "-s -w"

VERSION = $(shell grep -oE 'VERSION = "[^"]+"' internal/const.go | cut -d'"' -f2)

RUN_OS_BIN =
ifeq ($(OS),Windows_NT)
	RUN_OS_BIN := win.exe
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
.PHONY: clean version tests

build:
	$(GO_ENV_DARWIN_64) go build -o $(MAIN_BIN_PATH)-darwin $(MAIN_CMD)
	$(GO_ENV_LINUX_64) go build -o $(MAIN_BIN_PATH)-linux $(MAIN_CMD)
	$(GO_ENV_WIN_64) go build -o $(MAIN_BIN_PATH)-win.exe $(MAIN_CMD)
	$(GO_ENV_DARWIN_64) go build -o $(DB_COMPILER_BIN_PATH)-darwin $(DB_COMPILER_CMD)
	$(GO_ENV_LINUX_64) go build -o $(DB_COMPILER_BIN_PATH)-linux $(DB_COMPILER_CMD)
	$(GO_ENV_WIN_64) go build -o $(DB_COMPILER_BIN_PATH)-win.exe $(DB_COMPILER_CMD)
	$(GO_ENV_DARWIN_64) go build -o $(UPLOAD_DB_BIN_PATH)-darwin $(UPLOAD_DB_CMD)
	$(GO_ENV_LINUX_64) go build -o $(UPLOAD_DB_BIN_PATH)-linux $(UPLOAD_DB_CMD)
	$(GO_ENV_WIN_64) go build -o $(UPLOAD_DB_BIN_PATH)-win.exe $(UPLOAD_DB_CMD)

build-prod:
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o $(MAIN_BIN_PATH)-darwin $(MAIN_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o $(MAIN_BIN_PATH)-linux $(MAIN_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o $(MAIN_BIN_PATH)-win.exe $(MAIN_CMD)
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BIN_PATH)-darwin $(DB_COMPILER_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BIN_PATH)-linux $(DB_COMPILER_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BIN_PATH)-win.exe $(DB_COMPILER_CMD)
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o $(UPLOAD_DB_BIN_PATH)-darwin $(UPLOAD_DB_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o $(UPLOAD_DB_BIN_PATH)-linux $(UPLOAD_DB_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o $(UPLOAD_DB_BIN_PATH)-win.exe $(UPLOAD_DB_CMD)

build-arm64:
	GOOS=linux GOARM=7 GOARCH=arm64 go build $(GO_PROD_FLAGS) -o $(MAIN_BIN_PATH)-arm $(MAIN_CMD)
	GOOS=linux GOARM=7 GOARCH=arm64 go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BIN_PATH)-arm $(DB_COMPILER_CMD)
	GOOS=linux GOARM=7 GOARCH=arm64 go build $(GO_PROD_FLAGS) -o $(UPLOAD_DB_BIN_PATH)-arm $(UPLOAD_DB_CMD)

run: build
	$(MAIN_BIN_PATH)-$(RUN_OS_BIN)

clean:
	go clean
	rm $(MAIN_BIN_PATH)-arm
	rm $(MAIN_BIN_PATH)-darwin
	rm $(MAIN_BIN_PATH)-linux
	rm $(MAIN_BIN_PATH)-win.exe
	rm $(DB_COMPILER_BIN_PATH)-arm
	rm $(DB_COMPILER_BIN_PATH)-darwin
	rm $(DB_COMPILER_BIN_PATH)-linux
	rm $(DB_COMPILER_BIN_PATH)-win.exe
	rm $(UPLOAD_DB_BIN_PATH)-arm
	rm $(UPLOAD_DB_BIN_PATH)-darwin
	rm $(UPLOAD_DB_BIN_PATH)-linux
	rm $(UPLOAD_DB_BIN_PATH)-win.exe

version:
	@echo $(VERSION)

tests:
	go test -v ./tests/...
#cvepack bin
MAIN_CMD = ./cmd/cli/cvepack/cvepack.go
MAIN_BIN = cvepack
MAIN_BIN_PATH = ./bin/cli/$(MAIN_BIN)

#compile-advdb bin
DB_COMPILER_CMD = ./cmd/cli/advisory-database/compiler.go
DB_COMPILER_BIN = compiler
DB_COMPILER_BIN_PATH = ./bin/cli/advisory-database/$(DB_COMPILER_BIN)

#upload-compiled-advdb bin
DB_UPLOADER_CMD = ./cmd/cli/advisory-database/uploader.go
DB_UPLOADER_BIN = uploader
DB_UPLOADER_BIN_PATH = ./bin/cli/advisory-database/$(DB_UPLOADER_BIN)

GO_ENV_DARWIN_64 = GOARCH=amd64 GOOS=darwin
GO_ENV_LINUX_64 = GOARCH=amd64 GOOS=linux
GO_ENV_WIN_64 = GOARCH=amd64 GOOS=windows
GO_ENV_ARM_64 = GOARCH=arm64 GOOS=linux GOARM=7

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
	$(GO_ENV_DARWIN_64) go build -o ./bin/cli/darwin/$(MAIN_BIN_PATH)-darwin $(MAIN_CMD)
	$(GO_ENV_LINUX_64) go build -o ./bin/cli/darwin/$(MAIN_BIN_PATH)-linux $(MAIN_CMD)
	$(GO_ENV_WIN_64) go build -o ./bin/cli/darwin/$(MAIN_BIN_PATH)-win.exe $(MAIN_CMD)
	$(GO_ENV_DARWIN_64) go build -o $(DB_COMPILER_BIN_PATH)-darwin $(DB_COMPILER_CMD)
	$(GO_ENV_LINUX_64) go build -o $(DB_COMPILER_BIN_PATH)-linux $(DB_COMPILER_CMD)
	$(GO_ENV_WIN_64) go build -o $(DB_COMPILER_BIN_PATH)-win.exe $(DB_COMPILER_CMD)
	$(GO_ENV_DARWIN_64) go build -o $(DB_UPLOADER_BIN_PATH)-darwin $(DB_UPLOADER_CMD)
	$(GO_ENV_LINUX_64) go build -o $(DB_UPLOADER_BIN_PATH)-linux $(DB_UPLOADER_CMD)
	$(GO_ENV_WIN_64) go build -o $(DB_UPLOADER_BIN_PATH)-win.exe $(DB_UPLOADER_CMD)

build-prod:
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o $(MAIN_BIN_PATH)-darwin $(MAIN_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o $(MAIN_BIN_PATH)-linux $(MAIN_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o $(MAIN_BIN_PATH)-win.exe $(MAIN_CMD)
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BIN_PATH)-darwin $(DB_COMPILER_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BIN_PATH)-linux $(DB_COMPILER_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BIN_PATH)-win.exe $(DB_COMPILER_CMD)
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o $(DB_UPLOADER_BIN_PATH)-darwin $(DB_UPLOADER_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o $(DB_UPLOADER_BIN_PATH)-linux $(DB_UPLOADER_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o $(DB_UPLOADER_BIN_PATH)-win.exe $(DB_UPLOADER_CMD)

build-gh:
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o ./bin/darwin/$(MAIN_BIN) $(MAIN_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o ./bin/linux/$(MAIN_BIN) $(MAIN_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o ./bin/win/$(MAIN_BIN).exe $(MAIN_CMD)
	$(GO_ENV_ARM_64) go build $(GO_PROD_FLAGS) -o ./bin/arm/$(MAIN_BIN) $(MAIN_CMD)

build-arm64:
	GOOS=linux GOARM=7 GOARCH=arm64 go build $(GO_PROD_FLAGS) -o $(MAIN_BIN_PATH)-arm $(MAIN_CMD)
	GOOS=linux GOARM=7 GOARCH=arm64 go build $(GO_PROD_FLAGS) -o $(DB_COMPILER_BIN_PATH)-arm $(DB_COMPILER_CMD)
	GOOS=linux GOARM=7 GOARCH=arm64 go build $(GO_PROD_FLAGS) -o $(DB_UPLOADER_BIN_PATH)-arm $(DB_UPLOADER_CMD)

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
	rm $(DB_UPLOADER_BIN_PATH)-arm
	rm $(DB_UPLOADER_BIN_PATH)-darwin
	rm $(DB_UPLOADER_BIN_PATH)-linux
	rm $(DB_UPLOADER_BIN_PATH)-win.exe
	rm ./bin/darwin/$(MAIN_BIN)
	rm ./bin/linux/$(MAIN_BIN)
	rm ./bin/win/$(MAIN_BIN).exe
	rm ./bin/arm/$(MAIN_BIN)

version:
	@echo $(VERSION)

tests:
	go test -v ./tests/...
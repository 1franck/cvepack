#cvepack bin
MAIN_CMD = ./cmd/cli/cvepack/cvepack.go
MAIN_BIN = cvepack

#compile-advdb bin
DB_COMPILER_CMD = ./cmd/cli/advisory-database/compiler.go
DB_COMPILER_BIN = compiler

#upload-compiled-advdb bin
DB_UPLOADER_CMD = ./cmd/cli/advisory-database/uploader.go
DB_UPLOADER_BIN = uploader

# GO ENV & FLAGS
GO_ENV_DARWIN_64 = GOARCH=amd64 GOOS=darwin
GO_ENV_LINUX_64 = GOARCH=amd64 GOOS=linux
GO_ENV_WIN_64 = GOARCH=amd64 GOOS=windows
GO_ENV_ARM_64 = GOARCH=arm64 GOOS=linux GOARM=7
GO_PROD_FLAGS = -ldflags "-s -w"

# cvepack version
VERSION = $(shell grep -oE 'VERSION = "[^"]+"' internal/const.go | cut -d'"' -f2)

RUN_OS_BIN =
ifeq ($(OS),Windows_NT)
	RUN_OS_BIN := ./bin/win
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		RUN_OS_BIN := ./bin/darwin
	endif
	ifeq ($(UNAME_S),Linux)
		ifeq ($(shell uname -m),armv7l)
			RUN_OS_BIN := ./bin/arm
		endif
	endif
	ifeq ($(UNAME_S),Linux)
		RUN_OS_BIN := ./bin/linux
	endif
endif

.DEFAULT_GOAL := run
.PHONY: clean version tests

build:
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o ./bin/darwin/cli/$(MAIN_BIN) $(MAIN_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o ./bin/linux/cli/$(MAIN_BIN) $(MAIN_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o ./bin/win/cli/$(MAIN_BIN).exe $(MAIN_CMD)
	$(GO_ENV_ARM_64) go build $(GO_PROD_FLAGS) -o ./bin/arm/cli/$(MAIN_BIN) $(MAIN_CMD)

build-tools:
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o ./bin/darwin/cli/tools/$(DB_COMPILER_BIN) $(DB_COMPILER_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o ./bin/linux/cli/tools/$(DB_COMPILER_BIN) $(DB_COMPILER_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o ./bin/win/cli/tools/$(DB_COMPILER_BIN).exe $(DB_COMPILER_CMD)
	$(GO_ENV_ARM_64) go build $(GO_PROD_FLAGS) -o ./bin/arm/cli/tools/$(DB_COMPILER_BIN) $(DB_COMPILER_CMD)
	$(GO_ENV_DARWIN_64) go build $(GO_PROD_FLAGS) -o ./bin/darwin/cli/tools/$(DB_UPLOADER_BIN) $(DB_UPLOADER_CMD)
	$(GO_ENV_LINUX_64) go build $(GO_PROD_FLAGS) -o ./bin/linux/cli/tools/$(DB_UPLOADER_BIN) $(DB_UPLOADER_CMD)
	$(GO_ENV_WIN_64) go build $(GO_PROD_FLAGS) -o ./bin/win/cli/tools/$(DB_UPLOADER_BIN).exe $(DB_UPLOADER_CMD)
	$(GO_ENV_ARM_64) go build $(GO_PROD_FLAGS) -o ./bin/arm/cli/tools/$(DB_UPLOADER_BIN) $(DB_UPLOADER_CMD)

run: build
	$(RUN_OS_BIN)/cli/$(MAIN_BIN)

clean:
	go clean
	rm ./bin/darwin/cli/$(MAIN_BIN)
	rm ./bin/linux/cli/$(MAIN_BIN)
	rm ./bin/win/cli/$(MAIN_BIN).exe
	rm ./bin/arm/cli/$(MAIN_BIN)
	rm ./bin/darwin/cli/tools/$(DB_COMPILER_BIN)
	rm ./bin/linux/cli/tools/$(DB_COMPILER_BIN)
	rm ./bin/win/cli/tools/$(DB_COMPILER_BIN).exe
	rm ./bin/arm/cli/tools/$(DB_COMPILER_BIN)
	rm ./bin/darwin/cli/tools/$(DB_UPLOADER_BIN)
	rm ./bin/linux/cli/tools/$(DB_UPLOADER_BIN)
	rm ./bin/win/cli/tools/$(DB_UPLOADER_BIN).exe
	rm ./bin/arm/cli/tools/$(DB_UPLOADER_BIN)

version:
	@echo $(VERSION)

tests:
	go test -v ./tests/...
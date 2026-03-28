.PHONY: build clean tidy lint install help

BUILD_DIR ?= .build
PKG_NAME=archimate-mcp

build:
	@[ -d ${BUILD_DIR} ] || mkdir -p ${BUILD_DIR}
	CGO_ENABLED=0 go build -o ${BUILD_DIR}/${PKG_NAME} .
	@file  ${BUILD_DIR}/${PKG_NAME}
	@du -h ${BUILD_DIR}/${PKG_NAME}

clean:
	rm -f ${BUILD_DIR}/${PKG_NAME}

tidy:
	go mod tidy

lint: ## Run linter
	@golangci-lint run --timeout 5m

install:
	@go install ./

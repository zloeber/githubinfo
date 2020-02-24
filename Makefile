SHELL := /bin/bash
.DEFAULT_GOAL := help
.PHONY: build build-alpine clean test help default

SRCS := $(shell find . -name '*.go')
LINTERS := \
	golang.org/x/lint/golint \
	github.com/kisielk/errcheck \
	honnef.co/go/tools/cmd/staticcheck

BIN_NAME := githubinfo
VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT := $(shell git rev-parse HEAD)
GIT_DIRTY := $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE := $(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "zloeber/githubinfo"
LDFLAGS := -X github.com/zloeber/githubinfo/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/zloeber/githubinfo/version.BuildDate=${BUILD_DATE}

.PHONY: help
help: ## Help
	@grep --no-filename -E '^[a-zA-Z_/-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: deps ## Compile the project.
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -v \
		-ldflags "$(LDFLAGS)" \
		-o bin/${BIN_NAME}

.PHONY: build-alpine
build-alpine: ## Compile optimized for alpine linux.
	go build \
		-ldflags '-w -linkmode external -extldflags "-static" $(LDFLAGS)' \
		-o bin/${BIN_NAME}

.PHONY: image
image: ## Build docker image
	docker build \
		--build-arg VERSION=${VERSION} \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		-t $(IMAGE_NAME):local .

.PHONY: tag
tag: ## Tag docker image
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

.PHONY: push
push: tag  ## Push tagged images to registry
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest

.PHONY: deps
deps: ## Install dependencies
	go get -d -v ./...

.PHONY: updatedeps
updatedeps: ## Update dependencies
	go get -d -v -u -f ./...

.PHONY: testdeps
testdeps: ## Install test deps
	go get -d -v -t ./...
	go get -v $(LINTERS)

.PHONY: updatetestdeps
updatetestdeps: ## Update test deps
	go get -d -v -t -u -f ./...
	go get -u -v $(LINTERS)

.PHONY: install
install: deps ## Install
	go install ./...

.PHONY: golint
golint: testdeps ## Code linting
	for file in $(SRCS); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

.PHONY: vet
vet: testdeps ## Code vetting
	go vet ./...

.PHONY: testdeps
errcheck: testdeps ## Error checking
	errcheck ./...

.PHONY: staticcheck
staticcheck: testdeps ## Static testing
	staticcheck ./...

.PHONY: lint
lint: golint vet errcheck staticcheck ## Lint the project

.PHONY: version
version: ## Go version
	@go version

.PHONY: test
test: testdeps lint ## Run tests
	go test -race ./..

.PHONY: clean
clean: ## Clean the directory tree.
	go clean -i ./...
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}
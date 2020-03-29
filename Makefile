SHELL := /bin/bash
.DEFAULT_GOAL := help

## Modify these for your project
VENDOR ?= zloeber
APP ?= githubinfo
IMAGE_NAME ?= $(VENDOR)/$(APP)
REPO ?= github.com/$(VENDOR)/$(APP)

SRCS := $(shell find . -name '*.go')
GO_VERSION := $(shell cat ./.tool-versions | grep golang | cut -f 2 -d " ")
#VERSION := $(git describe --tags `git rev-list --tags --max-count=1`)
VERSION := $(shell grep "const Version " pkg/version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT := $(shell git rev-parse HEAD)
RELEASE_VERSION ?= $(VERSION)-$(GIT_COMMIT)
GIT_DIRTY := $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE := $(shell date '+%Y-%m-%d-%H:%M:%S')
LDFLAGS := -X $(REPO)/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X $(REPO)/version.BuildDate=${BUILD_DATE}
LINTERS := \
	golang.org/x/lint/golint \
	github.com/kisielk/errcheck \
	honnef.co/go/tools/cmd/staticcheck

HELPERAPPS := \
	github.com/spf13/cobra/cobra

.PHONY: help
help: ## Help
	@grep --no-filename -E '^[a-zA-Z_%/-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: deps mod/tidy ## Compile the project.
	go build -v -ldflags "$(LDFLAGS)" -o bin ./cmd/...

.PHONY: docker/image
docker/image: mod/tidy ## Build docker image
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg LDFLAGS="$(LDFLAGS)" \
		-t $(IMAGE_NAME):local .

.PHONY: docker/tag
docker/tag: ## Tag docker image
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

.PHONY: docker/push
docker/push: docker/tag  ## Push tagged images to registry
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest

.PHONY: docker/run
docker/run: ## Run a local docker image for the app
	docker run -i -t --rm --name=$(APP) $(IMAGE_NAME):local

.PHONY: deps
deps: ## Install dependencies
	go get -d -v ./...
	go get -v -u $(HELPERAPPS)


.PHONY: deps/update
deps/update: ## Update dependencies
	go get -d -v -u -f ./...

.PHONY: mod/tidy
mod/tidy: ## Update module dependencies
	go mod tidy

.PHONY: test/deps
test/deps: ## Install test deps
	go get -d -v -t ./...
	go get -v -u $(LINTERS)

.PHONY: test/deps/update
test/deps/update: ## Update test deps
	go get -d -v -t -u -f ./...
	go get -u -v $(LINTERS)

.PHONY: install
install: deps ## Install
	go install ./...

.PHONY: golint
golint: test/deps ## Code linting
	for file in $(SRCS); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

.PHONY: vet
vet: test/deps ## Code vetting
	go vet ./...

.PHONY: test/deps
errcheck: test/deps ## Error checking
	errcheck ./...

.PHONY: staticcheck
staticcheck: test/deps ## Static testing
	staticcheck ./...

.PHONY: lint
lint: golint vet errcheck staticcheck ## Lint, vet, errcheck, and staticcheck

.PHONY: version
version: ## Go version
	@go version

.PHONY: test
test: test/deps lint ## Run tests
	go test -race ./..

.PHONY: clean
clean: ## Clean the directory tree.
	go clean -i ./...
	@test ! -e bin/${APP} || rm bin/${APP}

.PHONY: show
show: ## Show various build settings
	@echo "VENDOR: $(VENDOR)"
	@echo "APP: $(APP)"
	@echo "IMAGE_NAME: $(IMAGE_NAME)"
	@echo "REPO: $(REPO)"
	@echo "GO_VERSION: $(GO_VERSION)"
	@echo "VERSION: $(VERSION)"
	@echo "GIT_COMMIT: $(GIT_COMMIT)"
	@echo "GIT_DIRTY: $(GIT_DIRTY)"
	@echo "BUILD_DATE: $(BUILD_DATE)"
	@echo "RELEASE_VERSION: $(RELEASE_VERSION)"
	@echo "LDFLAGS: $(LDFLAGS)"

.PHONY: release
release:
	git add -all .
	git commit -m "release: commit before release"
	git tag -a v$(RELEASE_VERSION) -m "auto-release"
	git push origin master --tags

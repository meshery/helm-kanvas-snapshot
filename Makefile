# Copyright Meshery Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# include build/Makefile.core.mk


# .PHONY: all
# all: dep-check build
# ## Lint check
# golangci: error dep-check
# 	golangci-lint run --exclude-use-default

# ## Analyze error codes
# error: dep-check
# 	go run github.com/layer5io/meshkit/cmd/errorutil -d . analyze -i ./build -o ./build

# ## Runs meshkit error utility to update error codes.
# error-util:
# 	go run github.com/layer5io/meshkit/cmd/errorutil -d . update -i ./build -o ./build



# #-----------------------------------------------------------------------------
# # Dependencies
# #-----------------------------------------------------------------------------
# .PHONY: dep-check
# #.SILENT: dep-check

# INSTALLED_GO_VERSION=$(shell go version)

# dep-check:

# ifeq (,$(findstring $(GOVERSION), $(INSTALLED_GO_VERSION)))
# # Only send a warning.
# 	@echo "Dependency missing: go$(GOVERSION). Ensure 'go$(GOVERSION).x' is installed and available in your 'PATH'"
# 	@echo "GOVERSION: " $(GOVERSION)
# 	@echo "INSTALLED_GO_VERSION: " $(INSTALLED_GO_VERSION)
# endif


# OUTDIR := ./cmd/kanvas-snapshot/bin
# ARCH := amd64

# BINNAME_DARWIN ?= kanvas-snapshot-darwin-$(ARCH)
# BINNAME_LINUX ?= kanvas-snapshot-linux-$(ARCH)
# BINNAME_WINDOWS ?= kanvas-snapshot-windows-$(ARCH).exe


# LDFLAGS := "\
#     -X 'main.providerToken=$(PROVIDER_TOKEN)' \
#     -X 'main.MesheryCloudApiBaseUrl=$(MESHERY_CLOUD_API_BASE_URL)' \
#     -X 'main.MesheryApiBaseUrl=$(MESHERY_API_BASE_URL)'"
# .PHONY: build
# build:
# 	@echo "Building for all platforms..."
# 	@$(MAKE) $(BINNAME_DARWIN)
# 	@$(MAKE) $(BINNAME_LINUX)
# 	@$(MAKE) $(BINNAME_WINDOWS)

# # Build Helm plugin for Darwin (macOS)
# .PHONY: $(BINNAME_DARWIN)
# $(BINNAME_DARWIN):
# 	@echo "Building for Darwin..."
# 	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=darwin go build -ldflags=$(LDFLAGS) -o $(OUTDIR)/$(BINNAME_DARWIN) ./cmd/kanvas-snapshot/main.go

# # Build Helm plugin for Linux
# .PHONY: $(BINNAME_LINUX)
# $(BINNAME_LINUX):
# 	@echo "Building for Linux..."
# 	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=linux go build -ldflags=$(LDFLAGS) -o $(OUTDIR)/$(BINNAME_LINUX) ./cmd/kanvas-snapshot/main.go

# # Build Helm plugin for Windows
# .PHONY: $(BINNAME_WINDOWS)
# $(BINNAME_WINDOWS):
# 	@echo "Building for Windows..."
# 	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=windows go build -ldflags=$(LDFLAGS) -o $(OUTDIR)/$(BINNAME_WINDOWS) ./cmd/kanvas-snapshot/main.go

# # Clean up binaries
# .PHONY: clean
# clean:
# 	@echo "Cleaning up..."
# 	rm -rf $(OUTDIR)




GOFMT_FILES?=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
APP_NAME?=helm-kanvas-snapshot
APP_DIR?=$(shell git rev-parse --show-toplevel)
DEV?=${DEVBOX_TRUE}
SRC_PACKAGES=$(shell go list ./... | grep -v "mocks")
BUILD_ENVIRONMENT?=${ENVIRONMENT}
VERSION?=0.1.0
REVISION?=$(shell git rev-parse --verify HEAD)
DATE?=$(shell date)
PLATFORM?=$(shell go env GOOS)
ARCHITECTURE?=$(shell go env GOARCH)
GOVERSION?=$(shell go version | awk '{printf $$3}')
BUILD_WITH_FLAGS="-s -w -X 'github.com/meshery/helm-kanvas-snapshot/version.Version=${VERSION}' -X 'github.com/meshery/helm-kanvas-snapshot/version.Env=${BUILD_ENVIRONMENT}' -X 'github.com/meshery/helm-kanvas-snapshot/version.BuildDate=${DATE}' -X 'github.com/meshery/helm-kanvas-snapshot/version.Revision=${REVISION}' -X 'github.com/meshery/helm-kanvas-snapshot/version.Platform=${PLATFORM}/${ARCHITECTURE}' -X 'github.com/meshery/helm-kanvas-snapshot/version.GoVersion=${GOVERSION}'  -X 'main.providerToken=$(PROVIDER_TOKEN)' -X 'main.MesheryCloudApiBaseUrl=$(MESHERY_CLOUD_API_BASE_URL)'  -X 'main.MesheryApiBaseUrl=$(MESHERY_API_BASE_URL)'"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.PHONY: help
help: ## Prints help (only for targets with comments)
	@grep -E '^[a-zA-Z0-9._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

local.fmt: ## Lints all the Go code in the application.
	@gofmt -w $(GOFMT_FILES)
	$(GOBIN)/goimports -w $(GOFMT_FILES)
	$(GOBIN)/gofumpt -l -w $(GOFMT_FILES)
	$(GOBIN)/gci write $(GOFMT_FILES) --skip-generated

local.check: local.fmt ## Loads all dependencies
	@go mod tidy

local.build: local.check ## Generates the artifact with 'go build'
	@go build -o $(APP_NAME) -ldflags="-s -w"

local.snapshot: local.check ## Generates the artifact with 'go build'
	GOVERSION=${GOVERSION} BUILD_ENVIRONMENT=${BUILD_ENVIRONMENT} goreleaser build --snapshot --clean

local.push: local.build ## Pushes built artifact to specified location

local.run: local.build ## Builds the artifact and starts the service
	./${APP_NAME}

print_home:
	@echo ${ENVIRONMENT}


local.deploy: local.build ## Deploys locally built Helm plugin
	@mkdir -p ${HOME}/Library/helm/plugins/helm-kanvas-snapshot/bin
	@rm -rf ${HOME}/Library/helm/plugins/helm-kanvas-snapshot/bin/helm-kanvas-snapshot
	@cp helm-kanvas-snapshot ${HOME}/Library/helm/plugins/helm-kanvas-snapshot/bin/helm-kanvas-snapshot
	@cp plugin.yaml ${HOME}/Library/helm/plugins/helm-kanvas-snapshot/
	@cp install-binary.sh ${HOME}/Library/helm/plugins/helm-kanvas-snapshot/


publish: local.check ## Builds and publishes the app
	GOVERSION=${GOVERSION} BUILD_ENVIRONMENT=${BUILD_ENVIRONMENT} PLUGIN_PATH=${APP_DIR} goreleaser release --snapshot --clean

mock.publish: local.check ## Builds and mocks app release
	GOVERSION=${GOVERSION} BUILD_ENVIRONMENT=${BUILD_ENVIRONMENT} PLUGIN_PATH=${APP_DIR} goreleaser release --skip=publish --clean


install.hooks: ## Installs pre-push hooks for the repository
	${APP_DIR}/scripts/hook.sh ${APP_DIR}

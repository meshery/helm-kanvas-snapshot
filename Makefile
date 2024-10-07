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

include build/Makefile.core.mk


.PHONY: all
all: dep-check build
## Lint check 
golangci: error dep-check
	golangci-lint run --exclude-use-default

## Analyze error codes
error: dep-check
	go run github.com/layer5io/meshkit/cmd/errorutil -d . analyze -i ./build -o ./build

## Runs meshkit error utility to update error codes.
error-util:
	go run github.com/layer5io/meshkit/cmd/errorutil -d . update -i ./build -o ./build



#-----------------------------------------------------------------------------
# Dependencies
#-----------------------------------------------------------------------------
.PHONY: dep-check
#.SILENT: dep-check

INSTALLED_GO_VERSION=$(shell go version)

dep-check:

ifeq (,$(findstring $(GOVERSION), $(INSTALLED_GO_VERSION)))
# Only send a warning.
	@echo "Dependency missing: go$(GOVERSION). Ensure 'go$(GOVERSION).x' is installed and available in your 'PATH'"
	@echo "GOVERSION: " $(GOVERSION)
	@echo "INSTALLED_GO_VERSION: " $(INSTALLED_GO_VERSION)
endif


OUTDIR := ./cmd/kanvas-snapshot/bin
ARCH := amd64

BINNAME_DARWIN ?= kanvas-snapshot-darwin-$(ARCH)
BINNAME_LINUX ?= kanvas-snapshot-linux-$(ARCH)
BINNAME_WINDOWS ?= kanvas-snapshot-windows-$(ARCH).exe


LDFLAGS := "\
    -X 'main.MesheryCloudApiCookie=$(MESHERY_CLOUD_API_COOKIES)' \
    -X 'main.MesheryApiCookie=$(MESHERY_API_COOKIES)' \
    -X 'main.MesheryCloudApiBaseUrl=$(MESHERY_CLOUD_API_BASE_URL)' \
    -X 'main.MesheryApiBaseUrl=$(MESHERY_API_BASE_URL)' \
    -X 'main.SystemID=$(SYSTEM_ID)'"

.PHONY: build
build:
	@echo "Building for all platforms..."
	@$(MAKE) $(BINNAME_DARWIN)
	@$(MAKE) $(BINNAME_LINUX)
	@$(MAKE) $(BINNAME_WINDOWS)

# Build Helm plugin for Darwin (macOS)
.PHONY: $(BINNAME_DARWIN)
$(BINNAME_DARWIN):
	@echo "Building for Darwin..."
	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=darwin go build -ldflags=$(LDFLAGS) -o $(OUTDIR)/$(BINNAME_DARWIN) ./cmd/kanvas-snapshot/main.go

# Build Helm plugin for Linux
.PHONY: $(BINNAME_LINUX)
$(BINNAME_LINUX):
	@echo "Building for Linux..."
	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=linux go build -ldflags=$(LDFLAGS) -o $(OUTDIR)/$(BINNAME_LINUX) ./cmd/kanvas-snapshot/main.go

# Build Helm plugin for Windows
.PHONY: $(BINNAME_WINDOWS)
$(BINNAME_WINDOWS):
	@echo "Building for Windows..."
	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=windows go build -ldflags=$(LDFLAGS) -o $(OUTDIR)/$(BINNAME_WINDOWS) ./cmd/kanvas-snapshot/main.go

# Clean up binaries
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(OUTDIR)
